package utils

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/artnoi43/todong/lib/enums"
)

func EncodeBcrypt(plain []byte) ([]byte, error) {
	if len(plain) < 6 {
		return nil, enums.ErrPwTooShort
	}
	return bcrypt.GenerateFromPassword(plain, 14)
}

func DecodeBcrypt(hashed, plain []byte) error {
	return bcrypt.CompareHashAndPassword(hashed, plain)
}
