package middleware

import "github.com/aws/aws-lambda-go/events"

func ValidateUser(next func(events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error)) func(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return func(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		return next(request)
	}
}
