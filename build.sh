#!/bin/bash

cd antlr
antlr4 -o ./Go -listener -visitor -Dlanguage=Go -lib ./ *.g4
cd ..

#antlr4 -o ./Java -listener -visitor -Dlanguage=Java -lib ./ ./PegiParser.g4 ./PegiLexer.g4
