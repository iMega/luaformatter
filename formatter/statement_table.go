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

type tableStatement struct {
	FieldList *fieldlist
}

func (tableStatement) InnerStatement(prev, cur *element) (bool, statementIntf) {
	return false, &fieldlist{}
}

func (tableStatement) TypeOf() typeStatement {
	return tsTable
}

func (s *tableStatement) IsEnd(prev, cur *element) (bool, bool) {
	if cur.Token.Type == nClosingCurlyBracket {
		return true, true
	}

	return false, false
}

func (s *tableStatement) Append(el *element) {}

func (s *tableStatement) AppendStatement(st statementIntf) {
	if v, ok := st.(*fieldlist); ok {
		s.FieldList = v
	}
}

func (s *tableStatement) GetBody(prevSt statementIntf, cur *element) statementIntf {
	return prevSt
}

func (s *tableStatement) GetStatement(prev, cur *element) statementIntf {
	return nil
}

func (s *tableStatement) Format(c *Config, p printer, w io.Writer) error {
	if _, err := w.Write([]byte("{")); err != nil {
		return err
	}

	if s.FieldList != nil {
		if len(s.FieldList.List) > 1 {
			if err := newLine(w); err != nil {
				return err
			}

			p.ParentStatement = s.TypeOf()
			p.Pad = p.Pad + c.IndentSize
		}

		if err := s.FieldList.Format(c, p, w); err != nil {
			return err
		}
	}

	if _, err := w.Write([]byte("}")); err != nil {
		return err
	}

	return nil
}
