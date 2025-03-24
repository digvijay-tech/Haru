grammar haru;

import Variables, Expressions, Output, ControlFlow;

program: statement* ;

statement: varDecl ';' | assign ';' | printStmt ';' | ifStmt ;

WS: [ \t\r\n]+ -> skip ;