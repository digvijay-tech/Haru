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
		num, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			runtimeErr(fmt.Sprintf("parsing failed at literal: %s", text))
		}

		// size check for type assertion
		if num >= math.MinInt32 && num <= math.MaxInt32 {
			return Value{Value: text, Typ: "i32"}
		}

		// making sure the number is within 64bit integer width
		if num < math.MinInt64 || num > math.MaxInt64 {
			runtimeErr(fmt.Sprintf("number is too big: %s", text))
		}

		return Value{Value: text, Typ: "i64"}
	case isFloat(text):
		// parsing as f32
		if len(text) <= 39 {
			val, err := strconv.ParseFloat(text, 32)
			if err != nil {
				runtimeErr(fmt.Sprintf("parsing failed at literal: %s", text))
			}

			f32 := float32(val)
			return Value{Value: fmt.Sprintf("%f", f32), Typ: "f32"}
		}

		// parsing as f64
		if len(text) <= 309 {
			val, err := strconv.ParseFloat(text, 64)
			if err != nil {
				runtimeErr(fmt.Sprintf("parsing failed at literal: %s", text))
			}

			f64 := float64(val)
			return Value{Value: fmt.Sprintf("%f", f64), Typ: "f64"}
		}

		runtimeErr(fmt.Sprintf("number is too big for float64: %s", text))
		return nil // runtimeErr stopped execution but this is here to avoid syntax error
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

// VisitParenExpr evaluates the inner expression
func (v *HaruVisitor) VisitParenExpr(ctx *parser.ParenExprContext) any {
	return v.Visit(ctx.Expr())
}

/*******************************************************************************
***************************** ARITHMETIC EXPRESSIONS ***************************
********************************************************************************/

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

/*******************************************************************************
****************************** LOGICAL EXPRESSIONS *****************************
********************************************************************************/

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

/*******************************************************************************
**************************** COMPARISION EXPRESSIONS ***************************
********************************************************************************/

// VisitEqExpr evaluates Equals comparision operator (==)
func (v *HaruVisitor) VisitEqExpr(ctx *parser.EqExprContext) any {
	leftVal := v.Visit(ctx.Expr(0)).(Value)
	rightVal := v.Visit(ctx.Expr(1)).(Value)

	// type compatibility check
	if leftVal.Typ != rightVal.Typ {
		log.Fatalf("Type error: cannot compare types '%s' and '%s' using '=='", leftVal.Typ, rightVal.Typ)
	}

	result := leftVal.Value == rightVal.Value
	return Value{Value: fmt.Sprintf("%v", result), Typ: "bool"}
}

// VisitNeExpr evaluates Not Equals comparision operator (!=)
func (v *HaruVisitor) VisitNeExpr(ctx *parser.NeExprContext) any {
	leftVal := v.Visit(ctx.Expr(0)).(Value)
	rightVal := v.Visit(ctx.Expr(1)).(Value)

	// type compatibility check
	if leftVal.Typ != rightVal.Typ {
		log.Fatalf("Type error: cannot compare types '%s' and '%s' using '!='", leftVal.Typ, rightVal.Typ)
	}

	result := leftVal.Value != rightVal.Value
	return Value{Value: fmt.Sprintf("%v", result), Typ: "bool"}
}

// VisitLtExpr evaluates Less Than comparision operator (<)
func (v *HaruVisitor) VisitLtExpr(ctx *parser.LtExprContext) any {
	leftVal := v.Visit(ctx.Expr(0)).(Value)
	rightVal := v.Visit(ctx.Expr(1)).(Value)

	// strict type compatibility check
	if leftVal.Typ != rightVal.Typ {
		log.Fatalf("Type error: cannot compare types '%s' and '%s' using '<'", leftVal.Typ, rightVal.Typ)
	}

	// only numeric comparision is allowed
	if isNumericType(leftVal.Typ) {
		// converting to f64 to safely represent all integer and float variants
		l, err1 := convertToFloat64(leftVal.Value)
		r, err2 := convertToFloat64(rightVal.Value)

		if err1 != nil || err2 != nil {
			log.Fatalf("Type error: cannot convert '%s' and '%s' to numbers for comparison", leftVal.Typ, rightVal.Typ)
		}

		return Value{Value: fmt.Sprintf("%v", l < r), Typ: "bool"}
	}

	log.Fatalf("Type error: '<' not supported between values of type '%s'", leftVal.Typ)
	return Value{Value: "", Typ: "unknown"}
}

