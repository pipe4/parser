package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	pipe4parser "github.com/pipe4/parser/antlr/Go"
	"io/ioutil"
	"log"
)

func ParseFileSource(fileSource string) (File, error) {
	lexer := pipe4parser.NewPipe4Lexer(antlr.NewInputStream(fileSource))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := pipe4parser.NewPipe4Parser(stream)
	listener := &pipe4listener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.SourceFile())

	return listener.file, nil
}

func ParseFile(path string) (File, error) {
	log.Println("parse", path)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	return ParseFileSource(string(file))

}
