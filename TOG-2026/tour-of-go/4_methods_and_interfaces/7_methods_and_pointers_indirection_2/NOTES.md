## main.go — 2026-09-06

**Q: Whats happening in this file?**
This is the mirror image of the pointer-receiver example (`6_methods_and_pointer_indirection`) — same auto-conversion rule, opposite direction.

`Abs` (`main.go:12-14`) is a method with a **value receiver** `(v Vertex)`, and `AbsFunc` (`main.go:16-18`) is a plain function taking a **value** `Vertex`. Both just compute `√(X²+Y²)` — read-only, no mutation needed, so a value receiver/argument is the natural choice.

In `main` (`main.go:20-28`):
```go
v := Vertex{3, 4}
fmt.Println(v.Abs())     // v is already a value — direct match
fmt.Println(AbsFunc(v))  // same, direct match

p := &Vertex{4, 3}
fmt.Println(p.Abs())      // p is *Vertex, but Abs has a value receiver
fmt.Println(AbsFunc(*p))  // must explicitly dereference: *p
```

The interesting line is `p.Abs()`. `p` is `*Vertex`, and `Abs` is defined on `Vertex` — yet it compiles, because Go rewrites `p.Abs()` as `(*p).Abs()` automatically, dereferencing the pointer for you at the method-call site. `AbsFunc(*p)` needs the dereference written explicitly — `AbsFunc(p)` would be a compile error — because plain functions never auto-convert between value and pointer types, only methods do.

Combined with the pointer-receiver example, the full symmetry is:
| Receiver/param type | Called with value `v` | Called with pointer `p` |
|---|---|---|
| Pointer receiver method | `v.Scale(2)` → `(&v).Scale(2)` (if `v` addressable) | `p.Scale(3)` direct |
| Value receiver method | `v.Abs()` direct | `p.Abs()` → `(*p).Abs()` |
| Function taking pointer | `ScaleFunc(&v, ...)` — must write `&v` | `ScaleFunc(p, ...)` direct |
| Function taking value | `AbsFunc(v)` direct | `AbsFunc(*p)` — must write `*p` |

Output: `v.Abs()`, `AbsFunc(v)`, `p.Abs()`, and `AbsFunc(*p)` all compute the 3-4-5 triangle hypotenuse, so the program prints:
```
5
5
5
5
```
</content>
