//+build !dev

package main

import (
	"net/http"
	"path"

	"github.com/go-chi/chi"
)

//go:generate go-assets-builder -s /client/dist -o assets_gen.go client/dist

func NewAssetsHandler() http.Handler {
	r := chi.NewRouter()
	h := http.FileServer(Assets)
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		p := path.Clean(r.URL.Path)
		if _, ok := Assets.Files[p]; !ok {
			r.URL.Path = "/"
		}
		h.ServeHTTP(w, r)
	})
	return r
}
