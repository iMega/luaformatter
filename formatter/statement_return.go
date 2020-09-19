package formatter

type returnStatement struct {
	Explist explist
}

func (returnStatement) New() statementIntf {
	return &returnStatement{}
}

func (s *returnStatement) IsEnd(el *element) bool {
	branch := getsyntax(tokenID(nReturn))
	_, ok := branch[el.Token.Type]

	return !ok
}

func (s *returnStatement) HasSyntax(el element) bool {
	return false
}

func (s *returnStatement) Append(el *element) {
	if el.Token.Type == nReturn {
		return
	}
	s.Explist = append(s.Explist, newExp(el))
}

func (s *returnStatement) AppendStatement(st statementIntf) {
	s.Explist = append(s.Explist, newExp(st))
}
