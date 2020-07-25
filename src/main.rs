use mlvm::syntax::ast::Expr;
use mlvm::syntax::ast::Program;
use mlvm::syntax::ast::Type;

fn main() {
    let mut p = Program::new(vec![Expr::IntLt(Type::Float, 123)]);
    print!("before:\n{}", p);

    if let Err(err) = p.infer_types() {
        print!("\nunexpected error:\n{}\n", err);
        return;
    };

    print!("after:\n{}", p);
}
