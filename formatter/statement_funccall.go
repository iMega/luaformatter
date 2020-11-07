package formatter

import (
	"io"
)

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

func (s *funcCallStatement) GetBody(prevSt statementIntf, cur *element) statementIntf {
	return prevSt
}

func (s *funcCallStatement) Format(c *Config, p printer, w io.Writer) error {
	if st := s.Prefixexp; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	if _, err := w.Write([]byte("(")); err != nil {
		return err
	}

	if st := s.Explist; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	if _, err := w.Write([]byte(")")); err != nil {
		return err
	}

	return nil
}
