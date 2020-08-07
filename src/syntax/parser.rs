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
            TokenKind::Int => self.parse_intlt(),
            TokenKind::String => self.parse_stringlt(),
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
        debug_assert!(self.cur_token.kind == TokenKind::Int);

        let literal = &self.cur_token.literal;
        match literal.parse::<i64>() {
            Ok(v) => {
                self.advance_token();
                Ok(Expr::new_intlt(v))
            }
            Err(num_err) => {
                return Err(Error::from(num_err)
                    .emit_diagnosis_note(format!("Int literal token cannot be parsed: {}", literal))
                    .take());
            }
        }
    }
    fn parse_stringlt(&mut self) -> Result<Expr, Error> {
        debug_assert!(self.cur_token.kind == TokenKind::String);
        let tok = self.advance_token();

        // Drops the begining and tailing `"`.
        let mut bytes = tok.literal.into_bytes();
        debug_assert!(bytes.len() >= 2);
        bytes.remove(bytes.len() - 1);
        bytes.remove(0);

        Ok(Expr::new_stringlt(String::from_utf8(bytes).unwrap()))
    }
    fn parse_arraylt(&mut self) -> Result<Expr, Error> {
        unimplemented!();
    }
}

impl Parser<'_> {
    fn advance_token(&mut self) -> Box<Token> {
        let old_token = std::mem::replace(
            &mut self.cur_token,
            std::mem::replace(&mut self.peek_token, self.lexer.next_token()),
        );
        old_token
    }

    // fn check_current_token_kind(&self, kind: TokenKind) -> Result<(), Error> {
    //     match &self.cur_token.kind {
    //         v if *v == kind => Ok(()),
    //         v => Err(Error::new()
    //             .emit_diagnosis_note(format!(
    //                 "expect current token kind as {:?}, got {:?}",
    //                 kind, v
    //             ))
    //             .take()),
    //     }
    // }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_intlt() {
        let mut p = Parser::new(b"1234");
        let exprs = p.parse_ast().unwrap();
        assert_eq!(1, exprs.len());
        assert_eq!("IntLt::Int (1234)", exprs[0].to_string());
    }

    #[test]
    fn test_strlt() {
        let mut p = Parser::new(b"\"1234\"");
        let exprs = p.parse_ast().unwrap();
        assert_eq!(1, exprs.len());
        assert_eq!("StringLt(\"1234\")", exprs[0].to_string());
    }
}
