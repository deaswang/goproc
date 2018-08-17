package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/deaswang/goproc/proc"
)

// Processes get all processes base info
func Processes(w http.ResponseWriter, r *http.Request) {
	pid, err := proc.GetProcesses()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(pid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// Process get process detail info by pid
func Process(w http.ResponseWriter, r *http.Request, pid int) {

}

