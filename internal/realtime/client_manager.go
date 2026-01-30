package realtime

import "sync"

type ClientMangager struct {
	Mu      sync.RWMutex
	Clients map[string]*Client
}

func New() *ClientMangager {
	return &ClientMangager{
		Clients: map[string]*Client{},
	}
}

func (cm *ClientMangager) Get(Nickname string) (*Client, bool) {
	cm.Mu.RLock()
	cl, ok := cm.Clients[Nickname]
	cm.Mu.RUnlock()
	return cl, ok
}

func (cm *ClientMangager) Add(Nickname string, cl *Client) {
	cm.Mu.Lock()
	cm.Clients[Nickname] = cl
	cm.Mu.Unlock()
}

func (cm *ClientMangager) Remove(Nickname string) {
	cm.Mu.Lock()
	delete(cm.Clients, Nickname)
	cm.Mu.Unlock()
}