package domain

import (
	"time"
)

type OrganisationD struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	//Package   models.Package `json:"package"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
