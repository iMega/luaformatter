package formatter

const (
	noTmp = iota
	tmplVarList
	tmplFunc
	tmplFunctionCall
)

type template struct {
	Tmpl     int
	AddSpace bool
	LF       bool
}
