// Code generated from ./grammar/haru.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // haru

import "github.com/antlr4-go/antlr/v4"

type BaseharuVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseharuVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitPrintStmtStatement(ctx *PrintStmtStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitVarDeclStatement(ctx *VarDeclStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitAssignStmtStatement(ctx *AssignStmtStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitIfStmtStatement(ctx *IfStmtStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitArrayDeclStatement(ctx *ArrayDeclStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitArrayIndexAssignStatement(ctx *ArrayIndexAssignStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitArrayReassignStatement(ctx *ArrayReassignStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitFunctionDeclStatement(ctx *FunctionDeclStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitReturnStmtStatement(ctx *ReturnStmtStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitFunctionCallStatement(ctx *FunctionCallStatementContext) interface{} {
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

func (v *BaseharuVisitor) VisitIndexExpr(ctx *IndexExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitLeExpr(ctx *LeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitDivExpr(ctx *DivExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitFunctionCallExpr(ctx *FunctionCallExprContext) interface{} {
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

func (v *BaseharuVisitor) VisitLenFunctionExpr(ctx *LenFunctionExprContext) interface{} {
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

func (v *BaseharuVisitor) VisitUIntType(ctx *UIntTypeContext) interface{} {
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

func (v *BaseharuVisitor) VisitIfBlockOnly(ctx *IfBlockOnlyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitIfElseChain(ctx *IfElseChainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitElseIfBlock(ctx *ElseIfBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitElseBlock(ctx *ElseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitArrayDecl(ctx *ArrayDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitConstExplicitArrayDecl(ctx *ConstExplicitArrayDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitConstImplicitArrayDecl(ctx *ConstImplicitArrayDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitLetExplicitArrayDecl(ctx *LetExplicitArrayDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitLetImplicitArrayDecl(ctx *LetImplicitArrayDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitMutFixedArrayWithInit(ctx *MutFixedArrayWithInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitMutFixedArrayNoInit(ctx *MutFixedArrayNoInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitMutArrayExplicitWithInit(ctx *MutArrayExplicitWithInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitMutArrayExplicitNoInit(ctx *MutArrayExplicitNoInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitMutArrayImplicit(ctx *MutArrayImplicitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitFixedArrayType(ctx *FixedArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitArrayLiteralExprList(ctx *ArrayLiteralExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitEmptyArr(ctx *EmptyArrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitArrayIndexAssign(ctx *ArrayIndexAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitArrayReassign(ctx *ArrayReassignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitLenFunction(ctx *LenFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitFunctionDecl(ctx *FunctionDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitParamList(ctx *ParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitParam(ctx *ParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitReturnSignature(ctx *ReturnSignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitExprList(ctx *ExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseharuVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}
