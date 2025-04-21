grammar haru;

import Output, Expressions, Variables, ControlFlow, Arrays, Builtins, Functions;


program: statement* ;


statement: printStmt            # PrintStmtStatement
         | varDecl              # VarDeclStatement
         | assign               # AssignStmtStatement
         | ifStmt               # IfStmtStatement
         | arrayDecl            # ArrayDeclStatement
         | arrayItemAssign      # ArrayIndexAssignStatement
         | arrayReassign        # ArrayReassignStatement
         | functionDecl         # FunctionDeclStatement
         | returnStmt           # ReturnStmtStatement
         | functionCall ';'     # FunctionCallStatement
         ;


WS: [ \t\r\n]+ -> skip ;

