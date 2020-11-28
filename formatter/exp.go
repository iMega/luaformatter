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
	"sort"
)

type exp struct {
	Element   *element           // nil | false | true | Numeral | LiteralString | ‘...’
	Table     *tableStatement    // {
	Func      *functionStatement // function
	Binop     *element
	Unop      *element
	Exp       *exp
	Prefixexp *prefixexpStatement
}

func (exp) New() statementIntf {
	return &exp{}
}

func (exp) InnerStatement(prev, cur *element) statementIntf {
	switch cur.Token.Type {
	case nFunction:
		return &functionStatement{}
	case nCurlyBracket:
		return &tableStatement{}
	}

	return nil
}

func (exp) TypeOf() typeStatement {
	return tsExp
}

func (s *exp) IsEnd(prev, cur *element) (bool, bool) {
	if cur.Token.Type == nEnd {
		return false, true
	}

	var syntax = map[tokenID]map[tokenID]bool{
		nNumber: {
			nAddition:   false, // 3 +
			nInequality: false, // 3 ~=
			nEquality:   false, // 3 ==
			nAnd:        false, // 3 and
			nOr:         false, // 3 or
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

func (s *exp) HasSyntax(el element) bool {
	return false
}

func isExp(el *element) bool {
	exps := []int{
		nID,           // 1
		nFalse,        // 11
		nFunction,     // 13
		nNil,          // 18
		nTrue,         // 22
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

	idx := sort.SearchInts(exps, el.Token.Type)

	return idx < len(exps) && exps[idx] == el.Token.Type
}

func (s *exp) Append(el *element) {
	if el.Token.Type == nComma {
		return
	}

	if el.Token.Type == nAssign {
		return
	}

	if el.Token.Type == nIn {
		return
	}

	if el.Token.Type == nParentheses || el.Token.Type == nClosingParentheses {
		return
	}

	if el.Token.Type == nSquareBracket || el.Token.Type == nClosingSquareBracket {
		return
	}

	if s.Element == nil {
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

	s.Element = el
}

func (s *exp) AppendStatement(st statementIntf) {
	switch v := st.(type) {
	case *tableStatement:
		s.Table = v
	case *functionStatement:
		s.Func = v
	case *exp:
		s.Exp = v
	case *prefixexpStatement:
		if s.Element != nil {
			v.Element = s.Element
			s.Element = nil
		}

		s.Prefixexp = v
	}
}

func (s *exp) GetBody(prevSt statementIntf, cur *element) statementIntf {
	return prevSt
}

func (s *exp) Format(c *Config, p printer, w io.Writer) error {
	if st := s.Unop; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
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
