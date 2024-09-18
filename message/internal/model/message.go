package model

import "github.com/google/uuid"

// уровень
type Message struct {
	ID   uuid.UUID
	Text string
}
