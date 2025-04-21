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

	// avoid printing if function returned nothing
	if result == nil {
		return nil
	}

	// check for Value type
	if val, ok := result.(Value); ok {
		fmt.Println(val.Value)
	} else {
		// // avoid printing empty slices like [] or {}
		// if str, ok := result.(string); ok && strings.TrimSpace(str) == "" {
		// 	return nil
		// }

		fmt.Println("Output:", result)
	}

	return nil
}
