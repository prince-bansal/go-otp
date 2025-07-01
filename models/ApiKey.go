package models

import (
	"github.com/prince-bansal/go-otp/internal/features/apiKey/domain"
	domain2 "github.com/prince-bansal/go-otp/internal/features/organisation/domain"
	"gorm.io/gorm"
	"time"
)

type ApiKey struct {
	Id             int            `json:"id" gorm:"column:id"`
	Key            string         `json:"key" gorm:"column:api_key"`
	Salt           string         `json:"salt" gorm:"column:salt_hash"`
	OrganisationId int            `json:"organisation_id" gorm:"column:organisation_id"`
	Organisation   Organisation   `json:"organisation" gorm:"foreignKey:OrganisationId"`
	Expiry         time.Time      `json:"expiry" gorm:"column:expiry"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;default:null"`
	CreatedAt      time.Time      `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt      time.Time      `json:"updatedAt" gorm:"column:updated_at"`
}

func (m *ApiKey) FromDomain(d *domain.ApiKeyD) {
	m.OrganisationId = d.OrganisationId
	m.Expiry = d.Expiry
	m.Key = d.Key
	m.Salt = d.Salt
}

func (m *ApiKey) ToDomain() *domain.ApiKeyD {

	d := domain.ApiKeyD{
		Id:             m.Id,
		Key:            m.Key,
		OrganisationId: m.OrganisationId,
		Expiry:         m.Expiry,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
		Salt:           m.Salt,
	}

	if m.Organisation.Id != 0 {
		d.Organisation = domain2.OrganisationD{
			Id:        m.Organisation.Id,
			Name:      m.Organisation.Name,
			Email:     m.Organisation.Email,
			CreatedAt: m.Organisation.CreatedAt,
			UpdatedAt: m.Organisation.UpdatedAt,
		}
	}

	return &d
}
