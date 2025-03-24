// Code generated from haru.g4 by ANTLR 4.13.2. DO NOT EDIT.

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

func (v *BaseharuVisitor) VisitLetDecl(ctx *LetDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitLetInferDecl(ctx *LetInferDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitMutDecl(ctx *MutDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitMutInferDecl(ctx *MutInferDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitConstDecl(ctx *ConstDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitConstInferDecl(ctx *ConstInferDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitI8Type(ctx *I8TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitI16Type(ctx *I16TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitI32Type(ctx *I32TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitI64Type(ctx *I64TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitIntType(ctx *IntTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitUI8Type(ctx *UI8TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitUI16Type(ctx *UI16TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitUI32Type(ctx *UI32TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitUI64Type(ctx *UI64TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitUIType(ctx *UITypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitF32Type(ctx *F32TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitF64Type(ctx *F64TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitBoolType(ctx *BoolTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitStringType(ctx *StringTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitByteType(ctx *ByteTypeContext) interface{} {
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

func (v *BaseharuVisitor) VisitMulExpr(ctx *MulExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitDivExpr(ctx *DivExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitLitExpr(ctx *LitExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitSubExpr(ctx *SubExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitVarExpr(ctx *VarExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitAddExpr(ctx *AddExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitModExpr(ctx *ModExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitExpExpr(ctx *ExpExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitAssignStmt(ctx *AssignStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitPrintStatement(ctx *PrintStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}
