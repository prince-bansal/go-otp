package domain

import (
	"github.com/prince-bansal/go-otp/internal/features/organisation/domain"
	"time"
)

type ApiKeyD struct {
	Id             string               `json:"id"`
	Key            string               `json:"key"`
	OrganisationId string               `json:"organisationId"`
	Organisation   domain.OrganisationD `json:"organisation"`
	Expiry         time.Time            `json:"expiry"`
	CreatedAt      time.Time            `json:"createdAt"`
	UpdatedAt      time.Time            `json:"updatedAt"`
}
