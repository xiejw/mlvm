use self::super::expr::*;
use self::super::sym_table;
use crate::base::Error;
use sym_table::SymTable;

pub fn infer_type(
    expr: &mut Expr,
    sym_table: &mut SymTable,
) -> Result<(), Error> {
    let result = match expr {
        Expr::IntLt(ref mut tp, _) => infer_trivial_type(
            tp,
            Type::Int,
            "Int Literal should have type Int",
        ),
        Expr::FloatLt(ref mut tp, _) => infer_trivial_type(
            tp,
            Type::Float,
            "Float Literal should have type Float",
        ),
        Expr::ArrayLt(ref mut tp, ref mut values) => {
            let mut result = infer_trivial_type(
                tp,
                Type::Array,
                "Array Literal should have type Array",
            );

            if !result.is_err() {
                for (i, expr) in values.iter_mut().enumerate() {
                    result = infer_type_with_expectation(
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
