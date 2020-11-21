package formatter

import "io"

type fieldlist struct {
	List []*field
}

func (fieldlist) New() statementIntf {
	return &fieldlist{}
}

func (fieldlist) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (fieldlist) TypeOf() typeStatement {
	return tsFieldList
}

func (s *fieldlist) IsEnd(prev, cur *element) bool {
	if cur.Token.Type == nComma || prev.Token.Type == nComma {
		return false
	}

	if cur.Token.Type == nIn {
		return true
	}

	if cur.Token.Type == nDo || cur.Token.Type == nClosingCurlyBracket {
		return true
	}

	return false
}

func (s *fieldlist) Append(el *element) {}

func (s *fieldlist) AppendStatement(st statementIntf) {
	if v, ok := st.(*field); ok {
		s.List = append(s.List, v)
	}
}

func (s *fieldlist) GetBody(prevSt statementIntf, cur *element) statementIntf {
	return prevSt
}

func (s *fieldlist) Format(c *Config, p printer, w io.Writer) error {
	for i, v := range s.List {
		if err := v.Format(c, p, w); err != nil {
			return err
		}

		if i < len(s.List)-1 {
			if _, err := w.Write([]byte(", ")); err != nil {
				return err
			}
		}
	}

	return nil
}
