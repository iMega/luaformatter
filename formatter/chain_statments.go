package formatter

type chainStatments struct {
	chain []statementIntf
}

func (cs *chainStatments) Append(st statementIntf) {
	cs.chain = append(cs.chain, st)
}

func (cs *chainStatments) Reset() {
	cs.chain = nil
}

// func (cs *chainStatments) Prepend(st statementIntf) {
// 	cs.chain = append([]statementIntf{st}, cs.chain...)
// }

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
