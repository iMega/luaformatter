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
	IsLocal     bool
	Name        *element
	Parlist     *explist
	Body        statementIntf
	IsAnonymous bool
}

func (functionStatement) InnerStatement(prev, cur *element) statementIntf {
	return nil
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

	if el.Token.Type == nID {
		s.Name = el

		return
	}

	if s.Name == nil && el.Token.Type == nParentheses {
		s.IsAnonymous = true
	}
}

func (s *functionStatement) AppendStatement(st statementIntf) {
	if s.Parlist == nil {
		if v, ok := st.(*explist); ok {
			s.Parlist = v

			return
		}
	}
}

func (s *functionStatement) GetBody(prevSt statementIntf, cur *element) statementIntf {
	if cur.Token.Type != nClosingParentheses {
		return prevSt
	}

	if s.Body == nil {
		s.Body = new(body).New()
	}

	return s.Body
}

func (s *functionStatement) GetStatement(prev, cur *element) statementIntf {
	if prev != nil && prev.Token.Type == nParentheses {
		return &explist{}
	}

	return nil
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

	// if _, err := w.Write([]byte("function")); err != nil {
	// 	return err
	// }

	if s.Name != nil {
		if err := space(w); err != nil {
			return err
		}

		if err := s.Name.Format(c, p, w); err != nil {
			return err
		}
	}

	if _, err := w.Write([]byte("(")); err != nil {
		return err
	}

	if st := s.Parlist; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	if _, err := w.Write([]byte(")")); err != nil {
		return err
	}

	if err := newLine(w); err != nil {
		return err
	}

	inner := printer{
		Pad: p.Pad + c.IndentSize,
	}
	st, ok := s.Body.(*body)
	if ok {
		if err := st.Format(c, inner, w); err != nil {
			return err
		}
	}

	if st.Qty > 0 {
		// a = function()
		// end
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
