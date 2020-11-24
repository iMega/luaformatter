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

import (
	"io"
	"strings"

	"github.com/timtadh/lexmachine"
	"github.com/timtadh/lexmachine/machines"
)

type scanner struct {
	lmScan *lexmachine.Scanner
	err    error
	el     element
}

func newScanner(code []byte) (*scanner, error) {
	lexer := lexmachine.NewLexer()

	for k, v := range keywords {
		lexer.Add([]byte(strings.ToLower(v)), token(k))
	}

	lexer.Add([]byte(`([a-zA-Z_.][a-zA-Z0-9_.:]*)`), token(nID))
	lexer.Add([]byte(`\s*\n\s*\n\s*`), token(nLF))
	lexer.Add([]byte("( |\t|\f|\r|\n)+"), skip)
	lexer.Add([]byte(`--\[\[([^\]\]])*\]\]`), token(nCommentLong))
	lexer.Add([]byte(`--( |\S)*`), token(nComment))
	lexer.Add([]byte(`::([^::])*::`), token(nLabel))

	lexer.Add([]byte(`(")[^(")]*(")`), token(nString))
	lexer.Add([]byte(`(')[^(')]*(')`), token(nString))
	lexer.Add([]byte(`(\[\[)[^(\]\])]*(\]\])`), token(nString))

	if err := lexer.Compile(); err != nil {
		return nil, err
	}

	s, err := lexer.Scanner(code)
	if err != nil {
		return nil, err
	}

	return &scanner{
		lmScan: s,
	}, nil
}

func (s *scanner) Next() bool {
	t, err, eof := s.lmScan.Next()
	if eof {
		return false
	}

	if _, is := err.(*machines.UnconsumedInput); is {
		s.err = err

		return false
	}

	if err != nil {
		s.err = err

		return false
	}

	token := t.(*lexmachine.Token)
	s.el = element{
		Token: token,
	}

	return true
}

type element struct {
	Token    *lexmachine.Token
	Resolved bool
	NL       int
	AddSpace bool
}

func (s *element) Format(c *Config, p printer, w io.Writer) error {
	_, err := w.Write(s.Token.Lexeme)

	return err
}

func (s *scanner) Scan() (element, error) {
	return s.el, s.err
}

func token(nodeID int) lexmachine.Action {
	return func(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
		return s.Token(nodeID, string(m.Bytes), m), nil
	}
}

func skip(*lexmachine.Scanner, *machines.Match) (interface{}, error) {
	return nil, nil
}

const (
	nEmpty = iota
	nID
	nLF
	nSpace
	nCommentLong
	nComment
	nBreak
	nDo
	nElse
	nElseif
	nEnd
	nFalse
	nFor
	nFunction
	nGoto
	nIf
	nIn
	nLocal
	nNil
	nRepeat
	nReturn
	nThen
	nTrue
	nUntil
	nWhile
	nColon
	nSemiColon
	nParentheses
	nClosingParentheses
	nSquareBracket
	nClosingSquareBracket
	nCurlyBracket
	nClosingCurlyBracket
	nAssign
	nComma
	nSingleQuote
	nDoubleQuote
	nBranch
	nNumber
	nString
	nStar
	nVararg
	nThis
	nLabel

	//
	// binop
	//

	// Logical Operators
	nAnd
	nOr

	// Arithmetic Operators
	nAddition
	nSubtraction
	nMultiplication
	nFloatDivision
	nFloorDivision
	nModulo
	nExponentiation

	// Bitwise Operators
	nBitwiseAND
	nBitwiseOR
	nBitwiseExclusiveOR
	nLeftShift
	nRightShift

	// Length Operator
	nLength

	// Concatenation
	nConcat

	// Relational Operators
	nEquality
	nInequality
	nLessThan
	nGreaterThan
	nLessOrEqual
	nGreaterOrEqual

	//
	// unop ‘-’ | not | ‘#’ | ‘~’
	//

	// Logical Operators
	nNot
)

