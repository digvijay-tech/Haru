grammar haru;

import Output, Expressions, Variables, ControlFlow, Arrays, Builtins;


program: statement* ;


statement: printStmt            # PrintStmtStatement
         | varDecl              # VarDeclStatement
         | assign               # AssignStmtStatement
         | ifStmt               # IfStmtStatement
         | arrayDecl            # ArrayDeclStatement
         | arrayItemAssign      # ArrayIndexAssignStatement
         | arrayReassign        # ArrayReassignStatement
         ;


WS: [ \t\r\n]+ -> skip ;
