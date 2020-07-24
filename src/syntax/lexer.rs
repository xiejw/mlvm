use super::token;

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

        match self.ch {
            b'(' => {
                kind = Kind::Lparen;
                literal = "(".to_string();
                // std::str::from_utf8(self.bytes(0, 0).unwrap())
                //     .unwrap()
                //     .to_string();
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

        self.read_char();

        tok
    }
}

impl Lexer<'_> {
    pub fn bytes<'a>(self: &'a Self, start: usize, end: usize) -> Option<&'a [u8]> {
        if end > self.size || start >= end {
            return None;
        }
        Some(&self.input[start..end])
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
        assert_eq!(b"a", l.bytes(0, 1).unwrap());
        assert_eq!(b"ab", l.bytes(0, 2).unwrap());
        assert_eq!(true, l.bytes(0, 3).is_none());
        assert_eq!(true, l.bytes(0, 0).is_none());
    }

    #[test]
    fn test_lexer_next_tokens() {
        let mut l = Lexer::new(b"(");
        let tok1 = l.next_token();
        assert_eq!("(", tok1.literal);
        assert_eq!(Kind::Lparen, tok1.kind);
    }
}
