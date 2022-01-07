package model

import "github.com/google/uuid"

type Author struct {
	Id   uuid.UUID
	Name string
}
