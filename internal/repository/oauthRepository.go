package repository

import (
	"agahi-plus-plus/internal/dto"
)

type OAuthRepository interface {
	GetToken(dto dto.OAuthToken) (*dto.AccessTokenResponse, error)
	GetPhoneNumber(dto dto.PhoneNumber) (*dto.PhoneNumberResponse, error)
}
