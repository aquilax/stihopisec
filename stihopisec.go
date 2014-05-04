package stihopisec

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"html/template"
	"net/http"
)

func init() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
		Funcs: []template.FuncMap{
			{
				"mod": func(n int, mod int) int {
					return n % mod
				},
			},
		},
	}))
	m.Use(martini.Static("public_html"))
	m.Get("/", GetHomeIndex)
	m.Get("/stih", GetStihove)
	m.Get("/stih/show/:id", GetStih)
	http.Handle("/", m)
}

func GetHomeIndex(r render.Render, req *http.Request) {
	r.HTML(200, "index", nil)
}

func GetStihove(r render.Render, req *http.Request) {
	r.HTML(200, "stih", nil)
}

func GetStih(r render.Render, req *http.Request) {
	r.HTML(200, "stihove", nil)
}
