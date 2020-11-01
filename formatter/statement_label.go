package formatter

import "bytes"

type labelStatement struct {
	Element *element
}

func (labelStatement) New() statementIntf {
	return &labelStatement{}
}

func (labelStatement) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (labelStatement) TypeOf() typeStatement {
	return tsIf
}

func (s *labelStatement) IsEnd(prev, cur *element) bool {
	return s.Element != nil && cur.Token.Type == nLabel
}

func (s *labelStatement) Append(el *element) {
	el.Token.Lexeme = bytes.Trim(el.Token.Lexeme, "::")
	el.Token.Lexeme = bytes.TrimSpace(el.Token.Lexeme)
	el.Token.Value = string(el.Token.Lexeme)

	s.Element = el
}

func (s *labelStatement) AppendStatement(st statementIntf) {}
