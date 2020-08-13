package fs

import (
	"github.com/spf13/afero"
	"log"
	"os"
	"path/filepath"
)

var initializedFs *AbstractFS

type AbstractFS struct {
	f afero.Fs
}

func (fs *AbstractFS) init(dir string) {
	aferofs := afero.NewOsFs()
	if dir != "" {
		aferofs = afero.NewBasePathFs(aferofs, dir)
	}
	fs.f = aferofs
}

func (fs *AbstractFS) WriteFile(path, data string, force bool) error {
	if b, _ := fs.Exists(path); b {
		s, err := fs.ReadFile(path)
		if err != nil {
			return err
		}
		if s == data {
			log.Printf("%s already exists. Skip.", path)
			return nil
		}
	}
	d, _ := filepath.Split(path)

	if err := fs.MkdirAll(d); err != nil {
		return err
	}

	return afero.WriteFile(fs.f, path, []byte(data), os.ModePerm)
}

func (fs *AbstractFS) ReadFile(path string) (string, error) {
	b, err := afero.ReadFile(fs.f, path)
	return string(b), err
}

func (fs *AbstractFS) Exists(path string) (bool, error) {
	return afero.Exists(fs.f, path)
}

func (fs *AbstractFS) Mkdir(dir string) error {
	return fs.f.Mkdir(dir, os.ModePerm)
}

func (fs *AbstractFS) MkdirAll(path string) error {
	return fs.f.MkdirAll(path, os.ModePerm)
}

func NewAbstractFS(dir string) *AbstractFS {
	fs := &AbstractFS{}
	fs.init(dir)
	initializedFs = fs
	return fs
}

func Get() *AbstractFS {
	if initializedFs == nil {
		return NewAbstractFS("")
	}

	return initializedFs
}

