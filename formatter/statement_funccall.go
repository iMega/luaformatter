package formatter

type funcCallStatement struct {
	Prefixexp *prefixexpStatement
	Explist   *explist
}

func (funcCallStatement) New() statementIntf {
	return &funcCallStatement{}
}

func (funcCallStatement) InnerStatement(prev, cur *element) statementIntf {
	return &explist{}
}

func (funcCallStatement) TypeOf() typeStatement {
	return tsFuncCallStatement
}

func (s *funcCallStatement) IsEnd(prev, cur *element) bool {
	return true
}

func (s *funcCallStatement) Append(el *element) {
}

func (s *funcCallStatement) AppendStatement(st statementIntf) {
	switch v := st.(type) {
	case *prefixexpStatement:
		s.Prefixexp = v
	case *explist:
		s.Explist = v
	}
}
