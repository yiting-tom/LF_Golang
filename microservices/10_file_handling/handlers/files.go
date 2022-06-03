package handlers

import (
	"10_file_handling/files"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
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

// ServeHTTP implements the http.Handler interface.
func (f *Files) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.l.Info("<POST> received request for file: ", id, fn)

	f.saveFile(id, fn, w, r)
}

// saveFile saves the contents of the request body to the file.
func (f *Files) saveFile(id, path string, w http.ResponseWriter, r *http.Request) {
	f.l.Info("Save file: ", id, path)

	fp := filepath.Join(id, path)
	if err := f.store.Save(fp, r.Body); err != nil {
		f.l.Error("Failed to save file: ", err)
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
	}
}
