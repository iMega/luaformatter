package formatter

type functionStatement struct {
	LocalElement *element
	IDStatement  *element
	NamePart     *element
	ParlistPart  parlist
	Body         []Block
	EndElement   *element
	Anonymous    bool
}

func (functionStatement) New() statementIntf {
	return &functionStatement{}
}

func (s *functionStatement) IsEnd(el *element) bool {
	return el.Token.Type == nEnd
}

func (s *functionStatement) HasSyntax(el element) bool {
	return false
}

func (s *functionStatement) Append(el *element) {
	switch el.Token.Type {
	case nFunction:
		s.IDStatement = el

	case nLocal:
		s.LocalElement = el

	case nEnd:
		s.EndElement = el

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
