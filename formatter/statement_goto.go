package formatter

type gotoStatement struct {
	Element *element
}

func (gotoStatement) New() statementIntf {
	return &gotoStatement{}
}

func (gotoStatement) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (gotoStatement) TypeOf() typeStatement {
	return tsIf
}

func (s *gotoStatement) IsEnd(prev, cur *element) bool {
	return s.Element != nil
}

func (s *gotoStatement) Append(el *element) {
	if el.Token.Type == nGoto {
		return
	}

	s.Element = el
}

func (s *gotoStatement) AppendStatement(st statementIntf) {}
