// Code generated from ./grammar/haru.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // haru

import "github.com/antlr4-go/antlr/v4"

// haruListener is a complete listener for a parse tree produced by haruParser.
type haruListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterPrintStmtStatement is called when entering the PrintStmtStatement production.
	EnterPrintStmtStatement(c *PrintStmtStatementContext)

	// EnterVarDeclStatement is called when entering the VarDeclStatement production.
	EnterVarDeclStatement(c *VarDeclStatementContext)

	// EnterAssignStmtStatement is called when entering the AssignStmtStatement production.
	EnterAssignStmtStatement(c *AssignStmtStatementContext)

	// EnterIfStmtStatement is called when entering the IfStmtStatement production.
	EnterIfStmtStatement(c *IfStmtStatementContext)

	// EnterArrayDeclStatement is called when entering the ArrayDeclStatement production.
	EnterArrayDeclStatement(c *ArrayDeclStatementContext)

	// EnterArrayIndexAssignStatement is called when entering the ArrayIndexAssignStatement production.
	EnterArrayIndexAssignStatement(c *ArrayIndexAssignStatementContext)

	// EnterArrayReassignStatement is called when entering the ArrayReassignStatement production.
	EnterArrayReassignStatement(c *ArrayReassignStatementContext)

	// EnterFunctionDeclStatement is called when entering the FunctionDeclStatement production.
	EnterFunctionDeclStatement(c *FunctionDeclStatementContext)

	// EnterReturnStmtStatement is called when entering the ReturnStmtStatement production.
	EnterReturnStmtStatement(c *ReturnStmtStatementContext)

	// EnterFunctionCallStatement is called when entering the FunctionCallStatement production.
	EnterFunctionCallStatement(c *FunctionCallStatementContext)

	// EnterWhileLoopStatement is called when entering the WhileLoopStatement production.
	EnterWhileLoopStatement(c *WhileLoopStatementContext)

	// EnterPrintStatement is called when entering the PrintStatement production.
	EnterPrintStatement(c *PrintStatementContext)

	// EnterMulExpr is called when entering the MulExpr production.
	EnterMulExpr(c *MulExprContext)

	// EnterAndExpr is called when entering the AndExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterLitExpr is called when entering the LitExpr production.
	EnterLitExpr(c *LitExprContext)

	// EnterNeExpr is called when entering the NeExpr production.
	EnterNeExpr(c *NeExprContext)

	// EnterSubExpr is called when entering the SubExpr production.
	EnterSubExpr(c *SubExprContext)

	// EnterLtExpr is called when entering the LtExpr production.
	EnterLtExpr(c *LtExprContext)

	// EnterGtExpr is called when entering the GtExpr production.
	EnterGtExpr(c *GtExprContext)

	// EnterAddExpr is called when entering the AddExpr production.
	EnterAddExpr(c *AddExprContext)

	// EnterGeExpr is called when entering the GeExpr production.
	EnterGeExpr(c *GeExprContext)

	// EnterExpExpr is called when entering the ExpExpr production.
	EnterExpExpr(c *ExpExprContext)

	// EnterIndexExpr is called when entering the IndexExpr production.
	EnterIndexExpr(c *IndexExprContext)

	// EnterLeExpr is called when entering the LeExpr production.
	EnterLeExpr(c *LeExprContext)

	// EnterOrExpr is called when entering the OrExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterDivExpr is called when entering the DivExpr production.
	EnterDivExpr(c *DivExprContext)

	// EnterFunctionCallExpr is called when entering the FunctionCallExpr production.
	EnterFunctionCallExpr(c *FunctionCallExprContext)

	// EnterEqExpr is called when entering the EqExpr production.
	EnterEqExpr(c *EqExprContext)

	// EnterVarExpr is called when entering the VarExpr production.
	EnterVarExpr(c *VarExprContext)

	// EnterNotExpr is called when entering the NotExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterModExpr is called when entering the ModExpr production.
	EnterModExpr(c *ModExprContext)

	// EnterParenExpr is called when entering the ParenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterLenFunctionExpr is called when entering the LenFunctionExpr production.
	EnterLenFunctionExpr(c *LenFunctionExprContext)

	// EnterAssignStmt is called when entering the AssignStmt production.
	EnterAssignStmt(c *AssignStmtContext)

	// EnterIntLiteral is called when entering the IntLiteral production.
	EnterIntLiteral(c *IntLiteralContext)

	// EnterFloatLiteral is called when entering the FloatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// EnterTrueLiteral is called when entering the TrueLiteral production.
	EnterTrueLiteral(c *TrueLiteralContext)

	// EnterFalseLiteral is called when entering the FalseLiteral production.
	EnterFalseLiteral(c *FalseLiteralContext)

	// EnterStringLiteral is called when entering the StringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterByteLiteral is called when entering the ByteLiteral production.
	EnterByteLiteral(c *ByteLiteralContext)

	// EnterLetDecl is called when entering the LetDecl production.
	EnterLetDecl(c *LetDeclContext)

	// EnterLetInferDecl is called when entering the LetInferDecl production.
	EnterLetInferDecl(c *LetInferDeclContext)

	// EnterMutDecl is called when entering the MutDecl production.
	EnterMutDecl(c *MutDeclContext)

	// EnterMutInferDecl is called when entering the MutInferDecl production.
	EnterMutInferDecl(c *MutInferDeclContext)

	// EnterConstDecl is called when entering the ConstDecl production.
	EnterConstDecl(c *ConstDeclContext)

	// EnterConstInferDecl is called when entering the ConstInferDecl production.
	EnterConstInferDecl(c *ConstInferDeclContext)

	// EnterI8Type is called when entering the I8Type production.
	EnterI8Type(c *I8TypeContext)

	// EnterI16Type is called when entering the I16Type production.
	EnterI16Type(c *I16TypeContext)

	// EnterI32Type is called when entering the I32Type production.
	EnterI32Type(c *I32TypeContext)

	// EnterI64Type is called when entering the I64Type production.
	EnterI64Type(c *I64TypeContext)

	// EnterIntType is called when entering the IntType production.
	EnterIntType(c *IntTypeContext)

	// EnterUI8Type is called when entering the UI8Type production.
	EnterUI8Type(c *UI8TypeContext)

	// EnterUI16Type is called when entering the UI16Type production.
	EnterUI16Type(c *UI16TypeContext)

	// EnterUI32Type is called when entering the UI32Type production.
	EnterUI32Type(c *UI32TypeContext)

	// EnterUI64Type is called when entering the UI64Type production.
	EnterUI64Type(c *UI64TypeContext)

	// EnterUIntType is called when entering the UIntType production.
	EnterUIntType(c *UIntTypeContext)

	// EnterF32Type is called when entering the F32Type production.
	EnterF32Type(c *F32TypeContext)

	// EnterF64Type is called when entering the F64Type production.
	EnterF64Type(c *F64TypeContext)

	// EnterBoolType is called when entering the BoolType production.
	EnterBoolType(c *BoolTypeContext)

	// EnterStringType is called when entering the StringType production.
	EnterStringType(c *StringTypeContext)

	// EnterByteType is called when entering the ByteType production.
	EnterByteType(c *ByteTypeContext)

	// EnterIfBlockOnly is called when entering the IfBlockOnly production.
	EnterIfBlockOnly(c *IfBlockOnlyContext)

	// EnterIfElseChain is called when entering the IfElseChain production.
	EnterIfElseChain(c *IfElseChainContext)

	// EnterElseIfBlock is called when entering the elseIfBlock production.
	EnterElseIfBlock(c *ElseIfBlockContext)

	// EnterElseBlock is called when entering the elseBlock production.
	EnterElseBlock(c *ElseBlockContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterArrayDecl is called when entering the arrayDecl production.
	EnterArrayDecl(c *ArrayDeclContext)

	// EnterConstExplicitArrayDecl is called when entering the ConstExplicitArrayDecl production.
	EnterConstExplicitArrayDecl(c *ConstExplicitArrayDeclContext)

	// EnterConstImplicitArrayDecl is called when entering the ConstImplicitArrayDecl production.
	EnterConstImplicitArrayDecl(c *ConstImplicitArrayDeclContext)

	// EnterLetExplicitArrayDecl is called when entering the LetExplicitArrayDecl production.
	EnterLetExplicitArrayDecl(c *LetExplicitArrayDeclContext)

	// EnterLetImplicitArrayDecl is called when entering the LetImplicitArrayDecl production.
	EnterLetImplicitArrayDecl(c *LetImplicitArrayDeclContext)

	// EnterMutFixedArrayWithInit is called when entering the MutFixedArrayWithInit production.
	EnterMutFixedArrayWithInit(c *MutFixedArrayWithInitContext)

	// EnterMutFixedArrayNoInit is called when entering the MutFixedArrayNoInit production.
	EnterMutFixedArrayNoInit(c *MutFixedArrayNoInitContext)

	// EnterMutArrayExplicitWithInit is called when entering the MutArrayExplicitWithInit production.
	EnterMutArrayExplicitWithInit(c *MutArrayExplicitWithInitContext)

	// EnterMutArrayExplicitNoInit is called when entering the MutArrayExplicitNoInit production.
	EnterMutArrayExplicitNoInit(c *MutArrayExplicitNoInitContext)

	// EnterMutArrayImplicit is called when entering the MutArrayImplicit production.
	EnterMutArrayImplicit(c *MutArrayImplicitContext)

	// EnterArrayType is called when entering the arrayType production.
	EnterArrayType(c *ArrayTypeContext)

	// EnterFixedArrayType is called when entering the fixedArrayType production.
	EnterFixedArrayType(c *FixedArrayTypeContext)

	// EnterArrayLiteralExprList is called when entering the ArrayLiteralExprList production.
	EnterArrayLiteralExprList(c *ArrayLiteralExprListContext)

	// EnterEmptyArr is called when entering the EmptyArr production.
	EnterEmptyArr(c *EmptyArrContext)

	// EnterArrayIndexAssign is called when entering the ArrayIndexAssign production.
	EnterArrayIndexAssign(c *ArrayIndexAssignContext)

	// EnterArrayReassign is called when entering the arrayReassign production.
	EnterArrayReassign(c *ArrayReassignContext)

	// EnterLenFunction is called when entering the lenFunction production.
	EnterLenFunction(c *LenFunctionContext)

	// EnterFunctionDecl is called when entering the functionDecl production.
	EnterFunctionDecl(c *FunctionDeclContext)

	// EnterParamList is called when entering the paramList production.
	EnterParamList(c *ParamListContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterReturnSignature is called when entering the returnSignature production.
	EnterReturnSignature(c *ReturnSignatureContext)

	// EnterReturnStmt is called when entering the returnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterExprList is called when entering the exprList production.
	EnterExprList(c *ExprListContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterWhileLoop is called when entering the whileLoop production.
	EnterWhileLoop(c *WhileLoopContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitPrintStmtStatement is called when exiting the PrintStmtStatement production.
	ExitPrintStmtStatement(c *PrintStmtStatementContext)

	// ExitVarDeclStatement is called when exiting the VarDeclStatement production.
	ExitVarDeclStatement(c *VarDeclStatementContext)

	// ExitAssignStmtStatement is called when exiting the AssignStmtStatement production.
	ExitAssignStmtStatement(c *AssignStmtStatementContext)

	// ExitIfStmtStatement is called when exiting the IfStmtStatement production.
	ExitIfStmtStatement(c *IfStmtStatementContext)

	// ExitArrayDeclStatement is called when exiting the ArrayDeclStatement production.
	ExitArrayDeclStatement(c *ArrayDeclStatementContext)

	// ExitArrayIndexAssignStatement is called when exiting the ArrayIndexAssignStatement production.
	ExitArrayIndexAssignStatement(c *ArrayIndexAssignStatementContext)

	// ExitArrayReassignStatement is called when exiting the ArrayReassignStatement production.
	ExitArrayReassignStatement(c *ArrayReassignStatementContext)

	// ExitFunctionDeclStatement is called when exiting the FunctionDeclStatement production.
	ExitFunctionDeclStatement(c *FunctionDeclStatementContext)

	// ExitReturnStmtStatement is called when exiting the ReturnStmtStatement production.
	ExitReturnStmtStatement(c *ReturnStmtStatementContext)

	// ExitFunctionCallStatement is called when exiting the FunctionCallStatement production.
	ExitFunctionCallStatement(c *FunctionCallStatementContext)

	// ExitWhileLoopStatement is called when exiting the WhileLoopStatement production.
	ExitWhileLoopStatement(c *WhileLoopStatementContext)

	// ExitPrintStatement is called when exiting the PrintStatement production.
	ExitPrintStatement(c *PrintStatementContext)

	// ExitMulExpr is called when exiting the MulExpr production.
	ExitMulExpr(c *MulExprContext)

	// ExitAndExpr is called when exiting the AndExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitLitExpr is called when exiting the LitExpr production.
	ExitLitExpr(c *LitExprContext)

	// ExitNeExpr is called when exiting the NeExpr production.
	ExitNeExpr(c *NeExprContext)

	// ExitSubExpr is called when exiting the SubExpr production.
	ExitSubExpr(c *SubExprContext)

	// ExitLtExpr is called when exiting the LtExpr production.
	ExitLtExpr(c *LtExprContext)

	// ExitGtExpr is called when exiting the GtExpr production.
	ExitGtExpr(c *GtExprContext)

	// ExitAddExpr is called when exiting the AddExpr production.
	ExitAddExpr(c *AddExprContext)

	// ExitGeExpr is called when exiting the GeExpr production.
	ExitGeExpr(c *GeExprContext)

	// ExitExpExpr is called when exiting the ExpExpr production.
	ExitExpExpr(c *ExpExprContext)

	// ExitIndexExpr is called when exiting the IndexExpr production.
	ExitIndexExpr(c *IndexExprContext)

	// ExitLeExpr is called when exiting the LeExpr production.
	ExitLeExpr(c *LeExprContext)

	// ExitOrExpr is called when exiting the OrExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitDivExpr is called when exiting the DivExpr production.
	ExitDivExpr(c *DivExprContext)

	// ExitFunctionCallExpr is called when exiting the FunctionCallExpr production.
	ExitFunctionCallExpr(c *FunctionCallExprContext)

	// ExitEqExpr is called when exiting the EqExpr production.
	ExitEqExpr(c *EqExprContext)

	// ExitVarExpr is called when exiting the VarExpr production.
	ExitVarExpr(c *VarExprContext)

	// ExitNotExpr is called when exiting the NotExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitModExpr is called when exiting the ModExpr production.
	ExitModExpr(c *ModExprContext)

	// ExitParenExpr is called when exiting the ParenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitLenFunctionExpr is called when exiting the LenFunctionExpr production.
	ExitLenFunctionExpr(c *LenFunctionExprContext)

	// ExitAssignStmt is called when exiting the AssignStmt production.
	ExitAssignStmt(c *AssignStmtContext)

	// ExitIntLiteral is called when exiting the IntLiteral production.
	ExitIntLiteral(c *IntLiteralContext)

	// ExitFloatLiteral is called when exiting the FloatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)

	// ExitTrueLiteral is called when exiting the TrueLiteral production.
	ExitTrueLiteral(c *TrueLiteralContext)

	// ExitFalseLiteral is called when exiting the FalseLiteral production.
	ExitFalseLiteral(c *FalseLiteralContext)

	// ExitStringLiteral is called when exiting the StringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitByteLiteral is called when exiting the ByteLiteral production.
	ExitByteLiteral(c *ByteLiteralContext)

	// ExitLetDecl is called when exiting the LetDecl production.
	ExitLetDecl(c *LetDeclContext)

	// ExitLetInferDecl is called when exiting the LetInferDecl production.
	ExitLetInferDecl(c *LetInferDeclContext)

	// ExitMutDecl is called when exiting the MutDecl production.
	ExitMutDecl(c *MutDeclContext)

	// ExitMutInferDecl is called when exiting the MutInferDecl production.
	ExitMutInferDecl(c *MutInferDeclContext)

	// ExitConstDecl is called when exiting the ConstDecl production.
	ExitConstDecl(c *ConstDeclContext)

	// ExitConstInferDecl is called when exiting the ConstInferDecl production.
	ExitConstInferDecl(c *ConstInferDeclContext)

	// ExitI8Type is called when exiting the I8Type production.
	ExitI8Type(c *I8TypeContext)

	// ExitI16Type is called when exiting the I16Type production.
	ExitI16Type(c *I16TypeContext)

	// ExitI32Type is called when exiting the I32Type production.
	ExitI32Type(c *I32TypeContext)

	// ExitI64Type is called when exiting the I64Type production.
	ExitI64Type(c *I64TypeContext)

	// ExitIntType is called when exiting the IntType production.
	ExitIntType(c *IntTypeContext)

	// ExitUI8Type is called when exiting the UI8Type production.
	ExitUI8Type(c *UI8TypeContext)

	// ExitUI16Type is called when exiting the UI16Type production.
	ExitUI16Type(c *UI16TypeContext)

	// ExitUI32Type is called when exiting the UI32Type production.
	ExitUI32Type(c *UI32TypeContext)

	// ExitUI64Type is called when exiting the UI64Type production.
	ExitUI64Type(c *UI64TypeContext)

	// ExitUIntType is called when exiting the UIntType production.
	ExitUIntType(c *UIntTypeContext)

	// ExitF32Type is called when exiting the F32Type production.
	ExitF32Type(c *F32TypeContext)

	// ExitF64Type is called when exiting the F64Type production.
	ExitF64Type(c *F64TypeContext)

	// ExitBoolType is called when exiting the BoolType production.
	ExitBoolType(c *BoolTypeContext)

	// ExitStringType is called when exiting the StringType production.
	ExitStringType(c *StringTypeContext)

	// ExitByteType is called when exiting the ByteType production.
	ExitByteType(c *ByteTypeContext)

	// ExitIfBlockOnly is called when exiting the IfBlockOnly production.
	ExitIfBlockOnly(c *IfBlockOnlyContext)

	// ExitIfElseChain is called when exiting the IfElseChain production.
	ExitIfElseChain(c *IfElseChainContext)

	// ExitElseIfBlock is called when exiting the elseIfBlock production.
	ExitElseIfBlock(c *ElseIfBlockContext)

	// ExitElseBlock is called when exiting the elseBlock production.
	ExitElseBlock(c *ElseBlockContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitArrayDecl is called when exiting the arrayDecl production.
	ExitArrayDecl(c *ArrayDeclContext)

	// ExitConstExplicitArrayDecl is called when exiting the ConstExplicitArrayDecl production.
	ExitConstExplicitArrayDecl(c *ConstExplicitArrayDeclContext)

	// ExitConstImplicitArrayDecl is called when exiting the ConstImplicitArrayDecl production.
	ExitConstImplicitArrayDecl(c *ConstImplicitArrayDeclContext)

	// ExitLetExplicitArrayDecl is called when exiting the LetExplicitArrayDecl production.
	ExitLetExplicitArrayDecl(c *LetExplicitArrayDeclContext)

	// ExitLetImplicitArrayDecl is called when exiting the LetImplicitArrayDecl production.
	ExitLetImplicitArrayDecl(c *LetImplicitArrayDeclContext)

	// ExitMutFixedArrayWithInit is called when exiting the MutFixedArrayWithInit production.
	ExitMutFixedArrayWithInit(c *MutFixedArrayWithInitContext)

	// ExitMutFixedArrayNoInit is called when exiting the MutFixedArrayNoInit production.
	ExitMutFixedArrayNoInit(c *MutFixedArrayNoInitContext)

	// ExitMutArrayExplicitWithInit is called when exiting the MutArrayExplicitWithInit production.
	ExitMutArrayExplicitWithInit(c *MutArrayExplicitWithInitContext)

	// ExitMutArrayExplicitNoInit is called when exiting the MutArrayExplicitNoInit production.
	ExitMutArrayExplicitNoInit(c *MutArrayExplicitNoInitContext)

	// ExitMutArrayImplicit is called when exiting the MutArrayImplicit production.
	ExitMutArrayImplicit(c *MutArrayImplicitContext)

	// ExitArrayType is called when exiting the arrayType production.
	ExitArrayType(c *ArrayTypeContext)

	// ExitFixedArrayType is called when exiting the fixedArrayType production.
	ExitFixedArrayType(c *FixedArrayTypeContext)

	// ExitArrayLiteralExprList is called when exiting the ArrayLiteralExprList production.
	ExitArrayLiteralExprList(c *ArrayLiteralExprListContext)

	// ExitEmptyArr is called when exiting the EmptyArr production.
	ExitEmptyArr(c *EmptyArrContext)

	// ExitArrayIndexAssign is called when exiting the ArrayIndexAssign production.
	ExitArrayIndexAssign(c *ArrayIndexAssignContext)

	// ExitArrayReassign is called when exiting the arrayReassign production.
	ExitArrayReassign(c *ArrayReassignContext)

	// ExitLenFunction is called when exiting the lenFunction production.
	ExitLenFunction(c *LenFunctionContext)

	// ExitFunctionDecl is called when exiting the functionDecl production.
	ExitFunctionDecl(c *FunctionDeclContext)

	// ExitParamList is called when exiting the paramList production.
	ExitParamList(c *ParamListContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitReturnSignature is called when exiting the returnSignature production.
	ExitReturnSignature(c *ReturnSignatureContext)

	// ExitReturnStmt is called when exiting the returnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitExprList is called when exiting the exprList production.
	ExitExprList(c *ExprListContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitWhileLoop is called when exiting the whileLoop production.
	ExitWhileLoop(c *WhileLoopContext)
}
