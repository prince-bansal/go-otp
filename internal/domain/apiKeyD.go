package domain

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	API_KEY_LENGTH      = 11
	API_KEY_SUB_LENGTHS = 4
)

type ApiKeyD struct {
	Id             int           `json:"id"`
	Key            string        `json:"key"`
	Salt           string        `json:"salt"`
	OrganisationId int           `json:"organisationId"`
	Organisation   OrganisationD `json:"organisation"`
	Expiry         time.Time     `json:"expiry"`
	CreatedAt      time.Time     `json:"createdAt"`
	UpdatedAt      time.Time     `json:"updatedAt"`
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
	d.Key = key
	return key
}

func (d *ApiKeyD) GenerateSalt() {
	saltBytes := sha256.Sum256([]byte(d.Key))
	d.Salt = hex.EncodeToString(saltBytes[:])
}

func (d *ApiKeyD) HashKey() error {

	d.GenerateSalt()
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(d.Key), 10)
	if err != nil {
		return err
	}
	d.Key = string(hashedBytes)

	fmt.Printf("hashed for key %s is %s\n", d.Key, string(hashedBytes))
	return nil
}

func (d *ApiKeyD) CompareKey(raw string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(d.Key), []byte(raw))
	if err != nil {
		return false, err
	}
	return true, nil
}
