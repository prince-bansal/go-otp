package apiKey

import (
	"github.com/prince-bansal/go-otp/internal/features/apiKey/domain"
	"time"
)

var organisationId = "test-1"

var apiKeys = []domain.ApiKeyD{
	{
		Id:             "1",
		Key:            "ANSJD-SDGF-SVASBD-SVVSBV",
		OrganisationId: organisationId,
		Expiry:         time.Now(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	},
	{
		Id:             "1",
		Key:            "EDCSE-DSVDSC-ASVD-ASVCS",
		OrganisationId: organisationId,
		Expiry:         time.Now(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	},
}
