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

func TestPointerDereference(t *testing.T) {
	input := `
		--- TESTING POINTER DEREFERENCE
		print "TESTING POINTER DEREFERENCE";

		let x: i8 = 10;
		let xptr: *i8 = &x;

		mut y: bool = false;
		mut yptr: *bool = &y;

		print xptr;
		print yptr;

		print *xptr;
		print *yptr;

		--- *xptr = 0; --- Runtime Error: pointer 'xptr' is immutable

		--- TESTING REASSIGNMENT
		print "====================";
		print "TESTING REASSIGNMENT";

		mut a: string = "Haru";
		mut aptr: *string = &a;

		--- dereferencing the pointer to get value stored in variable a
		print *aptr;

		--- mutating the value of variable a 
		*aptr = "Haru World";
		print a;

		mut b: i8 = 100;
		mut bptr: *i8 = &b;

		print *bptr;
		--- *bptr = 128; --- Runtime Error: value 128 out of range for type i8

		mut c: byte = "C";
		mut cptr: *byte = &c;

		print *cptr;
		*cptr = "D";
		print c;


		--- testing pointers for scoping issues
		if (true) {
			print "======== IN IF BLOCK ========";
			mut temp: i32 = 20;
			mut tempptr: *i32 = &temp;
			
			print *tempptr;

			*tempptr = 40;

			print temp;
			print *tempptr;
		}

		--- print temp;		--- Runtime Error: undefined variable 'temp'
		--- print *tempptr; --- Runtime Error: undefined variable 'tempptr'
`

	expected := `TESTING POINTER DEREFERENCE
x
y
10
false
====================
TESTING REASSIGNMENT
Haru
Haru World
100
67
68
======== IN IF BLOCK ========
20
40
40
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
