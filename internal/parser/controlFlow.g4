grammar ControlFlow;

ifStmt: 'if' expr '{' statement* '}' # IfStatement ;

ID: [a-zA-Z][a-zA-Z0-9]* ;
