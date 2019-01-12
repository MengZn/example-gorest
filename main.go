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
	//Method:POST URL:localhost:8080 HEADERS: Content-Type:application/json
	//BODY:{
	//     "name": "test3",
	//     "value": "100"
	// }
	resource.AddMethod(server.Post, "", exchange.CreateExChange)
	//Method:DELETE URL:localhost:8080/{name} HEADERS: Content-Type:application/json
	resource.AddMethod(server.Delete, "/{name}", exchange.DelExChange)

	apiServerHandler := server.NewAPIServerHandler()
	apiServerHandler.RegisterResource(resource)
	apiServerHandler.ServeHTTP()
}
