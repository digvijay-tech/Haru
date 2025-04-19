package interpreter_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/digvijay-tech/Haru/internal/interpreter"
	"github.com/digvijay-tech/Haru/internal/parser"
	"github.com/digvijay-tech/Haru/internal/preprocessor"
	utils_test "github.com/digvijay-tech/Haru/tests/interpreter/utils"
)

func TestExplicitLetArrays(t *testing.T) {
	input := `
	--- EXPLICITE LET ARRAYS DECLARATIONS
	print "EXPLICITE LET ARRAYS DECLARATIONS";
	let nums: []i8 = [10, 20, 30, 40, 50, 60, 70, 80, 90, 100];
	let alpha: []byte = ["a", "b", "c", "d"];
	let bools: []bool = [true, false, true, true];
	let floats: []f64 = [2.3, 2.4, 23.12341];
	let colors: []string = ['Orange', "Blue", 'Red', "Green", "Teal"];
	let uis: []uint = [1, 23, 255];

	print nums;
	print alpha;
	print bools;
	print floats;
	print colors;
	print uis;

	let red = colors[2];
	print red;
`

	expected := `EXPLICITE LET ARRAYS DECLARATIONS
[10,20,30,40,50,60,70,80,90,100]
[97,98,99,100]
[true,false,true,true]
[2.300000,2.400000,23.123409]
["Orange","Blue","Red","Green","Teal"]
[1,23,255]
Red
`

	// cleaning source input with custom preprocessor
	cleanIn := preprocessor.PreProcess(input)

	// setup ANTLR input and parsing
	is := antlr.NewInputStream(cleanIn)

	lexer := parser.NewharuLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewharuParser(stream)
	tree := p.Program()

	// capturing the output
	output := utils_test.CaptureOutput(func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()

		visitor := interpreter.NewHaruVisitor()
		visitor.Visit(tree)
	})

	// normalizing out for Windows
	output = strings.ReplaceAll(output, "\r\n", "\n")

	if output != expected {
		t.Errorf("\nExpected:\n%s\nGot:\n%s", expected, output)
	}
}
