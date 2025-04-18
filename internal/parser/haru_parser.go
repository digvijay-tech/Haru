// Code generated from ./grammar/haru.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // haru

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type haruParser struct {
	*antlr.BaseParser
}

var HaruParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func haruParserInit() {
	staticData := &HaruParserStaticData
	staticData.LiteralNames = []string{
		"", "'print'", "';'", "'!'", "'('", "')'", "'**'", "'*'", "'/'", "'%'",
		"'+'", "'-'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", "'&&'",
		"'||'", "'['", "','", "']'", "'='", "'true'", "'false'", "'let'", "':'",
		"'mut'", "'const'", "'i8'", "'i16'", "'i32'", "'i64'", "'int'", "'ui8'",
		"'ui16'", "'ui32'", "'ui64'", "'uint'", "'f32'", "'f64'", "'bool'",
		"'string'", "'byte'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "WS", "ID", "NUMBER", "FLOAT",
		"STRING", "BYTE",
	}
	staticData.RuleNames = []string{
		"program", "statement", "printStmt", "expr", "assign", "literal", "varDecl",
		"type",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 50, 183, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 5, 0, 18, 8, 0, 10, 0, 12,
		0, 21, 9, 0, 1, 1, 1, 1, 1, 1, 3, 1, 26, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3,
		43, 8, 3, 10, 3, 12, 3, 46, 9, 3, 3, 3, 48, 8, 3, 1, 3, 1, 3, 1, 3, 3,
		3, 53, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 97, 8, 3, 10, 3,
		12, 3, 100, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 3, 5, 107, 8, 5, 1, 5,
		1, 5, 3, 5, 111, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 118, 8, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 140, 8, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 164, 8, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 3, 7, 181, 8, 7, 1, 7, 0, 1, 6, 8, 0, 2, 4, 6, 8, 10, 12,
		14, 0, 0, 224, 0, 19, 1, 0, 0, 0, 2, 25, 1, 0, 0, 0, 4, 27, 1, 0, 0, 0,
		6, 52, 1, 0, 0, 0, 8, 101, 1, 0, 0, 0, 10, 117, 1, 0, 0, 0, 12, 163, 1,
		0, 0, 0, 14, 180, 1, 0, 0, 0, 16, 18, 3, 2, 1, 0, 17, 16, 1, 0, 0, 0, 18,
		21, 1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 1, 1, 0, 0,
		0, 21, 19, 1, 0, 0, 0, 22, 26, 3, 4, 2, 0, 23, 26, 3, 12, 6, 0, 24, 26,
		3, 8, 4, 0, 25, 22, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 25, 24, 1, 0, 0, 0,
		26, 3, 1, 0, 0, 0, 27, 28, 5, 1, 0, 0, 28, 29, 3, 6, 3, 0, 29, 30, 5, 2,
		0, 0, 30, 5, 1, 0, 0, 0, 31, 32, 6, 3, -1, 0, 32, 33, 5, 3, 0, 0, 33, 53,
		3, 6, 3, 19, 34, 35, 5, 4, 0, 0, 35, 36, 3, 6, 3, 0, 36, 37, 5, 5, 0, 0,
		37, 53, 1, 0, 0, 0, 38, 47, 5, 20, 0, 0, 39, 44, 3, 6, 3, 0, 40, 41, 5,
		21, 0, 0, 41, 43, 3, 6, 3, 0, 42, 40, 1, 0, 0, 0, 43, 46, 1, 0, 0, 0, 44,
		42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 48, 1, 0, 0, 0, 46, 44, 1, 0, 0,
		0, 47, 39, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 53,
		5, 22, 0, 0, 50, 53, 5, 46, 0, 0, 51, 53, 3, 10, 5, 0, 52, 31, 1, 0, 0,
		0, 52, 34, 1, 0, 0, 0, 52, 38, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 52, 51,
		1, 0, 0, 0, 53, 98, 1, 0, 0, 0, 54, 55, 10, 17, 0, 0, 55, 56, 5, 6, 0,
		0, 56, 97, 3, 6, 3, 18, 57, 58, 10, 16, 0, 0, 58, 59, 5, 7, 0, 0, 59, 97,
		3, 6, 3, 17, 60, 61, 10, 15, 0, 0, 61, 62, 5, 8, 0, 0, 62, 97, 3, 6, 3,
		16, 63, 64, 10, 14, 0, 0, 64, 65, 5, 9, 0, 0, 65, 97, 3, 6, 3, 15, 66,
		67, 10, 13, 0, 0, 67, 68, 5, 10, 0, 0, 68, 97, 3, 6, 3, 14, 69, 70, 10,
		12, 0, 0, 70, 71, 5, 11, 0, 0, 71, 97, 3, 6, 3, 13, 72, 73, 10, 11, 0,
		0, 73, 74, 5, 12, 0, 0, 74, 97, 3, 6, 3, 12, 75, 76, 10, 10, 0, 0, 76,
		77, 5, 13, 0, 0, 77, 97, 3, 6, 3, 11, 78, 79, 10, 9, 0, 0, 79, 80, 5, 14,
		0, 0, 80, 97, 3, 6, 3, 10, 81, 82, 10, 8, 0, 0, 82, 83, 5, 15, 0, 0, 83,
		97, 3, 6, 3, 9, 84, 85, 10, 7, 0, 0, 85, 86, 5, 16, 0, 0, 86, 97, 3, 6,
		3, 8, 87, 88, 10, 6, 0, 0, 88, 89, 5, 17, 0, 0, 89, 97, 3, 6, 3, 7, 90,
		91, 10, 5, 0, 0, 91, 92, 5, 18, 0, 0, 92, 97, 3, 6, 3, 6, 93, 94, 10, 4,
		0, 0, 94, 95, 5, 19, 0, 0, 95, 97, 3, 6, 3, 5, 96, 54, 1, 0, 0, 0, 96,
		57, 1, 0, 0, 0, 96, 60, 1, 0, 0, 0, 96, 63, 1, 0, 0, 0, 96, 66, 1, 0, 0,
		0, 96, 69, 1, 0, 0, 0, 96, 72, 1, 0, 0, 0, 96, 75, 1, 0, 0, 0, 96, 78,
		1, 0, 0, 0, 96, 81, 1, 0, 0, 0, 96, 84, 1, 0, 0, 0, 96, 87, 1, 0, 0, 0,
		96, 90, 1, 0, 0, 0, 96, 93, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1,
		0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 7, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101,
		102, 5, 46, 0, 0, 102, 103, 5, 23, 0, 0, 103, 104, 3, 6, 3, 0, 104, 9,
		1, 0, 0, 0, 105, 107, 5, 11, 0, 0, 106, 105, 1, 0, 0, 0, 106, 107, 1, 0,
		0, 0, 107, 108, 1, 0, 0, 0, 108, 118, 5, 47, 0, 0, 109, 111, 5, 11, 0,
		0, 110, 109, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112,
		118, 5, 48, 0, 0, 113, 118, 5, 24, 0, 0, 114, 118, 5, 25, 0, 0, 115, 118,
		5, 49, 0, 0, 116, 118, 5, 50, 0, 0, 117, 106, 1, 0, 0, 0, 117, 110, 1,
		0, 0, 0, 117, 113, 1, 0, 0, 0, 117, 114, 1, 0, 0, 0, 117, 115, 1, 0, 0,
		0, 117, 116, 1, 0, 0, 0, 118, 11, 1, 0, 0, 0, 119, 120, 5, 26, 0, 0, 120,
		121, 5, 46, 0, 0, 121, 122, 5, 27, 0, 0, 122, 123, 3, 14, 7, 0, 123, 124,
		5, 23, 0, 0, 124, 125, 3, 6, 3, 0, 125, 126, 5, 2, 0, 0, 126, 164, 1, 0,
		0, 0, 127, 128, 5, 26, 0, 0, 128, 129, 5, 46, 0, 0, 129, 130, 5, 23, 0,
		0, 130, 131, 3, 6, 3, 0, 131, 132, 5, 2, 0, 0, 132, 164, 1, 0, 0, 0, 133,
		134, 5, 28, 0, 0, 134, 135, 5, 46, 0, 0, 135, 136, 5, 27, 0, 0, 136, 139,
		3, 14, 7, 0, 137, 138, 5, 23, 0, 0, 138, 140, 3, 6, 3, 0, 139, 137, 1,
		0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 142, 5, 2, 0,
		0, 142, 164, 1, 0, 0, 0, 143, 144, 5, 28, 0, 0, 144, 145, 5, 46, 0, 0,
		145, 146, 5, 23, 0, 0, 146, 147, 3, 6, 3, 0, 147, 148, 5, 2, 0, 0, 148,
		164, 1, 0, 0, 0, 149, 150, 5, 29, 0, 0, 150, 151, 5, 46, 0, 0, 151, 152,
		5, 27, 0, 0, 152, 153, 3, 14, 7, 0, 153, 154, 5, 23, 0, 0, 154, 155, 3,
		6, 3, 0, 155, 156, 5, 2, 0, 0, 156, 164, 1, 0, 0, 0, 157, 158, 5, 29, 0,
		0, 158, 159, 5, 46, 0, 0, 159, 160, 5, 23, 0, 0, 160, 161, 3, 6, 3, 0,
		161, 162, 5, 2, 0, 0, 162, 164, 1, 0, 0, 0, 163, 119, 1, 0, 0, 0, 163,
		127, 1, 0, 0, 0, 163, 133, 1, 0, 0, 0, 163, 143, 1, 0, 0, 0, 163, 149,
		1, 0, 0, 0, 163, 157, 1, 0, 0, 0, 164, 13, 1, 0, 0, 0, 165, 181, 5, 30,
		0, 0, 166, 181, 5, 31, 0, 0, 167, 181, 5, 32, 0, 0, 168, 181, 5, 33, 0,
		0, 169, 181, 5, 34, 0, 0, 170, 181, 5, 35, 0, 0, 171, 181, 5, 36, 0, 0,
		172, 181, 5, 37, 0, 0, 173, 181, 5, 38, 0, 0, 174, 181, 5, 39, 0, 0, 175,
		181, 5, 40, 0, 0, 176, 181, 5, 41, 0, 0, 177, 181, 5, 42, 0, 0, 178, 181,
		5, 43, 0, 0, 179, 181, 5, 44, 0, 0, 180, 165, 1, 0, 0, 0, 180, 166, 1,
		0, 0, 0, 180, 167, 1, 0, 0, 0, 180, 168, 1, 0, 0, 0, 180, 169, 1, 0, 0,
		0, 180, 170, 1, 0, 0, 0, 180, 171, 1, 0, 0, 0, 180, 172, 1, 0, 0, 0, 180,
		173, 1, 0, 0, 0, 180, 174, 1, 0, 0, 0, 180, 175, 1, 0, 0, 0, 180, 176,
		1, 0, 0, 0, 180, 177, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 179, 1, 0,
		0, 0, 181, 15, 1, 0, 0, 0, 13, 19, 25, 44, 47, 52, 96, 98, 106, 110, 117,
		139, 163, 180,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// haruParserInit initializes any static state used to implement haruParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewharuParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func HaruParserInit() {
	staticData := &HaruParserStaticData
	staticData.once.Do(haruParserInit)
}

// NewharuParser produces a new parser instance for the optional input antlr.TokenStream.
func NewharuParser(input antlr.TokenStream) *haruParser {
	HaruParserInit()
	this := new(haruParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &HaruParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "haru.g4"

	return this
}

// haruParser tokens.
const (
	haruParserEOF    = antlr.TokenEOF
	haruParserT__0   = 1
	haruParserT__1   = 2
	haruParserT__2   = 3
	haruParserT__3   = 4
	haruParserT__4   = 5
	haruParserT__5   = 6
	haruParserT__6   = 7
	haruParserT__7   = 8
	haruParserT__8   = 9
	haruParserT__9   = 10
	haruParserT__10  = 11
	haruParserT__11  = 12
	haruParserT__12  = 13
	haruParserT__13  = 14
	haruParserT__14  = 15
	haruParserT__15  = 16
	haruParserT__16  = 17
	haruParserT__17  = 18
	haruParserT__18  = 19
	haruParserT__19  = 20
	haruParserT__20  = 21
	haruParserT__21  = 22
	haruParserT__22  = 23
	haruParserT__23  = 24
	haruParserT__24  = 25
	haruParserT__25  = 26
	haruParserT__26  = 27
	haruParserT__27  = 28
	haruParserT__28  = 29
	haruParserT__29  = 30
	haruParserT__30  = 31
	haruParserT__31  = 32
	haruParserT__32  = 33
	haruParserT__33  = 34
	haruParserT__34  = 35
	haruParserT__35  = 36
	haruParserT__36  = 37
	haruParserT__37  = 38
	haruParserT__38  = 39
	haruParserT__39  = 40
	haruParserT__40  = 41
	haruParserT__41  = 42
	haruParserT__42  = 43
	haruParserT__43  = 44
	haruParserWS     = 45
	haruParserID     = 46
	haruParserNUMBER = 47
	haruParserFLOAT  = 48
	haruParserSTRING = 49
	haruParserBYTE   = 50
)

// haruParser rules.
const (
	haruParserRULE_program   = 0
	haruParserRULE_statement = 1
	haruParserRULE_printStmt = 2
	haruParserRULE_expr      = 3
	haruParserRULE_assign    = 4
	haruParserRULE_literal   = 5
	haruParserRULE_varDecl   = 6
	haruParserRULE_type      = 7
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, haruParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&70369616592898) != 0 {
		{
			p.SetState(16)
			p.Statement()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyAll(ctx *StatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssignStmtStatementContext struct {
	StatementContext
}

func NewAssignStmtStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignStmtStatementContext {
	var p = new(AssignStmtStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *AssignStmtStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStmtStatementContext) Assign() IAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *AssignStmtStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterAssignStmtStatement(s)
	}
}

func (s *AssignStmtStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitAssignStmtStatement(s)
	}
}

func (s *AssignStmtStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitAssignStmtStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintStmtStatementContext struct {
	StatementContext
}

func NewPrintStmtStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintStmtStatementContext {
	var p = new(PrintStmtStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *PrintStmtStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStmtStatementContext) PrintStmt() IPrintStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintStmtContext)
}

func (s *PrintStmtStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterPrintStmtStatement(s)
	}
}

func (s *PrintStmtStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitPrintStmtStatement(s)
	}
}

func (s *PrintStmtStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitPrintStmtStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarDeclStatementContext struct {
	StatementContext
}

func NewVarDeclStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarDeclStatementContext {
	var p = new(VarDeclStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *VarDeclStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclStatementContext) VarDecl() IVarDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *VarDeclStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterVarDeclStatement(s)
	}
}

func (s *VarDeclStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitVarDeclStatement(s)
	}
}

