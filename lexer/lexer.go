package lexer

type TokenType string

const (
	// SPECIAL TOKENS
	ILLEGAL    = "ILLEGAL"
	EOF        = "EOF"
	WHITESPACE = "WHITESPACE"

	// LITERALS
	INT  = "INT"
	CHAR = "CHAR"

	// OPERATORS
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"
	MODULUS  = "%"
	// Special Operators
	INT_DIVIDE = "//"
	POWER      = "**"
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input        string
	position     int
	readposition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readposition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readposition]
	}
	l.position = l.readposition
	l.readposition++
}

func (l *Lexer) NextToken() Token {
	var token Token

	switch l.ch {
	case '+':
		token = Token{Type: PLUS, Value: "+"}
	case '-':
		token = Token{Type: MINUS, Value: "-"}
	case '*':
		switch l.lengthOperator() {
		case 1:
			token = Token{Type: MULTIPLY, Value: "*"}
		case 2:
			token = Token{Type: POWER, Value: "**"}
		default:
			token = Token{Type: ILLEGAL, Value: ""}
		}
		return token
	case '/':
		switch l.lengthOperator() {
		case 1:
			token = Token{Type: DIVIDE, Value: "/"}
		case 2:
			token = Token{Type: INT_DIVIDE, Value: "//"}
		default:
			token = Token{Type: ILLEGAL, Value: ""}
		}
		return token
	case '%':
		token = Token{Type: MODULUS, Value: "%"}
	case ' ':
		token = Token{Type: WHITESPACE, Value: " "}
	case 0:
		token = Token{Type: EOF, Value: ""}
	default:
		if isDigit(l.ch) {
			token.Value = l.readNumber()
			token.Type = INT
			return token
		} else {
			token = Token{Type: ILLEGAL, Value: string(l.ch)}
		}
	}
	l.readChar()
	return token

}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) lengthOperator() int {
	ch := l.ch
	length := 0
	for l.ch == ch {
		length++
		l.readChar()
	}
	return length
}

func isDigit(ch byte) bool {
	return ('0' <= ch && ch <= '9')

}
