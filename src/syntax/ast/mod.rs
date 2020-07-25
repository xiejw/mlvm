mod expr;
use crate::base::Error;

pub use expr::*;

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

        for (i, ref mut expr) in self.exprs.iter_mut().enumerate() {
            match expr {
                Expr::IntLt(ref mut tp, _) => match tp {
                    Type::Int => {}
                    Type::Unknown => *tp = Type::Int,
                    _ => panic!("123"),
                },
                _ => {
                    return Err(Error::new()
                        .emit_diagnosis_note_str("un supported expr type yet")
                        .take());
                }
            }
            println!("handle {}: {}", i, expr);
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
}
