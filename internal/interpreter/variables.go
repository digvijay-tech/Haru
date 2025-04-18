package interpreter

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/digvijay-tech/Haru/internal/parser"
)

// VisitVarDeclStatement passes control to actual const, let, and mut declaration handlers
func (v *HaruVisitor) VisitVarDeclStatement(ctx *parser.VarDeclStatementContext) any {
	decl := ctx.GetChild(0).(antlr.ParseTree)
	return v.Visit(decl)
}

// VisitExplicitLetDecl evaluates immutable let variables declared with type
func (v *HaruVisitor) VisitExplicitLetDecl(ctx *parser.LetDeclContext) any {
	varName := ctx.ID().GetText()
	varType := ctx.Type_().GetText()

	// evaluating the expression to get Value
	result := v.Visit(ctx.Expr())
	val, ok := result.(Value)
	if !ok {
		runtimeErr("Invalid value for 'let' declaration")
	}

	// the variable type and literal type are same
	if val.Typ == varType {
		// creating a value instance for storing let value and its type
		let := Value{Value: val.Value, Typ: varType}

		// adding to symbol table
		v.symbolTable[varName] = let
		return let
	}

	// attempting to convert the type and value of literal/exprVal to a compatible type
	updatedValue, err := convertType(val.Value, val.Typ, varType)
	if err != nil {
		runtimeErr(fmt.Sprintf("type conversion failed for %s: %v", varName, err))
	}

	// asserting type to be Value
	let := updatedValue.(Value)
	v.symbolTable[varName] = let

	return let
}

// VisitImplicitConstDecl evaluates constants declared without a type
func (v *HaruVisitor) VisitImplicitLetDecl(ctx *parser.LetInferDeclContext) any {
	varName := ctx.ID().GetText()

	// evaluating expression to get value assigned to the let
	result := v.Visit(ctx.Expr())
	val, ok := result.(Value)
	if !ok {
		runtimeErr("Invalid let value")
	}

	// infering type from literal
	// VisitLitExpr will convert numeric literals to either i32 or i64 and floats to f32 or f64
	updatedValue, err := convertType(val.Value, val.Typ, val.Typ)
	if err != nil {
		runtimeErr(fmt.Sprintf("Type inference failed for %s: %v", varName, err))
	}

	let := updatedValue.(Value)
	v.symbolTable[varName] = let

	return let
}

// VisitExplicitMutDecl evaluates mut declared with a type
func (v *HaruVisitor) VisitExplicitMutDecl(ctx *parser.MutDeclContext) any {
	varName := ctx.ID().GetText()
	varType := ctx.Type_().GetText()

	// the variable has an initializer
	if ctx.Expr() != nil {
		result := v.Visit(ctx.Expr())
		val, ok := result.(Value)
		if !ok {
			runtimeErr(fmt.Sprintf("invalid value in 'mut' declaration for '%s'", varName))
		}

		// converting to declared type
		updatedVal, err := convertType(val.Value, val.Typ, varType)
		if err != nil {
			runtimeErr(fmt.Sprintf("type conversion failed for '%s': %v", varName, err))
		}

		mut := updatedVal.(Value)
		mut.isMutable = true
		v.symbolTable[varName] = mut

		return mut
	}

	// variable is uninitialized
	// generating zero value based on type declared
	zeroedMut, err := zeroValueFor(varType)

	if err != nil {
		runtimeErr(err.Error())
	}

	// marking as mutable and adding to symbol table
	zeroedMut.isMutable = true
	v.symbolTable[varName] = zeroedMut

	return zeroedMut
}
