// Print statement evaluation
package interpreter

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/digvijay-tech/Haru/internal/parser"
)

func (v *HaruVisitor) VisitPrintStmtStatement(ctx *parser.PrintStmtStatementContext) any {
	return v.Visit(ctx.GetChild(0).(antlr.ParseTree)) // Visits actual printStmt
}

// VisitPrintStatement handles the print statement
func (v *HaruVisitor) VisitPrintStatement(ctx *parser.PrintStatementContext) any {
	// walks the parse tree for the expression and eventually return the result of type Value
	exprCtx := ctx.Expr()
	result := v.Visit(exprCtx)

	// verifying the type result to be Value
	if val, ok := result.(Value); ok {
		fmt.Println(val.Value)
	} else {
		// printing non-Value result
		fmt.Println("Output:", result)
	}

	return nil
}
