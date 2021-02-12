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

type prefixexpStatement struct {
	Element          *element
	FuncCall         *funcCallStatement
	FieldAccessorExp *exp
	FieldAccessor    *element
	Prefixexp        *prefixexpStatement
	Comments         map[uint64]*element
	Enclosed         bool
	IsUnknow         bool
}

func (s *prefixexpStatement) InnerStatement(prev, cur *element) (bool, statement) {
	if cur.Token.Type == nSquareBracket {
		return false, &exp{}
	}

	if cur.Token.Type == nParentheses {
		if prev != nil && prev.Token.Type == nID {
			return false, &funcCallStatement{} // Element: cur
		}

		s.Enclosed = true

		return true, &exp{} // why?
	}

	if cur.Token.Type == nString {
		return false, &funcCallStatement{}
	}

	if cur.Token.Type == nCurlyBracket {
		return false, &funcCallStatement{}
	}

	if cur.Token.Type == nDot {
		if prev != nil && prev.Token.Type == nClosingParentheses {
			return true, &prefixexpStatement{}
		}
	}

	if cur.Token.Type == nColon {
		if prev != nil && prev.Token.Type == nClosingParentheses {
			return true, &prefixexpStatement{}
		}
	}

	return false, nil
}

func (s *prefixexpStatement) TypeOf() typeStatement {
	if s.IsUnknow {
		return tsUnknow
	}

	return tsPrefixexpStatement
}

func (s *prefixexpStatement) IsEnd(prev, cur *element) (bool, bool) {
	if cur.Token.Type == nString {
		return false, false
	}

	if cur.Token.Type == nCurlyBracket { // function_call {
		return false, false
	}

	if cur.Token.Type == nParentheses { // function call
		return false, false
	}

	if cur.Token.Type == nSquareBracket { // function call
		return false, false
	}

	if cur.Token.Type == nClosingSquareBracket { // function call
		return false, false
	}

	if prev != nil && prev.Token.Type == nClosingSquareBracket && cur.Token.Type == nClosingParentheses {
		// porsche_handler(vehicles["Porsche"])
		// vehicles["Porsche"] = nil
		return false, true
	}

	if cur.Token.Type == nClosingParentheses && s.Enclosed { // (a-b) and
		return true, true
	}

	if prev != nil && prev.Token.Type == nDot && cur.Token.Type == nID { // function call
		return false, false // .id
	}

	if prev != nil && prev.Token.Type == nColon && cur.Token.Type == nID { // function call
		return false, false // :id
	}

	if cur.Token.Type == nComma { // assignment statement
		return false, true
	}

	if cur.Token.Type == nDot { // prefixexpStatement
		return false, false
	}

	if prev != nil && prev.Token.Type == nID && cur.Token.Type == nAssign {
		return false, false
	}

	if prev != nil && prev.Token.Type == nClosingSquareBracket && cur.Token.Type == nAssign {
		return false, false
	}

	if prev != nil && prev.Token.Type == nParentheses && isExp(cur) {
		return false, false
	}

	if cur.Token.Type == nCommentLong {
		return false, false
	}

	return false, true //cur.Token.Type == nAssign
}

func (s *prefixexpStatement) Append(el *element) {
	if el.Token.Type == nSquareBracket {
		return
	}

	if el.Token.Type == nClosingSquareBracket {
		return
	}

	if el.Token.Type == nClosingParentheses {
		return
	}

	// note: before nDot
	if el.Token.Type == nCommentLong {
		if s.Comments == nil {
			s.Comments = make(map[uint64]*element)
		}

		s.Comments[uint64(len(s.Comments))] = el

		return
	}

	if el.Token.Type == nDot || el.Token.Type == nColon {
		s.FieldAccessor = el

		return
	}

	if el.Token.Type == nEnd {
		return
	}

	if s.FieldAccessor != nil {
		if s.FieldAccessor.Token.Type == nDot { // .id
			s.Element = s.FieldAccessor
			s.FieldAccessor = el
		}

		if s.FieldAccessor.Token.Type == nColon { // :id
			s.Element = s.FieldAccessor
			s.FieldAccessor = el
		}

		return
	}

	s.Element = el
}

func (s *prefixexpStatement) AppendStatement(st statement) {
	switch v := st.(type) {
	case *prefixexpStatement:
		s.Prefixexp = v
	case *exp:
		s.FieldAccessorExp = v
	case *funcCallStatement:
		v.Prefixexp = &prefixexpStatement{
			Element: s.Element,
		}
		s.Element = nil
		s.FuncCall = v
	}
}

func (s *prefixexpStatement) GetBody(prevSt statement, cur *element) statement {
	return prevSt
}

func (s *prefixexpStatement) GetStatement(prev, cur *element) statement {
	if cur.Token.Type == nAssign {
		return &assignmentStatement{}
	}

	if cur.Token.Type == nComma {
		return &assignmentStatement{}
	}

	if cur.Token.Type == nParentheses || cur.Token.Type == nString {
		return &funcCallStatement{
			// TODO Prefixexp: s,
		}
	}

	if cur.Token.Type == nCurlyBracket {
		return &funcCallStatement{}
	}

	if cur.Token.Type == nSquareBracket {
		return &prefixexpStatement{}
	}

	if cur.Token.Type == nDot {
		return &prefixexpStatement{}
	}

	return nil
}

func (s *prefixexpStatement) Format(c *Config, p printer, w io.Writer) error {
	if s.Comments != nil {
		for i := 0; i < len(s.Comments); i++ {
			if _, err := w.Write([]byte("--[[ ")); err != nil {
				return err
			}

			if err := s.Comments[uint64(i)].Format(c, p, w); err != nil {
				return err
			}

			if _, err := w.Write([]byte(" ]]")); err != nil {
				return err
			}
		}
	}

	if st := s.FuncCall; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	if st := s.Element; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	if st := s.FieldAccessorExp; st != nil {
		def := []byte("[")
		if s.Enclosed {
			def = []byte("(")
		}

		if _, err := w.Write(def); err != nil {
			return err
		}

		if err := st.Format(c, p, w); err != nil {
			return err
		}

		def = []byte("]")
		if s.Enclosed {
			def = []byte(")")
		}

		if _, err := w.Write(def); err != nil {
			return err
		}
	}

	if st := s.FieldAccessor; st != nil {
		// if _, err := w.Write([]byte(".")); err != nil {
		// 	return err
		// }

		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	if st := s.Prefixexp; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	return nil
}
