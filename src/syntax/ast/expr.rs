use std::fmt;
use std::rc::Rc;

#[derive(PartialEq, Debug)]
pub enum Type {
    Unknown,
    Int,
    Float,
    Dim(Rc<String>),
    Shape(Vec<Rc<String>>),
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
            Type::Array => write!(f, "Array"),
            _ => panic!("unsuported type"),
        }
    }
}

#[derive(PartialEq, Debug)]
pub enum Expr {
    ID(Type, String),
    IntLt(Type, i64),
    FloatLt(Type, f32),
    ShapeLt(Type, Vec<Expr>),
    DimLt(Type, Rc<String>),
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
        Expr::ShapeLt(
            Type::Unknown,
            dims.iter()
                .map(|x| Expr::DimLt(Type::Unknown, Rc::new(x.to_string())))
                .collect(),
        )
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
            Expr::IntLt(tp, v) => write!(f, "Int::{} ({})", tp, v),
            Expr::FloatLt(tp, v) => write!(f, "Float::{} ({:.2})", tp, v),
            Expr::StringLt(_, v) => write!(f, "Str(\"{}\")", v),
            Expr::DimLt(_, s) => write!(f, "Dim({})", s),
            Expr::ShapeLt(_, l) => {
                let _ = write!(f, "Shape(");
                Expr::write_list(f, &l);
                write!(f, ")")
            }
            Expr::ArrayLt(ref tp, l) => {
                let _ = write!(f, "Array::{} (", tp);
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
        assert_eq!("Float::Float (123.00)", expr.to_string());
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
            assert_eq!(r#"Shape(Dim(@a))"#, expr.to_string());
        }
        {
            let expr = Expr::new_shapelt(&vec!["@a", "@b"]);
            assert_eq!(r#"Shape(Dim(@a), Dim(@b))"#, expr.to_string());
        }
    }

    #[test]
    fn test_arraylt() {
        {
            let expr = Expr::new_arraylt(&vec![1.0]);
            assert_eq!(
                r#"Array::Array (Float::Float (1.00))"#,
                expr.to_string()
            );
        }
        {
            let expr = Expr::new_arraylt(&vec![1.0, 2.0]);
            assert_eq!(
                r#"Array::Array (Float::Float (1.00), Float::Float (2.00))"#,
                expr.to_string()
            );
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
