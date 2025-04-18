package interpreter

import (
	"fmt"
	"strconv"

	"github.com/digvijay-tech/Haru/internal/parser"
)

// VisitIfStmt routes visitor to either IfBlockOnly block or IfElseChain block
func (v *HaruVisitor) VisitIfStmt(ctx parser.IIfStmtContext) any {
	switch t := ctx.(type) {
	case *parser.IfBlockOnlyContext:
		return v.VisitIfBlockOnly(t)
	case *parser.IfElseChainContext:
		fmt.Println("If else chain")
		return nil
	default:
		runtimeErr(fmt.Sprintf("unknown ifStmt variant: %T", t))
	}

	// for syntax validation, has no actual use
	return nil
}

// VisitIfBlockOnly evaluates if block only
func (v *HaruVisitor) VisitIfBlockOnly(ctx *parser.IfBlockOnlyContext) any {
	result := v.Visit(ctx.Expr())
	cond, ok := result.(Value)

	// ensuring condition is boolean
	if !ok || cond.Typ != "bool" {
		runtimeErr("if condition must be boolean")
	}

	// handles invalid values
	isTrue, err := strconv.ParseBool(cond.Value)

	if err != nil {
		runtimeErr("invalid boolean value in if condition")
	}

	// if true, visits all statements wrapped in the if block
	if isTrue {
		for _, stmt := range ctx.Block().AllStatement() {
			v.Visit(stmt)
		}
	}

	return nil
}
