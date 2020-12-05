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

type document struct {
	MaxWidth int
	Body     statementIntf
}

func newExp(elOrStat interface{}) *exp {
	e := &exp{}

	return e
}

type Block struct {
	Statement statement
	Return    *returnStatement
}

func (b *Block) Format(c *Config, p printer, w io.Writer) error {
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

	return nil
}

type statementIntf interface {
	Append(*element)
	AppendStatement(statementIntf)
	InnerStatement(prev, cur *element) statementIntf
	IsEnd(prev, cur *element) (bool, bool)
	TypeOf() typeStatement
	GetBody(prevSt statementIntf, cur *element) statementIntf
	GetStatement(prev, cur *element) statementIntf
}

type typeStatement int

const (
	tsNone = iota
	tsAssignment
	tsFunction
	tsIf
	tsBody
	tsReturn
	tsExp
	tsExpList
	tsPrefixexpStatement
	tsFuncCallStatement
	tsTable
	tsFieldList
	tsField
)

type statement struct {
	Assignment *assignmentStatement
	FuncCall   *funcCallStatement
	Label      *labelStatement
	Break      *breakStatement
	Goto       *gotoStatement
	Do         *doStatement
	While      *whileStatement
	Repeat     *repeatStatement
	If         *ifStatement
	For        *forStatement
	Function   *functionStatement
	Comment    *commentStatement
	NewLine    *newlineStatement
}

// chunk ::= block
// block ::= {stat} [retstat]
// stat ::=  ‘;’ |
//      varlist ‘=’ explist |
//      functioncall |
//      label |
//      break |
//      goto Name |
// b     do block end |
// b     while exp do block end |
// b     repeat block until exp |
// b     if exp then block {elseif exp then block} [else block] end |
// b     for Name ‘=’ exp ‘,’ exp [‘,’ exp] do block end |
// b     for namelist in explist do block end |
// b     function funcname funcbody |
// b     local function Name funcbody |
//      local namelist [‘=’ explist]
// retstat ::= return [explist] [‘;’]
// label ::= ‘::’ Name ‘::’
// funcname ::= Name {‘.’ Name} [‘:’ Name]
// varlist ::= var {‘,’ var}
// var ::=  Name | prefixexp ‘[’ exp ‘]’ | prefixexp ‘.’ Name
// namelist ::= Name {‘,’ Name}
// explist ::= exp {‘,’ exp}
// exp ::=  nil | false | true | Numeral | LiteralString | ‘...’ | functiondef |
//      prefixexp | tableconstructor | exp binop exp | unop exp
// prefixexp ::= var | functioncall | ‘(’ exp ‘)’
// functioncall ::=  prefixexp args | prefixexp ‘:’ Name args
// args ::=  ‘(’ [explist] ‘)’ | tableconstructor | LiteralString
// functiondef ::= function funcbody
// funcbody ::= ‘(’ [parlist] ‘)’ block end
// parlist ::= namelist [‘,’ ‘...’] | ‘...’
// tableconstructor ::= ‘{’ [fieldlist] ‘}’
// fieldlist ::= field {fieldsep field} [fieldsep]
// field ::= ‘[’ exp ‘]’ ‘=’ exp | Name ‘=’ exp | exp
// fieldsep ::= ‘,’ | ‘;’
// binop ::=  ‘+’ | ‘-’ | ‘*’ | ‘/’ | ‘//’ | ‘^’ | ‘%’ |
//      ‘&’ | ‘~’ | ‘|’ | ‘>>’ | ‘<<’ | ‘..’ |
//      ‘<’ | ‘<=’ | ‘>’ | ‘>=’ | ‘==’ | ‘~=’ |
//      and | or
// unop ::= ‘-’ | not | ‘#’ | ‘~’
