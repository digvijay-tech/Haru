grammar Variables;

varDecl: 'let' ID ':' type '=' expr # LetDecl
       | 'let' ID '=' expr          # LetInferDecl
       | 'mut' ID ':' type ('=' expr)? # MutDecl
       | 'mut' ID '=' expr          # MutInferDecl
       | 'const' ID ':' type '=' expr # ConstDecl
       | 'const' ID '=' expr        # ConstInferDecl ;

type: 'i8'    # I8Type
    | 'i16'   # I16Type
    | 'i32'   # I32Type
    | 'i64'   # I64Type
    | 'int'   # IntType
    | 'ui8'   # UI8Type
    | 'ui16'  # UI16Type
    | 'ui32'  # UI32Type
    | 'ui64'  # UI64Type
    | 'ui'    # UIType
    | 'f32'   # F32Type
    | 'f64'   # F64Type
    | 'bool'  # BoolType
    | 'string'# StringType
    | 'byte'  # ByteType ;

literal: NUMBER       # IntLiteral
       | FLOAT        # FloatLiteral
       | 'true'       # TrueLiteral
       | 'false'      # FalseLiteral
       | STRING       # StringLiteral
       | BYTE         # ByteLiteral ;

ID: [a-zA-Z][a-zA-Z0-9]* ;
NUMBER: [0-9]+ ;
FLOAT: [0-9]+ '.' [0-9]+ ;
STRING: ('"' (ESC|.)*? '"') | ('\'' (ESC|.)*? '\'') ;
BYTE: '0b' [0-1]+ ;
fragment ESC: '\\' ['"\\] ;