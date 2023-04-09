package usecases

import (
	"crypto/rsa"
	"database/sql"

	headerFile "github.com/Arun4rangan/api-ShopScaner/headerFile"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v4"
)

// AuthUseCase holds all business related functions for auth
type AuthUseCase struct {
	db          *sqlx.DB
	authStore   rfrl.AuthStore
	clientStore rfrl.ClientStore
	fireStore   rfrl.FireStoreClient
}

// NewAuthUseCase creates new AuthUseCase
func NewAuthUseCase(
	db sqlx.DB,
	authStore rfrl.AuthStore,
	clientStore rfrl.ClientStore,
	fireStore rfrl.FireStoreClient,
) *AuthUseCase {
	return &AuthUseCase{&db, authStore, clientStore, fireStore}
}


// SignupEmail allows user to signup with email
func (au *AuthUseCase) SignupEmail(
	password string,
	token string,
	email string,
	firstName string,
	lastName string,
	photo string,
	about string,
	isTutor null.Bool,
) (*rfrl.Client, *rfrl.Auth, error) {
	hash, hashError := hashAndSalt([]byte(password))

	if hashError != nil {
		return nil, nil, errors.Wrap(hashError, "SignupEmail")
	}

	newClient := rfrl.NewClient(firstName, lastName, about, email, photo, isTutor, "", "", null.Int{}, "")
	auth := rfrl.Auth{
		Email:        null.StringFrom(email),
		PasswordHash: hash,
	}

	var err = new(error)
	var tx *sqlx.Tx

	tx, *err = au.db.Beginx()

	if *err != nil {
		return nil, nil, errors.Wrap(*err, "SignupWithToken")
	}

	defer rfrl.HandleTransactions(tx, err)

	var createdClient *rfrl.Client

	createdClient, *err = au.clientStore.CreateClient(tx, newClient)

	if *err != nil {
		return nil, nil, *err
	}

	var createdAuth *rfrl.Auth
	createdAuth, *err = au.authStore.CreateWithEmail(tx, &auth, createdClient.ID)

	if *err != nil {
		return nil, nil, *err
	}

	*err = au.fireStore.CreateClient(
		createdClient.ID,
		createdClient.Photo.String,
		createdClient.FirstName.String,
		createdClient.LastName.String,
	)

	if *err != nil {
		return nil, nil, *err
	}

	var firebaseToken string
	firebaseToken, *err = au.fireStore.CreateLoginToken(createdClient.ID)
	createdAuth.FirebaseToken = firebaseToken

	return createdClient, createdAuth, *err
}

// LoginEmail allows user to login with email by checking password hash against the has the passed in
func (au *AuthUseCase) LoginEmail(email string, password string) (*rfrl.Client, *rfrl.Auth, error) {
	c, auth, err := au.authStore.GetByEmail(au.db, email)

	if err != nil {
		return nil, nil, err
	}

	err = bcrypt.CompareHashAndPassword(
		auth.PasswordHash,
		[]byte(password),
	)

	if err != nil {
		return nil, nil, errors.Wrap(err, "LoginEmail")
	}

	firebaseToken, err := au.fireStore.CreateLoginToken(c.ID)
	auth.FirebaseToken = firebaseToken

	return c, auth, err
}


func hashAndSalt(pwd []byte) ([]byte, error) {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "hashAndSalt")
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return hash, nil
}