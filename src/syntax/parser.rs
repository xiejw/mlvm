use self::super::lexer::Lexer;
use self::super::token::Token;
// use self::super::expr::*;
use crate::base::Error;

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
//     pub fn parse_ast(&mut self) -> Result<Vec<  {
}
