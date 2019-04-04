# Lambda Router

Use standard Go net/http Mux (or other routers) on AWS Lambda with AWS API Gateway.

## Install

`go get github.com/adnsio/lambda-router`

## Examples

With `go-chi`:
```golang
package main

import (
    "net/http"
    "github.com/go-chi/chi"
    "github.com/adnsio/lambda-router"
)

func main() {
    mux := chi.NewRouter()
    mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("welcome"))
    })
    lambda_router.Start(mux)
}
```
