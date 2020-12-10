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
	"bytes"
	"io"
)

type commentStatement struct {
	Element *element
}

func (commentStatement) InnerStatement(prev, cur *element) (bool, statementIntf) {
	return false, nil
}

func (commentStatement) TypeOf() typeStatement {
	return tsNone
}

func (s *commentStatement) IsEnd(prev, cur *element) (bool, bool) {
	return false, true
}

func (s *commentStatement) Append(el *element) {
	if el.Token.Type == nComment {
		el.Token.Lexeme = bytes.TrimLeft(el.Token.Lexeme, "--")
		el.Token.Lexeme = bytes.TrimSpace(el.Token.Lexeme)
		el.Token.Value = string(el.Token.Lexeme)
	}
	s.Element = el
}

func (s *commentStatement) AppendStatement(st statementIntf) {}

func (s *commentStatement) GetBody(prevSt statementIntf, cur *element) statementIntf {
	return prevSt
}

func (s *commentStatement) GetStatement(prev, cur *element) statementIntf {
	return nil
}

func (s *commentStatement) Format(c *Config, p printer, w io.Writer) error {
	if _, err := w.Write([]byte("-- ")); err != nil {
		return err
	}

	if err := s.Element.Format(c, p, w); err != nil {
		return err
	}

	return nil
}
