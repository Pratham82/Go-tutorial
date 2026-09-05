## main.go — 2026-09-05

This file uses a **pointer receiver-style function** (`*Vertex`). Compare
directly against the sibling folder `../5_pointers_and_functions_without_pointers/main.go`,
which is identical except `Scale` takes `Vertex` by value.

```go
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))  // 50
}
```

**Q: What's happening in the Scale function, and how do the two folders differ?**
`&v` passes the address of `main`'s struct (not a copy of the struct
itself). `Scale`'s parameter `v` is a pointer holding that address — there
is only **one** `Vertex` in memory here. `v.X = v.X * f` auto-dereferences
the pointer (shorthand for `(*v).X = (*v).X * f`) and writes through it,
directly editing `main`'s box. So after `Scale` returns, `main`'s `v` is
genuinely `{30, 40}`, and `Abs(v)` = `√(30²+40²)` = `50`.

Contrast: in the `_without_pointers` sibling, `Scale(v Vertex, f float64)`
receives a full **copy** of the struct in a separate memory box. Editing
that copy never touches `main`'s original, so `main`'s `v` stays `{3, 4}`
and `Abs(v)` = `5`. Ran both files with `go run` to confirm: this one
prints `50`, the sibling prints `5`.

**Q: How are the pointers taking effect here?**
Mental model: a variable is a named box in memory. `Vertex` parameter =
Go copies the whole box → function gets its own private box → edits can't
reach the original. `*Vertex` parameter = Go copies only an *address* →
no second box is created → the function's `v` and `main`'s `v` are two
different names pointing at the **same** box → an edit "through" either
name changes the one box both are looking at. It's not about scope — it's
about whether there are one or two boxes involved.

**Q: (paraphrased) So without the pointer, main's v is changed but not
updated where main is accessing it — and with the pointer we always get
the latest value because v is updated inside Scale, regardless of scope?**
Two small corrections landed on this:
1. Without the pointer, `main`'s `v` isn't "changed but not updated" — it
   is **never touched at all**. The struct that gets edited is a separate,
   throwaway copy that dies when `Scale` returns; `main`'s box was a
   different piece of memory the whole time.
2. It's not about *scope* (where a name is visible) — it's about
   **identity**: same memory address (pointer) vs. two different
   addresses (value copy). With a pointer, `Scale` and `main` share the
   exact same box, so an edit through either name is visible to both —
   not because of scope rules, but because they're literally the same box.
