package formatter

type ifStatement struct {
	Exp        *exp
	Body       []Block
	ElseIfPart []*elseifStatement
	ElsePart   *elseStatement
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
	return cur.Token.Type == nEnd
}

func (s *ifStatement) Append(el *element) {}

func (s *ifStatement) AppendStatement(st statementIntf) {
	switch v := st.(type) {
	case *exp:
		s.Exp = v
	case *elseifStatement:
		s.ElseIfPart = append(s.ElseIfPart, v)
	case *elseStatement:
		s.ElsePart = v
	default:
		s.Body = append(s.Body, newBlock(st))
	}
}

type elseifStatement struct {
	Exp  *exp
	Body []Block
}

func (elseifStatement) New() statementIntf {
	return &elseifStatement{}
}

func (elseifStatement) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (elseifStatement) TypeOf() typeStatement {
	return tsIf
}

func (s *elseifStatement) IsEnd(prev, cur *element) bool {
	return cur.Token.Type == nEnd || cur.Token.Type == nElse
}

func (s *elseifStatement) Append(el *element) {}

func (s *elseifStatement) AppendStatement(st statementIntf) {
	if v, ok := st.(*exp); ok {
		s.Exp = v
		return
	}

	s.Body = append(s.Body, newBlock(st))
}

type elseStatement struct {
	Body []Block
}

func (elseStatement) New() statementIntf {
	return &elseStatement{}
}

func (elseStatement) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (elseStatement) TypeOf() typeStatement {
	return tsIf
}

func (s *elseStatement) IsEnd(prev, cur *element) bool {
	return false
}

func (s *elseStatement) Append(el *element) {}

func (s *elseStatement) AppendStatement(st statementIntf) {
	s.Body = append(s.Body, newBlock(st))
}
