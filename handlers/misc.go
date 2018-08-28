package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/deaswang/goproc/proc"
)

// Misc handle the misc file
func Misc(w http.ResponseWriter, r *http.Request) {
	misc, err := proc.GetMisc("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(misc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
