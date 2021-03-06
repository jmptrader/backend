package registration

import (
	"log"

	"fairlance.io/mailer"
)

type RegistrationContext struct {
	RegisteredUserRepository *RegisteredUserRepository
	Mailer                   mailer.Mailer
}

func NewContext(dbName string) *RegistrationContext {
	registeredUserRepository, err := NewRegisteredUserRepository(dbName)
	if err != nil {
		log.Fatalf("Failed to open user repository: %q", err.Error())
	}

	// Setup context
	context := &RegistrationContext{
		RegisteredUserRepository: registeredUserRepository,
		Mailer: mailer.MailgunMailer{},
	}

	return context
}
