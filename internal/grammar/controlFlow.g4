grammar ControlFlow;

ifStmt
    : 'if' '(' expr ')' block                             # IfBlockOnly
    | 'if' '(' expr ')' block elseIfBlock* elseBlock?     # IfElseChain
    ;

elseIfBlock
    : 'else' 'if' '(' expr ')' block
    ;

elseBlock
    : 'else' block
    ;

block
    : '{' statement* '}'
    ;


ID: [a-zA-Z][a-zA-Z0-9]* ;
