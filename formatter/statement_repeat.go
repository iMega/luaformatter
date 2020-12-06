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

type repeatStatement struct {
	Body statementIntf
	Exp  *exp
}

func (repeatStatement) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (repeatStatement) TypeOf() typeStatement {
	return tsRepeat
}

func (s *repeatStatement) IsEnd(prev, cur *element) (bool, bool) {
	if prev != nil && prev.Token.Type == nRepeat {
		return false, false
	}

	if cur.Token.Type == nUntil {
		return false, false
	}

	if prev != nil && prev.Token.Type == nUntil {
		return false, false
	}

	return false, true
}

func (s *repeatStatement) Append(el *element) {}

func (s *repeatStatement) AppendStatement(st statementIntf) {
	if v, ok := st.(*exp); ok {
		s.Exp = v
	}
}

func (s *repeatStatement) GetBody(prevSt statementIntf, cur *element) statementIntf {
	if s.Body == nil {
		s.Body = new(body).New()
	}

	return s.Body
}

func (s *repeatStatement) GetStatement(prev, cur *element) statementIntf {
	if prev.Token.Type == nUntil {
		return &exp{}
	}

	return nil
}
