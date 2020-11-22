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
		// case nTable:
		// return &tableconstructor{}
	}

	return nil
}

func (exp) TypeOf() typeStatement {
	return tsExp
}

func (s *exp) IsEnd(prev, cur *element) bool {
	if cur.Token.Type == nEnd {
		return true
	}

	var syntax = map[tokenID]map[tokenID]bool{
		nNumber: {
			nAddition:   false,
			nInequality: false,
			nEquality:   false,
		},
		nID: {
			nAddition:      false,
			nInequality:    false,
			nEquality:      false,
			nSquareBracket: false,
			nParentheses:   false,
			nString:        false,
			nLessThan:      false, // alignment < 100
		},
		nAddition: {
			nNumber: false,
			nID:     false,
		},
		nInequality: {
			nNumber: false,
		},
		nEquality: {
			nNumber: false,
			nString: false,
			nID:     false, // if name == searched then
		},
		nSquareBracket: {
			nString: false,
		},
		nParentheses: {
			nID: false,
		},
		nAssign: {
			nFunction: false, // b = function() end
			nID:       false, // c = b()
		},
		nSubtraction: {
			nNumber: false, // -1
		},
		nLessThan: {
			nNumber: false, // alignment < 100
		},
	}

	v, ok := syntax[tokenID(prev.Token.Type)]
	if !ok {
		return true
	}

	res, ok := v[tokenID(cur.Token.Type)]
	if !ok {
		return true
	}

	return res
}

func (s *exp) HasSyntax(el element) bool {
	return false
}

func isExp(el *element) bool {
	exps := []int{nNil, nFalse, nTrue, nNumber, nString, nVararg, nFunction,
		nID, nCurlyBracket} // tableconstructor | exp binop exp | unop exp
	for _, e := range exps {
		if e == el.Token.Type {
			return true
		}
	}

	return false
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

	if st := s.Prefixexp; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	return nil
}
