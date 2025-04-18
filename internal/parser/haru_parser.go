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
		"'string'", "'byte'", "'if'", "'else'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "WS", "ID",
		"NUMBER", "FLOAT", "STRING", "BYTE",
	}
	staticData.RuleNames = []string{
		"program", "statement", "printStmt", "expr", "assign", "literal", "varDecl",
		"type", "ifStmt", "elseIfBlock", "elseBlock", "block",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 54, 234, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 5, 0, 26, 8, 0, 10, 0, 12, 0, 29, 9, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 3, 1, 35, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 52, 8, 3, 10, 3,
		12, 3, 55, 9, 3, 3, 3, 57, 8, 3, 1, 3, 1, 3, 1, 3, 3, 3, 62, 8, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 106, 8, 3, 10, 3, 12, 3, 109, 9, 3,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 3, 5, 117, 8, 5, 1, 5, 1, 5, 3, 5,
		121, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 128, 8, 5, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 150, 8, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 174, 8, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 3, 7, 191, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 5, 8, 205, 8, 8, 10, 8, 12, 8, 208, 9, 8, 1, 8, 3,
		8, 211, 8, 8, 3, 8, 213, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 5, 11, 227, 8, 11, 10, 11, 12, 11, 230,
		9, 11, 1, 11, 1, 11, 1, 11, 0, 1, 6, 12, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 0, 0, 276, 0, 27, 1, 0, 0, 0, 2, 34, 1, 0, 0, 0, 4, 36, 1,
		0, 0, 0, 6, 61, 1, 0, 0, 0, 8, 110, 1, 0, 0, 0, 10, 127, 1, 0, 0, 0, 12,
		173, 1, 0, 0, 0, 14, 190, 1, 0, 0, 0, 16, 212, 1, 0, 0, 0, 18, 214, 1,
		0, 0, 0, 20, 221, 1, 0, 0, 0, 22, 224, 1, 0, 0, 0, 24, 26, 3, 2, 1, 0,
		25, 24, 1, 0, 0, 0, 26, 29, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 28, 1,
		0, 0, 0, 28, 1, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 30, 35, 3, 4, 2, 0, 31,
		35, 3, 12, 6, 0, 32, 35, 3, 8, 4, 0, 33, 35, 3, 16, 8, 0, 34, 30, 1, 0,
		0, 0, 34, 31, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 34, 33, 1, 0, 0, 0, 35, 3,
		1, 0, 0, 0, 36, 37, 5, 1, 0, 0, 37, 38, 3, 6, 3, 0, 38, 39, 5, 2, 0, 0,
		39, 5, 1, 0, 0, 0, 40, 41, 6, 3, -1, 0, 41, 42, 5, 3, 0, 0, 42, 62, 3,
		6, 3, 19, 43, 44, 5, 4, 0, 0, 44, 45, 3, 6, 3, 0, 45, 46, 5, 5, 0, 0, 46,
		62, 1, 0, 0, 0, 47, 56, 5, 20, 0, 0, 48, 53, 3, 6, 3, 0, 49, 50, 5, 21,
		0, 0, 50, 52, 3, 6, 3, 0, 51, 49, 1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 51,
		1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 57, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0,
		56, 48, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 62, 5,
		22, 0, 0, 59, 62, 5, 50, 0, 0, 60, 62, 3, 10, 5, 0, 61, 40, 1, 0, 0, 0,
		61, 43, 1, 0, 0, 0, 61, 47, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 60, 1,
		0, 0, 0, 62, 107, 1, 0, 0, 0, 63, 64, 10, 17, 0, 0, 64, 65, 5, 6, 0, 0,
		65, 106, 3, 6, 3, 18, 66, 67, 10, 16, 0, 0, 67, 68, 5, 7, 0, 0, 68, 106,
		3, 6, 3, 17, 69, 70, 10, 15, 0, 0, 70, 71, 5, 8, 0, 0, 71, 106, 3, 6, 3,
		16, 72, 73, 10, 14, 0, 0, 73, 74, 5, 9, 0, 0, 74, 106, 3, 6, 3, 15, 75,
		76, 10, 13, 0, 0, 76, 77, 5, 10, 0, 0, 77, 106, 3, 6, 3, 14, 78, 79, 10,
		12, 0, 0, 79, 80, 5, 11, 0, 0, 80, 106, 3, 6, 3, 13, 81, 82, 10, 11, 0,
		0, 82, 83, 5, 12, 0, 0, 83, 106, 3, 6, 3, 12, 84, 85, 10, 10, 0, 0, 85,
		86, 5, 13, 0, 0, 86, 106, 3, 6, 3, 11, 87, 88, 10, 9, 0, 0, 88, 89, 5,
		14, 0, 0, 89, 106, 3, 6, 3, 10, 90, 91, 10, 8, 0, 0, 91, 92, 5, 15, 0,
		0, 92, 106, 3, 6, 3, 9, 93, 94, 10, 7, 0, 0, 94, 95, 5, 16, 0, 0, 95, 106,
		3, 6, 3, 8, 96, 97, 10, 6, 0, 0, 97, 98, 5, 17, 0, 0, 98, 106, 3, 6, 3,
		7, 99, 100, 10, 5, 0, 0, 100, 101, 5, 18, 0, 0, 101, 106, 3, 6, 3, 6, 102,
		103, 10, 4, 0, 0, 103, 104, 5, 19, 0, 0, 104, 106, 3, 6, 3, 5, 105, 63,
		1, 0, 0, 0, 105, 66, 1, 0, 0, 0, 105, 69, 1, 0, 0, 0, 105, 72, 1, 0, 0,
		0, 105, 75, 1, 0, 0, 0, 105, 78, 1, 0, 0, 0, 105, 81, 1, 0, 0, 0, 105,
		84, 1, 0, 0, 0, 105, 87, 1, 0, 0, 0, 105, 90, 1, 0, 0, 0, 105, 93, 1, 0,
		0, 0, 105, 96, 1, 0, 0, 0, 105, 99, 1, 0, 0, 0, 105, 102, 1, 0, 0, 0, 106,
		109, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 7, 1,
		0, 0, 0, 109, 107, 1, 0, 0, 0, 110, 111, 5, 50, 0, 0, 111, 112, 5, 23,
		0, 0, 112, 113, 3, 6, 3, 0, 113, 114, 5, 2, 0, 0, 114, 9, 1, 0, 0, 0, 115,
		117, 5, 11, 0, 0, 116, 115, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 118,
		1, 0, 0, 0, 118, 128, 5, 51, 0, 0, 119, 121, 5, 11, 0, 0, 120, 119, 1,
		0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 128, 5, 52, 0,
		0, 123, 128, 5, 24, 0, 0, 124, 128, 5, 25, 0, 0, 125, 128, 5, 53, 0, 0,
		126, 128, 5, 54, 0, 0, 127, 116, 1, 0, 0, 0, 127, 120, 1, 0, 0, 0, 127,
		123, 1, 0, 0, 0, 127, 124, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127, 126,
		1, 0, 0, 0, 128, 11, 1, 0, 0, 0, 129, 130, 5, 26, 0, 0, 130, 131, 5, 50,
		0, 0, 131, 132, 5, 27, 0, 0, 132, 133, 3, 14, 7, 0, 133, 134, 5, 23, 0,
		0, 134, 135, 3, 6, 3, 0, 135, 136, 5, 2, 0, 0, 136, 174, 1, 0, 0, 0, 137,
		138, 5, 26, 0, 0, 138, 139, 5, 50, 0, 0, 139, 140, 5, 23, 0, 0, 140, 141,
		3, 6, 3, 0, 141, 142, 5, 2, 0, 0, 142, 174, 1, 0, 0, 0, 143, 144, 5, 28,
		0, 0, 144, 145, 5, 50, 0, 0, 145, 146, 5, 27, 0, 0, 146, 149, 3, 14, 7,
		0, 147, 148, 5, 23, 0, 0, 148, 150, 3, 6, 3, 0, 149, 147, 1, 0, 0, 0, 149,
		150, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 152, 5, 2, 0, 0, 152, 174,
		1, 0, 0, 0, 153, 154, 5, 28, 0, 0, 154, 155, 5, 50, 0, 0, 155, 156, 5,
		23, 0, 0, 156, 157, 3, 6, 3, 0, 157, 158, 5, 2, 0, 0, 158, 174, 1, 0, 0,
		0, 159, 160, 5, 29, 0, 0, 160, 161, 5, 50, 0, 0, 161, 162, 5, 27, 0, 0,
		162, 163, 3, 14, 7, 0, 163, 164, 5, 23, 0, 0, 164, 165, 3, 6, 3, 0, 165,
		166, 5, 2, 0, 0, 166, 174, 1, 0, 0, 0, 167, 168, 5, 29, 0, 0, 168, 169,
		5, 50, 0, 0, 169, 170, 5, 23, 0, 0, 170, 171, 3, 6, 3, 0, 171, 172, 5,
		2, 0, 0, 172, 174, 1, 0, 0, 0, 173, 129, 1, 0, 0, 0, 173, 137, 1, 0, 0,
		0, 173, 143, 1, 0, 0, 0, 173, 153, 1, 0, 0, 0, 173, 159, 1, 0, 0, 0, 173,
		167, 1, 0, 0, 0, 174, 13, 1, 0, 0, 0, 175, 191, 5, 30, 0, 0, 176, 191,
		5, 31, 0, 0, 177, 191, 5, 32, 0, 0, 178, 191, 5, 33, 0, 0, 179, 191, 5,
		34, 0, 0, 180, 191, 5, 35, 0, 0, 181, 191, 5, 36, 0, 0, 182, 191, 5, 37,
		0, 0, 183, 191, 5, 38, 0, 0, 184, 191, 5, 39, 0, 0, 185, 191, 5, 40, 0,
		0, 186, 191, 5, 41, 0, 0, 187, 191, 5, 42, 0, 0, 188, 191, 5, 43, 0, 0,
		189, 191, 5, 44, 0, 0, 190, 175, 1, 0, 0, 0, 190, 176, 1, 0, 0, 0, 190,
		177, 1, 0, 0, 0, 190, 178, 1, 0, 0, 0, 190, 179, 1, 0, 0, 0, 190, 180,
		1, 0, 0, 0, 190, 181, 1, 0, 0, 0, 190, 182, 1, 0, 0, 0, 190, 183, 1, 0,
		0, 0, 190, 184, 1, 0, 0, 0, 190, 185, 1, 0, 0, 0, 190, 186, 1, 0, 0, 0,
		190, 187, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 190, 189, 1, 0, 0, 0, 191,
		15, 1, 0, 0, 0, 192, 193, 5, 45, 0, 0, 193, 194, 5, 4, 0, 0, 194, 195,
		3, 6, 3, 0, 195, 196, 5, 5, 0, 0, 196, 197, 3, 22, 11, 0, 197, 213, 1,
		0, 0, 0, 198, 199, 5, 45, 0, 0, 199, 200, 5, 4, 0, 0, 200, 201, 3, 6, 3,
		0, 201, 202, 5, 5, 0, 0, 202, 206, 3, 22, 11, 0, 203, 205, 3, 18, 9, 0,
		204, 203, 1, 0, 0, 0, 205, 208, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 206,
		207, 1, 0, 0, 0, 207, 210, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 209, 211,
		3, 20, 10, 0, 210, 209, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 213, 1,
		0, 0, 0, 212, 192, 1, 0, 0, 0, 212, 198, 1, 0, 0, 0, 213, 17, 1, 0, 0,
		0, 214, 215, 5, 46, 0, 0, 215, 216, 5, 45, 0, 0, 216, 217, 5, 4, 0, 0,
		217, 218, 3, 6, 3, 0, 218, 219, 5, 5, 0, 0, 219, 220, 3, 22, 11, 0, 220,
		19, 1, 0, 0, 0, 221, 222, 5, 46, 0, 0, 222, 223, 3, 22, 11, 0, 223, 21,
		1, 0, 0, 0, 224, 228, 5, 47, 0, 0, 225, 227, 3, 2, 1, 0, 226, 225, 1, 0,
		0, 0, 227, 230, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0,
		229, 231, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 231, 232, 5, 48, 0, 0, 232,
		23, 1, 0, 0, 0, 17, 27, 34, 53, 56, 61, 105, 107, 116, 120, 127, 149, 173,
		190, 206, 210, 212, 228,
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
	haruParserT__44  = 45
	haruParserT__45  = 46
	haruParserT__46  = 47
	haruParserT__47  = 48
	haruParserWS     = 49
	haruParserID     = 50
	haruParserNUMBER = 51
	haruParserFLOAT  = 52
	haruParserSTRING = 53
	haruParserBYTE   = 54
)