func (s *VarDeclStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitVarDeclStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, haruParserRULE_statement)
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case haruParserT__0:
		localctx = NewPrintStmtStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.PrintStmt()
		}

	case haruParserT__25, haruParserT__27, haruParserT__28:
		localctx = NewVarDeclStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)
			p.VarDecl()
		}

	case haruParserID:
		localctx = NewAssignStmtStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(24)
			p.Assign()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintStmtContext is an interface to support dynamic dispatch.
type IPrintStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrintStmtContext differentiates from other interfaces.
	IsPrintStmtContext()
}

type PrintStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintStmtContext() *PrintStmtContext {
	var p = new(PrintStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_printStmt
	return p
}

func InitEmptyPrintStmtContext(p *PrintStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_printStmt
}

func (*PrintStmtContext) IsPrintStmtContext() {}

func NewPrintStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintStmtContext {
	var p = new(PrintStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_printStmt

	return p
}

func (s *PrintStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintStmtContext) CopyAll(ctx *PrintStmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrintStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrintStatementContext struct {
	PrintStmtContext
}

func NewPrintStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintStatementContext {
	var p = new(PrintStatementContext)

	InitEmptyPrintStmtContext(&p.PrintStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrintStmtContext))

	return p
}

func (s *PrintStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStatementContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterPrintStatement(s)
	}
}

func (s *PrintStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitPrintStatement(s)
	}
}

