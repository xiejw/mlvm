use std::fmt;

pub struct Program {}

pub enum Expr {
    ID(String),
    IntLt(i64),
    FloatLt(f32),
    ShapeLt(Vec<Box<Expr>>),
    ArrayLt(Vec<Box<Expr>>),
    StringLt(String),
    FnCall(Box<Expr>, Vec<Box<Expr>>),
}

impl Expr {
    pub fn new_id(v: &str) -> Expr {
        Expr::ID(v.to_string())
    }

    pub fn new_shape(dims: &[&str]) -> Expr {
        Expr::ShapeLt(dims.iter().map(|x| Box::new(Expr::new_id(x))).collect())
    }
}

impl fmt::Display for Expr {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match &self {
            Expr::ID(v) => write!(f, "ID({})", v),
            Expr::IntLt(v) => write!(f, "Int({})", v),
            Expr::FloatLt(v) => write!(f, "Float({:.2})", v),
            Expr::StringLt(v) => write!(f, "Str(\"{}\")", v),
            Expr::ShapeLt(l) => {
                let _ = write!(f, "Shape(");
                Expr::write_list(f, &l);
                write!(f, ")")
            }
            _ => panic!("unsupported yet"),
        }
    }
}

impl Expr {
    fn write_list(f: &mut fmt::Formatter<'_>, list: &Vec<Box<Expr>>) {
        let len = list.len();
        for (i, e) in list.iter().enumerate() {
            let _ = write!(f, "{}", e);
            if i != len - 1 {
                let _ = write!(f, ", ");
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_id() {
        let expr = Expr::new_id("abc");
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

    #[test]
    fn test_shapelt() {
        let expr = Expr::new_shape(&vec!["@a"]);
        assert_eq!(r#"Shape(ID(@a))"#, expr.to_string());
    }
}
