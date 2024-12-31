package dispatcher

import (
	"net/http"
)

type Router struct {
	Routes map[string]map[string]http.HandlerFunc // path -> method -> handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if methodHandlers, ok := r.Routes[req.URL.Path]; ok {
		if handler, ok := methodHandlers[req.Method]; ok {
			handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}

// addRoute is a helper to add a new route
func (r *Router) AddRoute(method, pattern string, handler http.HandlerFunc) {
	if r.Routes == nil {
		r.Routes = make(map[string]map[string]http.HandlerFunc)
	}
	if r.Routes[pattern] == nil {
		r.Routes[pattern] = make(map[string]http.HandlerFunc)
	}
	r.Routes[pattern][method] = handler
}
