package formatter

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

// type codeblock struct {
// 	Formatter Formatter
// 	Pad       uint64
// 	Start     []element
// 	Body      map[uint64]codeblock
// 	QtyBlocks uint64
// 	End       []element
// 	List      map[uint64]ListItem
// }

// func (b *codeblock) AddBlock(cb codeblock) {
// 	b.Body[b.QtyBlocks] = cb
// 	b.QtyBlocks++
// }

// type fieldlist struct {
// 	FirstSqExp *codeblock
// 	FirstExp   *codeblock
// 	SecondExp  *codeblock
// }

// func (i *fieldlist) GetCodeBlock(tokenID int) *codeblock {
// 	cb := &codeblock{}
// 	switch tokenID {
// 	case nSquareBracket:
// 		i.FirstSqExp = cb
// 	case nEq:
// 		i.SecondExp = cb
// 	default:
// 		i.FirstExp = cb
// 	}
// 	return cb
// }

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

type field struct {
	Name   element
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
)

type statement struct {
	Assignment   *assignmentStatement
	FunctionCall *functionCallStatement
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

type doStatement struct {
	DoElement  *element
	Body       []Block
	EndElement *element
}

type whileStatement struct {
	StartElement *element
	doStatement
}

type repeatStatement struct {
	StartElement *element
	Body         []Block
	EndElement   *element
	Exp          exp
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
