package ttl

import (
	"log"
	"time"

	"github.com/Hari08-dev/flashkv/internal/datastore"
)

type Cleaner struct {
	store    *datastore.Store
	interval time.Duration
}

func New(store *datastore.Store, interval time.Duration) *Cleaner {
	return &Cleaner{
		store:    store,
		interval: interval,
	}
}

func (c *Cleaner) Start() {

	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	log.Println("TTL cleaner started")

	for range ticker.C {
		c.store.DeleteExpired()
	}
}
