package interpreter

import (
	"fmt"

	"github.com/digvijay-tech/Haru/internal/parser"
)

// VisitImmutablePointerDecl evaluates immutable pointer declarations
func (v *HaruVisitor) VisitImmutablePointerDecl(ctx *parser.ImmutablePointerDeclContext) any {
	ptrName := ctx.ID(0).GetText()
	targetVar := ctx.ID(1).GetText()
	ptrType := ctx.Type_().GetText()

	// check if the target variable exists
	targetVal, ok := v.resolve(targetVar)
	if !ok {
		runtimeErr(fmt.Sprintf("cannot take address of undefined variable '%s'", targetVar))
	}

	// ensuring the type is exactly the same
	if targetVal.Typ != ptrType {
		runtimeErr(fmt.Sprintf("type mismatch pointer to '%s' declared, but '%s' has type '%s'", ptrType, targetVar, targetVal.Typ))
	}

	// adding pointer to symbol table
	v.declare(ptrName, Value{
		Value: targetVar,     // using the variable name as reference
		Typ:   "*" + ptrType, // pointer type
	})

	return nil
}

func (v *HaruVisitor) VisitMutablePointerDecl(ctx *parser.MutablePointerDeclContext) any {
	ptrName := ctx.ID(0).GetText()
	targetVar := ctx.ID(1).GetText()
	ptrType := ctx.Type_().GetText()

	// check if the target variable exists
	targetVal, ok := v.resolve(targetVar)
	if !ok {
		runtimeErr(fmt.Sprintf("cannot take address of undefined variable '%s'", targetVar))
	}

	// ensuring the type is exactly the same
	if targetVal.Typ != ptrType {
		runtimeErr(fmt.Sprintf("type mismatch pointer to '%s' declared, but '%s' has type '%s'", ptrType, targetVar, targetVal.Typ))
	}

	// target variable must also be mutable
	if !targetVal.isMutable {
		runtimeErr(fmt.Sprintf("cannot create mutable pointer to immutable variable '%s'", targetVar))
	}

	// adding pointer to symbol table
	v.declare(ptrName, Value{
		Value:     targetVar,     // using the variable name as reference
		Typ:       "*" + ptrType, // pointer type
		isMutable: true,          // allowing read and write
	})

	return nil
}
