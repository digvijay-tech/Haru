// Evaluating Expressions
package interpreter

import "github.com/digvijay-tech/Haru/internal/parser"

// Temp
// VisitLitExpr handles literal expressions
func (v *HaruVisitor) VisitLitExpr(ctx *parser.LitExprContext) interface{} {
	text := ctx.GetText()

	switch {
	case isInt(text):
		return Value{Value: text, Typ: "int"}
	case isFloat(text):
		return Value{Value: text, Typ: "float"}
	case text == "true" || text == "false":
		return Value{Value: text, Typ: "bool"}
	case isByte(text):
		return Value{Value: text, Typ: "byte"}
	case isString(text):
		return Value{Value: stripQuotes(text), Typ: "string"}
	default:
		return Value{Value: text, Typ: "unknown"}
	}
}