func (s *PrintStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitPrintStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) PrintStmt() (localctx IPrintStmtContext) {
	localctx = NewPrintStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, haruParserRULE_printStmt)
	localctx = NewPrintStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(27)
		p.Match(haruParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(28)
		p.expr(0)
	}
	{
		p.SetState(29)
		p.Match(haruParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MulExprContext struct {
	ExprContext
}

func NewMulExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulExprContext {
	var p = new(MulExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MulExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MulExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MulExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterMulExpr(s)
	}
}

func (s *MulExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitMulExpr(s)
	}
}

func (s *MulExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitMulExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndExprContext struct {
	ExprContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AndExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LitExprContext struct {
	ExprContext
}

func NewLitExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LitExprContext {
	var p = new(LitExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LitExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitExprContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LitExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterLitExpr(s)
	}
}

func (s *LitExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitLitExpr(s)
	}
}

func (s *LitExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitLitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NeExprContext struct {
	ExprContext
}

func NewNeExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NeExprContext {
	var p = new(NeExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NeExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *NeExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterNeExpr(s)
	}
}

func (s *NeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitNeExpr(s)
	}
}

func (s *NeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitNeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SubExprContext struct {
	ExprContext
}

func NewSubExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubExprContext {
	var p = new(SubExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *SubExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *SubExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SubExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterSubExpr(s)
	}
}

func (s *SubExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitSubExpr(s)
	}
}

func (s *SubExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitSubExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LtExprContext struct {
	ExprContext
}

func NewLtExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LtExprContext {
	var p = new(LtExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LtExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LtExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *LtExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LtExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterLtExpr(s)
	}
}

func (s *LtExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitLtExpr(s)
	}
}

func (s *LtExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitLtExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type GtExprContext struct {
	ExprContext
}

func NewGtExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GtExprContext {
	var p = new(GtExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *GtExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GtExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *GtExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GtExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterGtExpr(s)
	}
}

func (s *GtExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitGtExpr(s)
	}
}

func (s *GtExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitGtExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddExprContext struct {
	ExprContext
}

func NewAddExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExprContext {
	var p = new(AddExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AddExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterAddExpr(s)
	}
}

func (s *AddExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitAddExpr(s)
	}
}

func (s *AddExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitAddExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type GeExprContext struct {
	ExprContext
}

func NewGeExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GeExprContext {
	var p = new(GeExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *GeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *GeExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterGeExpr(s)
	}
}

func (s *GeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitGeExpr(s)
	}
}

func (s *GeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitGeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpExprContext struct {
	ExprContext
}

func NewExpExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpExprContext {
	var p = new(ExpExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExpExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExpExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExpExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterExpExpr(s)
	}
}

func (s *ExpExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitExpExpr(s)
	}
}

func (s *ExpExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitExpExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LeExprContext struct {
	ExprContext
}

func NewLeExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LeExprContext {
	var p = new(LeExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *LeExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterLeExpr(s)
	}
}

func (s *LeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitLeExpr(s)
	}
}

func (s *LeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitLeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExprContext struct {
	ExprContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *OrExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayExprContext struct {
	ExprContext
}

func NewArrayExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayExprContext {
	var p = new(ArrayExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ArrayExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArrayExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrayExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterArrayExpr(s)
	}
}

func (s *ArrayExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitArrayExpr(s)
	}
}

func (s *ArrayExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitArrayExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type DivExprContext struct {
	ExprContext
}

func NewDivExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivExprContext {
	var p = new(DivExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *DivExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *DivExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DivExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterDivExpr(s)
	}
}

func (s *DivExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitDivExpr(s)
	}
}

func (s *DivExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitDivExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqExprContext struct {
	ExprContext
}

func NewEqExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqExprContext {
	var p = new(EqExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EqExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EqExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EqExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterEqExpr(s)
	}
}

func (s *EqExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitEqExpr(s)
	}
}

func (s *EqExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitEqExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarExprContext struct {
	ExprContext
}

func NewVarExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarExprContext {
	var p = new(VarExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *VarExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarExprContext) ID() antlr.TerminalNode {
	return s.GetToken(haruParserID, 0)
}

func (s *VarExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterVarExpr(s)
	}
}

func (s *VarExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitVarExpr(s)
	}
}

func (s *VarExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitVarExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExprContext struct {
	ExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ModExprContext struct {
	ExprContext
}

func NewModExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModExprContext {
	var p = new(ModExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ModExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ModExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ModExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterModExpr(s)
	}
}

func (s *ModExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitModExpr(s)
	}
}

func (s *ModExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitModExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExprContext struct {
	ExprContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitParenExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *haruParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, haruParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case haruParserT__2:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(32)
			p.Match(haruParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(33)
			p.expr(19)
		}

	case haruParserT__3:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(34)
			p.Match(haruParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(35)
			p.expr(0)
		}
		{
			p.SetState(36)
			p.Match(haruParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__19:
		localctx = NewArrayExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(38)
			p.Match(haruParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2181431120889880) != 0 {
			{
				p.SetState(39)
				p.expr(0)
			}
			p.SetState(44)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == haruParserT__20 {
				{
					p.SetState(40)
					p.Match(haruParserT__20)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(41)
					p.expr(0)
				}

				p.SetState(46)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(49)
			p.Match(haruParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserID:
		localctx = NewVarExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(50)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__10, haruParserT__23, haruParserT__24, haruParserNUMBER, haruParserFLOAT, haruParserSTRING, haruParserBYTE:
		localctx = NewLitExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(51)
			p.Literal()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(54)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(55)
					p.Match(haruParserT__5)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(56)
					p.expr(18)
				}

			case 2:
				localctx = NewMulExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(57)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(58)
					p.Match(haruParserT__6)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(59)
					p.expr(17)
				}

			case 3:
				localctx = NewDivExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(60)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(61)
					p.Match(haruParserT__7)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(62)
					p.expr(16)
				}

			case 4:
				localctx = NewModExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(63)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(64)
					p.Match(haruParserT__8)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(65)
					p.expr(15)
				}

			case 5:
				localctx = NewAddExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(66)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(67)
					p.Match(haruParserT__9)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(68)
					p.expr(14)
				}

			case 6:
				localctx = NewSubExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(69)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(70)
					p.Match(haruParserT__10)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(71)
					p.expr(13)
				}

			case 7:
				localctx = NewLtExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(72)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(73)
					p.Match(haruParserT__11)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(74)
					p.expr(12)
				}

			case 8:
				localctx = NewGtExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(75)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(76)
					p.Match(haruParserT__12)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(77)
					p.expr(11)
				}

			case 9:
				localctx = NewLeExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(78)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(79)
					p.Match(haruParserT__13)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(80)
					p.expr(10)
				}

			case 10:
				localctx = NewGeExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(81)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(82)
					p.Match(haruParserT__14)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(83)
					p.expr(9)
				}

			case 11:
				localctx = NewEqExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(84)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(85)
					p.Match(haruParserT__15)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(86)
					p.expr(8)
				}

			case 12:
				localctx = NewNeExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(87)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(88)
					p.Match(haruParserT__16)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(89)
					p.expr(7)
				}

			case 13:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(91)
					p.Match(haruParserT__17)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(92)
					p.expr(6)
				}

			case 14:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(93)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(94)
					p.Match(haruParserT__18)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(95)
					p.expr(5)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_assign
	return p
}

