package templates

import (
	"os"
	"strings"
	"text/template"
)

const tmpl = `local {{join .Namelist ", "}} = {{join .Explist ", "}}`
const tmplFunc = `
function .Name({{join .Args ", "}})
	{{.Body}}
end
`

const tmplLocalFunc = `
{{pad}}local function .Name({{join .Args ", "}})
{{pad}}{{indentSize}}{{.Body}}
{{pad}}end
`

// local function Name(
//     arg1,
//     arg2,
// )
//     body()
// end
const tmplLocalFuncLong = `
{{pad}}local function .Name(
{{range .Args}}
{{indentSize}}{{.}},
{{end}}
)
{{pad}}{{indentSize}}{{.Body}}
{{pad}}end
`

type Block struct {
	Namelist []string
	Explist  []string
}

var funcs = template.FuncMap{"join": strings.Join}

func String(b Block) {
	t, _ := template.New("tmpl").Funcs(funcs).Parse(tmpl)
	t.Execute(os.Stdout, b)
}
