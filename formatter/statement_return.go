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

type returnStatement struct {
	Explist *explist
}

func (returnStatement) InnerStatement(prev, cur *element) (bool, statement) {
	return false, nil
}

func (returnStatement) TypeOf() typeStatement {
	return tsNone
}

func (s *returnStatement) IsEnd(prev, cur *element) (bool, bool) {
	return false, !isExp(cur)
}

func (s *returnStatement) Append(el *element) {
	if el == nil || el.Token.Type == nReturn {
		return
	}
}

func (s *returnStatement) AppendStatement(st statement) {
	el, ok := st.(*explist)
	if !ok {
		return
	}

	s.Explist = el
}

func (s *returnStatement) GetBody(prevSt statement, cur *element) statement {
	return prevSt
}

func (s *returnStatement) GetStatement(prev, cur *element) statement {
	if isExp(cur) {
		return &explist{}
	}

	return nil
}

func (s *returnStatement) Format(c *Config, p printer, w io.Writer) error {
	if _, err := w.Write([]byte("return")); err != nil {
		return err
	}

	if st := s.Explist; st != nil {
		if _, err := w.Write([]byte(" ")); err != nil {
			return err
		}

		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	return nil
}
