package services

import (
	"os"
	"reflect"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/yellyoshua/elections-app/server/repository"
)

func TestGraphqlService(t *testing.T) {
	os.Setenv("DATABASE_URI", "mongodb://root:dbpwd@localhost:27017")

	repository.Initialize()

	srv := NewGraphqlService()

	var expected string = "*services.Service"

	if reflect.TypeOf(srv).String() != expected {
		t.Errorf("Should be a pointer of Service interface %s", reflect.TypeOf(srv).String())
	}

	users, err := srv.GetUsers(graphql.ResolveParams{})

	if err != nil {
		t.Errorf("Error getting users, error: %s %v", err, users)
	}

}
