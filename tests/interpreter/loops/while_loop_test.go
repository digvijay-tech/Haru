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

func TestPrintLiterals(t *testing.T) {
	input := `
	--- TESTING WHILE LOOP
	print "TESTING WHILE LOOP";

	mut i: i8 = 1;

	while (i <= 10) {
		print i;
		i = i + 1;
	}

	print "=================";
	--- loop with local scope per iteration
	mut j = 0;

	while (j < 3) {
		const msg: string = "Loop at j = ";
		print msg;
		print j;

		j = j + 1;
	}

	--- print msg; --- Runtime Error: undefined variable 'msg'

	print "=================";
	--- iterating over an array
	mut fruites: []string = ["apple", "cherry", "mango", "banana", "watermelon"];

	mut fi: i8 = 0;
	while (fi < len(fruites)) {
		print "Printing: " + fruites[fi];

		fi = fi + 1;
	}

	print "=================";
	--- iterating over an array
	mut x = 0;

	while (x < 2) {
		mut j = 0;

		while (j < 3) {
			print j;
			j = j + 1;
			print "END";
		}

		x = x + 1;
	}

	print x;
`

	expected := `TESTING WHILE LOOP
1
2
3
4
5
6
7
8
9
10
=================
Loop at j = 
0
Loop at j = 
1
Loop at j = 
2
=================
Printing: apple
Printing: cherry
Printing: mango
Printing: banana
Printing: watermelon
=================
0
END
1
END
2
END
0
END
1
END
2
END
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