func InitEmptyAssignContext(p *AssignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_assign
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) CopyAll(ctx *AssignContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssignStmtContext struct {
	AssignContext
}

func NewAssignStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignStmtContext {
	var p = new(AssignStmtContext)

	InitEmptyAssignContext(&p.AssignContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignContext))

	return p
}

func (s *AssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStmtContext) ID() antlr.TerminalNode {
	return s.GetToken(haruParserID, 0)
}

func (s *AssignStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterAssignStmt(s)
	}
}

func (s *AssignStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitAssignStmt(s)
	}
}

func (s *AssignStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitAssignStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, haruParserRULE_assign)
	localctx = NewAssignStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Match(haruParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Match(haruParserT__22)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyAll(ctx *LiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TrueLiteralContext struct {
	LiteralContext
}

func NewTrueLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueLiteralContext {
	var p = new(TrueLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *TrueLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterTrueLiteral(s)
	}
}

func (s *TrueLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitTrueLiteral(s)
	}
}

func (s *TrueLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitTrueLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringLiteralContext struct {
	LiteralContext
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext {
	var p = new(StringLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(haruParserSTRING, 0)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatLiteralContext struct {
	LiteralContext
}

func NewFloatLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(haruParserFLOAT, 0)
}

func (s *FloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitFloatLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type ByteLiteralContext struct {
	LiteralContext
}

func NewByteLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ByteLiteralContext {
	var p = new(ByteLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *ByteLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ByteLiteralContext) BYTE() antlr.TerminalNode {
	return s.GetToken(haruParserBYTE, 0)
}

func (s *ByteLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterByteLiteral(s)
	}
}

func (s *ByteLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitByteLiteral(s)
	}
}

func (s *ByteLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitByteLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntLiteralContext struct {
	LiteralContext
}

func NewIntLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntLiteralContext {
	var p = new(IntLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *IntLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLiteralContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(haruParserNUMBER, 0)
}

func (s *IntLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterIntLiteral(s)
	}
}

func (s *IntLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitIntLiteral(s)
	}
}

func (s *IntLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitIntLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type FalseLiteralContext struct {
	LiteralContext
}

func NewFalseLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalseLiteralContext {
	var p = new(FalseLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *FalseLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterFalseLiteral(s)
	}
}

func (s *FalseLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitFalseLiteral(s)
	}
}

func (s *FalseLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitFalseLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, haruParserRULE_literal)
	var _la int

	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIntLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == haruParserT__10 {
			{
				p.SetState(105)
				p.Match(haruParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(108)
			p.Match(haruParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewFloatLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == haruParserT__10 {
			{
				p.SetState(109)
				p.Match(haruParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(112)
			p.Match(haruParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewTrueLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(113)
			p.Match(haruParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewFalseLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(114)
			p.Match(haruParserT__24)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(115)
			p.Match(haruParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewByteLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(116)
			p.Match(haruParserBYTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclContext() *VarDeclContext {
	var p = new(VarDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_varDecl
	return p
}

func InitEmptyVarDeclContext(p *VarDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_varDecl
}

func (*VarDeclContext) IsVarDeclContext() {}

func NewVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclContext {
	var p = new(VarDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_varDecl

	return p
}

func (s *VarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclContext) CopyAll(ctx *VarDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MutInferDeclContext struct {
	VarDeclContext
}

func NewMutInferDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MutInferDeclContext {
	var p = new(MutInferDeclContext)

	InitEmptyVarDeclContext(&p.VarDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDeclContext))

	return p
}

func (s *MutInferDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutInferDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(haruParserID, 0)
}

func (s *MutInferDeclContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MutInferDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterMutInferDecl(s)
	}
}

func (s *MutInferDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitMutInferDecl(s)
	}
}

func (s *MutInferDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitMutInferDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConstInferDeclContext struct {
	VarDeclContext
}

func NewConstInferDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstInferDeclContext {
	var p = new(ConstInferDeclContext)

	InitEmptyVarDeclContext(&p.VarDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDeclContext))

	return p
}

func (s *ConstInferDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstInferDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(haruParserID, 0)
}

func (s *ConstInferDeclContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConstInferDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterConstInferDecl(s)
	}
}

func (s *ConstInferDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitConstInferDecl(s)
	}
}

func (s *ConstInferDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitConstInferDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type LetInferDeclContext struct {
	VarDeclContext
}

func NewLetInferDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetInferDeclContext {
	var p = new(LetInferDeclContext)

	InitEmptyVarDeclContext(&p.VarDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDeclContext))

	return p
}

func (s *LetInferDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetInferDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(haruParserID, 0)
}

func (s *LetInferDeclContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LetInferDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterLetInferDecl(s)
	}
}

func (s *LetInferDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitLetInferDecl(s)
	}
}

func (s *LetInferDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitLetInferDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type LetDeclContext struct {
	VarDeclContext
}

func NewLetDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetDeclContext {
	var p = new(LetDeclContext)

	InitEmptyVarDeclContext(&p.VarDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDeclContext))

	return p
}

func (s *LetDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(haruParserID, 0)
}

func (s *LetDeclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *LetDeclContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LetDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterLetDecl(s)
	}
}

func (s *LetDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitLetDecl(s)
	}
}

func (s *LetDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitLetDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConstDeclContext struct {
	VarDeclContext
}

func NewConstDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstDeclContext {
	var p = new(ConstDeclContext)

	InitEmptyVarDeclContext(&p.VarDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDeclContext))

	return p
}

func (s *ConstDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(haruParserID, 0)
}

func (s *ConstDeclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ConstDeclContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConstDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterConstDecl(s)
	}
}

func (s *ConstDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitConstDecl(s)
	}
}

func (s *ConstDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitConstDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type MutDeclContext struct {
	VarDeclContext
}

func NewMutDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MutDeclContext {
	var p = new(MutDeclContext)

	InitEmptyVarDeclContext(&p.VarDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDeclContext))

	return p
}

func (s *MutDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(haruParserID, 0)
}

func (s *MutDeclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *MutDeclContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MutDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterMutDecl(s)
	}
}

func (s *MutDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitMutDecl(s)
	}
}

func (s *MutDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitMutDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) VarDecl() (localctx IVarDeclContext) {
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, haruParserRULE_varDecl)
	var _la int

	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLetDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Match(haruParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.Match(haruParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.Type_()
		}
		{
			p.SetState(123)
			p.Match(haruParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.expr(0)
		}
		{
			p.SetState(125)
			p.Match(haruParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewLetInferDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(127)
			p.Match(haruParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)
			p.Match(haruParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.expr(0)
		}
		{
			p.SetState(131)
			p.Match(haruParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewMutDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(133)
			p.Match(haruParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.Match(haruParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.Type_()
		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == haruParserT__22 {
			{
				p.SetState(137)
				p.Match(haruParserT__22)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(138)
				p.expr(0)
			}

		}
		{
			p.SetState(141)
			p.Match(haruParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewMutInferDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(143)
			p.Match(haruParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)
			p.Match(haruParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.expr(0)
		}
		{
			p.SetState(147)
			p.Match(haruParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewConstDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(149)
			p.Match(haruParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.Match(haruParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)
			p.Type_()
		}
		{
			p.SetState(153)
			p.Match(haruParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.expr(0)
		}
		{
			p.SetState(155)
			p.Match(haruParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewConstInferDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(157)
			p.Match(haruParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)
			p.Match(haruParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.expr(0)
		}
		{
			p.SetState(161)
			p.Match(haruParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) CopyAll(ctx *TypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type F64TypeContext struct {
	TypeContext
}

func NewF64TypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *F64TypeContext {
	var p = new(F64TypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *F64TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *F64TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterF64Type(s)
	}
}

func (s *F64TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitF64Type(s)
	}
}

func (s *F64TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitF64Type(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolTypeContext struct {
	TypeContext
}

func NewBoolTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolTypeContext {
	var p = new(BoolTypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *BoolTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterBoolType(s)
	}
}

func (s *BoolTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitBoolType(s)
	}
}

func (s *BoolTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitBoolType(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringTypeContext struct {
	TypeContext
}

func NewStringTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringTypeContext {
	var p = new(StringTypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *StringTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterStringType(s)
	}
}

func (s *StringTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitStringType(s)
	}
}

func (s *StringTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitStringType(s)

	default:
		return t.VisitChildren(s)
	}
}

type I8TypeContext struct {
	TypeContext
}

func NewI8TypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *I8TypeContext {
	var p = new(I8TypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *I8TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *I8TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterI8Type(s)
	}
}

func (s *I8TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitI8Type(s)
	}
}

func (s *I8TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitI8Type(s)

	default:
		return t.VisitChildren(s)
	}
}

type ByteTypeContext struct {
	TypeContext
}

func NewByteTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ByteTypeContext {
	var p = new(ByteTypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *ByteTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ByteTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterByteType(s)
	}
}

func (s *ByteTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitByteType(s)
	}
}

func (s *ByteTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitByteType(s)

	default:
		return t.VisitChildren(s)
	}
}

type I16TypeContext struct {
	TypeContext
}

func NewI16TypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *I16TypeContext {
	var p = new(I16TypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *I16TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *I16TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterI16Type(s)
	}
}

func (s *I16TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitI16Type(s)
	}
}

func (s *I16TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitI16Type(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntTypeContext struct {
	TypeContext
}

func NewIntTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntTypeContext {
	var p = new(IntTypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *IntTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterIntType(s)
	}
}

func (s *IntTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitIntType(s)
	}
}

func (s *IntTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitIntType(s)

	default:
		return t.VisitChildren(s)
	}
}

type I64TypeContext struct {
	TypeContext
}

func NewI64TypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *I64TypeContext {
	var p = new(I64TypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *I64TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *I64TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterI64Type(s)
	}
}

func (s *I64TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitI64Type(s)
	}
}

func (s *I64TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitI64Type(s)

	default:
		return t.VisitChildren(s)
	}
}

type I32TypeContext struct {
	TypeContext
}

func NewI32TypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *I32TypeContext {
	var p = new(I32TypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *I32TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *I32TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterI32Type(s)
	}
}

func (s *I32TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitI32Type(s)
	}
}

func (s *I32TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitI32Type(s)

	default:
		return t.VisitChildren(s)
	}
}

type UI32TypeContext struct {
	TypeContext
}

func NewUI32TypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UI32TypeContext {
	var p = new(UI32TypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *UI32TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UI32TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterUI32Type(s)
	}
}

func (s *UI32TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitUI32Type(s)
	}
}

func (s *UI32TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitUI32Type(s)

	default:
		return t.VisitChildren(s)
	}
}

type UI64TypeContext struct {
	TypeContext
}

func NewUI64TypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UI64TypeContext {
	var p = new(UI64TypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *UI64TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UI64TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterUI64Type(s)
	}
}

func (s *UI64TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitUI64Type(s)
	}
}

func (s *UI64TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitUI64Type(s)

	default:
		return t.VisitChildren(s)
	}
}

type UI16TypeContext struct {
	TypeContext
}

func NewUI16TypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UI16TypeContext {
	var p = new(UI16TypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *UI16TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UI16TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterUI16Type(s)
	}
}

func (s *UI16TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitUI16Type(s)
	}
}

func (s *UI16TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitUI16Type(s)

	default:
		return t.VisitChildren(s)
	}
}

type F32TypeContext struct {
	TypeContext
}

func NewF32TypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *F32TypeContext {
	var p = new(F32TypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *F32TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *F32TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterF32Type(s)
	}
}

func (s *F32TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitF32Type(s)
	}
}

func (s *F32TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitF32Type(s)

	default:
		return t.VisitChildren(s)
	}
}

type UI8TypeContext struct {
	TypeContext
}

func NewUI8TypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UI8TypeContext {
	var p = new(UI8TypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *UI8TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UI8TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterUI8Type(s)
	}
}

func (s *UI8TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitUI8Type(s)
	}
}

func (s *UI8TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitUI8Type(s)

	default:
		return t.VisitChildren(s)
	}
}

type UIntTypeContext struct {
	TypeContext
}

func NewUIntTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UIntTypeContext {
	var p = new(UIntTypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *UIntTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UIntTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterUIntType(s)
	}
}

func (s *UIntTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitUIntType(s)
	}
}

func (s *UIntTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitUIntType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, haruParserRULE_type)
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case haruParserT__29:
		localctx = NewI8TypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(165)
			p.Match(haruParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__30:
		localctx = NewI16TypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(166)
			p.Match(haruParserT__30)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__31:
		localctx = NewI32TypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(167)
			p.Match(haruParserT__31)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__32:
		localctx = NewI64TypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(168)
			p.Match(haruParserT__32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__33:
		localctx = NewIntTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(169)
			p.Match(haruParserT__33)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__34:
		localctx = NewUI8TypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(170)
			p.Match(haruParserT__34)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__35:
		localctx = NewUI16TypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(171)
			p.Match(haruParserT__35)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__36:
		localctx = NewUI32TypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(172)
			p.Match(haruParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__37:
		localctx = NewUI64TypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(173)
			p.Match(haruParserT__37)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__38:
		localctx = NewUIntTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(174)
			p.Match(haruParserT__38)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__39:
		localctx = NewF32TypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(175)
			p.Match(haruParserT__39)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__40:
		localctx = NewF64TypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(176)
			p.Match(haruParserT__40)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__41:
		localctx = NewBoolTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(177)
			p.Match(haruParserT__41)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__42:
		localctx = NewStringTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(178)
			p.Match(haruParserT__42)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__43:
		localctx = NewByteTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(179)
			p.Match(haruParserT__43)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *haruParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *haruParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
