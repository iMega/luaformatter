package formatter

type repeatStatement struct {
	Body []Block
	Exp  *exp
}

func (repeatStatement) New() statementIntf {
	return &repeatStatement{}
}

func (repeatStatement) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (repeatStatement) TypeOf() typeStatement {
	return tsIf
}

func (s *repeatStatement) IsEnd(prev, cur *element) bool {
	if prev != nil && prev.Token.Type == nRepeat {
		return false
	}

	if cur.Token.Type == nUntil {
		return false
	}

	if prev != nil && prev.Token.Type == nUntil {
		return false
	}

	return true
}

func (s *repeatStatement) Append(el *element) {}

func (s *repeatStatement) AppendStatement(st statementIntf) {
	if v, ok := st.(*exp); ok {
		s.Exp = v

		return
	}

	s.Body = append(s.Body, newBlock(st))
}

func (s *repeatStatement) GetBody(prevSt statementIntf, cur *element) statementIntf {
	return prevSt
}
