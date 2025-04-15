// Print statement evaluation
package interpreter

import (
	"fmt"

	"github.com/digvijay-tech/Haru/internal/parser"
)

// VisitPrintStatement handles the print statement
func (v *HaruVisitor) VisitPrintStatement(ctx *parser.PrintStatementContext) any {
	// walks the parse tree for the expression and eventually return the result of type Value
	result := v.Visit(ctx.Expr())

	// verifying the type result to be Value
	if val, ok := result.(Value); ok {
		fmt.Println(val.Value)
	} else {
		// printing non-Value result
		fmt.Println("Output:", result)
	}

	return nil
}
