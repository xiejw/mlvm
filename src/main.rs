mod mlvm {
    pub mod syntax {
        pub mod lexer {

            #[derive(Debug)]
            pub enum Token {
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
        }
    }
}

fn main() {
    println!("Hello, world! {:?}", mlvm::syntax::lexer::Token::Lparen);
}
