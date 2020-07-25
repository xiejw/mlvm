use mlvm::syntax::ast::Expr;
use mlvm::syntax::ast::Program;
use mlvm::syntax::ast::Type;

fn main() {
    let mut p = Program::new(vec![Expr::IntLt(Type::Unknown, 123)]);
    p.infer_types().unwrap();
}
