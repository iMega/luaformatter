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

func (field) InnerStatement(prev, cur *element) (bool, statement) {
	return false, &exp{}
}

func (field) TypeOf() typeStatement {
	return tsField
}

func (s *field) IsEnd(prev, cur *element) (bool, bool) {
	if cur.Token.Type == nIn {
		return false, true
	}

	if cur.Token.Type == nClosingCurlyBracket {
		return false, true
	}

	if cur.Token.Type == nClosingSquareBracket {
		s.Square = true // exeption
	}

	return false, cur.Token.Type == nComma || cur.Token.Type == nDo
}

func (s *field) Append(el *element) {
	if el.Token.Type == nComment || el.Token.Type == nCommentLong {
		if s.Val == nil {
			s.Key.Append(el)

			return
		}

		s.Val.Append(el)
	}
}

func (s *field) AppendStatement(st statement) {
	v, ok := st.(*exp)
	if !ok {
		return
	}

	if s.Key == nil {
		s.Key = v

		return
	}

	if s.Key.Element == nil &&
		s.Key.Table == nil &&
		s.Key.Func == nil &&
		s.Key.Binop == nil &&
		s.Key.Unop == nil &&
		s.Key.Exp == nil &&
		s.Key.Prefixexp == nil {

		v.Comments = s.Key.Comments
		s.Key = v

		return
	}

	s.Val = v
}

func allFieldsStructAreNil(in ...interface{}) bool {
	for _, v := range in {
		if v != nil {
			return false
		}
	}

	return true
}

func (s *field) GetBody(prevSt statement, cur *element) statement {
	return prevSt
}

func (s *field) GetStatement(prev, cur *element) statement {
	if isExp(cur) {
		return &exp{}
	}

	return nil
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
