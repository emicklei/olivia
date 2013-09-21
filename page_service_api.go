package main

import (
	"github.com/emicklei/go-restful"
)

type PageService struct{}

func (p PageService) AddWebserviceTo(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/rss").Consumes("*/*").Produces("text/html")
	ws.Route(ws.GET("/nu.nl/algemeen").To(p.getNU_Algemeen))
	container.Add(ws)
}
