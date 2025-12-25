package main

import (
	"context"
	"os"
	"serverless/app"
	"serverless/dtos"
	"serverless/handlers/auth"
	"serverless/middleware"
	"strings"

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

	lambda.Start(func(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		route := strings.TrimPrefix(request.RawPath, "/prod")
		switch route{
		case auth.RegisterRoute:
			return middleware.ValidationPipeline(app.ApiHandler.HandleRegister, &dtos.UserRegisterDto{})(request)
		case auth.LoginRoute:
			return app.ApiHandler.HandleLogin(request)
		case "/health-check":
			return events.APIGatewayV2HTTPResponse{
				StatusCode: 200,
				Body: "Running",
			}, nil
		default:
			return events.APIGatewayV2HTTPResponse{
				StatusCode: 404,
				Body:       "No Route found",
			}, nil
		}
	})
}
