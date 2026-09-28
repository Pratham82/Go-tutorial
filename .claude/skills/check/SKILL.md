---
name: check
description: Reviews a partially solved exercism Go exercise up to a given function, giving hints and nudges instead of solutions. Use when the user says "check", "check my solution", "check till <func>", or asks whether their exercism work so far is correct.
---

# Check

Review the user's in-progress exercism exercise and nudge them toward fixes
without handing over the answer.

## Steps

1. **Find the exercise.** Use the file or directory the user names. Otherwise,
   pick the most recently modified exercise under `exercism/go/`.
2. **Scope the review.** If a function name is given as an argument, treat that
   function and every function above it in the file as "done". Without an
   argument, treat every function that is not still
   `panic("Please implement …")` as done. Ignore the unimplemented ones.
3. **Run the tests** for only the done functions:
   `go test -run '<TestFuncA|TestFuncB|...>'` from the exercise directory.
   Read the relevant cases in `*_test.go` to understand what each failure is
   actually testing (e.g. helper types the tests define).
4. **Report per function:** pass or fail. For each failure:
   - Show the failing test output.
   - Give a hint that points to the concept: the Go feature that's missing,
     or the edge case the test covers. Phrase it as a question or a pointer
     ("how does Go let you ask what's inside an interface?").
   - Offer a second, more specific hint if they want it. Do **not** write the
     fixing code.
5. **Style and idiom** notes: brief, framed as questions, never rewrites.
   Don't nitpick passing code unless something is genuinely non-idiomatic.
6. Give full code **only** if the user explicitly asks for the solution.
