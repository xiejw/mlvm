use std::fmt;

pub struct Program {}

pub enum Expr{
    ID(String),
    IntLt(i64),
    FloatLt(f32),
    ShaptLt(Vec<Box<Expr>>),
    ArrayLt(Vec<Box<Expr>>),
    StringLt(String),
    FnCall(Box<Expr>, Vec<Box<Expr>>),
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_id() {
        let expr = Expr::ID("abc".to_string());
        match &expr {
            Expr::ID(v) =>  assert_eq!("abc", v),
            _ => panic!("abc"),
        }
    }
}
