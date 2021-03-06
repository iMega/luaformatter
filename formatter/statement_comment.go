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
	Element   *element
	IsNewline bool
}

func (commentStatement) InnerStatement(prev, cur *element) (bool, statement) {
	return false, nil
}

func (commentStatement) TypeOf() typeStatement {
	return tsComment
}

func (s *commentStatement) IsEnd(prev, cur *element) (bool, bool) {
	return false, true
}

func (s *commentStatement) Append(el *element) {
	s.Element = el
}

func (s *commentStatement) AppendStatement(st statement) {}

func (s *commentStatement) GetBody(prevSt statement, cur *element) statement {
	return prevSt
}

func (s *commentStatement) GetStatement(prev, cur *element) statement {
	return nil
}

func (s *commentStatement) Format(c *Config, p printer, w io.Writer) error {
	prefix := []byte("-- ")
	if len(s.Element.Token.Lexeme) == 0 || bytes.HasPrefix(s.Element.Token.Lexeme, []byte("---")) {
		prefix = []byte("--")
	}

	if _, err := w.Write(prefix); err != nil {
		return err
	}

	if err := s.Element.Format(c, p, w); err != nil {
		return err
	}

	return nil
}
