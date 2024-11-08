package main

import (
	"embed"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"text/template"
	"time"

	"zombiezen.com/go/sqlite"
)

const (
	dbName        = "hw.db"
	dbPath        = "assets/" + dbName
	indexPageFile = "index.html"
	// resPageFile   = "res.html"
	resPageFile = "index.html"
	debug       = !true
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
		if debug && err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/r", func(wt http.ResponseWriter, r *http.Request) {
		w := r.FormValue("w")
		e, _ := searchByRoot(conn, w)
		d := ResData{w, e, true, w}
		if err := tmpl.ExecuteTemplate(wt, resPageFile, &d); debug && err != nil {
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
		if err := tmpl.ExecuteTemplate(wt, resPageFile, &d); debug && err != nil {
			log.Fatal(err)
		}
	})
	http.Handle("/assets/", http.FileServerFS(pubDir))

	port := "8001"
	if len(os.Args) == 2 {
		port = os.Args[1]
	}

	fmt.Println("Running:")
	fmt.Println("Localy: http://localhost:" + port)

	if runtime.GOOS == "linux" {
		fmt.Printf("Internet?: http://%s:%s\n", localIp(), port)
	} else {
		fmt.Println("Find your ip please :D")
	}

	run := make(chan bool)

	go func(c chan<- bool) {
		if err = http.ListenAndServe(":"+port, nil); err != nil {
			fmt.Println("encountered err:", err)
			fmt.Printf("It may be that the port %q is already used (defaut: 8001)\n", port)
			fmt.Println("You can specify port number by progname follwed by the port nubmer")
			fmt.Println("Example: `hw 8080`")
			c <- false
		}
	}(run)

	fmt.Println("Holdon... opening on your brwoser")

	go func(c chan<- bool) {
		time.Sleep(2 * time.Second)
		c <- true
	}(run)

	runn := <-run
	if !runn {
		return
	}

	openBrower("http://localhost:" + port)
	select {}
}

func openBrower(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		return
	}
	cmd.Run()
}

func localIp() string {
	if runtime.GOOS == "windows" {
		return "localhost"
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "localhost"
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}
	return "localhost"
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
