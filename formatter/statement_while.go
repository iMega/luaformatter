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

type whileStatement struct {
	Exp *exp
	Do  *doStatement
}

func (whileStatement) New() statementIntf {
	return &whileStatement{}
}

func (whileStatement) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (whileStatement) TypeOf() typeStatement {
	return tsIf
}

func (s *whileStatement) IsEnd(prev, cur *element) (bool, bool) {
	return false, cur.Token.Type == nEnd
}

func (s *whileStatement) Append(el *element) {}

func (s *whileStatement) AppendStatement(st statementIntf) {
	switch v := st.(type) {
	case *exp:
		s.Exp = v
	case *doStatement:
		s.Do = v
	}
}

func (s *whileStatement) GetBody(prevSt statementIntf, cur *element) statementIntf {
	return prevSt
}
