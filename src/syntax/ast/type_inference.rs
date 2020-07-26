use self::super::expr::*;
use self::super::sym_table;
use crate::base::Error;
use sym_table::SymTable;

pub fn infer_type<'a>(
    expr: &'a mut Expr,
    sym_table: &mut SymTable,
) -> Result<&'a Type, Error> {
    match expr {
        Expr::IntLt(tp, _) => infer_trivial_type(
            tp,
            Type::Int,
            "Int Literal should have type Int",
        ),
        Expr::FloatLt(tp, _) => infer_trivial_type(
            tp,
            Type::Float,
            "Float Literal should have type Float",
        ),
        Expr::DimLt(tp, dim) => {
            match tp {
                Type::Unknown => {
                *tp = Type::Dim(dim.clone());
                }
                Type::Dim(dim_in_type) => {
                    if dim !=  dim_in_type {
                        Err(
                    } else {
                        Ok(tp)
                    }
                }
                _ => panic!("unsupported"),
            }
            Ok(tp)
        }
        Expr::ArrayLt(tp, values) => {
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

#[cfg(test)]
mod tests {
    use super::*;

    // #[test]
    // fn test_id() {
    //     let expr = Expr::new_id("abc");
    //     let st = SymTable{};
    //     infer_type(expr, st).unwrap());
    // }

    #[test]
    fn test_intlt() {
        let expr = &mut Expr::new_intlt(123);
        assert_eq!("Int::Int (123)", expr.to_string());
        let st = &mut SymTable {};
        infer_type(expr, st).unwrap();
        assert_eq!("Int::Int (123)", expr.to_string());
    }

    #[test]
    fn test_intlt_unknown() {
        let expr = &mut Expr::IntLt(Type::Unknown, 123);
        assert_eq!("Int::?? (123)", expr.to_string());
        let st = &mut SymTable {};
        infer_type(expr, st).unwrap();
        assert_eq!("Int::Int (123)", expr.to_string());
    }

    #[test]
    #[should_panic = "should have type Int. Got"]
    fn test_intlt_wrong() {
        let expr = &mut Expr::IntLt(Type::Float, 123);
        let st = &mut SymTable {};
        infer_type(expr, st).unwrap();
    }

    #[test]
    fn test_floatlt() {
        let expr = &mut Expr::new_floatlt(123.0);
        assert_eq!("Float::Float (123.00)", expr.to_string());
        let st = &mut SymTable {};
        infer_type(expr, st).unwrap();
        assert_eq!("Float::Float (123.00)", expr.to_string());
    }

    #[test]
    fn test_floatlt_unknown() {
        let expr = &mut Expr::FloatLt(Type::Unknown, 123.0);
        assert_eq!("Float::?? (123.00)", expr.to_string());
        let st = &mut SymTable {};
        infer_type(expr, st).unwrap();
        assert_eq!("Float::Float (123.00)", expr.to_string());
    }

    #[test]
    #[should_panic = "should have type Float. Got"]
    fn test_floatlt_wrong() {
        let expr = &mut Expr::FloatLt(Type::Int, 123.0);
        let st = &mut SymTable {};
        infer_type(expr, st).unwrap();
    }

    // #[test]
    // fn test_stringlt() {
    //     let expr = Expr::new_stringlt("abc".to_string());
    //     assert_eq!(r#"Str("abc")"#, expr.to_string());
    // }

    // #[test]
    // fn test_shapelt() {
    //     {
    //         let expr = Expr::new_shapelt(&vec!["@a"]);
    //         assert_eq!(r#"Shape(ID(@a))"#, expr.to_string());
    //     }
    //     {
    //         let expr = Expr::new_shapelt(&vec!["@a", "@b"]);
    //         assert_eq!(r#"Shape(ID(@a), ID(@b))"#, expr.to_string());
    //     }
    // }

    #[test]
    fn test_arraylt() {
        let expr = &mut Expr::new_arraylt(&vec![1.0, 2.0]);
        assert_eq!(
            r#"Array::Array (Float::Float (1.00), Float::Float (2.00))"#,
            expr.to_string()
        );
        let st = &mut SymTable {};
        infer_type(expr, st).unwrap();
        assert_eq!(
            r#"Array::Array (Float::Float (1.00), Float::Float (2.00))"#,
            expr.to_string()
        );
    }

    #[test]
    fn test_arraylt_unknown() {
        let expr = &mut Expr::ArrayLt(
            Type::Unknown,
            vec![
                Expr::FloatLt(Type::Unknown, 1.),
                Expr::FloatLt(Type::Unknown, 2.),
            ],
        );
        assert_eq!(
            r#"Array::?? (Float::?? (1.00), Float::?? (2.00))"#,
            expr.to_string()
        );
        let st = &mut SymTable {};
        infer_type(expr, st).unwrap();
        assert_eq!(
            r#"Array::Array (Float::Float (1.00), Float::Float (2.00))"#,
            expr.to_string()
        );
    }

    #[test]
    #[should_panic = "should have type Array. Got"]
    fn test_arraylt_wrong_array_type() {
        let expr = &mut Expr::ArrayLt(
            Type::Int,
            vec![
                Expr::FloatLt(Type::Unknown, 1.),
                Expr::FloatLt(Type::Unknown, 2.),
            ],
        );
        let st = &mut SymTable {};
        infer_type(expr, st).unwrap();
    }

    #[test]
    #[should_panic = "Array element should only have Float type element"]
    fn test_arraylt_wrong_element_type() {
        let expr = &mut Expr::ArrayLt(
            Type::Unknown,
            vec![Expr::new_intlt(23), Expr::FloatLt(Type::Unknown, 2.)],
        );
        let st = &mut SymTable {};
        infer_type(expr, st).unwrap();
    }

    // #[test]
    // fn test_fn_call() {
    //     let expr = Expr::FnCall(
    //         Type::Unknown,
    //         Box::new(Expr::new_id("f")),
    //         vec![Expr::new_id("a"), Expr::new_id("b")],
    //     );
    //     assert_eq!("Fn(ID(f), ID(a), ID(b))", expr.to_string());
    // }
}
