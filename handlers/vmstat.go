package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/deaswang/goproc/proc"
)

// Vmstat handle the vmstat file
func Vmstat(w http.ResponseWriter, r *http.Request) {
	vmstat, err := proc.GetVmStat("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(vmstat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
