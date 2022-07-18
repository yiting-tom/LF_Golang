package files

import (
	"io"
	"os"
)

// Storage defines the interface for file operations.
type Storage interface {
	Save(path string, contents io.Reader) error
	Get(path string) (*os.File, error)
}
