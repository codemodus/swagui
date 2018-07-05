package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/codemodus/swagui"
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

	var v swagui.Version
	switch version {
	case 1:
		v = swagui.V1
	case 2:
		v = swagui.V2
	default:
		v = swagui.V3
	}

	opts := &swagui.Options{Version: v}
	ui, err := swagui.New(opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err = http.ListenAndServe(port, ui.Handler(defdef)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
