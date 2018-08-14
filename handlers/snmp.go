package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/deaswang/goproc/proc"
)

// Snmp handle the snmp file
func Snmp(w http.ResponseWriter, r *http.Request) {
	snmp, err := proc.GetSnmp("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(snmp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
