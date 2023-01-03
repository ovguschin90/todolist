package router

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

type routeNode struct {
	method   string
	pattern  *regexp.Regexp
	handler  http.HandlerFunc
	children map[rune]*routeNode
	isLeaf   bool
}

type Router struct {
	root *routeNode
}

func New() *Router {
	return &Router{root: &routeNode{children: make(map[rune]*routeNode)}}
}

func (rt *Router) AddRoute(method, pattern string, handler http.HandlerFunc) {
	currentNode := rt.insert(pattern)
	currentNode.pattern = regexp.MustCompile(wrapPattern(pattern))
	currentNode.method = method
	currentNode.handler = handler
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	path := r.URL.Path
	route := rt.search(path)

	if route == nil {
		http.NotFound(w, r)
		return
	}

	if route.pattern != nil {
		ok := route.pattern.MatchString(path)
		if !ok {
			http.NotFound(w, r)
			return
		}

		if route.method != method {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		route.handler.ServeHTTP(w, r)
		return
	}
	http.NotFound(w, r)
}

func (rt *Router) RoutesList() {
	rt.walk(rt.root, "")
}

func (rt *Router) insert(route string) *routeNode {
	currentNode := rt.root
	for _, c := range route {
		if currentNode.children[c] == nil {
			currentNode.children[c] = &routeNode{children: make(map[rune]*routeNode)}
		}
		currentNode = currentNode.children[c]
	}

	currentNode.isLeaf = true

	return currentNode
}

func (rt *Router) search(route string) *routeNode {
	currentNode := rt.root

	for _, c := range route {
		if currentNode.children[c] == nil {
			return nil
		}
		currentNode = currentNode.children[c]
	}

	return currentNode
}

func (rt *Router) walk(node *routeNode, prefix string) {
	if node.isLeaf {
		logrus.Info(strings.Join([]string{node.method, unwrapPattern(node.pattern.String())}, " "))
	}

	for c, child := range node.children {
		rt.walk(child, prefix+string(c))
	}
}

func wrapPattern(pattern string) string {
	return fmt.Sprintf("^%s$", pattern)
}

func unwrapPattern(pattern string) string {
	return strings.Trim(pattern, "^$")
}
