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

type field struct {
	Key   *exp
	Val   *exp
	Sqare bool
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

func (s *field) IsEnd(prev, cur *element) bool {
	return cur.Token.Type == nComma
}

func (s *field) Append(el *element) {}

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
