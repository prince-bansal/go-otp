package organisation

import (
	"github.com/prince-bansal/go-otp/internal/features/organisation/domain"
	"github.com/prince-bansal/go-otp/models"
	"time"
)

var Organisations = []domain.OrganisationD{
	{
		Id:        "1",
		Name:      "Organisation 1",
		Email:     "org1@mail.com",
		Package:   models.GOLD,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:        "2",
		Name:      "Organisation 2",
		Email:     "org2@mail.com",
		Package:   models.GOLD,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}
