pub struct Program {}

pub enum Expr{
    ID(String),
    IntLt(i64),
    FnCall(Box<Expr>),
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
