package files

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Local is an implementation of the Storage interface which
// works with the local disk.
type Local struct {
	basePath    string
	maxFileSize int
}

// fullPath returns the full path for the given path.
func (l *Local) fullPath(path string) string {
	return filepath.Join(l.basePath, path)
}

// NewLocal creates a new Local filesystem with the given base path.
// maxSize is the maximum size of a file in bytes.
// basePath is the base directory to save files to.
func NewLocal(basePath string, maxSize int) (*Local, error) {
	// Get the absolute path of the base path.
	p, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}

	return &Local{basePath: p}, nil
}

// Save the contes of the reader to the given path.
// path is a relative path to the base path.
func (l *Local) Save(path string, contents io.Reader) error {
	// Get the full path for the file
	fp := l.fullPath(path)

	// Get the dir and make sure it exists
	d := filepath.Dir(fp)
	if err := os.MkdirAll(d, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", d, err)
	}

	// Delete if it exists
	if _, err := os.Stat(fp); err == nil {
		if err = os.Remove(fp); err != nil {
			return fmt.Errorf("failed to remove file %s: %w", fp, err)
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("failed to stat file %s: %w", fp, err)
	}

	// Create the file
	f, err := os.Create(fp)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", fp, err)
	}
	defer f.Close()

	// Copy the contents
	if _, err = io.Copy(f, contents); err != nil {
		return fmt.Errorf("failed to copy contents to file %s: %w", fp, err)
	}

	return nil
}

// Get the file and return a Reader
// the calling function is responsible for closing the Reader.
func (l *Local) Get(path string) (*os.File, error) {
	// Get the full path for the file
	fp := l.fullPath(path)

	// Open the file
	f, err := os.Open(fp)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", fp, err)
	}

	return f, nil
}
