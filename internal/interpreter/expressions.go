package interpreter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/digvijay-tech/Haru/internal/parser"
)

// evalExpr evaluates an expression node and returns its value and type
func (v *HaruVisitor) evalExpr(ctx parser.IExprContext) (string, string) {
	switch ctx.(type) {
	case *parser.NotExprContext:
		not := ctx.(*parser.NotExprContext)
		val, typ := v.evalExpr(not.Expr())
		if typ != "bool" {
			fmt.Println("Error: Logical NOT (!) requires a boolean operand")
			return "false", "bool"
		}
		boolVal, _ := strconv.ParseBool(val)
		return strconv.FormatBool(!boolVal), "bool"

	case *parser.ParenExprContext:
		return v.evalExpr(ctx.(*parser.ParenExprContext).Expr())

	case *parser.ExpExprContext:
		exp := ctx.(*parser.ExpExprContext)
		leftVal, leftTyp := v.evalExpr(exp.Expr(0))
		rightVal, rightTyp := v.evalExpr(exp.Expr(1))
		if !isIntegerType(leftTyp) || !isIntegerType(rightTyp) {
			fmt.Printf("Error: Exponentiation (**) requires integer operands, got %s and %s\n", leftTyp, rightTyp)
			return "0", "unknown"
		}
		resultTyp := promoteNumericTypes(leftTyp, rightTyp)
		left, _ := strconv.ParseInt(leftVal, 10, 64)
		right, _ := strconv.ParseInt(rightVal, 10, 64)
		if right < 0 {
			fmt.Println("Error: Negative exponents not supported")
			return "0", resultTyp
		}
		result := int64(1)
		for i := int64(0); i < right; i++ {
			result *= left
		}
		return strconv.FormatInt(result, 10), resultTyp

	case *parser.MulExprContext:
		mul := ctx.(*parser.MulExprContext)
		leftVal, leftTyp := v.evalExpr(mul.Expr(0))
		rightVal, rightTyp := v.evalExpr(mul.Expr(1))
		if !isNumericType(leftTyp) || !isNumericType(rightTyp) {
			fmt.Printf("Error: Multiplication (*) requires numeric operands, got %s and %s\n", leftTyp, rightTyp)
			return "0", "unknown"
		}
		resultTyp := promoteNumericTypes(leftTyp, rightTyp)
		if isFloatType(resultTyp) {
			left, _ := strconv.ParseFloat(leftVal, 64)
			right, _ := strconv.ParseFloat(rightVal, 64)
			return strconv.FormatFloat(left*right, 'f', -1, 64), resultTyp
		}
		left, _ := strconv.ParseInt(leftVal, 10, 64)
		right, _ := strconv.ParseInt(rightVal, 10, 64)
		return strconv.FormatInt(left*right, 10), resultTyp

	case *parser.DivExprContext:
		div := ctx.(*parser.DivExprContext)
		leftVal, leftTyp := v.evalExpr(div.Expr(0))
		rightVal, rightTyp := v.evalExpr(div.Expr(1))
		if !isNumericType(leftTyp) || !isNumericType(rightTyp) {
			fmt.Printf("Error: Division (/) requires numeric operands, got %s and %s\n", leftTyp, rightTyp)
			return "0", "unknown"
		}
		leftFloat, _ := strconv.ParseFloat(leftVal, 64)
		rightFloat, _ := strconv.ParseFloat(rightVal, 64)
		if rightFloat == 0 {
			fmt.Println("Error: Division by zero")
			return "0.0", "f64"
		}
		return strconv.FormatFloat(leftFloat/rightFloat, 'f', -1, 64), "f64"

	case *parser.ModExprContext:
		mod := ctx.(*parser.ModExprContext)
		leftVal, leftTyp := v.evalExpr(mod.Expr(0))
		rightVal, rightTyp := v.evalExpr(mod.Expr(1))
		if !isIntegerType(leftTyp) || !isIntegerType(rightTyp) {
			fmt.Printf("Error: Modulus (%%) requires integer operands, got %s and %s\n", leftTyp, rightTyp)
			return "0", "unknown"
		}
		resultTyp := promoteNumericTypes(leftTyp, rightTyp)
		left, _ := strconv.ParseInt(leftVal, 10, 64)
		right, _ := strconv.ParseInt(rightVal, 10, 64)
		if right == 0 {
			fmt.Println("Error: Modulus by zero")
			return "0", resultTyp
		}
		return strconv.FormatInt(left%right, 10), resultTyp

	case *parser.AddExprContext:
		add := ctx.(*parser.AddExprContext)
		leftVal, leftTyp := v.evalExpr(add.Expr(0))
		rightVal, rightTyp := v.evalExpr(add.Expr(1))
		if !isNumericType(leftTyp) || !isNumericType(rightTyp) {
			fmt.Printf("Error: Addition (+) requires numeric operands, got %s and %s\n", leftTyp, rightTyp)
			return "0", "unknown"
		}
		resultTyp := promoteNumericTypes(leftTyp, rightTyp)
		if isFloatType(resultTyp) {
			left, _ := strconv.ParseFloat(leftVal, 64)
			right, _ := strconv.ParseFloat(rightVal, 64)
			return strconv.FormatFloat(left+right, 'f', -1, 64), resultTyp
		}
		left, _ := strconv.ParseInt(leftVal, 10, 64)
		right, _ := strconv.ParseInt(rightVal, 10, 64)
		return strconv.FormatInt(left+right, 10), resultTyp

	case *parser.SubExprContext:
		sub := ctx.(*parser.SubExprContext)
		leftVal, leftTyp := v.evalExpr(sub.Expr(0))
		rightVal, rightTyp := v.evalExpr(sub.Expr(1))
		if !isNumericType(leftTyp) || !isNumericType(rightTyp) {
			fmt.Printf("Error: Subtraction (-) requires numeric operands, got %s and %s\n", leftTyp, rightTyp)
			return "0", "unknown"
		}
		resultTyp := promoteNumericTypes(leftTyp, rightTyp)
		if isFloatType(resultTyp) {
			left, _ := strconv.ParseFloat(leftVal, 64)
			right, _ := strconv.ParseFloat(rightVal, 64)
			return strconv.FormatFloat(left-right, 'f', -1, 64), resultTyp
		}
		left, _ := strconv.ParseInt(leftVal, 10, 64)
		right, _ := strconv.ParseInt(rightVal, 10, 64)
		return strconv.FormatInt(left-right, 10), resultTyp

	case *parser.LtExprContext:
		lt := ctx.(*parser.LtExprContext)
		leftVal, leftTyp := v.evalExpr(lt.Expr(0))
		rightVal, rightTyp := v.evalExpr(lt.Expr(1))
		if !isNumericType(leftTyp) || !isNumericType(rightTyp) {
			fmt.Printf("Error: Comparison (<) requires numeric operands, got %s and %s\n", leftTyp, rightTyp)
			return "false", "bool"
		}
		if isFloatType(leftTyp) || isFloatType(rightTyp) {
			left, _ := strconv.ParseFloat(leftVal, 64)
			right, _ := strconv.ParseFloat(rightVal, 64)
			return strconv.FormatBool(left < right), "bool"
		}
		left, _ := strconv.ParseInt(leftVal, 10, 64)
		right, _ := strconv.ParseInt(rightVal, 10, 64)
		return strconv.FormatBool(left < right), "bool"

	case *parser.GtExprContext:
		gt := ctx.(*parser.GtExprContext)
		leftVal, leftTyp := v.evalExpr(gt.Expr(0))
		rightVal, rightTyp := v.evalExpr(gt.Expr(1))
		if !isNumericType(leftTyp) || !isNumericType(rightTyp) {
			fmt.Printf("Error: Comparison (>) requires numeric operands, got %s and %s\n", leftTyp, rightTyp)
			return "false", "bool"
		}
		if isFloatType(leftTyp) || isFloatType(rightTyp) {
			left, _ := strconv.ParseFloat(leftVal, 64)
			right, _ := strconv.ParseFloat(rightVal, 64)
			return strconv.FormatBool(left > right), "bool"
		}
		left, _ := strconv.ParseInt(leftVal, 10, 64)
		right, _ := strconv.ParseInt(rightVal, 10, 64)
		return strconv.FormatBool(left > right), "bool"

	case *parser.LeExprContext:
		le := ctx.(*parser.LeExprContext)
		leftVal, leftTyp := v.evalExpr(le.Expr(0))
		rightVal, rightTyp := v.evalExpr(le.Expr(1))
		if !isNumericType(leftTyp) || !isNumericType(rightTyp) {
			fmt.Printf("Error: Comparison (<=) requires numeric operands, got %s and %s\n", leftTyp, rightTyp)
			return "false", "bool"
		}
		if isFloatType(leftTyp) || isFloatType(rightTyp) {
			left, _ := strconv.ParseFloat(leftVal, 64)
			right, _ := strconv.ParseFloat(rightVal, 64)
			return strconv.FormatBool(left <= right), "bool"
		}
		left, _ := strconv.ParseInt(leftVal, 10, 64)
		right, _ := strconv.ParseInt(rightVal, 10, 64)
		return strconv.FormatBool(left <= right), "bool"

	case *parser.GeExprContext:
		ge := ctx.(*parser.GeExprContext)
		leftVal, leftTyp := v.evalExpr(ge.Expr(0))
		rightVal, rightTyp := v.evalExpr(ge.Expr(1))
		if !isNumericType(leftTyp) || !isNumericType(rightTyp) {
			fmt.Printf("Error: Comparison (>=) requires numeric operands, got %s and %s\n", leftTyp, rightTyp)
			return "false", "bool"
		}
		if isFloatType(leftTyp) || isFloatType(rightTyp) {
			left, _ := strconv.ParseFloat(leftVal, 64)
			right, _ := strconv.ParseFloat(rightVal, 64)
			return strconv.FormatBool(left >= right), "bool"
		}
		left, _ := strconv.ParseInt(leftVal, 10, 64)
		right, _ := strconv.ParseInt(rightVal, 10, 64)
		return strconv.FormatBool(left >= right), "bool"

	case *parser.EqExprContext:
		eq := ctx.(*parser.EqExprContext)
		leftVal, leftTyp := v.evalExpr(eq.Expr(0))
		rightVal, rightTyp := v.evalExpr(eq.Expr(1))
		if leftTyp != rightTyp {
			fmt.Printf("Error: Equality (==) requires matching types, got %s and %s\n", leftTyp, rightTyp)
			return "false", "bool"
		}
		switch leftTyp {
		case "i8", "i16", "i32", "i64", "int", "ui8", "ui16", "ui32", "ui64", "ui":
			left, _ := strconv.ParseInt(leftVal, 10, 64)
			right, _ := strconv.ParseInt(rightVal, 10, 64)
			return strconv.FormatBool(left == right), "bool"
		case "f32", "f64":
			left, _ := strconv.ParseFloat(leftVal, 64)
			right, _ := strconv.ParseFloat(rightVal, 64)
			return strconv.FormatBool(left == right), "bool"
		case "bool":
			left, _ := strconv.ParseBool(leftVal)
			right, _ := strconv.ParseBool(rightVal)
			return strconv.FormatBool(left == right), "bool"
		case "string":
			return strconv.FormatBool(leftVal == rightVal), "bool"
		default:
			fmt.Printf("Error: Equality (==) not supported for type %s\n", leftTyp)
			return "false", "bool"
		}

	case *parser.NeExprContext:
		ne := ctx.(*parser.NeExprContext)
		leftVal, leftTyp := v.evalExpr(ne.Expr(0))
		rightVal, rightTyp := v.evalExpr(ne.Expr(1))
		if leftTyp != rightTyp {
			fmt.Printf("Error: Inequality (!=) requires matching types, got %s and %s\n", leftTyp, rightTyp)
			return "false", "bool"
		}
		switch leftTyp {
		case "i8", "i16", "i32", "i64", "int", "ui8", "ui16", "ui32", "ui64", "ui":
			left, _ := strconv.ParseInt(leftVal, 10, 64)
			right, _ := strconv.ParseInt(rightVal, 10, 64)
			return strconv.FormatBool(left != right), "bool"
		case "f32", "f64":
			left, _ := strconv.ParseFloat(leftVal, 64)
			right, _ := strconv.ParseFloat(rightVal, 64)
			return strconv.FormatBool(left != right), "bool"
		case "bool":
			left, _ := strconv.ParseBool(leftVal)
			right, _ := strconv.ParseBool(rightVal)
			return strconv.FormatBool(left != right), "bool"
		case "string":
			return strconv.FormatBool(leftVal != rightVal), "bool"
		default:
			fmt.Printf("Error: Inequality (!=) not supported for type %s\n", leftTyp)
			return "false", "bool"
		}

	case *parser.AndExprContext:
		and := ctx.(*parser.AndExprContext)
		leftVal, leftTyp := v.evalExpr(and.Expr(0))
		rightVal, rightTyp := v.evalExpr(and.Expr(1))
		if leftTyp != "bool" || rightTyp != "bool" {
			fmt.Printf("Error: Logical AND (&&) requires boolean operands, got %s and %s\n", leftTyp, rightTyp)
			return "false", "bool"
		}
		left, _ := strconv.ParseBool(leftVal)
		right, _ := strconv.ParseBool(rightVal)
		return strconv.FormatBool(left && right), "bool"

	case *parser.OrExprContext:
		or := ctx.(*parser.OrExprContext)
		leftVal, leftTyp := v.evalExpr(or.Expr(0))
		rightVal, rightTyp := v.evalExpr(or.Expr(1))
		if leftTyp != "bool" || rightTyp != "bool" {
			fmt.Printf("Error: Logical OR (||) requires boolean operands, got %s and %s\n", leftTyp, rightTyp)
			return "false", "bool"
		}
		left, _ := strconv.ParseBool(leftVal)
		right, _ := strconv.ParseBool(rightVal)
		return strconv.FormatBool(left || right), "bool"

	case *parser.VarExprContext:
		id := ctx.(*parser.VarExprContext).ID().GetText()
		if v, exists := v.Vars[id]; exists {
			return v.Value, v.Typ
		}
		fmt.Printf("Error: Variable %s not declared\n", id)
		return "0", "unknown"

	case *parser.LitExprContext:
		val := ctx.(*parser.LitExprContext).Literal().GetText()
		return val, inferType(val)
	}
	return "0", "unknown"
}

