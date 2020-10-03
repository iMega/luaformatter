package formatter

type ifStatement struct {
	Exp *exp
	// StartElement *element
	// ThenElement  *element
	Body       []Block
	ElseIfPart []elseifPart
	ElsePart   *elsePart
	// EndElement   *element
}

func (ifStatement) New() statementIntf {
	return &ifStatement{}
}

func (ifStatement) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (ifStatement) TypeOf() typeStatement {
	return tsIf
}

func (s *ifStatement) IsEnd(prev, cur *element) bool {
	return false
}

func (s *ifStatement) HasSyntax(el element) bool {
	return false
}

func (s *ifStatement) Append(el *element) {}

func (s *ifStatement) AppendStatement(st statementIntf) {
	s.Exp = st.(*exp)
}

type elseifPart struct {
	StartElement *element
	Exp          exp
	ThenElement  *element
	Body         []Block
}

type elsePart struct {
	StartElement *element
	Body         []Block
}
