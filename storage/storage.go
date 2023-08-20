package storage

import (
	"markplatts.org/scrud/models"
	"markplatts.org/scrud/storage/local_store"
)

type Storage interface {
	Insert(m *models.MessageStore) (string, error)
	Retrieve(id string) (*models.MessageStore, error)
	Update(id string, m *models.MessageStore) error
	Delete(id string) error
}

var StorageMechanism local_store.LocalStore = *local_store.NewLocalStore()