// haruParser rules.
const (
	haruParserRULE_program     = 0
	haruParserRULE_statement   = 1
	haruParserRULE_printStmt   = 2
	haruParserRULE_expr        = 3
	haruParserRULE_assign      = 4
	haruParserRULE_literal     = 5
	haruParserRULE_varDecl     = 6
	haruParserRULE_type        = 7
	haruParserRULE_ifStmt      = 8
	haruParserRULE_elseIfBlock = 9
	haruParserRULE_elseBlock   = 10
	haruParserRULE_block       = 11
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
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1161085151346690) != 0 {
		{
			p.SetState(24)
			p.Statement()
		}

		p.SetState(29)
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

type IfStmtStatementContext struct {
	StatementContext
}

func NewIfStmtStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStmtStatementContext {
	var p = new(IfStmtStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *IfStmtStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtStatementContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *IfStmtStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterIfStmtStatement(s)
	}
}

func (s *IfStmtStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitIfStmtStatement(s)
	}
}

func (s *IfStmtStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitIfStmtStatement(s)

	default:
		return t.VisitChildren(s)
	}
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
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case haruParserT__0:
		localctx = NewPrintStmtStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)
			p.PrintStmt()
		}

	case haruParserT__25, haruParserT__27, haruParserT__28:
		localctx = NewVarDeclStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.VarDecl()
		}

	case haruParserID:
		localctx = NewAssignStmtStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(32)
			p.Assign()
		}

	case haruParserT__44:
		localctx = NewIfStmtStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(33)
			p.IfStmt()
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
		p.SetState(36)
		p.Match(haruParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(37)
		p.expr(0)
	}
	{
		p.SetState(38)
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
	p.SetState(61)
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
			p.SetState(41)
			p.Match(haruParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(42)
			p.expr(19)
		}

	case haruParserT__3:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(43)
			p.Match(haruParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(44)
			p.expr(0)
		}
		{
			p.SetState(45)
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
			p.SetState(47)
			p.Match(haruParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34902897163503640) != 0 {
			{
				p.SetState(48)
				p.expr(0)
			}
			p.SetState(53)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == haruParserT__20 {
				{
					p.SetState(49)
					p.Match(haruParserT__20)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(50)
					p.expr(0)
				}

				p.SetState(55)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(58)
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
			p.SetState(59)
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
			p.SetState(60)
			p.Literal()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(107)
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
			p.SetState(105)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(63)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(64)
					p.Match(haruParserT__5)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(65)
					p.expr(18)
				}

			case 2:
				localctx = NewMulExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(66)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(67)
					p.Match(haruParserT__6)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(68)
					p.expr(17)
				}

			case 3:
				localctx = NewDivExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(69)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(70)
					p.Match(haruParserT__7)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(71)
					p.expr(16)
				}

			case 4:
				localctx = NewModExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(72)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(73)
					p.Match(haruParserT__8)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(74)
					p.expr(15)
				}

			case 5:
				localctx = NewAddExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(75)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(76)
					p.Match(haruParserT__9)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(77)
					p.expr(14)
				}

			case 6:
				localctx = NewSubExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(78)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(79)
					p.Match(haruParserT__10)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(80)
					p.expr(13)
				}

			case 7:
				localctx = NewLtExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(81)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(82)
					p.Match(haruParserT__11)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(83)
					p.expr(12)
				}

			case 8:
				localctx = NewGtExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(84)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(85)
					p.Match(haruParserT__12)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(86)
					p.expr(11)
				}

			case 9:
				localctx = NewLeExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(87)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(88)
					p.Match(haruParserT__13)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(89)
					p.expr(10)
				}

			case 10:
				localctx = NewGeExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(91)
					p.Match(haruParserT__14)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(92)
					p.expr(9)
				}

			case 11:
				localctx = NewEqExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(93)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(94)
					p.Match(haruParserT__15)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(95)
					p.expr(8)
				}

			case 12:
				localctx = NewNeExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(96)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(97)
					p.Match(haruParserT__16)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(98)
					p.expr(7)
				}

			case 13:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(99)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(100)
					p.Match(haruParserT__17)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(101)
					p.expr(6)
				}

			case 14:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(102)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(103)
					p.Match(haruParserT__18)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(104)
					p.expr(5)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(109)
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
		p.SetState(110)
		p.Match(haruParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.Match(haruParserT__22)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.expr(0)
	}
	{
		p.SetState(113)
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

	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIntLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == haruParserT__10 {
			{
				p.SetState(115)
				p.Match(haruParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(118)
			p.Match(haruParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewFloatLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == haruParserT__10 {
			{
				p.SetState(119)
				p.Match(haruParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(122)
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
			p.SetState(123)
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
			p.SetState(124)
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
			p.SetState(125)
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
			p.SetState(126)
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

	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLetDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(129)
			p.Match(haruParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.Match(haruParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Type_()
		}
		{
			p.SetState(133)
			p.Match(haruParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.expr(0)
		}
		{
			p.SetState(135)
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
			p.SetState(137)
			p.Match(haruParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Match(haruParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.expr(0)
		}
		{
			p.SetState(141)
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
			p.Match(haruParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.Type_()
		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == haruParserT__22 {
			{
				p.SetState(147)
				p.Match(haruParserT__22)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(148)
				p.expr(0)
			}

		}
		{
			p.SetState(151)
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
			p.SetState(153)
			p.Match(haruParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(155)
			p.Match(haruParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.expr(0)
		}
		{
			p.SetState(157)
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
			p.SetState(159)
			p.Match(haruParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
			p.Match(haruParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)
			p.Type_()
		}
		{
			p.SetState(163)
			p.Match(haruParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)
			p.expr(0)
		}
		{
			p.SetState(165)
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
			p.SetState(167)
			p.Match(haruParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Match(haruParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)
			p.expr(0)
		}
		{
			p.SetState(171)
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
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case haruParserT__29:
		localctx = NewI8TypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(175)
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
			p.SetState(176)
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
			p.SetState(177)
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
			p.SetState(178)
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
			p.SetState(179)
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
			p.SetState(180)
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
			p.SetState(181)
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
			p.SetState(182)
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
			p.SetState(183)
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
			p.SetState(184)
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
			p.SetState(185)
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
			p.SetState(186)
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
			p.SetState(187)
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
			p.SetState(188)
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
			p.SetState(189)
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

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_ifStmt
	return p
}

func InitEmptyIfStmtContext(p *IfStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_ifStmt
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) CopyAll(ctx *IfStmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfBlockOnlyContext struct {
	IfStmtContext
}

func NewIfBlockOnlyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfBlockOnlyContext {
	var p = new(IfBlockOnlyContext)

	InitEmptyIfStmtContext(&p.IfStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStmtContext))

	return p
}

func (s *IfBlockOnlyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfBlockOnlyContext) Expr() IExprContext {
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

func (s *IfBlockOnlyContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfBlockOnlyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterIfBlockOnly(s)
	}
}

func (s *IfBlockOnlyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitIfBlockOnly(s)
	}
}

func (s *IfBlockOnlyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitIfBlockOnly(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfElseChainContext struct {
	IfStmtContext
}

func NewIfElseChainContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfElseChainContext {
	var p = new(IfElseChainContext)

	InitEmptyIfStmtContext(&p.IfStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStmtContext))

	return p
}

func (s *IfElseChainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseChainContext) Expr() IExprContext {
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

func (s *IfElseChainContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfElseChainContext) AllElseIfBlock() []IElseIfBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseIfBlockContext); ok {
			len++
		}
	}

	tst := make([]IElseIfBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseIfBlockContext); ok {
			tst[i] = t.(IElseIfBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfElseChainContext) ElseIfBlock(i int) IElseIfBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseIfBlockContext); ok {
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

	return t.(IElseIfBlockContext)
}

func (s *IfElseChainContext) ElseBlock() IElseBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseBlockContext)
}

func (s *IfElseChainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterIfElseChain(s)
	}
}

func (s *IfElseChainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitIfElseChain(s)
	}
}

func (s *IfElseChainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitIfElseChain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, haruParserRULE_ifStmt)
	var _la int

	var _alt int

	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIfBlockOnlyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(192)
			p.Match(haruParserT__44)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.Match(haruParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.expr(0)
		}
		{
			p.SetState(195)
			p.Match(haruParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)
			p.Block()
		}

	case 2:
		localctx = NewIfElseChainContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(198)
			p.Match(haruParserT__44)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
			p.Match(haruParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)
			p.expr(0)
		}
		{
			p.SetState(201)
			p.Match(haruParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(202)
			p.Block()
		}
		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(203)
					p.ElseIfBlock()
				}

			}
			p.SetState(208)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == haruParserT__45 {
			{
				p.SetState(209)
				p.ElseBlock()
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

// IElseIfBlockContext is an interface to support dynamic dispatch.
type IElseIfBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Block() IBlockContext

	// IsElseIfBlockContext differentiates from other interfaces.
	IsElseIfBlockContext()
}

type ElseIfBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfBlockContext() *ElseIfBlockContext {
	var p = new(ElseIfBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_elseIfBlock
	return p
}

func InitEmptyElseIfBlockContext(p *ElseIfBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_elseIfBlock
}

func (*ElseIfBlockContext) IsElseIfBlockContext() {}

func NewElseIfBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfBlockContext {
	var p = new(ElseIfBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_elseIfBlock

	return p
}

func (s *ElseIfBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfBlockContext) Expr() IExprContext {
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

func (s *ElseIfBlockContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseIfBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterElseIfBlock(s)
	}
}

func (s *ElseIfBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitElseIfBlock(s)
	}
}

func (s *ElseIfBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitElseIfBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) ElseIfBlock() (localctx IElseIfBlockContext) {
	localctx = NewElseIfBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, haruParserRULE_elseIfBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Match(haruParserT__45)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.Match(haruParserT__44)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(216)
		p.Match(haruParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)
		p.expr(0)
	}
	{
		p.SetState(218)
		p.Match(haruParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(219)
		p.Block()
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

// IElseBlockContext is an interface to support dynamic dispatch.
type IElseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext

	// IsElseBlockContext differentiates from other interfaces.
	IsElseBlockContext()
}

type ElseBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseBlockContext() *ElseBlockContext {
	var p = new(ElseBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_elseBlock
	return p
}

func InitEmptyElseBlockContext(p *ElseBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_elseBlock
}

func (*ElseBlockContext) IsElseBlockContext() {}

func NewElseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseBlockContext {
	var p = new(ElseBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_elseBlock

	return p
}

func (s *ElseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseBlockContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterElseBlock(s)
	}
}

func (s *ElseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitElseBlock(s)
	}
}

func (s *ElseBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitElseBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) ElseBlock() (localctx IElseBlockContext) {
	localctx = NewElseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, haruParserRULE_elseBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.Match(haruParserT__45)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
		p.Block()
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStatement() []IStatementContext {
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

func (s *BlockContext) Statement(i int) IStatementContext {
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

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, haruParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Match(haruParserT__46)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1161085151346690) != 0 {
		{
			p.SetState(225)
			p.Statement()
		}

		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(231)
		p.Match(haruParserT__47)
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
