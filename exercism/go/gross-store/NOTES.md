# gross_store.go — 2026-09-21

## Solution and how it works

**Caveats / tricky parts by function:**

- **`Units()`** — every value must be a distinct multiple (`quarter_of_a_dozen=3`, `half_of_a_dozen=6`, `dozen=12`, `small_gross=120`, `gross=144`, `great_gross=1728`). Easy to accidentally copy-paste the same value for every key.

- **`AddItem`** — no need for an `if itemCount > 0 {...} else {...}` branch. A missing map key returns the zero value (`0` for `int`), so `bill[item] += unitValue` alone correctly handles both "new item" and "existing item" cases.

- **`RemoveItem`** — the trickiest one. Order of operations matters:
  - Compute `newCount := itemCount - unitValue` as a **local variable first**, before touching `bill`. If you write to `bill[item]` before validating, an invalid removal (result < 0) leaves the map mutated even though you return `false`.
  - Three outcomes on `newCount`: negative → return `false` without mutating `bill`; exactly zero → `delete(bill, item)` (don't leave a `0` entry sitting in the map); positive → write `bill[item] = newCount`.

- **`GetItem`** — the two-value map lookup `qty, ok := bill[item]` already returns exactly what the function signature wants. No need for an `if !ok {...}` branch that just returns the same `qty, ok` either way — `qty` is already `0` when `ok` is `false`, so `return qty, ok` alone is equivalent and simpler.
