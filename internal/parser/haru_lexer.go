// Code generated from ./grammar/haru.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type haruLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var HaruLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func harulexerLexerInit() {
	staticData := &HaruLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'print'", "';'", "'!'", "'('", "')'", "'**'", "'*'", "'/'", "'%'",
		"'+'", "'-'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", "'&&'",
		"'||'", "'['", "','", "']'", "'='", "'true'", "'false'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "WS", "ID", "NUMBER", "FLOAT", "STRING",
		"BYTE",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"WS", "ID", "NUMBER", "FLOAT", "STRING", "BYTE", "ESC",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 31, 194, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21,
		1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 25, 4, 25, 135, 8, 25, 11, 25, 12, 25, 136,
		1, 25, 1, 25, 1, 26, 1, 26, 5, 26, 143, 8, 26, 10, 26, 12, 26, 146, 9,
		26, 1, 27, 4, 27, 149, 8, 27, 11, 27, 12, 27, 150, 1, 28, 4, 28, 154, 8,
		28, 11, 28, 12, 28, 155, 1, 28, 1, 28, 4, 28, 160, 8, 28, 11, 28, 12, 28,
		161, 1, 29, 1, 29, 1, 29, 5, 29, 167, 8, 29, 10, 29, 12, 29, 170, 9, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 176, 8, 29, 10, 29, 12, 29, 179, 9,
		29, 1, 29, 3, 29, 182, 8, 29, 1, 30, 1, 30, 1, 30, 1, 30, 4, 30, 188, 8,
		30, 11, 30, 12, 30, 189, 1, 31, 1, 31, 1, 31, 2, 168, 177, 0, 32, 1, 1,
		3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23,
		12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41,
		21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59,
		30, 61, 31, 63, 0, 1, 0, 6, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 65, 90,
		97, 122, 3, 0, 48, 57, 65, 90, 97, 122, 1, 0, 48, 57, 1, 0, 48, 49, 3,
		0, 34, 34, 39, 39, 92, 92, 203, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0,
		0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0,
		0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0,
		0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1,
		0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51,
		1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0,
		59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 1, 65, 1, 0, 0, 0, 3, 71, 1, 0, 0, 0,
		5, 73, 1, 0, 0, 0, 7, 75, 1, 0, 0, 0, 9, 77, 1, 0, 0, 0, 11, 79, 1, 0,
		0, 0, 13, 82, 1, 0, 0, 0, 15, 84, 1, 0, 0, 0, 17, 86, 1, 0, 0, 0, 19, 88,
		1, 0, 0, 0, 21, 90, 1, 0, 0, 0, 23, 92, 1, 0, 0, 0, 25, 94, 1, 0, 0, 0,
		27, 96, 1, 0, 0, 0, 29, 99, 1, 0, 0, 0, 31, 102, 1, 0, 0, 0, 33, 105, 1,
		0, 0, 0, 35, 108, 1, 0, 0, 0, 37, 111, 1, 0, 0, 0, 39, 114, 1, 0, 0, 0,
		41, 116, 1, 0, 0, 0, 43, 118, 1, 0, 0, 0, 45, 120, 1, 0, 0, 0, 47, 122,
		1, 0, 0, 0, 49, 127, 1, 0, 0, 0, 51, 134, 1, 0, 0, 0, 53, 140, 1, 0, 0,
		0, 55, 148, 1, 0, 0, 0, 57, 153, 1, 0, 0, 0, 59, 181, 1, 0, 0, 0, 61, 183,
		1, 0, 0, 0, 63, 191, 1, 0, 0, 0, 65, 66, 5, 112, 0, 0, 66, 67, 5, 114,
		0, 0, 67, 68, 5, 105, 0, 0, 68, 69, 5, 110, 0, 0, 69, 70, 5, 116, 0, 0,
		70, 2, 1, 0, 0, 0, 71, 72, 5, 59, 0, 0, 72, 4, 1, 0, 0, 0, 73, 74, 5, 33,
		0, 0, 74, 6, 1, 0, 0, 0, 75, 76, 5, 40, 0, 0, 76, 8, 1, 0, 0, 0, 77, 78,
		5, 41, 0, 0, 78, 10, 1, 0, 0, 0, 79, 80, 5, 42, 0, 0, 80, 81, 5, 42, 0,
		0, 81, 12, 1, 0, 0, 0, 82, 83, 5, 42, 0, 0, 83, 14, 1, 0, 0, 0, 84, 85,
		5, 47, 0, 0, 85, 16, 1, 0, 0, 0, 86, 87, 5, 37, 0, 0, 87, 18, 1, 0, 0,
		0, 88, 89, 5, 43, 0, 0, 89, 20, 1, 0, 0, 0, 90, 91, 5, 45, 0, 0, 91, 22,
		1, 0, 0, 0, 92, 93, 5, 60, 0, 0, 93, 24, 1, 0, 0, 0, 94, 95, 5, 62, 0,
		0, 95, 26, 1, 0, 0, 0, 96, 97, 5, 60, 0, 0, 97, 98, 5, 61, 0, 0, 98, 28,
		1, 0, 0, 0, 99, 100, 5, 62, 0, 0, 100, 101, 5, 61, 0, 0, 101, 30, 1, 0,
		0, 0, 102, 103, 5, 61, 0, 0, 103, 104, 5, 61, 0, 0, 104, 32, 1, 0, 0, 0,
		105, 106, 5, 33, 0, 0, 106, 107, 5, 61, 0, 0, 107, 34, 1, 0, 0, 0, 108,
		109, 5, 38, 0, 0, 109, 110, 5, 38, 0, 0, 110, 36, 1, 0, 0, 0, 111, 112,
		5, 124, 0, 0, 112, 113, 5, 124, 0, 0, 113, 38, 1, 0, 0, 0, 114, 115, 5,
		91, 0, 0, 115, 40, 1, 0, 0, 0, 116, 117, 5, 44, 0, 0, 117, 42, 1, 0, 0,
		0, 118, 119, 5, 93, 0, 0, 119, 44, 1, 0, 0, 0, 120, 121, 5, 61, 0, 0, 121,
		46, 1, 0, 0, 0, 122, 123, 5, 116, 0, 0, 123, 124, 5, 114, 0, 0, 124, 125,
		5, 117, 0, 0, 125, 126, 5, 101, 0, 0, 126, 48, 1, 0, 0, 0, 127, 128, 5,
		102, 0, 0, 128, 129, 5, 97, 0, 0, 129, 130, 5, 108, 0, 0, 130, 131, 5,
		115, 0, 0, 131, 132, 5, 101, 0, 0, 132, 50, 1, 0, 0, 0, 133, 135, 7, 0,
		0, 0, 134, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0,
		136, 137, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 6, 25, 0, 0, 139,
		52, 1, 0, 0, 0, 140, 144, 7, 1, 0, 0, 141, 143, 7, 2, 0, 0, 142, 141, 1,
		0, 0, 0, 143, 146, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0,
		0, 145, 54, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 147, 149, 7, 3, 0, 0, 148,
		147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 150, 151,
		1, 0, 0, 0, 151, 56, 1, 0, 0, 0, 152, 154, 7, 3, 0, 0, 153, 152, 1, 0,
		0, 0, 154, 155, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0,
		156, 157, 1, 0, 0, 0, 157, 159, 5, 46, 0, 0, 158, 160, 7, 3, 0, 0, 159,
		158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 161, 162,
		1, 0, 0, 0, 162, 58, 1, 0, 0, 0, 163, 168, 5, 34, 0, 0, 164, 167, 3, 63,
		31, 0, 165, 167, 9, 0, 0, 0, 166, 164, 1, 0, 0, 0, 166, 165, 1, 0, 0, 0,
		167, 170, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 169,
		171, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 171, 182, 5, 34, 0, 0, 172, 177,
		5, 39, 0, 0, 173, 176, 3, 63, 31, 0, 174, 176, 9, 0, 0, 0, 175, 173, 1,
		0, 0, 0, 175, 174, 1, 0, 0, 0, 176, 179, 1, 0, 0, 0, 177, 178, 1, 0, 0,
		0, 177, 175, 1, 0, 0, 0, 178, 180, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 180,
		182, 5, 39, 0, 0, 181, 163, 1, 0, 0, 0, 181, 172, 1, 0, 0, 0, 182, 60,
		1, 0, 0, 0, 183, 184, 5, 48, 0, 0, 184, 185, 5, 98, 0, 0, 185, 187, 1,
		0, 0, 0, 186, 188, 7, 4, 0, 0, 187, 186, 1, 0, 0, 0, 188, 189, 1, 0, 0,
		0, 189, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 62, 1, 0, 0, 0, 191,
		192, 5, 92, 0, 0, 192, 193, 7, 5, 0, 0, 193, 64, 1, 0, 0, 0, 12, 0, 136,
		144, 150, 155, 161, 166, 168, 175, 177, 181, 189, 1, 6, 0, 0,
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

