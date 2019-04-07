package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/codemodus/swagui"
	"github.com/codemodus/swagui/suidata1"
	"github.com/codemodus/swagui/suidata2"
	"github.com/codemodus/swagui/suidata3"
)

func main() {
	var (
		port    = ":2288"
		defdef  = ""
		version = 0
	)

	flag.StringVar(&port, "port", port, "http port")
	flag.StringVar(&defdef, "def", defdef, "default definition")
	flag.IntVar(&version, "v", version, "swagger version")
	flag.Parse()

	var p swagui.Provider
	switch version {
	case 1:
		p = suidata1.New()
	case 2:
		p = suidata2.New()
	default:
		version = 3
		p = suidata3.New()
	}
	ui, err := swagui.New(http.NotFoundHandler(), p)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// TODO: add http caching, etc.

	msgfmt := "serving swaggerui v%d on %s, with default def %q"
	fmt.Printf(msgfmt, version, port, defdef)

	if err = http.ListenAndServe(port, ui.Handler(defdef)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
