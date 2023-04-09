package usecases

import (
	"database/sql"
	"strings"

	rfrl "github.com/Arun4rangan/api-rfrl/rfrl"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"gopkg.in/guregu/null.v4"
)

// ClientUseCase holds all business related functions for client
type ClientUseCase struct {
	db           *sqlx.DB
	clientStore  rfrl.ClientStore
	emailer      rfrl.EmailerUseCase
	authStore    rfrl.AuthStore
	fireStore    rfrl.FireStoreClient
	companyStore rfrl.CompanyStore
}