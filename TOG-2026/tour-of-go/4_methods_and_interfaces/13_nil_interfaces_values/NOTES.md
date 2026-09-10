## main.go — 2026-09-10

**Q: Check the file and explain the working.**
This example deliberately panics — it demonstrates a *truly nil interface*.

```go
var i I        // i == (nil, nil) — no value, no type
describe(i)    // prints (<nil>, <nil>)
i.M()          // panic
```

- Line 11 `describe(i)`: works. `%v`/`%T` just report the tuple — both slots empty → `(<nil>, <nil>)`.
- Line 12 `i.M()`: panics with `invalid memory address or nil pointer dereference`. Method dispatch on an interface needs the **type slot** to know which concrete `M` to call. Type slot is nil → nothing to dispatch to → runtime crash.

Actual run output:
```
(<nil>, <nil>)
panic: runtime error: invalid memory address or nil pointer dereference
...
main.main()
	.../13_nil_interfaces_values/main.go:12
exit status 2
```

Contrast with `12_interface_values_nil_values`:

| | type slot | `i.M()` result |
|---|---|---|
| `12`: `var t *T; i = t` | `*main.T` | works — dispatches to `(*T).M`, nil receiver handled by guard |
| `13`: `var i I` (nothing assigned) | nil | panics — no type to dispatch on |

Rule: an interface holding a nil *pointer* is still callable (it has a type); a nil *interface* is not callable at all.
