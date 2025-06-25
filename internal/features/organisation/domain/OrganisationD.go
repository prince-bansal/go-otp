package domain

import (
	"github.com/prince-bansal/go-otp/models"
	"time"
)

type OrganisationD struct {
	Id        string         `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Package   models.Package `json:"package"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}
