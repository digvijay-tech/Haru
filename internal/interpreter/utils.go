// Helper function
package interpreter

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func isInt(s string) bool {
	match, _ := regexp.MatchString(`^-?[0-9]+$`, s)
	return match
}

func isFloat(s string) bool {
	match, _ := regexp.MatchString(`^-?[0-9]+\.[0-9]+$`, s)
	return match
}

func isByte(s string) bool {
	return strings.HasPrefix(s, "0b")
}

func isString(s string) bool {
	return strings.HasPrefix(s, "\"") || strings.HasPrefix(s, "'")
}

func stripQuotes(s string) string {
	if len(s) >= 2 && (s[0] == '"' || s[0] == '\'') {
		return s[1 : len(s)-1]
	}
	return s
}

func convertToFloat64(val string) (float64, error) {
	return strconv.ParseFloat(val, 64)
}

func isNumericType(t string) bool {
	_, ok := typeCategory[t]
	return ok
}

func promoteType(t1, t2 string) string {
	// guarding against promoting type categories instead of real types
	if _, ok := bitSize[t1]; !ok {
		log.Fatalf("Type error: '%s' is not a concrete type", t1)
	}

	if _, ok := bitSize[t2]; !ok {
		log.Fatalf("Type error: '%s' is not a concrete type", t2)
	}

	// getting the category of given types
	cat1, ok1 := typeCategory[t1]
	cat2, ok2 := typeCategory[t2]

	// handles concatenation type
	if t1 == "string" && t2 == "string" {
		return "string"
	}

	// preventing invalid operations like string + number, or bool + anything
	if t1 == "string" || t2 == "string" || t1 == "bool" || t2 == "bool" {
		log.Fatalf("Type error: cannot operate on types '%s' and '%s'", t1, t2)
	}

	// fallback for non-numeric types
	if !ok1 || !ok2 {
		log.Fatalf("Type error: unknown types '%s' and '%s'", t1, t2)
	}

	// getting the rank of both categories
	r1 := typeRank[cat1]
	r2 := typeRank[cat2]

	// promoting to higher ranked category (int < uint < float)
	if r1 > r2 {
		return t1 // type on left is big
	}

	if r2 > r1 {
		return t2 // type on right is big
	}

	// same category choose wider bit size
	if bitSize[t1] > bitSize[t2] {
		return t1
	}

	if bitSize[t2] > bitSize[t1] {
		return t2
	}

	// same width return t2 by default (right-side wins)
	return t2
}

// compareNumbers will compare numeric types using the given operator
func compareNumbers(op string, leftVal, rightVal Value) bool {
	l, _ := convertToFloat64(leftVal.Value)
	r, _ := convertToFloat64(rightVal.Value)

	switch op {
	case "<":
		return l < r
	case "<=":
		return l <= r
	case ">":
		return l > r
	case ">=":
		return l >= r
	}

	// no match
	log.Fatalf("Unknown operator: %s", op)
	return false
}
