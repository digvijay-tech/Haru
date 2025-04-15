// Code generated from ./grammar/haru.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // haru

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by haruParser.
type haruVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by haruParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by haruParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by haruParser#PrintStatement.
	VisitPrintStatement(ctx *PrintStatementContext) interface{}

	// Visit a parse tree produced by haruParser#MulExpr.
	VisitMulExpr(ctx *MulExprContext) interface{}

	// Visit a parse tree produced by haruParser#AndExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by haruParser#LitExpr.
	VisitLitExpr(ctx *LitExprContext) interface{}

	// Visit a parse tree produced by haruParser#NeExpr.
	VisitNeExpr(ctx *NeExprContext) interface{}

	// Visit a parse tree produced by haruParser#SubExpr.
	VisitSubExpr(ctx *SubExprContext) interface{}

	// Visit a parse tree produced by haruParser#LtExpr.
	VisitLtExpr(ctx *LtExprContext) interface{}

	// Visit a parse tree produced by haruParser#GtExpr.
	VisitGtExpr(ctx *GtExprContext) interface{}

	// Visit a parse tree produced by haruParser#AddExpr.
	VisitAddExpr(ctx *AddExprContext) interface{}

	// Visit a parse tree produced by haruParser#GeExpr.
	VisitGeExpr(ctx *GeExprContext) interface{}

	// Visit a parse tree produced by haruParser#ExpExpr.
	VisitExpExpr(ctx *ExpExprContext) interface{}

	// Visit a parse tree produced by haruParser#LeExpr.
	VisitLeExpr(ctx *LeExprContext) interface{}

	// Visit a parse tree produced by haruParser#OrExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by haruParser#ArrayExpr.
	VisitArrayExpr(ctx *ArrayExprContext) interface{}

	// Visit a parse tree produced by haruParser#DivExpr.
	VisitDivExpr(ctx *DivExprContext) interface{}

	// Visit a parse tree produced by haruParser#EqExpr.
	VisitEqExpr(ctx *EqExprContext) interface{}

	// Visit a parse tree produced by haruParser#VarExpr.
	VisitVarExpr(ctx *VarExprContext) interface{}

	// Visit a parse tree produced by haruParser#NotExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by haruParser#ModExpr.
	VisitModExpr(ctx *ModExprContext) interface{}

	// Visit a parse tree produced by haruParser#ParenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by haruParser#AssignStmt.
	VisitAssignStmt(ctx *AssignStmtContext) interface{}

	// Visit a parse tree produced by haruParser#IntLiteral.
	VisitIntLiteral(ctx *IntLiteralContext) interface{}

	// Visit a parse tree produced by haruParser#FloatLiteral.
	VisitFloatLiteral(ctx *FloatLiteralContext) interface{}

	// Visit a parse tree produced by haruParser#TrueLiteral.
	VisitTrueLiteral(ctx *TrueLiteralContext) interface{}

	// Visit a parse tree produced by haruParser#FalseLiteral.
	VisitFalseLiteral(ctx *FalseLiteralContext) interface{}

	// Visit a parse tree produced by haruParser#StringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by haruParser#ByteLiteral.
	VisitByteLiteral(ctx *ByteLiteralContext) interface{}
}
