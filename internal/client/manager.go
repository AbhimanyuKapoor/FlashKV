package client

import "sync"

type Manager struct {
	mu sync.RWMutex

	clients map[string]*Client
}

func NewManager() *Manager {

	return &Manager{
		clients: make(map[string]*Client),
	}
}

func (m *Manager) Add(client *Client) {

	m.mu.Lock()
	defer m.mu.Unlock()

	m.clients[client.ID] = client
}

func (m *Manager) Remove(id string) {

	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.clients, id)
}

func (m *Manager) Count() int {

	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.clients)
}

func (m *Manager) Get(id string) (*Client, bool) {

	m.mu.RLock()
	defer m.mu.RUnlock()

	client, ok := m.clients[id]

	return client, ok
}
