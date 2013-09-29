package ui

import (
	"github.com/emicklei/renderbee"
)

type JQueryMobile struct {
	Version     string // 1.2.1
	CoreVersion string // 1.8.3
}

func NewJQueryMobileFragment(mobileVersion string, coreVersion) *renderbee.Fragment {
	return renderbee.NewFragment(JQueryMobile{mobileVersion,coreVersion},JQueryMobile_Template)
}

var JQueryMobile_Template = template.Must(template.New("NUFeed").Parse(`
<meta name="viewport" content="width=device-width, initial-scale=1"> 
<link rel="stylesheet" href="http://code.jquery.com/mobile/{{.Version}}/jquery.mobile-{{.Version}}.min.css" />
<script src="http://code.jquery.com/jquery-{{.CoreVersion}}.min.js"></script>
<script src="http://code.jquery.com/mobile/{{.Version}}/jquery.mobile-{{.Version}}.min.js"></script>
`))
