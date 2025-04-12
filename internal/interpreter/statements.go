package interpreter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/digvijay-tech/Haru/internal/parser"
)

// VisitProgram visits the program node, executing all statements
func (v *HaruVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	for _, stmt := range ctx.AllStatement() {
		v.Visit(stmt)
	}
	return nil
}

// VisitLetDecl visits a let declaration node
func (v *HaruVisitor) VisitLetDecl(ctx *parser.LetDeclContext) interface{} {
	id := ctx.ID().GetText()
	typ := ctx.Type_().GetText()
	val, valTyp := v.evalExpr(ctx.Expr())
	if valTyp == "unknown" {
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

// VisitMutDecl visits a mutable declaration node
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

// VisitAssignStmt visits an assignment node
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

// VisitPrintStatement visits a print statement node
func (v *HaruVisitor) VisitPrintStatement(ctx *parser.PrintStatementContext) interface{} {
	val, _ := v.evalExpr(ctx.Expr())
	fmt.Println("Output:", val)
	return nil
}

// VisitIfStatement visits an if statement node
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

// isTerminal checks if a node is a terminal
func isTerminal(node interface{}) bool {
	_, isTerm := node.(antlr.TerminalNode)
	return isTerm
}

// convertToType converts a value from one type to another
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
