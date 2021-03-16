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

// JWT _
type JWT struct {
	Secret string
	Exp    int64
}

// New this instance a interface with methods that Create, Verify and get Secret
func New(secret string) Auth {
	expireToken := time.Now().Add(time.Minute * 5).Unix()

	jwt := new(JWT)
	jwt.Exp = expireToken
	jwt.Secret = secret
	return jwt
}

// GetSecret _
func (j *JWT) GetSecret() string {
	return j.Secret
}

// VerifyToken method decript with secret key and return in callback
func (j *JWT) VerifyToken(tokenString string) (string, error) {
	var tokenValue string
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing with HMAC method not allowed")
		}
		return []byte(j.Secret), nil
	})

	if err != nil {
		return tokenValue, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tokenValue = claims["sub"].(string)
		return tokenValue, nil
	}

	return tokenValue, fmt.Errorf("Invaled token")
}

// CreateToken method create a token string encoded with secret key
func (j *JWT) CreateToken(tokenValue string) (string, error) {
	secret := []byte(j.Secret)
	claims := &jwt.StandardClaims{
		Issuer:    "auth-app",
		Subject:   tokenValue,
		Audience:  "any",
		ExpiresAt: j.Exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString(secret)
	return jwtToken, err
}
