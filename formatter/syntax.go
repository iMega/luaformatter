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

type tokenID int

const (
	InheritFormatter = iota
	FunctionFormatter
	ReturnFormatter
)

var (
	syntax = map[tokenID]branch{
		nSquareBracket: {
			nThis: &prefixexpStatement{},
		},
		nClosingSquareBracket: {
			// nID:          &prefixexpStatement{},
			nDot:         &prefixexpStatement{},  // ].
			nParentheses: &funcCallStatement{},   // ](
			nComma:       &assignmentStatement{}, // ],
		},
		nAssign: {
			nSubtraction:  &explist{}, // var = -
			nNumber:       &explist{}, // var = 1
			nID:           &explist{}, // var1 = var
			nString:       &explist{}, // var = ""
			nLength:       &explist{}, // var = #var1
			nFalse:        &explist{}, // var = false
			nNil:          &explist{}, // var = nil
			nTrue:         &explist{}, // var = true
			nFunction:     &functionStatement{},
			nCurlyBracket: &tableStatement{}, // table = {}
		},
		nFunction: {
			nThis: &functionStatement{},
		},
		nReturn: {
			nThis:     &returnStatement{},
			nFunction: &explist{},
			nID:       &explist{},
			nNumber:   &explist{},
			nFalse:    &explist{},
			nTrue:     &explist{},
		},
		nIf: {
			nThis:   &ifStatement{},
			nID:     &exp{}, // if id
			nLength: &exp{}, // if #
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
			nNumber: &exp{}, // + 3
			nID:     &exp{}, // + id
		},
		nSubtraction: {
			nNumber: &exp{}, // - 3
		},
		nMultiplication: {
			nNumber: &exp{}, // * 3
		},
		nFloatDivision: {
			nNumber: &exp{}, // / 3
		},
		nBitwiseAND:         {nNumber: &exp{}}, // & 3
		nModulo:             {nNumber: &exp{}}, // % 3
		nExponentiation:     {nNumber: &exp{}}, // ^ 3
		nLessThan:           {nNumber: &exp{}}, // < 3
		nLeftShift:          {nNumber: &exp{}}, // << 3
		nLessOrEqual:        {nNumber: &exp{}}, // <= 3
		nGreaterThan:        {nNumber: &exp{}}, // > 3
		nGreaterOrEqual:     {nNumber: &exp{}}, // >= 3
		nRightShift:         {nNumber: &exp{}}, // >> 3
		nBitwiseOR:          {nNumber: &exp{}}, // | 3
		nBitwiseExclusiveOR: {nNumber: &exp{}}, // ~ 3
		nInequality:         {nNumber: &exp{}}, // ~= 3
		nAnd: {
			nNumber: &exp{}, // and 3
			nLength: &exp{}, // and #
		},
		nOr: {
			nNumber: &exp{}, // or 3
			nLength: &exp{}, // or #
		},
		nConcat: {
			nID:     &exp{}, // .. id
			nString: &exp{}, // .. "string"
			nNumber: &exp{}, // .. 3
		},
		nEquality: {
			nString: &exp{},
			nNumber: &exp{}, // == 3
			nID:     &exp{}, // if name == searched then
		},
		nComma: {
			nID:          &exp{},
			nNumber:      &exp{},
			nFunction:    &exp{},
			nVararg:      &exp{},
			nSubtraction: &exp{}, // for num
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
		nLocal: {
			nID:       &assignmentStatement{},
			nFunction: &functionStatement{},
		},
		nComment: {
			nThis: &commentStatement{},
		},
		nLF: {
			nThis: &newlineStatement{},
		},
		nFor: {
			nThis: &forStatement{},
			nID:   &field{},
		},
		nIn: {
			nID: &explist{},
		},
		nCurlyBracket: {
			nThis: &tableStatement{},
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
