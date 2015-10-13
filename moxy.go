package moxy

import (
	"net/http"
	"strings"
)

// NewReverseProxy returns a new ReverseProxy that load-balances the proxy requests between multiple hosts
// It also allows to define a chain of filter functions to process the outgoing response(s)
func NewReverseProxy(hosts []string, filters []FilterFunc, pathToAppend string) *ReverseProxy {
	director := func(request *http.Request) {
		host, _ := pick(hosts)
		actualPath := "/" + strings.Split(request.URL.Path, pathToAppend)[1]

		request.URL.Scheme = "https"
		request.URL.Host = host
		request.URL.Path = actualPath
	}
	return &ReverseProxy{
		Transport: NewTransport(),
		Director:  director,
		Filters:   filters,
	}
}
