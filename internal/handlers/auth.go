package handlers

import (
	"chat/internal/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	bytes := make([]byte, 1024)
	_, err := io.Reader.Read(r.Body, bytes)
	if err != nil {
		http.Error(w, "error in reading the request", http.StatusInternalServerError)
		return
	}
	var client *models.Client
	err = json.Unmarshal(bytes, client)
	log.Println(client)
}
