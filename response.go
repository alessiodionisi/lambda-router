package lambda_router

import (
	"bytes"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

type Response struct {
	wroteHeader bool
	header      http.Header
	gwRes       events.APIGatewayProxyResponse
	buff        bytes.Buffer
}

func (res *Response) Header() http.Header {
	return res.header
}

func (res *Response) Write(b []byte) (int, error) {
	if !res.wroteHeader {
		res.WriteHeader(http.StatusOK)
	}

	return res.buff.Write(b)
}

func (res *Response) WriteHeader(statusCode int) {
	if res.wroteHeader {
		return
	}

	res.gwRes.StatusCode = statusCode

	headerMap := make(map[string]string)
	for k, v := range res.Header() {
		if len(v) > 0 {
			headerMap[k] = v[0]
		}
	}

	res.gwRes.Headers = headerMap
	res.wroteHeader = true
}

func (res *Response) End() events.APIGatewayProxyResponse {
	res.gwRes.IsBase64Encoded = false
	res.gwRes.Body = res.buff.String()
	return res.gwRes
}

func CreateResponse() *Response {
	return &Response{
		header: make(http.Header),
	}
}
