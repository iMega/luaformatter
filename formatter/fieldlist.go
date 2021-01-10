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
	return tsFieldList
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

	t := c.Alignment.Table
	if t.KeyValuePairs || t.Comments {
		if p.ParentStatement == tsTable {
			fl = s.Align(c, p)
		}
	}

	for i, v := range s.List {
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
			if err := p.WritePad(w); err != nil {
				return err
			}
		}

		if i == 0 && v.Key.Comments != nil {
			for _, com := range v.Key.Comments {
				if com.Token.Type != nComment {
					break
				}

				if _, err := w.Write([]byte("-- ")); err != nil {
					return err
				}

				if err := com.Format(c, p, w); err != nil {
					return err
				}

				if err := newLine(w); err != nil {
					return err
				}
			}
		}

		fieldPrinter := p
		fieldPrinter.SpacesBeforeAssign = fl[uint64(i)].Key
		if err := v.Format(c, fieldPrinter, w); err != nil {
			return err
		}

		if p.ParentStatement != tsTable {
			if i < len(s.List)-1 {
				if _, err := w.Write([]byte(", ")); err != nil {
					return err
				}
			}
		}

		if p.ParentStatement == tsTable {
			if _, err := w.Write([]byte(",")); err != nil {
				return err
			}

			if i+1 < len(s.List) {
				com := s.List[i+1].Key.Comments
				if com != nil && len(com) > 0 && com[0].Token.Type == nComment {
					if err := p.WriteSpaces(w, int(fl[uint64(i)].Val)); err != nil {
						return err
					}

					if _, err := w.Write([]byte(" -- ")); err != nil {
						return err
					}

					if _, err := w.Write(com[0].Token.Lexeme); err != nil {
						return err
					}
				}
			}

			if err := newLine(w); err != nil {
				return err
			}
		}
	}

	return nil
}

type fieldLength struct {
	Key uint8
	Val uint8
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

		if item.Val != nil && item.Val.Func != nil {
			for b, v := range alignBlock {
				res[b] = fieldLength{
					Key: MaxKeyLength - v.Key,
					Val: MaxValueLength - v.Val,
				}
			}

			alignBlock = make(map[uint64]fieldLength)
			MaxKeyLength = 0
			MaxValueLength = 0

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
