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

type explist struct {
	List []*exp // separator ,
}

func (explist) InnerStatement(prev, cur *element) (bool, statement) {
	return false, &exp{}
}

func (explist) TypeOf() typeStatement {
	return tsNone
}

func (s *explist) IsEnd(prev, cur *element) (bool, bool) {
	if cur.Token.Type == nComma || prev.Token.Type == nComma {
		return false, false
	}

	if cur.Resolved {
		return false, false
	}

	return false, true
}

func (s *explist) Append(el *element) {}

func (s *explist) AppendStatement(st statement) {
	if v, ok := st.(*exp); ok {
		s.List = append(s.List, v)
	}
}

func (s *explist) GetBody(prevSt statement, cur *element) statement {
	return prevSt
}

func (s *explist) GetStatement(prev, cur *element) statement {
	if prev != nil && prev.Token.Type == nComma && isExp(cur) {
		return &exp{}
	}

	return nil
}

func (s *explist) Format(c *Config, p printer, w io.Writer) error {
	isInLine := s.isInline(c, p, w)
	sep := []byte(", ")
	if !isInLine {
		sep = []byte(",")
		p.Pad += c.IndentSize
	}

	for i := 0; i < len(s.List); i++ {
		if !isInLine {
			if err := newLine(w); err != nil {
				return err
			}

			if err := p.WritePad(w); err != nil {
				return err
			}
		}

		if err := s.List[i].Format(c, p, w); err != nil {
			return err
		}

		if i < len(s.List)-1 {
			if _, err := w.Write(sep); err != nil {
				return err
			}
		}
	}

	if !isInLine {
		if err := newLine(w); err != nil {
			return err
		}

		p.Pad -= c.IndentSize
		if err := p.WritePad(w); err != nil {
			return err
		}
	}

	return nil
}

func (s *explist) isInline(c *Config, p printer, w io.Writer) bool {
	if len(s.List) == 1 {
		return true
	}

	for _, item := range s.List {
		if item.Table != nil || item.Func != nil {
			return false
		}
	}

	return true
}
