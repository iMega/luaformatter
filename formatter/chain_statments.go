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

import "errors"

type chainStatments struct {
	chain []statement
}

var errDuplicateStatements = errors.New("statement does exists")

func (cs *chainStatments) Append(st statement) error {
	for i := len(cs.chain) - 1; i >= 0; i-- {
		if cs.chain[i] == st {
			return errDuplicateStatements
		}
	}

	cs.chain = append(cs.chain, st)

	return nil
}

func (cs *chainStatments) ExtractPrev() statement {
	cs.chain = cs.chain[:len(cs.chain)-1]

	if len(cs.chain) == 0 {
		return nil
	}

	return cs.chain[len(cs.chain)-1]
}

func (cs *chainStatments) Prev() statement {
	if len(cs.chain) == 0 {
		return nil
	}

	return cs.chain[len(cs.chain)-1]
}

func (cs *chainStatments) Len() int {
	return len(cs.chain)
}

func (cs *chainStatments) First() statement {
	return cs.chain[0]
}

func (cs *chainStatments) ExctractStatement(ts typeStatement) statement {
	item := -1

	for i := len(cs.chain) - 1; i >= 0; i-- {
		if cs.chain[i].TypeOf() == ts {
			item = i
		}

		if cs.chain[i].TypeOf() == tsBody {
			break
		}
	}

	if item == -1 {
		return nil
	}

	v := cs.chain[item]
	cs.chain = cs.chain[:item]

	return v
}

func (cs *chainStatments) GetLastBody() statement {
	for i := len(cs.chain) - 1; i >= 0; i-- {
		if v, ok := cs.chain[i].(*body); ok {
			return v
		}
	}

	return nil
}
