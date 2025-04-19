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

func TestFixedUninitMutArrays(t *testing.T) {
	input := `
	--- FIXED UNINITIALIZED MUT ARRAYS DECLARATIONS
	print "FIXED UNINITIALIZED MUT ARRAYS DECLARATIONS";

	--- fixed arrays when uninitialised will user zero value as a placeholder untill new value is assigned
	mut ranks: [4]string;
	mut integers: [3]int;
	mut floats: [1]f32;
	mut bools: [10]bool;
	mut bytes: [100]byte;

	print ranks;
	print integers;
	print floats;
	print bools;
	print bytes;
	print bytes[10];
`

	expected := `FIXED UNINITIALIZED MUT ARRAYS DECLARATIONS
[,,,]
[0,0,0]
[0]
[false,false,false,false,false,false,false,false,false,false]
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
0
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
