package files

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Local is an implementation of the File interface that stores files locally
type Local struct {
	maxFileSize int
	basePath    string
}

// NewLocal returns a new Local
func NewLocal(basePath string, maxSize int) (*Local, error) {
	p, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}

	return &Local{basePath: p, maxFileSize: maxSize}, nil
}

// Save saves the contents of the io.Reader to the given path
func (l *Local) Save(path string, contents io.Reader) error {
	// Get the full path
	p := l.fullPath(path)

	// Create the directory if it doesn't exist
	d := filepath.Dir(p)
	err := os.Mkdir(d, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("error creating directory %s: %s", d, err)
	}

	// If the file already exists, we need to remove it first
	switch _, err := os.Stat(p); err {
	case nil:
		if err := os.Remove(p); err != nil {
			return fmt.Errorf("error removing file %s: %s", p, err)
		}
	case os.ErrNotExist:
		return fmt.Errorf("error getting file info: %s", err)
	}

	// Create the file
	f, err := os.Create(p)
	if err != nil {
		return fmt.Errorf("error creating file %s: %s", p, err)
	}
	defer f.Close()

	// Write the contents
	if _, err = io.Copy(f, contents); err != nil {
		return fmt.Errorf("error copying contents: %s", err)
	}

	return nil
}

// Get returns the os.File for the given path
func (l *Local) Get(path string) (*os.File, error) {
	// Get the full path
	p := l.fullPath(path)

	// Open the file
	f, err := os.Open(p)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %s", p, err)
	}

	return f, nil
}

// fullPath returns the full path to the file
func (l *Local) fullPath(path string) string {
	return filepath.Join(l.basePath, path)
}
