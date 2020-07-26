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
}
