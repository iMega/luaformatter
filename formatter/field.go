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

type field struct {
	Key    *exp
	Val    *exp
	Square bool
}

func (field) New() statementIntf {
	return &field{}
}

func (field) InnerStatement(prev, cur *element) statementIntf {
	return &exp{}
}

func (field) TypeOf() typeStatement {
	return tsField
}

func (s *field) IsEnd(prev, cur *element) (bool, bool) {
	if cur.Token.Type == nIn {
		return false, true
	}

	return false, cur.Token.Type == nComma || cur.Token.Type == nDo
}

func (s *field) Append(el *element) {
	if el.Token.Type == nSquareBracket {
		s.Square = true
	}
}

func (s *field) AppendStatement(st statementIntf) {
	v, ok := st.(*exp)
	if !ok {
		return
	}

	if s.Key == nil {
		s.Key = v

		return
	}

	s.Val = v
}

func (s *field) GetBody(prevSt statementIntf, cur *element) statementIntf {
	return prevSt
}

func (s *field) Format(c *Config, p printer, w io.Writer) error {
	if s.Square {
		if _, err := w.Write([]byte("[")); err != nil {
			return err
		}
	}

	if err := s.Key.Format(c, p, w); err != nil {
		return err
	}

	if s.Square {
		if _, err := w.Write([]byte("]")); err != nil {
			return err
		}
	}

	if s.Val == nil {
		return nil
	}

	if _, err := w.Write([]byte(" = ")); err != nil {
		return err
	}

	if err := s.Val.Format(c, p, w); err != nil {
		return err
	}

	return nil
}
