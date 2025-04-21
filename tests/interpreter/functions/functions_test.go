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
		const x = test(); --- needs semicolon
		print x;
`

	expected := `TESTING FUNCTIONS!
John's details:
20
Occupation: Engineer
Initial in Byte:
74
===============
I am second test!
===============
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
