grammar haru;

import Variables, Expressions, Output;

program: statement* ;

statement: varDecl ';' | assign ';' | printStmt ';' ;

WS: [ \t\r\n]+ -> skip ;
