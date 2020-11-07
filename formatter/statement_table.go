package formatter

type tableStatement struct {
	List []field
}

type field struct {
	Key   *exp
	Val   *exp
	Sqare bool
}

func (tableStatement) New() statementIntf {
	return &tableStatement{}
}

func (tableStatement) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (tableStatement) TypeOf() typeStatement {
	return tsIf
}

func (s *tableStatement) IsEnd(prev, cur *element) bool {
	return cur.Token.Type == nEnd
}

func (s *tableStatement) Append(el *element) {}

func (s *tableStatement) AppendStatement(st statementIntf) {
	// s.Body = append(s.Body, newBlock(st))
}

func (s *tableStatement) GetBody(prevSt statementIntf, cur *element) statementIntf {
	return prevSt
}
