// Custom types for Haru visitor
package interpreter

import "github.com/digvijay-tech/Haru/internal/parser"

// Value represents a variable or expression result
type Value struct {
	Value     string    // Literal string value
	Typ       string    // Datatype
	isMutable bool      // Marks as mutable
	Function  *Function // Remains nil if uninitialized
}

// Function represents funcation declaration properties such as parameters, name, return type, and function body
type Param struct {
	name string
	typ  string
}

type Function struct {
	params      []Param
	returnTypes []string
	body        parser.IBlockContext
}

// All supported datatypes in the langauge
var validTypes = map[string]bool{
	"i8": true, "i16": true, "i32": true, "i64": true, "int": true,
	"ui8": true, "ui16": true, "ui32": true, "ui64": true, "uint": true,
	"f32": true, "f64": true,
	"bool": true, "string": true, "byte": true,
}

// Type Category map for numeric type promotion
var typeCategory = map[string]string{
	"i8": "int", "i16": "int", "i32": "int", "i64": "int", "int": "int",
	"ui8": "uint", "ui16": "uint", "ui32": "uint", "ui64": "uint", "uint": "uint",
	"f32": "float", "f64": "float",
}

// Type Rank map for promotion during arithmetic operation
var typeRank = map[string]int{
	"int":   1,
	"uint":  2,
	"float": 3,
}

// Bit Size map for type promotion on edgecases
var bitSize = map[string]int{
	"i8": 8, "i16": 16, "i32": 32, "i64": 64, "int": 64,
	"ui8": 8, "ui16": 16, "ui32": 32, "ui64": 64, "uint": 64,
	"f32": 32, "f64": 64,
}
