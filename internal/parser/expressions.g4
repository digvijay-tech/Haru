grammar Expressions;

expr: '(' expr ')'          # ParenExpr
    | expr '**' expr        # ExpExpr
    | expr '*' expr         # MulExpr
    | expr '/' expr         # DivExpr
    | expr '%' expr         # ModExpr
    | expr '+' expr         # AddExpr
    | expr '-' expr         # SubExpr
    | ID                    # VarExpr
    | literal               # LitExpr ;

assign: ID '=' expr         # AssignStmt ;

ID: [a-zA-Z][a-zA-Z0-9]* ;
literal: NUMBER             # IntLiteral
       | FLOAT              # FloatLiteral
       | 'true'             # TrueLiteral
       | 'false'            # FalseLiteral
       | STRING             # StringLiteral
       | BYTE               # ByteLiteral ;

NUMBER: [0-9]+ ;
FLOAT: [0-9]+ '.' [0-9]+ ;
STRING: ('"' (ESC|.)*? '"') | ('\'' (ESC|.)*? '\'') ;
BYTE: '0b' [0-1]+ ;
fragment ESC: '\\' ['"\\] ;
