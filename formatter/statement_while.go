package formatter

type whileStatement struct {
	Exp *exp
	Do  *doStatement
}

func (whileStatement) New() statementIntf {
	return &whileStatement{}
}

func (whileStatement) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (whileStatement) TypeOf() typeStatement {
	return tsIf
}

func (s *whileStatement) IsEnd(prev, cur *element) bool {
	return cur.Token.Type == nEnd
}

func (s *whileStatement) Append(el *element) {}

func (s *whileStatement) AppendStatement(st statementIntf) {
	switch v := st.(type) {
	case *exp:
		s.Exp = v
	case *doStatement:
		s.Do = v
	}
}
