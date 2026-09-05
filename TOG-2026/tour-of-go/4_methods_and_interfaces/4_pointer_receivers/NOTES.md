## main.go — 2026-09-05

**Q: What all changes have been done in this folder?**
Builds on `1_methods` (`Vertex` struct + value-receiver `Abs`, line 12-14)
by adding one new method, `Scale` (line 16-19), which uses a **pointer
receiver** (`*Vertex`) instead of a value receiver — the first pointer
receiver seen so far in the tour.

**Q: What's happening in the Scale function?**
```go
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
```
- Receiver is `*Vertex`, so `v` is a pointer to the caller's actual struct,
  not a copy.
- `v.X = v.X * f` — Go auto-dereferences the pointer (equivalent to
  `(*v).X`), reads the current `X`, multiplies by `f`, and writes the
  result back into the *original* struct through the pointer. Same for
  `v.Y`.
- No return value — it mutates in place instead of returning a new value.

Why a pointer receiver is required here: `Abs` only reads fields and
returns a number, so a value receiver (copy) is fine. `Scale` needs to
actually change `X`/`Y` on the caller's struct — a value receiver would
only scale a throwaway copy, so it must take a pointer to reach the
original.

In `main`: `v.Scale(10)` is called on `v` directly (not `(&v).Scale(10)`)
— Go automatically takes `v`'s address because `v` is addressable and
`Scale` expects a pointer receiver. Result: `v` becomes `{30, 40}`, then
`v.Abs()` = `√(30²+40²)` = `√2500` = `50`.

Rule of thumb noted: once any method on a type needs a pointer receiver,
it's idiomatic to make *all* methods on that type use pointer receivers,
for consistency.
