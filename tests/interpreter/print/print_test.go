package interpreter_test

import (
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/digvijay-tech/Haru/internal/interpreter"
	"github.com/digvijay-tech/Haru/internal/parser"
	"github.com/digvijay-tech/Haru/internal/preprocessor"
	utils_test "github.com/digvijay-tech/Haru/tests/interpreter/utils"
)

func TestPrintLiterals(t *testing.T) {
	input := `
	--- Comment: testing print statement with basic literal values
	print 0;
	print 42;


		print 3.14; --- 3.140000
	print -137.9928; --- -137.992798
	print "hello";

	print 'hello';
	print " Hello World! "; --- test comment
	print true;
	print false;
	print 0b1010;
--- print "This message is ignored";

	--- TESTING PRINTING CONSTANTS
	const name: string = "Digvijaysinh Padhiyar";
	const isProgrammer: bool = !false;
	const iq: f32 = -50.0; --- -50.000000
	print name;
	print isProgrammer;
	print iq;
`

	expected := `0
42
3.140000
-137.992798
hello
hello
 Hello World! 
true
false
0b1010
Digvijaysinh Padhiyar
true
-50.000000
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
		visitor := interpreter.NewHaruVisitor()
		visitor.Visit(tree)
	})

	// normalizing out for Windows
	output = strings.ReplaceAll(output, "\r\n", "\n")

	if output != expected {
		t.Errorf("\nExpected:\n%s\nGot:\n%s", expected, output)
	}
}
