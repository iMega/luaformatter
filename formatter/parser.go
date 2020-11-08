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

		if currentStatement != nil {
			for ok := currentStatement.IsEnd(prevElement, curElement); ok; ok = currentStatement.IsEnd(prevElement, curElement) {
				cs := chainSt.ExtractPrev()
				if cs == nil {
					currentBody = doc.Bod
					chainSt.Append(currentBody)
					// currentBody.AppendStatement(currentStatement)
					// doc.AddBlock(newBlock(currentStatement))

					// el.Resolved = true
					currentStatement = currentBody

					break
				}

				currentStatement = cs
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

		// if curElement.Token.Type == nID && currentStatement.TypeOf() == tsBody {
		// 	s = map[tokenID]branch{
		// 		nID: {
		// 			nThis:        &prefixexpStatement{},
		// 			nParentheses: &funcCallStatement{},
		// 		},
		// 	}
		// }

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
					nString: &funcCallStatement{}, //local base = require "resty.core.base"
					//nString: &prefixexpStatement{},
				},
			}
			if currentStatement.TypeOf() == tsExp {
				s = map[tokenID]branch{
					nID: {
						//nString: &funcCallStatement{}, //local base = require "resty.core.base"
						nString: &prefixexpStatement{},
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
			isPrefixexpConvertAssignment := false

			if st.TypeOf() == tsAssignment && prevElement.Token.Type == nLocal {
				st.Append(prevElement)
			}

			if currentStatement == nil {
				chainSt.Append(st)

				if prevElement != nil {
					st.Append(prevElement)
				}
			} else {
				if currentStatement.TypeOf() == tsPrefixexpStatement && st.TypeOf() == tsFuncCallStatement {
					st.AppendStatement(chainSt.ExctractPrefixexp())
					// chainSt.Reset()
					if chainSt.Len() > 0 {
						chainSt.Prev().AppendStatement(st)
					}
					chainSt.Append(st)
				} else if st.TypeOf() == tsAssignment { //if currentStatement.TypeOf() == tsPrefixexpStatement && st.TypeOf() == tsAssignment {
					isPrefixexpConvertAssignment = true
					// st.AppendStatement()
					currentStatement = chainSt.ExctractPrefixexp()
					if chainSt.Len() > 0 {
						chainSt.Prev().AppendStatement(st)
					}
					chainSt.Append(st)
					// currentStatement = chainSt.First()
					// chainSt.Reset()
					// chainSt.Append(st)
					// } else if currentStatement.TypeOf() == tsExp && st.TypeOf() == tsPrefixexpStatement {
					// st.AppendStatement()
				} else {
					currentStatement.AppendStatement(st)
					chainSt.Append(st)
				}

				// if currentStatement.TypeOf() == tsPrefixexpStatement && st.TypeOf() == tsExp {
				// 	chainSt.Reset()
				// 	chainSt.Append(st)
				// }
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
				// st.AppendStatement(chainSt.First())
				// chainSt.Reset()
				// chainSt.Append(st)
			}

			currentStatement = st
		}

		if currentStatement != nil {
			currentStatement.Append(curElement)
		}

		prevElement = curElement
		curElement = nil
	}

	if chainSt.Len() > 0 {
		// currentStatement = chainSt.First()
	}

	// if currentStatement != nil {
	// 	doc.AddBlock(newBlock(currentStatement))
	// 	currentStatement = nil
	// }

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
