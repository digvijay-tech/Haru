// Helper function
package interpreter

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// isInt returns true if is s is a valid integer
func isInt(s string) bool {
	match, _ := regexp.MatchString(`^-?[0-9]+$`, s)
	return match
}

// isFloat returns true is s is a valid float
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
	// allowing Allow string + string (concatenation)
	if t1 == "string" && t2 == "string" {
		return "string"
	}

	// preventing invalid operations like string + number, or bool + anything
	if t1 == "string" || t2 == "string" || t1 == "bool" || t2 == "bool" {
		runtimeErr(fmt.Sprintf("cannot operate on types '%s' and '%s'", t1, t2))
	}

	// guarding against promoting type categories instead of real types
	if _, ok := bitSize[t1]; !ok {
		runtimeErr(fmt.Sprintf("'%s' is not a concrete type", t1))
	}

	if _, ok := bitSize[t2]; !ok {
		runtimeErr(fmt.Sprintf("'%s' is not a concrete type", t2))
	}

	// now guaranteed t1 and t2 are numeric
	if _, ok := bitSize[t1]; !ok {
		runtimeErr(fmt.Sprintf("'%s' is not a concrete numeric type", t1))
	}

	if _, ok := bitSize[t2]; !ok {
		runtimeErr(fmt.Sprintf("'%s' is not a concrete numeric type", t2))
	}

	// getting categories (numeric only)
	cat1 := typeCategory[t1]
	cat2 := typeCategory[t2]

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

