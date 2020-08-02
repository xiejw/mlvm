use crate::base::Error;

use self::super::ast::Expr;
use self::super::lexer::Lexer;
use self::super::token::Kind as TokenKind;
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
        let mut exprs = Vec::new();

        loop {
            if self.curToken.kind == TokenKind::Eof {
                break;
            }

            match self.parse_expr() {
                Ok(expr) => {
                    exprs.push(expr);
                }
                Err(mut err) => {
                    let i = exprs.len() + 1;
                    return Err(err
                        .emit_diagnosis_note(format!("failed to parse {}-th top level expr", i))
                        .take());
                }
            }
        }
        Ok(exprs)
    }
}

impl Parser<'_> {
    fn parse_expr(&mut self) -> Result<Expr, Error> {
        let r = match self.curToken.kind {
            TokenKind::Integer => self.parse_intlt(),
            _ => panic!("unsupported expr for parser"),
        };

        r.map_err(|mut err| {
            err.emit_diagnosis_note_str("failed to parse expression")
                .take()
        })
    }

    fn parse_id(&mut self) -> Result<Expr, Error> {
        unimplemented!();
    }
    fn parse_intlt(&mut self) -> Result<Expr, Error> {
        unimplemented!();
    }
    fn parse_strlt(&mut self) -> Result<Expr, Error> {
        unimplemented!();
    }
    fn parse_arraylt(&mut self) -> Result<Expr, Error> {
        unimplemented!();
    }
}
