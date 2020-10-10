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
		nLabel: {nThis: &labelStatement{}},
		nGoto:  {nThis: &gotoStatement{}},
		nBreak: {nThis: &breakStatement{}},
		nAddition: {
			nNumber: &exp{},
		},
		nInequality: {
			nNumber: &exp{},
		},
		nComma: {
			nID:       &exp{},
			nFunction: &exp{},
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
