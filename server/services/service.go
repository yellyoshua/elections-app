package services

import (
	"github.com/gin-gonic/gin"
	gql "github.com/graphql-go/graphql"
)

// REST _
type REST interface {
	HandlerHome(ctx *gin.Context)
	HandlerAPI(ctx *gin.Context)
	HandlerLoginUser(ctx *gin.Context)
}

// GRAPHQL Queries and Mutators
type GRAPHQL interface {
	// Queries
	GetUsers(params gql.ResolveParams) (interface{}, error)
	FindUser(params gql.ResolveParams) (interface{}, error)
	// Mutators
	UpdateUsers(params gql.ResolveParams) (interface{}, error)
}

// Service _
type Service int

// NewGraphqlService intance services
func NewGraphqlService() GRAPHQL {
	service := new(Service)
	return service
}

// NewRestService instance services
func NewRestService() REST {
	service := new(Service)
	return service
}
