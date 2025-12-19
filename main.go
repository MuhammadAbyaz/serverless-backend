package main

import (
	"context"
	"os"
	"serverless/app"
	"serverless/handlers/auth"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"gorm.io/driver/postgres"
)

func main() {
	// Can Pass Any Driver
	dsn := os.Getenv("CONNECTION_STRING")
	driver := postgres.New(postgres.Config{
		DSN: dsn,
	})
	app := app.NewApp(driver, dsn)

	lambda.Start(func (ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		switch request.RawPath {
		case auth.RegisterRoute:
			return app.ApiHandler.HandleRegister(request)
		case auth.LoginRoute:
			return app.ApiHandler.HandleLogin(request)
		default:
			return events.APIGatewayV2HTTPResponse{
				StatusCode: 404,
				Body:       "No Route found",
			}, nil
		}
	})
}
