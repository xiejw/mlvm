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
            Expr::IntLt(v) => write!(f, "Int({})", v),
            Expr::FloatLt(v) => write!(f, "Float({:.2})", v),
            Expr::StringLt(v) => write!(f, "Str(\"{}\")", v),
            _ => panic!("unsupported yet"),
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

    #[test]
    fn test_intlt() {
        let expr = Expr::IntLt(123);
        assert_eq!("Int(123)", expr.to_string());
    }

    #[test]
    fn test_floatlt() {
        let expr = Expr::FloatLt(123.0);
        assert_eq!("Float(123.00)", expr.to_string());
    }

    #[test]
    fn test_stringlt() {
        let expr = Expr::StringLt("abc".to_string());
        assert_eq!(r#"Str("abc")"#, expr.to_string());
    }
}
