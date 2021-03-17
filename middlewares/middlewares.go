package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yellyoshua/elections-app/constants"
	"github.com/yellyoshua/elections-app/models"
	"github.com/yellyoshua/elections-app/modules/authentication"
	"github.com/yellyoshua/elections-app/repository"
	"github.com/yellyoshua/elections-app/utils"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/validator.v2"
)

// MiddlewareConf _
type MiddlewareConf struct {
	SecretToken string
	SecretCSRF  string
	CSRFKey     string
}

type MiddlewareInstance interface {
	AuthRequiredMiddleware(ctx *gin.Context)
	CorsMiddleware(ctx *gin.Context)
	BodyLoginUser(ctx *gin.Context)
	CSRF(ctx *gin.Context)
}

type instanceStruct struct {
	repo repository.Repository
	conf MiddlewareConf
}

// New _
func New(conf MiddlewareConf) MiddlewareInstance {
	repo := repository.New()
	return NewWithRepository(repo, conf)
}

// NewWithRepository _
func NewWithRepository(repo repository.Repository, conf MiddlewareConf) MiddlewareInstance {
	return &instanceStruct{repo: repo, conf: conf}
}

// AuthRequiredMiddleware _
func (i *instanceStruct) AuthRequiredMiddleware(ctx *gin.Context) {
	var session models.Session
	auth := authentication.New(i.conf.SecretToken)
	authorization := ctx.GetHeader("Authorization")
	token := bearerToToken(authorization)

	col := i.repo.Col(constants.CollectionSessions)
	if err := col.FindOne(bson.M{"token": token}, &session); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		ctx.Request.Context().Done()
	}

	if len(session.Token) == 0 {
		ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf(constants.Unauthorized))
		ctx.Request.Context().Done()
	} else {
		tokenValue, errWithTokenFormat := auth.VerifyToken(token)

		if errWithTokenFormat != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			ctx.Request.Context().Done()
		} else {
			ctx.Set("Authorization", tokenValue)
			ctx.Next()
		}
	}
}

// CorsMiddleware habilitate external request
func (i *instanceStruct) CorsMiddleware(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	ctx.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	ctx.Header("Server", "Powered with Golang")
	ctx.Next()
}

// BodyLoginUser handle body request and valid fields
func (i *instanceStruct) BodyLoginUser(ctx *gin.Context) {
	var user models.BodyLoginUser
	body := ctx.Request.Body
	userValidator := validator.NewValidator()

	json.NewDecoder(body).Decode(&user)

	defer body.Close()

	if errs := userValidator.Validate(user); errs != nil {
		// the request did not include all of the User
		// struct fields, so send a http.StatusBadRequest
		// back or something
		ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.Request.Context().Done()
	} else {
		ctx.Next()
	}
}

// CSRF protect attacks of CSRF token
func (i *instanceStruct) CSRF(ctx *gin.Context) {
	CSRFsecret := i.conf.SecretCSRF
	CSRFKey := i.conf.CSRFKey

	uuid, _ := utils.GenerateUniqueID(nil)
	jwt := authentication.New(CSRFsecret)
	token, _ := jwt.CreateToken(uuid)

	ctx.Set(CSRFKey, token)

	ctx.Next()
}

func bearerToToken(bearer string) string {
	bearerPrefix := constants.BearerTokenTemplate
	bearerParsed := strings.TrimPrefix(bearer, bearerPrefix)
	return strings.TrimPrefix(bearerParsed, " ")
}
