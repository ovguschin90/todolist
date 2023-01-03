package router

import (
	"net/http"
	"regexp"
	"strings"
)

type Route struct {
	method  string
	pattern *regexp.Regexp
	handler http.HandlerFunc
}

type Router struct {
	routes []*Route
}

func (rt *Router) AddRoute(method, pattern string, handler http.HandlerFunc) {
	pattern = strings.Join([]string{"^", pattern, "$"}, "")
	route := &Route{
		method:  method,
		pattern: regexp.MustCompile(pattern),
		handler: handler,
	}
	rt.routes = append(rt.routes, route)
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	method := r.Method

	for _, route := range rt.routes {
		if route.pattern.MatchString(path) {
			if method != route.method {
				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
				return
			}
			route.handler.ServeHTTP(w, r)
			return
		}
	}

	http.NotFound(w, r)
}
