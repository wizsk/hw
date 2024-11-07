package main

import (
	"io"
	"log"
	"net/http"
	"text/template"

	"zombiezen.com/go/sqlite"
)

const (
	dbPath        = "assets/hw.db"
	indexPageFile = "index.html"
	resPageFile   = "res.html"
	debug         = true
)

type ResData struct {
	Word    string
	Entries Entries
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
		err := tmpl.ExecuteTemplate(w, indexPageFile, nil)
		if err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/r", func(wt http.ResponseWriter, r *http.Request) {
		w := r.FormValue("w")
		e, _ := searchByRoot(conn, w)
		d := ResData{w, e}
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
		d := ResData{w, e}
		if err := tmpl.ExecuteTemplate(wt, resPageFile, &d); err != nil {
			log.Fatal(err)
		}
	})

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
