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

type fieldlist struct {
	List []*field
}

func (fieldlist) InnerStatement(prev, cur *element) (bool, statement) {
	// fieldlist always returns a field. need will add
	// it will need to be added to the innerStatement.
	// return &field{}
	return false, nil
}

func (fieldlist) TypeOf() typeStatement {
	return tsNone
}

func (s *fieldlist) IsEnd(prev, cur *element) (bool, bool) {
	t := []int{
		nDo,                  // 7
		nIn,                  // 16
		nClosingCurlyBracket, // 32
	}

	return false, binarySearch(t, cur.Token.Type)
}

func (s *fieldlist) Append(el *element) {}

func (s *fieldlist) AppendStatement(st statement) {
	if v, ok := st.(*field); ok {
		s.List = append(s.List, v)
	}
}

func (s *fieldlist) GetBody(prevSt statement, cur *element) statement {
	return prevSt
}

func (s *fieldlist) GetStatement(prev, cur *element) statement {
	if cur.Token.Type == nComma {
		return nil
	}

	return &field{Square: cur.Token.Type == nSquareBracket}
}

func (s *fieldlist) Format(c *Config, p printer, w io.Writer) error {
	var fl map[uint64]fieldLength

	isInLine := s.isInline(c, p, w)
	t := c.Alignment.Table
	if t.KeyValuePairs || t.Comments {
		if p.ParentStatement == tsTable {
			fl = s.Align(c, p)
		}
	}

	for i := 0; i < len(s.List); i++ {
		v := s.List[i]

		if v.Key.Comments != nil {
			comPrinter := p
			comPrinter.Pad += c.IndentSize

			for n := 0; n < len(v.Key.Comments); n++ {
				com := v.Key.Comments[uint64(n)]
				if com.Token.Type != nComment {
					break
				}

				if i == 0 && n == 0 {
					if err := newLine(w); err != nil {
						return err
					}

					if err := comPrinter.WritePad(w); err != nil {
						return err
					}
				}

				pad := 1
				if n > 0 {
					pad = 0

					if err := newLine(w); err != nil {
						return err
					}

					if err := comPrinter.WritePad(w); err != nil {
						return err
					}
				}

				if i > 0 {
					if n == 0 {
						if err := comPrinter.WriteSpaces(w, int(fl[uint64(i-1)].Val)+pad); err != nil {
							return err
						}
					}
				}

				if _, err := w.Write([]byte("-- ")); err != nil {
					return err
				}

				if err := com.Format(c, comPrinter, w); err != nil {
					return err
				}
			}
		}

		if v.Key.Element == nil &&
			v.Key.Table == nil &&
			v.Key.Func == nil &&
			v.Key.Binop == nil &&
			v.Key.Unop == nil &&
			v.Key.Exp == nil &&
			v.Key.Prefixexp == nil {
			continue
		}

		if p.ParentStatement == tsTable {
			if i < len(s.List) && !isInLine {
				if err := newLine(w); err != nil {
					return err
				}
			}
		}

		fieldPrinter := p
		if isInLine {
			fieldPrinter.Pad = 0
		} else {
			fieldPrinter.Pad += c.IndentSize
		}
		fieldPrinter.SpacesBeforeAssign = fl[uint64(i)].Key
		if err := v.Format(c, fieldPrinter, w); err != nil {
			return err
		}

		if p.ParentStatement != tsTable || isInLine {
			if i < len(s.List)-1 {
				if _, err := w.Write([]byte(", ")); err != nil {
					return err
				}
			}
		}

		if p.ParentStatement == tsTable && !isInLine {
			if _, err := w.Write([]byte(",")); err != nil {
				return err
			}
		}
	}

	if p.ParentStatement == tsTable && !isInLine {
		if err := newLine(w); err != nil {
			return err
		}

		if err := p.WritePad(w); err != nil {
			return err
		}
	}

	return nil
}

type fieldLength struct {
	Key uint8
	Val uint8
}

func (s *fieldlist) isInline(c *Config, p printer, w io.Writer) bool {
	var curpos cursorPosition

	if v, ok := w.(cursorPositioner); ok {
		curpos = v.Cursor()
	}

	buf := bytes.NewBuffer([]byte("{}"))

	values := 0
	keyType := -1
	isVector := false

	for i := 0; i < len(s.List); i++ {
		item := s.List[i]

		isEnded := isEndedAlignedBlock(item)
		if isEnded || item.HasComment() {
			return false
		}

		if err := item.Key.Format(c, p, buf); err != nil {
			return false
		}

		if item.Key.Element != nil {
			if keyType == -1 {
				keyType = item.Key.Element.Token.Type
			}

			if keyType == item.Key.Element.Token.Type {
				isVector = true
			}
		}

		if s.List[i].Val != nil {
			values++
			buf.WriteString(" = ")
			if err := s.List[i].Val.Format(c, p, buf); err != nil {
				return false
			}
		}

		if i < len(s.List)-1 {
			buf.WriteString(", ")
		}

		curpos.Col += uint64(buf.Len())
		buf.Reset()
		if curpos.Col > uint64(c.MaxLineLength+1) {
			return false
		}

		if i > 2 && values > 0 {
			return false
		}
	}

	if !isVector {
		return false
	}

	if isVector && len(s.List) > 5 {
		return false
	}

	return true
}

func (s *fieldlist) Align(c *Config, p printer) map[uint64]fieldLength {
	var (
		MaxKeyLength   uint8
		MaxValueLength uint8

		res = make(map[uint64]fieldLength)

		alignBlock = make(map[uint64]fieldLength)
		w          = bytes.NewBuffer(nil)
	)

	for i := 0; i < len(s.List); i++ {
		item := s.List[i]

		isEnded := isEndedAlignedBlock(item)
		if isEnded || isStartAlignedBlock(item) {
			for b, v := range alignBlock {
				res[b] = fieldLength{
					Key: MaxKeyLength - v.Key,
					Val: MaxValueLength - v.Val,
				}
			}

			alignBlock = make(map[uint64]fieldLength)
			MaxKeyLength = 0
			MaxValueLength = 0
		}

		if isEnded {
			continue
		}

		if s.List[i].Square {
			w.WriteString("[]")
		}

		if err := s.List[i].Key.Format(c, p, w); err != nil {
			return res
		}

		kl := uint8(w.Len())
		w.Reset()

		if s.List[i].Val != nil {
			if err := s.List[i].Val.Format(c, p, w); err != nil {
				return res
			}
		}

		vl := uint8(w.Len())
		w.Reset()

		alignBlock[uint64(i)] = fieldLength{Key: kl, Val: vl}

		if MaxKeyLength < kl {
			MaxKeyLength = kl
		}

		if MaxValueLength < vl {
			MaxValueLength = vl
		}
	}

	for b, v := range alignBlock {
		res[b] = fieldLength{
			Key: MaxKeyLength - v.Key,
			Val: MaxValueLength - v.Val,
		}
	}

	return res
}

func isStartAlignedBlock(f *field) bool {
	if f.Key.Table != nil {
		return true
	}

	if f.Key.Comments != nil && len(f.Key.Comments) > 1 {
		for i := 0; i < len(f.Key.Comments); i++ {
			if i > 0 && f.Key.Comments[uint64(i)].Token.Type == nComment {
				return true
			}
		}
	}

	return false
}

func isEndedAlignedBlock(f *field) bool {
	return f.Val != nil && f.Val.Func != nil
}
