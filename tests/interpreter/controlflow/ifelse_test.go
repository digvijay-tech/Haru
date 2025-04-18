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

func TestIfElseStatements(t *testing.T) {
	input := `
		--- IF CONDITIONS ONLY
		print "TESTING IF CONDITIONS ONLY";

		mut isActive: bool;
		print isActive;

		if (!isActive) {
			print "not active";
		}

		isActive = 10 >= (5 * 1.5);
		if (isActive && 1 > 0.10) {
			print (5 * 1.5);
		}
		
		let name = "John Doe";
		print name;

		let age: ui8 = 15;
		print age;

		if (name == "John Doe" && age >= 20) {
			print name + " can drive";
		}

		if (name == "John Doe" && age <= 20) {
			print name + " cannot drive";
		}

		if (name != "") {
			print "Name is not empty!";
			if (name == "John Doe") {
				print "Name is " + name;

				if (age >= 20) {
					print name + " is old enough to drive!";
				}

				if (age <= 20) {
					print name + " is not old enough to drive!";
				}
			}
		}

		--- IF/ELSE CHAIN
`

	expected := `TESTING IF CONDITIONS ONLY
false
not active
7.5
John Doe
15
John Doe cannot drive
Name is not empty!
Name is John Doe
John Doe is not old enough to drive!
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
