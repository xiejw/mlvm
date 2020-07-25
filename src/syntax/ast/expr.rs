use std::fmt;

pub enum Type {
    Unknown,
    Int,
    Float,
    Dim,
    Shape,
    Array,
    String,
    Fn {
        inputs: Vec<Type>,
        outputs: Vec<Type>,
    },
}

impl fmt::Display for Type {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            Type::Unknown => write!(f, "??"),
            Type::Int => write!(f, "Int"),
            Type::Float => write!(f, "Float"),
            _ => panic!("unsuported type"),
        }
    }
}

pub enum Expr {
    ID(Type, String),
    IntLt(Type, i64),
    FloatLt(Type, f32),
    ShapeLt(Type, Vec<Expr>),
    ArrayLt(Type, Vec<Expr>),
    StringLt(Type, String),
    FnCall(Type, Box<Expr>, Vec<Expr>),
}

impl Expr {
    pub fn new_intlt(v: i64) -> Expr {
        Expr::IntLt(Type::Int, v)
    }

    pub fn new_floatlt(v: f32) -> Expr {
        Expr::FloatLt(Type::Float, v)
    }

    pub fn new_stringlt(v: String) -> Expr {
        Expr::StringLt(Type::String, v)
    }

    pub fn new_id(v: &str) -> Expr {
        Expr::ID(Type::Unknown, v.to_string())
    }

    pub fn new_shapelt(dims: &[&str]) -> Expr {
        Expr::ShapeLt(Type::Shape, dims.iter().map(|x| Expr::new_id(x)).collect())
    }

    pub fn new_arraylt(values: &[f32]) -> Expr {
        Expr::ArrayLt(
            Type::Array,
            values
                .iter()
                .map(|x| Expr::FloatLt(Type::Float, *x))
                .collect(),
        )
    }
}

impl fmt::Display for Expr {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match &self {
            Expr::ID(_, v) => write!(f, "ID({})", v),
            Expr::IntLt(ref tp, v) => write!(f, "Int::{} ({})", tp, v),
            Expr::FloatLt(_, v) => write!(f, "Float({:.2})", v),
            Expr::StringLt(_, v) => write!(f, "Str(\"{}\")", v),
            Expr::ShapeLt(_, l) => {
                let _ = write!(f, "Shape(");
                Expr::write_list(f, &l);
                write!(f, ")")
            }
            Expr::ArrayLt(_, l) => {
                let _ = write!(f, "Array(");
                Expr::write_list(f, &l);
                write!(f, ")")
            }
            Expr::FnCall(_, func, args) => {
                let _ = write!(f, "Fn({}, ", func);
                Expr::write_list(f, &args);
                write!(f, ")")
            }
        }
    }
}

impl Expr {
    fn write_list(f: &mut fmt::Formatter<'_>, list: &Vec<Expr>) {
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
        let expr = Expr::new_intlt(123);
        assert_eq!("Int::Int (123)", expr.to_string());
    }

    #[test]
    fn test_floatlt() {
        let expr = Expr::new_floatlt(123.0);
        assert_eq!("Float(123.00)", expr.to_string());
    }

    #[test]
    fn test_stringlt() {
        let expr = Expr::new_stringlt("abc".to_string());
        assert_eq!(r#"Str("abc")"#, expr.to_string());
    }

    #[test]
    fn test_shapelt() {
        {
            let expr = Expr::new_shapelt(&vec!["@a"]);
            assert_eq!(r#"Shape(ID(@a))"#, expr.to_string());
        }
        {
            let expr = Expr::new_shapelt(&vec!["@a", "@b"]);
            assert_eq!(r#"Shape(ID(@a), ID(@b))"#, expr.to_string());
        }
    }

    #[test]
    fn test_arraylt() {
        {
            let expr = Expr::new_arraylt(&vec![1.0]);
            assert_eq!(r#"Array(Float(1.00))"#, expr.to_string());
        }
        {
            let expr = Expr::new_arraylt(&vec![1.0, 2.0]);
            assert_eq!(r#"Array(Float(1.00), Float(2.00))"#, expr.to_string());
        }
    }

    #[test]
    fn test_fn_call() {
        let expr = Expr::FnCall(
            Type::Unknown,
            Box::new(Expr::new_id("f")),
            vec![Expr::new_id("a"), Expr::new_id("b")],
        );
        assert_eq!("Fn(ID(f), ID(a), ID(b))", expr.to_string());
    }
}
