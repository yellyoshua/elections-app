package authentication

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Initialize _
func Initialize() {}

// Auth interface of Authentication module
type Auth interface {
	CreateToken(tokenValue string) (string, error)
	VerifyToken(tokenString string) (string, error)
	GetSecret() string
}

// Token _
type Token struct {
	Secret string
	Exp    int64
}

type ExpirationTime = int64

func CreateExpirationTime(duration time.Duration) ExpirationTime {
	return time.Now().Add(duration).Unix()
}

// New this instance a interface with methods that Create, Verify and get Secret
func New(secret string, token_expiration_time ExpirationTime) Auth {
	return &Token{
		Secret: secret,
		Exp:    token_expiration_time,
	}
}

// GetSecret _
func (t *Token) GetSecret() string {
	return t.Secret
}

// VerifyToken method decript with secret key and return in callback
func (t *Token) VerifyToken(tokenString string) (string, error) {
	var tokenValue string
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing with HMAC method not allowed")
		}
		return []byte(t.Secret), nil
	})

	if err != nil {
		return tokenValue, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tokenValue = claims["sub"].(string)
		return tokenValue, nil
	}

	return tokenValue, fmt.Errorf("invaled token")
}

// CreateToken method create a token string encoded with secret key
func (t *Token) CreateToken(tokenValue string) (string, error) {
	secret := []byte(t.Secret)
	claims := &jwt.StandardClaims{
		Issuer:    "auth-app",
		Subject:   tokenValue,
		Audience:  "any",
		ExpiresAt: t.Exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString(secret)
	return jwtToken, err
}
