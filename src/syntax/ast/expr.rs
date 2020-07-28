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
            Type::Dim(dim) => write!(f, "{}", dim), // No starting Dim.
            _ => panic!("unsuported type"),
        }
    }
}

#[derive(PartialEq, Debug)]
pub enum Kind {
    ID(String),
    IntLt(i64),
    FloatLt(f32),
    ShapeLt(Vec<Expr>),
    DimLt(Rc<String>),
    ArrayLt(Vec<Expr>),
    StringLt(String),
    FnCall(Box<Expr>, Vec<Expr>),
}

#[derive(PartialEq, Debug)]
pub struct Expr {
    pub etype: Type,
    pub kind: Kind,
}

impl Expr {
    pub fn new_intlt(v: i64) -> Expr {
        Expr {
            etype: Type::Int,
            kind: Kind::IntLt(v),
        }
    }

    pub fn new_floatlt(v: f32) -> Expr {
        Expr {
            etype: Type::Float,
            kind: Kind::FloatLt(v),
        }
    }

    pub fn new_stringlt(v: String) -> Expr {
        Expr {
            etype: Type::String,
            kind: Kind::StringLt(v),
        }
    }

    pub fn new_id(v: &str) -> Expr {
        Expr {
            etype: Type::Unknown,
            kind: Kind::ID(v.to_string()),
        }
    }

    pub fn new_shapelt(dims: &[&str]) -> Expr {
        Expr {
            etype: Type::Unknown,
            kind: Kind::ShapeLt(
                dims.iter()
                    .map(|x| Expr {
                        etype: Type::Unknown,
                        kind: Kind::DimLt(Rc::new(x.to_string())),
                    })
                    .collect(),
            ),
        }
    }

    pub fn new_arraylt(values: &[f32]) -> Expr {
        Expr {
            etype: Type::Array,
            kind: Kind::ArrayLt(values.iter().map(|x| Expr::new_floatlt(*x)).collect()),
        }
    }
}

impl fmt::Display for Expr {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let tp = &self.etype;
        match &self.kind {
            Kind::ID(v) => write!(f, "ID({})", v),
            Kind::IntLt(v) => write!(f, "IntLt::{} ({})", tp, v),
            Kind::FloatLt(v) => write!(f, "FloatLt::{} ({:.2})", tp, v),
            Kind::StringLt(v) => write!(f, "StringLt(\"{}\")", v),
            Kind::DimLt(s) => write!(f, "DimLt::{} ({})", tp, s),
            Kind::ShapeLt(l) => {
                let _ = write!(f, "ShapeLt(");
                Expr::write_list(f, &l);
                write!(f, ")")
            }
            Kind::ArrayLt(l) => {
                let _ = write!(f, "ArrayLt::{} (", tp);
                Expr::write_list(f, &l);
                write!(f, ")")
            }
            Kind::FnCall(func, args) => {
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
        assert_eq!("IntLt::Int (123)", expr.to_string());
    }

    #[test]
    fn test_floatlt() {
        let expr = Expr::new_floatlt(123.0);
        assert_eq!("FloatLt::Float (123.00)", expr.to_string());
    }

    #[test]
    fn test_stringlt() {
        let expr = Expr::new_stringlt("abc".to_string());
        assert_eq!(r#"StringLt("abc")"#, expr.to_string());
    }

    #[test]
    fn test_shapelt() {
        {
            let expr = Expr::new_shapelt(&vec!["@a"]);
            assert_eq!(r#"ShapeLt(DimLt::?? (@a))"#, expr.to_string());
        }
        {
            let expr = Expr::new_shapelt(&vec!["@a", "@b"]);
            assert_eq!(
                r#"ShapeLt(DimLt::?? (@a), DimLt::?? (@b))"#,
                expr.to_string()
            );
        }
    }

    #[test]
    fn test_arraylt() {
        {
            let expr = Expr::new_arraylt(&vec![1.0]);
            assert_eq!(
                r#"ArrayLt::Array (FloatLt::Float (1.00))"#,
                expr.to_string()
            );
        }
        {
            let expr = Expr::new_arraylt(&vec![1.0, 2.0]);
            assert_eq!(
                r#"ArrayLt::Array (FloatLt::Float (1.00), FloatLt::Float (2.00))"#,
                expr.to_string()
            );
        }
    }

    #[test]
    fn test_fn_call() {
        let expr = Expr {
            etype: Type::Unknown,
            kind: Kind::FnCall(
                Box::new(Expr::new_id("f")),
                vec![Expr::new_id("a"), Expr::new_id("b")],
            ),
        };
        assert_eq!("Fn(ID(f), ID(a), ID(b))", expr.to_string());
    }
}
