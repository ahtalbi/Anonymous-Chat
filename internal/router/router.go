package router

import "net/http"

type Router struct {
	Mux *http.ServeMux
}

func New(mux *http.ServeMux) *Router {
	return &Router{
		Mux: mux,
	}
}

func (r *Router) Routes(routes map[string]func(w http.ResponseWriter, r *http.Request)) {
	for route, handler := range routes {
		r.Mux.HandleFunc(route, handler)
	}
}