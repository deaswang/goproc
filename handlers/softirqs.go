package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/deaswang/goproc/proc"
)

// Softirqs handle the spftirqs file
func Softirqs(w http.ResponseWriter, r *http.Request) {
	softirqs, err := proc.GetSoftirqs("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(softirqs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
