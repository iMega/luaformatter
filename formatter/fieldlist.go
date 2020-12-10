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

type fieldlist struct {
	List []*field
}

func (fieldlist) InnerStatement(prev, cur *element) (bool, statementIntf) {
	// fieldlist always returns a field. need will add
	// it will need to be added to the innerStatement.
	// return &field{}
	return true, nil
}

func (fieldlist) TypeOf() typeStatement {
	return tsFieldList
}

func (s *fieldlist) IsEnd(prev, cur *element) (bool, bool) {
	if cur.Token.Type == nIn {
		return false, true
	}

	if cur.Token.Type == nDo {
		return false, true
	}

	if cur.Token.Type == nClosingCurlyBracket {
		return false, true
	}

	return false, false
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

func (s *fieldlist) GetStatement(prev, cur *element) statementIntf {
	if cur.Token.Type == nID {
		return &field{} // fieldlist always returns a field. need will add
		// it will need to be added to the innerStatement
	}

	if cur.Token.Type == nSquareBracket {
		return &field{Square: true}
	}

	if prev != nil && prev.Token.Type == nComma {
		return &field{}
	}
	// return &field{}
	return nil
}

func (s *fieldlist) Format(c *Config, p printer, w io.Writer) error {
	for i, v := range s.List {
		if p.ParentStatement == tsTable {
			if err := p.WritePad(w); err != nil {
				return err
			}
		}

		if err := v.Format(c, p, w); err != nil {
			return err
		}

		if p.ParentStatement != tsTable {
			if i < len(s.List)-1 {
				if _, err := w.Write([]byte(", ")); err != nil {
					return err
				}
			}
		}

		if p.ParentStatement == tsTable {
			if _, err := w.Write([]byte(",")); err != nil {
				return err
			}

			if err := newLine(w); err != nil {
				return err
			}
		}
	}

	return nil
}
