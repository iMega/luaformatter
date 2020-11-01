package formatter

// Parse code.
func parse(code []byte) (*document, error) {
	var (
		prevElement      *element
		curElement       *element
		currentStatement statementIntf

		chainSt = &chainStatments{}
	)

	s, err := newScanner(code)
	if err != nil {
		return nil, err
	}

	doc := NewDocument()

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
				cs := chainSt.Prev()
				if cs == nil {
					doc.AddBlock(newBlock(currentStatement))

					// el.Resolved = true
					currentStatement = cs

					break
				}

				currentStatement = cs
			}
		}

		s := syntax
		// prefixexp assignment or function call
		if curElement.Token.Type == nID && currentStatement == nil {
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
			if currentStatement == nil {
				chainSt.Append(st)

				if prevElement != nil {
					st.Append(prevElement)
				}
			} else {
				if currentStatement.TypeOf() == tsPrefixexpStatement && st.TypeOf() == tsFuncCallStatement {
					st.AppendStatement(chainSt.First())
					chainSt.Reset()
					chainSt.Append(st)
				} else if currentStatement.TypeOf() == tsPrefixexpStatement && st.TypeOf() == tsAssignment {
					isPrefixexpConvertAssignment = true
					currentStatement = chainSt.First()
					chainSt.Reset()
					chainSt.Append(st)
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
		currentStatement = chainSt.First()
	}

	if currentStatement != nil {
		doc.AddBlock(newBlock(currentStatement))
		currentStatement = nil
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
