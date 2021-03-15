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

type doStatement struct {
	Body statement
}

func (doStatement) InnerStatement(prev, cur *element) (bool, statement) {
	return false, nil
}

func (doStatement) TypeOf() typeStatement {
	return tsNone
}

func (s *doStatement) IsEnd(prev, cur *element) (bool, bool) {
	if cur.Token.Type == nEnd {
		return true, true
	}

	return false, false
}

func (s *doStatement) Append(el *element) {}

func (s *doStatement) AppendStatement(st statement) {}

func (s *doStatement) GetBody(prevSt statement, cur *element) statement {
	if s.Body == nil {
		s.Body = new(body).New()
	}

	return s.Body
}

func (s *doStatement) GetStatement(prev, cur *element) statement {
	return nil
}

func (s *doStatement) Format(c *Config, p printer, w io.Writer) error {
	if err := writeKeyword(c, nDo, w); err != nil {
		return err
	}

	st, ok := s.Body.(*body)
	if ok {
		ip := p
		ip.Pad = p.Pad + c.IndentSize

		if err := st.Format(c, ip, w); err != nil {
			return err
		}

		if err := newLine(w); err != nil {
			return err
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

	if err := writeKeyword(c, nEnd, w); err != nil {
		return err
	}

	return nil
}
