grammar haru;

import Output, Expressions, Variables;

program: statement* ;

statement: printStmt # PrintStmtStatement
         | varDecl   # VarDeclStatement
         | assign     # AssignStmtStatement ;

WS: [ \t\r\n]+ -> skip ;
