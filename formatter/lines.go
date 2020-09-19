package formatter

import (
	"bytes"
)

type line struct {
	Type     int
	Length   int
	Elements []element
	Branch   branch
	Template template
}

func NewLine() line {
	return line{}
}

// var templates = map[template]func(l line){
// 	tmplVarList: func(l line) {
// 		el := make([][]byte, len(l.Elements))
// 		for i, e := range l.Elements {
// 			el[i] = e.Token.Lexeme
// 		}
// 		fmt.Println(string(bytes.Join(el, []byte(" "))))
// 	},
// 	tmplFunctionCall: func(l line) {
// 		var el [][]byte
// 		for _, e := range l.Elements {
// 			el = append(el, e.Token.Lexeme)
// 			if e.Token.Type == nEq {
// 				el = append(el, []byte(" "))
// 			}
// 		}
// 	},
// }

func (l line) Format() []byte {
	var buf bytes.Buffer
	var pad bool
	var start bool
	for _, e := range l.Elements {
		if e.Token.Type == nCurlyBracket && l.Length > 80 {
			pad = true
			start = true
			buf.Write(e.Token.Lexeme)
			buf.Write([]byte("\n"))
			continue
		}

		if e.Token.Type == nClosingCurlyBracket && l.Length > 80 {
			pad = false
			buf.Write(e.Token.Lexeme)
			buf.Write([]byte("\n"))
			continue
		}

		if pad && start {
			buf.Write(bytes.Repeat([]byte(" "), 4))
			start = false
		}

		buf.Write(e.Token.Lexeme)

		if pad && e.Token.Type == nComma {
			buf.Write([]byte("\n"))
			start = true
		}

		if !pad && e.AddSpace {
			buf.Write([]byte(" "))
		}
	}
	buf.Write([]byte("\n"))
	return buf.Bytes()
	// fmt.Println(string(bytes.Join(el, []byte(""))))
}

// func (l line) Format() {
// 	var splitter int
// 	for i, e := range l.Elements {
// 		if nEq == e.Token.Type {
// 			splitter = i
// 		}
// 	}

// 	nl := make([]string, splitter-1)
// 	for i, e := range l.Elements[1:splitter] {
// 		nl[i] = string(e.Token.Lexeme)
// 	}

// 	exp := make([]string, len(l.Elements)-splitter-1)
// 	for i, e := range l.Elements[splitter+1 : len(l.Elements)] {
// 		exp[i] = string(e.Token.Lexeme)
// 	}

// 	templates.String(templates.Block{
// 		Namelist: nl,
// 		Explist:  exp,
// 	})
// }
