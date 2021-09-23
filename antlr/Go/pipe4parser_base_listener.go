// Code generated from Pipe4Parser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Pipe4Parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePipe4ParserListener is a complete listener for a parse tree produced by Pipe4Parser.
type BasePipe4ParserListener struct{}

var _ Pipe4ParserListener = &BasePipe4ParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePipe4ParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePipe4ParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePipe4ParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePipe4ParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSourceFile is called when production sourceFile is entered.
func (s *BasePipe4ParserListener) EnterSourceFile(ctx *SourceFileContext) {}

// ExitSourceFile is called when production sourceFile is exited.
func (s *BasePipe4ParserListener) ExitSourceFile(ctx *SourceFileContext) {}

// EnterParserClause is called when production parserClause is entered.
func (s *BasePipe4ParserListener) EnterParserClause(ctx *ParserClauseContext) {}

// ExitParserClause is called when production parserClause is exited.
func (s *BasePipe4ParserListener) ExitParserClause(ctx *ParserClauseContext) {}

// EnterImportDecl is called when production importDecl is entered.
func (s *BasePipe4ParserListener) EnterImportDecl(ctx *ImportDeclContext) {}

// ExitImportDecl is called when production importDecl is exited.
func (s *BasePipe4ParserListener) ExitImportDecl(ctx *ImportDeclContext) {}

// EnterImportSpec is called when production importSpec is entered.
func (s *BasePipe4ParserListener) EnterImportSpec(ctx *ImportSpecContext) {}

// ExitImportSpec is called when production importSpec is exited.
func (s *BasePipe4ParserListener) ExitImportSpec(ctx *ImportSpecContext) {}

// EnterImportPath is called when production importPath is entered.
func (s *BasePipe4ParserListener) EnterImportPath(ctx *ImportPathContext) {}

// ExitImportPath is called when production importPath is exited.
func (s *BasePipe4ParserListener) ExitImportPath(ctx *ImportPathContext) {}

// EnterString_ is called when production string_ is entered.
func (s *BasePipe4ParserListener) EnterString_(ctx *String_Context) {}

// ExitString_ is called when production string_ is exited.
func (s *BasePipe4ParserListener) ExitString_(ctx *String_Context) {}

// EnterStatementList is called when production statementList is entered.
func (s *BasePipe4ParserListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BasePipe4ParserListener) ExitStatementList(ctx *StatementListContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasePipe4ParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasePipe4ParserListener) ExitStatement(ctx *StatementContext) {}

// EnterShortVarDecl is called when production shortVarDecl is entered.
func (s *BasePipe4ParserListener) EnterShortVarDecl(ctx *ShortVarDeclContext) {}