// inferType determines the type of a literal value
func inferType(val string) string {
	if val == "true" || val == "false" {
		return "bool"
	}
	if (val[0] == '"' && val[len(val)-1] == '"') || (val[0] == '\'' && val[len(val)-1] == '\'') {
		return "string"
	}
	if strings.HasPrefix(val, "0b") {
		return "byte"
	}
	if _, err := strconv.ParseFloat(val, 64); err == nil {
		if strings.Contains(val, ".") {
			return "f32"
		}
	}
	if _, err := strconv.Atoi(val); err == nil {
		return "i32"
	}
	return "unknown"
}

// isNumericType checks if a type is numeric
func isNumericType(typ string) bool {
	return isIntegerType(typ) || isFloatType(typ)
}

// isIntegerType checks if a type is an integer
func isIntegerType(typ string) bool {
	switch typ {
	case "i8", "i16", "i32", "i64", "int", "ui8", "ui16", "ui32", "ui64", "ui":
		return true
	}
	return false
}

// isFloatType checks if a type is a float
func isFloatType(typ string) bool {
	return typ == "f32" || typ == "f64"
}

// promoteNumericTypes determines the result type for binary operations
func promoteNumericTypes(leftTyp, rightTyp string) string {
	if leftTyp == "f64" || rightTyp == "f64" {
		return "f64"
	}
	if leftTyp == "f32" || rightTyp == "f32" {
		return "f32"
	}
	if leftTyp == "i64" || rightTyp == "i64" {
		return "i64"
	}
	if leftTyp == "ui64" || rightTyp == "ui64" {
		return "ui64"
	}
	if leftTyp == "i32" || rightTyp == "i32" || leftTyp == "int" || rightTyp == "int" {
		return "i32"
	}
	if leftTyp == "ui32" || rightTyp == "ui32" || leftTyp == "ui" || rightTyp == "ui" {
		return "ui32"
	}
	if leftTyp == "i16" || rightTyp == "i16" {
		return "i16"
	}
	if leftTyp == "ui16" || rightTyp == "ui16" {
		return "ui16"
	}
	if leftTyp == "i8" || rightTyp == "i8" {
		return "i8"
	}
	if leftTyp == "ui8" || rightTyp == "ui8" {
		return "ui8"
	}
	// default fallback
	return "i32"
}
