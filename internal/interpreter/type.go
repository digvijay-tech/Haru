// Custom types for Haru visitor
package interpreter

// Value represents a variable or expression result
type Value struct {
	Value string // Literal string value
	Typ   string // Datatype
}

// Type Category map for type promotion
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
