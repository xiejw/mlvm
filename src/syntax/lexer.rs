use super::token::Kind;
use std::ops::Range;

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

    pub fn next_token(&mut self) -> Box<token::Token> {
        let kind: Kind;
        let literal: String;
        let mut advance_one_char = true;

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
            b'[' => {
                kind = Kind::Lsbracket;
                literal = "[".to_string();
            }
            b']' => {
                kind = Kind::Rsbracket;
                literal = "]".to_string();
            }
            b'\\' => {
                kind = Kind::Backslash;
                literal = "\\".to_string();
            }
            b'"' => {
                kind = Kind::String;
                literal = self.read_string();
            }
            _ if Self::is_identifider_char(self.ch) => {
                kind = Kind::Identifier;
                literal = self.read_identifider();
                advance_one_char = false; // Skips the next read_char
            }
            _ if Self::is_digit(self.ch) => {
                let mut num_kind = Kind::Integer;
                literal = self.read_number(&mut num_kind);
                kind = num_kind;
                advance_one_char = false; // Skips the next read_char
            }
            _ => {
                kind = Kind::Illegal;
                literal = "".to_string();
                advance_one_char = false; // Skips the next read_char
            }
        }

        let tok = Box::new(token::Token {
            kind: kind,
            loc: self.loc.clone(),
            literal: literal,
        });

        // Advances to next char, if asked, and then returns.
        if advance_one_char {
            self.read_char();
        }
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

    fn is_digit(c: u8) -> bool {
        (b'0'..=b'9').contains(&c)
    }

    fn read_identifider(self: &mut Self) -> String {
        let start = self.pos;
        while Self::is_identifider_char(self.ch) {
            self.read_char()
        }
        self.substring(start..self.pos)
    }

    fn read_number(self: &mut Self, kind: &mut Kind) -> String {
        let mut hit_dec_pt = false;

        let start = self.pos;
        loop {
            let c = self.ch;
            if Self::is_digit(c) {
                self.read_char();
                continue;
            }

            // decimal pt should be hit at most once.
            if c == b'.' && !hit_dec_pt {
                hit_dec_pt = true;
                *kind = Kind::Float;
                self.read_char();
                continue;
            }

            break;
        }

        return self.substring(start..self.pos);
    }

    fn read_string(self: &mut Self) -> String {
        debug_assert!(self.ch == b'"');
        let start = self.pos;
        self.read_char();

        // TODO: handle EOF and newline.
        while self.ch != b'"' {
            self.read_char();
        }

        return self.substring(start..self.pos + 1);
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
        let mut l = Lexer::new(b"( ) \"efd\"abc_+z 1 \\[2.3 4.]");
        {
            let tok = l.next_token();
            assert_eq!("(", tok.literal);
            assert_eq!(Kind::Lparen, tok.kind);
        }
        {
            let tok = l.next_token();
            assert_eq!(")", tok.literal);
            assert_eq!(Kind::Rparen, tok.kind);
        }
        {
            let tok = l.next_token();
            assert_eq!(r#""efd""#, tok.literal);
            assert_eq!(Kind::String, tok.kind);
        }
        {
            let tok = l.next_token();
            assert_eq!("abc_+z", tok.literal);
            assert_eq!(Kind::Identifier, tok.kind);
        }
        {
            let tok = l.next_token();
            assert_eq!("1", tok.literal);
            assert_eq!(Kind::Integer, tok.kind);
        }
        {
            let tok = l.next_token();
            assert_eq!("\\", tok.literal);
            assert_eq!(Kind::Backslash, tok.kind);
        }
        {
            let tok = l.next_token();
            assert_eq!("[", tok.literal);
            assert_eq!(Kind::Lsbracket, tok.kind);
        }
        {
            let tok = l.next_token();
            assert_eq!("2.3", tok.literal);
            assert_eq!(Kind::Float, tok.kind);
        }
        {
            let tok = l.next_token();
            assert_eq!("4.", tok.literal);
            assert_eq!(Kind::Float, tok.kind);
        }
        {
            let tok = l.next_token();
            assert_eq!("]", tok.literal);
            assert_eq!(Kind::Rsbracket, tok.kind);
        }
        {
            let tok = l.next_token();
            assert_eq!("", tok.literal);
            assert_eq!(Kind::Eof, tok.kind);
        }
    }
}
