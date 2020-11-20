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

type tableStatement struct {
	List []field
}

func (tableStatement) New() statementIntf {
	return &tableStatement{}
}

func (tableStatement) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (tableStatement) TypeOf() typeStatement {
	return tsIf
}

func (s *tableStatement) IsEnd(prev, cur *element) bool {
	return cur.Token.Type == nEnd
}

func (s *tableStatement) Append(el *element) {}

func (s *tableStatement) AppendStatement(st statementIntf) {
	// s.Body = append(s.Body, newBlock(st))
}

func (s *tableStatement) GetBody(prevSt statementIntf, cur *element) statementIntf {
	return prevSt
}
