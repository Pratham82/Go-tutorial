## main.go — 2026-09-05

**Q: Can you explain this file and what's happening in it?**
The file contrasts two ways of scaling a `Vertex`: `Scale` (`main.go:28-31`), a method with a pointer receiver `(v *Vertex)`, and `ScaleFunc` (`main.go:33-36`), a plain function taking `*Vertex`. Both mutate `X` and `Y`, which is why they need a pointer rather than a value receiver/argument — a value would just be a copy, and the mutation would be lost.

In `main` (`main.go:38-47`):
```go
v := Vertex{3, 4}
v.Scale(2)        // v is a value, but Scale has a pointer receiver
ScaleFunc(&v, 10) // must explicitly pass &v

p := &Vertex{4, 3}
p.Scale(3)        // p is already a pointer
ScaleFunc(p, 8)   // p is already a pointer
```

**Q: Please explain in detail.**
The key rule: Go lets you call a pointer-receiver method on an addressable value directly — `v.Scale(2)` is silently rewritten by the compiler to `(&v).Scale(2)`. This is a compile-time syntactic convenience specific to method-call syntax (`x.M()`), not runtime magic.

This convenience does **not** extend to plain functions: `ScaleFunc(v, 10)` is a compile error because function argument passing requires an exact type match — Go won't auto-convert `Vertex` to `*Vertex` there. You must write `ScaleFunc(&v, 10)` explicitly. This asymmetry is exactly what the comment at the top of `main.go` is pointing out.

The auto-`&` trick only works when the value is *addressable* (e.g. a local variable like `v`). It fails on non-addressable values such as a struct literal (`Vertex{3,4}.Scale(2)`) or a map value (`m["a"].Scale(2)`) — not shown in this file, but the natural next thing to try to see the boundary of the rule.

Program output: `v` goes `{3,4}` → `Scale(2)` → `{6,8}` → `ScaleFunc(&v,10)` → `{60,80}`. `p` goes `{4,3}` → `Scale(3)` → `{12,9}` → `ScaleFunc(p,8)` → `{96,72}`. So `fmt.Println(v, p)` prints:
```
{60 80} &{96 72}
```
(`fmt` prefixes pointer-to-struct values with `&` when printing.)
</content>
