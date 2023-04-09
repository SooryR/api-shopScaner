package headerFile

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

// Client model
type Client struct {
	ID                   string      `db:"id" json:"id"`
	CreatedAt            time.Time   `db:"created_at" json:"createdAt"`
	UpdatedAt            time.Time   `db:"updated_at" json:"updatedAt"`
	FirstName            null.String `db:"first_name" json:"firstName"`
	LastName             null.String `db:"last_name" json:"lastName"`
	Email                null.String `db:"email" json:"email"`
	IsAdmin              null.Bool   `db:"is_admin" json:"-"`
	VerifiedEmail        null.Bool   `db:"verified_email" json:"verifiedEmail"`
}

// NewClient creates new client model struct
func NewClient(
	firstName string,
	lastName string,
	email string,
) *Client {
	client := Client{
		FirstName:         null.NewString(firstName, firstName != ""),
		LastName:          null.NewString(lastName, lastName != ""),
		Email:             null.NewString(email, email != ""),
	}

	return &client
}


type UpdateClientPayload struct {
	FirstName         string
	LastName          string
	Email             string
}

type ClientStore interface {
	GetClientFromID(db DB, id string) (*Client, error)
	CreateClient(db DB, client *Client) (*Client, error)
}
/*
type ClientUseCase interface {
	CreateClient(firstName string, lastName string, about string, email string, photo string, isTutor null.Bool) (*Client, error)
	UpdateClient(id string, updateParams UpdateClientPayload) (*Client, error)
	GetClient(id string) (*Client, error)
	GetClients(options GetClientsOptions) (*[]Client, error)
	CreateEmailVerification(clientID string, email string, emailType string) error
	VerifyEmail(clientID string, email string, emailType string, passCode string) (*Client, error)
	GetVerificationEmail(clientID string, emailType string) (string, error)
	DeleteVerificationEmail(clientID string, emailType string) error
	GetClientEvents(clientID string, start null.Time, end null.Time, state null.String) (*[]Event, error)
	CreateOrUpdateClientEducation(clientID string, institution string, degree string, fieldOfStudy string, startYear int, endYear int) error
	CreateClientWantingCompanyReferrals(clientID string, IsLookingForReferral bool, companyIds []int) error
	GetClientWantingCompanyReferrals(clientId string) ([]int, error)
}*/
