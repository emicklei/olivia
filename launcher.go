package main

import (
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
	"path"
)

func main() {
	wsContainer := restful.NewContainer()
	PageService{}.AddWebserviceTo(wsContainer)

	ws := new(restful.WebService)
	ws.Consumes("*/*")
	ws.Route(ws.GET("/static/{resource}").To(staticFromPathParam))
	wsContainer.Add(ws)

	log.Printf("start listening on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}

func staticFromPathParam(req *restful.Request, resp *restful.Response) {
	http.ServeFile(
		resp.ResponseWriter,
		req.Request,
		path.Join("/Users/emicklei/Workspaces/go/src/bitbucket.org/emicklei/olivia", req.PathParameter("resource")))
}
