mod expr;
use crate::base;
pub use expr::*;

pub struct SymTable {}

pub struct Program {
    pub exprs: Vec<Expr>,
    pub sym_table: SymTable,
}

impl Program {
    fn infer_types(&mut self) -> Result<(), base::Error> {
        if self.exprs.is_empty() {
            return Ok(());
        }

        for (i, expr) in self.exprs.iter_mut().enumerate() {
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
        let mut p = Program {
            exprs: Vec::new(),
            sym_table: SymTable {},
        };
        p.infer_types().unwrap();
    }
}
