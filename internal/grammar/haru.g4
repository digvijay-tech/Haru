grammar haru;

import Output, Expressions, Variables;

program: statement* ;

statement: printStmt # PrintStmtStatement
         | varDecl   # VarDeclStatement ;

WS: [ \t\r\n]+ -> skip ;
