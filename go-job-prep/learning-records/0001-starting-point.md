# Starting point: Tour of Go completed through Methods & Interfaces

The learner is a comfortable programmer in another language and has worked through
the Tour of Go up to and including the Methods & Interfaces module (evidence:
sibling repo folders with completed exercises and notes on methods, pointer
indirection, value vs. pointer receivers, and interfaces).

Implications for teaching:
- Syntax, control flow, structs, slices, maps, methods, and basic interfaces can
  be assumed as *exposure* (not yet mastery — watch for evidence).
- The gap toward the backend‑job mission is idiomatic error handling, testing,
  concurrency (`goroutine`/`channel`/`context`), and the `net/http` +
  `encoding/json` standard library.
- Next lesson targets idiomatic **error handling** — it appears in every backend
  codebase, is a common interview probe, and builds directly on the interface
  knowledge just acquired (`error` is just an interface).
