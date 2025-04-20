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

func TestMutArrayReassignmentWithIndex(t *testing.T) {
	input := `
	--- MUT ARRAY REASSIGNMENTS BASED ON INDEX
	print "MUT ARRAY REASSIGNMENTS BASED ON INDEX";
	mut poll: [2]bool;
	print poll;

	poll[0] = true;
	poll[1] = true;
	print poll;

	mut names: []string = ["Jack", "John", "Joma"];
	print names;

	let joma = names[2];
	let jimmy = "Jimmy";
	print "Replacing " + joma + " with " + jimmy + "!";

	names[2] = "Jimmy";

	print "New Names:";
	print names;

	--- names[0] = 19; --- Error: type mismatch cannot assign
	print names[0];

	mut nums: []i8 = [-121, -120, 0, 121, 120];
	print nums;

	nums[2] = 1;
	print nums;

	let primes: []int = [2, 3, 5, 7];
	print primes;

	--- primes[0] = 4; --- Error: cannot reassign

	print nums;
	nums[2] = primes[0];
	print nums;
`

	expected := `MUT ARRAY REASSIGNMENTS BASED ON INDEX
[false,false]
[true,true]
["Jack","John","Joma"]
Replacing Joma with Jimmy!
New Names:
["Jack","John","Jimmy"]
Jack
[-121,-120,0,121,120]
[-121,-120,1,121,120]
[2,3,5,7]
[-121,-120,1,121,120]
[-121,-120,2,121,120]
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
