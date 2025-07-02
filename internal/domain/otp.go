package domain

import (
	"math/rand"
	"strconv"
)

type Otp struct {
	Otp            string `json:"otp"`
	OrganisationId int    `json:"organisationId"`
	MobileNo       int    `json:"mobileNo"`
}

const HIGH = 999999
const LOW = 100000

func (d *Otp) GenerateOtp() string {
	otp := LOW + rand.Intn(HIGH-LOW+1)
	return strconv.Itoa(otp)
}

func (d *Otp) ToGenerateResponse() *OtpGenerateResponse {
	return &OtpGenerateResponse{Otp: d.Otp}
}
