package formatter

type breakStatement struct{}

func (breakStatement) New() statementIntf {
	return &breakStatement{}
}

func (breakStatement) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (breakStatement) TypeOf() typeStatement {
	return tsIf
}

func (s *breakStatement) IsEnd(prev, cur *element) bool {
	return true
}

func (s *breakStatement) Append(el *element) {}

func (s *breakStatement) AppendStatement(st statementIntf) {}
