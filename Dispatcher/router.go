package dispatcher

import (
	"net/http"
	"synapse/consolelogger"
)

type Route struct {
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

type Router struct {
	Routes []Route
}

// ServeHTTP makes Router implement the http.Handler interface.
// It finds a matching route for the incoming request, then calls the handler.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method

	// Try to match a route
	for _, route := range r.Routes {
		if route.Method == method && route.Pattern == path {
			route.Handler(w, req)
			return
		}
	}

	// If no route was matched, return 404
	http.NotFound(w, req)
}

// addRoute is a helper to add a new route
func (r *Router) AddRoute(method, pattern string, handler http.HandlerFunc) {
	route := Route{
		Method:  method,
		Pattern: pattern,
		Handler: handler,
	}
	consolelogger.DebugLog("Adding route: " + method + " " + pattern)
	r.Routes = append(r.Routes, route)
}
