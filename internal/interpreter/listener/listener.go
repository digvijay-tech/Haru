package listener

import (
	"fmt"
	"strconv"

	"github.com/digvijay-tech/Haru/internal/parser"
)

type HaruListener struct {
	*parser.BaseharuListener
	Vars map[string]Value
}

type Value struct {
	Value   string
	Typ     string
	IsMut   bool
	IsConst bool
}

func NewHaruListener() *HaruListener {
	return &HaruListener{
		Vars:             make(map[string]Value),
		BaseharuListener: &parser.BaseharuListener{},
	}
}

func (l *HaruListener) ExitLetDecl(ctx *parser.LetDeclContext) {
	id := ctx.ID().GetText()
	typ := ctx.Type_().GetText()
	literal := ctx.Literal()

	if literal == nil {
		fmt.Printf("Error: 'let' variable %s must have a value\n", id)
		return
	}

	val := literal.GetText()
	l.Vars[id] = Value{Value: val, Typ: typ, IsMut: false, IsConst: false}
	fmt.Printf("Declared let %s: %s = %s\n", id, typ, val)
}

func (l *HaruListener) ExitLetInferDecl(ctx *parser.LetInferDeclContext) {
	id := ctx.ID().GetText()
	val := ctx.Literal().GetText()
	typ := inferType(val)

	l.Vars[id] = Value{Value: val, Typ: typ, IsMut: false, IsConst: false}
	fmt.Printf("Declared let %s: %s = %s (inferred)\n", id, typ, val)
}

func (l *HaruListener) ExitMutDecl(ctx *parser.MutDeclContext) {
	id := ctx.ID().GetText()
	typ := ctx.Type_().GetText()
	literal := ctx.Literal()

	val := ""

	if literal != nil {
		val = literal.GetText()
		fmt.Printf("Declared mut %s: %s = %s\n", id, typ, val)
	} else {
		fmt.Printf("Declared mut %s: %s (uninitialized)\n", id, typ)
	}

	l.Vars[id] = Value{Value: val, Typ: typ, IsMut: true, IsConst: false}
}

func (l *HaruListener) ExitMutInferDecl(ctx *parser.MutInferDeclContext) {
	id := ctx.ID().GetText()
	val := ctx.Literal().GetText()
	typ := inferType(val)

	l.Vars[id] = Value{Value: val, Typ: typ, IsMut: true, IsConst: false}
	fmt.Printf("Declared mut %s: %s = %s (inferred)\n", id, typ, val)
}

func (l *HaruListener) ExitConstDecl(ctx *parser.ConstDeclContext) {
	id := ctx.ID().GetText()
	typ := ctx.Type_().GetText()
	literal := ctx.Literal()

	if literal == nil {
		fmt.Printf("Error: 'const' variable %s must have a value\n", id)
		return
	}

	val := literal.GetText()
	l.Vars[id] = Value{Value: val, Typ: typ, IsMut: false, IsConst: true}

	fmt.Printf("Declared const %s: %s = %s\n", id, typ, val)
}

func (l *HaruListener) ExitConstInferDecl(ctx *parser.ConstInferDeclContext) {
	id := ctx.ID().GetText()
	val := ctx.Literal().GetText()
	typ := inferType(val)

	l.Vars[id] = Value{Value: val, Typ: typ, IsMut: false, IsConst: true}

	fmt.Printf("Declared const %s: %s = %s (inferred)\n", id, typ, val)
}

func (l *HaruListener) ExitAssignStmt(ctx *parser.AssignStmtContext) {
	id := ctx.ID().GetText()
	val, typ := l.evalExpr(ctx.Expr())

	if v, exists := l.Vars[id]; exists {
		if !v.IsMut {
			fmt.Printf("Error: Cannot assign to immutable variable %s\n", id)
			return
		}

		if v.Typ != typ {
			fmt.Printf("Error: Type mismatch assigning %s to %s (expected %s, got %s)\n", val, id, v.Typ, typ)
			return
		}

		l.Vars[id] = Value{Value: val, Typ: typ, IsMut: true, IsConst: false}
		fmt.Printf("Assigned %s = %s\n", id, val)
	} else {
		fmt.Printf("Error: Variable %s not declared\n", id)
	}
}

func (l *HaruListener) ExitPrintStatement(ctx *parser.PrintStatementContext) {
	val, _ := l.evalExpr(ctx.Expr())
	fmt.Println("Output:", val)
}

