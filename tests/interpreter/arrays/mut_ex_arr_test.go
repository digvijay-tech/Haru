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

func TestExplicitMutArrays(t *testing.T) {
	input := `
	--- DYNAMIC EXPLICIT MUT ARRAYS DECLARATIONS
	print "DYNAMIC EXPLICIT MUT ARRAYS DECLARATIONS";
	mut n: []int = [5, 10, 15, 32, 33];
	mut m: []f32 = [1.5, 41.2, 0];
	mut o: []bool = [true, false];
	mut b: []byte = ["A", "B", "C"];
	mut s: []string = ["One", "Two"];

	print n;
	print m;
	print o;
	print b;
	print s;

	mut emp: []ui8 = [];
	print emp;
	--- print emp[0];	--- Runtime Error: array emp is empty
`

	expected := `DYNAMIC EXPLICIT MUT ARRAYS DECLARATIONS
[5,10,15,32,33]
[1.500000,41.200001,0.000000]
[true,false]
[65,66,67]
["One","Two"]
[]
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
