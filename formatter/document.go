package formatter

import (
	"io"
)

type document struct {
	MaxWidth  int
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

func (b *Block) Format(c *Config, w io.Writer) error {
	if s := b.Statement.Assignment; s != nil {
		if err := s.Format(w); err != nil {
			return err
		}
	}

	if s := b.Statement.FuncCall; s != nil {
		if err := s.Format(c, w); err != nil {
			return err
		}
	}

	if s := b.Statement.Function; s != nil {
		if err := s.Format(c, w); err != nil {
			return err
		}
	}

	if s := b.Return; s != nil {
		if err := s.Format(c, w); err != nil {
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
}

type typeStatement int

const (
	tsAssignment = iota
	tsFunction
	tsIf
	tsReturn
	tsExp
	tsExpList
	tsPrefixexpStatement
	tsFuncCallStatement
)

type statement struct {
	Assignment   *assignmentStatement
	FuncCall     *funcCallStatement
	Label        *labelStatement
	Break        *breakStatement
	Goto         *gotoStatement
	Do           *doStatement
	While        *whileStatement
	Repeat       *repeatStatement
	If           *ifStatement
	ForNumerical *forNumericalStatement
	ForGeneric   *forGenericStatement
	Function     *functionStatement
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
	// case *forNumericalStatement:
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

type forNumericalStatement struct {
	IDStatement *element
	VarPart     varPart
	LimitPart   exp
	StepPart    *exp
	doStatement
}

type varPart struct {
	Element *element
	Exp     exp
}

type forGenericStatement struct {
	IDStatement *element
	Namelist    namelist
	InElement   element
	Explist     explist
	doStatement
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
