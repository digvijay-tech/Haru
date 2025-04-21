package interpreter

import (
	"fmt"
	"strconv"

	"github.com/digvijay-tech/Haru/internal/parser"
)

// VisitFunctionBlock evaluates function body by visiting every statements node
func (v *HaruVisitor) VisitFunctionBlock(ctx *parser.BlockContext) any {
	for _, stmt := range ctx.AllStatement() {
		v.Visit(stmt)
	}

	return nil
}

// VisitFunctionDeclStatement evaluates function declaration by extracting function name,
// capturing return signature, and storing it in the globally defined scope stack
func (v *HaruVisitor) VisitFunctionDeclStatement(ctx *parser.FunctionDeclStatementContext) any {
	funcCtx := ctx.FunctionDecl().(*parser.FunctionDeclContext)
	funcName := funcCtx.ID().GetText()

	// making sure function name is unique in the current scope
	scope := v.currentScope()

	if _, exists := scope[funcName]; exists {
		runtimeErr(fmt.Sprintf("function '%s' already declared in this scope", funcName))
	}

	// parsing function params
	var params []Param

	if funcCtx.ParamList() != nil {
		for _, p := range funcCtx.ParamList().AllParam() {
			paramName := p.ID().GetText()
			paramType := p.Type_().GetText()

			// making sure given param type is a valid haru datatype
			if !validTypes[paramType] {
				runtimeErr(fmt.Sprintf("invalid parameter type '%s' in function '%s'", paramType, funcName))
			}

			params = append(params, Param{name: paramName, typ: paramType})
		}
	}

	// parsing return types
	var returnTypes []string

	if funcCtx.ReturnSignature() != nil {
		for _, t := range funcCtx.ReturnSignature().AllType_() {
			typ := t.GetText()

			// making sure given return type is a valid datatype
			if !validTypes[typ] {
				runtimeErr(fmt.Sprintf("invalid return type '%s' in function '%s'", typ, funcName))
			}

			returnTypes = append(returnTypes, t.GetText())
		}
	}

	// wrapping the function in a value to store it on the scope stack
	fn := &Function{
		params:      params,
		returnTypes: returnTypes,
		body:        funcCtx.Block(),
	}

	v.declare(funcName, Value{
		Typ:      "function",
		Function: fn,
	})

	return nil
}

// VisitFunctionCallStatement evaluates function calls which are used as statements
func (v *HaruVisitor) VisitFunctionCallStatement(ctx *parser.FunctionCallStatementContext) any {
	funcName := ctx.FunctionCall().ID().GetText()

	// checking if function is defined and is of type function
	val, ok := v.resolve(funcName)

	if !ok || val.Function == nil || val.Typ != "function" {
		runtimeErr(fmt.Sprintf("'%s' is not a callable function", funcName))
	}

	fn := val.Function

	// evaluatign arguments
	var args []Value

	if ctx.FunctionCall().ArgumentList() != nil {
		for _, expr := range ctx.FunctionCall().ArgumentList().AllExpr() {
			value := v.Visit(expr).(Value)

			args = append(args, value)
		}
	}

	// calling function does not have same number of arguments as function params
	if len(args) != len(fn.params) {
		runtimeErr(fmt.Sprintf("function '%s' expects %d arguments, got %d", funcName, len(fn.params), len(args)))
	}

	for i, p := range fn.params {
		actual := args[i].Typ
		expected := p.typ

		if actual != expected {
			// standardizing the argument when param is expecting a byte but received either string or i32
			if expected == "byte" && (actual == "string" || actual == "i32") {
				argVal := args[i].Value

				// argument is a string
				if len(argVal) == 1 {
					// single character string, use its ASCII code
					args[i] = Value{Value: strconv.Itoa(int(argVal[0])), Typ: "byte"}

					// jump to next iteration
					continue
				}

				// argument is i32
				n, err := strconv.ParseUint(argVal, 10, 8) // parsing into a byte 8bit unsigned integer

				if err != nil {
					runtimeErr(fmt.Sprintf("cannot convert string '%s' to byte for parameter '%s'", argVal, p.name))
				}

				// createing a new Value with byte type
				args[i] = Value{
					Value: fmt.Sprintf("%d", n),
					Typ:   "byte",
				}

				// jump to next iteration
				continue
			}

			// making sure non-numeric types are not converted
			if !(isNumericType(actual) && isNumericType(expected)) {
				runtimeErr(fmt.Sprintf("parameter '%s' expects '%s', got '%s'", p.name, expected, actual))
			}

			converted, err := convertType(args[i].Value, actual, expected)

			if err != nil {
				runtimeErr(fmt.Sprintf("parameter '%s' expects '%s', got '%s'", p.name, expected, actual))
			}

			args[i] = converted.(Value)
		}
	}

	// calling function in a new scope
	v.pushScope()

	// declaring local variables in the function body from arguments
	for i, p := range fn.params {
		v.declare(p.name, args[i])
	}

	// evaluating function body
	v.Visit(fn.body)

	// removing scope when function ends
	v.popScope()

	return nil
}

