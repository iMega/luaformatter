package formatter

import (
	"bytes"
	"io"
)

type commentStatement struct {
	Element *element
}

func (commentStatement) New() statementIntf {
	return &commentStatement{}
}

func (commentStatement) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (commentStatement) TypeOf() typeStatement {
	return tsNone
}

func (s *commentStatement) IsEnd(prev, cur *element) bool {
	return true
}

func (s *commentStatement) Append(el *element) {
	if el.Token.Type == nComment {
		el.Token.Lexeme = bytes.TrimLeft(el.Token.Lexeme, "--")
		el.Token.Lexeme = bytes.TrimSpace(el.Token.Lexeme)
		el.Token.Value = string(el.Token.Lexeme)
	}
	s.Element = el
}

func (s *commentStatement) AppendStatement(st statementIntf) {}

func (s *commentStatement) GetBody(prevSt statementIntf, cur *element) statementIntf {
	return prevSt
}

func (s *commentStatement) Format(c *Config, p printer, w io.Writer) error {
	if _, err := w.Write([]byte("-- ")); err != nil {
		return err
	}

	if err := s.Element.Format(c, p, w); err != nil {
		return err
	}

	return nil
}
