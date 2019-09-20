package model

import (
	"crypto/rand"
	"time"
)

type Id []byte

func NewId() Id {
	ret := make(Id, 20)
	if _, err := rand.Read(ret); err != nil {
		panic(err)
	}
	return ret
}

type Model struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
