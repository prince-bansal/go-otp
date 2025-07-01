package models

import (
	"github.com/prince-bansal/go-otp/internal/features/otp/domain"
	"gorm.io/gorm"
	"time"
)

type Otp struct {
	Id             int            `json:"id" gorm:"column:id"`
	Otp            string         `json:"otp" gorm:"column:otp"`
	OrganisationId int            `json:"organisationId" gorm:"column:organisation_id"`
	MobileNo       int            `json:"mobileNo" gorm:"column:mobile_number"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at"`
	CreatedAt      time.Time      `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt      time.Time      `json:"updatedAt" gorm:"column:updated_at"`
}

func (m *Otp) ToDomain() *domain.Otp {
	return &domain.Otp{
		Otp:            m.Otp,
		OrganisationId: m.OrganisationId,
	}
}

func (m *Otp) FromDomain(d *domain.Otp) {
	m.Otp = d.Otp
	m.OrganisationId = d.OrganisationId
	m.MobileNo = d.MobileNo
}
