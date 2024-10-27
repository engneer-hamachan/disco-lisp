package lexer

import (
	"disco/base"
	"disco/lexer/reader"
	"strconv"
	"strings"
	"unicode"
)

type Lexer struct {
	tok    rune
	val    interface{}
	reader reader.LexerReader
}

var reserved map[string]interface{} = make(
	map[string]interface{},
)

func New(lr reader.LexerReader) Lexer {
	reserved["t"] = 't'
	reserved["nil"] = rune(base.NIL)
	reserved["load"] = rune(base.LOAD)

	return Lexer{
		reader: lr,
	}
}

func (l *Lexer) Token() rune {
	return l.tok
}

func (l *Lexer) Value() interface{} {
	return l.val
}

func (l *Lexer) lexDigit() {
	var buf strings.Builder

	for {
		char := l.reader.Read()

		if !unicode.IsDigit(char) && char != '.' {
			l.reader.Unread()
			break
		}

		buf.WriteRune(char)
	}

	val, err := strconv.ParseInt(buf.String(), 10, 64)

	if err != nil {
		l.val, _ = strconv.ParseFloat(buf.String(), 64)
		l.tok = base.FLOAT

		return
	}

	l.val = val
	l.tok = base.INT
}

func (l *Lexer) lexSymbol() {
	l.tok = base.SYMBOL

	var buf strings.Builder

	for {
		char := l.reader.Read()

		if !isSymbolChar(char) {
			l.reader.Unread()
			break
		}

		buf.WriteRune(char)
	}

	str := buf.String()
	l.val = base.Intern(str)

	val, ok := reserved[str]
	if ok {
		l.tok = val.(rune)
	}
}

func (l *Lexer) lexString() {
	var buf strings.Builder

	for {
		char := l.reader.Read()

		if char == '"' {
			break
		}

		if char == '\\' {
			char = l.reader.Read()
		}

		buf.WriteRune(char)
	}

	l.val = buf.String()
}

func (l *Lexer) skipSpace() {
	char := l.reader.Read()

	for {
		if !unicode.IsSpace(char) || char == '\n' {
			break
		}

		char = l.reader.Read()
	}

	l.reader.Unread()
}

func (l *Lexer) skipLineComment() {
	var char rune

	for {
		char = l.reader.Read()

		if char == '\n' {
			break
		}
	}

	l.reader.Unread()
}

func (l *Lexer) Advance() bool {
	l.skipSpace()
	char := l.reader.Read()

	switch char {
	case '\n', '(', ')', '\'', '`', ',', '#', '@':
		l.tok = char

	case '"':
		l.lexString()
		l.tok = base.STRING

	case ';':
		l.skipLineComment()
		return l.Advance()

	default:
		if unicode.IsDigit(char) {
			l.reader.Unread()
			l.lexDigit()

			break
		}

		if isSymbolChar(char) {
			l.reader.Unread()
			l.lexSymbol()

			break
		}

		return false
	}

	return true
}

func isSymbolChar(c rune) bool {
	if unicode.IsSpace(c) ||
		c == '\n' ||
		c == '(' ||
		c == ')' ||
		c == base.NIL {

		return false
	}

	return true
}
