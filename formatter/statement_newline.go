package formatter

import (
	"io"
)

type newlineStatement struct{}

func (newlineStatement) New() statementIntf {
	return &newlineStatement{}
}

func (newlineStatement) InnerStatement(prev, cur *element) statementIntf {
	return nil
}

func (newlineStatement) TypeOf() typeStatement {
	return tsNone
}

func (s *newlineStatement) IsEnd(prev, cur *element) bool {
	return true
}

func (s *newlineStatement) Append(el *element) {}

func (s *newlineStatement) AppendStatement(st statementIntf) {}

func (s *newlineStatement) GetBody(prevSt statementIntf, cur *element) statementIntf {
	return prevSt
}

func (s *newlineStatement) Format(c *Config, p printer, w io.Writer) error {
	return newLine(w)
}
