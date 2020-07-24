use super::token;

use std::ops::Range;
use token::Kind;

pub struct Lexer<'a> {
    input: &'a [u8],
    size: usize,
    pos: usize,
    read_pos: usize,
    ch: u8,
    loc: token::Loc,
}

impl Lexer<'_> {
    pub fn new<'a>(input: &'a [u8]) -> Lexer<'a> {
        let size = input.len();
        let mut l = Lexer {
            input: input,
            size: size,
            pos: 0,
            read_pos: 0,
            ch: 0,
            loc: token::Loc {
                row: 0,
                col: 0,
                pos: 0,
            },
        };

        l.read_char();
        l
    }

    pub fn next_token(self: &mut Self) -> Box<token::Token> {
        let kind: Kind;
        let literal: String;

        self.skip_while_spaces();

        match self.ch {
            0 => {
                kind = Kind::Eof;
                literal = "".to_string();
            }
            b'(' => {
                kind = Kind::Lparen;
                literal = "(".to_string();
            }
            b')' => {
                kind = Kind::Rparen;
                literal = ")".to_string();
            }
            _ if Self::is_identifider_char(self.ch) => {
                kind = Kind::Identifier;
                literal = self.read_identifider();
            }
            _ => {
                kind = Kind::Illegal;
                literal = "".to_string();
            }
        }

        let tok = Box::new(token::Token {
            kind: kind,
            loc: self.loc.clone(),
            literal: literal,
        });

        // Advances to next char and then returns.
        self.read_char();
        return tok;
    }
}

impl Lexer<'_> {
    pub fn bytes<'a>(self: &'a Self, range: Range<usize>) -> Option<&'a [u8]> {
        if range.end > self.size || range.start >= range.end {
            return None;
        }
        Some(&self.input[range])
    }

    pub fn substring(self: &Self, range: Range<usize>) -> String {
        match self.bytes(range) {
            Some(s) => std::str::from_utf8(s).unwrap().to_string(),
            None => "".to_string(),
        }
    }
}

impl Lexer<'_> {
    fn read_char(self: &mut Self) {
        if self.read_pos >= self.size {
            self.ch = 0;
        } else {
            self.ch = self.input[self.read_pos];
        }

        self.pos = self.read_pos;
        self.read_pos += 1;
    }

    fn skip_while_spaces(self: &mut Self) {
        loop {
            match self.ch {
                b' ' | b'\n' | b'\t' => {
                    self.read_char();
                    continue;
                }
                _ => {
                    return;
                }
            }
        }
    }

    fn is_identifider_char(c: u8) -> bool {
        match c {
            b'a'..=b'z' => true,
            b'A'..=b'Z' => true,
            b'_' | b'+' => true,
            _ => false,
        }
    }

    fn read_identifider(self: &mut Self) -> String {
        let start = self.pos;
        while Self::is_identifider_char(self.ch) {
            self.read_char()
        }
        self.substring(start..self.pos)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_lexer_read_char() {
        let mut l = Lexer::new(b"ab");
        assert_eq!(b'a', l.ch);
        assert_eq!(1, l.read_pos);
        assert_eq!(0, l.pos);
        l.read_char();

        assert_eq!(b'b', l.ch);
        assert_eq!(2, l.read_pos);
        assert_eq!(1, l.pos);
    }

    #[test]
    fn test_lexer_bytes() {
        let l = Lexer::new(b"ab");
        assert_eq!(b"a", l.bytes(0..1).unwrap());
        assert_eq!(b"ab", l.bytes(0..2).unwrap());
        assert_eq!(true, l.bytes(0..3).is_none());
        assert_eq!(true, l.bytes(0..0).is_none());
    }

    #[test]
    fn test_lexer_next_tokens() {
        let mut l = Lexer::new(b"( ) abc_+");
        let tok1 = l.next_token();
        assert_eq!("(", tok1.literal);
        assert_eq!(Kind::Lparen, tok1.kind);

        let tok2 = l.next_token();
        assert_eq!(")", tok2.literal);
        assert_eq!(Kind::Rparen, tok2.kind);

        let tok3 = l.next_token();
        assert_eq!("abc_+", tok3.literal);
        assert_eq!(Kind::Identifier, tok3.kind);

        let tok4 = l.next_token();
        assert_eq!("", tok4.literal);
        assert_eq!(Kind::Eof, tok4.kind);
    }
}
