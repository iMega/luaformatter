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

// Parse code.
func parse(code []byte) (*document, error) {
	var (
		prevElement      *element
		curElement       *element
		currentStatement statementIntf
		currentBody      statementIntf

		chainSt = &chainStatments{}
	)

	s, err := newScanner(code)
	if err != nil {
		return nil, err
	}

	doc := NewDocument()
	b := new(body).New()
	doc.Bod = b
	chainSt.Append(b)
	currentBody = b
	currentStatement = b

	for s.Next() {
		el, err := s.Scan()
		if err != nil {
			return nil, err
		}

		curElement = &el

		// if prevElement != nil {
		// 	fmt.Printf("%s%s %s%s = ", mMagenta, TokenIDs[prevElement.Token.Type], prevElement.Token.Value, defaultStyle)
		// }

		// fmt.Printf("%s%s %s%s\n", mMagenta, TokenIDs[el.Token.Type], el.Token.Value, defaultStyle)

		for isBlockEnd, ok := currentStatement.IsEnd(prevElement, curElement); ok; isBlockEnd, ok = currentStatement.IsEnd(prevElement, curElement) {
			cs := chainSt.ExtractPrev()
			if cs == nil {
				currentBody = doc.Bod
				chainSt.Append(currentBody)
				currentStatement = currentBody

				break
			}

			currentStatement = cs

			if isBlockEnd {
				break
			}
		}

		b := currentStatement.GetBody(chainSt.GetLastBody(), curElement)
		if b != currentBody {
			currentBody = b
			if b != chainSt.First() {
				chainSt.Append(b)
				currentStatement = b
			}
		}

		s := syntax
		// prefixexp assignment or function call
		if curElement.Token.Type == nID && currentStatement.TypeOf() == tsBody {
			s = map[tokenID]branch{
				nID: {
					nThis:        &prefixexpStatement{},
					nParentheses: &funcCallStatement{},
				},
			}
			if prevElement != nil && prevElement.Token.Type == nLocal {
				s = syntax
			}
		}

		// global assignment statement
		if currentStatement.TypeOf() == tsPrefixexpStatement && curElement.Token.Type == nComma {
			s = map[tokenID]branch{
				nComma: {
					nThis: &assignmentStatement{},
				},
			}
		}

		// global assignment statement with one var
		if currentStatement.TypeOf() == tsPrefixexpStatement && curElement.Token.Type == nAssign {
			s = map[tokenID]branch{
				nAssign: {
					nThis: &assignmentStatement{},
				},
			}
		}

		if curElement.Token.Type == nParentheses && (prevElement != nil && prevElement.Token.Type == nID) {
			s = map[tokenID]branch{
				nID: {
					nParentheses: &funcCallStatement{},
				},
			}
			if currentStatement.TypeOf() == tsFunction {
				s = syntax
			}
		}

		if curElement.Token.Type == nString && (prevElement != nil && prevElement.Token.Type == nID) {
			s = map[tokenID]branch{
				nID: {
					nString:       &funcCallStatement{}, //local base = require "resty.core.base"
					nCurlyBracket: &funcCallStatement{}, //local base = require {}
				},
			}
			if currentStatement.TypeOf() == tsExp {
				s = map[tokenID]branch{
					nID: {
						nString:       &prefixexpStatement{}, // id = id "string"
						nCurlyBracket: &prefixexpStatement{},
					},
				}
			}
		}

		if curElement.Token.Type == nCurlyBracket && (prevElement != nil && prevElement.Token.Type == nID) {
			s = map[tokenID]branch{
				nID: {
					nCurlyBracket: &funcCallStatement{}, //local base = require {}
				},
			}
			if currentStatement.TypeOf() == tsExp {
				s = map[tokenID]branch{
					nID: {
						nString:       &prefixexpStatement{}, // id = id "string"
						nCurlyBracket: &prefixexpStatement{},
					},
				}
			}
		}

		if curElement.Token.Type == nParentheses && currentStatement.TypeOf() == tsExp {
			s = map[tokenID]branch{
				nParentheses: {
					nThis: &prefixexpStatement{},
					// nParentheses: &funcCallStatement{},
				},
			}
		}

		// if curElement.Token.Type == nCurlyBracket && currentStatement.TypeOf() == tsExp {
		// 	s = map[tokenID]branch{
		// 		nCurlyBracket: {
		// 			nThis: &prefixexpStatement{},
		// 		},
		// 	}
		// }

		if currentStatement.TypeOf() == tsField {
			s = map[tokenID]branch{
				nAssign: {
					nID:     &exp{},
					nNumber: &exp{},
					nString: &exp{}, // = "string"
				},
			}
		}

		if currentStatement.TypeOf() == tsFieldList {
			s = map[tokenID]branch{
				nComma: {
					nID:            &field{},
					nNumber:        &field{},
					nSubtraction:   &field{},
					nSquareBracket: &field{},
				},
				nFor: {
					nID: &field{},
				},
				nCurlyBracket: {
					nID:            &field{},
					nSquareBracket: &field{}, // table = {["and"] = v}
				},
			}
		}

		if currentStatement != nil && prevElement != nil {
			if currentStatement.TypeOf() == tsFunction && prevElement.Token.Type == nParentheses {
				s = map[tokenID]branch{
					nParentheses: {
						nID: &explist{},
					},
				}
			}
		}

		if st := getStatement(s, prevElement, curElement); st != nil {
			var assignmentWithOneVar statementIntf
			isPrefixexpConvertAssignment := false

			if st.TypeOf() == tsAssignment && prevElement.Token.Type == nLocal { // local a
				st.Append(prevElement)
			}

			if st.TypeOf() == tsAssignment && curElement.Token.Type == nAssign { // a = 1
				assignmentWithOneVar = st
			}

			if st.TypeOf() == tsField && curElement.Token.Type == nSquareBracket { // ["and"] = 100,
				st.Append(curElement)
			}

			if st.TypeOf() == tsFunction { // local function a()
				if prevElement != nil && prevElement.Token.Type == nLocal {
					st.Append(prevElement)
				}
			}

			if currentStatement == nil {
				chainSt.Append(st)

				if prevElement != nil {
					st.Append(prevElement)
				}
			} else {
				if currentStatement.TypeOf() == tsPrefixexpStatement && st.TypeOf() == tsFuncCallStatement {
					st.AppendStatement(chainSt.ExctractPrefixexp())
					if chainSt.Len() > 0 {
						chainSt.Prev().AppendStatement(st)
					}
					chainSt.Append(st)
				} else if currentStatement.TypeOf() == tsAssignment && st.TypeOf() == tsFunction {
					// myvar = function() end
					exl := &explist{}
					currentStatement.AppendStatement(exl)
					chainSt.Append(exl)

					ex := &exp{}
					exl.AppendStatement(ex)
					chainSt.Append(ex)

					currentStatement = ex

					currentStatement.AppendStatement(st)
					chainSt.Append(st)

				} else if currentStatement.TypeOf() == tsAssignment && st.TypeOf() == tsTable {
					// table = {}
					exl := &explist{}
					currentStatement.AppendStatement(exl)
					chainSt.Append(exl)

					ex := &exp{}
					exl.AppendStatement(ex)
					chainSt.Append(ex)

					currentStatement = ex

					currentStatement.AppendStatement(st)
					chainSt.Append(st)

				} else if st.TypeOf() == tsAssignment {
					isPrefixexpConvertAssignment = true
					currentStatement = chainSt.ExctractPrefixexp()
					if chainSt.Len() > 0 {
						chainSt.Prev().AppendStatement(st)
					}
					chainSt.Append(st)
				} else {
					currentStatement.AppendStatement(st)
					chainSt.Append(st)
				}
			}

			for inner := st.InnerStatement(prevElement, curElement); inner != nil; inner = st.InnerStatement(prevElement, curElement) {
				if st.TypeOf() != inner.TypeOf() {
					st.AppendStatement(inner)
					chainSt.Append(inner)
				}
				st = inner
			}

			if isPrefixexpConvertAssignment {
				st.AppendStatement(currentStatement)

				if curElement.Token.Type == nAssign {
					for chst := chainSt.ExtractPrev(); chst.TypeOf() != assignmentWithOneVar.TypeOf(); chst = chainSt.ExtractPrev() {
					}

					assignmentWithOneVar.Append(curElement)
					st = assignmentWithOneVar
				}
			}

			currentStatement = st
		}

		if currentStatement != nil {
			currentStatement.Append(curElement)
		}

		prevElement = curElement
		curElement = nil
	}

	return doc, nil
}

func getStatement(s map[tokenID]branch, prev, cur *element) statementIntf {
	var branch branch

	if cur.Resolved {
		return nil
	}

	if prev != nil && prev.Token.Type == nReturn {
		branch = getsyntax(s, tokenID(nReturn))
		if cb, ok := branch[cur.Token.Type]; ok {
			return cb.New()
		}
	}

	if prev != nil && prev.Token.Type == nComma {
		branch = getsyntax(s, tokenID(nComma))
		if cb, ok := branch[cur.Token.Type]; ok {
			return cb.New()
		}
	}

	branch = getsyntax(s, tokenID(cur.Token.Type))
	if cb, ok := branch[nThis]; ok {
		return cb.New()
	}

	if prev != nil {
		branch = getsyntax(s, tokenID(prev.Token.Type))
		if cb, ok := branch[cur.Token.Type]; ok {
			return cb.New()
		}
	}

	return nil
}
