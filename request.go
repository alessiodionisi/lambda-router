package lambda_gateway

import (
	"context"
	"encoding/base64"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func CreateRequest(ctx context.Context, req events.APIGatewayProxyRequest) (*http.Request, error) {
	pUrl, err := url.Parse(req.Path)
	if err != nil {
		return nil, err
	}

	pUrlQuery := pUrl.Query()
	for k, v := range req.QueryStringParameters {
		pUrlQuery.Set(k, v)
	}
	pUrl.RawQuery = pUrlQuery.Encode()

	reqBody := req.Body
	if req.IsBase64Encoded {
		b, err := base64.StdEncoding.DecodeString(reqBody)
		if err != nil {
			return nil, err
		}
		reqBody = string(b)
	}

	httpReq, err := http.NewRequest(req.HTTPMethod, pUrl.String(), strings.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	for k, v := range req.Headers {
		httpReq.Header.Set(k, v)
	}

	if httpReq.Header.Get("Content-Length") == "" && reqBody != "" {
		httpReq.Header.Set("Content-Length", strconv.Itoa(len(reqBody)))
	}

	httpReq.URL.Host = httpReq.Header.Get("Host")
	httpReq.Host = httpReq.URL.Host

	return httpReq, nil
}
