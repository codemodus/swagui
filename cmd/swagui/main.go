package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/codemodus/swagui"
)

func main() {
	var port, path string
	var version int

	flag.StringVar(&port, "port", ":2288", "http port")
	flag.StringVar(&path, "path", "", "path prefix")
	flag.IntVar(&version, "v", 0, "swagger version")
	flag.Parse()

	opts := &swagui.Options{
		PathPrefix: path,
		Version:    version,
	}

	ui, err := swagui.New(opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err = http.ListenAndServe(port, ui.Handler()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
