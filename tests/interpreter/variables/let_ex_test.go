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

func TestExplicitLet(t *testing.T) {
	input := `
	--- EXPLICITE LET DECLARATIONS
	print "EXPLICITE LET DECLARATIONS";	
	let a: i8 = 8;								
	let b: i32 = 100;							
	let c: int = a * b;                 		
	let d: f32 = a * b + c;             		
	let fname: string = "Digvijay";				
	let lname: string = "Padhiyar";				
	let fullname: string = fname + " " + lname;  
	let canDrive: bool = false;					
	let isProgrammer: bool = !canDrive;			

	print a;										--- 8
	print b;										--- 100
	print c;										--- 800
	print d;										--- 1600.000000
	print fullname;									--- Digvijay Padhiyar
	print canDrive;									--- false
	print isProgrammer;								--- true

	--- BOOLEAN
	let isHuman: bool = true;
	let canCode: bool = true;
	let result: bool = isHuman && isProgrammer && canCode;

	print result;						 --- true

	--- COMPARISON + TYPE MISMATCH CASES
	let x: i32 = 10;
	let y: i64 = 10;
	let xyEqual: bool = x == 10;

	print xyEqual;						 --- true

	--- let err: i8 = 9999;              --- should error: out of range
	--- let wrong: i8 = "oops";          --- should error: invalid type
	--- let undef: i32 = someVar;        --- should error: undefined var	
`

	expected := `EXPLICITE LET DECLARATIONS
8
100
800
1600.000000
Digvijay Padhiyar
false
true
true
true
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
