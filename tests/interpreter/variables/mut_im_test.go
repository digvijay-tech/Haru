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

func TestImplicitMut(t *testing.T) {
	input := `
	--- IMPLICIT MUT DECLARATIONS
	print "IMPLICIT MUT DECLARATIONS";
	mut name = "Haru";
	print name;

	mut age = 20000000;
	print age;

	mut x = 52;
	print (x + 11) * ( 5 / 2.4) * 12.7 / 7 / 35;

	mut a = 5;
	print a * (7.9 + 1) + (-3.43) / (7/2);
	
	a = 10;
	print a;

	a = a+1;
	print a;

	mut isLarge = a > 10;
	print isLarge;

	isLarge = !(10000 >= 0);
	print isLarge;

	mut q: f32;
	print q;

	q = 32.12;
	print q;
`

	expected := `IMPLICIT MUT DECLARATIONS
Haru
20000000
6.803571428571429
43.52
10
11
true
false
0
32.119999
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
