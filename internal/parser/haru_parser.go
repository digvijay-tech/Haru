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
		"", "';'", "'print'", "'!'", "'('", "')'", "'**'", "'*'", "'/'", "'%'",
		"'+'", "'-'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", "'&&'",
		"'||'", "'['", "']'", "'='", "'true'", "'false'", "'let'", "':'", "'mut'",
		"'const'", "'i8'", "'i16'", "'i32'", "'i64'", "'int'", "'ui8'", "'ui16'",
		"'ui32'", "'ui64'", "'uint'", "'f32'", "'f64'", "'bool'", "'string'",
		"'byte'", "'if'", "'else'", "'{'", "'}'", "','", "'len'", "'function'",
		"'return'", "'while'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "WS", "ID", "NUMBER", "FLOAT", "STRING", "BYTE",
	}
	staticData.RuleNames = []string{
		"program", "statement", "printStmt", "expr", "assign", "literal", "varDecl",
		"type", "ifStmt", "elseIfBlock", "elseBlock", "block", "arrayDecl",
		"constArrayDecl", "letArrayDecl", "mutArrayDecl", "arrayType", "fixedArrayType",
		"arrayLiteral", "arrayItemAssign", "arrayReassign", "lenFunction", "functionDecl",
		"paramList", "param", "returnSignature", "returnStmt", "exprList", "functionCall",
		"argumentList", "whileLoop",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 58, 461, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1, 0, 5,
		0, 64, 8, 0, 10, 0, 12, 0, 67, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 82, 8, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 104, 8, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 5, 3, 148, 8, 3, 10, 3, 12, 3, 151, 9, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 5, 3, 5, 159, 8, 5, 1, 5, 1, 5, 3, 5, 163, 8, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 170, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 3, 6, 192, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 216, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 233,
		8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 5, 8, 247, 8, 8, 10, 8, 12, 8, 250, 9, 8, 1, 8, 3, 8, 253, 8, 8,
		3, 8, 255, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 5, 11, 269, 8, 11, 10, 11, 12, 11, 272, 9, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 12, 3, 12, 279, 8, 12, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 3, 13, 295, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 311, 8, 14, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 3, 15, 347, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 362, 8, 18,
		10, 18, 12, 18, 365, 9, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 371, 8,
		18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		22, 1, 22, 3, 22, 395, 8, 22, 1, 22, 1, 22, 3, 22, 399, 8, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 5, 23, 406, 8, 23, 10, 23, 12, 23, 409, 9, 23,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25, 419, 8,
		25, 10, 25, 12, 25, 422, 9, 25, 1, 25, 1, 25, 1, 26, 1, 26, 3, 26, 428,
		8, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 5, 27, 435, 8, 27, 10, 27, 12,
		27, 438, 9, 27, 1, 28, 1, 28, 1, 28, 3, 28, 443, 8, 28, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 29, 5, 29, 450, 8, 29, 10, 29, 12, 29, 453, 9, 29, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 363, 1, 6, 31, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 0, 0, 509, 0, 65, 1, 0, 0, 0, 2, 81,
		1, 0, 0, 0, 4, 83, 1, 0, 0, 0, 6, 103, 1, 0, 0, 0, 8, 152, 1, 0, 0, 0,
		10, 169, 1, 0, 0, 0, 12, 215, 1, 0, 0, 0, 14, 232, 1, 0, 0, 0, 16, 254,
		1, 0, 0, 0, 18, 256, 1, 0, 0, 0, 20, 263, 1, 0, 0, 0, 22, 266, 1, 0, 0,
		0, 24, 278, 1, 0, 0, 0, 26, 294, 1, 0, 0, 0, 28, 310, 1, 0, 0, 0, 30, 346,
		1, 0, 0, 0, 32, 348, 1, 0, 0, 0, 34, 352, 1, 0, 0, 0, 36, 370, 1, 0, 0,
		0, 38, 372, 1, 0, 0, 0, 40, 380, 1, 0, 0, 0, 42, 385, 1, 0, 0, 0, 44, 390,
		1, 0, 0, 0, 46, 402, 1, 0, 0, 0, 48, 410, 1, 0, 0, 0, 50, 414, 1, 0, 0,
		0, 52, 425, 1, 0, 0, 0, 54, 431, 1, 0, 0, 0, 56, 439, 1, 0, 0, 0, 58, 446,
		1, 0, 0, 0, 60, 454, 1, 0, 0, 0, 62, 64, 3, 2, 1, 0, 63, 62, 1, 0, 0, 0,
		64, 67, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 1, 1, 0,
		0, 0, 67, 65, 1, 0, 0, 0, 68, 82, 3, 4, 2, 0, 69, 82, 3, 12, 6, 0, 70,
		82, 3, 8, 4, 0, 71, 82, 3, 16, 8, 0, 72, 82, 3, 24, 12, 0, 73, 82, 3, 38,
		19, 0, 74, 82, 3, 40, 20, 0, 75, 82, 3, 44, 22, 0, 76, 82, 3, 52, 26, 0,
		77, 78, 3, 56, 28, 0, 78, 79, 5, 1, 0, 0, 79, 82, 1, 0, 0, 0, 80, 82, 3,
		60, 30, 0, 81, 68, 1, 0, 0, 0, 81, 69, 1, 0, 0, 0, 81, 70, 1, 0, 0, 0,
		81, 71, 1, 0, 0, 0, 81, 72, 1, 0, 0, 0, 81, 73, 1, 0, 0, 0, 81, 74, 1,
		0, 0, 0, 81, 75, 1, 0, 0, 0, 81, 76, 1, 0, 0, 0, 81, 77, 1, 0, 0, 0, 81,
		80, 1, 0, 0, 0, 82, 3, 1, 0, 0, 0, 83, 84, 5, 2, 0, 0, 84, 85, 3, 6, 3,
		0, 85, 86, 5, 1, 0, 0, 86, 5, 1, 0, 0, 0, 87, 88, 6, 3, -1, 0, 88, 89,
		5, 3, 0, 0, 89, 104, 3, 6, 3, 21, 90, 91, 5, 4, 0, 0, 91, 92, 3, 6, 3,
		0, 92, 93, 5, 5, 0, 0, 93, 104, 1, 0, 0, 0, 94, 104, 5, 54, 0, 0, 95, 96,
		5, 54, 0, 0, 96, 97, 5, 20, 0, 0, 97, 98, 3, 6, 3, 0, 98, 99, 5, 21, 0,
		0, 99, 104, 1, 0, 0, 0, 100, 104, 3, 56, 28, 0, 101, 104, 3, 42, 21, 0,
		102, 104, 3, 10, 5, 0, 103, 87, 1, 0, 0, 0, 103, 90, 1, 0, 0, 0, 103, 94,
		1, 0, 0, 0, 103, 95, 1, 0, 0, 0, 103, 100, 1, 0, 0, 0, 103, 101, 1, 0,
		0, 0, 103, 102, 1, 0, 0, 0, 104, 149, 1, 0, 0, 0, 105, 106, 10, 19, 0,
		0, 106, 107, 5, 6, 0, 0, 107, 148, 3, 6, 3, 20, 108, 109, 10, 18, 0, 0,
		109, 110, 5, 7, 0, 0, 110, 148, 3, 6, 3, 19, 111, 112, 10, 17, 0, 0, 112,
		113, 5, 8, 0, 0, 113, 148, 3, 6, 3, 18, 114, 115, 10, 16, 0, 0, 115, 116,
		5, 9, 0, 0, 116, 148, 3, 6, 3, 17, 117, 118, 10, 15, 0, 0, 118, 119, 5,
		10, 0, 0, 119, 148, 3, 6, 3, 16, 120, 121, 10, 14, 0, 0, 121, 122, 5, 11,
		0, 0, 122, 148, 3, 6, 3, 15, 123, 124, 10, 13, 0, 0, 124, 125, 5, 12, 0,
		0, 125, 148, 3, 6, 3, 14, 126, 127, 10, 12, 0, 0, 127, 128, 5, 13, 0, 0,
		128, 148, 3, 6, 3, 13, 129, 130, 10, 11, 0, 0, 130, 131, 5, 14, 0, 0, 131,
		148, 3, 6, 3, 12, 132, 133, 10, 10, 0, 0, 133, 134, 5, 15, 0, 0, 134, 148,
		3, 6, 3, 11, 135, 136, 10, 9, 0, 0, 136, 137, 5, 16, 0, 0, 137, 148, 3,
		6, 3, 10, 138, 139, 10, 8, 0, 0, 139, 140, 5, 17, 0, 0, 140, 148, 3, 6,
		3, 9, 141, 142, 10, 7, 0, 0, 142, 143, 5, 18, 0, 0, 143, 148, 3, 6, 3,
		8, 144, 145, 10, 6, 0, 0, 145, 146, 5, 19, 0, 0, 146, 148, 3, 6, 3, 7,
		147, 105, 1, 0, 0, 0, 147, 108, 1, 0, 0, 0, 147, 111, 1, 0, 0, 0, 147,
		114, 1, 0, 0, 0, 147, 117, 1, 0, 0, 0, 147, 120, 1, 0, 0, 0, 147, 123,
		1, 0, 0, 0, 147, 126, 1, 0, 0, 0, 147, 129, 1, 0, 0, 0, 147, 132, 1, 0,
		0, 0, 147, 135, 1, 0, 0, 0, 147, 138, 1, 0, 0, 0, 147, 141, 1, 0, 0, 0,
		147, 144, 1, 0, 0, 0, 148, 151, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149,
		150, 1, 0, 0, 0, 150, 7, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 152, 153, 5,
		54, 0, 0, 153, 154, 5, 22, 0, 0, 154, 155, 3, 6, 3, 0, 155, 156, 5, 1,
		0, 0, 156, 9, 1, 0, 0, 0, 157, 159, 5, 11, 0, 0, 158, 157, 1, 0, 0, 0,
		158, 159, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 170, 5, 55, 0, 0, 161,
		163, 5, 11, 0, 0, 162, 161, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164,
		1, 0, 0, 0, 164, 170, 5, 56, 0, 0, 165, 170, 5, 23, 0, 0, 166, 170, 5,
		24, 0, 0, 167, 170, 5, 57, 0, 0, 168, 170, 5, 58, 0, 0, 169, 158, 1, 0,
		0, 0, 169, 162, 1, 0, 0, 0, 169, 165, 1, 0, 0, 0, 169, 166, 1, 0, 0, 0,
		169, 167, 1, 0, 0, 0, 169, 168, 1, 0, 0, 0, 170, 11, 1, 0, 0, 0, 171, 172,
		5, 25, 0, 0, 172, 173, 5, 54, 0, 0, 173, 174, 5, 26, 0, 0, 174, 175, 3,
		14, 7, 0, 175, 176, 5, 22, 0, 0, 176, 177, 3, 6, 3, 0, 177, 178, 5, 1,
		0, 0, 178, 216, 1, 0, 0, 0, 179, 180, 5, 25, 0, 0, 180, 181, 5, 54, 0,
		0, 181, 182, 5, 22, 0, 0, 182, 183, 3, 6, 3, 0, 183, 184, 5, 1, 0, 0, 184,
		216, 1, 0, 0, 0, 185, 186, 5, 27, 0, 0, 186, 187, 5, 54, 0, 0, 187, 188,
		5, 26, 0, 0, 188, 191, 3, 14, 7, 0, 189, 190, 5, 22, 0, 0, 190, 192, 3,
		6, 3, 0, 191, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 193, 1, 0, 0,
		0, 193, 194, 5, 1, 0, 0, 194, 216, 1, 0, 0, 0, 195, 196, 5, 27, 0, 0, 196,
		197, 5, 54, 0, 0, 197, 198, 5, 22, 0, 0, 198, 199, 3, 6, 3, 0, 199, 200,
		5, 1, 0, 0, 200, 216, 1, 0, 0, 0, 201, 202, 5, 28, 0, 0, 202, 203, 5, 54,
		0, 0, 203, 204, 5, 26, 0, 0, 204, 205, 3, 14, 7, 0, 205, 206, 5, 22, 0,
		0, 206, 207, 3, 6, 3, 0, 207, 208, 5, 1, 0, 0, 208, 216, 1, 0, 0, 0, 209,
		210, 5, 28, 0, 0, 210, 211, 5, 54, 0, 0, 211, 212, 5, 22, 0, 0, 212, 213,
		3, 6, 3, 0, 213, 214, 5, 1, 0, 0, 214, 216, 1, 0, 0, 0, 215, 171, 1, 0,
		0, 0, 215, 179, 1, 0, 0, 0, 215, 185, 1, 0, 0, 0, 215, 195, 1, 0, 0, 0,
		215, 201, 1, 0, 0, 0, 215, 209, 1, 0, 0, 0, 216, 13, 1, 0, 0, 0, 217, 233,
		5, 29, 0, 0, 218, 233, 5, 30, 0, 0, 219, 233, 5, 31, 0, 0, 220, 233, 5,
		32, 0, 0, 221, 233, 5, 33, 0, 0, 222, 233, 5, 34, 0, 0, 223, 233, 5, 35,
		0, 0, 224, 233, 5, 36, 0, 0, 225, 233, 5, 37, 0, 0, 226, 233, 5, 38, 0,
		0, 227, 233, 5, 39, 0, 0, 228, 233, 5, 40, 0, 0, 229, 233, 5, 41, 0, 0,
		230, 233, 5, 42, 0, 0, 231, 233, 5, 43, 0, 0, 232, 217, 1, 0, 0, 0, 232,
		218, 1, 0, 0, 0, 232, 219, 1, 0, 0, 0, 232, 220, 1, 0, 0, 0, 232, 221,
		1, 0, 0, 0, 232, 222, 1, 0, 0, 0, 232, 223, 1, 0, 0, 0, 232, 224, 1, 0,
		0, 0, 232, 225, 1, 0, 0, 0, 232, 226, 1, 0, 0, 0, 232, 227, 1, 0, 0, 0,
		232, 228, 1, 0, 0, 0, 232, 229, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 232,
		231, 1, 0, 0, 0, 233, 15, 1, 0, 0, 0, 234, 235, 5, 44, 0, 0, 235, 236,
		5, 4, 0, 0, 236, 237, 3, 6, 3, 0, 237, 238, 5, 5, 0, 0, 238, 239, 3, 22,
		11, 0, 239, 255, 1, 0, 0, 0, 240, 241, 5, 44, 0, 0, 241, 242, 5, 4, 0,
		0, 242, 243, 3, 6, 3, 0, 243, 244, 5, 5, 0, 0, 244, 248, 3, 22, 11, 0,
		245, 247, 3, 18, 9, 0, 246, 245, 1, 0, 0, 0, 247, 250, 1, 0, 0, 0, 248,
		246, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 252, 1, 0, 0, 0, 250, 248,
		1, 0, 0, 0, 251, 253, 3, 20, 10, 0, 252, 251, 1, 0, 0, 0, 252, 253, 1,
		0, 0, 0, 253, 255, 1, 0, 0, 0, 254, 234, 1, 0, 0, 0, 254, 240, 1, 0, 0,
		0, 255, 17, 1, 0, 0, 0, 256, 257, 5, 45, 0, 0, 257, 258, 5, 44, 0, 0, 258,
		259, 5, 4, 0, 0, 259, 260, 3, 6, 3, 0, 260, 261, 5, 5, 0, 0, 261, 262,
		3, 22, 11, 0, 262, 19, 1, 0, 0, 0, 263, 264, 5, 45, 0, 0, 264, 265, 3,
		22, 11, 0, 265, 21, 1, 0, 0, 0, 266, 270, 5, 46, 0, 0, 267, 269, 3, 2,
		1, 0, 268, 267, 1, 0, 0, 0, 269, 272, 1, 0, 0, 0, 270, 268, 1, 0, 0, 0,
		270, 271, 1, 0, 0, 0, 271, 273, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 273,
		274, 5, 47, 0, 0, 274, 23, 1, 0, 0, 0, 275, 279, 3, 26, 13, 0, 276, 279,
		3, 28, 14, 0, 277, 279, 3, 30, 15, 0, 278, 275, 1, 0, 0, 0, 278, 276, 1,
		0, 0, 0, 278, 277, 1, 0, 0, 0, 279, 25, 1, 0, 0, 0, 280, 281, 5, 28, 0,
		0, 281, 282, 5, 54, 0, 0, 282, 283, 5, 26, 0, 0, 283, 284, 3, 32, 16, 0,
		284, 285, 5, 22, 0, 0, 285, 286, 3, 36, 18, 0, 286, 287, 5, 1, 0, 0, 287,
		295, 1, 0, 0, 0, 288, 289, 5, 28, 0, 0, 289, 290, 5, 54, 0, 0, 290, 291,
		5, 22, 0, 0, 291, 292, 3, 36, 18, 0, 292, 293, 5, 1, 0, 0, 293, 295, 1,
		0, 0, 0, 294, 280, 1, 0, 0, 0, 294, 288, 1, 0, 0, 0, 295, 27, 1, 0, 0,
		0, 296, 297, 5, 25, 0, 0, 297, 298, 5, 54, 0, 0, 298, 299, 5, 26, 0, 0,
		299, 300, 3, 32, 16, 0, 300, 301, 5, 22, 0, 0, 301, 302, 3, 36, 18, 0,
		302, 303, 5, 1, 0, 0, 303, 311, 1, 0, 0, 0, 304, 305, 5, 25, 0, 0, 305,
		306, 5, 54, 0, 0, 306, 307, 5, 22, 0, 0, 307, 308, 3, 36, 18, 0, 308, 309,
		5, 1, 0, 0, 309, 311, 1, 0, 0, 0, 310, 296, 1, 0, 0, 0, 310, 304, 1, 0,
		0, 0, 311, 29, 1, 0, 0, 0, 312, 313, 5, 27, 0, 0, 313, 314, 5, 54, 0, 0,
		314, 315, 5, 26, 0, 0, 315, 316, 3, 34, 17, 0, 316, 317, 5, 22, 0, 0, 317,
		318, 3, 36, 18, 0, 318, 319, 5, 1, 0, 0, 319, 347, 1, 0, 0, 0, 320, 321,
		5, 27, 0, 0, 321, 322, 5, 54, 0, 0, 322, 323, 5, 26, 0, 0, 323, 324, 3,
		34, 17, 0, 324, 325, 5, 1, 0, 0, 325, 347, 1, 0, 0, 0, 326, 327, 5, 27,
		0, 0, 327, 328, 5, 54, 0, 0, 328, 329, 5, 26, 0, 0, 329, 330, 3, 32, 16,
		0, 330, 331, 5, 22, 0, 0, 331, 332, 3, 36, 18, 0, 332, 333, 5, 1, 0, 0,
		333, 347, 1, 0, 0, 0, 334, 335, 5, 27, 0, 0, 335, 336, 5, 54, 0, 0, 336,
		337, 5, 26, 0, 0, 337, 338, 3, 32, 16, 0, 338, 339, 5, 1, 0, 0, 339, 347,
		1, 0, 0, 0, 340, 341, 5, 27, 0, 0, 341, 342, 5, 54, 0, 0, 342, 343, 5,
		22, 0, 0, 343, 344, 3, 36, 18, 0, 344, 345, 5, 1, 0, 0, 345, 347, 1, 0,
		0, 0, 346, 312, 1, 0, 0, 0, 346, 320, 1, 0, 0, 0, 346, 326, 1, 0, 0, 0,
		346, 334, 1, 0, 0, 0, 346, 340, 1, 0, 0, 0, 347, 31, 1, 0, 0, 0, 348, 349,
		5, 20, 0, 0, 349, 350, 5, 21, 0, 0, 350, 351, 3, 14, 7, 0, 351, 33, 1,
		0, 0, 0, 352, 353, 5, 20, 0, 0, 353, 354, 5, 55, 0, 0, 354, 355, 5, 21,
		0, 0, 355, 356, 3, 14, 7, 0, 356, 35, 1, 0, 0, 0, 357, 358, 5, 20, 0, 0,
		358, 363, 3, 6, 3, 0, 359, 360, 5, 48, 0, 0, 360, 362, 3, 6, 3, 0, 361,
		359, 1, 0, 0, 0, 362, 365, 1, 0, 0, 0, 363, 364, 1, 0, 0, 0, 363, 361,
		1, 0, 0, 0, 364, 366, 1, 0, 0, 0, 365, 363, 1, 0, 0, 0, 366, 367, 5, 21,
		0, 0, 367, 371, 1, 0, 0, 0, 368, 369, 5, 20, 0, 0, 369, 371, 5, 21, 0,
		0, 370, 357, 1, 0, 0, 0, 370, 368, 1, 0, 0, 0, 371, 37, 1, 0, 0, 0, 372,
		373, 5, 54, 0, 0, 373, 374, 5, 20, 0, 0, 374, 375, 3, 6, 3, 0, 375, 376,
		5, 21, 0, 0, 376, 377, 5, 22, 0, 0, 377, 378, 3, 6, 3, 0, 378, 379, 5,
		1, 0, 0, 379, 39, 1, 0, 0, 0, 380, 381, 5, 54, 0, 0, 381, 382, 5, 22, 0,
		0, 382, 383, 3, 36, 18, 0, 383, 384, 5, 1, 0, 0, 384, 41, 1, 0, 0, 0, 385,
		386, 5, 49, 0, 0, 386, 387, 5, 4, 0, 0, 387, 388, 3, 6, 3, 0, 388, 389,
		5, 5, 0, 0, 389, 43, 1, 0, 0, 0, 390, 391, 5, 50, 0, 0, 391, 392, 5, 54,
		0, 0, 392, 394, 5, 4, 0, 0, 393, 395, 3, 46, 23, 0, 394, 393, 1, 0, 0,
		0, 394, 395, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396, 398, 5, 5, 0, 0, 397,
		399, 3, 50, 25, 0, 398, 397, 1, 0, 0, 0, 398, 399, 1, 0, 0, 0, 399, 400,
		1, 0, 0, 0, 400, 401, 3, 22, 11, 0, 401, 45, 1, 0, 0, 0, 402, 407, 3, 48,
		24, 0, 403, 404, 5, 48, 0, 0, 404, 406, 3, 48, 24, 0, 405, 403, 1, 0, 0,
		0, 406, 409, 1, 0, 0, 0, 407, 405, 1, 0, 0, 0, 407, 408, 1, 0, 0, 0, 408,
		47, 1, 0, 0, 0, 409, 407, 1, 0, 0, 0, 410, 411, 5, 54, 0, 0, 411, 412,
		5, 26, 0, 0, 412, 413, 3, 14, 7, 0, 413, 49, 1, 0, 0, 0, 414, 415, 5, 12,
		0, 0, 415, 420, 3, 14, 7, 0, 416, 417, 5, 48, 0, 0, 417, 419, 3, 14, 7,
		0, 418, 416, 1, 0, 0, 0, 419, 422, 1, 0, 0, 0, 420, 418, 1, 0, 0, 0, 420,
		421, 1, 0, 0, 0, 421, 423, 1, 0, 0, 0, 422, 420, 1, 0, 0, 0, 423, 424,
		5, 13, 0, 0, 424, 51, 1, 0, 0, 0, 425, 427, 5, 51, 0, 0, 426, 428, 3, 54,
		27, 0, 427, 426, 1, 0, 0, 0, 427, 428, 1, 0, 0, 0, 428, 429, 1, 0, 0, 0,
		429, 430, 5, 1, 0, 0, 430, 53, 1, 0, 0, 0, 431, 436, 3, 6, 3, 0, 432, 433,
		5, 48, 0, 0, 433, 435, 3, 6, 3, 0, 434, 432, 1, 0, 0, 0, 435, 438, 1, 0,
		0, 0, 436, 434, 1, 0, 0, 0, 436, 437, 1, 0, 0, 0, 437, 55, 1, 0, 0, 0,
		438, 436, 1, 0, 0, 0, 439, 440, 5, 54, 0, 0, 440, 442, 5, 4, 0, 0, 441,
		443, 3, 58, 29, 0, 442, 441, 1, 0, 0, 0, 442, 443, 1, 0, 0, 0, 443, 444,
		1, 0, 0, 0, 444, 445, 5, 5, 0, 0, 445, 57, 1, 0, 0, 0, 446, 451, 3, 6,
		3, 0, 447, 448, 5, 48, 0, 0, 448, 450, 3, 6, 3, 0, 449, 447, 1, 0, 0, 0,
		450, 453, 1, 0, 0, 0, 451, 449, 1, 0, 0, 0, 451, 452, 1, 0, 0, 0, 452,
		59, 1, 0, 0, 0, 453, 451, 1, 0, 0, 0, 454, 455, 5, 52, 0, 0, 455, 456,
		5, 4, 0, 0, 456, 457, 3, 6, 3, 0, 457, 458, 5, 5, 0, 0, 458, 459, 3, 22,
		11, 0, 459, 61, 1, 0, 0, 0, 29, 65, 81, 103, 147, 149, 158, 162, 169, 191,
		215, 232, 248, 252, 254, 270, 278, 294, 310, 346, 363, 370, 394, 398, 407,
		420, 427, 436, 442, 451,
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
	haruParserT__48  = 49
	haruParserT__49  = 50
	haruParserT__50  = 51
	haruParserT__51  = 52
	haruParserWS     = 53
	haruParserID     = 54
	haruParserNUMBER = 55
	haruParserFLOAT  = 56
	haruParserSTRING = 57
	haruParserBYTE   = 58
)

