package swagui

import (
	"github.com/codemodus/swagui/bindata1"
	"github.com/codemodus/swagui/bindata2"
)

type assetFinder interface {
	Asset(string) ([]byte, error)
}

type data1 struct{}

func (d *data1) Asset(name string) ([]byte, error) {
	return bindata1.Asset(name)
}

type data2 struct{}

func (d *data2) Asset(name string) ([]byte, error) {
	return bindata2.Asset(name)
}
