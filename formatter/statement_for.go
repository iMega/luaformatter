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

type forStatement struct {
	FieldList *fieldlist
	Explist   *explist
	Body      statementIntf
}

func (forStatement) New() statementIntf {
	return &forStatement{}
}

func (forStatement) InnerStatement(prev, cur *element) statementIntf {
	return &fieldlist{}
}

func (forStatement) TypeOf() typeStatement {
	return tsIf
}

func (s *forStatement) IsEnd(prev, cur *element) bool {
	return cur.Token.Type == nEnd
}

func (s *forStatement) Append(el *element) {}

func (s *forStatement) AppendStatement(st statementIntf) {
	switch v := st.(type) {
	case *fieldlist:
		s.FieldList = v
	case *explist:
		s.Explist = v
	}
}

func (s *forStatement) GetBody(prevSt statementIntf, cur *element) statementIntf {
	if cur.Token.Type != nDo {
		return prevSt
	}

	if s.Body == nil {
		s.Body = new(body).New()
	}

	return s.Body
}
