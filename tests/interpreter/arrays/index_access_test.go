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

func TestIndexBasedAccess(t *testing.T) {
	input := `
	--- INDEX BASED ACCESS
	print "INDEX BASED ACCESS";

	const a: []i8 = [8, 7, 6, 5, 4, 3, 2, 1];
	print a;
	mut b: ui8 = 0;
	print a[b];

	const colors: []string = ['Orange', "Blue", 'Red', "Green", "Teal"];
	print colors;
	print colors[b+1*2]; --- Red
	--- print colors[5]; --- Runtime Error: index 5 out of bounds

	const secondColor: string = colors[1]; --- assigning Blue to this constant
	print secondColor;


	--- print y[0]; --- Runtime Error: undefined array 'y'
`

	expected := `INDEX BASED ACCESS
[8,7,6,5,4,3,2,1]
8
["Orange","Blue","Red","Green","Teal"]
Red
Blue
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
