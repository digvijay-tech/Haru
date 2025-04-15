grammar Expressions;

expr: '!' expr              # NotExpr
    | '(' expr ')'          # ParenExpr
    | expr '**' expr        # ExpExpr
    | expr '*' expr         # MulExpr
    | expr '/' expr         # DivExpr
    | expr '%' expr         # ModExpr
    | expr '+' expr         # AddExpr
    | expr '-' expr         # SubExpr
    | expr '<' expr         # LtExpr
    | expr '>' expr         # GtExpr
    | expr '<=' expr        # LeExpr
    | expr '>=' expr        # GeExpr
    | expr '==' expr        # EqExpr
    | expr '!=' expr        # NeExpr
    | expr '&&' expr        # AndExpr
    | expr '||' expr        # OrExpr
    | '[' (expr (',' expr)*)? ']' # ArrayExpr
    | ID                    # VarExpr
    | literal               # LitExpr ;

assign: ID '=' expr         # AssignStmt ;

literal: ('-')? NUMBER      # IntLiteral
       | ('-')? FLOAT       # FloatLiteral
       | 'true'             # TrueLiteral
       | 'false'            # FalseLiteral
       | STRING             # StringLiteral
       | BYTE               # ByteLiteral ;

ID: [a-zA-Z][a-zA-Z0-9]* ;
NUMBER: [0-9]+ ;
FLOAT: [0-9]+ '.' [0-9]+ ;
STRING: ('"' (ESC|.)*? '"') | ('\'' (ESC|.)*? '\'') ;
BYTE: '0b' [0-1]+ ;
fragment ESC: '\\' ['"\\] ;

