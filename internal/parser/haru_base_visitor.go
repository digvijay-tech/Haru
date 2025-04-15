// Code generated from ./grammar/haru.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // haru

import "github.com/antlr4-go/antlr/v4"

type BaseharuVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseharuVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitPrintStatement(ctx *PrintStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitMulExpr(ctx *MulExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitLitExpr(ctx *LitExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitNeExpr(ctx *NeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitSubExpr(ctx *SubExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitLtExpr(ctx *LtExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitGtExpr(ctx *GtExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitAddExpr(ctx *AddExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitGeExpr(ctx *GeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitExpExpr(ctx *ExpExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitLeExpr(ctx *LeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitArrayExpr(ctx *ArrayExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitDivExpr(ctx *DivExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitEqExpr(ctx *EqExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitVarExpr(ctx *VarExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitModExpr(ctx *ModExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitAssignStmt(ctx *AssignStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitIntLiteral(ctx *IntLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitFloatLiteral(ctx *FloatLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitTrueLiteral(ctx *TrueLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitFalseLiteral(ctx *FalseLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitByteLiteral(ctx *ByteLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
