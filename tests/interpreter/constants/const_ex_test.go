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

func TestConstantsExplicitConstants(t *testing.T) {
	input := `
	--- EXPLICITE CONSTANT DECLARATIONS
	print "EXPLICITE CONSTANT DECLARATIONS";
	const a:i8 = 1;
	const b: f64 = 3.14; --- 3.140000
	const c: bool = true;
	const d: string = "hello";

	print a;
	print b;
	print c;
	print d;
	print a + 4 * 2; --- expression with constant

	const e: bool = false;
	const f: bool = !e;
	const g: bool = !e == true && (e != f); --- true
	const z: string = d; --- assigning value from other constant
	--- z = "new value"; --- will stop execution from this line

	print e;
	print f;
	print e && f;
	print e || f;
	print g;
	print z;
`

	expected := `EXPLICITE CONSTANT DECLARATIONS
1
3.140000
true
hello
9
false
true
false
true
true
hello
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
