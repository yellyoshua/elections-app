package graphql

// package graphql

// import (
// 	"fmt"
// 	"time"

// 	gql "github.com/graphql-go/graphql"
// 	"github.com/yellyoshua/elections-app/constants"
// 	"github.com/yellyoshua/elections-app/models"
// 	"github.com/yellyoshua/elections-app/repository"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// // GraphqlResolvers Queries and Mutators
// type GraphqlModulesResolvers interface {
// 	// Queries -> only return a response
// 	GetUsers(params gql.ResolveParams) (interface{}, error)
// 	FindUserByID(params gql.ResolveParams) (interface{}, error)
// 	FindUserByUsername(params gql.ResolveParams) (interface{}, error)
// 	// Mutators -> return a response and update info
// 	CreateUser(params gql.ResolveParams) (interface{}, error)
// 	UpdateUser(params gql.ResolveParams) (interface{}, error)
// }

// type gqlresolverstruct struct {
// 	repo repository.Repository
// }

// func (r *gqlresolverstruct) GetUsers(params gql.ResolveParams) (interface{}, error) {
// 	col := r.repo.Col(constants.CollectionUsers)

// 	filter := bson.D{}
// 	var users []models.User
// 	err := col.Find(filter, &users)

// 	if err != nil {
// 		return make([]models.User, 0), err
// 	}
// 	return users, nil
// }

// func (r *gqlresolverstruct) FindUserByID(params gql.ResolveParams) (interface{}, error) {
// 	var err error
// 	var userID primitive.ObjectID

// 	col := r.repo.Col(constants.CollectionUsers)
// 	userIDString, _ := params.Args["id"].(string)

// 	userID, err = primitive.ObjectIDFromHex(userIDString)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var user models.User
// 	err = col.FindByID(userID, &user)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// func (r *gqlresolverstruct) FindUserByUsername(params gql.ResolveParams) (interface{}, error) {

// 	col := r.repo.Col(constants.CollectionUsers)
// 	username, _ := params.Args["username"].(string)

// 	var user models.User
// 	err := col.FindOne(bson.M{"username": username}, &user)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// func (r *gqlresolverstruct) CreateUser(params gql.ResolveParams) (interface{}, error) {
// 	col := r.repo.Col(constants.CollectionUsers)
// 	name, _ := params.Args["name"].(string)
// 	surname, _ := params.Args["surname"].(string)
// 	username, _ := params.Args["username"].(string)
// 	email, _ := params.Args["email"].(string)
// 	password, _ := params.Args["password"].(string)

// 	user := models.User{
// 		Name:     name,
// 		Surname:  surname,
// 		Username: username,
// 		Fullname: fmt.Sprintf("%s %s", name, surname),
// 		Email:    email,
// 		Password: password,
// 		Active:   true,
// 		Verified: false,
// 		Created:  time.Now().Unix(),
// 	}

// 	id, err := col.InsertOne(user)
// 	if err != nil {
// 		return nil, err
// 	}

// 	user.ID = id
// 	return user, err
// }

// func (r *gqlresolverstruct) UpdateUser(params gql.ResolveParams) (interface{}, error) {
// 	var err error
// 	var userID primitive.ObjectID

// 	col := r.repo.Col(constants.CollectionUsers)
// 	userIDString, _ := params.Args["userID"].(string)
// 	delete(params.Args, "userID")

// 	userID, err = primitive.ObjectIDFromHex(userIDString)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = col.UpdateOne(bson.M{"_id": userID}, params.Args)

// 	if err != nil {
// 		return nil, err
// 	}

// 	params.Args["_id"] = userID
// 	return params.Args, nil
// }
