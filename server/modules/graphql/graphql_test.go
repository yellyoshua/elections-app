package graphql

import (
	"os"
	"reflect"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/yellyoshua/elections-app/server/repository"
)

// TODO: Test graphql all queries
func TestInitialize(t *testing.T) {
	gql := Initialize()

	var expected string = "*graphql.Service"

	if reflect.TypeOf(gql).String() != expected {
		t.Error("Should be a pointer of Service interface")
	}
}

func TestHandler(t *testing.T) {
	Initialize()

	handler := Handler()

	var expected string = "*handler.Handler"

	if reflect.TypeOf(handler).String() != expected {
		t.Error("Should be a http handler")
	}
}

var queryUsers string = `
	{
		users {
			_id
		}
	}
`

var queryCreateUser string = `
	{
		createUser(
			name: "DemoName", surname:"DemoSurname",username:"demo",email:"demo@demo.com",password:"demodotcom"
		){
			_id
			username
		}
	}
`

var queryCreateUser1 string = `
	{
		createUser(
			name: "DemoName", surname:"DemoSurname",username:"demo1",email:"demo1@demo.com",password:"demodotcom"
		){
			_id
			username
		}
	}
`

var queryCreateUser2 string = `
	{
		createUser(
			name: "DemoName", surname:"DemoSurname",username:"demo2",email:"demo2@demo.com",password:"demodotcom"
		){
			_id
			username
		}
	}
`

var queryCreateUser3 string = `
	{
		createUser(
			name: "DemoName", surname:"DemoSurname",username:"demo3",email:"demo3@demo.com",password:"demodotcom"
		){
			_id
			username
		}
	}
`

func TestGraphqlService(t *testing.T) {
	os.Setenv("DATABASE_NAME", "golangtest")
	os.Setenv("DATABASE_URI", "mongodb://root:dbpwd@localhost:27017")
	var indexes bool = false

	repository.Initialize(indexes)
	srv := NewGraphqlService()
	schema, err := setupSchemas(srv)

	if err != nil {
		t.Errorf("Error creating schemas %v", err)
	}

	userCreationQueries := []string{
		queryCreateUser,
		queryCreateUser1,
		queryCreateUser2,
		queryCreateUser3,
	}

	for _, query := range userCreationQueries {
		params := graphql.Params{Schema: schema, RequestString: query}

		if r := graphql.Do(params); len(r.Errors) > 0 {
			t.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
		}
	}

	params := graphql.Params{Schema: schema, RequestString: queryUsers}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		t.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	// rJSON, _ := json.Marshal(r)

	// t.Fatalf("Bla Bla: %s", rJSON)

	// var expected string = "*graphql.Service"

	// if reflect.TypeOf(srv).String() != expected {
	// 	t.Errorf("Should be a pointer of Service interface %s", reflect.TypeOf(srv).String())
	// }

	// if users, err := srv.GetUsers(graphql.ResolveParams{}); err != nil {
	// 	t.Errorf("Error getting users, error: %s | %v", err, users)
	// }

	// if user, err := srv.CreateUser(graphql.ResolveParams{Args: argsCreateUser}); err == nil {
	// 	t.Errorf("Error creating user, error: %s | %v", err, user)
	// }

}
