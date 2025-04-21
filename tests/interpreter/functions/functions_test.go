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

func TestFunctions(t *testing.T) {
	input := `
		--- TESTING FUNCTIONS
		print "TESTING FUNCTIONS";

		function test(name: string, age: ui8, isEngineer: bool, initial: byte) {
			print "From Test Function!";
			print name + "'s " + "details:";
			print age;

			if (!isEngineer) {
				print "Occupation: Not Engineer";
			} else {
			 	print "Occupation: Engineer";
			}

			print "Initial in Byte:";
			print initial;
		}

		--- function call as statement
		test("John", 20, true, "J"); --- needs semicolon

		--- print isEngineer; --- Runtime Error: undefined variable 'isEngineer'

		print "===============";

		function test2() {
			print "I am second test!";
		}

		test2();

		print "===============";
		--- function call as expression

		function test3() <bool> {
			return false;
		}

		const x = test3();
		print x; --- false


		print "===============";

		function test4(x: i8, y: i8) <i32> {
			print "inside";
			return x + y;
			print "this won't run";  --- skipped
		}

		const add1: i32 = test4(10, 25);
		print add1;


		--- using return for early termination
		function test5(x: string, y: string) {
			const xLen: uint = len(x);
			const yLen: uint = len(y);

			if (xLen > yLen) {
				print "first argument has more characters!";
			} else {
				print "second argument has more characters!"; 
			}
		}

		test5("John", "Jo");
`

	expected := `TESTING FUNCTIONS
From Test Function!
John's details:
20
Occupation: Engineer
Initial in Byte:
74
===============
I am second test!
===============
false
===============
inside
35
first argument has more characters!
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
