package local_store

import (
	"github.com/google/uuid"
	"markplatts.org/scrud/models"
)

type LocalStore struct{
	messageList map[string]models.MessageStore
}

func NewLocalStore() *LocalStore {
    return &LocalStore{
        messageList: make(map[string]models.MessageStore),
    }
}

func (ls *LocalStore) GetMessages() map[string]models.MessageStore {
	return ls.messageList
}

func (ls *LocalStore) Insert(m *models.MessageStore) (string, error) {
	id := uuid.New().String()
	ls.messageList[id] = *m
	return id, nil
}

func (ls *LocalStore) Retrieve(id string) (*models.MessageStore, error) {
	ms := ls.messageList[id]
	return &ms, nil
}