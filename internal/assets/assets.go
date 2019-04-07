package assets

import (
	"bufio"
	"bytes"
)

// AssetFunc returns bytes by file name.
type AssetFunc func(string) ([]byte, error)

// Assets wraps an accessor function, configured aliases, and a cache. Access
// is provided to assets along with access with data modification.
type Assets struct {
	assetFn AssetFunc
	aliases map[string]string
	cache   map[string][]byte
}

// New sets up Assets.
func New(fn AssetFunc, aliases map[string]string) *Assets {
	return &Assets{
		assetFn: fn,
		aliases: aliases,
		cache:   make(map[string][]byte),
	}
}

// Asset accesses assets using available aliases.
func (as *Assets) Asset(name string) ([]byte, error) {
	return as.assetFn(valueOrKey(as.aliases, name))
}

// ModifiedAsset accesses and modifies (a to b) assets using available aliases.
func (as *Assets) ModifiedAsset(name, a, b string) ([]byte, error) {
	key := name + a + b

	bs, ok := as.cache[key]
	if ok {
		return bs, nil
	}

	bs, err := as.Asset(name)
	if err != nil {
		return nil, err
	}

	bs, err = replace(bs, a, b)
	if err != nil {
		return nil, err
	}

	as.cache[key] = bs

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
