package datastore

import (
	"sync"
	"time"
)

type Store struct {
	mu   sync.RWMutex
	data map[string]Entry
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]Entry),
	}
}

func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = Entry{
		Value: value,
	}
}

func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	entry, ok := s.data[key]
	s.mu.RUnlock()

	if !ok {
		return "", false
	}

	// Expired?
	if !entry.ExpiresAt.IsZero() && time.Now().After(entry.ExpiresAt) {
		s.Delete(key)
		return "", false
	}

	return entry.Value, true
}

func (s *Store) Delete(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.data[key]; !ok {
		return false
	}

	delete(s.data, key)
	return true
}

func (s *Store) Exists(key string) bool {
	_, ok := s.Get(key)
	return ok
}

func (s *Store) Expire(key string, seconds int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	entry, ok := s.data[key]
	if !ok {
		return false
	}

	entry.ExpiresAt = time.Now().Add(time.Duration(seconds) * time.Second)

	s.data[key] = entry

	return true
}

func (s *Store) TTL(key string) int {
	s.mu.RLock()
	entry, ok := s.data[key]
	s.mu.RUnlock()

	if !ok {
		return -2
	}

	if entry.ExpiresAt.IsZero() {
		return -1
	}

	ttl := int(time.Until(entry.ExpiresAt).Seconds())

	if ttl < 0 {
		s.Delete(key)
		return -2
	}

	return ttl
}

func (s *Store) Keys() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	keys := make([]string, 0, len(s.data))

	for key := range s.data {
		keys = append(keys, key)
	}

	return keys
}

func (s *Store) DeleteExpired() {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()

	for key, entry := range s.data {

		if entry.ExpiresAt.IsZero() {
			continue
		}

		if now.After(entry.ExpiresAt) {
			delete(s.data, key)
		}
	}
}

func (s *Store) Export() map[string]Entry {

	s.mu.RLock()
	defer s.mu.RUnlock()

	copyMap := make(map[string]Entry)

	for k, v := range s.data {
		copyMap[k] = v
	}

	return copyMap
}

func (s *Store) Import(data map[string]Entry) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = data
}
