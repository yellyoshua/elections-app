package services

import (
	gql "github.com/graphql-go/graphql"
	"github.com/yellyoshua/elections-app/logger"
	"github.com/yellyoshua/elections-app/server/models"
	"github.com/yellyoshua/elections-app/server/repository"
	"go.mongodb.org/mongo-driver/bson"
)

// GetUsers __
func (s *Service) GetUsers(params gql.ResolveParams) (interface{}, error) {
	col := repository.NewRepository(repository.CollectionUsers)
	var users []models.User = make([]models.User, 0)
	filter := bson.M{}
	err := col.Find(filter, &users)
	return users, err
}

// FindUser __
func (s *Service) FindUser(params gql.ResolveParams) (interface{}, error) {
	col := repository.NewRepository(repository.CollectionUsers)
	userID, _ := params.Args["id"].(string)
	var user models.User
	filter := bson.D{bson.E{Key: "_id", Value: userID}}
	err := col.FindOne(filter, &user)
	return user, err
}

// UpdateUsers {params models.UpdateUserGQL}
func (s *Service) UpdateUsers(params gql.ResolveParams) (interface{}, error) {
	username, _ := params.Args["username"].(string)

	return models.User{
		Username: username,
	}, nil
}

func checkError(err error) {
	if err != nil {
		logger.ServerFatal("repository error: %v", err)
	}
}
