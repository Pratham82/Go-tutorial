# Interfaces are implemented implicitly

A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

## main.go — 2026-09-10

**Q: What's the issue in this function definition — `greet(name string) string { return name }` inside `type I2 interface{...}`?**
An interface can only declare **method signatures** — it's a contract, not an implementation. Putting a body (`{ return name }`) directly inside the interface declaration is a syntax error: `syntax error: unexpected {, expected semicolon or newline in interface type`. Fix: drop the body, keep just the signature (`greet(name string) string`), and implement it separately on a concrete type via a receiver.

**Q: In this code why do we have a `Foo` struct? And also what's a struct?**
A struct groups related data into a named type (`type Name struct { fields }`). `Foo struct{}` (`main.go:25`) is an *empty* struct — zero fields — used purely as a label/vehicle to attach a method to.

`Foo` exists because interfaces can't hold implementations (per the question above) — some concrete type has to carry the real code. `func (f Foo) greet(...)` (`main.go:27-30`) attaches the actual implementation to `Foo`. `Foo` needs no fields because `greet` doesn't need any state — it just echoes back its `name` argument. Same pattern as `MyFloat`/`Vertex` from file 9, just with a type that has no data.

**Q: What's the function name here — `func (f Foo) greet(name string) string { return name }`?**
The method name is `greet`. `(f Foo)` is the receiver (`f` is the receiver's local var name, `Foo` its type) — the receiver is what makes `greet` a *method* on `Foo` rather than a standalone function. Its full identity is `Foo.greet`; called as `f.greet("Ann")`, never as bare `greet("Ann")`.

**Q: How to call this in main — is `foo.greet(myName)` correct?**
Yes. `var foo = Foo{}` then `foo.greet(myName)` is correct — `greet` is a method on `Foo`, so calling it on a `Foo` value works. Since `greet` prints `name` and returns it, this prints `Prathamesh` (the return value is ignored, which is fine).

One refinement suggested and since adopted in the code: since `i` (`main.go:32`) is declared through the interface type (`var i I = T{"hello"}`), declaring `foo` the same way — `var foo I2 = Foo{}` (`main.go:35`) — makes both examples symmetric, calling `greet` *through* the `I2` interface rather than directly on the concrete `Foo` type. Same output either way.

**Q: Can you add a comment explaining why we need that invocation for calling greet/foo?**
Added above `foo.greet(myName)` (`main.go:38`): a comment noting that `greet` must be called on `foo` (a receiver) because it's a method, not a standalone function — there's no bare `greet(myName)` to call.
