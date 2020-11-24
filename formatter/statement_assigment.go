// Copyright © 2020 Dmitry Stoletov <info@imega.ru>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

func (s *assignmentStatement) IsEnd(prev, cur *element) (bool, bool) {
	if s.HasEqPart && s.VarList != nil && s.Explist != nil {
		if len(s.VarList.List) == len(s.Explist.List) {
			return false, true
		}
	}

	if cur.Token.Type == nAssign {
		return false, false
	}

	return false, !isExp(cur)
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
	// case *functionStatement:
	// s.Explist.List = append(s.Explist.List, newExp(v))

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

	if st := s.Explist; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	return nil
}
