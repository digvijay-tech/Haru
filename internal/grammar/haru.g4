grammar haru;

import Output, Expressions, Variables, ControlFlow;

program: statement* ;

statement: printStmt # PrintStmtStatement
         | varDecl   # VarDeclStatement
         | assign    # AssignStmtStatement
         | ifStmt    # IfStmtStatement
         ;

WS: [ \t\r\n]+ -> skip ;
