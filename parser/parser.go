package parser

import (
	"github.com/Laaman03/monkey/token"
	"github.com/Laaman03/monkey/lexer"
	"github.com/Laaman03/monkey/ast"
)

type Parser struct{
	l *lexer.Lexer
	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// Read 2 tokens so both curToken and peekToken are set
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
