// Code generated from Pipe4Parser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Pipe4Parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by Pipe4Parser.
type Pipe4ParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Pipe4Parser#sourceFile.
	VisitSourceFile(ctx *SourceFileContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#parserClause.
	VisitParserClause(ctx *ParserClauseContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#importDecl.
	VisitImportDecl(ctx *ImportDeclContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#importSpec.
	VisitImportSpec(ctx *ImportSpecContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#importPath.
	VisitImportPath(ctx *ImportPathContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#string_.
	VisitString_(ctx *String_Context) interface{}

	// Visit a parse tree produced by Pipe4Parser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#shortVarDecl.
	VisitShortVarDecl(ctx *ShortVarDeclContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#pipe.
	VisitPipe(ctx *PipeContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#pipeUnit.
	VisitPipeUnit(ctx *PipeUnitContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#pipeUnitIdentifers.
	VisitPipeUnitIdentifers(ctx *PipeUnitIdentifersContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#pipeUnitExpression.
	VisitPipeUnitExpression(ctx *PipeUnitExpressionContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#primaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#operand.
	VisitOperand(ctx *OperandContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#basicLit.
	VisitBasicLit(ctx *BasicLitContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#operandName.
	VisitOperandName(ctx *OperandNameContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#unaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#sliceValue.
	VisitSliceValue(ctx *SliceValueContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#sliceValueSingleLine.
	VisitSliceValueSingleLine(ctx *SliceValueSingleLineContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#sliceValueMultiLine.
	VisitSliceValueMultiLine(ctx *SliceValueMultiLineContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#sliceValueMultiLineLine.
	VisitSliceValueMultiLineLine(ctx *SliceValueMultiLineLineContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#messageValue.
	VisitMessageValue(ctx *MessageValueContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#keyedElementKV.
	VisitKeyedElementKV(ctx *KeyedElementKVContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#keyedElementIdentifer.
	VisitKeyedElementIdentifer(ctx *KeyedElementIdentiferContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#keyedElement.
	VisitKeyedElement(ctx *KeyedElementContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#messageValueSingleLine.
	VisitMessageValueSingleLine(ctx *MessageValueSingleLineContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#messageValueMultiLine.
	VisitMessageValueMultiLine(ctx *MessageValueMultiLineContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#lineKeyedElement.
	VisitLineKeyedElement(ctx *LineKeyedElementContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#key.
	VisitKey(ctx *KeyContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#element.
	VisitElement(ctx *ElementContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#sliceType.
	VisitSliceType(ctx *SliceTypeContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#typeT.
	VisitTypeT(ctx *TypeTContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#typeLit.
	VisitTypeLit(ctx *TypeLitContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#interfaceType.
	VisitInterfaceType(ctx *InterfaceTypeContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#fieldDecl.
	VisitFieldDecl(ctx *FieldDeclContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#integer.
	VisitInteger(ctx *IntegerContext) interface{}

	// Visit a parse tree produced by Pipe4Parser#eos.
	VisitEos(ctx *EosContext) interface{}
}
