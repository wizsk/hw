package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/wizsk/hw/db"
)

const (
	debug    = true
	progName = "hw"
	version  = "1.0"

	indexPageFile = "index.html"
	// rootExplainPageFile = "roots.html"
	rootExplainPageFile = "index.html"
	rootSuggtTmplFile   = "suggestion.html" // for root and text search
	// resPageFile   = "res.html"
	resPageFile = "index.html"
	defaultPort = "8080"

	ResultLimit    = 50 // for root and text search
	RootSuggtLimit = 6  // for root and text search
)

var (
	//go:embed assets/pub/*
	pubDir embed.FS

	//go:embed ui/src/*
	uiTmpls embed.FS

	// always open browser
	willOpenBrowser = true // always open browser

	// default port 8080
	port = defaultPort
)

const usages = progName + `: [port] [COMMANDS...]
PORT:
	Just the port number. (default: ` + defaultPort + `)

COMMANDS:
	nobrowser, nb
		don't open browser
	version
		print version number
`

func main() {
	parseAragsAndFlags()

	var err error
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
		if r.URL.Path != "/" {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}
		err := tmpl.ExecuteTemplate(w, indexPageFile, ResData{IsRes: false})
		if debug && err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/roots", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, rootExplainPageFile, ResData{IsRootPage: true})
		if debug && err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/r", func(wt http.ResponseWriter, r *http.Request) {
		w := r.FormValue("w")
		e := db.SearchByRoot(w, ResultLimit).HTML()
		d := ResData{w, e, true, false, w, nil}
		if err := tmpl.ExecuteTemplate(wt, resPageFile, &d); debug && err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/t", func(wt http.ResponseWriter, r *http.Request) {
		w := r.FormValue("w")
		e := db.SearchByTxt(w, ResultLimit, "").HTML()
		d := ResData{w, e, true, false, w, nil}
		if err := tmpl.ExecuteTemplate(wt, resPageFile, &d); debug && err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/sugg", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Post request only", http.StatusBadRequest)
			return
		}

		v := strings.TrimSpace(r.FormValue("w"))
		if v == "" {
			http.Error(w, "No words provided", http.StatusNotFound)
			return
		}
		s := db.RootSuggestion(v, RootSuggtLimit)
		if s == nil {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		if tmpl.ExecuteTemplate(w, rootSuggtTmplFile, ResData{Suggestions: s}) != nil {
			http.Error(w, "Something went wrong!", http.StatusInternalServerError)
			return
		}
	})

	http.Handle("/assets/", http.FileServerFS(pubDir))

	serveErr := make(chan struct{})

	curPort := 8080
	portTrylimit := 1
	if port != defaultPort {
		// err was already checked while parsing args
		curPort, _ = strconv.Atoi(port)
	} else {
		portTrylimit = 10
	}
	success := false

loop:
	for i := 0; i < portTrylimit; i++ {
		go func(c chan<- struct{}) {
			if err = http.ListenAndServe(fmt.Sprintf(":%d", curPort), nil); err != nil {
				c <- struct{}{}
			}
		}(serveErr)

		select {
		case <-serveErr:
			if port != defaultPort {
				success = false
				break loop
			}
			curPort++
			continue loop
		case <-time.Tick(2 * time.Second):
			success = true
			break loop
		}
	}

	if !success {
		fmt.Printf("Could not start the server at port %d\n", curPort)
		fmt.Printf("It may be that the port %d is already used (defaut: %s)\n", curPort, defaultPort)
		fmt.Println("You can specify port number by progname follwed by the port nubmer")
		fmt.Println("Example: `hw 8081`")
		os.Exit(1)
	}

	fmt.Println("Running:")
	fmt.Printf("Localy: http://localhost:%d\n", curPort)

	if runtime.GOOS == "linux" {
		fmt.Printf("Internet?: http://%s:%d\n", localIp(), curPort)
	} else {
		fmt.Println("Find your ip please.")
	}

	openBrower("http://localhost:" + port)
	select {}
}
