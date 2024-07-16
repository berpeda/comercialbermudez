package main

import (
	"context"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"

	// "github.com/aws/aws-lambda-go/events"
	"github.com/berpeda/comercialbermudez/database"
	"github.com/berpeda/comercialbermudez/handlers"

	"github.com/berpeda/comercialbermudez/awsgo"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	var region string = "eu-west-1"
	awsgo.StartAWS(region)

	if !ValidParameters() {
		panic("The parameters that should be send are SecretName, UrlPrefix")
	}

	var res *events.APIGatewayProxyResponse
	prefix := os.Getenv("UrlPrefix")
	path := strings.Replace(request.RawPath, prefix, "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers

	database.ReadScecret()

	status, message := handlers.Handlers(path, method, body, header, request)
	//

	headersRes := map[string]string{
		"Content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(message),
		Headers:    headersRes,
	}

	return res, nil
}

func ValidParameters() bool {
	// verifies if all required environment variables are established
	requiredEnvVars := []string{"SecretName", "UrlPrefix"}

	for _, envVar := range requiredEnvVars {
		if _, ok := os.LookupEnv(envVar); !ok {
			return false
		}
	}

	return true
}
