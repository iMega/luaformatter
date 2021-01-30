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
	Blocks map[uint64]statement
	Qty    uint64
}

func (body) New() statement {
	return &body{
		Blocks: make(map[uint64]statement),
	}
}

func (b *body) GetBody(prevSt statement, cur *element) statement {
	return prevSt
}

func (body) InnerStatement(prev, cur *element) (bool, statement) {
	return false, nil
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

	if cur.Token.Type == nUntil {
		return false, true
	}

	return false, false
}

func (b *body) Append(el *element) {}

func (b *body) AppendStatement(st statement) {
	if _, ok := st.(*prefixexpStatement); ok {
		return
	}

	b.Blocks[b.Qty] = st
	b.Qty++
}

func (b *body) GetStatement(prev, cur *element) statement {
	if prev != nil && prev.Token.Type == nLocal {
		if cur.Token.Type == nID {
			return &assignmentStatement{}
		}
	}

	switch cur.Token.Type {
	case nID:
		return &prefixexpStatement{IsUnknow: true}

	case nFor:
		return &forStatement{}

	case nFunction:
		return &functionStatement{}

	case nReturn:
		return &returnStatement{}

	case nIf:
		return &ifStatement{}

	case nLabel:
		return &labelStatement{}

	case nGoto:
		return &gotoStatement{}

	case nBreak:
		return &breakStatement{}

	case nDo:
		return &doStatement{}

	case nWhile:
		return &whileStatement{}

	case nRepeat:
		return &repeatStatement{}

	case nComment:
		return &commentStatement{
			IsNewline: prev != nil && prev.Token.StartLine != cur.Token.StartLine ||
				prev == nil && cur.Token.TC > 1,
		}

	case nCommentLong:
		return &commentStatement{}

	case nLF:
		return &newlineStatement{}
	}

	return nil
}

func (b *body) Format(c *Config, p printer, w io.Writer) error {
	for i := 0; i < int(b.Qty); i++ {
		st := b.Blocks[uint64(i)]

		if !p.IgnoreFirstPad {
			_, newlineOk := st.(*newlineStatement)
			commentSt, commentOk := st.(*commentStatement)
			if !newlineOk && !commentOk || commentSt != nil && commentSt.IsNewline {
				if err := newLine(w); err != nil {
					return err
				}

				if err := p.WritePad(w); err != nil {
					return err
				}
			}

			if commentSt != nil && !commentSt.IsNewline {
				if err := space(w); err != nil {
					return err
				}
			}
		}

		p.IgnoreFirstPad = false

		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	return nil
}
