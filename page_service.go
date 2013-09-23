package main

import (
	"fmt"
	"github.com/SlyMarbo/rss"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/renderbee"
	"html/template"
	"net/http"
)

func (p PageService) getNU_Algemeen(req *restful.Request, resp *restful.Response) {
	feed, err := rss.Fetch("http://www.nu.nl/feeds/rss/algemeen.rss")
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	feedFragment := renderbee.NewFragment(feed, NUFeed_Template)
	nupage := renderbee.NewFragmentMap(NUPage_Template)
	nupage.Add("NUFeed", feedFragment)

	canvas := renderbee.NewHtmlCanvas(resp.ResponseWriter)
	canvas.Render(nupage)
}

var NUFeed_Template = template.Must(template.New("NUFeed").Parse(`
<div class="items">
	{{range .Items}}
	<div class="item">
		<h3>{{.Title}}</h3>
		<h6>{{.Date}}</h6>
		<p>{{.Content}}</p>
	</div>
	{{end}}
</div>`))

var NUPage_Template = template.Must(template.New("NUPage").Parse(`
<html>
<body>
{{.Render "NUFeed"}}
</body>
</html>
`))