// ExitShortVarDecl is called when production shortVarDecl is exited.
func (s *BasePipe4ParserListener) ExitShortVarDecl(ctx *ShortVarDeclContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BasePipe4ParserListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BasePipe4ParserListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BasePipe4ParserListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BasePipe4ParserListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterPipe is called when production pipe is entered.
func (s *BasePipe4ParserListener) EnterPipe(ctx *PipeContext) {}

// ExitPipe is called when production pipe is exited.
func (s *BasePipe4ParserListener) ExitPipe(ctx *PipeContext) {}

// EnterPipeUnit is called when production pipeUnit is entered.
func (s *BasePipe4ParserListener) EnterPipeUnit(ctx *PipeUnitContext) {}

// ExitPipeUnit is called when production pipeUnit is exited.
func (s *BasePipe4ParserListener) ExitPipeUnit(ctx *PipeUnitContext) {}

// EnterPipeUnitIdentifers is called when production pipeUnitIdentifers is entered.
func (s *BasePipe4ParserListener) EnterPipeUnitIdentifers(ctx *PipeUnitIdentifersContext) {}

// ExitPipeUnitIdentifers is called when production pipeUnitIdentifers is exited.
func (s *BasePipe4ParserListener) ExitPipeUnitIdentifers(ctx *PipeUnitIdentifersContext) {}

// EnterPipeUnitExpression is called when production pipeUnitExpression is entered.
func (s *BasePipe4ParserListener) EnterPipeUnitExpression(ctx *PipeUnitExpressionContext) {}

// ExitPipeUnitExpression is called when production pipeUnitExpression is exited.
func (s *BasePipe4ParserListener) ExitPipeUnitExpression(ctx *PipeUnitExpressionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePipe4ParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePipe4ParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *BasePipe4ParserListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *BasePipe4ParserListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterOperand is called when production operand is entered.
func (s *BasePipe4ParserListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BasePipe4ParserListener) ExitOperand(ctx *OperandContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasePipe4ParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasePipe4ParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BasePipe4ParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BasePipe4ParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BasePipe4ParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BasePipe4ParserListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterBasicLit is called when production basicLit is entered.
func (s *BasePipe4ParserListener) EnterBasicLit(ctx *BasicLitContext) {}

// ExitBasicLit is called when production basicLit is exited.
func (s *BasePipe4ParserListener) ExitBasicLit(ctx *BasicLitContext) {}

// EnterOperandName is called when production operandName is entered.
func (s *BasePipe4ParserListener) EnterOperandName(ctx *OperandNameContext) {}

// ExitOperandName is called when production operandName is exited.
func (s *BasePipe4ParserListener) ExitOperandName(ctx *OperandNameContext) {}

// EnterUnaryExpr is called when production unaryExpr is entered.
func (s *BasePipe4ParserListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production unaryExpr is exited.
func (s *BasePipe4ParserListener) ExitUnaryExpr(ctx *UnaryExprContext) {}

// EnterSliceValue is called when production sliceValue is entered.
func (s *BasePipe4ParserListener) EnterSliceValue(ctx *SliceValueContext) {}

// ExitSliceValue is called when production sliceValue is exited.
func (s *BasePipe4ParserListener) ExitSliceValue(ctx *SliceValueContext) {}

// EnterSliceValueSingleLine is called when production sliceValueSingleLine is entered.
func (s *BasePipe4ParserListener) EnterSliceValueSingleLine(ctx *SliceValueSingleLineContext) {}

// ExitSliceValueSingleLine is called when production sliceValueSingleLine is exited.
func (s *BasePipe4ParserListener) ExitSliceValueSingleLine(ctx *SliceValueSingleLineContext) {}

// EnterSliceValueMultiLine is called when production sliceValueMultiLine is entered.
func (s *BasePipe4ParserListener) EnterSliceValueMultiLine(ctx *SliceValueMultiLineContext) {}

// ExitSliceValueMultiLine is called when production sliceValueMultiLine is exited.
func (s *BasePipe4ParserListener) ExitSliceValueMultiLine(ctx *SliceValueMultiLineContext) {}

// EnterSliceValueMultiLineLine is called when production sliceValueMultiLineLine is entered.
func (s *BasePipe4ParserListener) EnterSliceValueMultiLineLine(ctx *SliceValueMultiLineLineContext) {}

// ExitSliceValueMultiLineLine is called when production sliceValueMultiLineLine is exited.
func (s *BasePipe4ParserListener) ExitSliceValueMultiLineLine(ctx *SliceValueMultiLineLineContext) {}

// EnterMessageValue is called when production messageValue is entered.
func (s *BasePipe4ParserListener) EnterMessageValue(ctx *MessageValueContext) {}

// ExitMessageValue is called when production messageValue is exited.
func (s *BasePipe4ParserListener) ExitMessageValue(ctx *MessageValueContext) {}

// EnterKeyedElementKV is called when production keyedElementKV is entered.
func (s *BasePipe4ParserListener) EnterKeyedElementKV(ctx *KeyedElementKVContext) {}

// ExitKeyedElementKV is called when production keyedElementKV is exited.
func (s *BasePipe4ParserListener) ExitKeyedElementKV(ctx *KeyedElementKVContext) {}

// EnterKeyedElementIdentifer is called when production keyedElementIdentifer is entered.
func (s *BasePipe4ParserListener) EnterKeyedElementIdentifer(ctx *KeyedElementIdentiferContext) {}

// ExitKeyedElementIdentifer is called when production keyedElementIdentifer is exited.
func (s *BasePipe4ParserListener) ExitKeyedElementIdentifer(ctx *KeyedElementIdentiferContext) {}

// EnterKeyedElement is called when production keyedElement is entered.
func (s *BasePipe4ParserListener) EnterKeyedElement(ctx *KeyedElementContext) {}

// ExitKeyedElement is called when production keyedElement is exited.
func (s *BasePipe4ParserListener) ExitKeyedElement(ctx *KeyedElementContext) {}

// EnterMessageValueSingleLine is called when production messageValueSingleLine is entered.
func (s *BasePipe4ParserListener) EnterMessageValueSingleLine(ctx *MessageValueSingleLineContext) {}

// ExitMessageValueSingleLine is called when production messageValueSingleLine is exited.
func (s *BasePipe4ParserListener) ExitMessageValueSingleLine(ctx *MessageValueSingleLineContext) {}

// EnterMessageValueMultiLine is called when production messageValueMultiLine is entered.
func (s *BasePipe4ParserListener) EnterMessageValueMultiLine(ctx *MessageValueMultiLineContext) {}

// ExitMessageValueMultiLine is called when production messageValueMultiLine is exited.
func (s *BasePipe4ParserListener) ExitMessageValueMultiLine(ctx *MessageValueMultiLineContext) {}

// EnterLineKeyedElement is called when production lineKeyedElement is entered.
func (s *BasePipe4ParserListener) EnterLineKeyedElement(ctx *LineKeyedElementContext) {}

// ExitLineKeyedElement is called when production lineKeyedElement is exited.
func (s *BasePipe4ParserListener) ExitLineKeyedElement(ctx *LineKeyedElementContext) {}

// EnterKey is called when production key is entered.
func (s *BasePipe4ParserListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BasePipe4ParserListener) ExitKey(ctx *KeyContext) {}

// EnterElement is called when production element is entered.
func (s *BasePipe4ParserListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BasePipe4ParserListener) ExitElement(ctx *ElementContext) {}

// EnterSliceType is called when production sliceType is entered.
func (s *BasePipe4ParserListener) EnterSliceType(ctx *SliceTypeContext) {}

// ExitSliceType is called when production sliceType is exited.
func (s *BasePipe4ParserListener) ExitSliceType(ctx *SliceTypeContext) {}

// EnterTypeT is called when production typeT is entered.
func (s *BasePipe4ParserListener) EnterTypeT(ctx *TypeTContext) {}

// ExitTypeT is called when production typeT is exited.
func (s *BasePipe4ParserListener) ExitTypeT(ctx *TypeTContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BasePipe4ParserListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BasePipe4ParserListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterTypeLit is called when production typeLit is entered.
func (s *BasePipe4ParserListener) EnterTypeLit(ctx *TypeLitContext) {}

// ExitTypeLit is called when production typeLit is exited.
func (s *BasePipe4ParserListener) ExitTypeLit(ctx *TypeLitContext) {}

// EnterInterfaceType is called when production interfaceType is entered.
func (s *BasePipe4ParserListener) EnterInterfaceType(ctx *InterfaceTypeContext) {}

// ExitInterfaceType is called when production interfaceType is exited.
func (s *BasePipe4ParserListener) ExitInterfaceType(ctx *InterfaceTypeContext) {}

// EnterFieldDecl is called when production fieldDecl is entered.
func (s *BasePipe4ParserListener) EnterFieldDecl(ctx *FieldDeclContext) {}

// ExitFieldDecl is called when production fieldDecl is exited.
func (s *BasePipe4ParserListener) ExitFieldDecl(ctx *FieldDeclContext) {}

// EnterInteger is called when production integer is entered.
func (s *BasePipe4ParserListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BasePipe4ParserListener) ExitInteger(ctx *IntegerContext) {}

// EnterEos is called when production eos is entered.
func (s *BasePipe4ParserListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BasePipe4ParserListener) ExitEos(ctx *EosContext) {}
