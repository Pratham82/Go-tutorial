# logs_logs_logs.go — 2026-09-24

## Solution and how it works

**Types:**

- **`Logs`** — `type Logs map[rune]string`, maps a single emoji rune to the
  application name it signals.
- **`logs`** (package-level var) — `Logs{'❗': "recommendation", '🔍': "search", '☀': "weather"}`.
  A `type` declaration alone creates no data — you still need an actual
  variable of that type (or a literal built inline) before you can index
  into it.

**Caveats / tricky parts by function:**

- **`Application`** — iterates the log with `for _, r := range log`, which
  yields `rune`s directly (no manual byte/UTF-8 handling needed). Looks up
  each rune in `logs`; a missing key returns the zero value `""` for maps,
  which doubles as the "not found" check (`if foundRune != ""`).
  - Must return on the **first** match and stop — the tests care about
    reading order (e.g. `"🔍 search recommended product ❗"` → `"search"`,
    not the later `❗`). Overwriting a result variable on every iteration
    without an early `return` silently keeps the *last* match instead of
    the first, which is the wrong behavior here.
  - Falls through to `return "default"` after the loop if nothing matched.

- **`Replace`** — walks the log rune-by-rune and rebuilds the string via
  `strings.Builder`, writing `newRune` in place of `oldRune` and the
  original rune otherwise.
  - Initially used `WriteString(string(r))`, which allocates a throwaway
    string just to hand it to `WriteString`. `strings.Builder.WriteRune(r)`
    writes the rune directly — same result, one less allocation.
  - The if/else can collapse to picking *which* rune to write, then a
    single `WriteRune` call:
    ```go
    r := s
    if s == oldRune {
        r = newRune
    }
    foundRune.WriteRune(r)
    ```

- **`WithinLimit`** — counts runes, not bytes. `len(log)` counts UTF-8
  **bytes**, and multi-byte runes (emoji can be up to 4 bytes) make that
  count diverge from the actual character count the exercise wants.
  Two correct options:
  - Manual: `for range log { count++ }` (rune-by-rune, same idea as the
    other two functions).
  - Stdlib: `utf8.RuneCountInString(log) <= limit` (`unicode/utf8`) — same
    result in one line.

**Testing:** each exercise is its own Go module (separate `go.mod`, no
shared `go.work`), so plain `go test ./...` from the `exercism/go` root
doesn't reach into subfolders. Use `go test -C <exercise-dir> ...` from the
root, or `cd` into the exercise dir and drop `-C`. A `run-tests.sh` wrapper
script lives at `exercism/go/run-tests.sh` (usage:
`./run-tests.sh logs-logs-logs [TestName] [-v]`).
