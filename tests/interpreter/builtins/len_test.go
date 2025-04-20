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

func TestBuiltInLenFunc(t *testing.T) {
	input := `
	--- BUILT-IN LEN FUNCTION
	print "BUILT-IN LEN FUNCTION";
	let name = "Haru";
	let nums: []int = [1,2,3,4,5,6];
	mut empty: []byte = [];
	mut days = ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"];

	print len(name);		--- 4
	print len("Haru!");		--- 5
	print len(nums);		--- 6
	print len(empty);		--- 0
	print len(days);		--- 7

	--- storing values
	let nameCount = len(name);
	print nameCount;		--- 4

	let numCount = len(nums);
	print numCount;			--- 6

	let zero: i8 = len(empty);
	print zero;				--- 0

	const dayCount: int = len(days);
	print dayCount;			--- 7
`

	expected := `BUILT-IN LEN FUNCTION
4
5
6
0
7
4
6
0
7
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
