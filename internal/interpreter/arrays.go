package interpreter

import (
	"fmt"

	"github.com/digvijay-tech/Haru/internal/parser"
)

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
	fmt.Println("From VisitConstExplicitArrayDecl")
	return nil
}

// VisitConstImplicitArrayDecl evaluates array declared with const and but type is inferred by array literal
func (v *HaruVisitor) VisitConstImplicitArrayDecl(ctx *parser.ConstImplicitArrayDeclContext) any {
	fmt.Println("From VisitConstImplicitArrayDecl")
	return nil
}
