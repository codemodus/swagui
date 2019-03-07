package swagui

import (
	"bufio"
	"bytes"

	"github.com/codemodus/swagui/bindata1"
	"github.com/codemodus/swagui/bindata2"
	"github.com/codemodus/swagui/bindata3"
)

// Version wraps integer to ease the defining of the swagger ui version.
type Version int

// V{n} constants list the available swagger ui versions.
const (
	VLatest Version = iota
	V1
	V2
	V3
)

type assetFunc func(string) ([]byte, error)

type uiFiles struct {
	aliases    map[string]string
	assetFn    assetFunc
	cache      map[string][]byte
	defaultDef string
}

func newUIFiles(v Version) *uiFiles {
	fs := &uiFiles{
		aliases: make(map[string]string),
		cache:   make(map[string][]byte),
	}

	switch v {
	case V1:
		fs.aliases["swagger-ui.js"] = "swagger-ui.min.js"
		fs.assetFn = bindata1.Asset
		fs.defaultDef = "http://petstore.swagger.wordnik.com/api/api-docs.json"

	case V2:
		fs.aliases["swagger-ui.js"] = "swagger-ui.min.js"
		fs.assetFn = bindata2.Asset
		fs.defaultDef = "http://petstore.swagger.io/v2/swagger.json"

	default:
		fs.assetFn = bindata3.Asset
		fs.defaultDef = "https://petstore.swagger.io/v2/swagger.json"
	}

	return fs
}

func (fs *uiFiles) file(name string) ([]byte, error) {
	return fs.assetFn(valueOrKey(fs.aliases, name))
}

func (fs *uiFiles) defFilteredFile(name, definition string) ([]byte, error) {
	bs, ok := fs.cache[definition]
	if ok {
		return bs, nil
	}

	bs, err := fs.file(name)
	if err != nil {
		return nil, err
	}

	bs, err = replace(bs, fs.defaultDef, definition)
	if err != nil {
		return nil, err
	}

	fs.cache[definition] = bs

	return bs, nil
}

func valueOrKey(m map[string]string, key string) string {
	if val, ok := m[key]; ok {
		return val
	}

	return key
}

func replace(s []byte, a, b string) ([]byte, error) {
	abs, bbs := []byte(a), []byte(b)
	dif := len(bbs) - len(abs)
	bs := make([]byte, 0, len(s)+dif)

	sc := bufio.NewScanner(bytes.NewReader(s))
	for sc.Scan() {
		ss := sc.Bytes()
		rbs := bytes.Replace(ss, abs, bbs, 1)
		bs = append(bs, rbs...)
		bs = append(bs, '\n')
	}

	return bs, sc.Err()
}
