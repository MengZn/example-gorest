package main

import (
	"math/rand"
	"time"

	"github.com/go-rest/pkg/exchange"
	"github.com/go-rest/pkg/server"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	resource := server.NewResource("/")
	exchange := exchange.NewExChanger()
	resource.AddMethod(server.Post, "", exchange.CreateExChange)

	apiServerHandler := server.NewAPIServerHandler()
	apiServerHandler.RegisterResource(resource)
	apiServerHandler.ServeHTTP()
}
