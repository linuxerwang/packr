package packr

import (
	"encoding/json"
	"path/filepath"
	"sync"
)

var gil = &sync.Mutex{}
var data = map[string][]byte{}

// PackBytes packs bytes for a file into a box.
func PackBytes(box string, name string, bb []byte) {
	gil.Lock()
	defer gil.Unlock()
	data[filepath.Join(box, name)] = bb
}

// PackJSONBytes packs JSON encoded bytes for a file into a box.
func PackJSONBytes(box string, name string, jbb string) error {
	bb := []byte{}
	err := json.Unmarshal([]byte(jbb), &bb)
	if err != nil {
		return err
	}
	PackBytes(box, name, bb)
	return nil
}
