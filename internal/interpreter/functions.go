package interpreter

import (
	"fmt"

	"github.com/digvijay-tech/Haru/internal/parser"
)

// VisitFunctionDeclStatement evaluates function declaration by extracting function name,
// capturing return signature, and storing it in the globally defined function table
func (v *HaruVisitor) VisitFunctionDeclStatement(ctx *parser.FunctionDeclStatementContext) any {
	fname := ctx.FunctionDecl().ID().GetText()

	fmt.Println(fname)

	return nil
}
