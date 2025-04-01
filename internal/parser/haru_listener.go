// Code generated from haru.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // haru

import "github.com/antlr4-go/antlr/v4"

// haruListener is a complete listener for a parse tree produced by haruParser.
type haruListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

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

	// EnterUIType is called when entering the UIType production.
	EnterUIType(c *UITypeContext)

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

	// EnterLeExpr is called when entering the LeExpr production.
	EnterLeExpr(c *LeExprContext)

	// EnterOrExpr is called when entering the OrExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterDivExpr is called when entering the DivExpr production.
	EnterDivExpr(c *DivExprContext)

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

	// EnterAssignStmt is called when entering the AssignStmt production.
	EnterAssignStmt(c *AssignStmtContext)

	// EnterPrintStatement is called when entering the PrintStatement production.
	EnterPrintStatement(c *PrintStatementContext)

	// EnterIfStatement is called when entering the IfStatement production.
	EnterIfStatement(c *IfStatementContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

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

	// ExitUIType is called when exiting the UIType production.
	ExitUIType(c *UITypeContext)

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

	// ExitLeExpr is called when exiting the LeExpr production.
	ExitLeExpr(c *LeExprContext)

	// ExitOrExpr is called when exiting the OrExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitDivExpr is called when exiting the DivExpr production.
	ExitDivExpr(c *DivExprContext)

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

	// ExitAssignStmt is called when exiting the AssignStmt production.
	ExitAssignStmt(c *AssignStmtContext)

	// ExitPrintStatement is called when exiting the PrintStatement production.
	ExitPrintStatement(c *PrintStatementContext)

	// ExitIfStatement is called when exiting the IfStatement production.
	ExitIfStatement(c *IfStatementContext)
}
