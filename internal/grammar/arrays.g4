grammar Arrays;



arrayDecl
    : constArrayDecl
    | letArrayDecl
    | mutArrayDecl
    ;


constArrayDecl
    : 'const' ID ':' arrayType '=' arrayLiteral ';'                  # ConstExplicitArrayDecl
    | 'const' ID '=' arrayLiteral ';'                                # ConstImplicitArrayDecl
    ;


letArrayDecl
    : 'let' ID ':' arrayType '=' arrayLiteral ';'                   # LetExplicitArrayDecl
    | 'let' ID '=' arrayLiteral ';'                                 # LetImplicitArrayDecl
    ;


mutArrayDecl
    : 'mut' ID ':' fixedArrayType '=' arrayLiteral ';'              # MutFixedArrayWithInit
    | 'mut' ID ':' fixedArrayType ';'                               # MutFixedArrayNoInit
    | 'mut' ID ':' arrayType '=' arrayLiteral ';'                   # MutArrayExplicitWithInit
    | 'mut' ID ':' arrayType ';'                                    # MutArrayExplicitNoInit
    | 'mut' ID '=' arrayLiteral ';'                                 # MutArrayImplicit
    ;


arrayType
    : '[' ']' type
    ;


fixedArrayType
    : '[' NUMBER ']' type
    ;


arrayLiteral
    : '[' expr (',' expr)*? ']'                                  # ArrayLiteralExprList
    | '[' ']'                                                    # EmptyArr
    ;


arrayItemAssign
    : ID '[' expr ']' '=' expr ';'                               # ArrayIndexAssign
    ;


arrayReassign
    : ID '=' arrayLiteral ';'
    ;

