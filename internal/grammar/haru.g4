grammar haru;

import Output, Expressions, Variables, ControlFlow, Arrays, Builtins, Functions, Loops;


program: statement* ;


statement: printStmt            # PrintStmtStatement
         | varDecl              # VarDeclStatement
         | assign               # AssignStmtStatement
         | pointerAssign        # PointerAssignStmtStatement
         | ifStmt               # IfStmtStatement
         | arrayDecl            # ArrayDeclStatement
         | arrayItemAssign      # ArrayIndexAssignStatement
         | arrayReassign        # ArrayReassignStatement
         | functionDecl         # FunctionDeclStatement
         | returnStmt           # ReturnStmtStatement
         | functionCall ';'     # FunctionCallStatement
         | whileLoop            # WhileLoopStatement
         ;


WS: [ \t\r\n]+ -> skip ;

