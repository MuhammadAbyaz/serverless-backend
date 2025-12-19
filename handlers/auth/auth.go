package auth

import (
	"serverless/interfaces"
	"serverless/repositories"

	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func Login(request events.APIGatewayV2HTTPRequest, userRepository repositories.UserRepository) (events.APIGatewayV2HTTPResponse, error) {
	responseBody := interfaces.ResponseBody{
		Message: "Logged In",
		Data: map[string]any{
			"accessToken": "something",
		},
	}

	jsonBody, err := json.Marshal(responseBody)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       "Failed to Marshal JSON object",
		}, err
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       string(jsonBody),
	}, nil
}

func Register(request events.APIGatewayV2HTTPRequest, userRepository repositories.UserRepository) (events.APIGatewayV2HTTPResponse, error) {
	responseBody := interfaces.ResponseBody{
		Message: "Registered",
		Data: map[string]any{
			"accessToken": "something",
		},
	}

	jsonBody, err := json.Marshal(responseBody)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       "Failed to Marshal JSON object",
		}, err
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       string(jsonBody),
	}, nil
}
