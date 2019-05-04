package docsite

import (
	"io/ioutil"
	"net/http"
)

func ReadFile(fs http.FileSystem, path string) ([]byte, error) {
	f, err := fs.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}
