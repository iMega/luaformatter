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

type funcCallStatement struct {
	Prefixexp *prefixexpStatement
	Explist   *explist
}

func (funcCallStatement) New() statementIntf {
	return &funcCallStatement{}
}

func (funcCallStatement) InnerStatement(prev, cur *element) statementIntf {
	return &explist{}
}

func (funcCallStatement) TypeOf() typeStatement {
	return tsFuncCallStatement
}

func (s *funcCallStatement) IsEnd(prev, cur *element) (bool, bool) {
	if cur.Token.Type == nClosingParentheses {
		return true, true
	}

	return false, true
}

func (s *funcCallStatement) Append(el *element) {
}

func (s *funcCallStatement) AppendStatement(st statementIntf) {
	switch v := st.(type) {
	case *prefixexpStatement:
		s.Prefixexp = v
	case *explist:
		s.Explist = v
	}
}

func (s *funcCallStatement) GetBody(prevSt statementIntf, cur *element) statementIntf {
	return prevSt
}

func (s *funcCallStatement) GetStatement(prev, cur *element) statementIntf {
	return nil
}

func (s *funcCallStatement) Format(c *Config, p printer, w io.Writer) error {
	if st := s.Prefixexp; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	if _, err := w.Write([]byte("(")); err != nil {
		return err
	}

	if st := s.Explist; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	if _, err := w.Write([]byte(")")); err != nil {
		return err
	}

	return nil
}
