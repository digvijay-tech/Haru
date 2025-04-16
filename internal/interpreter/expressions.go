// Evaluating Expressions
package interpreter

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/digvijay-tech/Haru/internal/parser"
)

// VisitLitExpr handles literal expressions
func (v *HaruVisitor) VisitLitExpr(ctx *parser.LitExprContext) any {
	text := ctx.GetText()

	switch {
	case isInt(text):
		return Value{Value: text, Typ: "i32"} // same as Rust
	case isFloat(text):
		return Value{Value: text, Typ: "f64"} // same as Rust
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

// VisitAddExpr evaluates addition expressions
func (v *HaruVisitor) VisitAddExpr(ctx *parser.AddExprContext) any {
	leftVal := v.Visit(ctx.Expr(0)).(Value)
	rightVal := v.Visit(ctx.Expr(1)).(Value)

	// handles string concatenation
	if leftVal.Typ == "string" && rightVal.Typ == "string" {
		return Value{Value: leftVal.Value + rightVal.Value, Typ: "string"}
	}

	// resolves type promotion for numeric type only
	// non-numeric types will result in fatal
	resultType := promoteType(leftVal.Typ, rightVal.Typ)

	// convert to float64 for uniform calculation
	l, err1 := convertToFloat64(leftVal.Value)
	r, err2 := convertToFloat64(rightVal.Value)

	if err1 != nil || err2 != nil {
		log.Fatalf("Type error: cannot convert operands '%s' and '%s' to numbers", leftVal.Value, rightVal.Value)
	}

	sum := l + r
	return Value{Value: fmt.Sprintf("%v", sum), Typ: resultType}
}

// VisitSubExpr evaluates subtraction expressions
func (v *HaruVisitor) VisitSubExpr(ctx *parser.SubExprContext) any {
	leftVal := v.Visit(ctx.Expr(0)).(Value)
	rightVal := v.Visit(ctx.Expr(1)).(Value)

	// resolves type promotion
	resultType := promoteType(leftVal.Typ, rightVal.Typ)

	// convert to float64 for uniform calculation
	l, err1 := convertToFloat64(leftVal.Value)
	r, err2 := convertToFloat64(rightVal.Value)

	if err1 != nil || err2 != nil {
		log.Fatalf("Type error: cannot convert operands '%s' and '%s' to numbers", leftVal.Value, rightVal.Value)
	}

	difference := l - r
	return Value{Value: fmt.Sprintf("%v", difference), Typ: resultType}
}

// VisitMulExpr evaluates multiplication expressions
func (v *HaruVisitor) VisitMulExpr(ctx *parser.MulExprContext) any {
	leftVal := v.Visit(ctx.Expr(0)).(Value)
	rightVal := v.Visit(ctx.Expr(1)).(Value)

	// resolves type promotion
	resultType := promoteType(leftVal.Typ, rightVal.Typ)

	// convert to float64 for uniform calculation
	l, err1 := convertToFloat64(leftVal.Value)
	r, err2 := convertToFloat64(rightVal.Value)

	if err1 != nil || err2 != nil {
		log.Fatalf("Type error: cannot convert operands '%s' and '%s' to numbers", leftVal.Value, rightVal.Value)
	}

	product := l * r
	return Value{Value: fmt.Sprintf("%v", product), Typ: resultType}
}

// VisitDivExpr evaluates division expressions
func (v *HaruVisitor) VisitDivExpr(ctx *parser.DivExprContext) any {
	leftVal := v.Visit(ctx.Expr(0)).(Value)
	rightVal := v.Visit(ctx.Expr(1)).(Value)

	// resolves type promotion
	resultType := promoteType(leftVal.Typ, rightVal.Typ)

	// convert to float64 for uniform calculation
	l, err1 := convertToFloat64(leftVal.Value)
	r, err2 := convertToFloat64(rightVal.Value)

	if err1 != nil || err2 != nil {
		log.Fatalf("Type error: cannot convert operands '%s' and '%s' to numbers", leftVal.Value, rightVal.Value)
	}

	if r == 0 {
		log.Fatal("Runtime error: division by zero")
	}

	quotient := l / r
	return Value{Value: fmt.Sprintf("%v", quotient), Typ: resultType}
}

// VisitModExpr evaluates modulus expressions
func (v *HaruVisitor) VisitModExpr(ctx *parser.ModExprContext) any {
	leftVal := v.Visit(ctx.Expr(0)).(Value)
	rightVal := v.Visit(ctx.Expr(1)).(Value)

	// resolves type promotion
	resultType := promoteType(leftVal.Typ, rightVal.Typ)

	// convert to float64 for uniform calculation
	l, err1 := convertToFloat64(leftVal.Value)
	r, err2 := convertToFloat64(rightVal.Value)

	if err1 != nil || err2 != nil {
		log.Fatalf("Type error: cannot convert operands '%s' and '%s' to numbers", leftVal.Value, rightVal.Value)
	}

	if r == 0 {
		log.Fatal("Runtime error: modulo by zero")
	}

	mod := float64(int64(l) % int64(r))
	return Value{Value: fmt.Sprintf("%v", mod), Typ: resultType}
}

// VisitExpExpr evaluates exponent expressions
// TODO: Exponentiation (**) is left-associative, so 2 ** 3 ** 2 = 64
// Should be right-associative to return 512 or support pow(x, y) instead
// Revisit this later to implement right-associativity
func (v *HaruVisitor) VisitExpExpr(ctx *parser.ExpExprContext) any {
	leftVal := v.Visit(ctx.Expr(0)).(Value)
	rightVal := v.Visit(ctx.Expr(1)).(Value)

	// resolves type promotion
	resultType := promoteType(leftVal.Typ, rightVal.Typ)

	// convert to float64 for uniform calculation
	l, err1 := convertToFloat64(leftVal.Value)
	r, err2 := convertToFloat64(rightVal.Value)

	if err1 != nil || err2 != nil {
		log.Fatalf("Type error: cannot convert operands '%s' and '%s' to numbers", leftVal.Value, rightVal.Value)
	}

	exp := math.Pow(l, r)
	return Value{Value: fmt.Sprintf("%v", exp), Typ: resultType}
}

// VisitParenExpr evaluates the inner expression
func (v *HaruVisitor) VisitParenExpr(ctx *parser.ParenExprContext) any {
	return v.Visit(ctx.Expr())
}

// VisitNotExpr evaluates logical NOT (!)
func (v *HaruVisitor) VisitNotExpr(ctx *parser.NotExprContext) any {
	val := v.Visit(ctx.Expr()).(Value)

	if val.Typ != "bool" {
		log.Fatalf("Type error: '!' operator can only be applied to boolean, got '%s'", val.Typ)
	}

	boolVal, err := strconv.ParseBool(val.Value)
	if err != nil {
		log.Fatalf("Runtime error: invalid boolean value '%s'", val.Value)
	}

	return Value{
		Value: fmt.Sprintf("%v", !boolVal),
		Typ:   "bool",
	}
}

// VisitAndExpr evaluates logical AND (&&)
func (v *HaruVisitor) VisitAndExpr(ctx *parser.AndExprContext) any {
	leftVal := v.Visit(ctx.Expr(0)).(Value)
	rightVal := v.Visit(ctx.Expr(1)).(Value)

	if leftVal.Typ != "bool" || rightVal.Typ != "bool" {
		log.Fatalf("Logical AND requires boolean operands, got %s and %s", leftVal.Typ, rightVal.Typ)
	}

	l, _ := strconv.ParseBool(leftVal.Value)
	r, _ := strconv.ParseBool(rightVal.Value)

	return Value{Value: fmt.Sprintf("%v", l && r), Typ: "bool"}
}

// VisitOrExpr evaluates logical OR (||)
func (v *HaruVisitor) VisitOrExpr(ctx *parser.OrExprContext) any {
	leftVal := v.Visit(ctx.Expr(0)).(Value)
	rightVal := v.Visit(ctx.Expr(1)).(Value)

	if leftVal.Typ != "bool" || rightVal.Typ != "bool" {
		log.Fatalf("Logical OR requires boolean operands, got %s and %s", leftVal.Typ, rightVal.Typ)
	}

	l, _ := strconv.ParseBool(leftVal.Value)
	r, _ := strconv.ParseBool(rightVal.Value)

	return Value{Value: fmt.Sprintf("%v", l || r), Typ: "bool"}
}
