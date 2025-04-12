package interpreter

// represents variable's value and metadata
type Value struct {
	Value   string
	Typ     string
	IsMut   bool
	IsConst bool
}
