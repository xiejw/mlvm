package parser

type Parser struct {
	input []byte
}

func New(input []byte) *Parser {
	return &Parser{
		input: input,
	}
}
