package formatter

import (
	"io"
)

// var ::=  Name | prefixexp ‘[’ exp ‘]’ | prefixexp ‘.’ Name
// prefixexp ::= var | functioncall | ‘(’ exp ‘)’
// functioncall ::=  prefixexp args | prefixexp ‘:’ Name args
// args ::=  ‘(’ [explist] ‘)’ | tableconstructor | LiteralString

// Name ‘[’ exp1 ‘]’ ‘[’ exp2 ‘]’ ‘.’ Name2 ‘(’ [explist] ‘)’
type prefixexpStatement struct {
	Element          *element
	FuncCall         *funcCallStatement
	FieldAccessorExp *exp
	FieldAccessor    *element
	Prefixexp        *prefixexpStatement
	OneValue         *exp
	IsVar            bool
}

func (prefixexpStatement) New() statementIntf {
	return &prefixexpStatement{}
}

func (s *prefixexpStatement) InnerStatement(prev, cur *element) statementIntf {
	if cur.Token.Type == nSquareBracket {
		// s.IsVar = true
		return &exp{}
	}

	if cur.Token.Type == nParentheses {
		return &funcCallStatement{}
	}

	if cur.Token.Type == nString {
		return &funcCallStatement{}
	}

	return nil
}

func (prefixexpStatement) TypeOf() typeStatement {
	return tsPrefixexpStatement
}

func (s *prefixexpStatement) IsEnd(prev, cur *element) bool {
	if cur.Token.Type == nString {
		return false
	}

	if cur.Token.Type == nParentheses { // function call
		return false
	}

	if cur.Token.Type == nSquareBracket { // function call
		return false
	}

	if cur.Token.Type == nClosingSquareBracket { // function call
		return false
	}

	if cur.Token.Type == nID { // function call
		return false
	}

	if cur.Token.Type == nComma { // assignment statement
		return false
	}

	return true //cur.Token.Type == nAssign
}

func (s *prefixexpStatement) Append(el *element) {
	if el.Token.Type == nSquareBracket {
		return
	}

	if el.Token.Type == nClosingSquareBracket {
		return
	}

	if el.Token.Type == nClosingParentheses {
		return
	}

	if el.Token.Type == nEnd {
		return
	}

	s.Element = el
}

func (s *prefixexpStatement) AppendStatement(st statementIntf) {
	switch v := st.(type) {
	case *prefixexpStatement:
		s.Prefixexp = v
	case *exp:
		s.FieldAccessorExp = v
	case *funcCallStatement:
		v.Prefixexp = &prefixexpStatement{
			Element: s.Element,
		}
		s.Element = nil
		s.FuncCall = v
	}
}

func (s *prefixexpStatement) GetBody(prevSt statementIntf, cur *element) statementIntf {
	return prevSt
}

func (s *prefixexpStatement) Format(c *Config, p printer, w io.Writer) error {
	if st := s.Element; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	if st := s.FieldAccessorExp; st != nil {
		if _, err := w.Write([]byte("[")); err != nil {
			return err
		}

		if err := st.Format(c, p, w); err != nil {
			return err
		}

		if _, err := w.Write([]byte("]")); err != nil {
			return err
		}
	}

	if st := s.Prefixexp; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	if st := s.FuncCall; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	return nil
}
