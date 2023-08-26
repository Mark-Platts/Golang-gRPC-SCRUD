package local_store

import (
	"github.com/google/uuid"
	"markplatts.org/scrud/models"
	"google.golang.org/grpc/status"
    "google.golang.org/grpc/codes"
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
	ms, ok := ls.messageList[id]
	if ok {
		return &ms, nil
	}
	return nil, status.Errorf(codes.NotFound, "No message found for provided id.")
}

func (ls *LocalStore) Update(id string, m *models.MessageStore) error {
	_, ok := ls.messageList[id]
	if ok {
		ls.messageList[id] = *m
		return nil
	}
	return status.Errorf(codes.NotFound, "No message found for provided id.")
}

func (ls *LocalStore) Delete(id string) error {
	_, ok := ls.messageList[id]
	if ok {
		delete(ls.messageList, id)
		return nil
	}
	return status.Errorf(codes.NotFound, "No message found for provided id.")
}

// func (ls *LocalStore) List() ([]*models.MessageStore, error) {

// }