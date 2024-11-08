package main

import (
	"embed"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"zombiezen.com/go/sqlite"
)

const (
	dbName        = "hw.db"
	dbPath        = "assets/" + dbName
	indexPageFile = "index.html"
	// resPageFile   = "res.html"
	resPageFile = "index.html"
	debug       = true
)

var (
	//go:embed assets/pub/*
	pubDir embed.FS
	//go:embed assets/hw.db
	dataBase embed.FS

	//go:embed ui/src/*
	uiTmpls embed.FS
)

type ResData struct {
	Word     string
	Entries  Entries
	IsRes    bool   // is result page
	PreInVal string // previous input value
}

func main() {
	dbTmpPath := filepath.Join(os.TempDir(), dbName)
	{
		d, err := dataBase.Open(dbPath)
		if err != nil {
			panic(err)
		}
		dt, err := os.Create(dbTmpPath)
		if err != nil {
			panic(err)
		}
		if _, err := io.Copy(dt, d); err != nil {
			panic(err)
		}
		d.Close()
		dt.Close()
	}

	conn, err := sqlite.OpenConn(dbTmpPath, sqlite.OpenReadOnly)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var tmpl templateWraper
	if debug {
		tmpl = &tmplW{}
	} else {
		tmpl, err = openEmbdedTmpl()
		if err != nil {
			log.Fatal(err)
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, indexPageFile, ResData{IsRes: false})
		if err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/r", func(wt http.ResponseWriter, r *http.Request) {
		w := r.FormValue("w")
		e, _ := searchByRoot(conn, w)
		d := ResData{w, e, true, w}
		if err := tmpl.ExecuteTemplate(wt, resPageFile, &d); err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/t", func(wt http.ResponseWriter, r *http.Request) {
		w := r.FormValue("w")
		e, err := searchByTxt(conn, w)
		if err != nil {
			log.Fatal(err)
		}
		d := ResData{w, e, true, w}
		if err := tmpl.ExecuteTemplate(wt, resPageFile, &d); err != nil {
			log.Fatal(err)
		}
	})

	http.Handle("/assets/", http.FileServerFS(pubDir))
	panic(http.ListenAndServe(":8001", nil))
}

type templateWraper interface {
	ExecuteTemplate(wr io.Writer, name string, data any) error
}

type tmplW struct{}

func (tp *tmplW) ExecuteTemplate(w io.Writer, name string, data any) error {
	t, err := template.ParseGlob("ui/src/*")
	if err != nil {
		return err
	}
	return t.ExecuteTemplate(w, name, data)
}

func openEmbdedTmpl() (templateWraper, error) {
	return template.ParseFS(uiTmpls, "ui/src/*")

}