// haruParser rules.
const (
	haruParserRULE_program         = 0
	haruParserRULE_statement       = 1
	haruParserRULE_printStmt       = 2
	haruParserRULE_expr            = 3
	haruParserRULE_assign          = 4
	haruParserRULE_literal         = 5
	haruParserRULE_varDecl         = 6
	haruParserRULE_type            = 7
	haruParserRULE_ifStmt          = 8
	haruParserRULE_elseIfBlock     = 9
	haruParserRULE_elseBlock       = 10
	haruParserRULE_block           = 11
	haruParserRULE_arrayDecl       = 12
	haruParserRULE_constArrayDecl  = 13
	haruParserRULE_letArrayDecl    = 14
	haruParserRULE_mutArrayDecl    = 15
	haruParserRULE_arrayType       = 16
	haruParserRULE_fixedArrayType  = 17
	haruParserRULE_arrayLiteral    = 18
	haruParserRULE_arrayItemAssign = 19
	haruParserRULE_arrayReassign   = 20
	haruParserRULE_lenFunction     = 21
	haruParserRULE_functionDecl    = 22
	haruParserRULE_paramList       = 23
	haruParserRULE_param           = 24
	haruParserRULE_returnSignature = 25
	haruParserRULE_returnStmt      = 26
	haruParserRULE_exprList        = 27
	haruParserRULE_functionCall    = 28
	haruParserRULE_argumentList    = 29
	haruParserRULE_whileLoop       = 30
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
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&25913290479632388) != 0 {
		{
			p.SetState(62)
			p.Statement()
		}

		p.SetState(67)
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

type ArrayReassignStatementContext struct {
	StatementContext
}

func NewArrayReassignStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayReassignStatementContext {
	var p = new(ArrayReassignStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ArrayReassignStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayReassignStatementContext) ArrayReassign() IArrayReassignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayReassignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayReassignContext)
}

