//+build dev

package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewAssetsHandler() http.Handler {
	url, _ := url.Parse("http://127.0.0.1:3000")
	return httputil.NewSingleHostReverseProxy(url)
}
