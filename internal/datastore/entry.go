package datastore

import "time"

type Entry struct {
	Value *Value

	ExpiresAt time.Time
}
