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

type exp struct {
	Element   *element
	Table     *tableStatement
	Func      *functionStatement
	Binop     *element
	Unop      *element
	Exp       *exp
	Prefixexp *prefixexpStatement
	Comments  map[uint64]*element
}

func (exp) InnerStatement(prev, cur *element) (bool, statement) {
	switch cur.Token.Type {
	case nFunction:
		return false, &functionStatement{}

	case nCurlyBracket:
		return false, &tableStatement{}

	case nParentheses:
		if prev != nil {
			if prev.Token.Type == nID || prev.Token.Type == nClosingSquareBracket {
				return false, nil // funcCallStatement, asd["adf"]()
			}

			if prev.Token.Type == nFunction {
				return false, nil
			}
		}

		return false, &prefixexpStatement{Enclosed: true}
	}

	return false, nil
}

func (exp) TypeOf() typeStatement {
	return tsExp
}

type tokenID int

func (s *exp) IsEnd(prev, cur *element) (bool, bool) {
	if cur.Token.Type == nEnd {
		return false, true
	}

	if isExp(prev) && isBinop(cur) {
		return false, false
	}

	if isBinop(prev) && isExp(cur) {
		return false, false
	}

	if isUnop(prev) && isExp(cur) {
		return false, false
	}

	if prev != nil && prev.Token.Type == nSquareBracket && isExp(cur) {
		return false, false
	}

	if cur.Token.Type == nParentheses {
		return false, false
	}

	if prev != nil && prev.Token.Type == nDot || cur.Token.Type == nDot {
		return false, false
	}

	if isBinop(cur) {
		return false, false
	}

	if prev != nil && prev.Token.Type == nParentheses && isExp(cur) {
		return false, false // ((1+1)+1)+1
	}

	if prev != nil && prev.Token.Type == nAssign && isExp(cur) {
		return false, false // = exp
	}

	if cur.Token.Type == nCommentLong {
		return false, false
	}

	// return false, true

	var syntax = map[tokenID]map[tokenID]bool{
		nNumber: {
			nSubtraction:        false, // 3 -
			nConcat:             false, // 3 ..
			nMultiplication:     false, // 3 *
			nFloatDivision:      false, // 3 /
			nBitwiseAND:         false, // 3 &
			nModulo:             false, // 3 %
			nExponentiation:     false, // 3 ^
			nAddition:           false, // 3 +
			nLessThan:           false, // 3 <
			nLeftShift:          false, // 3 <<
			nLessOrEqual:        false, // 3 <=
			nEquality:           false, // 3 ==
			nGreaterThan:        false, // 3 >
			nGreaterOrEqual:     false, // 3 >=
			nBitwiseOR:          false, // 3 |
			nBitwiseExclusiveOR: false, // 3 ~
			nInequality:         false, // 3 ~=
			nAnd:                false, // 3 and
			nOr:                 false, // 3 or
		},
		nID: {
			nAddition:      false, // id +
			nInequality:    false, // id ~=
			nEquality:      false, // id ==
			nSquareBracket: false, // id[
			nParentheses:   false, // id(
			nCurlyBracket:  false, // id {
			nString:        false, // id "string"
			nLessThan:      false, // id <
			nGreaterThan:   false, // id >
			nAnd:           false, // id and
		},
		nAnd: {
			nLength: false, // and #
		},
		nOr: {
			nLength: false, // or #
		},
		nAddition: {
			nNumber: false, // + 3
			nID:     false, // + id
		},
		nInequality: {
			nNumber: false, // ~= 3
		},
		nEquality: {
			nNumber: false, // == 3
			nString: false, // == "string"
			nID:     false, // == id
		},
		nSquareBracket: {
			nString: false, // ["string"
		},
		nParentheses: {
			nID:           false, // (id
			nString:       false, // ("string"
			nCurlyBracket: false, // ({
		},
		nAssign: {
			nFunction: false, // = function
			nID:       false, // = id
		},
		nSubtraction: {
			nNumber: false, // -3
		},
		nLessThan: {
			nNumber: false, // < 3
		},
		nMultiplication: {
			nNumber: false, // * 3
		},
		nFloatDivision: {
			nNumber: false, // / 3
		},
		nGreaterThan: {
			nNumber: false, // > 3
		},
		nLength: {
			nID: false, // #id
		},
		nClosingParentheses: {
			nConcat: false, // ) ..
		},
		nConcat: {
			nID:     false, // .. id
			nString: false, // .. "string"
		},
		nString: {
			nConcat: false, // "string" ..
		},
	}

	v, ok := syntax[tokenID(prev.Token.Type)]
	if !ok {
		return false, true
	}

	res, ok := v[tokenID(cur.Token.Type)]
	if !ok {
		return false, true
	}

	return false, res
}

