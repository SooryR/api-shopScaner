package usecases

import (
	"database/sql"
	"strings"

	headerFile "github.com/SooryR/api-ShopScaner/headerFile"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"gopkg.in/guregu/null.v4"
)

// ClientUseCase holds all business related functions for client
type ClientUseCase struct {
	db           *sqlx.DB
	clientStore  headerFile.ClientStore
	emailer      headerFile.EmailerUseCase
	authStore    headerFile.AuthStore
	fireStore    headerFile.FireStoreClient
	companyStore headerFile.CompanyStore
}