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
		return v.VisitIfElseChain(t)
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

	// if true, visit all statements wrapped in the if block
	if isTrue {
		// adding local scope
		v.pushScope()

		for _, stmt := range ctx.Block().AllStatement() {
			v.Visit(stmt)
		}

		// removing local scope
		v.popScope()
	}

	return nil
}

// VisitIfElseChain evaluates if-else, if-elseif and if-elseif-else blocks
func (v *HaruVisitor) VisitIfElseChain(ctx *parser.IfElseChainContext) any {
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

	// if true, visit all statements wrapped in the if block
	if isTrue {
		// adding local scope
		v.pushScope()

		for _, stmt := range ctx.Block().AllStatement() {
			v.Visit(stmt)
		}

		// removing local scope
		v.popScope()

		// prevents furthur execution beyond if statement
		return nil
	}

	// evaluating else-if blocks
	for _, elseIf := range ctx.AllElseIfBlock() {
		result := v.Visit(elseIf.Expr())
		cond, ok := result.(Value)

		if !ok || cond.Typ != "bool" {
			runtimeErr("else if condition must be boolean")
		}

		isTrue, err := strconv.ParseBool(cond.Value)
		if err != nil {
			runtimeErr("invalid boolean value in else if condition")
		}

		if isTrue {
			// adding local scope
			v.pushScope()

			// calling visit if there are any nested statements
			for _, stmt := range elseIf.Block().AllStatement() {
				v.Visit(stmt)
			}

			// removing local scope
			v.popScope()

			// prevents furthur execution beyond else-if statement
			return nil
		}
	}

	// evaluating else block if present
	if ctx.ElseBlock() != nil {
		// adding local scope
		v.pushScope()

		for _, stmt := range ctx.ElseBlock().Block().AllStatement() {
			v.Visit(stmt)
		}

		// removing local scope
		v.popScope()
	}

	return nil
}
