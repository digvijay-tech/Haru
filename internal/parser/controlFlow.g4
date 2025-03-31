grammar ControlFlow;

ifStmt: 'if' expr '{' ifBody=statement* '}' ('else' '{' elseBody=statement* '}')? # IfStatement ;

ID: [a-zA-Z][a-zA-Z0-9]* ;
