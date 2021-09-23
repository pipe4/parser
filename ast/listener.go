package ast

import (
	parser "github.com/pipe4/parser/antlr/Go"
)

type pipe4listener struct {
	*parser.BasePipe4ParserListener

	file File
}