// enforcing range safe type conversions
// will always return Value if error is nil
func convertType(value, fromType, toType string) (any, error) {
	// helper function to check integer range
	checkIntRange := func(val string, min, max int64) error {
		parsedVal, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse %s for type %s", val, toType)
		}

		if parsedVal < min || parsedVal > max {
			return fmt.Errorf("value %s out of range for type %s", val, toType)
		}

		return nil
	}

	// helper function to check unsigned integer range
	checkUintRange := func(val string, max uint64) error {
		parsedVal, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse %s for type %s", val, toType)
		}

		if parsedVal > max {
			return fmt.Errorf("value %s out of range for type %s", val, toType)
		}

		return nil
	}

	// helper function to check float range
	checkFloatRange := func(val string, min, max float64) error {
		parsedVal, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return fmt.Errorf("failed to parse %s for type %s", val, toType)
		}

		if parsedVal < min || parsedVal > max {
			return fmt.Errorf("value %s out of range for type %s", val, toType)
		}

		return nil
	}

	// converting literal based on variable's type
	switch toType {
	case "i8":
		// checks if the value of literal fits within the range
		if err := checkIntRange(value, math.MinInt8, math.MaxInt8); err != nil {
			return nil, err
		}

		// value fits in range, converting value of type string to i8
		val, err := strconv.ParseInt(value, 10, 8) // always returns int64
		if err != nil {
			return nil, err
		}

		// converting int64 to int8
		parsedVal := int8(val)
		return Value{Value: fmt.Sprintf("%d", parsedVal), Typ: "i8"}, nil

	case "i16":
		// checks if the value of literal fits within the range
		if err := checkIntRange(value, math.MinInt16, math.MaxInt16); err != nil {
			return nil, err
		}

		// value fits in range, converting value of type string to i16
		val, err := strconv.ParseInt(value, 10, 16) // always returns int64
		if err != nil {
			return nil, err
		}

		// converting int64 to int16
		parsedVal := int16(val)
		return Value{Value: fmt.Sprintf("%d", parsedVal), Typ: "i16"}, nil

	case "i32":
		// checks if the value of literal fits within the range
		if err := checkIntRange(value, math.MinInt32, math.MaxInt32); err != nil {
			return nil, err
		}

		// value fits in range, converting value of type string to i32
		val, err := strconv.ParseInt(value, 10, 32) // always returns int64
		if err != nil {
			return nil, err
		}

		// converting int64 to int32
		parsedVal := int32(val)
		return Value{Value: fmt.Sprintf("%d", parsedVal), Typ: "i32"}, nil

	case "i64":
		// checks if the value of literal fits within the range
		if err := checkIntRange(value, math.MinInt64, math.MaxInt64); err != nil {
			return nil, err
		}

		// value fits in range, converting value of type string to i64
		val, err := strconv.ParseInt(value, 10, 64) // always returns int64
		if err != nil {
			return nil, err
		}

		return Value{Value: fmt.Sprintf("%d", val), Typ: "i64"}, nil

	case "int":
		// int is 32-bit integer
		// checks if the value of literal fits within the range
		if err := checkIntRange(value, math.MinInt32, math.MaxInt32); err != nil {
			return nil, err
		}

		// value fits in range, converting value of type string to i32
		val, err := strconv.ParseInt(value, 10, 32) // always returns int64
		if err != nil {
			return nil, err
		}

		// converting int64 to int32
		parsedVal := int32(val)
		return Value{Value: fmt.Sprintf("%d", parsedVal), Typ: "int"}, nil

	case "ui8":
		// checks if the value of literal fits within the range
		if err := checkUintRange(value, math.MaxUint8); err != nil {
			return nil, err
		}

		// value fits in range, converting value of type string to ui8
		val, err := strconv.ParseUint(value, 10, 8) // always returns uint64
		if err != nil {
			return nil, err
		}

		// converting uint64 to uint8
		parsedVal := uint8(val)
		return Value{Value: fmt.Sprintf("%d", parsedVal), Typ: "ui8"}, nil

	case "ui16":
		// checks if the value of literal fits within the range
		if err := checkUintRange(value, math.MaxUint16); err != nil {
			return nil, err
		}

		// value fits in range, converting value of type string to ui16
		val, err := strconv.ParseUint(value, 10, 16) // always returns uint64
		if err != nil {
			return nil, err
		}

		// converting uint64 to uint16
		parsedVal := uint16(val)
		return Value{Value: fmt.Sprintf("%d", parsedVal), Typ: "ui16"}, nil

	case "ui32":
		// checks if the value of literal fits within the range
		if err := checkUintRange(value, math.MaxUint32); err != nil {
			return nil, err
		}

		// value fits in range, converting value of type string to ui32
		val, err := strconv.ParseUint(value, 10, 32) // always returns uint64
		if err != nil {
			return nil, err
		}

		// converting uint64 to uint32
		parsedVal := uint32(val)
		return Value{Value: fmt.Sprintf("%d", parsedVal), Typ: "ui32"}, nil

	case "ui64":
		// checks if the value of literal fits within the range
		if err := checkUintRange(value, math.MaxUint64); err != nil {
			return nil, err
		}

		// value fits in range, converting value of type string to ui64
		val, err := strconv.ParseUint(value, 10, 64) // always returns uint64
		if err != nil {
			return nil, err
		}

		return Value{Value: fmt.Sprintf("%d", val), Typ: "ui64"}, nil

	case "uint":
		// uint is 32-bit unsigned integer
		// checks if the value of literal fits within the range
		if err := checkUintRange(value, math.MaxUint32); err != nil {
			return nil, err
		}

		// value fits in range, converting value of type string to ui32
		val, err := strconv.ParseUint(value, 10, 32) // always returns uint64
		if err != nil {
			return nil, err
		}

		// converting uint64 to uint32
		parsedVal := uint32(val)
		return Value{Value: fmt.Sprintf("%d", parsedVal), Typ: "uint"}, nil

	case "f32":
		// checks if the value of literal fits within the range
		// minus sign before math.MaxFloat32 gets the minimum float32 possible
		if err := checkFloatRange(value, -math.MaxFloat32, math.MaxFloat32); err != nil {
			return nil, err
		}

		// value fits in range, converting value of type string to f32
		val, err := strconv.ParseFloat(value, 32) // always returns float64
		if err != nil {
			return nil, err
		}

		// converting float64 to float32
		parsedVal := float32(val)
		return Value{Value: fmt.Sprintf("%f", parsedVal), Typ: "f32"}, nil

	case "f64":
		// checks if the value of literal fits within the range
		// minus sign before math.MaxFloat64 gets the minimum float64 possible
		if err := checkFloatRange(value, -math.MaxFloat64, math.MaxFloat64); err != nil {
			return nil, err
		}

		// value fits in range, converting value of type string to f64
		val, err := strconv.ParseFloat(value, 64) // always returns float64
		if err != nil {
			return nil, err
		}

		return Value{Value: fmt.Sprintf("%f", val), Typ: "f64"}, nil

	case "bool":
		if value == "true" || value == "false" {
			return Value{Value: value, Typ: "bool"}, nil
		}

		return nil, fmt.Errorf("invalid boolean value %s", value)

	case "string":
		return Value{Value: value, Typ: "string"}, nil

	case "byte":
		if len(value) == 1 {
			// single character string, use its ASCII code
			return Value{Value: strconv.Itoa(int(value[0])), Typ: "byte"}, nil
		}

		// try parsing it as a number
		val, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid byte value: %v", err)
		}

		// range of byte in numeric form must be equal to uint8
		if val < 0 || val > math.MaxUint8 {
			return nil, fmt.Errorf("value %d out of range for byte", val)
		}

		return Value{Value: strconv.Itoa(int(val)), Typ: "byte"}, nil
	}

	// type is unknown
	return nil, fmt.Errorf("unsupported type conversion from %s to %s", fromType, toType)
}

// zeroValueFor takes in a type and produces Value struct for that type's zero value. It does not change isMutable to true
func zeroValueFor(typ string) (Value, error) {
	// all numeric type will (0)
	if isNumericType(typ) {
		val := Value{Value: "0", Typ: typ}
		return val, nil
	}

	// zero type of bool is false
	if typ == "bool" {
		val := Value{Value: "false", Typ: typ}
		return val, nil
	}

	// zero type of string is empty string literal
	if typ == "string" {
		val := Value{Value: "", Typ: typ}
		return val, nil
	}

	// zero type of byte is (0)
	if typ == "byte" {
		val := Value{Value: "0", Typ: typ}
		return val, nil
	}

	// no type match found
	// returning error with empty Value struct as nil can't be returned
	return Value{}, fmt.Errorf("unsupported type '%s' in 'mut' declaration", typ)
}
