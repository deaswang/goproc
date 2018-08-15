package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/deaswang/goproc/handlers"
)

// Router is custom router
type Router struct {
	http.ServeMux
}

// define router list
var routerList = map[string]http.HandlerFunc{
	"/":           index,
	"/cpuinfo":    handlers.Cpuinfo,
	"/buddyinfo":  handlers.Buddyinfo,
	"/diskstats":  handlers.Diskstats,
	"/interrupts": handlers.Interrupts,
	"/loadavg":    handlers.Loadavg,
	"/locks":      handlers.Locks,
	"/meminfo":    handlers.Meminfo,
	"/misc":       handlers.Misc,
	"/mounts":     handlers.Mounts,
	"/partitions": handlers.Partitions,
	"/snmp":       handlers.Snmp,
	"/softirqs":   handlers.Softirqs,
	"/stat":       handlers.Stat,
	"/uptime":     handlers.Uptime,
	"/version":    handlers.Version,
}

// ServeHTTP custom route
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri := r.RequestURI
	uris := strings.Split(uri, "/")
	for _, u := range uris {
		if u == "" {
			continue
		}
		pid, err := strconv.Atoi(u)
		if err != nil {
			break
		} else {
			handlers.Process(w, r, pid)
			return
		}
	}
	router.ServeMux.ServeHTTP(w, r)
}

// InitRouter init router
func InitRouter(needToken bool) http.Handler {
	router := new(Router)
	for pattern, handlerFunc := range routerList {
		router.HandleFunc(pattern, withAuth(needToken, handlerFunc))
	}
	return router
}
