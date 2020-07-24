use super::token;

pub struct Lexer {
    input: Vec<u8>,
    size: usize,
    pos: usize,
    read_pos: usize,
    ch: u8,
    loc: token::Loc,
}

impl Lexer {
    pub fn new(input: Vec<u8>) -> Lexer {
        let size = input.len();
        Lexer {
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
        }
    }
}

impl Lexer {
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
    fn test_lexer() {
        let mut l = Lexer::new(b"ab".to_vec());
        l.read_char();
        assert_eq!(b'a', l.ch);
    }
}
