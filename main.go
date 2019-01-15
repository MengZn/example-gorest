package main

import (
	"example-gorest/pkg/exchange"
	"example-gorest/pkg/server"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	resource := server.NewWebService("/")
	exchange := exchange.NewExChanger()
	//Method:POST URL:localhost:8080 HEADERS: Content-Type:application/json
	//BODY:{
	//     "name": "test3",
	//     "value": "100"
	// }
	resource.AddMethod(server.Post, "", exchange.CreateExChange)
	//Method:DELETE URL:localhost:8080/{name} HEADERS: Content-Type:application/json
	resource.AddMethod(server.Delete, "/{name}", exchange.DelExChange)
	//Method:GET URL:localhost:8080/{name} HEADERS: Content-Type:application/json
	resource.AddMethod(server.Get, "/{name}", exchange.GetExChange)
	//Method:PUT URL:localhost:8080/{name} HEADERS: Content-Type:application/json
	// {
	// 	"name": "test2",
	// 	"value": "300"
	// }
	resource.AddMethod(server.Put, "/{name}", exchange.EditExChange)

	apiServerHandler := server.NewAPIServerHandler()
	apiServerHandler.RegisterResource(resource)
	apiServerHandler.ServeHTTP()
}
