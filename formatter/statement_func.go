package formatter

type functionStatement struct {
	LocalElement bool
	NamePart     *element
	ParlistPart  parlist
	Body         []Block
	Anonymous    bool
}

func (functionStatement) New() statementIntf {
	return &functionStatement{}
}

func (functionStatement) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (functionStatement) TypeOf() typeStatement {
	return tsFunction
}

func (s *functionStatement) IsEnd(prev, cur *element) bool {
	if cur.Token.Type == nEnd {
		cur.Resolved = true

		return true
	}

	return false
}

func (s *functionStatement) HasSyntax(el element) bool {
	return false
}

func (s *functionStatement) Append(el *element) {
	switch el.Token.Type {
	case nLocal:
		s.LocalElement = true

	case nParentheses:
		s.Anonymous = s.NamePart == nil

	case nID:
		if s.Anonymous {
			s.ParlistPart = append(s.ParlistPart, el)

			return
		}

		s.NamePart = el

	case nVararg:
		s.ParlistPart = append(s.ParlistPart, el)
	}
}

func (s *functionStatement) AppendStatement(st statementIntf) {
	if v, ok := st.(*returnStatement); ok {
		s.Body = append(s.Body, Block{Return: v})

		return
	}

	s.Body = append(s.Body, Block{Statement: newStatement(st)})
}
