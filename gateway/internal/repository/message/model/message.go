package model

import "github.com/google/uuid"

type Message struct {
	ID   uuid.UUID
	Text string
}
