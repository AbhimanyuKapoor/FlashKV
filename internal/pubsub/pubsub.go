package pubsub

import (
	"sync"
)

type PubSub struct {
	mu   sync.RWMutex
	subs map[string][]chan string
}

func New() *PubSub {
	return &PubSub{
		subs: make(map[string][]chan string),
	}
}

func (p *PubSub) Subscribe(channel string) chan string {
	p.mu.Lock()
	defer p.mu.Unlock()

	ch := make(chan string, 10)

	p.subs[channel] = append(p.subs[channel], ch)

	return ch
}

func (p *PubSub) Publish(channel, message string) {

	p.mu.RLock()
	defer p.mu.RUnlock()

	for _, sub := range p.subs[channel] {
		select {
		case sub <- message:
		default:
		}
	}
}

func (p *PubSub) Unsubscribe(channel string, target chan string) {

	p.mu.Lock()
	defer p.mu.Unlock()

	subs := p.subs[channel]

	for i, ch := range subs {

		if ch == target {

			close(ch)

			p.subs[channel] = append(subs[:i], subs[i+1:]...)

			break
		}
	}
}
