package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/deaswang/goproc/proc"
)

// Vmstat handle the vmstat file
func Vmstat(w http.ResponseWriter, r *http.Request) {
	vmstat, err := proc.GetVMStat("")
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
