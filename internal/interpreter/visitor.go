// Main parse tree visitor file
package interpreter

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/digvijay-tech/Haru/internal/parser"
)

// Main Haru visitor struct
type HaruVisitor struct {
	parser.BaseharuVisitor
	Vars map[string]Value
}

// NewHaruVisitor initializes the visitor with empty state
func NewHaruVisitor() *HaruVisitor {
	return &HaruVisitor{
		Vars: make(map[string]Value),
	}
}

// Visit dispatches to specific VisitX methods
func (v *HaruVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch ctx := tree.(type) {
	case *parser.ProgramContext:
		return v.VisitProgram(ctx)
	case *parser.PrintStatementContext:
		return v.VisitPrintStatement(ctx)
	case *parser.LitExprContext:
		return v.VisitLitExpr(ctx)
	}

	return v.VisitChildren(tree.(antlr.RuleNode))
}

// VisitChildren visits all child nodes
func (v *HaruVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		if t, ok := child.(antlr.ParseTree); ok {
			t.Accept(v)
		}
	}

	return nil
}

// VisitTerminal does nothing for terminals
func (v *HaruVisitor) VisitTerminal(node antlr.TerminalNode) any {
	return nil
}

// VisitErrorNode handles parsing errors
func (v *HaruVisitor) VisitErrorNode(node antlr.ErrorNode) any {
	fmt.Println("Parse error at:", node.GetText())
	return nil
}

// VisitProgram loops through all statements and evaluates them by calling v.Visit() method
func (v *HaruVisitor) VisitProgram(ctx *parser.ProgramContext) any {
	for _, stmt := range ctx.AllStatement() {
		v.Visit(stmt)
	}

	return nil
}
