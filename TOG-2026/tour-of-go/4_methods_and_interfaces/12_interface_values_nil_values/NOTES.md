## main.go — 2026-09-10

**Q: How is this working? (walkthrough of the whole file)**
This example contrasts a *nil interface* with an *interface holding a nil pointer*.

Setup in `main`:
```go
var i I     // i == (nil, nil)  → truly nil interface
var t *T    // t == nil, static type *T
```

Round 1:
```go
i = t       // i == (nil, *main.T)  → type slot filled, value slot nil
describe(i) // (<nil>, *main.T)
i.M()       // <nil>   — no panic
```
`i != nil` here, because an interface is nil only when *both* value and type are nil. `i.M()` works because dispatch only needs the type slot (`*main.T`); `(*T).M` has a pointer receiver so it runs with `t == nil`, and the guard at lines 14-17 (`if t == nil { fmt.Println("<nil>"); return }`) handles it.

Round 2:
```go
i = &T{"hello"}  // i == (&T{"hello"}, *main.T)
describe(i)      // (&{hello}, *main.T)
i.M()            // hello   — t non-nil, prints t.S
```

Output:
```
(<nil>, *main.T)
<nil>
(&{hello}, *main.T)
hello
```

Takeaways: (1) a method can be called on a nil receiver — write methods that handle nil gracefully; (2) an interface holding a nil concrete value is itself NOT nil. A *truly* nil interface (`var i I` with nothing assigned) calling `i.M()` would panic — no type in the tuple to dispatch on. This is also the classic gotcha: returning a nil `*T` as an `error` makes the caller's `err != nil` check pass.

**Q: Explain `var i I` vs `var t *T` again.**
Both are "nothing", but different kinds.

`var i I` — `I` is an interface type. Zero value is the empty tuple `(value: nil, type: nil)`. Nothing put in it. `i == nil` is true. Calling `i.M()` now panics — no type slot to dispatch on.

`var t *T` — `*T` is a concrete type. Zero value of a pointer is `nil`, so `t == nil` is true. But `t` is statically typed `*T` at compile time; it's just a pointer pointing at nothing, not a tuple.

`i = t` copies both parts of the concrete value into the tuple → `i == (nil, *main.T)`. Value slot nil, type slot now filled. So `i == nil` becomes **false**.

| Expression | value slot | type slot | `== nil`? |
|---|---|---|---|
| `var i I` | nil | nil | true |
| `var t *T` (pointer itself) | — | `*T` (static) | true |
| `i = t`, then check `i` | nil | `*main.T` | **false** |

That's why `i.M()` works: type slot has `*main.T`, Go finds `(*T).M`, calls it with the nil value as receiver `t`, and the `if t == nil` guard handles it.
