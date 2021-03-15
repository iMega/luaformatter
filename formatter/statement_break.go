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

import "io"

type breakStatement struct{}

func (breakStatement) InnerStatement(*element, *element) (bool, statement) {
	return false, nil
}

func (breakStatement) TypeOf() typeStatement {
	return tsNone
}

func (s *breakStatement) IsEnd(*element, *element) (bool, bool) {
	return false, true
}

func (s *breakStatement) Append(*element) {}

func (s *breakStatement) AppendStatement(st statement) {}

func (s *breakStatement) GetBody(statement, *element) statement {
	return nil
}

func (s *breakStatement) GetStatement(*element, *element) statement {
	return nil
}

func (s *breakStatement) Format(c *Config, p printer, w io.Writer) error {
	return writeKeyword(c, nBreak, w)
}
