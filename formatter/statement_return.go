package formatter

type returnStatement struct {
	Explist *explist
}

func (returnStatement) New() statementIntf {
	return &returnStatement{}
}

func (returnStatement) InnerStatement() statementIntf {
	return nil
}

func (returnStatement) TypeOf() typeStatement {
	return tsReturn
}

func (s *returnStatement) IsEnd(prev, cur *element) bool {
	if nReturn == cur.Token.Type {
		return false
	}

	branch := getsyntax(tokenID(nReturn))
	_, ok := branch[cur.Token.Type]

	return !ok
}

func (s *returnStatement) HasSyntax(el element) bool {
	return false
}

func (s *returnStatement) Append(el *element) {
	if el == nil || el.Token.Type == nReturn {
		return
	}
	s.Explist.List = append(s.Explist.List, newExp(el))
}

func (s *returnStatement) AppendStatement(st statementIntf) {
	el, ok := st.(*explist)
	if !ok {
		return
	}
	s.Explist = el
}
