package browser

import (
	"fmt"
	"strings"
)

func validateEmail(email string) error {
	if len(email) < 3 || len(email) > 254 {
		return fmt.Errorf("email length is invalid: %d characters", len(email))
	}

	atIndex := strings.Index(email, "@")
	if atIndex < 1 || atIndex > len(email)-5 {
		return fmt.Errorf("invalid position for @ symbol in email: %s", email)
	}

	localPart := email[:atIndex]
	domainPart := email[atIndex+1:]

	if len(localPart) < 1 || len(domainPart) < 3 {
		return fmt.Errorf("local part or domain part is too short in email: %s", email)
	}

	if strings.HasPrefix(localPart, ".") || strings.HasPrefix(domainPart, ".") {
		return fmt.Errorf("email cannot start with a dot: %s", email)
	}

	if strings.HasSuffix(localPart, ".") || strings.HasSuffix(domainPart, ".") {
		return fmt.Errorf("email cannot end with a dot: %s", email)
	}

	if strings.Contains(localPart, " ") || strings.Contains(domainPart, " ") {
		return fmt.Errorf("email cannot contain spaces: %s", email)
	}

	dotIndex := strings.LastIndex(domainPart, ".")
	if dotIndex < 1 || dotIndex > len(domainPart)-2 {
		return fmt.Errorf("invalid domain format in email: %s", email)
	}

	return nil
}
