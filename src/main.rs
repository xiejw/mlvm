use mlvm::syntax::ast::Expr;
use mlvm::syntax::ast::Kind;
use mlvm::syntax::ast::Program;
use mlvm::syntax::ast::Type::Unknown;
use std::rc::Rc;

fn main() {
    let mut p = Program::new(vec![
        Expr {
            etype: Unknown,
            kind: Kind::IntLt(23),
        },
        Expr {
            etype: Unknown,
            kind: Kind::ArrayLt(vec![
                Expr {
                    etype: Unknown,
                    kind: Kind::FloatLt(123.),
                },
                Expr::new_floatlt(7.89),
            ]),
        },
        Expr {
            etype: Unknown,
            kind: Kind::DimLt(Rc::new("@a".to_string())),
        },
    ]);
    print!("before:\n{}", p);

    if let Err(err) = p.infer_types() {
        print!("\nunexpected error:\n{}\n", err);
        return;
    };

    print!("after:\n{}", p);
}
