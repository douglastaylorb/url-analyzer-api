package services

import (
	"crypto/rand"
	"errors"
	"math/big"

	"github.com/douglastaylorb/url-analyzer-api/models"
)

const (
	lowercaseChars = "abcdefghijklmnopqrstuvwxyz"
	uppercaseChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars    = "0123456789"
	specialChars   = "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	defaultLength  = 8
)

func GeneratePassword(req models.PasswordRequest) (models.PasswordResponse, error) {
	if req.Length == 0 {
		req.Length = defaultLength
	}

	charSet := ""
	if req.Lowercase {
		charSet += lowercaseChars
	}
	if req.Uppercase {
		charSet += uppercaseChars
	}
	if req.Numbers {
		charSet += numberChars
	}
	if req.Special {
		charSet += specialChars
	}

	if charSet == "" {
		return models.PasswordResponse{}, errors.New("at least one character set must be selected")
	}
	password := make([]byte, req.Length)
	for i := range password {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
		if err != nil {
			return models.PasswordResponse{}, err
		}
		password[i] = charSet[randomIndex.Int64()]
	}
	return models.PasswordResponse{Password: string(password)}, nil
}
