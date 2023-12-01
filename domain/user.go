package domain

import "github.com/google/uuid"

type ExchangeCredential struct {
	UserUID   uuid.UUID `json:"user_uid"`
	APIKey    string    `json:"api_key"`
	ApiSecret string    `json:"api_secret"`
}
