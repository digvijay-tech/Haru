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

func TestImplicitLet(t *testing.T) {
	input := `
	--- IMPLICIT LET DECLARATIONS
	print "IMPLICITE LET DECLARATIONS";	
	let a = 5;								  --- 5
	let b = 2.5;							  --- 2.500000
	let sum = a + b;						  --- 7.500000
	let status = true;						  --- true
	let neg = !status;						  --- false
	let word = "Hi";						  --- Hi
	let exclaim = word + "!";				  --- Hi!
	let bigCalc = (2 + 3) * 4;				  --- 20
	let floatResult = 2.0 * (3 + 4);		  --- 14.000000
	let condition = a < 10;					  --- true
	let logic = (a < 10) && (b < 10);		  --- true
	let deep = ((a + 2) * (b + 1.5)) - 3;     --- (7 * 4.0) - 3 = 25.000000
	let flags = !false && status || false;    --- true
	let equality = 10 == 10.0;                --- true
	let inequality = sum != 7.5;              --- false
	let description = word + " " + "there!";  --- Hi there!
	let combo = a + 2 > b * 3;                --- 7 > 7.5 = false
	--- let mixed = "value: " + sum;		  --- Error: cannot operate on types 'string' and 'f32'

	print a;							--- 5
	print b;							--- 2.500000
	print sum;							--- 7.500000
	print status;						--- true
	print neg;							--- false
	print word;							--- Hi
	print exclaim;						--- Hi!
	print bigCalc;						--- 20
	print floatResult;					--- 14.000000
	print condition;					--- true
	print logic;						--- true
	print deep;							--- 25.0
	print flags;						--- true
	print equality;						--- true
	print inequality;					--- false
	print description;					--- Hi there!
	print combo;						--- false
`

	expected := `IMPLICITE LET DECLARATIONS
5
2.500000
7.500000
true
false
Hi
Hi!
20
14.000000
true
true
25.000000
true
true
false
Hi there!
false
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
