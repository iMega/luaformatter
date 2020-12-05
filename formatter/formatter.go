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

type Config struct {
	IndentSize    uint8 `mapstructure:"indent-size"`
	MaxLineLength uint8 `mapstructure:"max-line-length"`
}

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

func DefaultConfig() Config {
	return Config{
		IndentSize:    4,
		MaxLineLength: 80,
	}
}

type printer struct {
	Pad             uint8
	ParentStatement typeStatement
	IgnoreFirstPad  bool
}

func (p printer) WritePad(w io.Writer) error {
	b := bytes.Repeat([]byte(" "), int(p.Pad))
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

func newBlock(st statementIntf) Block {
	bl := Block{}

	switch v := st.(type) {
	case *assignmentStatement:
		bl.Statement = statement{Assignment: v}

	case *labelStatement:
		bl.Statement = statement{Label: v}

	case *gotoStatement:
		bl.Statement = statement{Goto: v}

	case *breakStatement:
		bl.Statement = statement{Break: v}

	case *doStatement:
		bl.Statement = statement{Do: v}

	case *whileStatement:
		bl.Statement = statement{While: v}

	case *repeatStatement:
		bl.Statement = statement{Repeat: v}

	case *functionStatement:
		bl.Statement = statement{Function: v}

	case *funcCallStatement:
		bl.Statement = statement{FuncCall: v}

	case *prefixexpStatement:
		bl.Statement = statement{
			FuncCall: &funcCallStatement{
				Prefixexp: v,
			},
		}

	case *ifStatement:
		bl.Statement = statement{If: v}

	case *returnStatement:
		bl.Return = v

	case *commentStatement:
		bl.Statement = statement{Comment: v}

	case *newlineStatement:
		bl.Statement = statement{NewLine: v}
	}

	return bl
}
