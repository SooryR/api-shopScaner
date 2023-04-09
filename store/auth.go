package store

import (
	headerFile "github.com/SooryR/api-ShopScaner/headerFile"
	"github.com/pkg/errors"
	"gopkg.in/guregu/null.v4"
)

const insertEmailAuth string = `
INSERT INTO auth (email, password_hash, auth_type, client_id)
VALUES ($1, $2, $3, $4)
RETURNING *
	`

// CreateWithEmail creates auth row with email in db
func (au *AuthStore) CreateWithEmail(db headerFile.DB, auth *headerFile.Auth, clientID string) (*headerFile.Auth, error) {
	var createdAuth headerFile.Auth
	row := db.QueryRowx(insertEmailAuth, auth.Email, auth.PasswordHash, headerFile.EMAIL, clientID)

	err := row.StructScan(&createdAuth)

	return &createdAuth, errors.Wrap(err, "CreateWithEmail")
}