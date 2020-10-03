package formatter

type exp struct {
	Element *element           // nil | false | true | Numeral | LiteralString | ‘...’
	Table   *tableconstructor  // {
	Func    *functionStatement // function
	Binop   *element
	Unop    *element
	Exp     *exp
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
			nAddition: false,
		},
	}

	v, ok := syntax[tokenID(prev.Token.Type)]
	if !ok {
		return false
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
	switch el.Token.Type {
	case nNot:
		s.Unop = el
	case nAddition:
	case nNegEq:
		s.Binop = el
	default:
		s.Element = el
	}
}

func (s *exp) AppendStatement(st statementIntf) {
	switch v := st.(type) {
	// case *tableconstructor:
	// s.Table = v
	case *functionStatement:
		s.Func = v
	case *exp:
		s.Exp = v
	}
}
