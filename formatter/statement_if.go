package formatter

import "io"

type ifStatement struct {
	Exp        *exp
	Body       statementIntf
	ElseIfPart []*elseifStatement
	ElsePart   *elseStatement
}

func (ifStatement) New() statementIntf {
	return &ifStatement{}
}

func (ifStatement) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (ifStatement) TypeOf() typeStatement {
	return tsIf
}

func (s *ifStatement) IsEnd(prev, cur *element) bool {
	return cur.Token.Type == nEnd
}

func (s *ifStatement) Append(el *element) {}

func (s *ifStatement) AppendStatement(st statementIntf) {
	switch v := st.(type) {
	case *exp:
		s.Exp = v
	case *elseifStatement:
		s.ElseIfPart = append(s.ElseIfPart, v)
	case *elseStatement:
		s.ElsePart = v
	case *prefixexpStatement:
		return
		// default:
		// s.Body = append(s.Body, newBlock(st))
	}
}

func (s *ifStatement) GetBody(prevSt statementIntf, cur *element) statementIntf {
	if cur.Token.Type != nThen {
		return prevSt
	}

	if s.Body == nil {
		s.Body = new(body).New()
	}

	return s.Body
}

func (s *ifStatement) Format(c *Config, p printer, w io.Writer) error {
	if _, err := w.Write([]byte("if ")); err != nil {
		return err
	}

	if s.Exp != nil {
		if err := s.Exp.Format(c, p, w); err != nil {
			return err
		}
	}

	if _, err := w.Write([]byte(" then")); err != nil {
		return err
	}

	if err := newLine(w); err != nil {
		return err
	}

	if st, ok := s.Body.(*body); ok {
		pp := p
		pp.Pad = p.Pad + 4
		if err := st.Format(c, pp, w); err != nil {
			return err
		}
	}

	if _, err := w.Write([]byte("end")); err != nil {
		return err
	}

	return nil
}

type elseifStatement struct {
	Exp  *exp
	Body []Block
}

func (elseifStatement) New() statementIntf {
	return &elseifStatement{}
}

func (elseifStatement) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (elseifStatement) TypeOf() typeStatement {
	return tsIf
}

func (s *elseifStatement) IsEnd(prev, cur *element) bool {
	return cur.Token.Type == nEnd || cur.Token.Type == nElse
}

func (s *elseifStatement) Append(el *element) {}

func (s *elseifStatement) AppendStatement(st statementIntf) {
	if v, ok := st.(*exp); ok {
		s.Exp = v

		return
	}

	s.Body = append(s.Body, newBlock(st))
}

func (s *elseifStatement) GetBody(prevSt statementIntf, cur *element) statementIntf {
	return prevSt
}

type elseStatement struct {
	Body []Block
}

func (elseStatement) New() statementIntf {
	return &elseStatement{}
}

func (elseStatement) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (elseStatement) TypeOf() typeStatement {
	return tsIf
}

func (s *elseStatement) IsEnd(prev, cur *element) bool {
	return cur.Token.Type == nEnd
}

func (s *elseStatement) Append(el *element) {}

func (s *elseStatement) AppendStatement(st statementIntf) {
	s.Body = append(s.Body, newBlock(st))
}

func (s *elseStatement) GetBody(prevSt statementIntf, cur *element) statementIntf {
	return prevSt
}
