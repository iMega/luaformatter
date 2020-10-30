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

func (l line) Format() []byte {
	var (
		buf   bytes.Buffer
		pad   bool
		start bool
	)

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
}
