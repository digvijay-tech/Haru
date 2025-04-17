// Reads and evaluates identifiers
package interpreter

import (
	"github.com/digvijay-tech/Haru/internal/parser"
)

func (v *HaruVisitor) VisitVarExpr(ctx *parser.VarExprContext) any {
	name := ctx.ID().GetText()

	if value, ok := v.symbolTable[name]; ok {
		return value
	}

	return nil
}
