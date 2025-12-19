package api

import (
	"serverless/handlers/auth"
	"serverless/interfaces"
	"serverless/repositories"

	"github.com/aws/aws-lambda-go/events"
)

type ApiHandler struct {
	DbInstance     *interfaces.IDbInstance
	UserRepository repositories.UserRepository
}

func NewApiHandler(dbInstance *interfaces.IDbInstance) *ApiHandler {
	userRepository := repositories.NewUserRepository(dbInstance)
	return &ApiHandler{
		DbInstance:     dbInstance,
		UserRepository: *userRepository,
	}
}

func (api ApiHandler) HandleRegister(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return auth.Register(request, api.UserRepository)
}

func (api ApiHandler) HandleLogin(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return auth.Login(request, api.UserRepository)
}
