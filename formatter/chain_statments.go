package formatter

type chainStatments struct {
	chain []statementIntf
}

func (cs *chainStatments) Append(st statementIntf) {
	cs.chain = append(cs.chain, st)
}

func (cs *chainStatments) Prev() statementIntf {
	if len(cs.chain) == 0 {
		return nil
	}

	cs.chain = cs.chain[:len(cs.chain)-1]

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
