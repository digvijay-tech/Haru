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

	// manages variable and function scopes
	scopes []map[string]Value
}

// NewHaruVisitor initializes the visitor with empty state
func NewHaruVisitor() *HaruVisitor {
	return &HaruVisitor{
		scopes: []map[string]Value{},
	}
}

// Visit dispatches to specific VisitX methods
func (v *HaruVisitor) Visit(tree antlr.ParseTree) any {
	switch ctx := tree.(type) {
	case *parser.ProgramContext:
		return v.VisitProgram(ctx)
	case *parser.PrintStmtStatementContext:
		return v.VisitPrintStmtStatement(ctx)
	case *parser.PrintStatementContext:
		return v.VisitPrintStatement(ctx)
	case *parser.LitExprContext:
		return v.VisitLitExpr(ctx)
	case *parser.AddExprContext:
		return v.VisitAddExpr(ctx)
	case *parser.SubExprContext:
		return v.VisitSubExpr(ctx)
	case *parser.MulExprContext:
		return v.VisitMulExpr(ctx)
	case *parser.DivExprContext:
		return v.VisitDivExpr(ctx)
	case *parser.ModExprContext:
		return v.VisitModExpr(ctx)
	case *parser.ExpExprContext:
		return v.VisitExpExpr(ctx)
	case *parser.ParenExprContext:
		return v.VisitParenExpr(ctx)
	case *parser.NotExprContext:
		return v.VisitNotExpr(ctx)
	case *parser.AndExprContext:
		return v.VisitAndExpr(ctx)
	case *parser.OrExprContext:
		return v.VisitOrExpr(ctx)
	case *parser.EqExprContext:
		return v.VisitEqExpr(ctx)
	case *parser.NeExprContext:
		return v.VisitNeExpr(ctx)
	case *parser.LtExprContext:
		return v.VisitLtExpr(ctx)
	case *parser.LeExprContext:
		return v.VisitLeExpr(ctx)
	case *parser.GtExprContext:
		return v.VisitGtExpr(ctx)
	case *parser.GeExprContext:
		return v.VisitGeExpr(ctx)
	case *parser.VarExprContext:
		return v.VisitVarExpr(ctx)
	case *parser.VarDeclStatementContext:
		return v.VisitVarDeclStatement(ctx)
	case *parser.ConstDeclContext:
		return v.VisitExplicitConstDecl(ctx)
	case *parser.ConstInferDeclContext:
		return v.VisitImplicitConstDecl(ctx)
	case *parser.LetDeclContext:
		return v.VisitExplicitLetDecl(ctx)
	case *parser.LetInferDeclContext:
		return v.VisitImplicitLetDecl(ctx)
	case *parser.MutDeclContext:
		return v.VisitExplicitMutDecl(ctx)
	case *parser.MutInferDeclContext:
		return v.VisitImplicitMutDecl(ctx)
	case *parser.AssignStmtStatementContext:
		return v.VisitMutReassignment(ctx)
	case *parser.IfStmtStatementContext:
		return v.VisitIfStmt(ctx.IfStmt())
	case *parser.ArrayLiteralExprListContext:
		return v.VisitArrayLiteralExprList(ctx)
	case *parser.ArrayDeclStatementContext:
		return v.VisitArrayDeclStatement(ctx)
	case *parser.IndexExprContext:
		return v.VisitIndexExpr(ctx)
	case *parser.ArrayReassignStatementContext:
		return v.VisitMutArrayReassignment(ctx)
	case *parser.ArrayIndexAssignStatementContext:
		return v.VisitArrayIndexAssignStatement(ctx)
	case *parser.LenFunctionExprContext:
		return v.VisitLenFunction(ctx)
	case *parser.FunctionDeclStatementContext:
		return v.VisitFunctionDeclStatement(ctx)
	case *parser.FunctionCallStatementContext:
		return v.VisitFunctionCallStatement(ctx)
	case *parser.FunctionCallExprContext:
		return v.VisitFunctionCallExpression(ctx)
	case *parser.BlockContext:
		return v.VisitFunctionBlock(ctx)
	case *parser.ReturnStmtStatementContext:
		return v.VisitReturnStatement(ctx)
	case *parser.WhileLoopStatementContext:
		return v.VisitWhileLoop(ctx)
	case *parser.ImmutablePointerDeclContext:
		return v.VisitImmutablePointerDecl(ctx)
	case *parser.MutablePointerDeclContext:
		return v.VisitMutablePointerDecl(ctx)
	case *parser.DerefExprContext:
		return v.VisitPointerDerefExpr(ctx)
	case *parser.PointerAssignStmtStatementContext:
		return v.VisitPointerAssignment(ctx)
	case *parser.EmptyArrContext:
		return []Value{}
	default:
		fmt.Printf("Reached: %T\n", tree)
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
	// setting up global scope
	v.pushScope()

	for _, stmt := range ctx.AllStatement() {
		v.Visit(stmt)
	}

	// cleaning up global scope
	v.popScope()

	return nil
}
