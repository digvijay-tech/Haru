package interpreter

import (
	"fmt"
	"regexp"
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
		runtimeErr(fmt.Sprintf("index for array '%s' must be an unsigned integer, got '%s'", varName, expr.Typ))
	}

	// converting to unsigned int
	index, err := strconv.ParseUint(expr.Value, 10, 64)
	if err != nil {
		runtimeErr(fmt.Sprintf("invalid index value for array '%s': %v", varName, err))
	}

	// getting the stringified value out from symbol table
	arrayVal, ok := v.symbolTable[varName]
	if !ok {
		runtimeErr(fmt.Sprintf("undefined array '%s'", varName))
	}

	// checking if array is empty/declared empty with mut
	if arrayVal.isMutable && arrayVal.Value == "[]" {
		runtimeErr(fmt.Sprintf("array %s is empty", varName))
	}

	// matching []type and [number]type pattern with regex
	var arrayTypePattern = regexp.MustCompile(`^\[\d*\]\w+$`)

	if !arrayTypePattern.Match([]byte(arrayVal.Typ)) {
		runtimeErr(fmt.Sprintf("'%s' is not an array", varName))
	}

	// removing [] from both sides
	raw := strings.Trim(arrayVal.Value, "[]")

	// splitting the string into a slice of strings
	parts := strings.Split(raw, ",")

	// making sure index in not out of bound
	if int(index) >= len((parts)) {
		runtimeErr(fmt.Sprintf("index %d out of bounds for array '%s'", index, varName))
	}

	// getting value without ""
	value := strings.Trim(parts[index], "\"")

	// cleanup removing [] and extracting type
	closing := strings.Index(arrayVal.Typ, "]")
	if closing == -1 || closing == len(arrayVal.Typ)-1 {
		runtimeErr(fmt.Sprintf("invalid array type format: %s", arrayVal.Typ))
	}

	cleanType := arrayVal.Typ[closing+1:]

	return Value{
		Value: value,
		Typ:   cleanType,
	}
}

