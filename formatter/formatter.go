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
)

const (
	defaultStyle = "\x1b[0m"
	mMagenta     = "\x1b[35m"
)

func Format(c Config, b []byte, w io.Writer) error {
	doc, err := parse(b)
	if err != nil {
		return err
	}

	if doc.Body == nil {
		return nil
	}

	p := printer{}

	if st, ok := doc.Body.(*body); ok {
		if err := st.Format(&c, p, w); err != nil {
			return err
		}
	}

	if err := newLine(w); err != nil {
		return err
	}

	return nil
}

type printer struct {
	ParentStatement typeStatement
	Pad             uint8

	SpacesBeforeAssign  uint8
	SpacesBeforeComment uint8

	IgnoreFirstPad bool
}

func (p printer) WritePad(w io.Writer) error {
	return p.WriteSpaces(w, int(p.Pad))
}

func (p printer) WriteSpaces(w io.Writer, count int) error {
	b := bytes.Repeat([]byte(" "), count)
	_, err := w.Write(b)

	return err
}

func newLine(w io.Writer) error {
	_, err := w.Write([]byte(newLineSymbol))

	return err
}

func space(w io.Writer) error {
	_, err := w.Write([]byte(" "))

	return err
}

func writeKeyword(c *Config, tokenType int, w io.Writer) error {
	raw := keywords[tokenType]

	if c.Highlight {
		raw = append([]byte("\x1b[1m"), raw...)
		raw = append(raw, "\x1b[0m"...)
	}

	_, err := w.Write(raw)

	return err
}

var tokens = map[int][]byte{
	nColon:     []byte(":"),
	nSemiColon: []byte(";"),
	nAssign:    []byte(`=`),
	nComma:     []byte(`,`),
	nDot:       []byte(`.`),
	nVararg:    []byte(`...`),

	// Arithmetic Operators
	nAddition:       []byte(`+`),
	nSubtraction:    []byte("-"),
	nMultiplication: []byte(`*`),
	nFloatDivision:  []byte("/"),
	nFloorDivision:  []byte("//"),
	nModulo:         []byte("%"),
	nExponentiation: []byte(`^`),

	// Bitwise Operators
	nBitwiseAND:         []byte("&"),
	nBitwiseOR:          []byte(`|`),
	nBitwiseExclusiveOR: []byte("~"),
	nLeftShift:          []byte("<<"),
	nRightShift:         []byte(">>"),

	// Length Operator
	nLength: []byte("#"),

	// Concatenation
	nConcat: []byte(`..`),

	// Relational Operators
	nEquality:       []byte("=="),
	nInequality:     []byte("~="),
	nLessThan:       []byte("<"),
	nGreaterThan:    []byte(">"),
	nLessOrEqual:    []byte("<="),
	nGreaterOrEqual: []byte(">="),

	nParentheses:          []byte(`(`),
	nClosingParentheses:   []byte(`)`),
	nSquareBracket:        []byte(`[`),
	nClosingSquareBracket: []byte(`]`),
	nCurlyBracket:         []byte(`{`),
	nClosingCurlyBracket:  []byte(`}`),
}

func isRelationalOperator(el *element) bool {
	ops := []int{
		nEquality,
		nInequality,
		nLessThan,
		nGreaterThan,
		nLessOrEqual,
		nGreaterOrEqual,
	}

	return binarySearch(ops, el.Token.Type)
}

func spaceAroundOption(c *Config, tokenType int, w io.Writer) error {
	raw := tokens[tokenType]

	raw = append([]byte(" "), raw...)
	raw = append(raw, " "...)

	_, err := w.Write(raw)

	return err
}
