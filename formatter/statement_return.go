package formatter

import "io"

type returnStatement struct {
	Explist *explist
}

func (returnStatement) New() statementIntf {
	return &returnStatement{}
}

func (returnStatement) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (returnStatement) TypeOf() typeStatement {
	return tsReturn
}

func (s *returnStatement) IsEnd(prev, cur *element) bool {
	if nReturn == cur.Token.Type {
		return false
	}

	branch := getsyntax(syntax, tokenID(nReturn))
	_, ok := branch[cur.Token.Type]

	return !ok
}

func (s *returnStatement) HasSyntax(el element) bool {
	return false
}

func (s *returnStatement) Append(el *element) {
	if el == nil || el.Token.Type == nReturn {
		return
	}

	s.Explist.List = append(s.Explist.List, newExp(el))
}

func (s *returnStatement) AppendStatement(st statementIntf) {
	el, ok := st.(*explist)
	if !ok {
		return
	}

	s.Explist = el
}

func (s *returnStatement) Format(c *Config, w io.Writer) error {
	if _, err := w.Write([]byte("return")); err != nil {
		return err
	}

	if st := s.Explist; st != nil {
		if _, err := w.Write([]byte(" ")); err != nil {
			return err
		}

		if err := st.Format(c, w); err != nil {
			return err
		}
	}

	return nil
}
