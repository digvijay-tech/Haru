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

func TestConstantsImplicitConstants(t *testing.T) {
	input := `
	print "IMPLICIT CONSTANT DECLARATION";

	const a = 1;
	const b = -1;
	const c = 0;
	const e = -0.1283198; 										--- -0.128320
	const f = 0.00012318; 										--- 0.000123
	const g = 120837182; 										--- 120837182
	const h = -120837182; 										--- -120837182
	const i = 5.92233720368547758079456724567924475765765; 		--- 5.922337
	const j = 0.2; 												--- 0.200000
	const k = 1.23; 											--- 1.230000
	const l = 234177613746487618; 								--- 234177613746487618
	const m = 4.1764176417647164913479364913756971364591637;	--- 4.176418
	const n = false;											--- false
	const o = !n;												--- true
	const p = "Ping!";											--- Ping!
	const q = 'queue';											--- queue
	const r = (10 * (a + 3)) - 4;								--- 36

	print a;
	print b;
	print c;
	print e;
	print f;
	print g;
	print h;
	print i;
	print j;
	print k;
	print l;
	print m;
	print n;
	print o;
	print p;
	print q;
	print r;
`

	expected := `IMPLICIT CONSTANT DECLARATION
1
-1
0
-0.128320
0.000123
120837182
-120837182
5.922337
0.200000
1.230000
234177613746487618
4.176418
false
true
Ping!
queue
36
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
