package utils

import (
	"github.com/lithammer/shortuuid"
)

func GenerateToken() string {
	token := shortuuid.New()
	return token
}
