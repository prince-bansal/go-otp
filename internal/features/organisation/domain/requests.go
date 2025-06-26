package domain

import (
	"github.com/go-playground/validator/v10"
)

type OrganisationRequest struct {
	Name  string `json:"name" validate:"required,min=3"`
	Email string `json:"email" validate:"required,email"`
}

func (d *OrganisationRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(d)
	return err
}
