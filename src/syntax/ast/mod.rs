mod expr;
pub use expr::*;

mod sym_table;
use sym_table::SymTable;
mod type_inference;

use crate::base::Error;
use std::fmt;

pub struct Program {
    exprs: Vec<Expr>,
    sym_table: SymTable,
}

impl Program {
    pub fn new(exprs: Vec<Expr>) -> Program {
        Program {
            exprs: exprs,
            sym_table: SymTable {},
        }
    }

    pub fn infer_types(&mut self) -> Result<(), Error> {
        if self.exprs.is_empty() {
            return Ok(());
        }

        let sym_table = &mut self.sym_table;
        for (i, expr) in self.exprs.iter_mut().enumerate() {
            let result = type_inference::infer_type(expr, sym_table);

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
        let _ = write!(f, "\n{{\n");
        for (i, expr) in self.exprs.iter().enumerate() {
            let _ = write!(f, "  %{:<3}: {}\n", i, expr);
        }
        let _ = write!(f, "}}\n");

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
    fn test_infer_intlt() {
        let mut p = Program::new(vec![Expr::new_intlt(123)]);
        p.infer_types().unwrap();
    }

    #[test]
    fn test_infer_intlt_with_unknown_type() {
        let mut p = Program::new(vec![Expr::IntLt(Type::Unknown, 123)]);
        p.infer_types().unwrap();
        assert_eq!(&Expr::IntLt(Type::Int, 123), &p.exprs[0]);
    }

    #[test]
    #[should_panic = "should have type Int. Got: Float"]
    fn test_infer_intlt_with_wrong_type() {
        let mut p = Program::new(vec![Expr::IntLt(Type::Float, 123)]);
        p.infer_types().unwrap();
    }

    #[test]
    fn test_infer_floatlt() {
        let mut p = Program::new(vec![Expr::new_floatlt(123.3)]);
        p.infer_types().unwrap();
    }

    #[test]
    fn test_infer_floatlt_with_unknown_type() {
        let mut p = Program::new(vec![Expr::FloatLt(Type::Unknown, 123.3)]);
        p.infer_types().unwrap();
        assert_eq!(&Expr::FloatLt(Type::Float, 123.3), &p.exprs[0]);
    }

    #[test]
    #[should_panic = "should have type Float. Got:"]
    fn test_infer_floatlt_with_wrong_type() {
        let mut p = Program::new(vec![Expr::FloatLt(Type::Int, 123.)]);
        p.infer_types().unwrap();
    }
}