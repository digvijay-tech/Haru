// Reads and evaluates identifiers
package interpreter

import (
	"fmt"

	"github.com/digvijay-tech/Haru/internal/parser"
)

// VisitVarExpr looks up a variable by name in the symbol table and returns its value, or raises a runtime error if it's undefined
func (v *HaruVisitor) VisitVarExpr(ctx *parser.VarExprContext) any {
	name := ctx.ID().GetText()

	if value, ok := v.resolve(name); ok {
		return value
	}

	runtimeErr(fmt.Sprintf("undefined variable '%s'", name))
	return nil // don't need this but have to keep it here for syntax warning
}
