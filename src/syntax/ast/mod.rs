mod expr;
pub use expr::*;

use crate::base::Error;
use std::fmt;

pub struct SymTable {}

pub struct Program {
    exprs: Vec<Expr>,
    // sym_table: SymTable,
}

impl Program {
    pub fn new(exprs: Vec<Expr>) -> Program {
        Program {
            exprs: exprs,
            //      sym_table: SymTable {},
        }
    }

    pub fn infer_types(&mut self) -> Result<(), Error> {
        if self.exprs.is_empty() {
            return Ok(());
        }

        for (i, expr) in self.exprs.iter_mut().enumerate() {
            let result = match expr {
                Expr::IntLt(ref mut tp, _) => match tp {
                    Type::Int => Ok(()),
                    Type::Unknown => {
                        *tp = Type::Int;
                        Ok(())
                    }
                    _ => Err(Error::new()
                        .emit_diagnosis_note(format!(
                            "Int Literal should have type Int. Got: {}",
                            tp
                        ))
                        .take()),
                },
                _ => Err(Error::new()
                    .emit_diagnosis_note_str("un supported expr type yet")
                    .take()),
            };

            if let Err(mut err) = result {
                return Err(err
                    .emit_diagnosis_note(format!(
                        "failed to infer type for {}-th expr: {}",
                        i, expr
                    ))
                    .take());
            }
        }

        Ok(())
    }
}

impl fmt::Display for Program {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        for (i, expr) in self.exprs.iter().enumerate() {
            let _ = write!(f, "%{}: {}\n", i, expr);
        }

        Ok(())
    }
}

#[cfg(test)]
mod tests {

    use self::super::*;

    #[test]
    fn test_infer() {
        let mut p = Program::new(Vec::new());
        p.infer_types().unwrap();
    }

    #[test]
    fn test_infer_int() {
        let mut p = Program::new(vec![Expr::new_intlt(123)]);
        p.infer_types().unwrap();
    }

    #[test]
    fn test_infer_int_with_unknown_type() {
        let mut p = Program::new(vec![Expr::IntLt(Type::Unknown, 123)]);
        p.infer_types().unwrap();
        assert_eq!(&Expr::IntLt(Type::Int, 123), &p.exprs[0]);
    }

    #[test]
    #[should_panic = "should have type Int. Got: Float"]
    fn test_infer_int_with_wrong_type() {
        let mut p = Program::new(vec![Expr::IntLt(Type::Float, 123)]);
        p.infer_types().unwrap();
    }
}