// VisitArrayDeclStatement routes visitor to one of the array category
func (v *HaruVisitor) VisitArrayDeclStatement(ctx *parser.ArrayDeclStatementContext) any {
	switch child := ctx.ArrayDecl().GetChild(0).(type) {
	case *parser.ConstExplicitArrayDeclContext:
		return v.VisitConstExplicitArrayDecl(child)
	case *parser.ConstImplicitArrayDeclContext:
		return v.VisitConstImplicitArrayDecl(child)
	case *parser.LetExplicitArrayDeclContext:
		return v.VisitLetExplicitArrayDecl(child)
	case *parser.LetImplicitArrayDeclContext:
		return v.VisitLetImplicitArrayDecl(child)
	case *parser.MutArrayExplicitWithInitContext:
		return v.VisitDynamicExplicitMutArrayDecl(child)
	case *parser.MutFixedArrayNoInitContext:
		return v.VisitMutFixedArrayNoInitDecl(child)
	case *parser.MutFixedArrayWithInitContext:
		return v.VisitMutFixedArrayInitDecl(child)
	case *parser.MutArrayExplicitNoInitContext:
		return v.VisitDynamicExplicitMutArrayUnInitDecl(child)
	case *parser.MutArrayImplicitContext:
		return v.VisitImplicitMutArrayDecl(child)

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

// VisitLetExplicitArrayDecl evaluates array declared with let, type and value/array literal
func (v *HaruVisitor) VisitLetExplicitArrayDecl(ctx *parser.LetExplicitArrayDeclContext) any {
	arrName := ctx.ID().GetText()
	arrType := ctx.ArrayType().Type_().GetText()

	// getting the parsed array literal from array literal context
	items, ok := v.Visit(ctx.ArrayLiteral()).([]Value)
	if !ok {
		runtimeErr(fmt.Sprintf("invalid/empty array literal in %s", arrName))
	}

	// array declared with let cannot be empty
	if len(items) == 0 {
		runtimeErr(fmt.Sprintf("const array '%s' cannot be empty", arrName))
	}

	// converting items to let's type
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

// VisitLetImplicitArrayDecl evaluates array declared with let and but type is inferred by associated array literal
func (v *HaruVisitor) VisitLetImplicitArrayDecl(ctx *parser.LetImplicitArrayDeclContext) any {
	arrName := ctx.ID().GetText()

	// getting the parsed array literal from array literal context
	items, ok := v.Visit(ctx.ArrayLiteral()).([]Value)
	if !ok {
		runtimeErr(fmt.Sprintf("invalid/empty array literal in %s", arrName))
	}

	// array declared with let cannot be empty
	if len(items) == 0 {
		runtimeErr(fmt.Sprintf("let array '%s' cannot be empty", arrName))
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

// VisitDynamicExplicitMutArrayDecl evaluates dynamic explicit array
func (v *HaruVisitor) VisitDynamicExplicitMutArrayDecl(ctx *parser.MutArrayExplicitWithInitContext) any {
	arrName := ctx.ID().GetText()
	arrType := ctx.ArrayType().Type_().GetText()

	// allowing [] empty array initialization
	if v.Visit(ctx.ArrayLiteral()) == nil {
		v.symbolTable[arrName] = Value{
			Value:     "[]",
			Typ:       "[]" + arrType,
			isMutable: true,
		}

		return nil
	}

	items, ok := v.Visit(ctx.ArrayLiteral()).([]Value)

	if !ok {
		runtimeErr(fmt.Sprintf("invalid array literal in mut '%s'", arrName))
	}

	// standardize + validate items
	var standardizedItems []Value

	for i, value := range items {
		newVal, err := convertType(value.Value, value.Typ, arrType)

		if err != nil {
			runtimeErr(fmt.Sprintf("type error at index %d in '%s': %s", i, arrName, err.Error()))
		}

		casted, ok := newVal.(Value)
		if !ok {
			runtimeErr(fmt.Sprintf("failed to cast value at index %d in '%s'", i, arrName))
		}

		standardizedItems = append(standardizedItems, casted)
	}

	// converting to serialized form
	var stringified []string

	for _, val := range standardizedItems {
		if arrType == "string" {
			stringified = append(stringified, fmt.Sprintf(`"%v"`, val.Value))
		} else {
			stringified = append(stringified, val.Value)
		}
	}

	serialized := "[" + strings.Join(stringified, ",") + "]"

	v.symbolTable[arrName] = Value{
		Value:     serialized,
		Typ:       "[]" + arrType,
		isMutable: true,
	}

	return nil
}

// VisitDynamicExplicitMutArrayUnInitDecl evevalutes dynamic arrays declared with type but no initialization
func (v *HaruVisitor) VisitDynamicExplicitMutArrayUnInitDecl(ctx *parser.MutArrayExplicitNoInitContext) any {
	arrName := ctx.ID().GetText()
	arrType := ctx.ArrayType().Type_().GetText()

	v.symbolTable[arrName] = Value{
		Value:     "[]",
		Typ:       "[]" + arrType,
		isMutable: true,
	}

	return nil
}

// VisitMutFixedArrayNoInitDecl evaluates fixed length uninitialized mutable arrays
// when uninitialised this function will use zero value as a placeholder until new value is assigned
func (v *HaruVisitor) VisitMutFixedArrayNoInitDecl(ctx *parser.MutFixedArrayNoInitContext) any {
	arrName := ctx.ID().GetText()
	arrType := ctx.FixedArrayType().Type_().GetText()

	// extracting length from type expression
	expr := ctx.FixedArrayType().NUMBER().GetText()

	// parsing expression into uint
	length, err := strconv.ParseUint(expr, 10, 64)

	if err != nil || length == 0 {
		runtimeErr(fmt.Sprintf("invalid array size for '%s'", arrName))
	}

	// generating zero values
	zeroVal, err := zeroValueFor(arrType)
	if err != nil {
		runtimeErr(err.Error())
	}

	// populating array with zero value
	var serializedItems []string

	for range int(length) {
		serializedItems = append(serializedItems, zeroVal.Value)
	}

	serialized := "[" + strings.Join(serializedItems, ",") + "]"

	v.symbolTable[arrName] = Value{
		Value:     serialized,
		Typ:       fmt.Sprintf("[%d]%s", length, arrType),
		isMutable: true,
	}

	return nil
}

// VisitMutFixedArrayInitDecl evaluates fixed length initialized mutable array
func (v *HaruVisitor) VisitMutFixedArrayInitDecl(ctx *parser.MutFixedArrayWithInitContext) any {
	arrName := ctx.ID().GetText()
	arrType := ctx.FixedArrayType().Type_().GetText()
	number := ctx.FixedArrayType().NUMBER().GetText()

	// evaluating expression to get assigned array literal
	items, ok := v.Visit(ctx.ArrayLiteral()).([]Value)

	if !ok {
		runtimeErr("invalid array literal")
	}

	// parsing to uint from fixed length expression
	length, err := strconv.ParseUint(number, 10, 64)

	if err != nil || length == 0 {
		runtimeErr(fmt.Sprintf("invalid array size for '%s'", arrName))
	}

	// items cannot be more than length
	if len(items) > int(length) {
		runtimeErr(fmt.Sprintf("cannot fit more items than %d in %s", length, arrName))
	}

	// converting items to declared type
	var serializedItems []string

	// populating array when its either half, full or empty
	for _, item := range items {
		// converting the assigned value
		converted, err := convertType(item.Value, item.Typ, arrType)
		if err != nil {
			runtimeErr(err.Error())
		}

		val := converted.(Value)

		serializedItems = append(serializedItems, val.Value)
	}

	// populating remaining uninitialized array with zero value
	zeroVal, err := zeroValueFor(arrType)

	if err != nil {
		runtimeErr(err.Error())
	}

	if len(serializedItems) < int(length) {
		diff := int(length) - len(serializedItems)

		for range diff {
			serializedItems = append(serializedItems, zeroVal.Value)
		}
	}

	serialized := "[" + strings.Join(serializedItems, ",") + "]"

	v.symbolTable[arrName] = Value{
		Value:     serialized,
		Typ:       fmt.Sprintf("[%d]%s", length, arrType),
		isMutable: true,
	}

	return nil
}

// VisitImplicitMutArrayDecl evaluates mutable array with no type, it will be inferred from first array literal
func (v *HaruVisitor) VisitImplicitMutArrayDecl(ctx *parser.MutArrayImplicitContext) any {
	arrName := ctx.ID().GetText()
	items, ok := v.Visit(ctx.ArrayLiteral()).([]Value)

	if !ok {
		runtimeErr("invalid array literal")
	}

	// assigned array literal must have atleast 1 element inside for type inference
	if len(items) < 1 {
		runtimeErr(fmt.Sprintf("array literal cannot be empty for %s", arrName))
	}

	// infer type from first array element
	inferredType := items[0].Typ

	// converting items to inferred type
	var serializedItems []string

	for _, value := range items {
		converted, err := convertType(value.Value, value.Typ, inferredType)

		if err != nil {
			runtimeErr(err.Error())
		}

		newValue := converted.(Value)

		serializedItems = append(serializedItems, newValue.Value)
	}

	serialized := "[" + strings.Join(serializedItems, ",") + "]"

	v.symbolTable[arrName] = Value{
		Value:     serialized,
		Typ:       "[]" + inferredType,
		isMutable: true,
	}

	return nil
}

// VisitMutArrayReassignment reassign mut with whole new array of same type
func (v *HaruVisitor) VisitMutArrayReassignment(ctx *parser.ArrayReassignStatementContext) any {
	arrName := ctx.ArrayReassign().ID().GetText()
	items, ok := v.Visit(ctx.ArrayReassign().ArrayLiteral()).([]Value)

	if !ok {
		runtimeErr("invalid array literal")
	}

	// prevent if variable is not defined and get variable if it exists
	variable, ok := v.symbolTable[arrName]

	if !ok {
		runtimeErr(fmt.Sprintf("undefined variable '%s'", arrName))
	}

	if !variable.isMutable {
		runtimeErr(fmt.Sprintf("cannot reassign to immutable '%s'", arrName))
	}

	// preventing fixed arrays to have more values than speficied in type declaration
	// fixed array reassignment must have exact number of items
	fixedLen, err := extractNumFromBrackets(variable.Typ)
	if err != nil {
		runtimeErr(err.Error())
	}

	if len(items) != fixedLen && fixedLen != 0 {
		runtimeErr(fmt.Sprintf("cannot reassign to %s it expects %d items got %d", arrName, fixedLen, len(items)))
	}

	// variable is mutable and converting array literals to its type
	var serializedItems []string

	// removing [] or [num] from type
	cleanType := stripArrayPrefix(variable.Typ)

	for _, item := range items {
		converted, err := convertType(item.Value, item.Typ, cleanType)

		if err != nil {
			runtimeErr(err.Error())
		}

		newValue := converted.(Value)

		serializedItems = append(serializedItems, newValue.Value)
	}

	serialized := "[" + strings.Join(serializedItems, ",") + "]"

	// updating the variable in symbol table
	v.symbolTable[arrName] = Value{
		Value:     serialized,
		Typ:       variable.Typ,
		isMutable: true,
	}

	return nil
}

// VisitArrayIndexAssignStatement reassigns to array element by index
func (v *HaruVisitor) VisitArrayIndexAssignStatement(ctx *parser.ArrayIndexAssignStatementContext) any {
	assignCtx := ctx.ArrayItemAssign().(*parser.ArrayIndexAssignContext)
	arrName := assignCtx.ID().GetText()

	// evaluate the index expression
	rawIndexVal := v.Visit(assignCtx.Expr(0))
	indexVal, ok := rawIndexVal.(Value)
	if !ok {
		runtimeErr(fmt.Sprintf("invalid index expression for array '%s'", arrName))
	}

	// only unsigned integer indexes are allowed
	index, err := strconv.ParseUint(indexVal.Value, 10, 64)
	if err != nil {
		runtimeErr(fmt.Sprintf("invalid index value for array '%s'", arrName))
	}

	// check if variable exists
	variable, ok := v.symbolTable[arrName]
	if !ok {
		runtimeErr(fmt.Sprintf("array '%s' doesn't exist", arrName))
	}

	// prevent updating immutable variable
	if !variable.isMutable {
		runtimeErr(fmt.Sprintf("cannot reassign, '%s' is immutable", arrName))
	}

	// evaluate the new value expression
	rawVal := v.Visit(assignCtx.Expr(1))
	newVal, ok := rawVal.(Value)
	if !ok {
		runtimeErr("invalid value for reassignment")
	}

	// clean type prefix from array type
	cleanType := stripArrayPrefix(variable.Typ)

	// type check + convert
	if newVal.Typ != cleanType && !(isNumericType(newVal.Typ) && isNumericType(cleanType)) {
		runtimeErr(fmt.Sprintf("type mismatch cannot assign '%s' to array of '%s'", newVal.Typ, cleanType))
	}

	converted, err := convertType(newVal.Value, newVal.Typ, cleanType)
	if err != nil {
		runtimeErr(err.Error())
	}

	updated := converted.(Value)

	// parse array from string
	rawString := strings.Trim(variable.Value, "[]")
	items := strings.Split(rawString, ",")

	if int(index) >= len(items) {
		runtimeErr(fmt.Sprintf("index %d out of bounds for array '%s'", index, arrName))
	}

	// preserving string quotes
	if cleanType == "string" {
		items[index] = fmt.Sprintf(`"%v"`, updated.Value)
	} else {
		items[index] = updated.Value
	}

	// re-serialize array
	serialized := "[" + strings.Join(items, ",") + "]"

	v.symbolTable[arrName] = Value{
		Value:     serialized,
		Typ:       variable.Typ,
		isMutable: true,
	}

	return nil
}
