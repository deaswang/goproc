package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/deaswang/goproc/proc"
)

// Loadavg handle the loadavg file
func Loadavg(w http.ResponseWriter, r *http.Request) {
	loadavg, err := proc.GetLoadAvg("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(loadavg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
