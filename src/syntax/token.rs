#[derive(Debug, Clone, PartialEq)]
pub enum Kind {
    Lparen,
    Rparen,
    Lsbracket,
    Rsbracket,
    Backslash,
    Identifier,
    Int,
    Float,
    String,
    Illegal,
    Eof,
}

#[derive(Clone)]
pub struct Loc {
    pub row: usize,
    pub col: usize,
    pub pos: usize,
}

#[derive(Clone)]
pub struct Token {
    pub kind: Kind,
    pub loc: Loc,
    pub literal: String,
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_token() {
        let tok = Token {
            kind: Kind::Lparen,
            loc: Loc {
                row: 1,
                col: 1,
                pos: 2,
            },
            literal: String::from("("),
        };
        assert_eq!(&tok.literal, "(");
    }
}
