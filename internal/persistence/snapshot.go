package persistence

import (
	"encoding/json"
	"os"

	"FlashKV/internal/datastore"
)

type Snapshot struct {
	Data map[string]datastore.Entry `json:"data"`
}

func Save(path string, store *datastore.Store) error {

	snapshot := Snapshot{
		Data: store.Export(),
	}

	bytes, err := json.MarshalIndent(snapshot, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, bytes, 0644)
}

func Load(path string, store *datastore.Store) error {

	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var snapshot Snapshot

	if err := json.Unmarshal(bytes, &snapshot); err != nil {
		return err
	}

	store.Import(snapshot.Data)

	return nil
}
