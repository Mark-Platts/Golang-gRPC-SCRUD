package models

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MessageStore struct {
	timeCreated timestamp.Timestamp
	message string 
}

func NewMessageStore(message string) *MessageStore {
	return &MessageStore{
		timeCreated: *timestamppb.Now(),
		message: message,
	}
}

func (ms MessageStore)GetTimeCreated() *timestamp.Timestamp {
	return &ms.timeCreated
}

func (ms MessageStore)GetMessage() string {
	return ms.message
}