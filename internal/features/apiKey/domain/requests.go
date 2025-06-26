package domain

type ApiKeyRequest struct {
	OrganisationId string `json:"organisationId"`
	ValidTill      int    `json:"validTill"`
}
