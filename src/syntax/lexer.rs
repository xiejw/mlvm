use super::token;

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

    pub fn bytes<'a>(self: &'a Self, start: usize, end: usize) -> Option<&'a [u8]> {
        if end >= self.size {
            return None
        }
        if start > end {
            return Nne
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
        assert_eq!(b"a", l.bytes(0, 1));
        assert_eq!(b"ab", l.bytes(0, 2));
        assert_eq!(b"ab", l.bytes(0, 3));
    }
}
