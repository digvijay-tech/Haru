package interpreter

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/digvijay-tech/Haru/internal/parser"
)

type HaruVisitor struct {
	parser.BaseharuVisitor
	Vars map[string]Value
}

type Value struct {
	Value   string
	Typ     string
	IsMut   bool
	IsConst bool
}

func NewHaruVisitor() *HaruVisitor {
	return &HaruVisitor{
		Vars: make(map[string]Value),
	}
}

func (v *HaruVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch ctx := tree.(type) {
	case *parser.ProgramContext:
		return v.VisitProgram(ctx)
	case *parser.LetDeclContext:
		return v.VisitLetDecl(ctx)
	case *parser.MutDeclContext:
		return v.VisitMutDecl(ctx)
	case *parser.AssignStmtContext:
		return v.VisitAssignStmt(ctx)
	case *parser.PrintStatementContext:
		return v.VisitPrintStatement(ctx)
	case *parser.IfStatementContext:
		return v.VisitIfStatement(ctx)
	}

	return v.VisitChildren(tree.(antlr.RuleNode))
}

func (v *HaruVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, child := range node.GetChildren() {
		if tree, ok := child.(antlr.ParseTree); ok {
			tree.Accept(v)
		}
	}

	return nil
}

func (v *HaruVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return nil
}

func (v *HaruVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	fmt.Println("Error in parsing:", node.GetText())
	return nil
}

func (v *HaruVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	for _, stmt := range ctx.AllStatement() {
		v.Visit(stmt)
	}

	return nil
}

func (v *HaruVisitor) VisitLetDecl(ctx *parser.LetDeclContext) interface{} {
	id := ctx.ID().GetText()
	typ := ctx.Type_().GetText()
	val := ctx.Literal().GetText()

	v.Vars[id] = Value{Value: val, Typ: typ, IsMut: false, IsConst: false}
	fmt.Printf("Declared let %s: %s = %s\n", id, typ, val)

	return nil
}

func (v *HaruVisitor) VisitMutDecl(ctx *parser.MutDeclContext) interface{} {
	id := ctx.ID().GetText()
	typ := ctx.Type_().GetText()
	val := ""

	if ctx.Literal() != nil {
		val = ctx.Literal().GetText()
		fmt.Printf("Declared mut %s: %s = %s\n", id, typ, val)
	} else {
		fmt.Printf("Declared mut %s: %s (uninitialized)\n", id, typ)
	}

	v.Vars[id] = Value{Value: val, Typ: typ, IsMut: true, IsConst: false}
	return nil
}

func (v *HaruVisitor) VisitAssignStmt(ctx *parser.AssignStmtContext) interface{} {
	id := ctx.ID().GetText()
	val, typ := v.evalExpr(ctx.Expr())

	if existing, exists := v.Vars[id]; exists {
		if !existing.IsMut {
			fmt.Printf("Error: Cannot assign to immutable variable %s\n", id)
			return nil
		}
		if existing.Typ != typ {
			fmt.Printf("Error: Type mismatch assigning %s to %s (expected %s, got %s)\n", val, id, existing.Typ, typ)
			return nil
		}

		v.Vars[id] = Value{Value: val, Typ: typ, IsMut: true, IsConst: false}
		fmt.Printf("Assigned %s = %s\n", id, val)
	} else {
		fmt.Printf("Error: Variable %s not declared\n", id)
	}

	return nil
}

func (v *HaruVisitor) VisitPrintStatement(ctx *parser.PrintStatementContext) interface{} {
	val, _ := v.evalExpr(ctx.Expr())
	fmt.Println("Output:", val)
	return nil
}

func (v *HaruVisitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	condVal, condTyp := v.evalExpr(ctx.Expr())
	if condTyp != "bool" {
		fmt.Println("Error: If condition must be a boolean")
		return nil
	}

	cond, err := strconv.ParseBool(condVal)
	if err != nil {
		fmt.Println("Error: Invalid boolean condition:", err)
		return nil
	}

	if cond {
		for _, stmt := range ctx.AllStatement() {
			v.Visit(stmt)
		}
	}

	return nil
}

func (v *HaruVisitor) evalExpr(ctx parser.IExprContext) (string, string) {
	switch ctx.(type) {
	case *parser.ParenExprContext:
		return v.evalExpr(ctx.(*parser.ParenExprContext).Expr())
	case *parser.ExpExprContext:
		exp := ctx.(*parser.ExpExprContext)
		baseVal, baseTyp := v.evalExpr(exp.Expr(0))
		expVal, expTyp := v.evalExpr(exp.Expr(1))

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
		leftVal, leftTyp := v.evalExpr(mul.Expr(0))
		rightVal, rightTyp := v.evalExpr(mul.Expr(1))

		if leftTyp != "int" || rightTyp != "int" {
			fmt.Println("Error: Multiplication only supports int for now")
			return "0", "int"
		}

		left, _ := strconv.Atoi(leftVal)
		right, _ := strconv.Atoi(rightVal)

		return strconv.Itoa(left * right), "int"
	case *parser.DivExprContext:
		div := ctx.(*parser.DivExprContext)
		leftVal, leftTyp := v.evalExpr(div.Expr(0))
		rightVal, rightTyp := v.evalExpr(div.Expr(1))

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

		return strconv.Itoa(left / right), "int"
	case *parser.ModExprContext:
		mod := ctx.(*parser.ModExprContext)
		leftVal, leftTyp := v.evalExpr(mod.Expr(0))
		rightVal, rightTyp := v.evalExpr(mod.Expr(1))

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

		return strconv.Itoa(left % right), "int"
	case *parser.AddExprContext:
		add := ctx.(*parser.AddExprContext)
		leftVal, leftTyp := v.evalExpr(add.Expr(0))
		rightVal, rightTyp := v.evalExpr(add.Expr(1))

		if leftTyp != "int" || rightTyp != "int" {
			fmt.Println("Error: Addition only supports int for now")
			return "0", "int"
		}

		left, _ := strconv.Atoi(leftVal)
		right, _ := strconv.Atoi(rightVal)

		return strconv.Itoa(left + right), "int"
	case *parser.SubExprContext:
		sub := ctx.(*parser.SubExprContext)
		leftVal, leftTyp := v.evalExpr(sub.Expr(0))
		rightVal, rightTyp := v.evalExpr(sub.Expr(1))

		if leftTyp != "int" || rightTyp != "int" {
			fmt.Println("Error: Subtraction only supports int for now")
			return "0", "int"
		}

		left, _ := strconv.Atoi(leftVal)
		right, _ := strconv.Atoi(rightVal)

		return strconv.Itoa(left - right), "int"
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
