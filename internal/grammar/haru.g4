grammar haru;

import Output, Expressions;

program: statement* ;

statement: printStmt ;

WS: [ \t\r\n]+ -> skip ;
