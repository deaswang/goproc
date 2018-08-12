package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/deaswang/goproc/proc"
)

// Cpuinfo handle the cpu info file
func Cpuinfo(w http.ResponseWriter, r *http.Request) {
	cpuinfo, err := proc.GetCPUInfo("")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(cpuinfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

