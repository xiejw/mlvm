pub mod token {
    #[derive(Debug)]
    pub enum Kind {
        Lparen,
        Rparen,
        Lsbracket,
        Rsbracket,
        Backslash,
        Identifier,
        Integer,
        Float,
        String,
        Illegal,
        Eof,
    }

    pub struct Loc {
        pub row: usize,
        pub col: usize,
        pub pos: usize,
    }

    pub struct Token {
        pub kind: Kind,
        pub loc: Loc,
        pub literal: String,
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_token() {
        let tok = token::Token {
            kind: token::Kind::Lparen,
            loc: token::Loc {
                row: 1,
                col: 1,
                pos: 2,
            },
            literal: String::from("("),
        };
        assert_eq!(&tok.literal, "(");
    }
}
