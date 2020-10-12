package formatter

type tokenID int

const (
	InheritFormatter = iota
	FunctionFormatter
	ReturnFormatter
)

var (
	syntax = map[tokenID]branch{
		nID: {
			nComma:  &assignmentStatement{},
			nAssing: &assignmentStatement{},
		},
		nAssing: {
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

func getsyntax(ID tokenID) branch {
	b, ok := syntax[ID]
	if !ok {
		return nil
	}

	return b
}
