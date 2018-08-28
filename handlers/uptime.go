package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/deaswang/goproc/proc"
)

// Uptime handle the uptime file
func Uptime(w http.ResponseWriter, r *http.Request) {
	uptime, err := proc.GetUptime("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(uptime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
