package handlers

import "net/http"

func (h *Handlers) CheckSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("SessionId")
	if err != nil || cookie.Value == "" {
		http.Error(w, "No Session", http.StatusUnauthorized)
		return
	}

	_, exists := h.ClientManager.GetClientBySessionId(cookie.Value)
	if !exists {
		http.Error(w, "No Session", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
