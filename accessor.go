package swagui

import (
	"github.com/codemodus/swagui/bindata1"
	"github.com/codemodus/swagui/bindata2"
	"github.com/codemodus/swagui/bindata3"
)

type accessor interface {
	Access(string) ([]byte, error)
}

type a1 struct{}

func (a *a1) Access(name string) ([]byte, error) {
	return bindata1.Asset(name)
}

type a2 struct{}

func (a *a2) Access(name string) ([]byte, error) {
	if name == "swagger-ui.js" {
		name = "swagger-ui.min.js"
	}

	return bindata2.Asset(name)
}

type a3 struct{}

func (a *a3) Access(name string) ([]byte, error) {
	return bindata3.Asset(name)
}

func accessorByVersion(v Version) accessor {
	switch v {
	case V1:
		return &a1{}
	case V2:
		return &a2{}
	default:
		return &a3{}
	}
}
