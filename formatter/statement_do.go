package formatter

type doStatement struct {
	Body []Block
}

func (doStatement) New() statementIntf {
	return &doStatement{}
}

func (doStatement) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (doStatement) TypeOf() typeStatement {
	return tsIf
}

func (s *doStatement) IsEnd(prev, cur *element) bool {
	return cur.Token.Type == nEnd
}

func (s *doStatement) Append(el *element) {}

func (s *doStatement) AppendStatement(st statementIntf) {
	s.Body = append(s.Body, newBlock(st))
}
