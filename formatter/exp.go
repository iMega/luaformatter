package formatter

type exp struct {
	Element   *element           // nil | false | true | Numeral | LiteralString | ‘...’
	Table     *tableconstructor  // {
	Func      *functionStatement // function
	Binop     *element
	Unop      *element
	Exp       *exp
	Prefixexp *prefixexpStatement
}

func (exp) New() statementIntf {
	return &exp{}
}

func (exp) InnerStatement(prev, cur *element) statementIntf {
	switch cur.Token.Type {
	case nFunction:
		return &functionStatement{}
		// case nTable:
		// return &tableconstructor{}
	}

	return nil
}

func (exp) TypeOf() typeStatement {
	return tsExp
}

func (s *exp) IsEnd(prev, cur *element) bool {
	if cur.Token.Type == nEnd {
		return true
	}

	var syntax = map[tokenID]map[tokenID]bool{
		nNumber: {
			nAddition:   false,
			nInequality: false,
			nEquality:   false,
		},
		nID: {
			nAddition:      false,
			nInequality:    false,
			nEquality:      false,
			nSquareBracket: false,
		},
		nAddition: {
			nNumber: false,
		},
		nInequality: {
			nNumber: false,
		},
		nEquality: {
			nNumber: false,
		},
		nSquareBracket: {
			nString: false,
		},
	}

	v, ok := syntax[tokenID(prev.Token.Type)]
	if !ok {
		return true
	}

	res, ok := v[tokenID(cur.Token.Type)]
	if !ok {
		return true
	}

	return res
}

func (s *exp) HasSyntax(el element) bool {
	return false
}

func (s *exp) Append(el *element) {
	if el.Token.Type == nComma {
		return
	}

	if el.Token.Type == nParentheses || el.Token.Type == nClosingParentheses {
		return
	}

	if s.Element == nil {
		switch el.Token.Type {
		case nSubtraction:
			s.Unop = el

			return

		case nNot:
			s.Unop = el

			return

		case nLength:
			s.Unop = el

			return

		case nBitwiseExclusiveOR:
			s.Unop = el

			return
		}
	}

	if el.Token.Type >= nAnd && el.Token.Type <= nGreaterOrEqual {
		s.Binop = el

		return
	}

	s.Element = el
}

func (s *exp) AppendStatement(st statementIntf) {
	switch v := st.(type) {
	// case *tableconstructor:
	// s.Table = v
	case *functionStatement:
		s.Func = v
	case *exp:
		s.Exp = v
	case *prefixexpStatement:
		if s.Element != nil {
			v.Element = s.Element
			s.Element = nil
		}

		s.Prefixexp = v
	}
}
