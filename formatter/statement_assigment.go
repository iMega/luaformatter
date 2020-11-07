package formatter

import "io"

type assignmentStatement struct {
	IsLocal   bool
	VarList   *explist
	HasEqPart bool
	Explist   *explist
}

func (assignmentStatement) New() statementIntf {
	return &assignmentStatement{}
}

func (assignmentStatement) InnerStatement(prev, cur *element) statementIntf {
	return &explist{}
}

func (assignmentStatement) TypeOf() typeStatement {
	return tsAssignment
}

func (s *assignmentStatement) IsEnd(prev, cur *element) bool {
	if s.HasEqPart && s.VarList != nil && s.Explist != nil {
		if len(s.VarList.List) == len(s.Explist.List) {
			return true
		}
	}

	return false
}

func (s *assignmentStatement) HasSyntax(el element) bool {
	return false
}

func (s *assignmentStatement) Append(el *element) {
	if el.Token.Type == nLocal {
		s.IsLocal = true

		return
	}

	if el.Token.Type == nAssign {
		s.HasEqPart = true

		return
	}

	// switch el.Token.Type {
	// case nNumber:
	// 	s.Explist.List = append(s.Explist.List, newExp(el))
	// }
}

func (s *assignmentStatement) AppendStatement(st statementIntf) {
	switch v := st.(type) {
	case *functionStatement:
		s.Explist.List = append(s.Explist.List, newExp(v))

	case *explist:
		if s.HasEqPart {
			s.Explist = v
		} else {
			s.VarList = v
		}
	}
}

func (s *assignmentStatement) GetBody(prevSt statementIntf, cur *element) statementIntf {
	return prevSt
}

func (s *assignmentStatement) Format(c *Config, p printer, w io.Writer) error {
	if s.IsLocal {
		if _, err := w.Write([]byte("local ")); err != nil {
			return err
		}
	}

	if err := s.VarList.Format(c, p, w); err != nil {
		return err
	}

	if s.HasEqPart {
		if _, err := w.Write([]byte(" = ")); err != nil {
			return err
		}
	}

	if err := s.Explist.Format(c, p, w); err != nil {
		return err
	}

	return nil
}
