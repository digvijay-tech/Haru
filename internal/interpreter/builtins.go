package interpreter

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/digvijay-tech/Haru/internal/parser"
)

// VisitLenFunction evaluates the built-in len(v) function and returns the length of v
func (v *HaruVisitor) VisitLenFunction(ctx *parser.LenFunctionExprContext) any {
	lenCtx := ctx.LenFunction()
	arg, ok := v.Visit(lenCtx.Expr()).(Value)

	if !ok {
		runtimeErr("invalid value for len()")
	}

	if arg.Typ == "string" {
		return Value{
			Value: fmt.Sprintf("%d", utf8.RuneCountInString(arg.Value)),
			Typ:   "i32",
		}
	}

	if isArrayType(arg.Typ) {
		trimmed := strings.Trim(arg.Value, "[]")

		if trimmed == "" {
			return Value{
				Value: "0",
				Typ:   "i32",
			}
		}

		items := strings.Split(trimmed, ",")

		return Value{
			Value: fmt.Sprintf("%d", len(items)),
			Typ:   "i32",
		}
	}

	// neither string nor array
	runtimeErr(fmt.Sprintf("len() not supported for type '%s'", arg.Typ))
	return nil
}

// VisitInputFunction
func (v *HaruVisitor) VisitInputFunction(ctx *parser.InputFunctionExprContext) any {
	prompt := ""

	if ctx.InputFunction().STRING() != nil {
		prompt = stripQuotes(ctx.InputFunction().STRING().GetText())
	}

	// display prompt if exists
	if prompt != "" {
		fmt.Print(prompt)
	}

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	// removing extra whitespaces before returning
	input = strings.TrimSpace(input)

	if err != nil {
		runtimeErr("failed to read input")
	}

	return Value{
		Value: input,
		Typ:   "string",
	}
}
