package main

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type fileSystem interface {
	Open(name string) (file, error)
	Stat(name string) (os.FileInfo, error)
	Walk(root string, walkFn filepath.WalkFunc) error
	ReadFile(filename string) ([]byte, error)
	WriteFile(filename string, data []byte, perm os.FileMode) error
}

type file interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	Stat() (os.FileInfo, error)
}

// osFS implements fileSystem using the local disk.
type osFS struct{}

func (osFS) Open(name string) (file, error)                   { return os.Open(name) }
func (osFS) Stat(name string) (os.FileInfo, error)            { return os.Stat(name) }
func (osFS) Walk(root string, walkFn filepath.WalkFunc) error { return filepath.Walk(root, walkFn) }
func (osFS) ReadFile(filename string) ([]byte, error)         { return ioutil.ReadFile(filename) }
func (osFS) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filename, data, perm)
}

type mockFS struct{}

func (mockFS) Open(name string) (file, error)                                 { return nil, nil }
func (mockFS) Stat(name string) (os.FileInfo, error)                          { return nil, nil }
func (mockFS) Walk(root string, walkFn filepath.WalkFunc) error               { return nil }
func (mockFS) ReadFile(filename string) ([]byte, error)                       { return []byte(`Test String`), nil }
func (mockFS) WriteFile(filename string, data []byte, perm os.FileMode) error { return nil }
