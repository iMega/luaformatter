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
	Body statementIntf
}

func (doStatement) InnerStatement(prev, cur *element) (bool, statementIntf) {
	return false, nil
}

func (doStatement) TypeOf() typeStatement {
	return tsIf
}

func (s *doStatement) IsEnd(prev, cur *element) (bool, bool) {
	return false, cur.Token.Type == nEnd
}

func (s *doStatement) Append(el *element) {}

func (s *doStatement) AppendStatement(st statementIntf) {}

func (s *doStatement) GetBody(prevSt statementIntf, cur *element) statementIntf {
	if s.Body == nil {
		s.Body = new(body).New()
	}

	return s.Body
}

func (s *doStatement) GetStatement(prev, cur *element) statementIntf {
	return nil
}

func (s *doStatement) Format(c *Config, p printer, w io.Writer) error {
	if _, err := w.Write([]byte("do")); err != nil {
		return err
	}

	if err := newLine(w); err != nil {
		return err
	}

	if st, ok := s.Body.(*body); ok {
		ip := p
		ip.Pad = p.Pad + 4

		if err := st.Format(c, ip, w); err != nil {
			return err
		}

		if err := newLine(w); err != nil {
			return err
		}
	}

	return nil
}
