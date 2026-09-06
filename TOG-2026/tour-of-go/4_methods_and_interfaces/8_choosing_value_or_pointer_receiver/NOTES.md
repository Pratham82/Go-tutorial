# Choosing a value or pointer receiver

There are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

In this example, both Scale and Abs are methods with receiver type *Vertex, even though the Abs method needn't modify its receiver.

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)

## main.go — 2026-09-06

**Q: Check this file and explain its working and what's happening here.**
This example moves past the value-vs-pointer mechanics (examples 6/7) to the actual design guidance: when should you choose a pointer receiver?

Both methods use pointer receivers:
- `Scale` (`main.go:12-15`) `(v *Vertex)` — mutates `X`/`Y`, so it *needs* a pointer receiver.
- `Abs` (`main.go:17-19`) `(v *Vertex)` — only reads `X`/`Y`, so it doesn't strictly need a pointer receiver (example 7 showed it working fine as a value receiver). It's given one anyway.

The point: once a type has one method with a pointer receiver, its other methods should generally use pointer receivers too, for consistency — even if a given method doesn't itself need to mutate anything. Mixing receiver types on the same type is legal but bad style, since it makes the type's method set inconsistent (a value receiver silently works with a copy while a pointer receiver on the same type works with the original).

In `main` (`main.go:21-27`):
```go
v := &Vertex{3, 4}                                    // v is already *Vertex
fmt.Printf("Before scaling: %+v, Abs %v\n", v, v.Abs())
v.Scale(5)
fmt.Printf("After scaling scaling: %+v, Abs %v\n", v, v.Abs())
```
Since `v` is declared as `&Vertex{3, 4}`, it's already a `*Vertex` — both `v.Abs()` and `v.Scale(5)` are direct calls, no indirection tricks needed (unlike examples 6/7, this file sidesteps that question by making `v` a pointer from the start).

`v.Scale(5)` multiplies `X` and `Y` by 5 in place (`{3,4}` → `{15,20}`), and `Abs()` recomputes the magnitude each time it's called: before scaling `√(9+16)=5`, after scaling `√(225+400)=25`.

Output:
```
Before scaling: &{X:3 Y:4}, Abs 5
After scaling scaling: &{X:15 Y:20}, Abs 25
```
(`%+v` on a pointer to struct prints field names, with `&` prefix for the pointer.)

**Takeaway:** the Go rule of thumb — if *any* method on a type needs a pointer receiver (usually because it mutates the receiver), give *all* its methods pointer receivers, for a uniform and predictable API.
