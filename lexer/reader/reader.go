package reader

import (
	"bufio"
)

type LexerReader struct {
	reader   bufio.Reader
	ungetFlg bool
	char     rune
}

func New(r bufio.Reader) LexerReader {
	return LexerReader{
		reader:   r,
		ungetFlg: false,
	}
}

func (lr *LexerReader) Read() rune {
	if lr.ungetFlg {
		lr.ungetFlg = false

		return lr.char
	}

	lr.char, _, _ = lr.reader.ReadRune()

	return lr.char
}

func (lr *LexerReader) Unread() {
	lr.ungetFlg = true
}
