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

import (
	"io"
)

type functionStatement struct {
	FuncCall    *funcCallStatement
	Body        statement
	IsLocal     bool
	IsAnonymous bool
}

func (functionStatement) InnerStatement(prev, cur *element) (bool, statement) {
	return false, nil
}

func (functionStatement) TypeOf() typeStatement {
	return tsFunction
}

func (s *functionStatement) IsEnd(prev, cur *element) (bool, bool) {
	if cur.Token.Type == nEnd {
		cur.Resolved = true

		return false, true
	}

	return false, false
}

func (s *functionStatement) Append(el *element) {
	if el.Token.Type == nLocal {
		s.IsLocal = true

		return
	}
}

func (s *functionStatement) AppendStatement(st statement) {
	if v, ok := st.(*funcCallStatement); ok {
		s.FuncCall = v

		return
	}

	if v, ok := st.(*prefixexpStatement); ok {
		s.FuncCall = &funcCallStatement{
			Prefixexp: v,
		}
	}
}

func (s *functionStatement) GetBody(prevSt statement, cur *element) statement {
	if s.FuncCall == nil {
		return prevSt
	}

	if s.Body == nil {
		s.Body = new(body).New()
	}

	return s.Body
}

func (s *functionStatement) GetStatement(prev, cur *element) statement {
	if prev != nil && prev.Token.Type == nParentheses {
		return &explist{}
	}

	if s.FuncCall == nil && cur.Token.Type == nParentheses {
		s.IsAnonymous = true

		return &funcCallStatement{}
	}

	return &prefixexpStatement{}
}

func (s *functionStatement) Format(c *Config, p printer, w io.Writer) error {
	if s.IsLocal {
		if _, err := w.Write([]byte("local ")); err != nil {
			return err
		}
	}

	if err := writeKeyword(c, nFunction, w); err != nil {
		return err
	}

	if !s.IsAnonymous {
		if err := space(w); err != nil {
			return err
		}
	}

	if st := s.FuncCall; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	inner := printer{
		Pad: p.Pad + c.IndentSize,
	}

	st, ok := s.Body.(*body)
	if ok {
		if err := st.Format(c, inner, w); err != nil {
			return err
		}

		if st.Qty > 0 {
			// a = function()
			// end
			if err := newLine(w); err != nil {
				return err
			}
		}
	}

	if st == nil || len(st.Blocks) == 0 {
		if err := newLine(w); err != nil {
			return err
		}
	}

	if err := p.WritePad(w); err != nil {
		return err
	}

	if _, err := w.Write([]byte("end")); err != nil {
		return err
	}

	return nil
}
