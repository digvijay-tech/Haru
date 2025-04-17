package interpreter

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/digvijay-tech/Haru/internal/parser"
)

// VisitVarDeclStatement passes control to actual const, let, and mut declaration handlers
func (v *HaruVisitor) VisitVarDeclStatement(ctx *parser.VarDeclStatementContext) any {
	decl := ctx.GetChild(0).(antlr.ParseTree)
	return v.Visit(decl)
}
