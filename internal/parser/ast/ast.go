package ast

// Expression represents any valid expression
type Expression interface {
	ExpNode()
}

// NumberLiteral represents a numeric value
type NumberLiteral struct {
	Value string
	Line  int
	Col   int
}

func (n *NumberLiteral) ExpNode() {}

// BinaryExpression represents an operation like x - y * z
type BinaryExpression struct {
	Left     Expression
	Operator string
	Right    Expression
	Line     int
	Col      int
}

func (b *BinaryExpression) ExpNode() {}

// Identifier represents a variable reference
type Identifier struct {
	Name string
	Line int
	Col  int
}

func (i *Identifier) ExpNode() {}

// Variable Declaration Statements, could be `let` or `mut` statement
type VarStatement struct {
	Identifier string     // Variable Name
	Type       string     // Data Type (optional if inferred)
	Value      Expression // Assigned expression
	Mutable    bool       // Whether the variable is mutable (defined with `mut` keyword)
	Line       int
	Col        int
}

// Variable assignment
type VariableAssignment struct {
	Left  *VarStatement // the variable being assign (`let` | mut)
	Right Expression    // the expression assigned to it
}

func (v *VariableAssignment) ExpNode() {}
