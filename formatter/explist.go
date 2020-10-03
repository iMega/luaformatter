package formatter

type explist struct {
	List []*exp // separator ,
}

func (explist) New() statementIntf {
	return &explist{}
}

func (explist) InnerStatement(prev, cur *element) statementIntf {
	return &exp{}
}

func (explist) TypeOf() typeStatement {
	return tsExpList
}

func (s *explist) IsEnd(prev, cur *element) bool {
	if cur.Token.Type == nComma || prev.Token.Type == nComma {
		return false
	}

	if cur.Resolved {
		return false
	}

	return true
}

func (s *explist) HasSyntax(el element) bool {
	return false
}

func (s *explist) Append(el *element) {}

func (s *explist) AppendStatement(st statementIntf) {
	if v, ok := st.(*exp); ok {
		s.List = append(s.List, v)
	}
}
