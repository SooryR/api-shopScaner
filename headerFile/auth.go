package headerFile

import (
	"crypto/rsa"
	"database/sql/driver"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"gopkg.in/guregu/null.v4"
)


// JWTClaims are custom claims extending default ones.
type JWTClaims struct {
	ClientID string `json:"id"`
	Email    string `json:"email"`
	Admin    bool   `json:"admin"`
	jwt.StandardClaims
}



// Auth model
type Auth struct {
	ID            int         `db:"id" json:"-"`
	CreatedAt     time.Time   `db:"created_at" json:"-"`
	UpdatedAt     time.Time   `db:"updated_at" json:"-"`
	Token         null.String `db:"token" json:"-"`
	AuthType      null.String `db:"auth_type" json:"-"`
	Email         null.String `db:"email" json:"-"`
	PasswordHash  []byte      `db:"password_hash" json:"-"`
	ClientID      string      `db:"client_id" json:"-"`
	Stage         SignUpFlow  `db:"sign_up_flow" json:"signUpStage"`
	Blocked       bool        `db:"blocked" json:"-"`
	FirebaseToken string      `json:"firebaseToken"`
}

type AuthStore interface {
	CreateWithEmail(db DB, auth *Auth, clientID string) (*Auth, error)
}

type AuthUseCase interface {
	SignupEmail(password string, token string, email string, firstName string, lastName string, photo string, about string, isTutor null.Bool) (*Client, *Auth, error)
	LoginEmail(email string, password string) (*Client, *Auth, error)
	GenerateToken(claims *JWTClaims, signingKey *rsa.PrivateKey) (string, error)
	UpdateSignUpFlow(clientID string, stage SignUpFlow) error
	BlockClient(clientID string, blocked bool) error
}
