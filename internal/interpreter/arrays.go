package interpreter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/digvijay-tech/Haru/internal/parser"
)

// VisitArrayLiteralExprList evaluates the array literal expression and returns parsed items array
func (v *HaruVisitor) VisitArrayLiteralExprList(ctx *parser.ArrayLiteralExprListContext) any {
	var items []Value

	for _, expr := range ctx.AllExpr() {
		// evaluates expressions but its currently of type any
		result := v.Visit(expr)

		// asserting type and ensuring type is a valid Value struct
		parsedItem, ok := result.(Value)

		if !ok {
			runtimeErr("invalid value in array")
		}

		items = append(items, parsedItem)
	}

	return items
}

// VisitIndexExpr evaluates expression to access item based on array index
func (v *HaruVisitor) VisitIndexExpr(ctx *parser.IndexExprContext) any {
	varName := ctx.ID().GetText()

	// evaluating expression to get index value
	expr, ok := v.Visit(ctx.Expr()).(Value)

	if !ok {
		runtimeErr(fmt.Sprintf("invalid index for %s", varName))
	}

	// only non-float numeric expressions are allowed
	if !isNumericType(expr.Typ) || expr.Typ == "f32" || expr.Typ == "f64" {
		runtimeErr(fmt.Sprintf("index for array '%s' must be an integer, got '%s'", varName, expr.Typ))
	}

	// converting to integer
	index, err := strconv.Atoi(expr.Value)
	if err != nil {
		runtimeErr(fmt.Sprintf("invalid index value for array '%s': %v", varName, err))
	}

	// getting the stringified value out from symbol table
	arrayVal, ok := v.symbolTable[varName]
	if !ok {
		runtimeErr(fmt.Sprintf("undefined array '%s'", varName))
	}

	if !strings.HasPrefix(arrayVal.Typ, "[]") {
		runtimeErr(fmt.Sprintf("'%s' is not an array", varName))
	}

	// removing [] from both sides
	raw := strings.Trim(arrayVal.Value, "[]")

	// splitting the string into a slice of strings
	parts := strings.Split(raw, ",")

	// making sure index in not out of bound
	if index < 0 || index >= len(parts) {
		runtimeErr(fmt.Sprintf("index %d out of bounds for array '%s'", index, varName))
	}

	// getting value without ""
	value := strings.Trim(parts[index], "\"")

	return Value{
		Value: value,
		Typ:   arrayVal.Typ[2:], // removing [] prefix
	}
}

// VisitArrayDeclStatement routes visitor to one of the array category
func (v *HaruVisitor) VisitArrayDeclStatement(ctx *parser.ArrayDeclStatementContext) any {
	switch child := ctx.ArrayDecl().GetChild(0).(type) {
	case *parser.ConstExplicitArrayDeclContext:
		return v.VisitConstExplicitArrayDecl(child)
	case *parser.ConstImplicitArrayDeclContext:
		return v.VisitConstImplicitArrayDecl(child)
	default:
		runtimeErr("unknown array declaration type")
	}

	// will never reach here but but need this for syntax validation
	return nil
}

// VisitConstExplicitArrayDecl evaluates array declared with const, type and value/array literal
func (v *HaruVisitor) VisitConstExplicitArrayDecl(ctx *parser.ConstExplicitArrayDeclContext) any {
	arrName := ctx.ID().GetText()
	arrType := ctx.ArrayType().Type_().GetText()

	// getting the parsed array literal from array literal context
	items, ok := v.Visit(ctx.ArrayLiteral()).([]Value)
	if !ok {
		runtimeErr(fmt.Sprintf("invalid/empty array literal in %s", arrName))
	}

	// array declared with const cannot be empty
	if len(items) == 0 {
		runtimeErr(fmt.Sprintf("const array '%s' cannot be empty", arrName))
	}

	// converting items to constant's type
	var standardizedItems []Value

	for i, value := range items {
		newVal, err := convertType(value.Value, value.Typ, arrType)

		if err != nil {
			runtimeErr(err.Error())
		}

		// asserting type from any to Value
		updatedVal, ok := newVal.(Value)

		if !ok {
			runtimeErr(fmt.Sprintf("type assertion failed for %s at index: %d", arrName, i))
		}

		standardizedItems = append(standardizedItems, updatedVal)
	}

	// ensuring all items have the exact same type as specified in declaration
	// and parsing them into str array
	var stringifiedItems []string

	for i, value := range standardizedItems {
		if value.Typ != arrType {
			runtimeErr(fmt.Sprintf("type mismatch in array '%s': expected all items to be '%s', found '%s' at index: %d", arrName, arrType, value.Typ, i))
		}

		// preserving quotes for string
		if arrType == "string" {
			stringifiedItems = append(stringifiedItems, fmt.Sprintf(`"%v"`, value.Value))
		} else {
			stringifiedItems = append(stringifiedItems, value.Value)
		}
	}

	// serializing final array value/literal by adding [] on both ends
	serialzed := "[" + strings.Join(stringifiedItems, ",") + "]"

	// storing in symbol table for global access
	v.symbolTable[arrName] = Value{
		Value: serialzed,
		Typ:   "[]" + arrType,
	}

	return nil
}

// VisitConstImplicitArrayDecl evaluates array declared with const and but type is inferred by associated array literal
func (v *HaruVisitor) VisitConstImplicitArrayDecl(ctx *parser.ConstImplicitArrayDeclContext) any {
	arrName := ctx.ID().GetText()

	// getting the parsed array literal from array literal context
	items, ok := v.Visit(ctx.ArrayLiteral()).([]Value)
	if !ok {
		runtimeErr(fmt.Sprintf("invalid/empty array literal in %s", arrName))
	}

	// array declared with const cannot be empty
	if len(items) == 0 {
		runtimeErr(fmt.Sprintf("const array '%s' cannot be empty", arrName))
	}

	// infering the type from first item
	inferredType := items[0].Typ

	// making sure all elements have the exact same type as the first element
	// and parsing them into str array
	var stringifiedItems []string

	for i, val := range items {
		if val.Typ != inferredType {
			runtimeErr(fmt.Sprintf("type mismatch in array '%s': expected all items to be '%s', found '%s' at index: %d", arrName, inferredType, val.Typ, i))
		}

		// preserving quotes for string
		if inferredType == "string" {
			stringifiedItems = append(stringifiedItems, fmt.Sprintf(`"%v"`, val.Value))
		} else {
			stringifiedItems = append(stringifiedItems, val.Value)
		}
	}

	// serializing final array value/literal by adding [] on both ends
	serialzed := "[" + strings.Join(stringifiedItems, ",") + "]"

	// storing in symbol table for global access
	v.symbolTable[arrName] = Value{
		Value: serialzed,
		Typ:   "[]" + inferredType,
	}

	return nil
}
