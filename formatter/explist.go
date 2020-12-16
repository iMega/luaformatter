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

type explist struct {
	List []*exp // separator ,
}

func (explist) InnerStatement(prev, cur *element) (bool, statementIntf) {
	return false, &exp{}
}

func (explist) TypeOf() typeStatement {
	return tsExpList
}

func (s *explist) IsEnd(prev, cur *element) (bool, bool) {
	if cur.Token.Type == nComma || prev.Token.Type == nComma {
		return false, false
	}

	if cur.Resolved {
		return false, false
	}

	return false, true
}

func (s *explist) Append(el *element) {}

func (s *explist) AppendStatement(st statementIntf) {
	if v, ok := st.(*exp); ok {
		s.List = append(s.List, v)
	}
}

func (s *explist) GetBody(prevSt statementIntf, cur *element) statementIntf {
	return prevSt
}

func (s *explist) GetStatement(prev, cur *element) statementIntf {
	if prev != nil && prev.Token.Type == nComma && isExp(cur) {
		return &exp{}
	}

	return nil
}

func (s *explist) Format(c *Config, p printer, w io.Writer) error {
	l := len(s.List)

	for idx, e := range s.List {
		if err := e.Format(c, p, w); err != nil {
			return err
		}

		if idx < l-1 {
			if _, err := w.Write([]byte(", ")); err != nil {
				return err
			}
		}
	}

	return nil
}