// haruLexerInit initializes any static state used to implement haruLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewharuLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func HaruLexerInit() {
	staticData := &HaruLexerLexerStaticData
	staticData.once.Do(harulexerLexerInit)
}

// NewharuLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewharuLexer(input antlr.CharStream) *haruLexer {
	HaruLexerInit()
	l := new(haruLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &HaruLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "haru.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// haruLexer tokens.
const (
	haruLexerT__0   = 1
	haruLexerT__1   = 2
	haruLexerT__2   = 3
	haruLexerT__3   = 4
	haruLexerT__4   = 5
	haruLexerT__5   = 6
	haruLexerT__6   = 7
	haruLexerT__7   = 8
	haruLexerT__8   = 9
	haruLexerT__9   = 10
	haruLexerT__10  = 11
	haruLexerT__11  = 12
	haruLexerT__12  = 13
	haruLexerT__13  = 14
	haruLexerT__14  = 15
	haruLexerT__15  = 16
	haruLexerT__16  = 17
	haruLexerT__17  = 18
	haruLexerT__18  = 19
	haruLexerT__19  = 20
	haruLexerT__20  = 21
	haruLexerT__21  = 22
	haruLexerT__22  = 23
	haruLexerT__23  = 24
	haruLexerT__24  = 25
	haruLexerWS     = 26
	haruLexerID     = 27
	haruLexerNUMBER = 28
	haruLexerFLOAT  = 29
	haruLexerSTRING = 30
	haruLexerBYTE   = 31
)
