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
		currentStatement statement
		currentBody      statement

		chainSt = &chainStatments{}
	)

	s, err := newScanner(code)
	if err != nil {
		return nil, err
	}

	doc := &document{}
	b := new(body).New()
	doc.Body = b
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
			if currentStatement.TypeOf() == tsUnknow && curElement.Token.Type == nComma {
				break
			}

			cs := chainSt.ExtractPrev()
			if cs == nil {
				currentBody = doc.Body
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

		if st := currentStatement.GetStatement(prevElement, curElement); st != nil {
			var assignmentWithOneVar statement

			isPrefixexpConvertAssignment := false

			if st.TypeOf() == tsAssignment || st.TypeOf() == tsFunction {
				if prevElement != nil && prevElement.Token.Type == nLocal {
					st.Append(prevElement)
				}
			}

			if st.TypeOf() == tsAssignment && curElement.Token.Type == nAssign { // a = 1
				assignmentWithOneVar = st

				if s := chainSt.ExctractStatement(tsAssignment); s != nil {
					currentStatement = s
					currentStatement.Append(curElement)
					chainSt.Append(currentStatement)

					st = &explist{}
				}
			}

			if currentStatement.TypeOf() == tsPrefixexpStatement && st.TypeOf() == tsFuncCallStatement {
				extracted := chainSt.ExctractStatement(tsUnknow)
				if extracted == nil {
					extracted = chainSt.ExctractStatement(tsPrefixexpStatement)
				}

				st.AppendStatement(extracted)
				chainSt.Prev().AppendStatement(st)
			} else if currentStatement.TypeOf() == tsUnknow && st.TypeOf() == tsFuncCallStatement {
				st.AppendStatement(chainSt.ExctractStatement(tsUnknow))
				chainSt.Prev().AppendStatement(st)
			} else if st.TypeOf() == tsAssignment {
				isPrefixexpConvertAssignment = true
				currentStatement = chainSt.ExctractStatement(tsUnknow)
				if currentStatement == nil {
					currentStatement = chainSt.ExctractStatement(tsPrefixexpStatement)
				}
				chainSt.Prev().AppendStatement(st)
			} else {
				currentStatement.AppendStatement(st)
			}

			chainSt.Append(st)

			for isBreak, inner := st.InnerStatement(prevElement, curElement); inner != nil; isBreak, inner = st.InnerStatement(prevElement, curElement) {
				// if st.TypeOf() != inner.TypeOf() {
				st.AppendStatement(inner)
				chainSt.Append(inner)
				// }

				st = inner

				if isBreak {
					break
				}
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

		currentStatement.Append(curElement)

		if curElement.Token.Type != nCommentLong && curElement.Token.Type != nComment {
			prevElement = curElement
		}

		curElement = nil
	}

	return doc, nil
}
