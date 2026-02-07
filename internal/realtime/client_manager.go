package realtime

import (
	"sync"

	"chat/internal/models"
)

type ClientMangager struct {
	Mu                sync.RWMutex
	Clients           map[string]*models.Client
	SessionIDToClient map[string]*models.Client
}

func New() *ClientMangager {
	return &ClientMangager{
		Clients:            make(map[string]*models.Client),
		SessionIDToClient:  make(map[string]*models.Client),
	}
}

func (cm *ClientMangager) Get(Nickname string) (*models.Client, bool) {
	cm.Mu.RLock()
	cl, ok := cm.Clients[Nickname]
	cm.Mu.RUnlock()
	return cl, ok
}

func (cm *ClientMangager) GetClientBySessionId(sessionID string) (*models.Client, bool) {
	cm.Mu.RLock()
	cl, exists := cm.SessionIDToClient[sessionID]
	cm.Mu.RUnlock()
	return cl, exists
}

func (cm *ClientMangager) Add(Nickname string, cl *models.Client) {
	cm.Mu.Lock()
	cm.Clients[Nickname] = cl
	cm.SessionIDToClient[cl.SessionId] = cl
	cm.Mu.Unlock()
}

func (cm *ClientMangager) Remove(Nickname string) {
	cm.Mu.Lock()
	cl, _ := cm.Get(Nickname)
	delete(cm.SessionIDToClient, cl.SessionId)
	delete(cm.Clients, Nickname)
	cm.Mu.Unlock()
}
