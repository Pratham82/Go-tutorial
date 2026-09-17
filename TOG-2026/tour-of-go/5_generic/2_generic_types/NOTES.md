## main.go — 2026-09-17

**Q: The `List[T]` struct has `next`, `val`, and `head int`. Give me suggestions for adding functionality, and how to invoke it in `main`.**
The `head int` field doesn't belong on the node struct — "head" is a property of the list as a whole (which node is first), not of every individual node. It should be dropped. The classic Tour of Go pattern is just:

```go
type List[T any] struct {
	next *List[T]
	val  T
}
```

Suggested methods to add:
- `Push(val T) *List[T]` — prepend a node, O(1). Must return the new head rather than mutate in place, since prepending changes which node is first.
- `Len() int` — walk `next` pointers counting nodes.
- `ToSlice() []T` — collect values in order into a slice.
- `Find(target T, eq func(a, b T) bool) *List[T]` — linear search; needs an `eq` callback because `T any` isn't `comparable`, so `==` can't be used directly. (Alternative: constrain `T` to `comparable` and use `==` inline instead of a callback — worth trying as a variant.)
- `PushBack`, `Remove(val)`, `String()` (`fmt.Stringer`) as further exercises. Note `PushBack` is O(n) since it must walk to the tail, unlike `Push` which is O(1).

Invocation in `main`:
```go
var list *List[int]
list = list.Push(3)
list = list.Push(2)
list = list.Push(1) // list is now: 1 -> 2 -> 3
fmt.Println(list.ToSlice()) // [1 2 3]
```

**Q: Explain the `Push` method line by line.**
```go
func (l *List[T]) Push(val T) *List[T] {
	return &List[T]{next: l, val: val}
}
```
- `l *List[T]` — receiver is a pointer to the current head node (or `nil` if the list is empty).
- `&List[T]{next: l, val: val}` — builds a new node whose `next` points back to the old head `l`, then takes its address with `&`.
- Returns the new node's address; the caller reassigns their variable (`list = list.Push(3)`) so it now points at the new head.
- Works even when `list` is `nil` at the start: calling a method on a nil pointer receiver is legal in Go as long as the method body never dereferences `l` — `Push` only *stores* `l` into `next`, never reads through it. So `Push` on a nil list correctly produces `&List[int]{next: nil, val: 3}`.

**Q: How do pointers work here? What are `*List` and `&List` called when referring to them?**
- `&` is the "address-of" operator — given a value, produces a pointer to it. `&x` reads as "the address of x."
- `*` means two different things depending on position:
  - In a *type* position (`*List[T]`) it means "pointer to `List[T]`" — read as "a pointer to List of T" or "a `List[T]` pointer."
  - In an *expression* position (`*p`) it means "dereference" — follow the pointer to get the value it points to.
- A variable like `l *List[T]` is called "a pointer" or "a pointer variable," not "a list" — it's *a pointer to* a list node.
- `&List[T]{next: l, val: val}` is read as "the address of a new `List[T]`" / "a pointer to a new List."
- Why pointers are required here: nodes reference each other (`next` points to another node); structs can't contain themselves by value (infinite size), so linked structures must use pointers. Using pointers also avoids copying whole node structs around on every call.
- Simple example:
```go
x := 5
p := &x          // p holds the address of x
fmt.Println(*p)  // dereference p → 5
*p = 10          // dereference and assign → changes x itself
fmt.Println(x)   // 10
```
Mental picture: `p` is a sticky note saying "go look over there"; `&x` writes the note, `*p` means "go look at what the note points to." In the list, each node's `next` is literally one of these sticky notes pointing at the next node (or `nil`, meaning "nothing after this").
