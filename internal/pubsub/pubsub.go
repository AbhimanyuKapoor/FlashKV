package pubsub

import (
	"sync"

	"FlashKV/internal/client"
)

type PubSub struct {
	mu sync.RWMutex

	channels map[string]map[string]*client.Client
}

func New() *PubSub {

	return &PubSub{
		channels: make(map[string]map[string]*client.Client),
	}
}

func (p *PubSub) Subscribe(channel string, clint *client.Client) {

	p.mu.Lock()
	defer p.mu.Unlock()

	if _, ok := p.channels[channel]; !ok {
		p.channels[channel] = make(map[string]*client.Client)
	}

	p.channels[channel][clint.ID] = clint
}

func (p *PubSub) Unsubscribe(channel string, client *client.Client) {

	p.mu.Lock()
	defer p.mu.Unlock()

	if subscribers, ok := p.channels[channel]; ok {

		delete(subscribers, client.ID)

		if len(subscribers) == 0 {
			delete(p.channels, channel)
		}
	}
}

func (p *PubSub) Publish(channel, message string) int {

	p.mu.RLock()
	defer p.mu.RUnlock()

	subscribers, ok := p.channels[channel]

	if !ok {
		return 0
	}

	for _, client := range subscribers {

		client.Conn.Write(
			[]byte(
				"[" + channel + "] " + message + "\n",
			),
		)
	}

	return len(subscribers)
}
