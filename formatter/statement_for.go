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

type forStatement struct {
	FieldList *fieldlist
	Explist   *explist
	Body      statement
}

func (forStatement) InnerStatement(prev, cur *element) (bool, statement) {
	return false, &fieldlist{}
}

func (forStatement) TypeOf() typeStatement {
	return tsFor
}

func (s *forStatement) IsEnd(prev, cur *element) (bool, bool) {
	return false, cur.Token.Type == nEnd
}

func (s *forStatement) Append(el *element) {}

func (s *forStatement) AppendStatement(st statement) {
	switch v := st.(type) {
	case *fieldlist:
		s.FieldList = v
	case *explist:
		s.Explist = v
	}
}

func (s *forStatement) GetBody(prevSt statement, cur *element) statement {
	if cur.Token.Type != nDo {
		return prevSt
	}

	if s.Body == nil {
		s.Body = new(body).New()
	}

	return s.Body
}

func (s *forStatement) GetStatement(prev, cur *element) statement {
	if cur.Token.Type == nFor {
		return &forStatement{}
	}

	// if cur.Token.Type == nIn {
	// 	return &explist{}
	// }

	if prev != nil && prev.Token.Type == nIn {
		return &explist{}
	}

	if isExp(cur) {
		return &field{}
	}

	return nil
}

func (s *forStatement) Format(c *Config, p printer, w io.Writer) error {
	if err := writeKeywordWithSpaceRight(c, nFor, w); err != nil {
		return err
	}

	if s.FieldList != nil {
		if err := s.FieldList.Format(c, p, w); err != nil {
			return err
		}
	}

	if s.Explist != nil {
		if err := writeKeywordWithSpace(c, nIn, w); err != nil {
			return err
		}

		if err := s.Explist.Format(c, p, w); err != nil {
			return err
		}
	}

	if st, ok := s.Body.(*body); ok {
		if err := space(w); err != nil {
			return err
		}

		np := p
		np.ParentStatement = s.TypeOf()
		np.IgnoreFirstPad = true

		if err := st.Format(c, np, w); err != nil {
			return err
		}
	}

	return nil
}