// VisitFunctionCallExpression evaluates function calls used as expressions
func (v *HaruVisitor) VisitFunctionCallExpression(ctx *parser.FunctionCallExprContext) any {
	funcName := ctx.FunctionCall().ID().GetText()

	// checking if function is defined and is of type function
	val, ok := v.resolve(funcName)

	if !ok || val.Function == nil || val.Typ != "function" {
		runtimeErr(fmt.Sprintf("'%s' is not a callable function", funcName))
	}

	fn := val.Function

	// evaluatign arguments
	var args []Value

	if ctx.FunctionCall().ArgumentList() != nil {
		for _, expr := range ctx.FunctionCall().ArgumentList().AllExpr() {
			value := v.Visit(expr).(Value)

			args = append(args, value)
		}
	}

	// calling function does not have same number of arguments as function params
	if len(args) != len(fn.params) {
		runtimeErr(fmt.Sprintf("function '%s' expects %d arguments, got %d", funcName, len(fn.params), len(args)))
	}

	for i, p := range fn.params {
		actual := args[i].Typ
		expected := p.typ

		if actual != expected {
			// standardizing the argument when param is expecting a byte but received either string or i32
			if expected == "byte" && (actual == "string" || actual == "i32") {
				argVal := args[i].Value

				// argument is a string
				if len(argVal) == 1 {
					// single character string, use its ASCII code
					args[i] = Value{Value: strconv.Itoa(int(argVal[0])), Typ: "byte"}

					// jump to next iteration
					continue
				}

				// argument is i32
				n, err := strconv.ParseUint(argVal, 10, 8) // parsing into a byte 8bit unsigned integer

				if err != nil {
					runtimeErr(fmt.Sprintf("cannot convert string '%s' to byte for parameter '%s'", argVal, p.name))
				}

				// createing a new Value with byte type
				args[i] = Value{
					Value: fmt.Sprintf("%d", n),
					Typ:   "byte",
				}

				// jump to next iteration
				continue
			}

			// making sure non-numeric types are not converted
			if !(isNumericType(actual) && isNumericType(expected)) {
				runtimeErr(fmt.Sprintf("parameter '%s' expects '%s', got '%s'", p.name, expected, actual))
			}

			converted, err := convertType(args[i].Value, actual, expected)

			if err != nil {
				runtimeErr(fmt.Sprintf("parameter '%s' expects '%s', got '%s'", p.name, expected, actual))
			}

			args[i] = converted.(Value)
		}
	}

	// calling function in a new scope
	v.pushScope()

	// declaring local variables in the function body from arguments
	for i, p := range fn.params {
		v.declare(p.name, args[i])
	}

	// evaluating function body
	v.Visit(fn.body)

	// removing scope when function ends
	v.popScope()

	// Placeholder return until `return` is supported
	return Value{
		Value: "0",
		Typ:   "i32", // or any default
	}
}
