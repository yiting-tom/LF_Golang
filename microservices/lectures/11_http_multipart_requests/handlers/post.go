package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

// Create a new file with
func (f *Files) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	fp := filepath.Join(id, fn)
	f.l.Info("Handle <POST> received request for file: ", fp)

	f.saveFile(id, fn, w, r)
}

// saveFile saves the contents of the request body to the file.
func (f *Files) saveFile(id, path string, w http.ResponseWriter, r *http.Request) {

	fp := filepath.Join(id, path)
	f.l.Info("Save file: ", fp)

	if err := f.store.Save(fp, r.Body); err != nil {
		f.l.Error("Failed to save file: ", err)
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("File saved"))
}