func isBinop(el *element) bool {
	binop := []int{
		nAnd,                // 44
		nOr,                 // 45
		nAddition,           // 46
		nSubtraction,        // 47
		nMultiplication,     // 48
		nFloatDivision,      // 49
		nModulo,             // 51
		nExponentiation,     // 52
		nBitwiseAND,         // 53
		nBitwiseOR,          // 54
		nBitwiseExclusiveOR, // 55
		nLeftShift,          // 56
		nRightShift,         // 57
		nConcat,             // 59
		nEquality,           // 60
		nInequality,         // 61
		nLessThan,           // 62
		nGreaterThan,        // 63
		nLessOrEqual,        // 64
		nGreaterOrEqual,     // 65
	}

	return binarySearch(binop, el.Token.Type)
}

func isUnop(el *element) bool {
	ops := []int{
		nSubtraction,        // 47
		nBitwiseExclusiveOR, // 55
		nLength,             // 58
		nNot,                // 66
	}

	return binarySearch(ops, el.Token.Type)
}

func isExp(el *element) bool {
	exps := []int{
		nID,           // 1
		nFalse,        // 11
		nFunction,     // 13
		nNil,          // 18
		nTrue,         // 22
		nParentheses,  // 27 // c = (1-2)
		nCurlyBracket, // 31
		nNumber,       // 38
		nString,       // 39
		nVararg,       // 41

		// unop
		nSubtraction,        // 47
		nBitwiseExclusiveOR, // 55
		nLength,             // 58
		nNot,                // 65

	} //  exp binop exp

	return binarySearch(exps, el.Token.Type)
}

func (s *exp) Append(el *element) {
	types := []int{
		nIn,                   // 16
		nReturn,               // 20
		nParentheses,          // 27
		nClosingParentheses,   // 28
		nSquareBracket,        // 29
		nClosingSquareBracket, // 30
		nClosingCurlyBracket,  // 32
		nAssign,               // 33
		nComma,                // 34
	}
	if binarySearch(types, el.Token.Type) {
		return
	}

	if s.Element == nil && s.Prefixexp == nil {
		switch el.Token.Type {
		case nSubtraction:
			s.Unop = el

			return

		case nNot:
			s.Unop = el

			return

		case nLength:
			s.Unop = el

			return

		case nBitwiseExclusiveOR:
			s.Unop = el

			return
		}
	}

	if el.Token.Type >= nAnd && el.Token.Type <= nGreaterOrEqual {
		s.Binop = el

		return
	}

	if el.Token.Type == nCommentLong {
		if s.Comments == nil {
			s.Comments = make(map[uint64]*element)
		}

		s.Comments[uint64(len(s.Comments))] = el

		return
	}

	s.Element = el
}

func (s *exp) AppendStatement(st statement) {
	switch v := st.(type) {
	case *tableStatement:
		s.Table = v
	case *functionStatement:
		s.Func = v
	case *exp:
		s.Exp = v
	case *funcCallStatement:
		s.Prefixexp = &prefixexpStatement{FuncCall: v} // a = func[0].call{}

	case *prefixexpStatement:
		if s.Element != nil {
			v.Element = s.Element
			s.Element = nil
		}

		s.Prefixexp = v
	}
}

func (s *exp) GetBody(prevSt statement, cur *element) statement {
	return prevSt
}

func (s *exp) GetStatement(prev, cur *element) statement {
	if prev != nil {
		switch prev.Token.Type {
		case nID:
			if cur.Token.Type == nString {
				return &prefixexpStatement{}
			}

			if cur.Token.Type == nParentheses {
				return &prefixexpStatement{} //funcCallStatement 111
			}

			if cur.Token.Type == nCurlyBracket {
				return &prefixexpStatement{}
			}
			// 		if cur.Token.Type == nParentheses {
			// 			return &funcCallStatement{}
			// 		}
		case nParentheses:
			if cur.Token.Type == nParentheses {
				return &prefixexpStatement{} // 1+((1+2)+2)
			}
		}
	}

	// funcCallStatement
	// if cur.Token.Type == nParentheses {
	// 	return &prefixexpStatement{} //
	// }

	if cur.Token.Type == nDot {
		return &prefixexpStatement{}
	}

	if cur.Token.Type == nSquareBracket {
		return &prefixexpStatement{}
	}

	if cur.Token.Type == nCurlyBracket {
		return &tableStatement{}
	}

	// if isExp(prev) && isBinop(cur) {
	// 	return false, false
	// }

	if isBinop(prev) && isExp(cur) {
		return &exp{}
	}

	return nil
}

func (s *exp) Format(c *Config, p printer, w io.Writer) error {
	if s.Comments != nil {
		for i := 0; i < len(s.Comments); i++ {
			if _, err := w.Write([]byte("--[[ ")); err != nil {
				return err
			}

			if err := s.Comments[uint64(i)].Format(c, p, w); err != nil {
				return err
			}

			if _, err := w.Write([]byte(" ]]")); err != nil {
				return err
			}
		}
	}

	if st := s.Unop; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}

		if st.Token.Type == nNot {
			if err := space(w); err != nil {
				return err
			}
		}
	}

	if st := s.Element; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	if st := s.Table; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	if st := s.Func; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	if st := s.Prefixexp; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	if st := s.Binop; st != nil {
		if err := space(w); err != nil {
			return err
		}

		if err := st.Format(c, p, w); err != nil {
			return err
		}

		if err := space(w); err != nil {
			return err
		}
	}

	if st := s.Exp; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	return nil
}
