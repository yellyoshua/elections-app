package graphql

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/yellyoshua/elections-app/repository"
	"golang.org/x/net/context"
)

// TODO: Test graphql all queries

type T struct {
	Query     string
	Schema    graphql.Schema
	Expected  interface{}
	Variables map[string]interface{}
}

var QueryCreateUser string = `
	query RootQuery($name: String!, $surname: String!, $username: String!, $email: String!, $password: String!) {
		createUser(
			name: $name, surname: $surname, username: $username, email: $email, password: $password
		){ username }
	}
`

var QueryFindByUsername string = `
	query RootQuery($username: String!) {
		findUserByUsername(username: $username) { username }
	}
`

var QueryUpdateUser string = `
	query RootQuery($userID: String!, $name: String!, $surname: String!, $fullName: String!, $username: String!, $email: String!, $verified: Boolean!, $active: Boolean!) {
		updateUser(
			name: $name, surname: $surname, fullName: $fullName, username: $username, email: $email, verified: $verified, active: $active
		) { username }
	}
`

func dropDBAndRepositoryStart(t *testing.T) {
	os.Setenv("DATABASE_NAME", "golangtest")
	os.Setenv("DATABASE_URI", "mongodb://root:dbpwd@localhost:27017")
	var indexes bool = false

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)

	defer cancel()

	repository.Initialize(indexes)
	col := repository.NewRepository(repository.CollectionUsers)

	if err := col.Database().Drop(ctx); err != nil {
		t.Fatalf("Error dropping database: %v", err)
	}
}

func TestGraphqlModule(t *testing.T) {
	dropDBAndRepositoryStart(t)

	var expectedGQL string = "*graphql.Service"
	gql := Initialize()

	if reflect.TypeOf(gql).String() != expectedGQL {
		t.Error("Should be a pointer of Service interface")
	}

	var expectedHandler string = "*handler.Handler"
	handler := Handler()

	if reflect.TypeOf(handler).String() != expectedHandler {
		t.Error("Should be a http handler")
	}

	schema, err := setupSchemas(gql)
	if err != nil {
		t.Fatalf("Error dropping collection: %v", err)
	}

	var Tests []T = []T{
		{
			Query:  QueryCreateUser,
			Schema: schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"createUser": map[string]interface{}{"username": "testuser"},
				},
			},
			Variables: map[string]interface{}{
				"name":     "TestName",
				"surname":  "TestSurname",
				"username": "testuser",
				"email":    "demo@demo.test",
				"password": "demodotcom",
			},
		},
		{
			Query:  QueryFindByUsername,
			Schema: schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"findUserByUsername": map[string]interface{}{"username": "testuser"},
				},
			},
			Variables: map[string]interface{}{
				"username": "testuser",
			},
		},
	}

	for _, test := range Tests {
		params := graphql.Params{
			Schema:         test.Schema,
			RequestString:  test.Query,
			VariableValues: test.Variables,
		}

		result := graphql.Do(params)
		if len(result.Errors) > 0 {
			t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
		}
		if !reflect.DeepEqual(result, test.Expected) {
			t.Fatalf("wrong result, query: %v, graphql result diff: %v", test.Query, diff(test.Expected, result))
		}
	}
}

func diff(want, got interface{}) []string {
	return []string{fmt.Sprintf("\ngot: %v", got), fmt.Sprintf("\nwant: %v\n", want)}
}
