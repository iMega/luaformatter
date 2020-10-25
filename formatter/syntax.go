package formatter

type tokenID int

const (
	InheritFormatter = iota
	FunctionFormatter
	ReturnFormatter
)

var (
	syntax = map[tokenID]branch{
		// nID: {
		// nThis:        &prefixexpStatement{},
		// nParentheses: &funcCallStatement{},
		// },
		nSquareBracket: {
			nThis: &prefixexpStatement{},
		},
		nClosingSquareBracket: {
			nID:          &prefixexpStatement{},
			nParentheses: &funcCallStatement{},
			nComma:       &assignmentStatement{},
		},
		nAssign: {
			nNumber:   &explist{},
			nFunction: &functionStatement{},
		},
		nFunction: {
			nThis: &functionStatement{},
		},
		nReturn: {
			nThis:     &returnStatement{},
			nFunction: &explist{},
			nID:       &explist{},
			nNumber:   &explist{},
		},
		nIf: {
			nThis: &ifStatement{},
			nID:   &exp{},
		},
		nElseif: {
			nThis: &elseifStatement{},
			nID:   &exp{},
		},
		nElse: {
			nThis: &elseStatement{},
			nID:   &exp{},
		},
		nLabel: {nThis: &labelStatement{}},
		nGoto:  {nThis: &gotoStatement{}},
		nBreak: {nThis: &breakStatement{}},
		nAddition: {
			nNumber: &exp{},
		},
		nEquality: {
			nNumber: &exp{},
		},
		nInequality: {
			nNumber: &exp{},
		},
		nComma: {
			nID:       &exp{},
			nNumber:   &exp{},
			nFunction: &exp{},
		},
		nDo: {
			nThis: &doStatement{},
		},
		nWhile: {
			nThis: &whileStatement{},
			nID:   &exp{},
		},
		nRepeat: {
			nThis: &repeatStatement{},
		},
		nUntil: {
			nID: &exp{},
		},
	}
)

type branch map[int]statementIntf

func getsyntax(s map[tokenID]branch, ID tokenID) branch {
	b, ok := s[ID]
	if !ok {
		return nil
	}

	return b
}
