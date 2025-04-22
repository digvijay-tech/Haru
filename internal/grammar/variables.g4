grammar Variables;


varDecl: 'let' ID ':' type '=' expr ';'          # LetDecl
       | 'let' ID '=' expr ';'                   # LetInferDecl
       | 'mut' ID ':' type ('=' expr)? ';'       # MutDecl
       | 'mut' ID '=' expr ';'                   # MutInferDecl
       | 'const' ID ':' type '=' expr ';'        # ConstDecl
       | 'const' ID '=' expr ';'                 # ConstInferDecl
       | 'let' ID ':' '*' type '=' '&' ID ';'    # ImmutablePointerDecl
       | 'mut' ID ':' '*' type '=' '&' ID ';'    # MutablePointerDecl
       ;


type: 'i8'    # I8Type
    | 'i16'   # I16Type
    | 'i32'   # I32Type
    | 'i64'   # I64Type
    | 'int'   # IntType
    | 'ui8'   # UI8Type
    | 'ui16'  # UI16Type
    | 'ui32'  # UI32Type
    | 'ui64'  # UI64Type
    | 'uint'  # UIntType
    | 'f32'   # F32Type
    | 'f64'   # F64Type
    | 'bool'  # BoolType
    | 'string'# StringType
    | 'byte'  # ByteType
    ;

