mod expr;
mod program;
mod sym_table;
mod type_inference;

// Only exposes the following symbols at this mod.
pub use expr::Expr;
pub use expr::Kind;
pub use expr::Type;
pub use program::Program;
