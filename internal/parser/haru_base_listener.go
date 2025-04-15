// Code generated from ./grammar/haru.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // haru

import "github.com/antlr4-go/antlr/v4"

// BaseharuListener is a complete listener for a parse tree produced by haruParser.
type BaseharuListener struct{}

var _ haruListener = &BaseharuListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseharuListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseharuListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseharuListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseharuListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseharuListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseharuListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseharuListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseharuListener) ExitStatement(ctx *StatementContext) {}

// EnterPrintStatement is called when production PrintStatement is entered.
func (s *BaseharuListener) EnterPrintStatement(ctx *PrintStatementContext) {}

// ExitPrintStatement is called when production PrintStatement is exited.
func (s *BaseharuListener) ExitPrintStatement(ctx *PrintStatementContext) {}

// EnterMulExpr is called when production MulExpr is entered.
func (s *BaseharuListener) EnterMulExpr(ctx *MulExprContext) {}

// ExitMulExpr is called when production MulExpr is exited.
func (s *BaseharuListener) ExitMulExpr(ctx *MulExprContext) {}

// EnterAndExpr is called when production AndExpr is entered.
func (s *BaseharuListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production AndExpr is exited.
func (s *BaseharuListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterLitExpr is called when production LitExpr is entered.
func (s *BaseharuListener) EnterLitExpr(ctx *LitExprContext) {}

// ExitLitExpr is called when production LitExpr is exited.
func (s *BaseharuListener) ExitLitExpr(ctx *LitExprContext) {}

// EnterNeExpr is called when production NeExpr is entered.
func (s *BaseharuListener) EnterNeExpr(ctx *NeExprContext) {}

// ExitNeExpr is called when production NeExpr is exited.
func (s *BaseharuListener) ExitNeExpr(ctx *NeExprContext) {}

// EnterSubExpr is called when production SubExpr is entered.
func (s *BaseharuListener) EnterSubExpr(ctx *SubExprContext) {}

// ExitSubExpr is called when production SubExpr is exited.
func (s *BaseharuListener) ExitSubExpr(ctx *SubExprContext) {}

// EnterLtExpr is called when production LtExpr is entered.
func (s *BaseharuListener) EnterLtExpr(ctx *LtExprContext) {}

// ExitLtExpr is called when production LtExpr is exited.
func (s *BaseharuListener) ExitLtExpr(ctx *LtExprContext) {}

// EnterGtExpr is called when production GtExpr is entered.
func (s *BaseharuListener) EnterGtExpr(ctx *GtExprContext) {}

// ExitGtExpr is called when production GtExpr is exited.
func (s *BaseharuListener) ExitGtExpr(ctx *GtExprContext) {}

// EnterAddExpr is called when production AddExpr is entered.
func (s *BaseharuListener) EnterAddExpr(ctx *AddExprContext) {}

// ExitAddExpr is called when production AddExpr is exited.
func (s *BaseharuListener) ExitAddExpr(ctx *AddExprContext) {}

// EnterGeExpr is called when production GeExpr is entered.
func (s *BaseharuListener) EnterGeExpr(ctx *GeExprContext) {}

// ExitGeExpr is called when production GeExpr is exited.
func (s *BaseharuListener) ExitGeExpr(ctx *GeExprContext) {}

// EnterExpExpr is called when production ExpExpr is entered.
func (s *BaseharuListener) EnterExpExpr(ctx *ExpExprContext) {}

// ExitExpExpr is called when production ExpExpr is exited.
func (s *BaseharuListener) ExitExpExpr(ctx *ExpExprContext) {}

// EnterLeExpr is called when production LeExpr is entered.
func (s *BaseharuListener) EnterLeExpr(ctx *LeExprContext) {}

// ExitLeExpr is called when production LeExpr is exited.
func (s *BaseharuListener) ExitLeExpr(ctx *LeExprContext) {}

// EnterOrExpr is called when production OrExpr is entered.
func (s *BaseharuListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production OrExpr is exited.
func (s *BaseharuListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterArrayExpr is called when production ArrayExpr is entered.
func (s *BaseharuListener) EnterArrayExpr(ctx *ArrayExprContext) {}

// ExitArrayExpr is called when production ArrayExpr is exited.
func (s *BaseharuListener) ExitArrayExpr(ctx *ArrayExprContext) {}

// EnterDivExpr is called when production DivExpr is entered.
func (s *BaseharuListener) EnterDivExpr(ctx *DivExprContext) {}

// ExitDivExpr is called when production DivExpr is exited.
func (s *BaseharuListener) ExitDivExpr(ctx *DivExprContext) {}

// EnterEqExpr is called when production EqExpr is entered.
func (s *BaseharuListener) EnterEqExpr(ctx *EqExprContext) {}

// ExitEqExpr is called when production EqExpr is exited.
func (s *BaseharuListener) ExitEqExpr(ctx *EqExprContext) {}

// EnterVarExpr is called when production VarExpr is entered.
func (s *BaseharuListener) EnterVarExpr(ctx *VarExprContext) {}

// ExitVarExpr is called when production VarExpr is exited.
func (s *BaseharuListener) ExitVarExpr(ctx *VarExprContext) {}

// EnterNotExpr is called when production NotExpr is entered.
func (s *BaseharuListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production NotExpr is exited.
func (s *BaseharuListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterModExpr is called when production ModExpr is entered.
func (s *BaseharuListener) EnterModExpr(ctx *ModExprContext) {}

// ExitModExpr is called when production ModExpr is exited.
func (s *BaseharuListener) ExitModExpr(ctx *ModExprContext) {}

// EnterParenExpr is called when production ParenExpr is entered.
func (s *BaseharuListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production ParenExpr is exited.
func (s *BaseharuListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterAssignStmt is called when production AssignStmt is entered.
func (s *BaseharuListener) EnterAssignStmt(ctx *AssignStmtContext) {}

// ExitAssignStmt is called when production AssignStmt is exited.
func (s *BaseharuListener) ExitAssignStmt(ctx *AssignStmtContext) {}

// EnterIntLiteral is called when production IntLiteral is entered.
func (s *BaseharuListener) EnterIntLiteral(ctx *IntLiteralContext) {}

// ExitIntLiteral is called when production IntLiteral is exited.
func (s *BaseharuListener) ExitIntLiteral(ctx *IntLiteralContext) {}

// EnterFloatLiteral is called when production FloatLiteral is entered.
func (s *BaseharuListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production FloatLiteral is exited.
func (s *BaseharuListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterTrueLiteral is called when production TrueLiteral is entered.
func (s *BaseharuListener) EnterTrueLiteral(ctx *TrueLiteralContext) {}

// ExitTrueLiteral is called when production TrueLiteral is exited.
func (s *BaseharuListener) ExitTrueLiteral(ctx *TrueLiteralContext) {}

// EnterFalseLiteral is called when production FalseLiteral is entered.
func (s *BaseharuListener) EnterFalseLiteral(ctx *FalseLiteralContext) {}

// ExitFalseLiteral is called when production FalseLiteral is exited.
func (s *BaseharuListener) ExitFalseLiteral(ctx *FalseLiteralContext) {}

// EnterStringLiteral is called when production StringLiteral is entered.
func (s *BaseharuListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production StringLiteral is exited.
func (s *BaseharuListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterByteLiteral is called when production ByteLiteral is entered.
func (s *BaseharuListener) EnterByteLiteral(ctx *ByteLiteralContext) {}

// ExitByteLiteral is called when production ByteLiteral is exited.
func (s *BaseharuListener) ExitByteLiteral(ctx *ByteLiteralContext) {}
