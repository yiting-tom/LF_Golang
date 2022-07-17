package files

import "io"

// Storage defines the interface for file operations.
type Storage interface {
	Save(path string, contents io.Reader) error
}
