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
	MaxWidth  int
	Bod       statementIntf
	Body      map[uint64]Block
	QtyBlocks uint64
}

func NewDocument() *document {
	return &document{
		Body: make(map[uint64]Block),
	}
}

func (d *document) AddBlock(b Block) {
	d.Body[d.QtyBlocks] = b
	d.QtyBlocks++
}

type varlist []element // separator ,

func newExp(elOrStat interface{}) *exp {
	e := &exp{}

	return e
}

type tableconstructor struct {
	FieldList fieldList
	Separator int
}

type fieldList struct {
	sqareField sqareField
	Field      field
	Exp        exp
}

type sqareField struct {
	ExpKey exp
	ExpVal exp
}

type functiondef struct {
	StartElement *element
	Parlist      parlist
	EndElement   *element
}

type parlist []*element

type binop struct {
	FirstExp  *exp
	Operator  element
	SecondExp exp
}

type unop struct {
	Operator *element
	Exp      *exp
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
	New() statementIntf
	Append(*element)
	AppendStatement(statementIntf)
	InnerStatement(prev, cur *element) statementIntf
	IsEnd(prev, cur *element) bool
	TypeOf() typeStatement
	GetBody(prevSt statementIntf, cur *element) statementIntf
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

func newStatement(st statementIntf) statement {
	stat := statement{}

	switch v := st.(type) {
	case *assignmentStatement:
		stat.Assignment = v
	// case *functionCallStatement:
	// 	stat.FunctionCall = v
	// case *labelStatement:
	// 	stat.Label = v
	// case *breakStatement:
	// 	stat.Break = v
	// case *gotoStatement:
	// 	stat.Goto = v
	// case *doStatement:
	// 	stat.Do = v
	// case *whileStatement:
	// 	stat.While = v
	// case *repeatStatement:
	// 	stat.Repeat = v
	// case *ifStatement:
	// 	stat.If = v
	// case *forStatement:
	// 	stat.ForNumerical = v
	// case *forGenericStatement:
	// 	stat.ForGeneric = v
	case *functionStatement:
		stat.Function = v
	}

	return stat
}

type functionCallStatement struct {
	Element *element
	Args    *explist
}

type varPart struct {
	Element *element
	Exp     exp
}

type namelist []*element

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
