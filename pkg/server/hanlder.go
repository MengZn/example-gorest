package server

import (
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
)

type Method int

const (
	Get Method = iota
	Post
	Put
	Delete
)

type APIServerHandler struct {
	GoRestfulContainer *restful.Container
	HttpServer         *http.Server
}

func (a *APIServerHandler) RegisterResource(res *Resource) {
	a.GoRestfulContainer.Add(res.WebService)
}

func NewAPIServerHandler() *APIServerHandler {
	gorestfulContainer := restful.NewContainer()
	// gorestfulContainer.ServeMux = http.NewServeMux()
	gorestfulContainer.Router(restful.CurlyRouter{})
	return &APIServerHandler{
		GoRestfulContainer: gorestfulContainer,
		HttpServer: &http.Server{
			Addr:    "0.0.0.0:8080",
			Handler: gorestfulContainer,
		},
	}
}
func (a *APIServerHandler) ServeHTTP() {
	log.Print("start listening on 0.0.0.0:8080")
	log.Fatal(a.HttpServer.ListenAndServe())
}
