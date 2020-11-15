package formatter

import (
	"io"
)

type functionStatement struct {
	IsLocal     bool
	Name        *element
	Parlist     *explist
	Body        statementIntf
	IsAnonymous bool
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
	if el.Token.Type == nLocal {
		s.IsLocal = true

		return
	}

	if el.Token.Type == nID {
		s.Name = el

		return
	}

	if s.Name == nil && el.Token.Type == nParentheses {
		s.IsAnonymous = true
	}
}

func (s *functionStatement) AppendStatement(st statementIntf) {
	if s.Parlist == nil {
		if v, ok := st.(*explist); ok {
			s.Parlist = v

			return
		}
	}
}

func (s *functionStatement) GetBody(prevSt statementIntf, cur *element) statementIntf {
	if cur.Token.Type != nClosingParentheses {
		return prevSt
	}

	if s.Body == nil {
		s.Body = new(body).New()
	}

	return s.Body
}

func (s *functionStatement) Format(c *Config, p printer, w io.Writer) error {
	if s.IsLocal {
		if _, err := w.Write([]byte("local ")); err != nil {
			return err
		}
	}

	if _, err := w.Write([]byte("function")); err != nil {
		return err
	}

	if s.Name != nil {
		if err := space(w); err != nil {
			return err
		}

		if err := s.Name.Format(c, p, w); err != nil {
			return err
		}
	}

	if _, err := w.Write([]byte("(")); err != nil {
		return err
	}

	if st := s.Parlist; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	if _, err := w.Write([]byte(")")); err != nil {
		return err
	}

	if err := newLine(w); err != nil {
		return err
	}

	inner := printer{
		Pad: p.Pad + c.IndentSize,
	}
	if st, ok := s.Body.(*body); ok {
		if err := st.Format(c, inner, w); err != nil {
			return err
		}
	}

	if err := newLine(w); err != nil {
		return err
	}

	if err := p.WritePad(w); err != nil {
		return err
	}

	if _, err := w.Write([]byte("end")); err != nil {
		return err
	}

	return nil
}
