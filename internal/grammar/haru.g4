grammar haru;

import Output, Expressions, Variables, ControlFlow, Arrays;


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
