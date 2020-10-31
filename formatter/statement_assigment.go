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
	return false
}

func (s *assignmentStatement) HasSyntax(el element) bool {
	return false
}

func (s *assignmentStatement) Append(el *element) {
	if el.Token.Type == nAssign {
		s.HasEqPart = true

		return
	}

	// if s.EqPart == nil {
	// 	s.Namelist = append(s.Namelist, el)
	// 	return
	// }

	switch el.Token.Type {
	case nNumber:
		s.Explist.List = append(s.Explist.List, newExp(el))
	}
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

func (s *assignmentStatement) Format(c *Config, p printer, w io.Writer) error {
	return nil
}
