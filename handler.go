package lambda_gateway

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

func Start(mux http.Handler) {
	if mux == nil {
		mux = http.DefaultServeMux
	}

	lambda.Start(func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		httpReq, err := CreateRequest(ctx, req)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}

		httpRes := CreateResponse()
		mux.ServeHTTP(httpRes, httpReq)

		return httpRes.End(), nil
	})
}
