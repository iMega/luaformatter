package formatter

import (
	"bytes"
	"io"
)

const (
	defaultStyle = "\x1b[0m"
	mMagenta     = "\x1b[35m"
)

type Config struct {
	IndentSize    uint8 `mapstructure:"indent-size"`
	MaxLineLength uint8 `mapstructure:"max-line-length"`
}

func Format(c Config, b []byte, w io.Writer) error {
	doc, err := parse(b)
	if err != nil {
		return err
	}

	p := printer{}

	for _, b := range doc.Body {
		b.Format(&c, p, w)
	}

	return nil
}

func DefaultConfig() Config {
	return Config{
		IndentSize:    4,
		MaxLineLength: 80,
	}
}

type printer struct {
	Pad uint8
}

func (p printer) WritePad(w io.Writer) error {
	b := bytes.Repeat([]byte(" "), int(p.Pad))
	_, err := w.Write(b)

	return err
}

func newLine(w io.Writer) error {
	_, err := w.Write([]byte(newLineSymbol))

	return err
}

func statementAppend(st statementIntf, elOrSt interface{}) {
	switch v := elOrSt.(type) {
	case *element:
		st.Append(v)
	case statementIntf:
		st.AppendStatement(v)
	}
}

func hasExplist(st statementIntf) bool {
	switch st.(type) {
	case *assignmentStatement:
		return true

	case *returnStatement:
		return true
	}

	return false
}

func getLastElement(chain []*element) *element {
	if len(chain) == 0 {
		return nil
	}

	return chain[len(chain)-1]
}

func getLastStatement(chain []statementIntf) statementIntf {
	if len(chain) == 0 {
		return nil
	}

	return chain[len(chain)-1]
}

func newBlock(st statementIntf) Block {
	bl := Block{}

	switch v := st.(type) {
	case *assignmentStatement:
		bl.Statement = statement{Assignment: v}

	case *labelStatement:
		bl.Statement = statement{Label: v}

	case *gotoStatement:
		bl.Statement = statement{Goto: v}

	case *breakStatement:
		bl.Statement = statement{Break: v}

	case *doStatement:
		bl.Statement = statement{Do: v}

	case *whileStatement:
		bl.Statement = statement{While: v}

	case *repeatStatement:
		bl.Statement = statement{Repeat: v}

	case *functionStatement:
		bl.Statement = statement{Function: v}

	case *funcCallStatement:
		bl.Statement = statement{FuncCall: v}

	case *prefixexpStatement:
		bl.Statement = statement{
			FuncCall: &funcCallStatement{
				Prefixexp: v,
			},
		}

	case *ifStatement:
		bl.Statement = statement{If: v}

	case *returnStatement:
		bl.Return = v

	case *commentStatement:
		bl.Statement = statement{Comment: v}
	}

	return bl
}
