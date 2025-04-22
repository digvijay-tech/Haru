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

func TestPointerDeclarations(t *testing.T) {
	input := `
		--- TESTING IMMUTABLE POINTER DECLARATIONS (READ-ONLY)
		print "TESTING IMMUTABLE POINTER DECLARATIONS (READ-ONLY)";
		let x: i32 = 10;
		let xptr: *i32 = &x;

		let y: string = 15;
		let yptr: *string = &y;

		print xptr;
		print yptr;
		
		print "==============";
		--- TESTING MUTABLE POINTER DECLARATIONS (READ AND WRITE)
		print "TESTING MUTABLE POINTER DECLARATIONS (READ AND WRITE)";
		mut a: f32 = 137.137;
		mut aptr: *f32 = &a;

		mut b: bool = true;
		mut bptr: *bool = &b;

		print aptr;
		print bptr;

		--- 
		let c: byte = 65;		--- read only/immutable variable
		
		--- assigning to mutable/read and write pointer
		--- mut cptr: *byte = &c;	--- Runtime Error: cannot create mutable pointer to immutable variable 'c'

`

	expected := `TESTING IMMUTABLE POINTER DECLARATIONS (READ-ONLY)
x
y
==============
TESTING MUTABLE POINTER DECLARATIONS (READ AND WRITE)
a
b
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
