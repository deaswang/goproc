package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/deaswang/goproc/proc"
)

// Locks handle the locks file
func Locks(w http.ResponseWriter, r *http.Request) {
	locks, err := proc.GetLocks("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(locks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
