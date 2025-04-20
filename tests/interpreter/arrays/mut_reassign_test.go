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

func TestMutArrayReassignment(t *testing.T) {
	input := `
	--- MUT ARRAY REASSIGNMENTS
	--- reassignment can only work with array of exactly same type

	print "MUT ARRAY REASSIGNMENTS";
	mut nums: []i8 = [1,2,3,4];
	print nums;

	nums = [4,3,2,1];
	print nums;
	
	let primes: []i32 = [2, 3, 5, 7];
	--- nums = primes; --- Runtime Error: type mismatch in assignmen
	
	nums = [10, 20, 30, 40, 50];
	print nums;

	nums = [1];
	print nums;

	mut days = ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"];
	print days;

	let currentDay = "Friday";
	let employee = "John";
	mut isWorking = false;

	if (employee == "John") {
		if (currentDay != "Saturday" && currentDay != "Sunday") {
			isWorking = true;
		} else {
			isWorking = false; 
		}
	}

	if (!isWorking) {
		print employee + " doesn't work on weekends!";
	}

	if (isWorking) {
		print employee + " is working today!";
	}

	days = ["Holiday"];
	print days;	

	days = [];
	print days;
`

	expected := `MUT ARRAY REASSIGNMENTS
[1,2,3,4]
[4,3,2,1]
[10,20,30,40,50]
[1]
[Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday]
John is working today!
[Holiday]
[]
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
