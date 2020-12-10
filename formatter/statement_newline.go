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

type newlineStatement struct{}

func (newlineStatement) InnerStatement(prev, cur *element) (bool, statementIntf) {
	return true, nil
}

func (newlineStatement) TypeOf() typeStatement {
	return tsNone
}

func (s *newlineStatement) IsEnd(prev, cur *element) (bool, bool) {
	return false, true
}

func (s *newlineStatement) Append(el *element) {}

func (s *newlineStatement) AppendStatement(st statementIntf) {}

func (s *newlineStatement) GetBody(prevSt statementIntf, cur *element) statementIntf {
	return prevSt
}

func (s *newlineStatement) GetStatement(prev, cur *element) statementIntf {
	return nil
}

func (s *newlineStatement) Format(c *Config, p printer, w io.Writer) error {
	return newLine(w)
}
