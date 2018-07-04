package swagui

import (
	"github.com/codemodus/swagui/bindata1"
	"github.com/codemodus/swagui/bindata2"
	"github.com/codemodus/swagui/bindata3"
)

// Version wraps integer to ease the defining of the swagger ui version.
type Version int

// V{n} constants list the available swagger ui versions.
const (
	VNewest Version = iota
	V1
	V2
	V3
)

var (
	indexFile = "index.html"
)

type assetFunc func(string) ([]byte, error)

type uiFiles struct {
	aliases map[string]string
	assetFn assetFunc
	store   map[string][]byte
	origDef string
}

func newUIFiles(v Version) *uiFiles {
	fs := &uiFiles{
		aliases: make(map[string]string),
		store:   make(map[string][]byte),
	}

	switch v {
	case V1:
		fs.aliases["swagger-ui.js"] = "swagger-ui.min.js"
		fs.assetFn = bindata1.Asset
		fs.origDef = "http://petstore.swagger.wordnik.com/api/api-docs.json"

	case V2:
		fs.aliases["swagger-ui.js"] = "swagger-ui.min.js"
		fs.assetFn = bindata2.Asset
		fs.origDef = "http://petstore.swagger.io/v2/swagger.json"

	default:
		fs.assetFn = bindata3.Asset
		fs.origDef = "https://petstore.swagger.io/v2/swagger.json"

	}

	return fs
}

func (fs *uiFiles) access(name string) ([]byte, error) {
	return fs.assetFn(resolveAlias(fs.aliases, name))
}

func (fs *uiFiles) accessIndex(definition string) ([]byte, error) {
	return accessIndex(fs.store, fs.assetFn, fs.origDef, definition)
}

func resolveAlias(aliases map[string]string, name string) string {
	if len(aliases) == 0 {
		return name
	}

	a, ok := aliases[name]
	if !ok {
		return name
	}

	return a
}

func accessIndex(store map[string][]byte, assetFn assetFunc, orig, def string) ([]byte, error) {
	bs, ok := store[def]
	if ok {
		return bs, nil
	}

	bs, err := assetFn(indexFile)
	if err != nil {
		return nil, err
	}

	bs, err = replaceDefinition(orig, def, bs)
	if err != nil {
		return nil, err
	}

	store[def] = bs

	return bs, nil
}

func replaceDefinition(orig, def string, bs []byte) ([]byte, error) {
	return bs, nil
}
