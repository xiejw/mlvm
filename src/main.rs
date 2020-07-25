use mlvm::syntax::ast::Expr;
use mlvm::syntax::ast::Program;

fn main() {
    let mut p = Program::new(vec![Expr::new_id("a")]);
    p.infer_types().unwrap();
}
