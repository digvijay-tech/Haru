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

func TestImplicitMutArrays(t *testing.T) {
	input := `
	--- IMPLICIT MUT ARRAY DECLARATIONS
	print "IMPLICIT MUT ARRAY DECLARATIONS";
	mut nums = [2194781841974619, 1, -12983];
	print nums;
	print nums[5-3];

	mut booleans = [true, false, false, true];
	print booleans;

	mut floats = [0.0012, 0.12948];
	print floats;

	mut names = ["John", "Jane", "Jimmy"];
	print names;
	print names[1];

	--- mut problem = []; --- Error: array literal cannot be empty
`

	expected := `IMPLICIT MUT ARRAY DECLARATIONS
[2194781841974619,1,-12983]
-12983
[true,false,false,true]
[0.001200,0.129480]
[John,Jane,Jimmy]
Jane
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
