package main

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	TokenIdentifier TokenType = "IDENTIFIER"
	TokenNumber     TokenType = "NUMBER"
	TokenOperator   TokenType = "OPERATOR"
	TokenEOF        TokenType = "EOF"
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
}

type Lexer struct {
	source string
	start  int
	current int
}

func NewLexer(source string) *Lexer {
	return &Lexer{source: source}
}

func (l *Lexer) Lex() []Token {
	var tokens []Token

	for !l.isAtEnd() {
		l.start = l.current
		token := l.scanToken()
		if token.Type != "" {
			tokens = append(tokens, token)
		}
	}
	tokens = append(tokens, Token{Type: TokenEOF, Lexeme: "", Literal: nil})
	return tokens
}

func (l *Lexer) scanToken() Token {
	c := l.advance()

	if unicode.IsLetter(c) {
		return l.identifier()
	}
	if unicode.IsDigit(c) {
		return l.number()
	}
	switch c {
	case '+', '-', '*', '/':
		return Token{Type: TokenOperator, Lexeme: string(c), Literal: nil}
	}
	return Token{}
}

func (l *Lexer) identifier() Token {
	for unicode.IsLetter(l.peek()) {
		l.advance()
	}
	text := l.source[l.start:l.current]
	return Token{Type: TokenIdentifier, Lexeme: text, Literal: nil}
}

func (l *Lexer) number() Token {
	for unicode.IsDigit(l.peek()) {
		l.advance()
	}
	text := l.source[l.start:l.current]
	return Token{Type: TokenNumber, Lexeme: text, Literal: text}
}

func (l *Lexer) advance() rune {
	r := rune(l.source[l.current])
	l.current++
	return r
}

func (l *Lexer) isAtEnd() bool {
	return l.current >= len(l.source)
}

func (l *Lexer) peek() rune {
	if l.isAtEnd() {
		return '\000'
	}
	return rune(l.source[l.current])
}

type Parser struct {
	tokens  []Token
	current int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) parse() interface{} {
	return p.expression()
}

func (p *Parser) expression() interface{} {
	return p.term()
}

func (p *Parser) term() interface{} {
	left := p.factor()

	for p.match(TokenOperator) {
		operator := p.previous()
		right := p.factor()
		left = fmt.Sprintf("(%v %s %v)", left, operator.Lexeme, right)
	}
	return left
}

func (p *Parser) factor() interface{} {
	token := p.advance()
	if token.Type == TokenNumber {
		return token.Lexeme
	}
	return nil
}

func (p *Parser) match(tokenType TokenType) bool {
	if p.check(tokenType) {
		p.advance()
		return true
	}
	return false
}

func (p *Parser) check(tokenType TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tokenType
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == TokenEOF
}

func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}

func main() {
	source := "3 + 5 * 2"
	lexer := NewLexer(source)
	tokens := lexer.Lex()

	fmt.Println("Tokenlər:")
	for _, token := range tokens {
		fmt.Println(token)
	}

	parser := NewParser(tokens)
	result := parser.parse()
	fmt.Printf("Nəticə: %v\n", result)
}
