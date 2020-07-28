use self::super::expr::*;
use self::super::sym_table;
use crate::base::Error;
use std::rc::Rc;
use sym_table::SymTable;

// Infers the type of `expr` from bottom-up.
pub fn infer_type<'a>(expr: &'a mut Expr, sym_table: &mut SymTable) -> Result<&'a Type, Error> {
    match &mut expr.kind {
        Kind::IntLt(_) => infer_trivial_type(&mut expr.etype, Type::Int, "Int Literal"),
        Kind::FloatLt(_) => infer_trivial_type(&mut expr.etype, Type::Float, "Float Literal"),
        Kind::DimLt(dim) => infer_dimlt_type(&mut expr.etype, dim),
        Kind::ArrayLt(values) => infer_array_type(&mut expr.etype, values, sym_table),
        k => Err(Error::new()
            .emit_diagnosis_note(format!("unsupported expr kind yet: {:?}", k))
            .take()),
    }
}

fn infer_trivial_type<'a>(
    tp: &'a mut Type,
    expected: Type,
    title: &str,
) -> Result<&'a Type, Error> {
    if *tp == expected {
        Ok(tp)
    } else if tp == &Type::Unknown {
        *tp = expected;
        Ok(tp)
    } else {
        Err(Error::new()
            .emit_diagnosis_note(format!(
                "{} should have type {}. Got: {}",
                title, expected, tp
            ))
            .take())
    }
}

fn infer_dimlt_type<'a>(etype: &'a mut Type, dim: &Rc<String>) -> Result<&'a Type, Error> {
    match &mut *etype {
        tp if *tp == Type::Unknown => {
            *tp = Type::Dim(dim.clone());
        }
        Type::Dim(ref dim_in_type) => {
            if dim != dim_in_type {
                return Err(Error::new()
                    .emit_diagnosis_note(format!(
                        "Dim type has size ({}), but the literal has size ({})",
                        dim, dim_in_type
                    ))
                    .take());
            }
        }
        etype => {
            return Err(Error::new()
                .emit_diagnosis_note(format!(
                    "Dim literal can only have Dim type. Got: {}",
                    etype
                ))
                .take());
        }
    };
    Ok(etype)
}

fn infer_array_type<'a>(
    etype: &'a mut Type,
    values: &mut Vec<Expr>,
    sym_table: &mut SymTable,
) -> Result<&'a Type, Error> {
    infer_trivial_type(&mut *etype, Type::Array, "Array Literal")?;
    infer_elements_with_same_type(
        values,
        &Type::Float,
        sym_table,
        "Array element should only have Float type element",
    )?;
    Ok(etype)
}

fn infer_elements_with_same_type(
    values: &mut Vec<Expr>,
    expected_type: &Type,
    sym_table: &mut SymTable,
    msg: &str,
) -> Result<(), Error> {
    for (i, expr) in values.iter_mut().enumerate() {
        let mut result = infer_type_with_expectation(expr, expected_type, sym_table);
        if let Err(ref mut err) = result {
            return Err(err
                .emit_diagnosis_note(format!("{}. At {}-th, type assertion failed", msg, i))
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
                    .emit_diagnosis_note(format!("expected type: {}, got: {}", expected_type, tp))
                    .take());
            }
            Ok(())
        }
        Err(err) => Err(err),
    }
}

#[cfg(test)]
mod tests {
    use super::Type::Unknown;
    use super::*;
    use std::rc::Rc;

    // #[test]
    // fn test_id() {
    //     let expr = Expr::new_id("abc");
    //     let st = SymTable{};
    //     infer_type(expr, st).unwrap());
    // }

    #[test]
    fn test_intlt() {
        let expr = &mut Expr::new_intlt(123);
        assert_eq!("IntLt::Int (123)", expr.to_string());
        let st = &mut SymTable {};
        infer_type(expr, st).unwrap();
        assert_eq!("IntLt::Int (123)", expr.to_string());
    }

    #[test]
    fn test_intlt_unknown() {
        let expr = &mut Expr {
            etype: Type::Unknown,
            kind: Kind::IntLt(123),
        };
        assert_eq!("IntLt::?? (123)", expr.to_string());
        let st = &mut SymTable {};
        infer_type(expr, st).unwrap();
        assert_eq!("IntLt::Int (123)", expr.to_string());
    }

    #[test]
    #[should_panic = "should have type Int. Got"]
    fn test_intlt_wrong() {
        let expr = &mut Expr {
            etype: Type::Float,
            kind: Kind::IntLt(123),
        };
        let st = &mut SymTable {};
        infer_type(expr, st).unwrap();
    }