func (l *HaruListener) evalExpr(ctx parser.IExprContext) (string, string) {
	switch ctx.(type) {
	case *parser.ParenExprContext:
		paren := ctx.(*parser.ParenExprContext)
		return l.evalExpr(paren.Expr())
	case *parser.ExpExprContext:
		exp := ctx.(*parser.ExpExprContext)
		baseVal, baseTyp := l.evalExpr(exp.Expr(0))
		expVal, expTyp := l.evalExpr(exp.Expr(1))

		if baseTyp != "int" || expTyp != "int" {
			fmt.Println("Error: Exponentiation only supports int for now")
			return "0", "int"
		}

		base, _ := strconv.Atoi(baseVal)
		exponent, _ := strconv.Atoi(expVal)

		if exponent < 0 {
			fmt.Println("Error: Negative exponents not supported")
			return "0", "int"
		}

		result := 1

		for i := 0; i < exponent; i++ {
			result *= base
		}

		return strconv.Itoa(result), "int"
	case *parser.MulExprContext:
		mul := ctx.(*parser.MulExprContext)
		leftVal, leftTyp := l.evalExpr(mul.Expr(0))
		rightVal, rightTyp := l.evalExpr(mul.Expr(1))

		if leftTyp != "int" || rightTyp != "int" {
			fmt.Println("Error: Multiplication only supports int for now")
			return "0", "int"
		}

		left, _ := strconv.Atoi(leftVal)
		right, _ := strconv.Atoi(rightVal)
		result := left * right

		return strconv.Itoa(result), "int"
	case *parser.DivExprContext:
		div := ctx.(*parser.DivExprContext)
		leftVal, leftTyp := l.evalExpr(div.Expr(0))
		rightVal, rightTyp := l.evalExpr(div.Expr(1))

		if leftTyp != "int" || rightTyp != "int" {
			fmt.Println("Error: Division only supports int for now")
			return "0", "int"
		}

		left, _ := strconv.Atoi(leftVal)
		right, _ := strconv.Atoi(rightVal)

		if right == 0 {
			fmt.Println("Error: Division by zero")
			return "0", "int"
		}

		result := left / right
		return strconv.Itoa(result), "int"
	case *parser.ModExprContext:
		mod := ctx.(*parser.ModExprContext)
		leftVal, leftTyp := l.evalExpr(mod.Expr(0))
		rightVal, rightTyp := l.evalExpr(mod.Expr(1))

		if leftTyp != "int" || rightTyp != "int" {
			fmt.Println("Error: Modulus only supports int for now")
			return "0", "int"
		}

		left, _ := strconv.Atoi(leftVal)
		right, _ := strconv.Atoi(rightVal)

		if right == 0 {
			fmt.Println("Error: Modulus by zero")
			return "0", "int"
		}

		result := left % right
		return strconv.Itoa(result), "int"
	case *parser.AddExprContext:
		add := ctx.(*parser.AddExprContext)
		leftVal, leftTyp := l.evalExpr(add.Expr(0))
		rightVal, rightTyp := l.evalExpr(add.Expr(1))

		if leftTyp != "int" || rightTyp != "int" {
			fmt.Println("Error: Addition only supports int for now")
			return "0", "int"
		}

		left, _ := strconv.Atoi(leftVal)
		right, _ := strconv.Atoi(rightVal)
		result := left + right

		return strconv.Itoa(result), "int"
	case *parser.SubExprContext:
		sub := ctx.(*parser.SubExprContext)
		leftVal, leftTyp := l.evalExpr(sub.Expr(0))
		rightVal, rightTyp := l.evalExpr(sub.Expr(1))

		if leftTyp != "int" || rightTyp != "int" {
			fmt.Println("Error: Subtraction only supports int for now")
			return "0", "int"
		}

		left, _ := strconv.Atoi(leftVal)
		right, _ := strconv.Atoi(rightVal)
		result := left - right

		return strconv.Itoa(result), "int"
	case *parser.VarExprContext:
		varExpr := ctx.(*parser.VarExprContext)
		id := varExpr.ID().GetText()

		if v, exists := l.Vars[id]; exists {
			return v.Value, v.Typ
		}

		fmt.Printf("Error: Variable %s not declared\n", id)

		return "0", "unknown"
	case *parser.LitExprContext:
		lit := ctx.(*parser.LitExprContext)
		val := lit.Literal().GetText()

		return val, inferType(val)
	}

	return "0", "unknown"
}

func inferType(val string) string {
	if _, err := strconv.Atoi(val); err == nil {
		return "int"
	}
	if _, err := strconv.ParseFloat(val, 64); err == nil {
		return "f32"
	}
	if val == "true" || val == "false" {
		return "bool"
	}
	if (val[0] == '"' && val[len(val)-1] == '"') || (val[0] == '\'' && val[len(val)-1] == '\'') {
		return "string"
	}
	return "unknown"
}
