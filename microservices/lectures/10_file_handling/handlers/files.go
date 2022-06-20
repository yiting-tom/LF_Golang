package handlers

import (
	"10_file_handling/files"

	log "github.com/sirupsen/logrus"
)

// Files is a handler for r/w files.
type Files struct {
	l     *log.Logger
	store files.Storage
}

// NewFiles creates a new Files handler.
func NewFiles(l *log.Logger, s files.Storage) *Files {
	return &Files{l, s}
}
