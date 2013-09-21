package main

import (
	"github.com/SlyMarbo/rss"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/renderbee"
	"html/template"
	"io"
	"net/http"
)

func (p PageService) getNU_Algemeen(req *restful.Request, resp *restful.Response) {
	feed, err := rss.Fetch("http://www.nu.nl/feeds/rss/algemeen.rss")
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	nufeed := &NUFeed{feed}
	nupage := renderbee.NewContainer(NUPage_Template)
	nupage.Add("NUFeed", nufeed)
	canvas := &renderbee.HtmlCanvas{resp.ResponseWriter} // NewHtmlCanvas ?
	canvas.Render(nupage)
}

var NUPage_Template = template.Must(template.New("NUPage").Parse(`
<html>
<body>
{{.Render "NUFeed"}}
</body>
</html>
`))

type NUFeed struct {
	Feed *rss.Feed
}

func (f NUFeed) RenderOn(hc *renderbee.HtmlCanvas) {
	io.WriteString(hc, "some items")
}
