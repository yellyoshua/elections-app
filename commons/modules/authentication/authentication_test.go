package authentication

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestEncriptionToken(t *testing.T) {
	var tokenValue string = "someOne"
	var secret string = "secret"

	auth := New(secret, CreateExpirationTime(time.Second*5))
	token, err := auth.CreateToken(tokenValue)

	if err != nil {
		t.Errorf("Error creating token: %v", err)
	}

	if len(token) == 0 {
		t.Error("Token is short")
	}

	tokenValueResult, errToken := auth.VerifyToken(token)

	if errToken != nil {
		t.Errorf("Error verifying token: %v", errToken)
	}

	if tokenValue != tokenValueResult {
		t.Errorf("Token not equal to expected: expected(%v) result(%v)", tokenValue, tokenValueResult)
	}

	if returnedSecret := auth.GetSecret(); secret != returnedSecret {
		t.Error("Error secret is not equal to expected")
	}
}

func TestEncriptionFailureToken(t *testing.T) {
	var secret string = "secret"

	auth := New(secret, CreateExpirationTime(time.Second*5))
	fakeToken := "adadasdasdas"

	_, errToken := auth.VerifyToken(fakeToken)

	if ve, ok := errToken.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed == 0 {
			t.Errorf("Error expected not malformed token, not: %v", errToken)
		}
	}

	if returnedSecret := auth.GetSecret(); secret != returnedSecret {
		t.Error("Error secret is not equal to expected")
	}
}

func TestEncriptionFailSecretToken(t *testing.T) {
	var secret string = "secretito"

	auth := New(secret, CreateExpirationTime(time.Second*5))
	fakeToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhbnkiLCJleHAiOjE2MDk5OTM2MzQsImlzcyI6ImF1dGgtYXBwIiwic3ViIjoic29tZU9uZSJ9.hjv7AzIZSCL_yv9emr5iFu7WJvdo-Qrjjd7Mwv1zCiM"

	_, errToken := auth.VerifyToken(fakeToken)

	if ve, ok := errToken.(*jwt.ValidationError); ok {
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) == 0 {
			t.Errorf("Error expected not valid/expired token, not: %v", errToken)
		}
	}

	if returnedSecret := auth.GetSecret(); secret != returnedSecret {
		t.Error("Error secret is not equal to expected")
	}
}