    #[test]
    fn test_floatlt() {
        let expr = &mut Expr::new_floatlt(123.0);
        assert_eq!("FloatLt::Float (123.00)", expr.to_string());
        let st = &mut SymTable {};
        infer_type(expr, st).unwrap();
        assert_eq!("FloatLt::Float (123.00)", expr.to_string());
    }

    #[test]
    fn test_floatlt_unknown() {
        let expr = &mut Expr {
            etype: Type::Unknown,
            kind: Kind::FloatLt(123.0),
        };
        assert_eq!("FloatLt::?? (123.00)", expr.to_string());
        let st = &mut SymTable {};
        infer_type(expr, st).unwrap();
        assert_eq!("FloatLt::Float (123.00)", expr.to_string());
    }

    #[test]
    #[should_panic = "should have type Float. Got"]
    fn test_floatlt_wrong() {
        let expr = &mut Expr {
            etype: Type::Int,
            kind: Kind::FloatLt(123.0),
        };
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
            r#"ArrayLt::Array (FloatLt::Float (1.00), FloatLt::Float (2.00))"#,
            expr.to_string()
        );
        let st = &mut SymTable {};
        infer_type(expr, st).unwrap();
        assert_eq!(
            r#"ArrayLt::Array (FloatLt::Float (1.00), FloatLt::Float (2.00))"#,
            expr.to_string()
        );
    }

    #[test]
    fn test_arraylt_unknown() {
        let expr = &mut Expr {
            etype: Unknown,
            kind: Kind::ArrayLt(vec![
                Expr {
                    etype: Unknown,
                    kind: Kind::FloatLt(1.0),
                },
                Expr {
                    etype: Unknown,
                    kind: Kind::FloatLt(2.0),
                },
            ]),
        };
        assert_eq!(
            r#"ArrayLt::?? (FloatLt::?? (1.00), FloatLt::?? (2.00))"#,
            expr.to_string()
        );
        let st = &mut SymTable {};
        infer_type(expr, st).unwrap();
        assert_eq!(
            r#"ArrayLt::Array (FloatLt::Float (1.00), FloatLt::Float (2.00))"#,
            expr.to_string()
        );
    }

    #[test]
    #[should_panic = "Array Literal should have type Array. Got"]
    fn test_arraylt_wrong_array_type() {
        let expr = &mut Expr {
            etype: Type::Int,
            kind: Kind::ArrayLt(vec![Expr::new_floatlt(1.0)]),
        };
        let st = &mut SymTable {};
        infer_type(expr, st).unwrap();
    }

    #[test]
    #[should_panic = "Array element should only have Float type elemen"]
    fn test_arraylt_wrong_ele_type() {
        let expr = &mut Expr {
            etype: Unknown,
            kind: Kind::ArrayLt(vec![Expr {
                etype: Unknown,
                kind: Kind::IntLt(1),
            }]),
        };
        let st = &mut SymTable {};
        infer_type(expr, st).unwrap();
    }

    #[test]
    fn test_dimlt() {
        let expr = &mut Expr {
            etype: Type::Dim(Rc::new("@a".to_string())),
            kind: Kind::DimLt(Rc::new("@a".to_string())),
        };
        assert_eq!(r#"DimLt::@a (@a)"#, expr.to_string());
        let st = &mut SymTable {};
        infer_type(expr, st).unwrap();
        assert_eq!(r#"DimLt::@a (@a)"#, expr.to_string());
    }

    #[test]
    fn test_dimlt_unknown() {
        let expr = &mut Expr {
            etype: Type::Unknown,
            kind: Kind::DimLt(Rc::new("@a".to_string())),
        };
        assert_eq!(r#"DimLt::?? (@a)"#, expr.to_string());
        let st = &mut SymTable {};
        infer_type(expr, st).unwrap();
        assert_eq!(r#"DimLt::@a (@a)"#, expr.to_string());
    }

    #[test]
    #[should_panic = "Dim type has size (@a)"]
    fn test_dimlt_wrong_dim() {
        let expr = &mut Expr {
            etype: Type::Dim(Rc::new("@b".to_string())),
            kind: Kind::DimLt(Rc::new("@a".to_string())),
        };
        let st = &mut SymTable {};
        infer_type(expr, st).unwrap();
    }

    #[test]
    #[should_panic = "Dim literal can only have Dim type. Got: Int"]
    fn test_dimlt_wrong_type() {
        let expr = &mut Expr {
            etype: Type::Int,
            kind: Kind::DimLt(Rc::new("@a".to_string())),
        };
        let st = &mut SymTable {};
        infer_type(expr, st).unwrap();
    }
}