func (s *ArrayReassignStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterArrayReassignStatement(s)
	}
}

func (s *ArrayReassignStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitArrayReassignStatement(s)
	}
}

func (s *ArrayReassignStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitArrayReassignStatement(s)

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

type ReturnStmtStatementContext struct {
	StatementContext
}

func NewReturnStmtStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStmtStatementContext {
	var p = new(ReturnStmtStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ReturnStmtStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtStatementContext) ReturnStmt() IReturnStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStmtContext)
}

func (s *ReturnStmtStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterReturnStmtStatement(s)
	}
}

func (s *ReturnStmtStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitReturnStmtStatement(s)
	}
}

func (s *ReturnStmtStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitReturnStmtStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type WhileLoopStatementContext struct {
	StatementContext
}

func NewWhileLoopStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileLoopStatementContext {
	var p = new(WhileLoopStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *WhileLoopStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileLoopStatementContext) WhileLoop() IWhileLoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileLoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileLoopContext)
}

func (s *WhileLoopStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterWhileLoopStatement(s)
	}
}

func (s *WhileLoopStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitWhileLoopStatement(s)
	}
}

func (s *WhileLoopStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitWhileLoopStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallStatementContext struct {
	StatementContext
}

func NewFunctionCallStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallStatementContext {
	var p = new(FunctionCallStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *FunctionCallStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallStatementContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *FunctionCallStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterFunctionCallStatement(s)
	}
}

func (s *FunctionCallStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitFunctionCallStatement(s)
	}
}

