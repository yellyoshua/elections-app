package authentication

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
)

func TestEncriptionToken(t *testing.T) {
	var tokenValue string = "someOne"
	var secret string = "secret"

	auth := NewAuthentication(secret)
	token, err := auth.CreateToken(tokenValue)

	if err != nil {
		t.Errorf("Error creating token: %v", err)
	}

	if len(token) == 0 {
		t.Error("Token is short")
	}

	auth.VerifyToken(token, func(tokenValueResult string, err error) {
		if err != nil {
			t.Errorf("Error verifying token: %v", err)
		}

		if tokenValue != tokenValueResult {
			t.Errorf("Token not equal to expected: expected(%v) result(%v)", tokenValue, tokenValueResult)
		}
	})

	if returnedSecret := auth.GetSecret(); secret != returnedSecret {
		t.Error("Error secret is not equal to expected")
	}
}

func TestEncriptionFailureToken(t *testing.T) {
	var secret string = "secret"

	auth := NewAuthentication(secret)
	fakeToken := "adadasdasdas"

	auth.VerifyToken(fakeToken, func(tokenValueResult string, err error) {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed == 0 {
				t.Errorf("Error expected not malformed token, not: %v", err)
			}
		}
	})

	if returnedSecret := auth.GetSecret(); secret != returnedSecret {
		t.Error("Error secret is not equal to expected")
	}
}

func TestEncriptionFailSecretToken(t *testing.T) {
	var secret string = "secretito"

	auth := NewAuthentication(secret)
	fakeToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhbnkiLCJleHAiOjE2MDk5OTM2MzQsImlzcyI6ImF1dGgtYXBwIiwic3ViIjoic29tZU9uZSJ9.hjv7AzIZSCL_yv9emr5iFu7WJvdo-Qrjjd7Mwv1zCiM"

	auth.VerifyToken(fakeToken, func(tokenValueResult string, err error) {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) == 0 {
				t.Errorf("Error expected not valid/expired token, not: %v", err)
			}
		}
	})

	if returnedSecret := auth.GetSecret(); secret != returnedSecret {
		t.Error("Error secret is not equal to expected")
	}
}
