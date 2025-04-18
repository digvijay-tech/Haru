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

func TestMutReassignment(t *testing.T) {
	input := `
	--- VARIABLE REASSIGNMENTS
	print "VARIABLE REASSIGNMENTS";
	mut a: i8 = 5;
	print a;
	a = 10;
	print a;

	mut b: i16 = 300;
	print b;
	b = -123;
	print b;

	mut c: i32 = 100000;
	print c;
	c = 42;
	print c;

	mut d: i64 = 9000000000;
	print d;
	d = -1;
	print d;

	mut e: int = 999;
	print e;
	e = 123456;
	print e;

	mut f: ui8 = 250;
	print f;
	f = 100;
	print f;

	mut g: ui16 = 65530;
	print g;
	g = 42;
	print g;

	mut h: ui32 = 1000000;
	print h;
	h = 123;
	print h;

	mut i: ui64 = 10000000000;
	print i;
	i = 987654321;
	print i;

	mut j: uint = 2048;
	print j;
	j = 4096;
	print j;

	mut k: f32 = 3.14;
	print k;
	k = 1.618;
	print k;

	mut l: f64 = 2.718281828;
	print l;
	l = 6.62607015;
	print l;

	mut m: bool = true;
	print m;
	m = false;
	print m;

	mut n: string = "Hello";
	print n;
	n = "World";
	print n;

	mut o: byte = 254;
	print o;
	o = 255;
	print o;

	mut p: bool = n != "hello";
	print p;
	p = !p;
	print p;
`

	expected := `VARIABLE REASSIGNMENTS
5
10
300
-123
100000
42
9000000000
-1
999
123456
250
100
65530
42
1000000
123
10000000000
987654321
2048
4096
3.140000
1.618000
2.718282
6.626070
true
false
Hello
World
254
255
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
