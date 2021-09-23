package ast

import (
	parser "github.com/pipe4/parser/antlr/Go"
)

func (l *pipe4listener) ExitImportPath(ctx * parser.ImportPathContext) {
	l.file.Import = append(l.file.Import, ctx.GetText())
}

