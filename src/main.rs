use mlvm::syntax::ast::Expr;
use mlvm::syntax::ast::Program;
use mlvm::syntax::ast::Type::Unknown;
use std::rc::Rc;

fn main() {
    let mut p = Program::new(vec![
        Expr::IntLt(Unknown, 23),
        Expr::ArrayLt(
            Unknown,
            vec![Expr::FloatLt(Unknown, 123.), Expr::new_floatlt(7.89)],
        ),
        Expr::DimLt(Unknown, Rc::new("@a".to_string())),
    ]);
    print!("before:\n{}", p);

    if let Err(err) = p.infer_types() {
        print!("\nunexpected error:\n{}\n", err);
        return;
    };

    print!("after:\n{}", p);
}
