package handlers

import (
	"net/http"
	"os"
	"path/filepath"
)

var fs = http.FileServer(http.Dir("web"))

func Home(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("web", filepath.Clean(r.URL.Path))

	if info, err := os.Stat(path); err == nil && !info.IsDir() {
		http.ServeFile(w, r, path)
		return
	}

	http.ServeFile(w, r, "web/index.html")
}
