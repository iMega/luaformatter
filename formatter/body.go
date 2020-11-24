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
)

type body struct {
	Blocks map[uint64]block
	Qty    uint64
}

func (body) New() statementIntf {
	return &body{
		Blocks: make(map[uint64]block),
	}
}

func (b *body) GetBody(prevSt statementIntf, cur *element) statementIntf {
	return prevSt
}

func (body) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (body) TypeOf() typeStatement { return tsBody }

func (b *body) IsEnd(prev, cur *element) (bool, bool) {
	if cur.Token.Type == nEnd {
		return false, true
	}

	if cur.Token.Type == nElseif {
		return false, true
	}

	if cur.Token.Type == nElse {
		return false, true
	}

	return false, false
}

func (b *body) Append(el *element) {}

func (b *body) AppendStatement(st statementIntf) {
	if _, ok := st.(*prefixexpStatement); ok {
		return
	}

	b.Blocks[b.Qty] = newBloc(st)
	b.Qty++
}

func (b *body) Format(c *Config, p printer, w io.Writer) error {
	for i := 0; i < int(b.Qty); i++ {
		st := b.Blocks[uint64(i)]
		if err := st.Format(c, p, w); err != nil {
			return err
		}

		if int(b.Qty)-1 == i {
			continue
		}

		if s := st.Statement.NewLine; s == nil {
			if err := newLine(w); err != nil {
				return err
			}
		}
	}

	return nil
}

func (b *block) Format(c *Config, p printer, w io.Writer) error {
	if !p.IgnoreFirstPad {
		if b.Statement.NewLine == nil {
			if err := p.WritePad(w); err != nil {
				return err
			}
		}
	}
	p.IgnoreFirstPad = false

	if s := b.Statement.Assignment; s != nil {
		if err := s.Format(c, p, w); err != nil {
			return err
		}
	}

	if st := b.Statement.Do; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	if st := b.Statement.If; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	if st := b.Statement.For; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	if s := b.Statement.FuncCall; s != nil {
		if err := s.Format(c, p, w); err != nil {
			return err
		}
	}

	if s := b.Statement.Function; s != nil {
		if err := s.Format(c, p, w); err != nil {
			return err
		}
	}

	if s := b.Return; s != nil {
		if err := s.Format(c, p, w); err != nil {
			return err
		}
	}

	if s := b.Statement.Comment; s != nil {
		if err := s.Format(c, p, w); err != nil {
			return err
		}
	}

	if s := b.Statement.NewLine; s != nil {
		if err := s.Format(c, p, w); err != nil {
			return err
		}
	}

	return nil
}

func newBloc(st statementIntf) block {
	bl := block{}

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

	case *forStatement:
		bl.Statement = statement{For: v}

	case *returnStatement:
		bl.Return = v

	case *commentStatement:
		bl.Statement = statement{Comment: v}

	case *newlineStatement:
		bl.Statement = statement{NewLine: v}
	}

	return bl
}

type block struct {
	Statement statement
	Return    *returnStatement
}
