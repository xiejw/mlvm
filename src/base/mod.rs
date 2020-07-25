use std::fmt;

#[derive(Debug)]
pub struct Error {
    notes: Vec<String>, // Reverse order
}

impl Error {
    pub fn new() -> Self {
        Error { notes: Vec::new() }
    }

    pub fn emit_diagnosis_note(&mut self, note: String) -> &mut Self {
        self.notes.push(note);
        self
    }

    pub fn emit_diagnosis_note_str(&mut self, note: &str) -> &mut Self {
        self.notes.push(note.to_string());
        self
    }
}

impl fmt::Display for Error {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let notes = &self.notes;
        if notes.is_empty() {
            return write!(f, "(no error message)");
        }

        let mut indent = String::from("");
        for i in (0..notes.len()).rev() {
            if i == 0 {
                let _ = write!(f, "{}+-> {}", indent, notes[i]);
            } else {
                let _ = write!(f, "{}+-+ {}\n", indent, notes[i]);
            }
            indent += "  ";
        }

        Ok(())
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_display_for_new_error() {
        assert_eq!("(no error message)", format!("{}", Error::new()));
    }

    #[test]
    fn test_display_for_error() {
        let mut err = Error::new();
        err.emit_diagnosis_note_str("during stack 1");
        err.emit_diagnosis_note_str("during stack 2");
        assert_eq!(
            r#"+-+ during stack 2
  +-> during stack 1"#,
            format!("{}", err)
        );
    }

    #[test]
    fn test_display_for_error_chain() {
        let mut err = Error::new();
        err.emit_diagnosis_note_str("during stack 1")
            .emit_diagnosis_note_str("during stack 2")
            .emit_diagnosis_note_str("during stack 3");
        assert_eq!(
            r#"+-+ during stack 3
  +-+ during stack 2
    +-> during stack 1"#,
            format!("{}", err)
        );
    }
}
