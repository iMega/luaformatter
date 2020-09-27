package formatter

type assignmentStatement struct {
	LocalElement  *element
	Namelist      namelist
	EqPart        *element
	Explist       explist
	LastTokenType tokenID
}

func (assignmentStatement) New() statementIntf {
	return &assignmentStatement{}
}

func (assignmentStatement) InnerStatement() statementIntf {
	return nil
}

func (assignmentStatement) TypeOf() typeStatement {
	return tsAssignment
}

func (s *assignmentStatement) IsEnd(prev, cur *element) bool {
	return len(s.Namelist) == len(s.Explist.List)
	// return el.Token.Type != nComma
}

func (s *assignmentStatement) HasSyntax(el element) bool {
	var syntax []tokenID

	switch s.LastTokenType {
	case nID:
		syntax = []tokenID{nComma, nEq}
	case nComma:
		syntax = []tokenID{
			nID,
			// exp
			nNumber,
		}
	case nEq:
		// nil | false | true | Numeral | LiteralString | ‘...’ | functiondef |
		// prefixexp | tableconstructor | exp binop exp | unop exp
		syntax = []tokenID{
			nNil,
			nFalse,
			nTrue,
			nNumber,
			nString,
			nVararg,
			nFunction,
			nCurlyBracket,
		}
	case nNumber:
		syntax = []tokenID{nComma}
	}

	for _, v := range syntax {
		if v == tokenID(el.Token.Type) {
			return true
		}
	}

	return false
}

func (s *assignmentStatement) Append(el *element) {
	s.LastTokenType = tokenID(el.Token.Type)

	if el.Token.Type == nEq {
		s.EqPart = el
		return
	}

	if s.EqPart == nil {
		s.Namelist = append(s.Namelist, el)
		return
	}

	switch el.Token.Type {
	case nNumber:
		s.Explist.List = append(s.Explist.List, newExp(el))
	}
}

func (s *assignmentStatement) AppendStatement(st statementIntf) {
	switch v := st.(type) {
	case *functionStatement:
		s.Explist.List = append(s.Explist.List, newExp(v))
	}
}
