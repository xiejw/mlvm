use crate::base::Error;

use self::super::ast::Expr;
use self::super::lexer::Lexer;
use self::super::token::Token;

pub struct Parser<'a> {
    lexer: Lexer<'a>,
    curToken: Box<Token>,
    peekToken: Box<Token>,
}

impl Parser<'_> {
    pub fn new<'a>(input: &'a [u8]) -> Parser<'a> {
        let mut lexer = Lexer::new(input);
        let curToken = lexer.next_token();
        let peekToken = lexer.next_token();
        Parser {
            lexer: lexer,
            curToken: curToken,
            peekToken: peekToken,
        }
    }
}

impl Parser<'_> {
    // Consumes all valid tokens in the lexer and parses the source program into Exprs.
    pub fn parse_ast(&mut self) -> Result<Vec<Expr>, Error> {
        unimplemented!()
    }
}
