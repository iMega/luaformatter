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
		chainElements    []*element
		chainStatments   []statementIntf
		currentStatement statementIntf
		currentBranch    branch
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

		fmt.Printf("%s%s %s%s\n", mMagenta, TokenIDs[el.Token.Type], el.Token.Value, defaultStyle)

		// if currentStatement != nil && currentStatement.HasSyntax(el) == false {
		// 	bl := Block{}

		// 	switch v := currentStatement.(type) {
		// 	case *assignmentStatement:
		// 		bl.Statement = statement{Assignment: v}
		// 	}

		// 	doc.AddBlock(bl)
		// 	currentBranch = nil
		// 	currentStatement = nil
		// 	chainElements = nil
		// }
		if currentStatement != nil && hasExplist(currentStatement) && currentStatement.IsEnd(&el) {
			if len(chainStatments) > 0 {
				chainStatments = chainStatments[:len(chainStatments)-1]
				currentStatement = getLastStatement(chainStatments)
			}
		}

		lastStatement := getLastStatement(chainStatments)

		st := getStatement(getLastElement(chainElements), &el)
		if st != nil {
			currentStatement = st
		}

		if currentStatement != lastStatement {
			chainStatments = append(chainStatments, currentStatement)
			if lastStatement != nil {
				lastStatement.AppendStatement(currentStatement)
			}
		}

		chainElements = append(chainElements, &el)

		if currentStatement == nil {
			continue
		}

		for _, i := range chainElements {
			currentStatement.Append(i)
		}
		chainElements = nil

		if currentStatement != nil && !hasExplist(currentStatement) && currentStatement.IsEnd(&el) {
			if len(chainStatments) > 0 {
				chainStatments = chainStatments[:len(chainStatments)-1]
				currentStatement = getLastStatement(chainStatments)
			}
		}

		_ = currentBranch
		_ = chainElements
		_ = chainStatments
	}

	if currentStatement != nil {
		bl := Block{}

		switch v := currentStatement.(type) {
		case *assignmentStatement:
			bl.Statement = statement{Assignment: v}
		}

		doc.AddBlock(bl)
		currentBranch = nil
		currentStatement = nil
		chainElements = nil
	}

	return doc, nil
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

	el = last
	if last == nil {
		el = curEl
	}

	branch := getsyntax(tokenID(el.Token.Type))
	if branch == nil {
		return nil
	}

	if cb, ok := branch[nThis]; ok {
		return cb.New()
	}

	cb, ok := branch[curEl.Token.Type]
	if !ok {
		return nil
	}

	return cb.New()
}
