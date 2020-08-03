use crate::base::Error;

use self::super::ast::Expr;
use self::super::lexer::Lexer;
use self::super::token::Kind as TokenKind;
use self::super::token::Token;

pub struct Parser<'a> {
    lexer: Lexer<'a>,
    cur_token: Box<Token>,
    peek_token: Box<Token>,
}

impl Parser<'_> {
    pub fn new<'a>(input: &'a [u8]) -> Parser<'a> {
        let mut lexer = Lexer::new(input);
        let cur_token = lexer.next_token();
        let peek_token = lexer.next_token();
        Parser {
            lexer: lexer,
            cur_token: cur_token,
            peek_token: peek_token,
        }
    }
}

impl Parser<'_> {
    // Consumes all valid tokens in the lexer and parses the source program into Exprs.
    pub fn parse_ast(&mut self) -> Result<Vec<Expr>, Error> {
        let mut exprs = Vec::new();

        loop {
            if self.cur_token.kind == TokenKind::Eof {
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
        let r = match self.cur_token.kind {
            TokenKind::Identifier => self.parse_id(),
            TokenKind::Integer => self.parse_intlt(),
            TokenKind::String => self.parse_strlt(),
            TokenKind::Lsbracket => self.parse_arraylt(),
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
        self.advance_token();
        unimplemented!();
    }
    fn parse_strlt(&mut self) -> Result<Expr, Error> {
        unimplemented!();
    }
    fn parse_arraylt(&mut self) -> Result<Expr, Error> {
        unimplemented!();
    }
}

impl Parser<'_> {
    fn advance_token(&mut self) {
        self.cur_token = std::mem::replace(&mut self.peek_token, self.lexer.next_token());
    }
}
