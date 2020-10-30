package formatter

import "io"

type explist struct {
	List []*exp // separator ,
}

func (explist) New() statementIntf {
	return &explist{}
}

func (explist) InnerStatement(prev, cur *element) statementIntf {
	return &exp{}
}

func (explist) TypeOf() typeStatement {
	return tsExpList
}

func (s *explist) IsEnd(prev, cur *element) bool {
	if cur.Token.Type == nComma || prev.Token.Type == nComma {
		return false
	}

	if cur.Resolved {
		return false
	}

	return true
}

func (s *explist) HasSyntax(el element) bool {
	return false
}

func (s *explist) Append(el *element) {}

func (s *explist) AppendStatement(st statementIntf) {
	if v, ok := st.(*exp); ok {
		s.List = append(s.List, v)
	}
}

func (s *explist) Format(c *Config, w io.Writer) error {
	l := len(s.List)
	for idx, e := range s.List {
		if err := e.Format(c, w); err != nil {
			return err
		}

		if idx < l-1 {
			if _, err := w.Write([]byte(", ")); err != nil {
				return err
			}
		}
	}

	return nil
}
