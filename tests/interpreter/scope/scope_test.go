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

func TestVariableScope(t *testing.T) {
	input := `
	--- TESTING VARIABLE SCOPE
	print "TESTING VARIABLE SCOPE";

	let alpha: i8 = 15;
	mut beta: ui16 = 10;
	const gamma: bool = false;

	if (5 >= 4) {
		print "=== BEGIN BLOCK ===";
		let x = 5;
		print x;	--- x is local to this block

		--- accessing global scope variables in a local scope
		print alpha;
		print beta;
		print gamma;
		print "==== END BLOCK ====";
	}

	if (true) {
		print "=== BEGIN BLOCK ===";

		--- can reassign the mutable variable declared in global scope
		beta = 100;
		print beta;

		mut name: string = "Your Friendly Local Mut!";
		print name;
		print "==== END BLOCK ====";
	}

	print "beta reassigned in local scope:";
	print beta; --- now changed to 100

	--- print name; --- Runtime Error: undefined variable 'name'
	--- print x; --- Runtime Error: undefined variable 'x'


	--- SHADOWING OUTER VARIABLE
	mut one = 1;
	let One = "One";

	if (true) {
		print "=== BEGIN BLOCK ===";

		mut one = 10;
		let One: string = "Not One";

		print one;	--- 10
		print One;	--- "Not One"
		print "==== END BLOCK ====";
	}

	print one;	--- 1
	print One;	--- "One"


	--- NESTED BLOCKS
	mut two: int = 2;

	if (true) {
		print "=== BEGIN BLOCK ===";

		const three = 3;

		if (!!true) {
			let four = 4;

			print two;		--- 2
			print three;	--- 3
			print four;		--- 4
		}

		--- print four; --- Runtime Error: undefined variable 'four'

		print "==== END BLOCK ====";
	}
`

	expected := `TESTING VARIABLE SCOPE
=== BEGIN BLOCK ===
5
15
10
false
==== END BLOCK ====
=== BEGIN BLOCK ===
100
Your Friendly Local Mut!
==== END BLOCK ====
beta reassigned in local scope:
100
=== BEGIN BLOCK ===
10
Not One
==== END BLOCK ====
1
One
=== BEGIN BLOCK ===
2
3
4
==== END BLOCK ====
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
