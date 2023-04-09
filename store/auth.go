package store

import (
	headerFile "github.com/Arun4rangan/api-ShopScaner/headerFile"
	"github.com/pkg/errors"
	"gopkg.in/guregu/null.v4"
)

const insertEmailAuth string = `
INSERT INTO auth (email, password_hash, auth_type, client_id)
VALUES ($1, $2, $3, $4)
RETURNING *
	`

// CreateWithEmail creates auth row with email in db
func (au *AuthStore) CreateWithEmail(db rfrl.DB, auth *rfrl.Auth, clientID string) (*rfrl.Auth, error) {
	var createdAuth rfrl.Auth
	row := db.QueryRowx(insertEmailAuth, auth.Email, auth.PasswordHash, rfrl.EMAIL, clientID)

	err := row.StructScan(&createdAuth)

	return &createdAuth, errors.Wrap(err, "CreateWithEmail")
}