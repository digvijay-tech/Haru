// Code generated from ./grammar/haru.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // haru

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by haruParser.
type haruVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by haruParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by haruParser#PrintStmtStatement.
	VisitPrintStmtStatement(ctx *PrintStmtStatementContext) interface{}

	// Visit a parse tree produced by haruParser#VarDeclStatement.
	VisitVarDeclStatement(ctx *VarDeclStatementContext) interface{}

	// Visit a parse tree produced by haruParser#AssignStmtStatement.
	VisitAssignStmtStatement(ctx *AssignStmtStatementContext) interface{}

	// Visit a parse tree produced by haruParser#PointerAssignStmtStatement.
	VisitPointerAssignStmtStatement(ctx *PointerAssignStmtStatementContext) interface{}

	// Visit a parse tree produced by haruParser#IfStmtStatement.
	VisitIfStmtStatement(ctx *IfStmtStatementContext) interface{}

	// Visit a parse tree produced by haruParser#ArrayDeclStatement.
	VisitArrayDeclStatement(ctx *ArrayDeclStatementContext) interface{}

	// Visit a parse tree produced by haruParser#ArrayIndexAssignStatement.
	VisitArrayIndexAssignStatement(ctx *ArrayIndexAssignStatementContext) interface{}

	// Visit a parse tree produced by haruParser#ArrayReassignStatement.
	VisitArrayReassignStatement(ctx *ArrayReassignStatementContext) interface{}

	// Visit a parse tree produced by haruParser#FunctionDeclStatement.
	VisitFunctionDeclStatement(ctx *FunctionDeclStatementContext) interface{}

	// Visit a parse tree produced by haruParser#ReturnStmtStatement.
	VisitReturnStmtStatement(ctx *ReturnStmtStatementContext) interface{}

	// Visit a parse tree produced by haruParser#FunctionCallStatement.
	VisitFunctionCallStatement(ctx *FunctionCallStatementContext) interface{}

	// Visit a parse tree produced by haruParser#WhileLoopStatement.
	VisitWhileLoopStatement(ctx *WhileLoopStatementContext) interface{}

	// Visit a parse tree produced by haruParser#PrintStatement.
	VisitPrintStatement(ctx *PrintStatementContext) interface{}

	// Visit a parse tree produced by haruParser#MulExpr.
	VisitMulExpr(ctx *MulExprContext) interface{}

	// Visit a parse tree produced by haruParser#AndExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by haruParser#DerefExpr.
	VisitDerefExpr(ctx *DerefExprContext) interface{}

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

	// Visit a parse tree produced by haruParser#IndexExpr.
	VisitIndexExpr(ctx *IndexExprContext) interface{}

	// Visit a parse tree produced by haruParser#LeExpr.
	VisitLeExpr(ctx *LeExprContext) interface{}

	// Visit a parse tree produced by haruParser#OrExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by haruParser#InputFunctionExpr.
	VisitInputFunctionExpr(ctx *InputFunctionExprContext) interface{}

	// Visit a parse tree produced by haruParser#DivExpr.
	VisitDivExpr(ctx *DivExprContext) interface{}

	// Visit a parse tree produced by haruParser#FunctionCallExpr.
	VisitFunctionCallExpr(ctx *FunctionCallExprContext) interface{}

	// Visit a parse tree produced by haruParser#AddressOfExpr.
	VisitAddressOfExpr(ctx *AddressOfExprContext) interface{}

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

	// Visit a parse tree produced by haruParser#LenFunctionExpr.
	VisitLenFunctionExpr(ctx *LenFunctionExprContext) interface{}

	// Visit a parse tree produced by haruParser#AssignStmt.
	VisitAssignStmt(ctx *AssignStmtContext) interface{}

	// Visit a parse tree produced by haruParser#PointerAssignStmt.
	VisitPointerAssignStmt(ctx *PointerAssignStmtContext) interface{}

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

	// Visit a parse tree produced by haruParser#ImmutablePointerDecl.
	VisitImmutablePointerDecl(ctx *ImmutablePointerDeclContext) interface{}

	// Visit a parse tree produced by haruParser#MutablePointerDecl.
	VisitMutablePointerDecl(ctx *MutablePointerDeclContext) interface{}

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

	// Visit a parse tree produced by haruParser#UIntType.
	VisitUIntType(ctx *UIntTypeContext) interface{}

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

	// Visit a parse tree produced by haruParser#IfBlockOnly.
	VisitIfBlockOnly(ctx *IfBlockOnlyContext) interface{}

	// Visit a parse tree produced by haruParser#IfElseChain.
	VisitIfElseChain(ctx *IfElseChainContext) interface{}

	// Visit a parse tree produced by haruParser#elseIfBlock.
	VisitElseIfBlock(ctx *ElseIfBlockContext) interface{}

	// Visit a parse tree produced by haruParser#elseBlock.
	VisitElseBlock(ctx *ElseBlockContext) interface{}

	// Visit a parse tree produced by haruParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by haruParser#arrayDecl.
	VisitArrayDecl(ctx *ArrayDeclContext) interface{}

	// Visit a parse tree produced by haruParser#ConstExplicitArrayDecl.
	VisitConstExplicitArrayDecl(ctx *ConstExplicitArrayDeclContext) interface{}

	// Visit a parse tree produced by haruParser#ConstImplicitArrayDecl.
	VisitConstImplicitArrayDecl(ctx *ConstImplicitArrayDeclContext) interface{}

	// Visit a parse tree produced by haruParser#LetExplicitArrayDecl.
	VisitLetExplicitArrayDecl(ctx *LetExplicitArrayDeclContext) interface{}

	// Visit a parse tree produced by haruParser#LetImplicitArrayDecl.
	VisitLetImplicitArrayDecl(ctx *LetImplicitArrayDeclContext) interface{}

	// Visit a parse tree produced by haruParser#MutFixedArrayWithInit.
	VisitMutFixedArrayWithInit(ctx *MutFixedArrayWithInitContext) interface{}

	// Visit a parse tree produced by haruParser#MutFixedArrayNoInit.
	VisitMutFixedArrayNoInit(ctx *MutFixedArrayNoInitContext) interface{}

	// Visit a parse tree produced by haruParser#MutArrayExplicitWithInit.
	VisitMutArrayExplicitWithInit(ctx *MutArrayExplicitWithInitContext) interface{}

	// Visit a parse tree produced by haruParser#MutArrayExplicitNoInit.
	VisitMutArrayExplicitNoInit(ctx *MutArrayExplicitNoInitContext) interface{}

	// Visit a parse tree produced by haruParser#MutArrayImplicit.
	VisitMutArrayImplicit(ctx *MutArrayImplicitContext) interface{}

	// Visit a parse tree produced by haruParser#arrayType.
	VisitArrayType(ctx *ArrayTypeContext) interface{}

	// Visit a parse tree produced by haruParser#fixedArrayType.
	VisitFixedArrayType(ctx *FixedArrayTypeContext) interface{}

	// Visit a parse tree produced by haruParser#ArrayLiteralExprList.
	VisitArrayLiteralExprList(ctx *ArrayLiteralExprListContext) interface{}

	// Visit a parse tree produced by haruParser#EmptyArr.
	VisitEmptyArr(ctx *EmptyArrContext) interface{}

	// Visit a parse tree produced by haruParser#ArrayIndexAssign.
	VisitArrayIndexAssign(ctx *ArrayIndexAssignContext) interface{}

	// Visit a parse tree produced by haruParser#arrayReassign.
	VisitArrayReassign(ctx *ArrayReassignContext) interface{}

	// Visit a parse tree produced by haruParser#lenFunction.
	VisitLenFunction(ctx *LenFunctionContext) interface{}

	// Visit a parse tree produced by haruParser#inputFunction.
	VisitInputFunction(ctx *InputFunctionContext) interface{}

	// Visit a parse tree produced by haruParser#functionDecl.
	VisitFunctionDecl(ctx *FunctionDeclContext) interface{}

	// Visit a parse tree produced by haruParser#paramList.
	VisitParamList(ctx *ParamListContext) interface{}

	// Visit a parse tree produced by haruParser#param.
	VisitParam(ctx *ParamContext) interface{}

	// Visit a parse tree produced by haruParser#returnSignature.
	VisitReturnSignature(ctx *ReturnSignatureContext) interface{}

	// Visit a parse tree produced by haruParser#returnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by haruParser#exprList.
	VisitExprList(ctx *ExprListContext) interface{}

	// Visit a parse tree produced by haruParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by haruParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by haruParser#whileLoop.
	VisitWhileLoop(ctx *WhileLoopContext) interface{}
}
