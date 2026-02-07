package handlers

import "chat/internal/realtime"

type Handlers struct {
	ClientManager *realtime.ClientMangager
}

func New(clientManager *realtime.ClientMangager) *Handlers {
	return &Handlers{
		ClientManager: clientManager,
	}
}
