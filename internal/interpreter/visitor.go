package interpreter

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/digvijay-tech/Haru/internal/parser"
)

// HaruVisitor is the main visitor for walking the parse tree
type HaruVisitor struct {
	parser.BaseharuVisitor
	Vars map[string]Value
}

// NewHaruVisitor creates a new visitor with an empty variable map
func NewHaruVisitor() *HaruVisitor {
	return &HaruVisitor{
		Vars: make(map[string]Value),
	}
}

// Visit dispatches to specific handlers based on context type
func (v *HaruVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch ctx := tree.(type) {
	case *parser.ProgramContext:
		return v.VisitProgram(ctx)
	case *parser.LetDeclContext:
		return v.VisitLetDecl(ctx)
	case *parser.MutDeclContext:
		return v.VisitMutDecl(ctx)
	case *parser.AssignStmtContext:
		return v.VisitAssignStmt(ctx)
	case *parser.PrintStatementContext:
		return v.VisitPrintStatement(ctx)
	case *parser.IfStatementContext:
		return v.VisitIfStatement(ctx)
	case antlr.TerminalNode:
		return nil
	}
	return v.VisitChildren(tree.(antlr.RuleNode))
}

// VisitChildren walks all child nodes
func (v *HaruVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, child := range node.GetChildren() {
		if tree, ok := child.(antlr.ParseTree); ok {
			tree.Accept(v)
		}
	}
	return nil
}

// VisitTerminal skips terminal nodes
func (v *HaruVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return nil
}

// VisitErrorNode reports parsing errors
func (v *HaruVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	fmt.Println("Error in parsing:", node.GetText())
	return nil
}
