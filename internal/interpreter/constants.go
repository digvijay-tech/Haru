// Evaluates constants
package interpreter

import (
	"fmt"

	"github.com/digvijay-tech/Haru/internal/parser"
)

func (v *HaruVisitor) VisitConstDecl(ctx *parser.ConstDeclContext) any {
	varName := ctx.ID().GetText()    // identifier
	varType := ctx.Type_().GetText() // user defined type

	// evaluating expression to get value assigned to the constant
	exprVal := v.Visit(ctx.Expr()).(Value)

	// the variable type and literal type are same
	if exprVal.Typ == varType {
		// creating a value instance for storing constant value and its type
		constant := Value{Value: exprVal.Value, Typ: varType}

		// adding to symbol table
		v.symbolTable[varName] = constant

		return constant
	}

	// attempting to convert the type and value of literal/exprVal to a compatible type
	updatedValue, err := convertType(exprVal.Value, exprVal.Typ, varType)
	if err != nil {
		return fmt.Errorf("type conversion failed for %s: %v", varName, err)
	}

	// asserting type to be Value
	constant := updatedValue.(Value)
	v.symbolTable[varName] = constant

	return constant
}
