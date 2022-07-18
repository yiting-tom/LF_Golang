package handlers

import (
	"io"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

func (f *Files) GetSingleImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.l.Info("Handle <GET> received request for file: ", id, fn)

	fp := filepath.Join(id, fn)
	fio, err := f.store.Get(fp)

	if err != nil {
		f.l.Error("Failed to get file: ", err)
		http.Error(w, "Failed to get file", http.StatusInternalServerError)
		return
	}

	// Copy the steam to the response body
	io.Copy(w, fio)

}
