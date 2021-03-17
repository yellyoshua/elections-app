package graphql

import (
	"fmt"
	"reflect"
	"testing"

	gql "github.com/graphql-go/graphql"
	"github.com/stretchr/testify/mock"
	"github.com/yellyoshua/elections-app/constants"
	mockRepo "github.com/yellyoshua/elections-app/mocks/repository"
	"github.com/yellyoshua/elections-app/models"
	"github.com/yellyoshua/elections-app/repository"
	"github.com/yellyoshua/elections-app/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: Test graphql all queries
type T struct {
	Query     string
	Schema    gql.Schema
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
		findUserByUsername(username: $username) { username, name }
	}
`

var QueryUpdateUser string = `
	query RootQuery($userID: String!, $name: String!, $surname: String!, $fullName: String!, $username: String!, $email: String!, $verified: Boolean!, $active: Boolean!) {
		updateUser(
			name: $name, surname: $surname, fullName: $fullName, username: $username, email: $email, verified: $verified, active: $active
		) { username }
	}
`

func TestGraphqlModule(t *testing.T) {
	sampleUser := models.User{
		Name:     "TestName",
		Surname:  "TestSurname",
		Username: "testuser",
		Email:    "demo@demo.test",
		Password: "demodotcom",
	}
	repoMock := new(mockRepo.Repository)
	repoClientMock := new(mockRepo.Collection)

	repoMock.On("Col", constants.CollectionUsers).Return(func(repo string) repository.Collection {
		return repoClientMock
	})

	repoClientMock.On("InsertOne", mock.Anything).Return(func(user interface{}) primitive.ObjectID {
		return primitive.NewObjectID()
	}, func(user interface{}) error {
		return nil
	})

	repoClientMock.On("FindOne", mock.AnythingOfType("primitive.M"), mock.Anything).Return(func(filter interface{}, dest interface{}) error {
		f := filter.(primitive.M)
		if f["username"] == sampleUser.Username {
			utils.ReflectValueTo(sampleUser, dest)
			return nil
		}

		return nil
	})

	schema, err := graphqlInit(repoMock)
	if err != nil {
		t.Fatalf("Error dropping collection: %v", err)
	}

	var graphqlQueries []T = []T{
		{
			Query:  QueryCreateUser,
			Schema: schema,
			Expected: &gql.Result{
				Data: map[string]interface{}{
					"createUser": map[string]interface{}{"username": sampleUser.Username},
				},
			},
			Variables: map[string]interface{}{
				"name":     sampleUser.Name,
				"surname":  sampleUser.Surname,
				"username": sampleUser.Username,
				"email":    sampleUser.Email,
				"password": sampleUser.Password,
			},
		},
		{
			Query:  QueryFindByUsername,
			Schema: schema,
			Expected: &gql.Result{
				Data: map[string]interface{}{
					"findUserByUsername": map[string]interface{}{
						"username": sampleUser.Username,
						"name":     sampleUser.Name,
					},
				},
			},
			Variables: map[string]interface{}{
				"username": sampleUser.Username,
			},
		},
		{
			Query:  QueryFindByUsername,
			Schema: schema,
			Expected: &gql.Result{
				Data: map[string]interface{}{
					"findUserByUsername": map[string]interface{}{
						"username": "",
						"name":     "",
					},
				},
			},
			Variables: map[string]interface{}{
				"username": "not-found-username",
			},
		},
	}

	for _, query := range graphqlQueries {
		params := gql.Params{
			Schema:         query.Schema,
			RequestString:  query.Query,
			VariableValues: query.Variables,
		}

		result := gql.Do(params)
		if len(result.Errors) > 0 {
			t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
		}

		if !reflect.DeepEqual(result, query.Expected) {
			t.Fatalf("wrong result, query: %v, graphql result diff: %v", query.Query, diff(query.Expected, result))
		}
	}

	repoClientMock.AssertExpectations(t)
	repoMock.AssertExpectations(t)
}

func diff(want, got interface{}) []string {
	return []string{fmt.Sprintf("\ngot: %v", got), fmt.Sprintf("\nwant: %v\n", want)}
}
