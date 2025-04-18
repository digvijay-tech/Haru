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

func TestExplicitMut(t *testing.T) {
	input := `
	--- EXPLICITE MUT DECLARATIONS
	print "EXPLICITE MUT DECLARATIONS";
	mut a: i8 = 5;
	mut b: i8;
	mut c: bool;
	mut d: bool = true;
	mut e: string;
	mut name: string = "Haru";
	mut greet: string = "Hello" + ' ' + name + "!";
	mut x: f32 = a * (7.9 + 1) + (-3.43) / (7/2);

	print a + 5.8;	--- 10.8
	print b;		--- 0
	print c;		--- false
	print d;		--- true
	print e;		--- (empty string)
	print name;		--- Haru!
	print greet;	--- Hello Haru!
	print x;		--- 43.520000
`

	expected := `EXPLICITE MUT DECLARATIONS
10.8
0
false
true

Haru
Hello Haru!
43.520000
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
