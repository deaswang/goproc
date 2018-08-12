package main

import (
	"net/http"

	"github.com/deaswang/goproc/handlers"
)

// single router struct
type Route struct {
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// define router list
var routes = Routes{
	Route{
		"/",
		index,
	},
	Route{
		"/cpuinfo",
		handlers.Cpuinfo,
	},
	Route{
		"/cpu",
		index,
	},
	Route{
		"/todos/",
		index,
	},
}

// init router
func InitRoute(need_token bool) http.Handler {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.HandleFunc(route.Pattern, withAuth(need_token, route.HandlerFunc))
	}
	return mux
}
