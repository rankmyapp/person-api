package entity

import (
	"regexp"
	"time"

	"github.com/rpuglielli/person-api/pkg/errors"
)

type Person struct {
	ID         string    `json:"id,omitempty" bson:"_id,omitempty"`
	ExternalID string    `json:"externalId" bson:"externalId"`
	FirstName  string    `json:"firstName" bson:"firstName"`
	LastName   string    `json:"lastName" bson:"lastName"`
	Email      string    `json:"email" bson:"email"`
	Phone      string    `json:"phone" bson:"phone"`
	Category   string    `json:"category" bson:"category"`
	Created    time.Time `json:"created" bson:"created"`
	Updated    time.Time `json:"updated" bson:"updated"`
}

func (p *Person) Validate() error {
	if p.ExternalID == "" {
		return errors.NewValidationError("externalId is required")
	}
	if p.FirstName == "" {
		return errors.NewValidationError("firstName is required")
	}
	if p.Email == "" {
		return errors.NewValidationError("email is required")
	}
	if p.Email != "" && !isValidEmail(p.Email) {
		return errors.NewValidationError("invalid email")
	}
	if len(p.FirstName) > 50 {
		return errors.NewValidationError("firstName cannot be longer than 50 characters")
	}
	if len(p.LastName) > 50 {
		return errors.NewValidationError("lastName cannot be longer than 50 characters")
	}

	return nil
}

func isValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
