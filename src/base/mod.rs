use std::fmt;

#[derive(Debug)]
pub struct Error {
    notes: Option<Vec<String>>, // Reverse order
}

impl Error {
    pub fn new() -> Self {
        Error { notes: None }
    }
    pub fn from(err: impl std::error::Error) -> Self {
        Error::new().emit_diagnosis_note(err.to_string()).take()
    }

    pub fn emit_diagnosis_note(&mut self, note: String) -> &mut Self {
        if let Some(ref mut notes) = self.notes {
            notes.push(note);
        } else {
            self.notes = Some(vec![note])
        }
        self
    }

    pub fn emit_diagnosis_note_str(&mut self, note: &str) -> &mut Self {
        return self.emit_diagnosis_note(note.to_string());
    }

    pub fn take(&mut self) -> Self {
        Error {
            notes: self.notes.take(),
        }
    }
}

impl fmt::Display for Error {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let _ = write!(f, "\n###################################################\n");
        if let Some(ref notes) = self.notes {
            let mut indent = String::from("");
            for i in (0..notes.len()).rev() {
                if i == 0 {
                    let _ = write!(f, "{}+-> {}", indent, notes[i]);
                } else {
                    let _ = write!(f, "{}+-+ {}\n", indent, notes[i]);
                }
                indent += "  ";
            }
        } else {
            let _ = write!(f, "(no error message)");
        }

        let _ = write!(f, "\n###################################################\n");
        Ok(())
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_display_for_new_error() {
        assert_eq!(
            r#"
###################################################
(no error message)
###################################################
"#,
            format!("{}", Error::new())
        );
    }

    #[test]
    fn test_display_for_error() {
        let err = Error::new()
            .emit_diagnosis_note_str("during stack 1")
            .emit_diagnosis_note_str("during stack 2")
            .take();
        assert_eq!(
            r#"
###################################################
+-+ during stack 2
  +-> during stack 1
###################################################
"#,
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
            r#"
###################################################
+-+ during stack 3
  +-+ during stack 2
    +-> during stack 1
###################################################
"#,
            format!("{}", err)
        );
    }
}
