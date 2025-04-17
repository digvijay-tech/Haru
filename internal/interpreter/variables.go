package interpreter

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/digvijay-tech/Haru/internal/parser"
)

// VisitVarDeclStatement passes control to actual const, let, and mut declaration handlers
func (v *HaruVisitor) VisitVarDeclStatement(ctx *parser.VarDeclStatementContext) any {
	decl := ctx.GetChild(0).(antlr.ParseTree)
	return v.Visit(decl)
}

// VisitExplicitLetDecl evaluates immutable let variables declared with type
// func (v *HaruVisitor) VisitExplicitLetDecl(ctx *parser.LetDeclContext) any {
// 	fmt.Println("From VisitExplicitLetDecl")
// 	varName := ctx.ID().GetText()
// 	varType := ctx.Type_().GetText()

// 	// evaluating the expression to get Value
// 	result := v.Visit(ctx.Expr())
// 	val, ok := result.(Value)
// 	if !ok {
// 		runtimeErr("Invalid value for 'let' declaration")
// 	}

// 	// the variable type and literal type are same
// 	if val.Typ == varType {
// 		// creating a value instance for storing let value and its type
// 		let := Value{Value: val.Value, Typ: varType}

// 		// adding to symbol table
// 		v.symbolTable[varName] = let
// 		fmt.Println(let)
// 		return let
// 	}

// 	// attempting to convert the type and value of literal/exprVal to a compatible type
// 	updatedValue, err := convertType(val.Value, val.Typ, varType)
// 	if err != nil {
// 		runtimeErr(fmt.Sprintf("type conversion failed for %s: %v", varName, err))
// 	}
// 	fmt.Println(updatedValue)
// 	// asserting type to be Value
// 	let := updatedValue.(Value)
// 	v.symbolTable[varName] = let

// 	return let
// }
