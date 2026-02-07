package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	"chat/internal/models"
)

func NewSessionID() (string, error) {
	b := make([]byte, 32) // 256 bits
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	client := &models.Client{}
	err := json.NewDecoder(r.Body).Decode(client)
	if err != nil {
		http.Error(w, "error json request : "+err.Error(), http.StatusInternalServerError)
		return
	}
	sessionId, err := NewSessionID()
	if err != nil {
		http.Error(w, "error generate session : "+err.Error(), http.StatusInternalServerError)
		return
	}
	client.SessionId = sessionId
	h.ClientManager.Add(client.Nickname, client)

	http.SetCookie(w, &http.Cookie{
		Name:     "SessionId",
		Value:    sessionId,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(time.Minute * 60),
	})
	w.WriteHeader(http.StatusOK)
}
