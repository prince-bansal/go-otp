package domain

type OtpVerifyRequest struct {
	Otp      string `json:"otp" `
	MobileNo int    `json:"mobileNo"`
}

type OtpGenerateRequest struct {
	MobileNo int `json:"mobileNo"`
}

type OtpGenerateResponse struct {
	Otp string `json:"otp"`
}
