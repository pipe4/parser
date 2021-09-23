// Code generated from Pipe4Parser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Pipe4Parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasePipe4ParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePipe4ParserVisitor) VisitSourceFile(ctx *SourceFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitParserClause(ctx *ParserClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitImportDecl(ctx *ImportDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitImportSpec(ctx *ImportSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitImportPath(ctx *ImportPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitString_(ctx *String_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitStatementList(ctx *StatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitShortVarDecl(ctx *ShortVarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitPipe(ctx *PipeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitPipeUnit(ctx *PipeUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitPipeUnitIdentifers(ctx *PipeUnitIdentifersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitPipeUnitExpression(ctx *PipeUnitExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitOperand(ctx *OperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitBasicLit(ctx *BasicLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitOperandName(ctx *OperandNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitSliceValue(ctx *SliceValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitSliceValueSingleLine(ctx *SliceValueSingleLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitSliceValueMultiLine(ctx *SliceValueMultiLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitSliceValueMultiLineLine(ctx *SliceValueMultiLineLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitMessageValue(ctx *MessageValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitKeyedElementKV(ctx *KeyedElementKVContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitKeyedElementIdentifer(ctx *KeyedElementIdentiferContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitKeyedElement(ctx *KeyedElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitMessageValueSingleLine(ctx *MessageValueSingleLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitMessageValueMultiLine(ctx *MessageValueMultiLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitLineKeyedElement(ctx *LineKeyedElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitKey(ctx *KeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitElement(ctx *ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitSliceType(ctx *SliceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitTypeT(ctx *TypeTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitTypeLit(ctx *TypeLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitInterfaceType(ctx *InterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitFieldDecl(ctx *FieldDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitInteger(ctx *IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePipe4ParserVisitor) VisitEos(ctx *EosContext) interface{} {
	return v.VisitChildren(ctx)
}
