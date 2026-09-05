## main.go — 2026-09-05

This file is the "no pointer" experiment the tour's comment suggests
(remove the `*` from the `Scale` parameter). Compare directly against the
sibling folder `../5_pointers_and_functions/main.go`, which is identical
except `Scale` takes `*Vertex`.

```go
func Scale(v Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(v, 10)
	fmt.Println(Abs(v))  // 5
}
```

**Q: What's happening here, and how does this differ from the pointer version?**
`Scale(v, 10)` passes `v` **by value** — Go copies the entire struct into
a new, separate box that only `Scale` can see. `v.X = v.X * f` inside
`Scale` edits that private copy. When `Scale` returns, the copy is
discarded; `main`'s original `v` was a different piece of memory the whole
time and was never touched. So `Abs(v)` prints `5`, unscaled. Ran both
files with `go run` to confirm: this one prints `5`, the pointer sibling
prints `50`.

**Q: How are pointers taking effect (or in this case, not)?**
No pointer is used anywhere in this file — every value is copied by
default in Go. Mental model: `Vertex` parameter = Go copies the whole
struct into a fresh box = two independent boxes exist (`main`'s and
`Scale`'s) = editing one can never affect the other, no matter what. This
is the baseline behavior that `*Vertex` (in the sibling file) exists to
override.

**Q: (paraphrased) So v is changed but not updated where main is accessing
it, since main only refers to that v — is that right?**
Small correction: `main`'s `v` isn't "changed but not updated" — it's
**never touched at all**. What gets changed is a completely separate copy
that only exists inside `Scale` and is discarded the moment `Scale`
returns. `main`'s `v` and `Scale`'s `v` are two different boxes in memory
from the start, not the same box viewed differently — that's why editing
one has zero effect on the other.
