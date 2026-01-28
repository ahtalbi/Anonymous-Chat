package handlers

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("web"))
	fs.ServeHTTP(w, r)
}
