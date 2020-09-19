package formatter

type exp struct {
	Element *element // nil | false | true | Numeral | LiteralString | ‘...’
	Table   *tableconstructor
	Func    *functionStatement
	Binop   *binop
	Unop    *unop
}

func (exp) New() statementIntf {
	return &exp{}
}

func (s *exp) IsEnd(el *element) bool {
	return false
}

func (s *exp) HasSyntax(el element) bool {
	return false
}

func (s *exp) Append(el *element) {}

func (s *exp) AppendStatement(st statementIntf) {}
