package server

import restful "github.com/emicklei/go-restful"

type Resource struct {
	// normally one would use DAO (data access object)
	WebService *restful.WebService
}

func NewWebService(path string) *Resource {
	ws := new(restful.WebService)
	ws.Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	return &Resource{WebService: ws}
}

func (r *Resource) AddMethod(method Method, path string, fn restful.RouteFunction) {
	var route *restful.RouteBuilder
	switch method {
	case Get:
		route = r.WebService.GET(path).To(fn)
	case Post:
		route = r.WebService.POST(path).To(fn)
	case Put:
		route = r.WebService.PUT(path).To(fn)
	case Delete:
		route = r.WebService.DELETE(path).To(fn)
	}
	r.WebService.Route(route)
}
