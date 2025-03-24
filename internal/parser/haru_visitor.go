// Code generated from haru.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // haru

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by haruParser.
type haruVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by haruParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by haruParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by haruParser#LetDecl.
	VisitLetDecl(ctx *LetDeclContext) interface{}

	// Visit a parse tree produced by haruParser#LetInferDecl.
	VisitLetInferDecl(ctx *LetInferDeclContext) interface{}

	// Visit a parse tree produced by haruParser#MutDecl.
	VisitMutDecl(ctx *MutDeclContext) interface{}

	// Visit a parse tree produced by haruParser#MutInferDecl.
	VisitMutInferDecl(ctx *MutInferDeclContext) interface{}

	// Visit a parse tree produced by haruParser#ConstDecl.
	VisitConstDecl(ctx *ConstDeclContext) interface{}

	// Visit a parse tree produced by haruParser#ConstInferDecl.
	VisitConstInferDecl(ctx *ConstInferDeclContext) interface{}

	// Visit a parse tree produced by haruParser#I8Type.
	VisitI8Type(ctx *I8TypeContext) interface{}

	// Visit a parse tree produced by haruParser#I16Type.
	VisitI16Type(ctx *I16TypeContext) interface{}

	// Visit a parse tree produced by haruParser#I32Type.
	VisitI32Type(ctx *I32TypeContext) interface{}

	// Visit a parse tree produced by haruParser#I64Type.
	VisitI64Type(ctx *I64TypeContext) interface{}

	// Visit a parse tree produced by haruParser#IntType.
	VisitIntType(ctx *IntTypeContext) interface{}

	// Visit a parse tree produced by haruParser#UI8Type.
	VisitUI8Type(ctx *UI8TypeContext) interface{}

	// Visit a parse tree produced by haruParser#UI16Type.
	VisitUI16Type(ctx *UI16TypeContext) interface{}

	// Visit a parse tree produced by haruParser#UI32Type.
	VisitUI32Type(ctx *UI32TypeContext) interface{}

	// Visit a parse tree produced by haruParser#UI64Type.
	VisitUI64Type(ctx *UI64TypeContext) interface{}

	// Visit a parse tree produced by haruParser#UIType.
	VisitUIType(ctx *UITypeContext) interface{}

	// Visit a parse tree produced by haruParser#F32Type.
	VisitF32Type(ctx *F32TypeContext) interface{}

	// Visit a parse tree produced by haruParser#F64Type.
	VisitF64Type(ctx *F64TypeContext) interface{}

	// Visit a parse tree produced by haruParser#BoolType.
	VisitBoolType(ctx *BoolTypeContext) interface{}

	// Visit a parse tree produced by haruParser#StringType.
	VisitStringType(ctx *StringTypeContext) interface{}

	// Visit a parse tree produced by haruParser#ByteType.
	VisitByteType(ctx *ByteTypeContext) interface{}

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

	// Visit a parse tree produced by haruParser#MulExpr.
	VisitMulExpr(ctx *MulExprContext) interface{}

	// Visit a parse tree produced by haruParser#DivExpr.
	VisitDivExpr(ctx *DivExprContext) interface{}

	// Visit a parse tree produced by haruParser#LitExpr.
	VisitLitExpr(ctx *LitExprContext) interface{}

	// Visit a parse tree produced by haruParser#SubExpr.
	VisitSubExpr(ctx *SubExprContext) interface{}

	// Visit a parse tree produced by haruParser#VarExpr.
	VisitVarExpr(ctx *VarExprContext) interface{}

	// Visit a parse tree produced by haruParser#AddExpr.
	VisitAddExpr(ctx *AddExprContext) interface{}

	// Visit a parse tree produced by haruParser#ModExpr.
	VisitModExpr(ctx *ModExprContext) interface{}

	// Visit a parse tree produced by haruParser#ParenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by haruParser#ExpExpr.
	VisitExpExpr(ctx *ExpExprContext) interface{}

	// Visit a parse tree produced by haruParser#AssignStmt.
	VisitAssignStmt(ctx *AssignStmtContext) interface{}

	// Visit a parse tree produced by haruParser#PrintStatement.
	VisitPrintStatement(ctx *PrintStatementContext) interface{}

	// Visit a parse tree produced by haruParser#IfStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}
}