var (
	Tokens []string

	keywords = map[int]string{
		nBreak:    "break",
		nDo:       "do",
		nElse:     "else",
		nElseif:   "elseif",
		nEnd:      "end",
		nFalse:    "false",
		nFor:      "for",
		nFunction: "function",
		nGoto:     "goto",
		nIf:       "if",
		nIn:       "in",
		nLocal:    "local",
		nNil:      "nil",

		nRepeat:    "repeat",
		nReturn:    "return",
		nThen:      "then",
		nTrue:      "true",
		nUntil:     "until",
		nWhile:     "while",
		nColon:     ":",
		nSemiColon: ";",
		nAssign:    `=`,
		nComma:     `,`,
		nStar:      `\*`,
		nVararg:    `\.\.\.`,

		// binop ::=  ‘+’ | ‘-’ | ‘*’ | ‘/’ | ‘//’ | ‘^’ | ‘%’ |
		//      ‘&’ | ‘~’ | ‘|’ | ‘>>’ | ‘<<’ | ‘..’ |
		//      ‘<’ | ‘<=’ | ‘>’ | ‘>=’ | ‘==’ | ‘~=’ |
		//      and | or

		// Logical Operators
		nAnd: "and",
		nOr:  "or",
		nNot: "not",

		// Arithmetic Operators
		nAddition:       `\+`,
		nSubtraction:    "-",
		nMultiplication: `\*`,
		nFloatDivision:  "/",
		nFloorDivision:  "//",
		nModulo:         "%",
		nExponentiation: `\^`,

		// Bitwise Operators
		nBitwiseAND:         "&",
		nBitwiseOR:          `\|`,
		nBitwiseExclusiveOR: "~",
		nLeftShift:          "<<",
		nRightShift:         ">>",

		// Length Operator
		nLength: "#",

		// Concatenation
		nConcat: `\.\.`,

		// Relational Operators
		nEquality:       "==",
		nInequality:     "~=",
		nLessThan:       "<",
		nGreaterThan:    ">",
		nLessOrEqual:    "<=",
		nGreaterOrEqual: ">=",

		nParentheses:          `\(`,
		nClosingParentheses:   `\)`,
		nSquareBracket:        `\[`,
		nClosingSquareBracket: `\]`,
		nCurlyBracket:         `\{`,
		nClosingCurlyBracket:  `\}`,
		// nSingleQuote:          `'`,
		// nDoubleQuote:          `"`,

		nNumber: `\d+(\.\d+)?`,
	}

	TokenIDs = map[int]string{
		nID:          "ID",
		nLF:          "LF",
		nSpace:       "Space",
		nCommentLong: "nComment",
		nComment:     "nComment",
		nAnd:         "nAnd",
		nBreak:       "nBreak",
		nDo:          "nDo",
		nElse:        "nElse",
		nElseif:      "nElseif",
		nEnd:         "nEnd",
		nFalse:       "nFalse",
		nFor:         "nFor",
		nFunction:    "nFunction",
		nGoto:        "nGoto",
		nIf:          "nIf",
		nIn:          "nIn",
		nLocal:       "nLocal",
		nNil:         "nNil",
		nNot:         "nNot",
		nOr:          "nOr",
		nRepeat:      "nRepeat",
		nReturn:      "nReturn",
		nThen:        "nThen",
		nTrue:        "nTrue",
		nUntil:       "nUntil",
		nWhile:       "nWhilee",
		nColon:       "nColon",
		nComma:       `nComma`,
		nAssign:      `nAssign`,
		nStar:        `nStar`,
		nVararg:      `nVararg`,
		nLabel:       "::",

		nSemiColon:            "nSemiColon",
		nParentheses:          "nParentheses",
		nClosingParentheses:   "nClosingParentheses",
		nSquareBracket:        `nSquareBracket`,
		nClosingSquareBracket: `nClosingSquareBracket`,
		nCurlyBracket:         `nCurlyBracket`,
		nClosingCurlyBracket:  `nClosingCurlyBracket`,
		nSingleQuote:          `nSingleQuote`,
		nDoubleQuote:          `nDoubleQuote`,

		nNumber: `Number`,
		nString: `String`,

		// Arithmetic Operators
		nAddition:       "+",
		nSubtraction:    "-",
		nMultiplication: "*",
		nFloatDivision:  "/",
		nFloorDivision:  "//",
		nModulo:         "%",
		nExponentiation: "^",

		// Length Operator
		nLength: "nLength",

		// Concatenation
		nConcat: "nConcat",

		// Relational Operators
		nEquality:       "nEquality",
		nInequality:     "nInequality",
		nLessThan:       "nLessThan",
		nGreaterThan:    "nGreaterThan",
		nLessOrEqual:    "nLessOrEqual",
		nGreaterOrEqual: "nGreaterOrEqual",
	}
)
