package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
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

// ServeHTTP handles the HTTP request.
func (f *Files) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.l.Info("serving file", "id", id, "filename", fn)

	f.saveFile(id, fn, w, r)
}

// saveFile saves the file to the store.
func (f *Files) saveFile(id, path string, w http.ResponseWriter, r *http.Request) {
	f.l.Info("saving file", "id", id, "filename", path)

	fp := filepath.Join(id, path)
	if err := f.store.Save(fp, r.Body); err != nil {
		f.l.Error("error saving file", "id", id, "filename", path, "err", err)
		http.Error(w, "error saving file", http.StatusInternalServerError)
		return
	}
}
