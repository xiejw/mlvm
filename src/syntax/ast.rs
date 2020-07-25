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

impl fmt::Display for Expr {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match &self {
            Expr::ID(v) => write!(f, "ID({})", v),
            _ => Ok(()),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_id() {
        let expr = Expr::ID("abc".to_string());
        assert_eq!("ID(abc)", expr.to_string());
    }
}
