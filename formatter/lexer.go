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
	"bytes"
	"io"
	"unicode"

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
		lexer.Add(v, token(k))
	}

	lexer.Add([]byte(`([a-zA-Z_][a-zA-Z0-9_:]*)`), token(nID))
	lexer.Add([]byte(`\s*\n\s*\n\s*`), token(nLF))
	lexer.Add([]byte("( |\t|\f|\r|\n)+"), skip)
	lexer.Add([]byte(`::[^:]*::`), token(nLabel))

	lexer.Add([]byte(`--[^[\n]+[^\n]*`), comment(nComment))
	lexer.Add([]byte(`--\[\[[^]]*\]\]`), commentLong(nCommentLong))
	lexer.Add([]byte(`--\n`), comment(nComment))
	lexer.Add([]byte(`--\r\n`), comment(nComment))

	lexer.Add([]byte(`"[^"]*"`), token(nString))
	lexer.Add([]byte(`'[^']*'`), token(nString))
	lexer.Add([]byte(`\[\[[^]]*\]\]`), token(nString))

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

	if token, ok := t.(*lexmachine.Token); ok {
		s.el = element{
			Token: token,
		}
	}

	return true
}

type element struct {
	Token    *lexmachine.Token
	NL       int
	Resolved bool
	AddSpace bool
}

func (s *element) Format(c *Config, p printer, w io.Writer) error {
	// if s.Token.Type == nString {
	// 	return s.FormatString(c, p, w)
	// }

	_, err := w.Write(s.Token.Lexeme)

	return err
}

func (s *element) FormatString(c *Config, p printer, w io.Writer) error {
	curpos, ok := w.(cursorPositioner)
	if !ok {
		return errCastingType
	}

	_ = curpos

	return nil
}

func (s *scanner) Scan() (element, error) {
	return s.el, s.err
}

func token(nodeID int) lexmachine.Action {
	return func(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
		return s.Token(nodeID, string(m.Bytes), m), nil
	}
}

func comment(nodeID int) lexmachine.Action {
	return func(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
		prefix := []byte("-- ")
		if !bytes.HasPrefix(m.Bytes, prefix) {
			prefix = []byte("--")
		}

		b := bytes.TrimPrefix(m.Bytes, prefix)
		b = bytes.TrimRightFunc(b, unicode.IsSpace)
		m.Bytes = b

		return s.Token(nodeID, string(b), m), nil
	}
}

func commentLong(nodeID int) lexmachine.Action {
	return func(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
		b := bytes.TrimPrefix(m.Bytes, []byte("--[["))
		b = bytes.TrimSuffix(b, []byte("]]"))
		b = bytes.TrimSpace(b)
		m.Bytes = b

		return s.Token(nodeID, string(b), m), nil
	}
}

func skip(*lexmachine.Scanner, *machines.Match) (interface{}, error) {
	return nil, nil
}

const (
	nID = iota
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
	nDot
	nSingleQuote
	nDoubleQuote
	nNumber
	nString
	nVararg
	nLabel

	// Logical Operators.
	nAnd
	nOr

	// Arithmetic Operators.
	nAddition
	nSubtraction
	nMultiplication
	nFloatDivision
	nFloorDivision
	nModulo
	nExponentiation

	// Bitwise Operators.
	nBitwiseAND
	nBitwiseOR
	nBitwiseExclusiveOR
	nLeftShift
	nRightShift

	// Length Operator.
	nLength

	// Concatenation.
	nConcat

	// Relational Operators.
	nEquality
	nInequality
	nLessThan
	nGreaterThan
	nLessOrEqual
	nGreaterOrEqual

	// Logical Operators.
	nNot
)

var (
	Tokens []string

	keywords = map[int][]byte{
		nBreak:    []byte("break"),
		nDo:       []byte("do"),
		nElse:     []byte("else"),
		nElseif:   []byte("elseif"),
		nEnd:      []byte("end"),
		nFalse:    []byte("false"),
		nFor:      []byte("for"),
		nFunction: []byte("function"),
		nGoto:     []byte("goto"),
		nIf:       []byte("if"),
		nIn:       []byte("in"),
		nLocal:    []byte("local"),
		nNil:      []byte("nil"),
		nRepeat:   []byte("repeat"),
		nReturn:   []byte("return"),
		nThen:     []byte("then"),
		nTrue:     []byte("true"),
		nUntil:    []byte("until"),
		nWhile:    []byte("while"),

		nColon:     []byte(":"),
		nSemiColon: []byte(";"),
		nAssign:    []byte(`=`),
		nComma:     []byte(`,`),
		nDot:       []byte(`\.`),
		nVararg:    []byte(`\.\.\.`),

		// binop ::=  ‘+’ | ‘-’ | ‘*’ | ‘/’ | ‘//’ | ‘^’ | ‘%’ |
		//      ‘&’ | ‘~’ | ‘|’ | ‘>>’ | ‘<<’ | ‘..’ |
		//      ‘<’ | ‘<=’ | ‘>’ | ‘>=’ | ‘==’ | ‘~=’ |
		//      and | or

		// Logical Operators
		nAnd: []byte("and"),
		nOr:  []byte("or"),
		nNot: []byte("not"),

		// Arithmetic Operators
		nAddition:       []byte(`\+`),
		nSubtraction:    []byte("-"),
		nMultiplication: []byte(`\*`),
		nFloatDivision:  []byte("/"),
		nFloorDivision:  []byte("//"),
		nModulo:         []byte("%"),
		nExponentiation: []byte(`\^`),

		// Bitwise Operators
		nBitwiseAND:         []byte("&"),
		nBitwiseOR:          []byte(`\|`),
		nBitwiseExclusiveOR: []byte("~"),
		nLeftShift:          []byte("<<"),
		nRightShift:         []byte(">>"),

		// Length Operator
		nLength: []byte("#"),

		// Concatenation
		nConcat: []byte(`\.\.`),

		// Relational Operators
		nEquality:       []byte("=="),
		nInequality:     []byte("~="),
		nLessThan:       []byte("<"),
		nGreaterThan:    []byte(">"),
		nLessOrEqual:    []byte("<="),
		nGreaterOrEqual: []byte(">="),

		nParentheses:          []byte(`\(`),
		nClosingParentheses:   []byte(`\)`),
		nSquareBracket:        []byte(`\[`),
		nClosingSquareBracket: []byte(`\]`),
		nCurlyBracket:         []byte(`\{`),
		nClosingCurlyBracket:  []byte(`\}`),
		// nSingleQuote:          `'`,
		// nDoubleQuote:          `"`,

		nNumber: []byte(`\d+(\.\d+)?`),
	}

	TokenIDs = map[int]string{
		nID:          "ID",
		nLF:          "LF",
		nSpace:       "Space",
		nCommentLong: "nCommentLong",
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
		nDot:         `nDot`,
		nAssign:      `nAssign`,
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
