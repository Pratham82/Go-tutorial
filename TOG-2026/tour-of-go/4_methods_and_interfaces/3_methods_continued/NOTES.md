# Methods on non-struct types

This step's point: **methods aren't limited to structs — you can declare a
method on any named type you define**, as long as it's defined in the same
package.

```go
type MyFloat float64
```

This creates a brand-new named type `MyFloat`, whose underlying
representation is `float64`. It's a distinct type from plain `float64` even
though the data is identical — you need an explicit conversion to go
between them.

```go
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
```

- The receiver `(f MyFloat)` attaches `Abs` to the custom type. This works
  even though `MyFloat` isn't a struct — it's just a number under the hood.
- Inside, `f` behaves like a normal `MyFloat`/float value: `f < 0` compares
  it, `-f` negates it.
- `float64(-f)` / `float64(f)` — explicit conversions back to plain
  `float64`, needed because the function's return type is `float64`, not
  `MyFloat`. Even though they hold the same bits, Go won't implicitly
  convert between named types.

```go
func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
```

- `math.Sqrt2` is a plain `float64` constant (~1.41421356...). `-math.Sqrt2`
  negates it, then `MyFloat(...)` converts that `float64` into `MyFloat`.
- `f.Abs()` calls the method: since `f` is negative, it returns
  `float64(-f)` → the absolute value, `1.41421356...`.

## Key takeaway vs. previous steps

- Steps 1-2 (`1_methods`, `2_methods_are_functions`) showed methods on a
  **struct** (`Vertex`).
- This step shows methods work on **any named type**, including one based
  on a primitive like `float64`.
- Hard rule: the method's receiver type must be defined in the **same
  package** as the method — you can't add methods to types from other
  packages (e.g. you can't do `func (f float64) Abs() float64` directly on
  the builtin `float64`, which is exactly why this exercise wraps it in
  `MyFloat` first).

## main.go — 2026-09-05

**Q: Can you explain what's happening here in main.go?**
Covered above: `MyFloat` is a named type over `float64` (line 8), `Abs`
(lines 10-16) is a method on that named type showing methods work on any
named type, not just structs, and `main` (lines 18-21) converts
`-math.Sqrt2` into a `MyFloat` and calls `.Abs()` on it, returning
`1.41421356...`.
