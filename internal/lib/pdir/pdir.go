package pdir

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// RmRf remove o diretório e seu conteúdo recursivamente
func RmRf(path string) error {
	p, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, d := range p {
		os.RemoveAll(filepath.Join(path, d.Name()))
	}
	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}
