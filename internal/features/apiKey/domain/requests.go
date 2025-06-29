package domain

import "github.com/go-playground/validator/v10"

type ApiKeyRequest struct {
	OrganisationId int    `json:"organisationId" validate:"required"`
	Expiry         string `json:"expiry" validate:"datetime=2006-01-01"`
}

func (r *ApiKeyRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)
	return err
}
