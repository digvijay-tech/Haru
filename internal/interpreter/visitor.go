package interpreter

import (
	"fmt"
	"strconv"
	"strings"

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
	case antlr.TerminalNode: // Skip terminals like ';'
		return nil
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
	val, valTyp := v.evalExpr(ctx.Expr())

	if valTyp == "unknown" { // Check for error case from evalExpr
		fmt.Printf("Error: Invalid expression for let %s\n", id)
		return nil
	}

	convertedVal, err := convertToType(val, valTyp, typ)
	if err != nil {
		fmt.Printf("Error: Type mismatch in let %s: expected %s, got %s (%s)\n", id, typ, valTyp, err)
		return nil
	}
	fmt.Printf("Declared let %s: %s = %s\n", id, typ, convertedVal)
	v.Vars[id] = Value{Value: convertedVal, Typ: typ, IsMut: false, IsConst: false}
	return nil
}

func (v *HaruVisitor) VisitMutDecl(ctx *parser.MutDeclContext) interface{} {
	id := ctx.ID().GetText()
	typ := ctx.Type_().GetText()
	var val string

	if ctx.Expr() != nil {
		var valTyp string
		val, valTyp = v.evalExpr(ctx.Expr())
		convertedVal, err := convertToType(val, valTyp, typ)
		if err != nil {
			fmt.Printf("Error: Type mismatch in mut %s: expected %s, got %s (%s)\n", id, typ, valTyp, err)
			return nil
		}
		val = convertedVal
		fmt.Printf("Declared mut %s: %s = %s\n", id, typ, val)
	} else {
		val = ""
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

	children := ctx.GetChildren()

	var ifBlock, elseBlock []antlr.ParseTree
	inElse := false

	for _, child := range children {
		if pt, ok := child.(antlr.ParseTree); ok {
			text := pt.GetText()
			if text == "else" {
				inElse = true
				continue
			}

			if !isTerminal(child) && child != ctx.Expr() && text != "{" && text != "}" {
				if inElse {
					elseBlock = append(elseBlock, pt)
				} else {
					ifBlock = append(ifBlock, pt)
				}
			}
		}
	}

	if cond {
		for _, stmt := range ifBlock {
			v.Visit(stmt)
		}
	} else if len(elseBlock) > 0 {
		for _, stmt := range elseBlock {
			v.Visit(stmt)
		}
	}

	return nil
}

// Add helper function
func isTerminal(node interface{}) bool {
	_, isTerm := node.(antlr.TerminalNode)
	return isTerm
}

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

func inferType(val string) string {
	if val == "true" || val == "false" {
		return "bool"
	}
	if (val[0] == '"' && val[len(val)-1] == '"') || (val[0] == '\'' && val[len(val)-1] == '\'') {
		return "string"
	}
	if strings.HasPrefix(val, "0b") {
		return "byte" // Byte literals like 0b1010
	}
	if _, err := strconv.ParseFloat(val, 64); err == nil {
		if strings.Contains(val, ".") {
			return "f32" // Default to f32 for floats; refine to f64 later if needed
		}
	}
	if _, err := strconv.Atoi(val); err == nil {
		return "i32" // Default to i32 for integers; we'll handle others via explicit types
	}
	return "unknown"
}

/******* HELPER FUNCTIONS *******/

func isNumericType(typ string) bool {
	return isIntegerType(typ) || isFloatType(typ)
}

func isIntegerType(typ string) bool {
	switch typ {
	case "i8", "i16", "i32", "i64", "int", "ui8", "ui16", "ui32", "ui64", "ui":
		return true
	}
	return false
}

func isFloatType(typ string) bool {
	return typ == "f32" || typ == "f64"
}

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
	return "i32" // default fallback
}

// helper function for type conversion
func convertToType(val, fromTyp, toTyp string) (string, error) {
	if fromTyp == toTyp {
		return val, nil
	}
	switch toTyp {
	case "i8":
		if isIntegerType(fromTyp) {
			i, _ := strconv.ParseInt(val, 10, 8)
			return strconv.FormatInt(i, 10), nil
		}
	case "i16":
		if isIntegerType(fromTyp) {
			i, _ := strconv.ParseInt(val, 10, 16)
			return strconv.FormatInt(i, 10), nil
		}
	case "i32", "int":
		if isIntegerType(fromTyp) {
			i, _ := strconv.ParseInt(val, 10, 32)
			return strconv.FormatInt(i, 10), nil
		}
	case "i64":
		if isIntegerType(fromTyp) {
			i, _ := strconv.ParseInt(val, 10, 64)
			return strconv.FormatInt(i, 10), nil
		}
	case "ui8":
		if isIntegerType(fromTyp) {
			u, _ := strconv.ParseUint(val, 10, 8)
			return strconv.FormatUint(u, 10), nil
		}
	case "ui16":
		if isIntegerType(fromTyp) {
			u, _ := strconv.ParseUint(val, 10, 16)
			return strconv.FormatUint(u, 10), nil
		}
	case "ui32", "ui":
		if isIntegerType(fromTyp) {
			u, _ := strconv.ParseUint(val, 10, 32)
			return strconv.FormatUint(u, 10), nil
		}
	case "ui64":
		if isIntegerType(fromTyp) {
			u, _ := strconv.ParseUint(val, 10, 64)
			return strconv.FormatUint(u, 10), nil
		}
	case "f32":
		if isNumericType(fromTyp) {
			f, _ := strconv.ParseFloat(val, 32)
			return strconv.FormatFloat(f, 'f', -1, 32), nil
		}
	case "f64":
		if isNumericType(fromTyp) {
			f, _ := strconv.ParseFloat(val, 64)
			return strconv.FormatFloat(f, 'f', -1, 64), nil
		}
	case "bool":
		if fromTyp == "bool" {
			return val, nil
		}
	case "string":
		if fromTyp == "string" {
			return val, nil
		}
	case "byte":
		if fromTyp == "byte" || strings.HasPrefix(val, "0b") {
			return val, nil
		}
	}
	return "", fmt.Errorf("cannot convert %s to %s", fromTyp, toTyp)
}
