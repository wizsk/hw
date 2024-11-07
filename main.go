package main

import (
	"embed"
	"io"
	"log"
	"net/http"
	"text/template"

	"zombiezen.com/go/sqlite"
)

const (
	dbPath        = "assets/hw.db"
	indexPageFile = "index.html"
	// resPageFile   = "res.html"
	resPageFile = "index.html"
	debug       = true
)

var (
	//go:embed assets/pub/*
	pubDir embed.FS
)

type ResData struct {
	Word     string
	Entries  Entries
	IsRes    bool   // is result page
	PreInVal string // previous input value
}

func main() {
	conn, err := sqlite.OpenConn(dbPath, sqlite.OpenReadOnly)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var tmpl templateWraper
	if debug {
		tmpl = &tmplW{}
	} else {
		tmpl, err = newTmpl()
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
	t, err := newTmpl()
	if err != nil {
		return err
	}
	return t.ExecuteTemplate(w, name, data)
}

func newTmpl() (templateWraper, error) {
	return template.ParseGlob("ui/src/*")
}
