package middleware

import (
	"serverless/interfaces"

	"github.com/aws/aws-lambda-go/events"
)

func ValidationPipeline(next func(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error), dto interfaces.DtoInterface) func(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return func(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		body := request.Body
		if err := dto.Validate(body); err != nil {
			return events.APIGatewayV2HTTPResponse{
				StatusCode: 400,
				Body:       err.Error(),
			}, nil
		}
		return next(request)
	}
}
