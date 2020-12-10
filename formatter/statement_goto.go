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

type gotoStatement struct {
	Element *element
}

func (gotoStatement) InnerStatement(prev, cur *element) (bool, statementIntf) {
	return false, nil
}

func (gotoStatement) TypeOf() typeStatement {
	return tsIf
}

func (s *gotoStatement) IsEnd(prev, cur *element) (bool, bool) {
	return false, s.Element != nil
}

func (s *gotoStatement) Append(el *element) {
	if el.Token.Type == nGoto {
		return
	}

	s.Element = el
}

func (s *gotoStatement) AppendStatement(st statementIntf) {}

func (s *gotoStatement) GetBody(prevSt statementIntf, cur *element) statementIntf {
	return prevSt
}

func (s *gotoStatement) GetStatement(prev, cur *element) statementIntf {
	return nil
}
