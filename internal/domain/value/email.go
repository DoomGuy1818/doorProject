package value

import (
	"fmt"
	"strings"
)

type Email string

func NewEmail(email string) (Email, error) {
	if strings.Contains(email, "@") {
		return Email(email), nil
	}
	return "", fmt.Errorf("invalid email format")
}
