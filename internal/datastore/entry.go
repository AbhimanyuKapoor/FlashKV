package datastore

import "time"

type Entry struct {
	Value     string    `json:"value"`
	ExpiresAt time.Time `json:"expires_at"`
}
