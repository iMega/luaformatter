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

type ifStatement struct {
	Exp        *exp
	Body       statement
	ElseIfPart []*elseifStatement
	ElsePart   *elseStatement
}

func (ifStatement) InnerStatement(prev, cur *element) (bool, statement) {
	return false, nil
}

func (ifStatement) TypeOf() typeStatement {
	return tsNone
}

func (s *ifStatement) IsEnd(prev, cur *element) (bool, bool) {
	if cur.Token.Type == nEnd {
		return true, true
	}

	return false, false
}

func (s *ifStatement) Append(el *element) {}

func (s *ifStatement) AppendStatement(st statement) {
	switch v := st.(type) {
	case *exp:
		s.Exp = v
	case *elseifStatement:
		s.ElseIfPart = append(s.ElseIfPart, v)
	case *elseStatement:
		s.ElsePart = v
	}
}

func (s *ifStatement) GetBody(prevSt statement, cur *element) statement {
	if cur.Token.Type != nThen {
		return prevSt
	}

	if s.Body == nil {
		s.Body = new(body).New()
	}

	return s.Body
}

func (s *ifStatement) GetStatement(prev, cur *element) statement {
	if cur.Token.Type == nElseif {
		return &elseifStatement{}
	}

	if cur.Token.Type == nElse {
		return &elseStatement{}
	}

	if isExp(cur) || isBinop(cur) {
		return &exp{}
	}

	return nil
}

func (s *ifStatement) Format(c *Config, p printer, w io.Writer) error {
	if err := writeKeywordWithSpaceRight(c, nIf, w); err != nil {
		return err
	}

	curpos := getCursorPosition(w)
	buf := bytes.NewBuffer(nil)
	if err := s.Exp.Format(c, p, buf); err != nil {
		return err
	}
	curpos.Col += uint64(buf.Len())
	buf.Reset()
	np := p
	np.IfStatementExpLong = curpos.Col > uint64(c.MaxLineLength+1)

	if np.IfStatementExpLong {
		if err := p.WriteSpaces(w, int(c.IndentSize-3)); err != nil {
			return err
		}
	}

	if err := s.Exp.Format(c, np, w); err != nil {
		return err
	}

	if np.IfStatementExpLong {
		if err := newLine(w); err != nil {
			return err
		}
	} else {
		if err := space(w); err != nil {
			return err
		}
	}

	if err := writeKeyword(c, nThen, w); err != nil {
		return err
	}

	if st, ok := s.Body.(*body); ok {
		ip := p
		ip.Pad = p.Pad + c.IndentSize

		if err := st.Format(c, ip, w); err != nil {
			return err
		}

		if err := newLine(w); err != nil {
			return err
		}
	}

	for _, i := range s.ElseIfPart {
		if err := i.Format(c, p, w); err != nil {
			return err
		}
	}

	if st := s.ElsePart; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	if err := p.WritePad(w); err != nil {
		return err
	}

	if err := writeKeyword(c, nEnd, w); err != nil {
		return err
	}

	return nil
}

type elseifStatement struct {
	Exp  *exp
	Body statement
}

func (elseifStatement) InnerStatement(prev, cur *element) (bool, statement) {
	return false, nil
}

func (elseifStatement) TypeOf() typeStatement {
	return tsNone
}

func (s *elseifStatement) IsEnd(prev, cur *element) (bool, bool) {
	if cur.Token.Type == nElseif {
		return false, true
	}

	return false, cur.Token.Type == nEnd || cur.Token.Type == nElse
}

func (s *elseifStatement) Append(el *element) {}

func (s *elseifStatement) AppendStatement(st statement) {
	if v, ok := st.(*exp); ok {
		s.Exp = v
	}
}

func (s *elseifStatement) GetBody(prevSt statement, cur *element) statement {
	if cur.Token.Type != nThen {
		return prevSt
	}

	if s.Body == nil {
		s.Body = new(body).New()
	}

	return s.Body
}

func (s *elseifStatement) GetStatement(prev, cur *element) statement {
	if cur.Token.Type == nElseif {
		return &elseifStatement{}
	}

	if isExp(cur) {
		return &exp{}
	}

	return nil
}

func (s *elseifStatement) Format(c *Config, p printer, w io.Writer) error {
	if err := p.WritePad(w); err != nil {
		return err
	}

	if err := writeKeywordWithSpaceRight(c, nElseif, w); err != nil {
		return err
	}

	if s.Exp != nil {
		if err := s.Exp.Format(c, p, w); err != nil {
			return err
		}
	}

	if err := writeKeywordWithSpaceLeft(c, nThen, w); err != nil {
		return err
	}

	if st, ok := s.Body.(*body); ok {
		ip := p
		ip.Pad = p.Pad + c.IndentSize

		if err := st.Format(c, ip, w); err != nil {
			return err
		}

		if err := newLine(w); err != nil {
			return err
		}
	}

	return nil
}

type elseStatement struct {
	Body statement
}

func (elseStatement) InnerStatement(prev, cur *element) (bool, statement) {
	return false, nil
}

func (elseStatement) TypeOf() typeStatement {
	return tsNone
}

func (s *elseStatement) IsEnd(prev, cur *element) (bool, bool) {
	return false, cur.Token.Type == nEnd
}

func (s *elseStatement) Append(el *element) {}

func (s *elseStatement) AppendStatement(st statement) {}

func (s *elseStatement) GetBody(statement, *element) statement {
	if s.Body == nil {
		s.Body = new(body).New()
	}

	return s.Body
}

func (s *elseStatement) GetStatement(prev, cur *element) statement {
	return nil
}

func (s *elseStatement) Format(c *Config, p printer, w io.Writer) error {
	if err := p.WritePad(w); err != nil {
		return err
	}

	if err := writeKeyword(c, nElse, w); err != nil {
		return err
	}

	if st, ok := s.Body.(*body); ok {
		ip := p
		ip.Pad = p.Pad + c.IndentSize

		if err := st.Format(c, ip, w); err != nil {
			return err
		}

		if err := newLine(w); err != nil {
			return err
		}
	}

	return nil
}
