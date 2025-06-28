package models

import "time"

type Organisation struct {
	Id        string    `json:"id" gorm:"id"`
	Name      string    `json:"name" gorm:"name"`
	Email     string    `json:"email" gorm:"email"`
	Package   Package   `json:"package" gorm:"package"`
	CreatedAt time.Time `json:"createdAt" gorm:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"updated_at"`
}
