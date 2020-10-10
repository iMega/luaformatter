package formatter

type tokenID int

// function f () body end
//
// nFunction nSpace [nID] [nSpace] nParentheses BODY nEnd

const (
	InheritFormatter = iota
	FunctionFormatter
	ReturnFormatter
)

type branch map[int]statementIntf

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
		// nNumber: {
		// 	nAddition: &exp{},
		// },
		nAddition: {
			nNumber: &exp{},
		},
		nNegEq: {
			nNumber: &exp{},
		},
		// nLocal: {
		// 	nFunction: FunctionFormatter,
		// },
		// nFunction: {
		// 	nID:          InheritFormatter,
		// 	nParentheses: InheritFormatter,
		// },
		// nParentheses: {
		// 	nID:                 InheritFormatter,
		// 	nClosingParentheses: InheritFormatter,
		// },
		// nClosingParentheses: {
		// 	nEnd: InheritFormatter, //Указать что блок закрыт
		// },
		// nID: {
		// 	nParentheses:        InheritFormatter,
		// 	nComma:              InheritFormatter,
		// 	nClosingParentheses: InheritFormatter,
		// },
		nComma: {
			nID:       &exp{},
			nFunction: &exp{},
		},
		// nReturn: {
		// 	nID:    ReturnFormatter,
		// 	nFalse: ReturnFormatter,
		// },
	}

	// syntax = map[tokenID]branch{
	// 	nLocal: {
	// 		nID:       template{AddSpace: true},
	// 		nFunction: template{},
	// 	},
	// 	nID: {
	// 		nCurlyBracket: template{},
	// 	},
	// 	nClosingCurlyBracket: {},
	// }

	// block = map[tokenID]branch{
	// 	nID: {
	// 		nEq:          template{AddSpace: true},
	// 		nParentheses: template{},
	// 		nComma:       template{},
	// 	},
	// 	nEq: {
	// 		nID:     template{Tmpl: tmplFunctionCall},
	// 		nString: template{Tmpl: tmplVarList},
	// 		nFalse:  template{},
	// 		nTrue:   template{},
	// 		nNumber: template{},
	// 	},
	// 	nString: {
	// 		nClosingParentheses: template{},
	// 		nComma:              template{AddSpace: true},
	// 	},
	// 	nNumber: {
	// 		nClosingParentheses: template{},
	// 		nComma:              template{AddSpace: true},
	// 		nStar:               template{},
	// 	},
	// 	nStar: {
	// 		nNumber: template{},
	// 	},
	// 	nParentheses: {
	// 		nString: template{},
	// 	},
	// 	nClosingParentheses: {},
	// 	nCurlyBracket: {
	// 		nID: template{AddSpace: true},
	// 	},
	// 	nFalse: {
	// 		nComma: template{AddSpace: true},
	// 	},
	// 	nComma: {
	// 		nID:                  template{AddSpace: true},
	// 		nComment:             template{AddSpace: true, LF: true},
	// 		nClosingCurlyBracket: template{},
	// 	},
	// 	nComment:             {},
	// 	nClosingCurlyBracket: {
	// 		//nComma: template{},
	// 		//nComment: template{LF: true},
	// 	},
	// }
	// levels = map[int]int{
	// 	nThen: 1,
	// 	nDo:   1,
	// 	nEnd:  -1,
	// }
)

// type branch map[int]template

// func (b branch) Next(tokenType int) (branch, template) {
// 	tmpl, ok := b[tokenType]
// 	if !ok {
// 		return nil, template{}
// 	}

// 	return block[tokenID(tokenType)], tmpl
// }

func getsyntax(ID tokenID) branch {
	b, ok := syntax[ID]
	if !ok {
		return nil
	}

	return b
}

// func getSyntax(ID tokenID) branch {
// 	b, ok := syntax[ID]
// 	if !ok {
// 		return nil
// 	}

// 	return b
// }
