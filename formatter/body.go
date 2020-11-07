package formatter

import "io"

type body struct {
	Blocks map[uint64]block
	Qty    uint64
}

func (body) New() statementIntf {
	return &body{
		Blocks: make(map[uint64]block),
	}
}

func (b *body) GetBody(prevSt statementIntf, cur *element) statementIntf {
	return prevSt
}

func (body) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (body) TypeOf() typeStatement { return tsBody }

func (b *body) IsEnd(prev, cur *element) bool {
	if cur.Token.Type == nEnd {
		return true
	}

	if cur.Token.Type == nElseif {
		return true
	}

	return false
}

func (b *body) Append(el *element) {}

func (b *body) AppendStatement(st statementIntf) {
	if _, ok := st.(*prefixexpStatement); ok {
		return
	}

	b.Blocks[b.Qty] = newBloc(st)
	b.Qty++
}

func (b *body) Format(c *Config, p printer, w io.Writer) error {
	for i := 0; i < int(b.Qty); i++ {
		b := b.Blocks[uint64(i)]
		if err := b.Format(c, p, w); err != nil {
			return err
		}
	}

	return nil
}

func (b *block) Format(c *Config, p printer, w io.Writer) error {
	if err := p.WritePad(w); err != nil {
		return err
	}

	if s := b.Statement.Assignment; s != nil {
		if err := s.Format(c, p, w); err != nil {
			return err
		}
	}

	if st := b.Statement.If; st != nil {
		if err := st.Format(c, p, w); err != nil {
			return err
		}
	}

	if s := b.Statement.FuncCall; s != nil {
		if err := s.Format(c, p, w); err != nil {
			return err
		}
	}

	if s := b.Statement.Function; s != nil {
		if err := s.Format(c, p, w); err != nil {
			return err
		}
	}

	if s := b.Return; s != nil {
		if err := s.Format(c, p, w); err != nil {
			return err
		}
	}

	if s := b.Statement.Comment; s != nil {
		if err := s.Format(c, p, w); err != nil {
			return err
		}
	}

	if s := b.Statement.NewLine; s != nil {
		if err := s.Format(c, p, w); err != nil {
			return err
		}
	}

	if s := b.Statement.NewLine; s == nil {
		newLine(w)
	}

	return nil
}

func newBloc(st statementIntf) block {
	bl := block{}

	switch v := st.(type) {
	case *assignmentStatement:
		bl.Statement = statement{Assignment: v}

	case *labelStatement:
		bl.Statement = statement{Label: v}

	case *gotoStatement:
		bl.Statement = statement{Goto: v}

	case *breakStatement:
		bl.Statement = statement{Break: v}

	case *doStatement:
		bl.Statement = statement{Do: v}

	case *whileStatement:
		bl.Statement = statement{While: v}

	case *repeatStatement:
		bl.Statement = statement{Repeat: v}

	case *functionStatement:
		bl.Statement = statement{Function: v}

	case *funcCallStatement:
		bl.Statement = statement{FuncCall: v}

	case *prefixexpStatement:
		bl.Statement = statement{
			FuncCall: &funcCallStatement{
				Prefixexp: v,
			},
		}

	case *ifStatement:
		bl.Statement = statement{If: v}

	case *returnStatement:
		bl.Return = v

	case *commentStatement:
		bl.Statement = statement{Comment: v}

	case *newlineStatement:
		bl.Statement = statement{NewLine: v}
	}

	return bl
}

type block struct {
	Statement statement
	Return    *returnStatement
}
