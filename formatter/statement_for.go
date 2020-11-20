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

type forNumericalStatement struct {
	VarPart   *field
	LimitPart *exp
	StepPart  *exp
	Body      statementIntf
}

func (forNumericalStatement) New() statementIntf {
	return &forNumericalStatement{}
}

func (forNumericalStatement) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (forNumericalStatement) TypeOf() typeStatement {
	return tsIf
}

func (s *forNumericalStatement) IsEnd(prev, cur *element) bool {
	return cur.Token.Type == nEnd
}

func (s *forNumericalStatement) Append(el *element) {}

func (s *forNumericalStatement) AppendStatement(st statementIntf) {
	switch v := st.(type) {
	case *field:
		s.VarPart = v
	case *exp:
		if s.LimitPart == nil {
			s.LimitPart = v
		} else {
			s.StepPart = v
		}
	}
}

func (s *forNumericalStatement) GetBody(prevSt statementIntf, cur *element) statementIntf {
	if cur.Token.Type != nDo {
		return prevSt
	}

	if s.Body == nil {
		s.Body = new(body).New()
	}

	return s.Body
}

type forGenericStatement struct {
	IDStatement *element
	Namelist    namelist
	InElement   element
	Explist     explist
	doStatement
}
