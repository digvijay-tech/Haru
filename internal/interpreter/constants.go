// Evaluates constants
package interpreter

import (
	"fmt"

	"github.com/digvijay-tech/Haru/internal/parser"
)

// VisitExplicitConstDecl evaluates constants declared with type
func (v *HaruVisitor) VisitExplicitConstDecl(ctx *parser.ConstDeclContext) any {
	varName := ctx.ID().GetText()    // identifier
	varType := ctx.Type_().GetText() // user defined type

	// evaluating expression to get value assigned to the constant
	result := v.Visit(ctx.Expr())
	val, ok := result.(Value)

	if !ok {
		runtimeErr("Invalid constant value")
	}

	// the variable type and literal type are same
	if val.Typ == varType {
		// creating a value instance for storing constant value and its type
		constant := Value{Value: val.Value, Typ: varType}

		// adding to symbol table
		v.symbolTable[varName] = constant

		return constant
	}

	// attempting to convert the type and value of literal/exprVal to a compatible type
	updatedValue, err := convertType(val.Value, val.Typ, varType)
	if err != nil {
		runtimeErr(fmt.Sprintf("type conversion failed for %s: %v", varName, err))
	}

	// asserting type to be Value
	constant := updatedValue.(Value)
	v.symbolTable[varName] = constant

	return constant
}

// VisitImplicitConstDecl evaluates constants declared without a type
func (v *HaruVisitor) VisitImplicitConstDecl(ctx *parser.ConstInferDeclContext) any {
	varName := ctx.ID().GetText() // identifier

	// evaluating expression to get value assigned to the constant
	result := v.Visit(ctx.Expr())
	val, ok := result.(Value)
	if !ok {
		runtimeErr("Invalid constant value")
	}

	// infering type from literal
	// VisitLitExpr will convert numeric literals to either i32 or i64 and floats to f32 or f64
	updatedValue, err := convertType(val.Value, val.Typ, val.Typ)
	if err != nil {
		runtimeErr(fmt.Sprintf("Type inference failed for %s: %v", varName, err))
	}

	constant := updatedValue.(Value)
	v.symbolTable[varName] = constant

	return constant
}
