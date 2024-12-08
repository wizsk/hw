package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"text/template"

	"github.com/wizsk/hw/db"
)

type ResData struct {
	Word        string
	Entries     db.HEntries
	IsRes       bool   // is result page
	IsRootPage  bool   // is result page
	PreInVal    string // previous input value
	Suggestions []string
}

func unkownCmd(c string) {
	fmt.Printf("Unkown command: %q\n", c)
	printUsagesAndExit()
}

func printVersionAndExit() {
	fmt.Printf("%s version v%s %s/%s\n", progName, version, runtime.GOOS, runtime.GOARCH)
	os.Exit(0)
}

func printUsagesAndExit() {
	fmt.Print(usages)
	os.Exit(0)
}

func parseAragsAndFlags() {
	for _, v := range os.Args[1:] {
		switch v {
		case "help", "--help", "-help", "-h", "--h":
			printUsagesAndExit()

		case "nb", "nobrowser":
			willOpenBrowser = false

		case "version":
			printVersionAndExit()

		default:
			if len(v) == 4 {
				if _, err := strconv.Atoi(v); err == nil {
					port = v
					continue
				}
			}
			unkownCmd(v)
		}
	}
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

// only if gloabal var 'willOpenBrowser == true' then open
func openBrower(url string) {
	if !willOpenBrowser || debug {
		return
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "linux":
		// not running graphycally
		session := os.Getenv("XDG_SESSION_TYPE")

		if !(session == "wayland" || session == "x11") {
			fmt.Println("[WARNINING] Your not running x11 or wayland. No browsers will be opened")
			return
		}
		if exec.Command("command", "-v", "xdg-open").Run() != nil {
			fmt.Println("[WARNINING] xdg-open command not found")
			return
		}
		cmd = exec.Command("xdg-open", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		return
	}
	addAtrribute(cmd)
	cmd.Start()
	fmt.Println("Holdon... opening on your brwoser")
}
