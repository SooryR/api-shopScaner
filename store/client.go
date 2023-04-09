package store

import (
	"database/sql"

	headerFile "github.com/SooryR/api-ShopScaner/headerFile"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"gopkg.in/guregu/null.v4"
)

// ClientStore holds all store related functions for client
type ClientStore struct{}

// NewClientStore creates new clientStore
func NewClientStore() *ClientStore {
	return &ClientStore{}
}

const (
	getClientByIDSQL string = `
SELECT * FROM client
WHERE client.id = $1
	`
	getClientByIDsSQL string = `
	SELECT * FROM client
	WHERE client.id IN (?)
`
)

func (cl *ClientStore) GetClientFromID(db headerFile.DB, id string) (*headerFile.Client, error) {
	var m headerFile.Client
	err := db.QueryRowx(getClientByIDSQL, id).StructScan(&m)
	if err != nil {
		return nil, errors.Wrap(err, "GetClientFromID")
	}
	return &m, nil
}

// CreateClient creates a new row for a client in the database
func (cl *ClientStore) CreateClient(db headerFile.DB, client *headerFile.Client) (*headerFile.Client, error) {
	columns := []string{"first_name", "last_name", "email"}
	values := make([]interface{}, 0)
	values = append(values,
		client.FirstName,
		client.LastName,
		client.Email,
	)

	query := sq.Insert("client").
		Columns(columns...).
		Values(values...).
		Suffix("RETURNING *")

	sql, args, err := query.
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "CreateClient")
	}

	row := db.QueryRowx(
		sql,
		args...,
	)

	var m headerFile.Client

	err = row.StructScan(&m)
	return &m, errors.Wrap(err, "CreateClient")
}
