package auth
import (
	"fmt"
	"net/http"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)
type (
	// SignUpPayload is the struct used to hold payload from /signup
	SignUpPayload struct {
		Email           string `json:"email" validate:"email"`
		Name            string `json:"name"`
		Token           string `json:"token"`
		Password        string `json:"password" validate:"gte=10,required_with=email"`
		Type            string `json:"type" validate:"required,oneof= GOOGLE LINKEDIN EMAIL"`
	}
)

func (h *Handler) signupGoogle(payload SignUpPayload) (string, error) {
	email := "arun.ranga@hotmail.ca"

	claims := &jwtCustomClaims{
		email,
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	return GenerateToken(claims)
}

func (h *Handler) signupLinkedIn(payload SignUpPayload) (string, error) {
	email := "arun.ranga@hotmail.ca"

	claims := &jwtCustomClaims{
		email,
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	return GenerateToken(claims)
}

func (h *Handler) signupEmail(payload SignUpPayload) (string, error) {
	email := "arun.ranga@hotmail.ca"

	claims := &jwtCustomClaims{
		email,
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	return GenerateToken(claims)
}

// SignUpPayloadValidation validates user inputs
func (h *Handler) SignUpPayloadValidation(sl validator.StructLevel) {

	payload := sl.Current().Interface().(SignUpPayload)

	switch payload.Type {
	case GOOGLE:
		if len(payload.Token) == 0 {
			sl.ReportError(payload.Token, "token", "Token", "validToken", "")
		}
	case LINKEDIN:
		if len(payload.Token) == 0 {
			sl.ReportError(payload.Token, "token", "Token", "validToken", "")
		}
	case EMAIL:
		if len(payload.Email) == 0 {
			sl.ReportError(payload.Email, "email", "Email", "validEmail", "")
		}
		if len(payload.Password) < 10 {
			sl.ReportError(payload.Email, "password", "Password", "validPassworrd", "")
		}
	}
	// plus can do more, even with different tag than "fnameorlname"
}

// Signup endpoint
func (h *Handler) Signup(c echo.Context) error {
	payload := SignUpPayload{}

	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var token string
	var error error

	switch payload.Type {
	case GOOGLE:
		token, error = h.signupGoogle(payload)
	case LINKEDIN:
		token, error = h.signupLinkedIn(payload)
	case EMAIL:
		token, error = h.signupEmail(payload)
	default:
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Login type - %s is not supported", payload.Token))
	}
	if error != nil {
		panic(fmt.Sprintf("%v", error))
	}
	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}