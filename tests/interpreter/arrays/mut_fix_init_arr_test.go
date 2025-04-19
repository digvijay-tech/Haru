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

func TestFixedInitMutArrays(t *testing.T) {
	input := `
	--- FIXED INITIALIZED MUT ARRAYS DECLARATIONS
	print "FIXED INITIALIZED MUT ARRAYS DECLARATIONS";
	mut nums: [6]ui16 = [1,2,3,4,5]; 		--- half initialized
	mut chars: [2]byte = ["a"];      		--- half initialized
	mut bools: [4]bool = [];		 		--- empty
	mut strs: [2]string = ["One", "Two"]; 	--- fully initialized

	print nums;
	print chars;
	print bools;
	print strs;

	print nums[nums[4]-nums[3]]; --- 2
`

	expected := `FIXED INITIALIZED MUT ARRAYS DECLARATIONS
[1,2,3,4,5,0]
[97,0]
[false,false,false,false]
[One,Two]
2
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