func (s *FunctionCallStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitFunctionCallStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayDeclStatementContext struct {
	StatementContext
}

func NewArrayDeclStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayDeclStatementContext {
	var p = new(ArrayDeclStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ArrayDeclStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayDeclStatementContext) ArrayDecl() IArrayDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayDeclContext)
}

func (s *ArrayDeclStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterArrayDeclStatement(s)
	}
}

func (s *ArrayDeclStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitArrayDeclStatement(s)
	}
}

func (s *ArrayDeclStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitArrayDeclStatement(s)

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

type ArrayIndexAssignStatementContext struct {
	StatementContext
}

func NewArrayIndexAssignStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayIndexAssignStatementContext {
	var p = new(ArrayIndexAssignStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ArrayIndexAssignStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayIndexAssignStatementContext) ArrayItemAssign() IArrayItemAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayItemAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayItemAssignContext)
}

func (s *ArrayIndexAssignStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterArrayIndexAssignStatement(s)
	}
}

func (s *ArrayIndexAssignStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitArrayIndexAssignStatement(s)
	}
}

func (s *ArrayIndexAssignStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitArrayIndexAssignStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionDeclStatementContext struct {
	StatementContext
}

func NewFunctionDeclStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionDeclStatementContext {
	var p = new(FunctionDeclStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *FunctionDeclStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclStatementContext) FunctionDecl() IFunctionDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclContext)
}

func (s *FunctionDeclStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterFunctionDeclStatement(s)
	}
}

func (s *FunctionDeclStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitFunctionDeclStatement(s)
	}
}

