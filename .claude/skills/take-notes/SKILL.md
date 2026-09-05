---
name: take-notes
description: Saves a summary of the Go concepts learned in the current session into today's dated notes folder. Use when the user says "take-notes", "take notes", or asks to save/record what was learned today.
---

# Take Notes

Summarize the Go concepts covered in this session and save them to a dated
notes folder at the repo root, so learning progress accumulates day over day.

## Steps

1. Determine today's date in `YYYY-MM-DD` format (use the current date from
   context; do not guess).
2. Ensure `notes/<YYYY-MM-DD>/` exists at the repo root (create it if
   missing).
3. Review the current conversation and any files touched this session
   (e.g. new/edited files under `TOG-2026/`, `learn-go-fast/`,
   `go-projects/`, etc.) to identify the Go concepts, syntax, and exercises
   actually learned or practiced today. Only include what was genuinely
   covered — do not pad with generic Go trivia.
4. Write (or append to, if it already exists) `notes/<YYYY-MM-DD>/notes.md`
   with a concise Markdown summary structured like:

   ```markdown
   # Notes - <YYYY-MM-DD>

   ## Concepts covered
   - <concept>: <one-line explanation / key takeaway>

   ## Code / exercises
   - <path to file worked on> - <what it demonstrates>

   ## Questions / things to revisit
   - <anything the user was unsure about, if applicable>
   ```

   - If the file already exists for today (e.g. take-notes was called
     earlier in the day), merge new content in rather than duplicating
     sections or overwriting prior notes.
   - Keep entries terse — bullet points, not paragraphs. Prioritize the
     "why it matters" / non-obvious bits over restating documentation.
5. Confirm to the user what was saved and where (file path), with a short
   summary of the concepts recorded.

## Notes

- Never invent concepts that weren't actually discussed or practiced in the
  session.
- This is a personal learning log, not documentation — write it in plain,
  practical terms tied to the actual code/examples from the session.
