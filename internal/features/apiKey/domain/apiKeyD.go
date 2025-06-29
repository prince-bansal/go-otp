package domain

import (
	"crypto/rand"
	"github.com/prince-bansal/go-otp/internal/features/organisation/domain"
	"time"
)

const (
	API_KEY_LENGTH      = 11
	API_KEY_SUB_LENGTHS = 4
)

type ApiKeyD struct {
	Id             int                  `json:"id"`
	Key            string               `json:"key"`
	OrganisationId int                  `json:"organisationId"`
	Organisation   domain.OrganisationD `json:"organisation"`
	Expiry         time.Time            `json:"expiry"`
	CreatedAt      time.Time            `json:"createdAt"`
	UpdatedAt      time.Time            `json:"updatedAt"`
}

func (d *ApiKeyD) GenerateKey() string {
	key := ""

	subLength :=
		API_KEY_SUB_LENGTHS
	totalLength :=
		API_KEY_LENGTH

	for len(key) < totalLength {
		salt := rand.Text()
		key += salt[:subLength]

		if len(key) < totalLength {
			key += "-"
		}
	}

	return key
}
