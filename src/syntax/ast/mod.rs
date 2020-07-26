mod expr;
pub use expr::*;

use crate::base::Error;
use std::fmt;

pub struct SymTable {}

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
            let result = Program::infer_type(expr, sym_table);

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

    fn infer_type(
        expr: &mut Expr,
        sym_table: &mut SymTable,
    ) -> Result<(), Error> {
        let result = match expr {
            Expr::IntLt(ref mut tp, _) => Program::infer_trivial_type(
                tp,
                Type::Int,
                "Int Literal should have type Int",
            ),
            Expr::FloatLt(ref mut tp, _) => Program::infer_trivial_type(
                tp,
                Type::Float,
                "Float Literal should have type Float",
            ),
            Expr::ArrayLt(ref mut tp, ref mut values) => {
                let mut result = Program::infer_trivial_type(
                    tp,
                    Type::Array,
                    "Array Literal should have type Array",
                );

                if !result.is_err() {
                    for (i, expr) in values.iter_mut().enumerate() {
                        result = Program::infer_type_with_expectation(
                            expr,
                            &Type::Float,
                            sym_table,
                        );
                        if let Err(ref mut err) = result {
                            err.emit_diagnosis_note(format!(
                                    "Array element should only have Float type element. At {}-th, type assertion failed", i));
                            break;
                        }
                    }
                }

                result
            }
            _ => Err(Error::new()
                .emit_diagnosis_note_str("un supported expr type yet")
                .take()),
        };

        result
    }

    fn infer_trivial_type(
        tp: &mut Type,
        expected: Type,
        msg: &str,
    ) -> Result<(), Error> {
        if *tp == expected {
            Ok(())
        } else if tp == &Type::Unknown {
            *tp = expected;
            Ok(())
        } else {
            Err(Error::new()
                .emit_diagnosis_note(format!("{}. Got: {}", msg, tp))
                .take())
        }
    }

    fn infer_type_with_expectation(
        expr: &mut Expr,
        expected_type: &Type,
        sym_table: &mut SymTable,
    ) -> Result<(), Error> {
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
