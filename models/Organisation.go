package models

import (
	"github.com/prince-bansal/go-otp/internal/features/organisation/domain"
	"time"
)

type Organisation struct {
	Id    int    `json:"id" gorm:"id"`
	Name  string `json:"name" gorm:"name"`
	Email string `json:"email" gorm:"email"`
	//Package   Package   `json:"package" gorm:"package"`
	CreatedAt time.Time `json:"createdAt" gorm:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"updated_at"`
}

func (m *Organisation) ToDomain() *domain.OrganisationD {
	return &domain.OrganisationD{
		Id:    m.Id,
		Name:  m.Name,
		Email: m.Email,
		//Package:   m.Package,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (m *Organisation) FromDomain(d *domain.OrganisationD) {
	m.Id = d.Id
	m.Name = d.Name
	m.Email = d.Email
	m.CreatedAt = d.CreatedAt
	m.UpdatedAt = d.UpdatedAt
}
