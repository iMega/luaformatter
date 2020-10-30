package formatter

import (
	"io"
)

const (
	defaultStyle = "\x1b[0m"
	mMagenta     = "\x1b[35m"
)

type Config struct {
	IndentSize int `mapstructure:"indent-size"`
}

func Format(c Config, b []byte, w io.Writer) error {
	doc, err := parse(b)
	if err != nil {
		return err
	}

	for _, b := range doc.Body {
		b.Format(&c, w)
	}

	return nil
}

func DefaultConfig() Config {
	return Config{
		IndentSize: 4,
	}
}

func newLine(w io.Writer) error {
	_, err := w.Write([]byte(newLineSymbol))

	return err
}

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

					el.Resolved = true
					currentStatement = cs

					break
				}

				currentStatement = cs
			}
		}

		s := syntax
		if curElement.Token.Type == nID && currentStatement == nil {
			s = map[tokenID]branch{
				nID: {
					nThis:        &prefixexpStatement{},
					nParentheses: &funcCallStatement{},
				},
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
			if currentStatement == nil {
				chainSt.Append(st)

				if prevElement != nil {
					st.Append(prevElement)
				}
			} else {
				isPrefixexpConvertAssignment := false
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

func newBlock(st statementIntf) Block {
	bl := Block{}

	switch v := st.(type) {
	case *assignmentStatement:
		bl.Statement = statement{Assignment: v}

	case *labelStatement:
		bl.Statement = statement{Label: v}

	case *gotoStatement:
		bl.Statement = statement{Goto: v}

	case *breakStatement:
		bl.Statement = statement{Break: v}

	case *doStatement:
		bl.Statement = statement{Do: v}

	case *whileStatement:
		bl.Statement = statement{While: v}

	case *repeatStatement:
		bl.Statement = statement{Repeat: v}

	case *functionStatement:
		bl.Statement = statement{Function: v}

	case *funcCallStatement:
		bl.Statement = statement{FuncCall: v}

	case *prefixexpStatement:
		bl.Statement = statement{
			FuncCall: &funcCallStatement{
				Prefixexp: v,
			},
		}

	case *ifStatement:
		bl.Statement = statement{If: v}

	case *returnStatement:
		bl.Return = v
	}

	return bl
}
