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

type chainStatments struct {
	chain []statementIntf
}

func (cs *chainStatments) Append(st statementIntf) {
	for i := len(cs.chain) - 1; i >= 0; i-- {
		if cs.chain[i] == st {
			return
		}
	}

	cs.chain = append(cs.chain, st)
}

func (cs *chainStatments) Reset() {
	cs.chain = nil
}

func (cs *chainStatments) ExtractPrev() statementIntf {
	if len(cs.chain) == 0 {
		return nil
	}

	cs.chain = cs.chain[:len(cs.chain)-1]

	if len(cs.chain) == 0 {
		return nil
	}

	return cs.chain[len(cs.chain)-1]
}

func (cs *chainStatments) Prev() statementIntf {
	if len(cs.chain) == 0 {
		return nil
	}

	return cs.chain[len(cs.chain)-1]
}

func (cs *chainStatments) Len() int {
	return len(cs.chain)
}

func (cs *chainStatments) First() statementIntf {
	return cs.chain[0]
}

func (cs *chainStatments) ExctractAssignStatement() statementIntf {
	item := -1
	for i := len(cs.chain) - 1; i >= 0; i-- {
		if _, ok := cs.chain[i].(*assignmentStatement); ok {
			item = i
		}
	}

	if item == -1 {
		return nil
	}

	v, _ := cs.chain[item].(*assignmentStatement)
	cs.chain = cs.chain[:item]

	return v
}

func (cs *chainStatments) ExctractPrefixexp() statementIntf {
	item := -1
	for i := len(cs.chain) - 1; i >= 0; i-- {
		if _, ok := cs.chain[i].(*prefixexpStatement); ok {
			item = i
		}
	}

	if item == -1 {
		return nil
	}

	v, _ := cs.chain[item].(*prefixexpStatement)
	cs.chain = cs.chain[:item]

	return v
}

func (cs *chainStatments) GetLastBody() statementIntf {
	for i := len(cs.chain) - 1; i >= 0; i-- {
		if v, ok := cs.chain[i].(*body); ok {
			return v
		}
	}

	return nil
}
