---
name: file-notes
description: Saves a Q&A-style note about how a specific Go file/example works, right alongside that file. Use when the user says "file-notes", "note this", "save this explanation", or has been asking questions about how a particular file/snippet works and wants that captured.
---

# File Notes

Capture question-and-answer style notes about a *specific* file the user is
running/reading right now, and save them next to that file — so the
explanation stays attached to the exact example it's about, not buried in a
daily log.

## Steps

1. Identify the file (or files) the discussion has been about. If it's
   ambiguous which file the user means, ask before writing anything.
2. In the same directory as that file, create or append to `NOTES.md`.
3. Structure each entry as:

   ```markdown
   ## <filename> — <YYYY-MM-DD>

   **Q: <question the user asked>**
   <concise answer, grounded in the actual code — reference line numbers
   or snippets where useful>

   **Q: <next question>**
   <answer>
   ```

   - If `NOTES.md` already has an entry for this file from earlier today,
     append new Q&A pairs under the existing heading instead of creating a
     duplicate heading.
   - If the user asked several questions about the same file in one
     sitting, batch them into one entry rather than one call per question
     — trigger the write when they say "note this" / "save that" / the
     skill is invoked, not after every single question.
4. Keep answers tied to what was actually asked and actually true of the
   code — don't pad with generic Go documentation. Quote the relevant
   line(s) of code when it makes the answer clearer.
5. Confirm to the user which file's `NOTES.md` was updated.

## Notes

- This is deliberately separate from the `take-notes` skill: `take-notes`
  rolls up a whole session's concepts into a dated `notes/` folder at the
  repo root; `file-notes` keeps granular Q&A pinned next to the specific
  file it explains, so future-you finds it while reading that file again.
- Never overwrite existing `NOTES.md` content — always append.