func (s *FunctionDeclStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitFunctionDeclStatement(s)

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
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPrintStmtStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(68)
			p.PrintStmt()
		}

	case 2:
		localctx = NewVarDeclStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(69)
			p.VarDecl()
		}

	case 3:
		localctx = NewAssignStmtStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(70)
			p.Assign()
		}

	case 4:
		localctx = NewIfStmtStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(71)
			p.IfStmt()
		}

	case 5:
		localctx = NewArrayDeclStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(72)
			p.ArrayDecl()
		}

	case 6:
		localctx = NewArrayIndexAssignStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(73)
			p.ArrayItemAssign()
		}

	case 7:
		localctx = NewArrayReassignStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(74)
			p.ArrayReassign()
		}

	case 8:
		localctx = NewFunctionDeclStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(75)
			p.FunctionDecl()
		}

	case 9:
		localctx = NewReturnStmtStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(76)
			p.ReturnStmt()
		}

	case 10:
		localctx = NewFunctionCallStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(77)
			p.FunctionCall()
		}
		{
			p.SetState(78)
			p.Match(haruParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewWhileLoopStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(80)
			p.WhileLoop()
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
		p.SetState(83)
		p.Match(haruParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
		p.expr(0)
	}
	{
		p.SetState(85)
		p.Match(haruParserT__0)
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

type IndexExprContext struct {
	ExprContext
}

func NewIndexExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexExprContext {
	var p = new(IndexExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IndexExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexExprContext) ID() antlr.TerminalNode {
	return s.GetToken(haruParserID, 0)
}

func (s *IndexExprContext) Expr() IExprContext {
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

func (s *IndexExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterIndexExpr(s)
	}
}

func (s *IndexExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitIndexExpr(s)
	}
}

func (s *IndexExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitIndexExpr(s)

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

type FunctionCallExprContext struct {
	ExprContext
}

func NewFunctionCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallExprContext {
	var p = new(FunctionCallExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FunctionCallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallExprContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *FunctionCallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterFunctionCallExpr(s)
	}
}

func (s *FunctionCallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitFunctionCallExpr(s)
	}
}

func (s *FunctionCallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitFunctionCallExpr(s)

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

type LenFunctionExprContext struct {
	ExprContext
}

func NewLenFunctionExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LenFunctionExprContext {
	var p = new(LenFunctionExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LenFunctionExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LenFunctionExprContext) LenFunction() ILenFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILenFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILenFunctionContext)
}

func (s *LenFunctionExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterLenFunctionExpr(s)
	}
}

func (s *LenFunctionExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitLenFunctionExpr(s)
	}
}

func (s *LenFunctionExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitLenFunctionExpr(s)

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
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(88)
			p.Match(haruParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)
			p.expr(21)
		}

	case 2:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(90)
			p.Match(haruParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.expr(0)
		}
		{
			p.SetState(92)
			p.Match(haruParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewVarExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(94)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewIndexExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(95)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)
			p.Match(haruParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(97)
			p.expr(0)
		}
		{
			p.SetState(98)
			p.Match(haruParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewFunctionCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(100)
			p.FunctionCall()
		}

	case 6:
		localctx = NewLenFunctionExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(101)
			p.LenFunction()
		}

	case 7:
		localctx = NewLitExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(102)
			p.Literal()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(147)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(105)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(106)
					p.Match(haruParserT__5)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(107)
					p.expr(20)
				}

			case 2:
				localctx = NewMulExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(108)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(109)
					p.Match(haruParserT__6)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(110)
					p.expr(19)
				}

			case 3:
				localctx = NewDivExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(111)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(112)
					p.Match(haruParserT__7)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(113)
					p.expr(18)
				}

			case 4:
				localctx = NewModExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(114)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(115)
					p.Match(haruParserT__8)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(116)
					p.expr(17)
				}

			case 5:
				localctx = NewAddExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(117)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(118)
					p.Match(haruParserT__9)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(119)
					p.expr(16)
				}

			case 6:
				localctx = NewSubExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(120)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(121)
					p.Match(haruParserT__10)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(122)
					p.expr(15)
				}

			case 7:
				localctx = NewLtExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(123)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(124)
					p.Match(haruParserT__11)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(125)
					p.expr(14)
				}

			case 8:
				localctx = NewGtExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(126)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(127)
					p.Match(haruParserT__12)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(128)
					p.expr(13)
				}

			case 9:
				localctx = NewLeExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(129)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(130)
					p.Match(haruParserT__13)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(131)
					p.expr(12)
				}

			case 10:
				localctx = NewGeExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(132)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(133)
					p.Match(haruParserT__14)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(134)
					p.expr(11)
				}

			case 11:
				localctx = NewEqExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(135)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(136)
					p.Match(haruParserT__15)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(137)
					p.expr(10)
				}

			case 12:
				localctx = NewNeExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(138)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(139)
					p.Match(haruParserT__16)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(140)
					p.expr(9)
				}

			case 13:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(141)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(142)
					p.Match(haruParserT__17)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(143)
					p.expr(8)
				}

			case 14:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, haruParserRULE_expr)
				p.SetState(144)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(145)
					p.Match(haruParserT__18)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(146)
					p.expr(7)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
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
		p.SetState(152)
		p.Match(haruParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(haruParserT__21)
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
		p.Match(haruParserT__0)
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

	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIntLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == haruParserT__10 {
			{
				p.SetState(157)
				p.Match(haruParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(160)
			p.Match(haruParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewFloatLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == haruParserT__10 {
			{
				p.SetState(161)
				p.Match(haruParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(164)
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
			p.SetState(165)
			p.Match(haruParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewFalseLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(166)
			p.Match(haruParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(167)
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
			p.SetState(168)
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

	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLetDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)
			p.Match(haruParserT__24)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)
			p.Match(haruParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
			p.Type_()
		}
		{
			p.SetState(175)
			p.Match(haruParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.expr(0)
		}
		{
			p.SetState(177)
			p.Match(haruParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewLetInferDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(179)
			p.Match(haruParserT__24)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(181)
			p.Match(haruParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.expr(0)
		}
		{
			p.SetState(183)
			p.Match(haruParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewMutDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(185)
			p.Match(haruParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Match(haruParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.Type_()
		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == haruParserT__21 {
			{
				p.SetState(189)
				p.Match(haruParserT__21)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(190)
				p.expr(0)
			}

		}
		{
			p.SetState(193)
			p.Match(haruParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewMutInferDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(195)
			p.Match(haruParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.Match(haruParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.expr(0)
		}
		{
			p.SetState(199)
			p.Match(haruParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewConstDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(201)
			p.Match(haruParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(202)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)
			p.Match(haruParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)
			p.Type_()
		}
		{
			p.SetState(205)
			p.Match(haruParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)
			p.expr(0)
		}
		{
			p.SetState(207)
			p.Match(haruParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewConstInferDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(209)
			p.Match(haruParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(211)
			p.Match(haruParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(212)
			p.expr(0)
		}
		{
			p.SetState(213)
			p.Match(haruParserT__0)
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
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case haruParserT__28:
		localctx = NewI8TypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(217)
			p.Match(haruParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__29:
		localctx = NewI16TypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(218)
			p.Match(haruParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__30:
		localctx = NewI32TypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(219)
			p.Match(haruParserT__30)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__31:
		localctx = NewI64TypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(220)
			p.Match(haruParserT__31)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__32:
		localctx = NewIntTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(221)
			p.Match(haruParserT__32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__33:
		localctx = NewUI8TypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(222)
			p.Match(haruParserT__33)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__34:
		localctx = NewUI16TypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(223)
			p.Match(haruParserT__34)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__35:
		localctx = NewUI32TypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(224)
			p.Match(haruParserT__35)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__36:
		localctx = NewUI64TypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(225)
			p.Match(haruParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__37:
		localctx = NewUIntTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(226)
			p.Match(haruParserT__37)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__38:
		localctx = NewF32TypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(227)
			p.Match(haruParserT__38)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__39:
		localctx = NewF64TypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(228)
			p.Match(haruParserT__39)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__40:
		localctx = NewBoolTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(229)
			p.Match(haruParserT__40)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__41:
		localctx = NewStringTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(230)
			p.Match(haruParserT__41)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case haruParserT__42:
		localctx = NewByteTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(231)
			p.Match(haruParserT__42)
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

	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIfBlockOnlyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(234)
			p.Match(haruParserT__43)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)
			p.Match(haruParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(236)
			p.expr(0)
		}
		{
			p.SetState(237)
			p.Match(haruParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.Block()
		}

	case 2:
		localctx = NewIfElseChainContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(240)
			p.Match(haruParserT__43)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(241)
			p.Match(haruParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.expr(0)
		}
		{
			p.SetState(243)
			p.Match(haruParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)
			p.Block()
		}
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(245)
					p.ElseIfBlock()
				}

			}
			p.SetState(250)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == haruParserT__44 {
			{
				p.SetState(251)
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
		p.SetState(256)
		p.Match(haruParserT__44)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(257)
		p.Match(haruParserT__43)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
		p.Match(haruParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(259)
		p.expr(0)
	}
	{
		p.SetState(260)
		p.Match(haruParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
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
		p.SetState(263)
		p.Match(haruParserT__44)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(264)
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
		p.SetState(266)
		p.Match(haruParserT__45)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&25913290479632388) != 0 {
		{
			p.SetState(267)
			p.Statement()
		}

		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(273)
		p.Match(haruParserT__46)
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

// IArrayDeclContext is an interface to support dynamic dispatch.
type IArrayDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ConstArrayDecl() IConstArrayDeclContext
	LetArrayDecl() ILetArrayDeclContext
	MutArrayDecl() IMutArrayDeclContext

	// IsArrayDeclContext differentiates from other interfaces.
	IsArrayDeclContext()
}

type ArrayDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayDeclContext() *ArrayDeclContext {
	var p = new(ArrayDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_arrayDecl
	return p
}

func InitEmptyArrayDeclContext(p *ArrayDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_arrayDecl
}

func (*ArrayDeclContext) IsArrayDeclContext() {}

func NewArrayDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayDeclContext {
	var p = new(ArrayDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_arrayDecl

	return p
}

func (s *ArrayDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayDeclContext) ConstArrayDecl() IConstArrayDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstArrayDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstArrayDeclContext)
}

func (s *ArrayDeclContext) LetArrayDecl() ILetArrayDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILetArrayDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILetArrayDeclContext)
}

func (s *ArrayDeclContext) MutArrayDecl() IMutArrayDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMutArrayDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMutArrayDeclContext)
}

func (s *ArrayDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterArrayDecl(s)
	}
}

func (s *ArrayDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitArrayDecl(s)
	}
}

func (s *ArrayDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitArrayDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) ArrayDecl() (localctx IArrayDeclContext) {
	localctx = NewArrayDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, haruParserRULE_arrayDecl)
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case haruParserT__27:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(275)
			p.ConstArrayDecl()
		}

	case haruParserT__24:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(276)
			p.LetArrayDecl()
		}

	case haruParserT__26:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(277)
			p.MutArrayDecl()
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

// IConstArrayDeclContext is an interface to support dynamic dispatch.
type IConstArrayDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsConstArrayDeclContext differentiates from other interfaces.
	IsConstArrayDeclContext()
}

type ConstArrayDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstArrayDeclContext() *ConstArrayDeclContext {
	var p = new(ConstArrayDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_constArrayDecl
	return p
}

func InitEmptyConstArrayDeclContext(p *ConstArrayDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_constArrayDecl
}

func (*ConstArrayDeclContext) IsConstArrayDeclContext() {}

func NewConstArrayDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstArrayDeclContext {
	var p = new(ConstArrayDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_constArrayDecl

	return p
}

func (s *ConstArrayDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstArrayDeclContext) CopyAll(ctx *ConstArrayDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ConstArrayDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstArrayDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ConstExplicitArrayDeclContext struct {
	ConstArrayDeclContext
}

func NewConstExplicitArrayDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstExplicitArrayDeclContext {
	var p = new(ConstExplicitArrayDeclContext)

	InitEmptyConstArrayDeclContext(&p.ConstArrayDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConstArrayDeclContext))

	return p
}

func (s *ConstExplicitArrayDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstExplicitArrayDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(haruParserID, 0)
}

func (s *ConstExplicitArrayDeclContext) ArrayType() IArrayTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayTypeContext)
}

func (s *ConstExplicitArrayDeclContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *ConstExplicitArrayDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterConstExplicitArrayDecl(s)
	}
}

func (s *ConstExplicitArrayDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitConstExplicitArrayDecl(s)
	}
}

func (s *ConstExplicitArrayDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitConstExplicitArrayDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConstImplicitArrayDeclContext struct {
	ConstArrayDeclContext
}

func NewConstImplicitArrayDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstImplicitArrayDeclContext {
	var p = new(ConstImplicitArrayDeclContext)

	InitEmptyConstArrayDeclContext(&p.ConstArrayDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConstArrayDeclContext))

	return p
}

func (s *ConstImplicitArrayDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstImplicitArrayDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(haruParserID, 0)
}

func (s *ConstImplicitArrayDeclContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *ConstImplicitArrayDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterConstImplicitArrayDecl(s)
	}
}

func (s *ConstImplicitArrayDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitConstImplicitArrayDecl(s)
	}
}

func (s *ConstImplicitArrayDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitConstImplicitArrayDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) ConstArrayDecl() (localctx IConstArrayDeclContext) {
	localctx = NewConstArrayDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, haruParserRULE_constArrayDecl)
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewConstExplicitArrayDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(280)
			p.Match(haruParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(281)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(282)
			p.Match(haruParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(283)
			p.ArrayType()
		}
		{
			p.SetState(284)
			p.Match(haruParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)
			p.ArrayLiteral()
		}
		{
			p.SetState(286)
			p.Match(haruParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewConstImplicitArrayDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(288)
			p.Match(haruParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(289)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(290)
			p.Match(haruParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(291)
			p.ArrayLiteral()
		}
		{
			p.SetState(292)
			p.Match(haruParserT__0)
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

// ILetArrayDeclContext is an interface to support dynamic dispatch.
type ILetArrayDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLetArrayDeclContext differentiates from other interfaces.
	IsLetArrayDeclContext()
}

type LetArrayDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLetArrayDeclContext() *LetArrayDeclContext {
	var p = new(LetArrayDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_letArrayDecl
	return p
}

func InitEmptyLetArrayDeclContext(p *LetArrayDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_letArrayDecl
}

func (*LetArrayDeclContext) IsLetArrayDeclContext() {}

func NewLetArrayDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LetArrayDeclContext {
	var p = new(LetArrayDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_letArrayDecl

	return p
}

func (s *LetArrayDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *LetArrayDeclContext) CopyAll(ctx *LetArrayDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LetArrayDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetArrayDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LetImplicitArrayDeclContext struct {
	LetArrayDeclContext
}

func NewLetImplicitArrayDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetImplicitArrayDeclContext {
	var p = new(LetImplicitArrayDeclContext)

	InitEmptyLetArrayDeclContext(&p.LetArrayDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*LetArrayDeclContext))

	return p
}

func (s *LetImplicitArrayDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetImplicitArrayDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(haruParserID, 0)
}

func (s *LetImplicitArrayDeclContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *LetImplicitArrayDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterLetImplicitArrayDecl(s)
	}
}

func (s *LetImplicitArrayDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitLetImplicitArrayDecl(s)
	}
}

func (s *LetImplicitArrayDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitLetImplicitArrayDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type LetExplicitArrayDeclContext struct {
	LetArrayDeclContext
}

func NewLetExplicitArrayDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetExplicitArrayDeclContext {
	var p = new(LetExplicitArrayDeclContext)

	InitEmptyLetArrayDeclContext(&p.LetArrayDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*LetArrayDeclContext))

	return p
}

func (s *LetExplicitArrayDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetExplicitArrayDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(haruParserID, 0)
}

func (s *LetExplicitArrayDeclContext) ArrayType() IArrayTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayTypeContext)
}

func (s *LetExplicitArrayDeclContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *LetExplicitArrayDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterLetExplicitArrayDecl(s)
	}
}

func (s *LetExplicitArrayDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitLetExplicitArrayDecl(s)
	}
}

func (s *LetExplicitArrayDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitLetExplicitArrayDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) LetArrayDecl() (localctx ILetArrayDeclContext) {
	localctx = NewLetArrayDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, haruParserRULE_letArrayDecl)
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLetExplicitArrayDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(296)
			p.Match(haruParserT__24)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(297)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(298)
			p.Match(haruParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)
			p.ArrayType()
		}
		{
			p.SetState(300)
			p.Match(haruParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(301)
			p.ArrayLiteral()
		}
		{
			p.SetState(302)
			p.Match(haruParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewLetImplicitArrayDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(304)
			p.Match(haruParserT__24)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(305)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(306)
			p.Match(haruParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(307)
			p.ArrayLiteral()
		}
		{
			p.SetState(308)
			p.Match(haruParserT__0)
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

// IMutArrayDeclContext is an interface to support dynamic dispatch.
type IMutArrayDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMutArrayDeclContext differentiates from other interfaces.
	IsMutArrayDeclContext()
}

type MutArrayDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMutArrayDeclContext() *MutArrayDeclContext {
	var p = new(MutArrayDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_mutArrayDecl
	return p
}

func InitEmptyMutArrayDeclContext(p *MutArrayDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_mutArrayDecl
}

func (*MutArrayDeclContext) IsMutArrayDeclContext() {}

func NewMutArrayDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MutArrayDeclContext {
	var p = new(MutArrayDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_mutArrayDecl

	return p
}

func (s *MutArrayDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *MutArrayDeclContext) CopyAll(ctx *MutArrayDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *MutArrayDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutArrayDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MutArrayExplicitNoInitContext struct {
	MutArrayDeclContext
}

func NewMutArrayExplicitNoInitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MutArrayExplicitNoInitContext {
	var p = new(MutArrayExplicitNoInitContext)

	InitEmptyMutArrayDeclContext(&p.MutArrayDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*MutArrayDeclContext))

	return p
}

func (s *MutArrayExplicitNoInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutArrayExplicitNoInitContext) ID() antlr.TerminalNode {
	return s.GetToken(haruParserID, 0)
}

func (s *MutArrayExplicitNoInitContext) ArrayType() IArrayTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayTypeContext)
}

func (s *MutArrayExplicitNoInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterMutArrayExplicitNoInit(s)
	}
}

func (s *MutArrayExplicitNoInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitMutArrayExplicitNoInit(s)
	}
}

func (s *MutArrayExplicitNoInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitMutArrayExplicitNoInit(s)

	default:
		return t.VisitChildren(s)
	}
}

type MutArrayExplicitWithInitContext struct {
	MutArrayDeclContext
}

func NewMutArrayExplicitWithInitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MutArrayExplicitWithInitContext {
	var p = new(MutArrayExplicitWithInitContext)

	InitEmptyMutArrayDeclContext(&p.MutArrayDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*MutArrayDeclContext))

	return p
}

func (s *MutArrayExplicitWithInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutArrayExplicitWithInitContext) ID() antlr.TerminalNode {
	return s.GetToken(haruParserID, 0)
}

func (s *MutArrayExplicitWithInitContext) ArrayType() IArrayTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayTypeContext)
}

func (s *MutArrayExplicitWithInitContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *MutArrayExplicitWithInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterMutArrayExplicitWithInit(s)
	}
}

func (s *MutArrayExplicitWithInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitMutArrayExplicitWithInit(s)
	}
}

func (s *MutArrayExplicitWithInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitMutArrayExplicitWithInit(s)

	default:
		return t.VisitChildren(s)
	}
}

type MutArrayImplicitContext struct {
	MutArrayDeclContext
}

func NewMutArrayImplicitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MutArrayImplicitContext {
	var p = new(MutArrayImplicitContext)

	InitEmptyMutArrayDeclContext(&p.MutArrayDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*MutArrayDeclContext))

	return p
}

func (s *MutArrayImplicitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutArrayImplicitContext) ID() antlr.TerminalNode {
	return s.GetToken(haruParserID, 0)
}

func (s *MutArrayImplicitContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *MutArrayImplicitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterMutArrayImplicit(s)
	}
}

func (s *MutArrayImplicitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitMutArrayImplicit(s)
	}
}

func (s *MutArrayImplicitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitMutArrayImplicit(s)

	default:
		return t.VisitChildren(s)
	}
}

type MutFixedArrayNoInitContext struct {
	MutArrayDeclContext
}

func NewMutFixedArrayNoInitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MutFixedArrayNoInitContext {
	var p = new(MutFixedArrayNoInitContext)

	InitEmptyMutArrayDeclContext(&p.MutArrayDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*MutArrayDeclContext))

	return p
}

func (s *MutFixedArrayNoInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutFixedArrayNoInitContext) ID() antlr.TerminalNode {
	return s.GetToken(haruParserID, 0)
}

func (s *MutFixedArrayNoInitContext) FixedArrayType() IFixedArrayTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFixedArrayTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFixedArrayTypeContext)
}

func (s *MutFixedArrayNoInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterMutFixedArrayNoInit(s)
	}
}

