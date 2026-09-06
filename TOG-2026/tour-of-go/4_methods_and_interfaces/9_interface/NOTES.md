# Interfaces

An interface type is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.

Note: There is an error in the example code on line 22. Vertex (the value type) doesn't implement Abser because the Abs method is defined only on *Vertex (the pointer type).

## main.go — 2026-09-06

**Q: Can you explain what's happening in this file?**
This file introduces **interfaces**, and it directly builds on files 6-8 about value vs. pointer receivers — it contains a deliberate compile error to make that connection explicit.

```go
type Abser interface {
	Abs() float64
}
```
`Abser` is an interface: a method-set requirement, not an implementation. Any type with an `Abs() float64` method automatically satisfies it — Go's satisfaction is implicit/structural, no `implements` keyword.

Two implementers:
- `MyFloat` (`main.go:27-35`) implements `Abs` with a **value receiver** → the method set of `MyFloat` includes `Abs`.
- `Vertex` (`main.go:37-43`) implements `Abs` with a **pointer receiver** → the method set of `Vertex` does *not* include `Abs`; only `*Vertex`'s method set does.

In `main` (`main.go:12-25`):
```go
var a Abser
f := MyFloat(-math.Sqrt2)
v := Vertex{3, 4}

a = f  // OK — MyFloat's method set has Abs
a = &v // OK — *Vertex's method set has Abs

a = v  // COMPILE ERROR — commented in the code as intentional
```
Line 22 (`a = v`) is the whole point of the exercise. `v` is a `Vertex` value, not `*Vertex`. Since `Abs` is defined only on `*Vertex`, plain `Vertex` doesn't satisfy `Abser`, so this fails to compile:
```
cannot use v (variable of type Vertex) as Abser value in assignment:
Vertex does not implement Abser (method Abs has pointer receiver)
```
Key rule: **interface satisfaction checks the method set of the exact static type being assigned** — unlike a direct method call `v.Abs()` (which *would* auto-convert to `(&v).Abs()` per file 7's rule), interface assignment is a type-conformance check, not a method call on a known addressable variable, so the auto-`&` convenience does not apply here. Fix would be `a = &v`.

Separately, `Vertex.Abs()` (`main.go:41-43`) has a bug — `math.Sqrt(v.X + v.X + v.Y*v.Y)` adds `v.X + v.X` (`2*v.X`) instead of squaring it (`v.X*v.X`), unlike the correct version in files 7/8. For `Vertex{3,4}` this would give `√22 ≈ 4.69` instead of the correct `5`. It never actually executes though, since line 22 doesn't compile as written.

**Q: How does interfaces work in go and why are we getting an error when we are doing a = v?**
An interface (`type Abser interface { Abs() float64 }`) is just a method-signature contract. A concrete type satisfies it automatically (structural typing, no `implements` keyword) if that type's **method set** contains every method the interface requires.

Method sets depend on receiver kind:
- Value receiver `func (t T) M()` → `M` is in the method set of both `T` and `*T`.
- Pointer receiver `func (t *T) M()` → `M` is in the method set of `*T` **only**, not `T`.

Here `MyFloat` uses a value receiver (`main.go:29-35`) so `MyFloat`'s method set has `Abs`. `Vertex` uses a pointer receiver (`main.go:41-43`) so only `*Vertex`'s method set has `Abs` — plain `Vertex`'s does not.

`a = v` fails because assigning to an interface variable is a **type-conformance check against the value's static type** — the compiler checks whether `Vertex` (not `*Vertex`) has `Abs` in its method set, finds it doesn't, and rejects it. This differs from a direct call like `v.Abs()`, which the compiler is free to rewrite as `(&v).Abs()` per the method-call auto-addressing rule (file 7) — that rewrite only applies to call syntax, not to interface assignment, since assignment isn't a method call at all. Fix: `a = &v`.

**Q: What's *Vertex called, and what's &v called — how do you pronounce them, and what's their usage?**
- `*Vertex` is a **pointer type** ("pointer to Vertex"), used wherever a type is expected (receiver declarations, params, variable types) — e.g. `func (v *Vertex) Abs() float64` (`main.go:41`) declares the receiver's type as "pointer to Vertex."
- `&v` is the **address-of** expression ("address of v"), applied to a value to produce a pointer to it — e.g. `a = &v` (`main.go:18`) produces a `*Vertex` value pointing at `v`.
- `*` has a second meaning on an expression (not a type): dereference — `*p` means "the value `p` points to," the inverse of `&`.
- Relationship: `p := &v` (value → pointer) and `q := *p` (pointer → value) are inverse operations; `*Vertex` as a type name is a separate (third) usage of `*`, purely declarative.
- Why it matters here: `v` alone (type `Vertex`) doesn't satisfy `Abser`; `&v` (type `*Vertex`) does, because `&` converts the value into the pointer type that actually carries `Abs` in its method set.

**Q: So what are we exactly trying to learn here?**
Two layers:
1. **Big picture:** interfaces let unrelated concrete types (`MyFloat`, a number; `Vertex`, a struct) be used interchangeably as long as they share a method — this is Go's form of polymorphism (no inheritance, just "has the right methods"). It's the same mechanism behind `io.Reader`, `error`, `sort.Interface`, etc.
2. **Specific lesson:** interface satisfaction is checked via **method sets**, and method sets depend on receiver kind (value vs. pointer) — directly tying back to files 6-8. The deliberate `a = v` compile error exists to show that this isn't just a style choice: if a type's methods use pointer receivers, you must pass `*T` (not `T`) to satisfy interfaces with it.

One-sentence takeaway: interfaces are satisfied structurally by method sets, and method sets depend on receiver type — so choosing pointer vs. value receivers determines what your type can be used as.

**Q: What's `a` in here exactly, what is its purpose?**
`a` (`main.go:13`, `var a Abser`) is a variable of **interface type**. Unlike a concrete-type variable, it's a "box" holding two things: the concrete type currently stored (e.g. `MyFloat` or `*Vertex`) and the value of that type. Its declared type (`Abser`) never changes, but what's inside it can be reassigned to any type satisfying `Abser`.

Its purpose in this file is purely demonstrative — a test bed for interface-satisfaction rules: `a = f` succeeds (`MyFloat` satisfies `Abser`), `a = &v` succeeds (`*Vertex` satisfies `Abser`), `a = v` fails (`Vertex` doesn't). There's no real business logic riding on `a`; it exists to show which assignments compile. `fmt.Println(a.Abs())` (`main.go:23`) is the payoff line — it calls `.Abs()` on `a` without needing to know which concrete type is actually inside it (this line never executes though, since the file has a compile error above it on line 22).
