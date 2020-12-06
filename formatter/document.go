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

type document struct {
	MaxWidth int
	Body     statementIntf
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
	tsRepeat
)

// chunk ::= block
// block ::= {stat} [retstat]
// stat ::=  ‘;’ |
//      varlist ‘=’ explist |
//      functioncall |
//      label |
//      break |
//      goto Name |
//      do block end |
//      while exp do block end |
//      repeat block until exp |
//      if exp then block {elseif exp then block} [else block] end |
//      for Name ‘=’ exp ‘,’ exp [‘,’ exp] do block end |
//      for namelist in explist do block end |
//      function funcname funcbody |
//      local function Name funcbody |
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