func (s *MutFixedArrayNoInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitMutFixedArrayNoInit(s)
	}
}

func (s *MutFixedArrayNoInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitMutFixedArrayNoInit(s)

	default:
		return t.VisitChildren(s)
	}
}

type MutFixedArrayWithInitContext struct {
	MutArrayDeclContext
}

func NewMutFixedArrayWithInitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MutFixedArrayWithInitContext {
	var p = new(MutFixedArrayWithInitContext)

	InitEmptyMutArrayDeclContext(&p.MutArrayDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*MutArrayDeclContext))

	return p
}

func (s *MutFixedArrayWithInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutFixedArrayWithInitContext) ID() antlr.TerminalNode {
	return s.GetToken(haruParserID, 0)
}

func (s *MutFixedArrayWithInitContext) FixedArrayType() IFixedArrayTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFixedArrayTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFixedArrayTypeContext)
}

func (s *MutFixedArrayWithInitContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *MutFixedArrayWithInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterMutFixedArrayWithInit(s)
	}
}

func (s *MutFixedArrayWithInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitMutFixedArrayWithInit(s)
	}
}

func (s *MutFixedArrayWithInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitMutFixedArrayWithInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) MutArrayDecl() (localctx IMutArrayDeclContext) {
	localctx = NewMutArrayDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, haruParserRULE_mutArrayDecl)
	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMutFixedArrayWithInitContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(312)
			p.Match(haruParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(313)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(314)
			p.Match(haruParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(315)
			p.FixedArrayType()
		}
		{
			p.SetState(316)
			p.Match(haruParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(317)
			p.ArrayLiteral()
		}
		{
			p.SetState(318)
			p.Match(haruParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewMutFixedArrayNoInitContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(320)
			p.Match(haruParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(321)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(322)
			p.Match(haruParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)
			p.FixedArrayType()
		}
		{
			p.SetState(324)
			p.Match(haruParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewMutArrayExplicitWithInitContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(326)
			p.Match(haruParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(327)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.Match(haruParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)
			p.ArrayType()
		}
		{
			p.SetState(330)
			p.Match(haruParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.ArrayLiteral()
		}
		{
			p.SetState(332)
			p.Match(haruParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewMutArrayExplicitNoInitContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(334)
			p.Match(haruParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)
			p.Match(haruParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(337)
			p.ArrayType()
		}
		{
			p.SetState(338)
			p.Match(haruParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewMutArrayImplicitContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(340)
			p.Match(haruParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
			p.Match(haruParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(342)
			p.Match(haruParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			p.ArrayLiteral()
		}
		{
			p.SetState(344)
			p.Match(haruParserT__0)
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

// IArrayTypeContext is an interface to support dynamic dispatch.
type IArrayTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext

	// IsArrayTypeContext differentiates from other interfaces.
	IsArrayTypeContext()
}

type ArrayTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayTypeContext() *ArrayTypeContext {
	var p = new(ArrayTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_arrayType
	return p
}

func InitEmptyArrayTypeContext(p *ArrayTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_arrayType
}

func (*ArrayTypeContext) IsArrayTypeContext() {}

func NewArrayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayTypeContext {
	var p = new(ArrayTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_arrayType

	return p
}

func (s *ArrayTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayTypeContext) Type_() ITypeContext {
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

func (s *ArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterArrayType(s)
	}
}

func (s *ArrayTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitArrayType(s)
	}
}

func (s *ArrayTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitArrayType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) ArrayType() (localctx IArrayTypeContext) {
	localctx = NewArrayTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, haruParserRULE_arrayType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(348)
		p.Match(haruParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(349)
		p.Match(haruParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(350)
		p.Type_()
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

// IFixedArrayTypeContext is an interface to support dynamic dispatch.
type IFixedArrayTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	Type_() ITypeContext

	// IsFixedArrayTypeContext differentiates from other interfaces.
	IsFixedArrayTypeContext()
}

type FixedArrayTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFixedArrayTypeContext() *FixedArrayTypeContext {
	var p = new(FixedArrayTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_fixedArrayType
	return p
}

func InitEmptyFixedArrayTypeContext(p *FixedArrayTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_fixedArrayType
}

func (*FixedArrayTypeContext) IsFixedArrayTypeContext() {}

func NewFixedArrayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FixedArrayTypeContext {
	var p = new(FixedArrayTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_fixedArrayType

	return p
}

func (s *FixedArrayTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FixedArrayTypeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(haruParserNUMBER, 0)
}

func (s *FixedArrayTypeContext) Type_() ITypeContext {
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

func (s *FixedArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FixedArrayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FixedArrayTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterFixedArrayType(s)
	}
}

func (s *FixedArrayTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitFixedArrayType(s)
	}
}

func (s *FixedArrayTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitFixedArrayType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) FixedArrayType() (localctx IFixedArrayTypeContext) {
	localctx = NewFixedArrayTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, haruParserRULE_fixedArrayType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(352)
		p.Match(haruParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(353)
		p.Match(haruParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(354)
		p.Match(haruParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(355)
		p.Type_()
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

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_arrayLiteral
	return p
}

func InitEmptyArrayLiteralContext(p *ArrayLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_arrayLiteral
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) CopyAll(ctx *ArrayLiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArrayLiteralExprListContext struct {
	ArrayLiteralContext
}

func NewArrayLiteralExprListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayLiteralExprListContext {
	var p = new(ArrayLiteralExprListContext)

	InitEmptyArrayLiteralContext(&p.ArrayLiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArrayLiteralContext))

	return p
}

func (s *ArrayLiteralExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralExprListContext) AllExpr() []IExprContext {
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

func (s *ArrayLiteralExprListContext) Expr(i int) IExprContext {
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

func (s *ArrayLiteralExprListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterArrayLiteralExprList(s)
	}
}

func (s *ArrayLiteralExprListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitArrayLiteralExprList(s)
	}
}

func (s *ArrayLiteralExprListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitArrayLiteralExprList(s)

	default:
		return t.VisitChildren(s)
	}
}

type EmptyArrContext struct {
	ArrayLiteralContext
}

func NewEmptyArrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmptyArrContext {
	var p = new(EmptyArrContext)

	InitEmptyArrayLiteralContext(&p.ArrayLiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArrayLiteralContext))

	return p
}

func (s *EmptyArrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyArrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterEmptyArr(s)
	}
}

func (s *EmptyArrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitEmptyArr(s)
	}
}

func (s *EmptyArrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitEmptyArr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, haruParserRULE_arrayLiteral)
	var _alt int

	p.SetState(370)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		localctx = NewArrayLiteralExprListContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(357)
			p.Match(haruParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(358)
			p.expr(0)
		}
		p.SetState(363)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1+1 {
				{
					p.SetState(359)
					p.Match(haruParserT__47)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(360)
					p.expr(0)
				}

			}
			p.SetState(365)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(366)
			p.Match(haruParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewEmptyArrContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(368)
			p.Match(haruParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(369)
			p.Match(haruParserT__20)
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

// IArrayItemAssignContext is an interface to support dynamic dispatch.
type IArrayItemAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArrayItemAssignContext differentiates from other interfaces.
	IsArrayItemAssignContext()
}

type ArrayItemAssignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayItemAssignContext() *ArrayItemAssignContext {
	var p = new(ArrayItemAssignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_arrayItemAssign
	return p
}

func InitEmptyArrayItemAssignContext(p *ArrayItemAssignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_arrayItemAssign
}

func (*ArrayItemAssignContext) IsArrayItemAssignContext() {}

func NewArrayItemAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayItemAssignContext {
	var p = new(ArrayItemAssignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_arrayItemAssign

	return p
}

func (s *ArrayItemAssignContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayItemAssignContext) CopyAll(ctx *ArrayItemAssignContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ArrayItemAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayItemAssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArrayIndexAssignContext struct {
	ArrayItemAssignContext
}

func NewArrayIndexAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayIndexAssignContext {
	var p = new(ArrayIndexAssignContext)

	InitEmptyArrayItemAssignContext(&p.ArrayItemAssignContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArrayItemAssignContext))

	return p
}

func (s *ArrayIndexAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayIndexAssignContext) ID() antlr.TerminalNode {
	return s.GetToken(haruParserID, 0)
}

func (s *ArrayIndexAssignContext) AllExpr() []IExprContext {
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

func (s *ArrayIndexAssignContext) Expr(i int) IExprContext {
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

func (s *ArrayIndexAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterArrayIndexAssign(s)
	}
}

func (s *ArrayIndexAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitArrayIndexAssign(s)
	}
}

func (s *ArrayIndexAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitArrayIndexAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) ArrayItemAssign() (localctx IArrayItemAssignContext) {
	localctx = NewArrayItemAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, haruParserRULE_arrayItemAssign)
	localctx = NewArrayIndexAssignContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)
		p.Match(haruParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(373)
		p.Match(haruParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(374)
		p.expr(0)
	}
	{
		p.SetState(375)
		p.Match(haruParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(376)
		p.Match(haruParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(377)
		p.expr(0)
	}
	{
		p.SetState(378)
		p.Match(haruParserT__0)
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

// IArrayReassignContext is an interface to support dynamic dispatch.
type IArrayReassignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	ArrayLiteral() IArrayLiteralContext

	// IsArrayReassignContext differentiates from other interfaces.
	IsArrayReassignContext()
}

type ArrayReassignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayReassignContext() *ArrayReassignContext {
	var p = new(ArrayReassignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_arrayReassign
	return p
}

func InitEmptyArrayReassignContext(p *ArrayReassignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_arrayReassign
}

func (*ArrayReassignContext) IsArrayReassignContext() {}

func NewArrayReassignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayReassignContext {
	var p = new(ArrayReassignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_arrayReassign

	return p
}

func (s *ArrayReassignContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayReassignContext) ID() antlr.TerminalNode {
	return s.GetToken(haruParserID, 0)
}

func (s *ArrayReassignContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *ArrayReassignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayReassignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayReassignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterArrayReassign(s)
	}
}

func (s *ArrayReassignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitArrayReassign(s)
	}
}

func (s *ArrayReassignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitArrayReassign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) ArrayReassign() (localctx IArrayReassignContext) {
	localctx = NewArrayReassignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, haruParserRULE_arrayReassign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(380)
		p.Match(haruParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(381)
		p.Match(haruParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(382)
		p.ArrayLiteral()
	}
	{
		p.SetState(383)
		p.Match(haruParserT__0)
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

// ILenFunctionContext is an interface to support dynamic dispatch.
type ILenFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsLenFunctionContext differentiates from other interfaces.
	IsLenFunctionContext()
}

type LenFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLenFunctionContext() *LenFunctionContext {
	var p = new(LenFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_lenFunction
	return p
}

func InitEmptyLenFunctionContext(p *LenFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_lenFunction
}

func (*LenFunctionContext) IsLenFunctionContext() {}

func NewLenFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LenFunctionContext {
	var p = new(LenFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_lenFunction

	return p
}

func (s *LenFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *LenFunctionContext) Expr() IExprContext {
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

func (s *LenFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LenFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LenFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterLenFunction(s)
	}
}

func (s *LenFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitLenFunction(s)
	}
}

func (s *LenFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitLenFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) LenFunction() (localctx ILenFunctionContext) {
	localctx = NewLenFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, haruParserRULE_lenFunction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(385)
		p.Match(haruParserT__48)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(386)
		p.Match(haruParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(387)
		p.expr(0)
	}
	{
		p.SetState(388)
		p.Match(haruParserT__4)
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

// IFunctionDeclContext is an interface to support dynamic dispatch.
type IFunctionDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Block() IBlockContext
	ParamList() IParamListContext
	ReturnSignature() IReturnSignatureContext

	// IsFunctionDeclContext differentiates from other interfaces.
	IsFunctionDeclContext()
}

type FunctionDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclContext() *FunctionDeclContext {
	var p = new(FunctionDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_functionDecl
	return p
}

func InitEmptyFunctionDeclContext(p *FunctionDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_functionDecl
}

func (*FunctionDeclContext) IsFunctionDeclContext() {}

func NewFunctionDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclContext {
	var p = new(FunctionDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_functionDecl

	return p
}

func (s *FunctionDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(haruParserID, 0)
}

func (s *FunctionDeclContext) Block() IBlockContext {
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

func (s *FunctionDeclContext) ParamList() IParamListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *FunctionDeclContext) ReturnSignature() IReturnSignatureContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnSignatureContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnSignatureContext)
}

func (s *FunctionDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterFunctionDecl(s)
	}
}

func (s *FunctionDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitFunctionDecl(s)
	}
}

func (s *FunctionDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitFunctionDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) FunctionDecl() (localctx IFunctionDeclContext) {
	localctx = NewFunctionDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, haruParserRULE_functionDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		p.Match(haruParserT__49)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(391)
		p.Match(haruParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(392)
		p.Match(haruParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == haruParserID {
		{
			p.SetState(393)
			p.ParamList()
		}

	}
	{
		p.SetState(396)
		p.Match(haruParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(398)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == haruParserT__11 {
		{
			p.SetState(397)
			p.ReturnSignature()
		}

	}
	{
		p.SetState(400)
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

// IParamListContext is an interface to support dynamic dispatch.
type IParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParam() []IParamContext
	Param(i int) IParamContext

	// IsParamListContext differentiates from other interfaces.
	IsParamListContext()
}

type ParamListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamListContext() *ParamListContext {
	var p = new(ParamListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_paramList
	return p
}

func InitEmptyParamListContext(p *ParamListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_paramList
}

func (*ParamListContext) IsParamListContext() {}

func NewParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamListContext {
	var p = new(ParamListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_paramList

	return p
}

func (s *ParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamListContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *ParamListContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
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

	return t.(IParamContext)
}

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterParamList(s)
	}
}

func (s *ParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitParamList(s)
	}
}

func (s *ParamListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitParamList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) ParamList() (localctx IParamListContext) {
	localctx = NewParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, haruParserRULE_paramList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(402)
		p.Param()
	}
	p.SetState(407)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == haruParserT__47 {
		{
			p.SetState(403)
			p.Match(haruParserT__47)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(404)
			p.Param()
		}

		p.SetState(409)
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

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Type_() ITypeContext

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_param
	return p
}

func InitEmptyParamContext(p *ParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_param
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) ID() antlr.TerminalNode {
	return s.GetToken(haruParserID, 0)
}

func (s *ParamContext) Type_() ITypeContext {
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

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitParam(s)
	}
}

func (s *ParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, haruParserRULE_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(410)
		p.Match(haruParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(411)
		p.Match(haruParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(412)
		p.Type_()
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

// IReturnSignatureContext is an interface to support dynamic dispatch.
type IReturnSignatureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllType_() []ITypeContext
	Type_(i int) ITypeContext

	// IsReturnSignatureContext differentiates from other interfaces.
	IsReturnSignatureContext()
}

type ReturnSignatureContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnSignatureContext() *ReturnSignatureContext {
	var p = new(ReturnSignatureContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_returnSignature
	return p
}

func InitEmptyReturnSignatureContext(p *ReturnSignatureContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_returnSignature
}

func (*ReturnSignatureContext) IsReturnSignatureContext() {}

func NewReturnSignatureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnSignatureContext {
	var p = new(ReturnSignatureContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_returnSignature

	return p
}

func (s *ReturnSignatureContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnSignatureContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *ReturnSignatureContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
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

	return t.(ITypeContext)
}

func (s *ReturnSignatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnSignatureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnSignatureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterReturnSignature(s)
	}
}

func (s *ReturnSignatureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitReturnSignature(s)
	}
}

func (s *ReturnSignatureContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitReturnSignature(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) ReturnSignature() (localctx IReturnSignatureContext) {
	localctx = NewReturnSignatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, haruParserRULE_returnSignature)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(414)
		p.Match(haruParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(415)
		p.Type_()
	}
	p.SetState(420)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == haruParserT__47 {
		{
			p.SetState(416)
			p.Match(haruParserT__47)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(417)
			p.Type_()
		}

		p.SetState(422)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(423)
		p.Match(haruParserT__12)
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

// IReturnStmtContext is an interface to support dynamic dispatch.
type IReturnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExprList() IExprListContext

	// IsReturnStmtContext differentiates from other interfaces.
	IsReturnStmtContext()
}

type ReturnStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStmtContext() *ReturnStmtContext {
	var p = new(ReturnStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_returnStmt
	return p
}

func InitEmptyReturnStmtContext(p *ReturnStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_returnStmt
}

func (*ReturnStmtContext) IsReturnStmtContext() {}

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_returnStmt

	return p
}

func (s *ReturnStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStmtContext) ExprList() IExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) ReturnStmt() (localctx IReturnStmtContext) {
	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, haruParserRULE_returnStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(425)
		p.Match(haruParserT__50)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&559009303772530712) != 0 {
		{
			p.SetState(426)
			p.ExprList()
		}

	}
	{
		p.SetState(429)
		p.Match(haruParserT__0)
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

// IExprListContext is an interface to support dynamic dispatch.
type IExprListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsExprListContext differentiates from other interfaces.
	IsExprListContext()
}

type ExprListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprListContext() *ExprListContext {
	var p = new(ExprListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_exprList
	return p
}

func InitEmptyExprListContext(p *ExprListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_exprList
}

func (*ExprListContext) IsExprListContext() {}

func NewExprListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprListContext {
	var p = new(ExprListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_exprList

	return p
}

func (s *ExprListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprListContext) AllExpr() []IExprContext {
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

func (s *ExprListContext) Expr(i int) IExprContext {
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

func (s *ExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterExprList(s)
	}
}

func (s *ExprListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitExprList(s)
	}
}

func (s *ExprListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitExprList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) ExprList() (localctx IExprListContext) {
	localctx = NewExprListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, haruParserRULE_exprList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(431)
		p.expr(0)
	}
	p.SetState(436)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == haruParserT__47 {
		{
			p.SetState(432)
			p.Match(haruParserT__47)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(433)
			p.expr(0)
		}

		p.SetState(438)
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

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	ArgumentList() IArgumentListContext

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) ID() antlr.TerminalNode {
	return s.GetToken(haruParserID, 0)
}

func (s *FunctionCallContext) ArgumentList() IArgumentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, haruParserRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(439)
		p.Match(haruParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(440)
		p.Match(haruParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(442)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&559009303772530712) != 0 {
		{
			p.SetState(441)
			p.ArgumentList()
		}

	}
	{
		p.SetState(444)
		p.Match(haruParserT__4)
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

// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_argumentList
	return p
}

func InitEmptyArgumentListContext(p *ArgumentListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_argumentList
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) AllExpr() []IExprContext {
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

func (s *ArgumentListContext) Expr(i int) IExprContext {
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

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterArgumentList(s)
	}
}

func (s *ArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitArgumentList(s)
	}
}

func (s *ArgumentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitArgumentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) ArgumentList() (localctx IArgumentListContext) {
	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, haruParserRULE_argumentList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(446)
		p.expr(0)
	}
	p.SetState(451)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == haruParserT__47 {
		{
			p.SetState(447)
			p.Match(haruParserT__47)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(448)
			p.expr(0)
		}

		p.SetState(453)
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

// IWhileLoopContext is an interface to support dynamic dispatch.
type IWhileLoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Block() IBlockContext

	// IsWhileLoopContext differentiates from other interfaces.
	IsWhileLoopContext()
}

type WhileLoopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileLoopContext() *WhileLoopContext {
	var p = new(WhileLoopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_whileLoop
	return p
}

func InitEmptyWhileLoopContext(p *WhileLoopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = haruParserRULE_whileLoop
}

func (*WhileLoopContext) IsWhileLoopContext() {}

func NewWhileLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileLoopContext {
	var p = new(WhileLoopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = haruParserRULE_whileLoop

	return p
}

func (s *WhileLoopContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileLoopContext) Expr() IExprContext {
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

func (s *WhileLoopContext) Block() IBlockContext {
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

func (s *WhileLoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileLoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileLoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.EnterWhileLoop(s)
	}
}

func (s *WhileLoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(haruListener); ok {
		listenerT.ExitWhileLoop(s)
	}
}

func (s *WhileLoopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case haruVisitor:
		return t.VisitWhileLoop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *haruParser) WhileLoop() (localctx IWhileLoopContext) {
	localctx = NewWhileLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, haruParserRULE_whileLoop)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(454)
		p.Match(haruParserT__51)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(455)
		p.Match(haruParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(456)
		p.expr(0)
	}
	{
		p.SetState(457)
		p.Match(haruParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(458)
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
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
