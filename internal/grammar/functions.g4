grammar Functions;



functionDecl
    : 'function' ID '(' paramList? ')' returnSignature? block
    ;


paramList
    : param (',' param)*
    ;


param
    : ID ':' type
    ;


returnSignature
    : '<' type (',' type)* '>'
    ;


returnStmt
    : 'return' exprList? ';'
    ;


exprList
    : expr (',' expr)*
    ;


block
    : '{' statement* '}'
    ;


functionCall
    : ID '(' argumentList? ')'
    ;


argumentList
    : expr (',' expr)* ;

