package formatter

import (
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
	lexer.Add([]byte(`([a-zA-Z_][a-zA-Z0-9_.:]*)`), token(nID))
	lexer.Add([]byte("( |\t|\f|\r|\n)+"), skip)
	lexer.Add([]byte(`--\[\[([^\]\]])*\]\]`), token(nCommentLong))
	lexer.Add([]byte(`--( |\S)*`), token(nComment))
	lexer.Add([]byte(`\n\s*\n`), token(nLF))

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
	nWhil
	nNegEq
	nColon
	nSemiColon
	nParentheses
	nClosingParentheses
	nSquareBracket
	nClosingSquareBracket
	nCurlyBracket
	nClosingCurlyBracket
	nAssing
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
		nWhil:      "while",
		nColon:     ":",
		nSemiColon: ";",
		nAssing:    `=`,
		nComma:     `,`,
		nStar:      `\*`,
		nVararg:    `\.\.\.`,
		nLabel:     "::",

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
		nString: `('|"|\[\[)[^('|"|\]\])]*('|"|\]\])`,
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
		nWhil:        "nWhile",
		nNegEq:       "nNegEq",
		nColon:       "nColon",
		nComma:       `nComma`,
		nAssing:      `nAssing`,
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

		// Relational Operators
		nEquality:       "nEquality",
		nInequality:     "nInequality",
		nLessThan:       "nLessThan",
		nGreaterThan:    "nGreaterThan",
		nLessOrEqual:    "nLessOrEqual",
		nGreaterOrEqual: "nGreaterOrEqual",
	}
)
