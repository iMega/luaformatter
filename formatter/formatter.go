package formatter

import (
	"fmt"
)

const (
	defaultStyle = "\x1b[0m"
	mMagenta     = "\x1b[35m"
)

// Format code format
func Parse(code []byte) (*document, error) {
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

		if prevElement != nil {
			fmt.Printf("%s%s %s%s = ", mMagenta, TokenIDs[prevElement.Token.Type], prevElement.Token.Value, defaultStyle)
		}

		fmt.Printf("%s%s %s%s\n", mMagenta, TokenIDs[el.Token.Type], el.Token.Value, defaultStyle)

		if currentStatement != nil {
			for ok := currentStatement.IsEnd(prevElement, curElement); ok; ok = currentStatement.IsEnd(prevElement, curElement) {
				cs := chainSt.Prev()
				if cs == nil {
					bl := Block{}

					switch v := currentStatement.(type) {
					case *assignmentStatement:
						bl.Statement = statement{Assignment: v}
					case *functionStatement:
						bl.Statement = statement{Function: v}
					}

					doc.AddBlock(bl)
					currentStatement = cs

					break
				}
				currentStatement = cs
			}
		}

		if st := getStatement(prevElement, curElement); st != nil {
			if currentStatement == nil {
				chainSt.Append(st)
				if prevElement != nil {
					st.Append(prevElement)
				}
			} else {
				// if currentStatement.TypeOf() != st.TypeOf() {
				currentStatement.AppendStatement(st)
				chainSt.Append(st)
				// }

				for inner := st.InnerStatement(); inner != nil; inner = nil {
					if st.TypeOf() != inner.TypeOf() {
						st.AppendStatement(inner)
						chainSt.Append(inner)
					}
					st = inner
				}
			}
			currentStatement = st
		}

		if currentStatement != nil {
			currentStatement.Append(curElement)

			if currentStatement.IsEnd(prevElement, curElement) {
				currentStatement = chainSt.Prev()
			}
		}

		prevElement = curElement
		curElement = nil
	}

	if chainSt.Len() > 0 {
		currentStatement = chainSt.First()
	}

	if currentStatement != nil {
		bl := Block{}

		switch v := currentStatement.(type) {
		case *assignmentStatement:
			bl.Statement = statement{Assignment: v}
		case *functionStatement:
			bl.Statement = statement{Function: v}
		case *returnStatement:
			bl.Return = v
		}

		doc.AddBlock(bl)
		currentStatement = nil
	}

	return doc, nil
}

func statementAppend(st statementIntf, elOrSt interface{}) {
	switch v := elOrSt.(type) {
	case *element:
		st.Append(v)
	case statementIntf:
		st.AppendStatement(v)
	}
}

func hasExplist(st statementIntf) bool {
	switch st.(type) {
	case *assignmentStatement:
		return true

	case *returnStatement:
		return true
	}

	return false
}

func getLastElement(chain []*element) *element {
	if len(chain) == 0 {
		return nil
	}

	return chain[len(chain)-1]
}

func getLastStatement(chain []statementIntf) statementIntf {
	if len(chain) == 0 {
		return nil
	}

	return chain[len(chain)-1]
}

func getStatement(last, curEl *element) statementIntf {
	var el *element

	branch := getsyntax(tokenID(curEl.Token.Type))
	if cb, ok := branch[nThis]; ok {
		return cb.New()
	}
	//

	el = last
	if last == nil {
		el = curEl
	}

	branch = getsyntax(tokenID(el.Token.Type))
	if branch == nil {
		branch = getsyntax(tokenID(curEl.Token.Type))
		if cb, ok := branch[nThis]; ok {
			return cb.New()
		}
		return nil
	}

	if cb, ok := branch[curEl.Token.Type]; ok {
		return cb.New()
	}

	// if cb, ok := branch[nThis]; ok {
	// 	return cb.New()
	// }

	return nil
}
