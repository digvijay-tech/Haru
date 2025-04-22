package interpreter

import (
	"strconv"
	"time"

	"github.com/digvijay-tech/Haru/internal/parser"
)

// VisitWhileLoop evaluates while loop statements
func (v *HaruVisitor) VisitWhileLoop(ctx *parser.WhileLoopStatementContext) any {
	for {
		// evaluating loop condition
		cond := v.Visit(ctx.WhileLoop().Expr())
		condVal, ok := cond.(Value)

		if !ok || condVal.Typ != "bool" {
			runtimeErr("while condition must evaluate to a boolean")
		}

		isTrue, err := strconv.ParseBool(condVal.Value)

		if err != nil {
			runtimeErr("invalid boolean value in while loop condition")
		}

		if !isTrue {
			break
		}

		// pushing a new scope for each iteration (if inner variables are declared)
		v.pushScope()

		for _, stmt := range ctx.WhileLoop().Block().AllStatement() {
			v.Visit(stmt)
			time.Sleep(50 * time.Millisecond)
		}

		// cleaning up scope after each iteration
		v.popScope()
	}

	return nil
}
