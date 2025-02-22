package tokenizer

type TokenType string

const (
	IDENTIFIER   TokenType = "IDENTIFIER"
	KEYWORD      TokenType = "KEYWORD"
	OPERATOR     TokenType = "OPERATOR"
	NUMBER       TokenType = "NUMBER"
	STRING       TokenType = "STRING"
	PUNCTUATION  TokenType = "PUNCTUATION"
	DATATYPE     TokenType = "DATATYPE"
	TUPLE_TYPE   TokenType = "TUPLETYPE"
	UNIDENTIFIED TokenType = "UNIDENTIFIED"
	ERROR        TokenType = "ERROR"
	EOF          TokenType = "EOF"
)

type Token struct {
	Type  TokenType
	Value string
	Line  int
	Col   int
}

var KeywordsTable = map[string]TokenType{
	"let": KEYWORD, "mut": KEYWORD, "fn": KEYWORD, "return": KEYWORD,
	"if": KEYWORD, "else": KEYWORD, "elseif": KEYWORD, "match": KEYWORD,
	"while": KEYWORD, "for": KEYWORD, "import": KEYWORD, "export": KEYWORD,
	"break": KEYWORD, "continue": KEYWORD, "true": KEYWORD, "false": KEYWORD,
	"struct": KEYWORD, "enum": KEYWORD, "const": KEYWORD, "type": KEYWORD,
	"null": KEYWORD, "spawn": KEYWORD, "await": KEYWORD, "lock": KEYWORD,
	"from": KEYWORD, "in": KEYWORD,
}

var DatatypesTable = map[string]TokenType{
	"int": DATATYPE, "uint": DATATYPE, "i8": DATATYPE, "i16": DATATYPE, "i32": DATATYPE, "i64": DATATYPE,
	"ui8": DATATYPE, "ui16": DATATYPE, "ui32": DATATYPE, "ui64": DATATYPE, "f32": DATATYPE, "f64": DATATYPE,
	"char": DATATYPE, "bool": DATATYPE, "string": DATATYPE, "byte": DATATYPE, "map": DATATYPE,
	"void": DATATYPE,
}
