package interpreter_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/digvijay-tech/Haru/internal/interpreter"
	"github.com/digvijay-tech/Haru/internal/parser"
	"github.com/digvijay-tech/Haru/internal/preprocessor"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = stdout
	buf.ReadFrom(r)
	return buf.String()
}

func TestArithmeticExpressions(t *testing.T) {
	input := `
		--- ARITHMETIC EXPRESSIONS
		print "ARITHMETIC EXPRESSIONS";
		print 3 - 10;
		print 7.947 - 0.97;
		print 2 + 3 * 4;
		print (2 + 3) * 4;
		print 10 - 3 - 2;
		print 10 - (3 - 2);
		print (2 + 3) * (4 + 5);
		print 2 ** (3 ** 2);
		print (2 ** 3) ** 2;
		print 100 / (5 * (2 + 3));
		print (2 + 3) * (4 + 1) - 6;
		print ((20 - (4 * 2)) / 2) + 1;
		print 100 / (5 + 5) + (6 * 3 - 2);

		--- LOGICAL EXPRESSIONS
		print "LOGICAL EXPRESSIONS";
		print !true;					--- false
		print !false;					--- true
		print !!true;					--- true
		print !!false;					--- false
		print true && true;         	--- true
		print true && false;        	--- false
		print false || true;        	--- true
		print false || false;       	--- false
		print true && false || true; 	--- true (evaluates as (true && false) || true)
		print (true || false) && false; --- false

		--- COMPARISION EXPRESSIONS
		print "COMPARISION EXPRESSIONS";
		print 3 < 5;           				--- true
		print 10 > 7;          				--- true
		print 4 <= 4;          				--- true
		print 6 >= 10;         				--- false
		print 2 < 2;           				--- false
		print 5 >= 5;          				--- true
		print 5.0 < 10.0;	   				--- true
		print 3.0 >= 3.1;      				--- false
		print (2 + 3) * 4 > 10;     		--- true evaluates as (5 * 4 = 20) > 10
		print (5 * 2) <= (3 + 7);   		--- true evaluates as 10 <= 10
		print (10 - 2) > (2 * 4);   		--- false evaluates as 8 > 8
		print 3 + 2 < 6 && 7 > 2;   		--- true evaluates as 5 < 6 && true
		print 2 + 3 > 4 || 1 > 10;  		--- true evaluates as 5 > 4 || false
		print ((2 + 2) * (3 + 1)) >= 16; 	--- true evaluates as (4 * 4 = 16) >= 16
		print ((10 - 5) * 2) < (3 * 4);  	--- true evaluates as (5 * 2 = 10) < 12
		print (5 + 3) > (2 * 2) + 1;     	--- true evaluates as 8 > 5
		print ((2 + 3) * (4 + 1) > 20) || (6 / 2 + 1 < 5 && 3 * 3 == 9); --- true
		print 5 == 5;						--- true
		print 3.14 == 3.14;					--- true
		print 'hello' == 'hello';			--- true
		print 10 == (3 + 7); 				--- true
		print 5 != 3;						--- true
		print 3.14 != 2.71;					--- true
		print 'hello' != 'hello';			--- false
		print 5 != (2 + 3);					--- false
		print "Digvijay" == "digvijay";		--- false
		print true == true;					--- true
		print false == true;				--- false
	`

	expected := `ARITHMETIC EXPRESSIONS
-7
6.977
14
20
5
9
45
512
64
4
19
7
26
LOGICAL EXPRESSIONS
false
true
true
false
true
false
true
false
true
false
COMPARISION EXPRESSIONS
true
true
true
false
false
true
true
false
true
true
false
true
true
true
true
true
true
true
true
true
true
true
true
false
false
false
true
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
	output := captureOutput(func() {
		visitor := interpreter.NewHaruVisitor()
		visitor.Visit(tree)
	})

	// normalizing out for Windows
	output = strings.ReplaceAll(output, "\r\n", "\n")

	if output != expected {
		t.Errorf("\nExpected:\n%s\nGot:\n%s", expected, output)
	}
}
