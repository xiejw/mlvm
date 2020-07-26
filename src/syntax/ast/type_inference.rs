use self::super::expr::*;
use self::super::sym_table;
use crate::base::Error;
use sym_table::SymTable;

pub fn infer_type<'a>(
    expr: &'a mut Expr,
    sym_table: &mut SymTable,
) -> Result<&'a Type, Error> {
    match expr {
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
            {
                // Check tp.
                let result = infer_trivial_type(
                    tp,
                    Type::Array,
                    "Array Literal should have type Array",
                );

                if result.is_err() {
                    return Err(result.unwrap_err());
                }
            }
            {
                let result = infer_elements_with_same_type(
                    values,
                    &Type::Float,
                    sym_table,
                    "Array element should only have Float type element",
                );
                if result.is_err() {
                    return Err(result.unwrap_err());
                }
            }

            Ok(tp)
        }
        _ => Err(Error::new()
            .emit_diagnosis_note(format!("unsupported expr type yet: {}", expr))
            .take()),
    }
}

fn infer_trivial_type<'a>(
    tp: &'a mut Type,
    expected: Type,
    msg: &str,
) -> Result<&'a Type, Error> {
    if *tp == expected {
        Ok(tp)
    } else if tp == &Type::Unknown {
        *tp = expected;
        Ok(tp)
    } else {
        Err(Error::new()
            .emit_diagnosis_note(format!("{}. Got: {}", msg, tp))
            .take())
    }
}

fn infer_elements_with_same_type(
    values: &mut Vec<Expr>,
    expected_type: &Type,
    sym_table: &mut SymTable,
    msg: &str,
) -> Result<(), Error> {
    for (i, expr) in values.iter_mut().enumerate() {
        let mut result =
            infer_type_with_expectation(expr, expected_type, sym_table);
        if let Err(ref mut err) = result {
            return Err(err
                .emit_diagnosis_note(format!(
                    "{}. At {}-th, type assertion failed",
                    msg, i
                ))
                .take());
        }
    }
    Ok(())
}

fn infer_type_with_expectation(
    expr: &mut Expr,
    expected_type: &Type,
    sym_table: &mut SymTable,
) -> Result<(), Error> {
    match infer_type(expr, sym_table) {
        Ok(tp) => {
            if tp != expected_type {
                return Err(Error::new()
                    .emit_diagnosis_note(format!(
                        "expected type: {}, got: {}",
                        expected_type, tp
                    ))
                    .take());
            }
            Ok(())
        }
        Err(err) => Err(err),
    }
}
