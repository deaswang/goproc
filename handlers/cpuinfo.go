package handlers

import (
	"net/http"
	"encoding/json"
	"strings"

	"github.com/deaswang/goproc/proc"
)

// Cpuinfo handle the cpu info file
func Cpuinfo(w http.ResponseWriter, r *http.Request) {
	cpuinfo, err := proc.GetCPUInfo("")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	r.ParseForm()
	value, ok := r.Form["data"]
	if ok {
		for _, v := range value {
			if strings.ToLower(v) == "mhz" {
				var mhz = make(map[int]float64)
				for _, p := range cpuinfo.Processors {
					mhz[p.ID] = p.MHz
				}
				b, err := json.Marshal(mhz)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.WriteHeader(http.StatusOK)
				w.Write(b)
				return
			}
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		b, err := json.Marshal(cpuinfo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