// VisitLeExpr evaluates Less Than Equals comparision operator (<=)
func (v *HaruVisitor) VisitLeExpr(ctx *parser.LeExprContext) any {
	leftVal := v.Visit(ctx.Expr(0)).(Value)
	rightVal := v.Visit(ctx.Expr(1)).(Value)

	// strict type compatibility check
	if leftVal.Typ != rightVal.Typ {
		log.Fatalf("Type error: cannot compare types '%s' and '%s' using '<='", leftVal.Typ, rightVal.Typ)
	}

	// only numeric comparision is allowed
	if isNumericType(leftVal.Typ) {
		// converting to f64 to safely represent all integer and float variants
		l, err1 := convertToFloat64(leftVal.Value)
		r, err2 := convertToFloat64(rightVal.Value)

		if err1 != nil || err2 != nil {
			log.Fatalf("Type error: cannot convert '%s' and '%s' to numbers for comparison", leftVal.Typ, rightVal.Typ)
		}

		return Value{Value: fmt.Sprintf("%v", l <= r), Typ: "bool"}
	}

	log.Fatalf("Type error: '<=' not supported between values of type '%s'", leftVal.Typ)
	return Value{Value: "", Typ: "unknown"}
}

// VisitGtExpr evaluates Greater Than comparision operator (>)
func (v *HaruVisitor) VisitGtExpr(ctx *parser.GtExprContext) any {
	leftVal := v.Visit(ctx.Expr(0)).(Value)
	rightVal := v.Visit(ctx.Expr(1)).(Value)

	// strict type compatibility check
	if leftVal.Typ != rightVal.Typ {
		log.Fatalf("Type error: cannot compare types '%s' and '%s' using '>'", leftVal.Typ, rightVal.Typ)
	}

	// only numeric comparision is allowed
	if isNumericType(leftVal.Typ) {
		// converting to f64 to safely represent all integer and float variants
		l, err1 := convertToFloat64(leftVal.Value)
		r, err2 := convertToFloat64(rightVal.Value)

		if err1 != nil || err2 != nil {
			log.Fatalf("Type error: cannot convert '%s' and '%s' to numbers for comparison", leftVal.Typ, rightVal.Typ)
		}

		return Value{Value: fmt.Sprintf("%v", l > r), Typ: "bool"}
	}

	log.Fatalf("Type error: '>' not supported between values of type '%s'", leftVal.Typ)
	return Value{Value: "", Typ: "unknown"}
}

// VisitGeExpr evaluates Greater Than Equals comparision operator (>=)
func (v *HaruVisitor) VisitGeExpr(ctx *parser.GeExprContext) any {
	leftVal := v.Visit(ctx.Expr(0)).(Value)
	rightVal := v.Visit(ctx.Expr(1)).(Value)

	// strict type compatibility check
	if leftVal.Typ != rightVal.Typ {
		log.Fatalf("Type error: cannot compare types '%s' and '%s' using '>='", leftVal.Typ, rightVal.Typ)
	}

	// only numeric comparision is allowed
	if isNumericType(leftVal.Typ) {
		// converting to f64 to safely represent all integer and float variants
		l, err1 := convertToFloat64(leftVal.Value)
		r, err2 := convertToFloat64(rightVal.Value)

		if err1 != nil || err2 != nil {
			log.Fatalf("Type error: cannot convert '%s' and '%s' to numbers for comparison", leftVal.Typ, rightVal.Typ)
		}

		return Value{Value: fmt.Sprintf("%v", l >= r), Typ: "bool"}
	}

	log.Fatalf("Type error: '>=' not supported between values of type '%s'", leftVal.Typ)
	return Value{Value: "", Typ: "unknown"}
}
