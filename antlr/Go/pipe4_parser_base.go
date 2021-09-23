package parser

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Pipe4ParserBase implementation.
type Pipe4ParserBase struct {
	*antlr.BaseParser
}

// Returns true if on the current index of the parser's
// token stream a token exists on the Hidden channel which
// either is a line terminator, or is a multi line comment that
// contains a line terminator.
func (p *Pipe4ParserBase) lineTerminatorAhead() bool {
	// Get the token ahead of the current index.
	offset := 1
	possibleIndexEosToken := p.GetCurrentToken().GetTokenIndex() - offset

	if possibleIndexEosToken == -1 {
		return true
	}

	ahead := p.GetTokenStream().Get(possibleIndexEosToken)

	for ahead.GetChannel() == antlr.LexerHidden {
		if ahead.GetTokenType() == Pipe4LexerTERMINATOR {
			return true
		}
		if ahead.GetTokenType() == Pipe4LexerWS {
			offset++
			possibleIndexEosToken = p.GetCurrentToken().GetTokenIndex() - offset
			ahead = p.GetTokenStream().Get(possibleIndexEosToken)
		}
		if ahead.GetTokenType() == Pipe4LexerCOMMENT || ahead.GetTokenType() == Pipe4LexerLINE_COMMENT {
			if strings.Contains(ahead.GetText(), "\r") || strings.Contains(ahead.GetText(), "\n") {
				return true
			} else {
				offset++
				possibleIndexEosToken = p.GetCurrentToken().GetTokenIndex() - offset
				ahead = p.GetTokenStream().Get(possibleIndexEosToken)
			}
		}
	}

	return false
}

func (p *Pipe4ParserBase) noTerminatorBetween(tokenOffset int) bool {
	var s antlr.TokenStream
	s = p.GetTokenStream()
	stream := s.(*antlr.CommonTokenStream)
	tokens := stream.GetHiddenTokensToLeft(stream.LT(tokenOffset).GetTokenIndex(), -1)
	if tokens == nil {
		return true
	}

	for _, token := range tokens {
		if strings.Contains(token.GetText(), "\n") {
			return false
		}
	}
	return true
}

func (p *Pipe4ParserBase) noTerminatorAfterParams(tokenOffset int) bool {
	stream := p.GetTokenStream()
	leftParams := 1
	rightParams := 0
	var tokenType int

	if stream.LT(tokenOffset).GetTokenType() == Pipe4LexerL_PAREN {
		for leftParams != rightParams {
			tokenOffset++
			tokenType = stream.LT(tokenOffset).GetTokenType()
			if tokenType == Pipe4LexerL_PAREN {
				leftParams++
			} else if tokenType == Pipe4LexerR_PAREN {
				rightParams++
			}
		}
		tokenOffset++
		return p.noTerminatorBetween(tokenOffset)
	}
	return true
}

func (p *Pipe4ParserBase) checkPreviousTokenText(text string) bool {
	stream := p.GetTokenStream()
	return stream.LT(1).GetText() == text
}