package models

import "time"

type ApiKey struct {
	Id           string       `json:"id"`
	Key          string       `json:"key"`
	Organisation Organisation `json:"organisation"`
	Expiry       time.Time    `json:"expiry"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    time.Time    `json:"updatedAt"`
}
