## main.go — 2026-09-11

Context: file has the tour's original `MyError` plus a user-added `CustomError` with matching `Error()` method and a `runErrorLog()` returning `&CustomError{...}`.

**Q: Is `Error` a keyword in Go?**
No. `Error` is an ordinary method name — you could call it anything. It matters *only* because the standard library defines:
```go
type error interface {
	Error() string
}
```
Read that as a checklist: "to be an `error`, a type must have a method named `Error`, taking nothing, returning `string`." Any type that has that method automatically satisfies `error` — no `implements` keyword, same structural matching as `Stringer`/`String()`.

Proof it's not a keyword: rename the method to `Describe()` and the build fails with `*CustomError does not implement error (missing method Error)`. Rename back and it compiles. The name *is* the contract.

Chain: you name the method `Error` → that makes your struct fit the `error` interface → functions declaring `error` as return type can return it → `fmt.Println` finds the `Error()` method and calls it.

| Interface | Required method | Called by |
|---|---|---|
| `fmt.Stringer` | `String() string` | `fmt` when printing |
| `error` | `Error() string` | `fmt` when printing; `if err != nil` by convention |

**Q: Explain the working of `MyError` and `CustomError`.**
Both do the same job (carry failure info + satisfy `error`), just different fields. They're independent types — nothing links them; each meets the interface on its own.

`MyError`: struct with `When time.Time`, `What string`. `func (e *MyError) Error() string` returns `fmt.Sprintf("at %v, %s", e.When, e.What)`. `run()` returns `&MyError{time.Now(), "it didn't work"}` — fields filled *positionally* in declaration order.

`CustomError`: struct with `ErrorName, ErrorCode, ErrorTitle, ErrorSubtitle`. `func (e *CustomError) Error() string` prints them in the order Code, Name, Subtitle, Title (Sprintf arg order is the author's choice). `runErrorLog()` returns `&CustomError{"Error 1", 404, "Not found", "Value not found"}` — positional, so literal order is *declaration* order (Name, Code, Title, Subtitle) even though `Error()` prints a different order. Use field-keyed literals (`ErrorCode: 404, ...`) to avoid the mismatch.

`main` handles both identically: `if err := fn(); err != nil { fmt.Println(err) }` — call, nil-check for failure, `fmt.Println` calls `err.Error()` to render.

**Q: What is the `(e *MyError)` part before the method name?**
It's the *receiver* — what turns a plain function into a method attached to a type.
```go
func (e *MyError) Error() string {
//   │  │         │       └ return type
//   │  │         └ method name
//   │  └ receiver TYPE (*MyError = pointer to MyError)
//   └ receiver NAME (the variable used inside the body)
```
Basically a special first parameter moved in front of the name, which is what enables dot-call syntax `x.Error()` instead of `Error(x)`. When you call `err.Error()`, Go passes `err` in as `e`, so inside the body `e` *is* that specific value; `e.When` / `e.What` read its fields. Same idea as `this`/`self`, except you name it and type it explicitly. Convention: a short letter.

Value vs pointer receiver:
- `func (e MyError) Error() string` — `e` is a *copy* of the struct.
- `func (e *MyError) Error() string` — `e` points at the *original*; no copy.

**Q: Why `&MyError{...}` / `&CustomError{...}` in the runner functions?**
Because the `Error()` method has a *pointer receiver* (`*MyError`), so only `*MyError` is in the method set that satisfies `error` — plain `MyError` does not.

| expression | type | satisfies `error`? |
|---|---|---|
| `MyError{...}` | `MyError` | ✗ |
| `&MyError{...}` | `*MyError` | ✓ |

`run()` returns `error`, so it must return something satisfying the interface → needs `&`. `&MyError{...}` = construct the struct, then take its address (Go heap-allocates automatically since the pointer escapes via return).

Drop the `&` and the build fails: `MyError does not implement error (method Error has pointer receiver)`.

If the method used a *value* receiver (`func (e MyError) Error() string`), both `MyError{...}` and `&MyError{...}` would satisfy `error` and the `&` would be optional. Go convention for errors is pointer receivers, hence `return &SomeError{...}` almost everywhere.
