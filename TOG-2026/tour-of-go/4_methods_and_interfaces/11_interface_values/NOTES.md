## main.go — 2026-09-10

**Q: How is this working? (walkthrough of the whole file)**
An interface value is a `(value, type)` tuple.
- `var i I` starts as `(nil, nil)`.
- `i = &T{"Hello"}` → tuple is `(&T{"Hello"}, *main.T)`. `describe(i)` prints `(&{Hello}, *main.T)`. `i.M()` runs `(*T).M` → prints `Hello`.
- `i = F(math.Pi)` → same variable `i`, tuple is now `(3.141592653589793, main.F)`. `i.M()` runs `F.M` → prints the float.
- `M` on `T` has a **pointer receiver** (`func (t *T) M()` at line 16), so only `*T` satisfies `I` — that's why the code stores `&T{}`, not `T{}`. `F.M` (line 22) has a **value receiver**, so `F` satisfies `I` directly.
- Static type of `i` is always `I`; the dynamic type in the tuple changes per assignment, and `i.M()` dispatches on that dynamic type.

Output:
```
(&{Hello}, *main.T)
Hello
(3.141592653589793, main.F)
3.141592653589793
```

**Q: How does the describe function work?**
```go
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```
`i` is passed twice, once per verb. `%v` = value in default format (`&{Hello}`, then `3.141592653589793`). `%T` = concrete type (`*main.T`, then `main.F`; `main.` is the package name). `fmt` uses reflection at runtime, so even though the parameter's static type is `I`, `%v`/`%T` read the concrete value and type out of the interface tuple.

**Q: How to define `func (s string) M()` and call it on a string like `n = "Prathamesh"`?**
Can't attach a method to `string` directly — Go only allows methods on types declared in the same package, and `string` is a builtin (`cannot define new methods on non-local type string`). Use a named type, same as `type F float64` in this file:
```go
type S string

func (s S) M() {
	fmt.Println(s)
}
```
Then `i = S("Prathamesh"); i.M()` → prints `Prathamesh`; `describe(i)` → `(Prathamesh, main.S)`. `S("...")` is a conversion (string literal defaults to `string`). Value receiver, so both `S` and `*S` satisfy `I`. `S` is distinct from `string` but still supports `len`, `+`, indexing since its underlying type is `string`.

**Q: Isn't `n = S("Prathamesh"); n.M()` valid?**
Only if `n` is already declared, because `=` is assignment not declaration. If `n` was never declared → `undefined: n`; use `n := S("Prathamesh")`. Valid if `var n I` earlier (S satisfies I) or `var n S` earlier. The snippet on its own is missing a declaration.
