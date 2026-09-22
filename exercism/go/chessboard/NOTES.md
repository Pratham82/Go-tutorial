# chessboard.go — 2026-09-22

## Solution and how it works

**Types:**

- **`File`** — `type File []bool`, one bool per rank (index 0 = rank 1, ... index 7 = rank 8).
- **`Chessboard`** — `type Chessboard map[string]File`, keyed by file letter `"A"`-`"H"`. Go is case-sensitive, so a stray `ChessBoard`/`Files` typo (capitalization or wrong name) leaves the type referenced in the function signatures undefined — the name has to match exactly everywhere.

**Caveats / tricky parts by function:**

- **`CountInFile`** — straightforward: `cb[file]` gives the `File` slice directly, then range over it counting `true` values.

- **`CountInRank`** — a rank cuts *across* files, so you need to visit each file key (`"A"` through `"H"`) and check one index per file, rather than ranging over a single slice.
  - First pass used a nested loop (`for i, v := range cb[v]`) just to find the element at `i+1 == rank`, an O(8×8) scan for something indexable directly.
  - Optimized to `cb[v][rank-1]` — direct indexing, O(8) total.
  - Direct indexing removes the implicit bounds safety the loop-and-compare version had (an out-of-range rank just never matched before, returning 0 safely). Once indexing directly, an invalid `rank` (e.g. 100) panics with "index out of range" — the test suite includes exactly this case. Fix: guard once *before* the loop, since rank validity is the same for every file: `if rank < 1 || rank > 8 { return 0 }`.

- **`CountAll`** — counts squares present (not occupied), so no need for the explicit `"A"`-`"H"` list or a per-index check. Range over the map directly and sum `len(file)` per entry — order doesn't matter since it's just a sum.

- **`CountOccupied`** — same shape as `CountAll` but counting `true` values instead of `len`. Still needs a nested loop internally (must inspect every bool in every file), but the outer loop can range over `cb` directly rather than building/using an ordered files slice — unlike `CountInRank`, there's no specific index being targeted, so key order is irrelevant.
