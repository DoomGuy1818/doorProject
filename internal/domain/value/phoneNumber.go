package value

import (
	"fmt"
	"strings"
)

type PhoneNumber string

func NewPhoneNumber(phone string) (Email, error) {
	if strings.HasPrefix(phone, "+7") {
		return Email(phone), nil
	}
	return "", fmt.Errorf("invalid phone number")
}
