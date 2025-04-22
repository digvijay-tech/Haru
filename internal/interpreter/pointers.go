package interpreter

import (
	"fmt"
	"strconv"
	"strings"

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

// VisitMutablePointerDecl evaluates mutable pointer declarations
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

// VisitPointerDerefExpr evaluates pointer deferencing expression and returns the value of pointee variable
func (v *HaruVisitor) VisitPointerDerefExpr(ctx *parser.DerefExprContext) any {
	// evaluating pointer expression which should be a variable
	raw := v.Visit(ctx.Expr())
	ptr, ok := raw.(Value)

	if !ok || !strings.HasPrefix(ptr.Typ, "*") {
		runtimeErr("dereference error: not a pointer")
	}

	// getting the variable name stored in the pointer's value
	targetVar := ptr.Value

	// resolving the pointee variable
	val, ok := v.resolve(targetVar)

	if !ok {
		runtimeErr(fmt.Sprintf("dereference error: unknown variable '%s'", targetVar))
	}

	return val
}

// VisitPointerAssignment evaluating pointer reassignment expressions for mutable variables only
func (v *HaruVisitor) VisitPointerAssignment(ctx *parser.PointerAssignStmtStatementContext) any {
	assignCtx := ctx.PointerAssign().(*parser.PointerAssignStmtContext)
	ptrName := assignCtx.ID().GetText()

	// lookup pointer variable in symbol table
	ptr, ok := v.resolve(ptrName)

	if !ok {
		runtimeErr(fmt.Sprintf("undefined pointer '%s'", ptrName))
	}

	// validating pointer type
	if !strings.HasPrefix(ptr.Typ, "*") {
		runtimeErr(fmt.Sprintf("'%s' is not a pointer", ptrName))
	}

	// ensuring pointer itself is mutable
	if !ptr.isMutable {
		runtimeErr(fmt.Sprintf("pointer '%s' is immutable", ptrName))
	}

	// getting the variable this pointer points to
	targetVar := ptr.Value

	// getting the value of pointee
	pointee, ok := v.resolve(targetVar)

	if !ok {
		runtimeErr(fmt.Sprintf("pointer '%s' points to undefined variable '%s'", ptrName, targetVar))
	}

	// ensuring the pointee is mutable
	if !pointee.isMutable {
		runtimeErr(fmt.Sprintf("cannot assign through pointer to immutable variable '%s'", targetVar))
	}

	// evaluating new value being assigned
	raw := v.Visit(assignCtx.Expr())

	newVal, ok := raw.(Value)
	if !ok {
		runtimeErr("invalid value in pointer assignment")
	}

	// type check
	expectedType := pointee.Typ

	if newVal.Typ != expectedType {
		// allowing reassignment for string or i32 to byte
		if expectedType == "byte" && (newVal.Typ == "string" || newVal.Typ == "i32") {
			raw := newVal.Value

			// string to byte, only single character is allowed
			if newVal.Typ == "string" && len(raw) == 1 {
				newVal = Value{Value: strconv.Itoa(int(raw[0])), Typ: "byte"}
			} else {
				// i32 to byte, only 0
				n, err := strconv.ParseUint(raw, 10, 8)

				if err != nil {
					runtimeErr(fmt.Sprintf("cannot convert '%s' to byte", raw))
				}

				newVal = Value{Value: fmt.Sprintf("%d", n), Typ: "byte"}
			}
		} else if isNumericType(newVal.Typ) && isNumericType(expectedType) {
			converted, err := convertType(newVal.Value, newVal.Typ, expectedType)

			if err != nil {
				runtimeErr(err.Error())
			}

			newVal = converted.(Value)
		} else {
			runtimeErr(fmt.Sprintf("type mismatch: expected '%s', got '%s'", expectedType, newVal.Typ))
		}
	}

	// performing the reassignment
	v.assign(targetVar, newVal)

	return nil
}
