package model

import "github.com/gocql/gocql"

type Message struct {
	ID   gocql.UUID
	Text string
}
