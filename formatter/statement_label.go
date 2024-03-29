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

type labelStatement struct {
	Element *element
}

func (labelStatement) InnerStatement(prev, cur *element) (bool, statement) {
	return false, nil
}

func (labelStatement) TypeOf() typeStatement {
	return tsLabel
}

func (s *labelStatement) IsEnd(prev, cur *element) (bool, bool) {
	return false, true
}

func (s *labelStatement) Append(el *element) {
	el.Token.Lexeme = bytes.Trim(el.Token.Lexeme, "::")
	el.Token.Lexeme = bytes.TrimSpace(el.Token.Lexeme)
	el.Token.Value = string(el.Token.Lexeme)

	s.Element = el
}

func (s *labelStatement) AppendStatement(st statement) {}

func (s *labelStatement) GetBody(prevSt statement, cur *element) statement {
	return prevSt
}

func (s *labelStatement) GetStatement(prev, cur *element) statement {
	return nil
}

func (s *labelStatement) Format(c *Config, p printer, w io.Writer) error {
	return nil
}
