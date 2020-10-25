package formatter

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

	return nil
}

func (prefixexpStatement) TypeOf() typeStatement {
	return tsPrefixexpStatement
}

func (s *prefixexpStatement) IsEnd(prev, cur *element) bool {
	return cur.Token.Type == nAssign
}

func (s *prefixexpStatement) Append(el *element) {
	if el.Token.Type == nSquareBracket {
		return
	}

	if el.Token.Type == nClosingSquareBracket {
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
	}
}
