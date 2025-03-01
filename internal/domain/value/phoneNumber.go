package value

import (
	"fmt"
	"strings"
)

type PhoneNumber string

func NewPhoneNumber(phone string) (PhoneNumber, error) {
	if strings.HasPrefix(phone, "+7") {
		return PhoneNumber(phone), nil
	}
	return "", fmt.Errorf("invalid phone number")
}
