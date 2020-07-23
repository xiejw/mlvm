pub mod token {

    #[derive(Debug)]
    pub enum Type {
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

    struct Loc {
        Row: usize,
        Col: usize,
        Pos: usize,
    }

    struct Token {}
}
