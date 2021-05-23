package database

import "time"

type Path struct {
	Id        int64     `json:"id"`
	Path      string    `json:"path" pg:"path,unique,notnull"`
	Type      string    `json:"type"`
	Data      string    `json:"data,omitempty"`
	CreatedAt time.Time `json:"createdAt" pg:"default:now()"`
	UpdatedAt time.Time `json:"updatedAt" pg:"default:now()"`
}
