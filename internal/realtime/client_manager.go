package realtime

import (
	"chat/internal/models"
	"sync"
)

type ClientMangager struct {
	Mu      sync.RWMutex
	Clients map[string]*models.Client
}

func New() *ClientMangager {
	return &ClientMangager{
		Clients: map[string]*models.Client{},
	}
}

func (cm *ClientMangager) Get(Nickname string) (*models.Client, bool) {
	cm.Mu.RLock()
	cl, ok := cm.Clients[Nickname]
	cm.Mu.RUnlock()
	return cl, ok
}

func (cm *ClientMangager) Add(Nickname string, cl *models.Client) {
	cm.Mu.Lock()
	cm.Clients[Nickname] = cl
	cm.Mu.Unlock()
}

func (cm *ClientMangager) Remove(Nickname string) {
	cm.Mu.Lock()
	delete(cm.Clients, Nickname)
	cm.Mu.Unlock()
}