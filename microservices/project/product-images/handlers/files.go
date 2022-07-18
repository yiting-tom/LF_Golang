package handlers

import (
	"github.com/sirupsen/logrus"
	"github.com/yiting-tom/LF_Golang/microservices/project/product-images/files"
)

// Files is a handler that provides file storage.
type Files struct {
	store files.Storage
	l     *logrus.Logger
}

// NewFiles returns a new Files handler.
func NewFiles(s files.Storage, l *logrus.Logger) *Files {
	return &Files{store: s, l: l}
}
