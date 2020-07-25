use std::fmt;

pub struct Program {}

pub enum Type {
    Unknown,
    Int,
    Float,
    Dim,
    Shape,
    Array,
    String,
    Fn(Vec<Type>, Vec<Type>), // inputs, outputs
}

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

    pub fn new_array(values: &[f32]) -> Expr {
        Expr::ArrayLt(values.iter().map(|x| Box::new(Expr::FloatLt(*x))).collect())
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
            Expr::ArrayLt(l) => {
                let _ = write!(f, "Array(");
                Expr::write_list(f, &l);
                write!(f, ")")
            }
            Expr::FnCall(func, args) => {
                let _ = write!(f, "Fn({}, ", func);
                Expr::write_list(f, &args);
                write!(f, ")")
            }
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
        {
            let expr = Expr::new_shape(&vec!["@a"]);
            assert_eq!(r#"Shape(ID(@a))"#, expr.to_string());
        }
        {
            let expr = Expr::new_shape(&vec!["@a", "@b"]);
            assert_eq!(r#"Shape(ID(@a), ID(@b))"#, expr.to_string());
        }
    }

    #[test]
    fn test_arraylt() {
        {
            let expr = Expr::new_array(&vec![1.0]);
            assert_eq!(r#"Array(Float(1.00))"#, expr.to_string());
        }
        {
            let expr = Expr::new_array(&vec![1.0, 2.0]);
            assert_eq!(r#"Array(Float(1.00), Float(2.00))"#, expr.to_string());
        }
    }

    #[test]
    fn test_fn_call() {
        let expr = Expr::FnCall(
            Box::new(Expr::new_id("f")),
            vec![Box::new(Expr::new_id("a")), Box::new(Expr::new_id("b"))],
        );
        assert_eq!("Fn(ID(f), ID(a), ID(b))", expr.to_string());
    }
}
