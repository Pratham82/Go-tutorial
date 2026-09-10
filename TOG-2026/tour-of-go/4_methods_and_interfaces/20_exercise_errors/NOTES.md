## main.go — 2026-09-11

Exercise: make `Sqrt` return a custom `error` for negative input. Final working version:
```go
type ErrorNegativeSqrt float64

func (e ErrorNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrorNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}
```
Output:
```
1.4142135623730951 <nil>
-2 cannot Sqrt negative number: -2
```

**Q: What went wrong along the way? (iteration notes)**
- `func (e ErrorNegativeSqrt) Error(n float64) string` — adding a parameter breaks the `error` interface, which is exactly `Error() string` (no params). The number must come from the *receiver* `e`, not an argument. `ErrorNegativeSqrt` is a `float64` underneath, so `e` holds the value.
- `fmt.Sprintf("... %v")` with no argument → `go vet` error `format %v reads arg #1, but call has 0 args`, runtime prints `%!v(MISSING)`. Pass `float64(e)`.
- Why `float64(e)` and not just `e`? `e`'s type `ErrorNegativeSqrt` has an `Error()` method, so `fmt` would call it again → infinite recursion. Converting to plain `float64` (no methods) breaks the cycle.
- `return x, error.Error(x)` — `error` is a type/interface, not a value; can't call a method on it. Return an *instance* of your type: `ErrorNegativeSqrt(x)` (a conversion of the float into the named type). That value satisfies `error` via the method.
- `return Sqrt(x)` in the non-negative branch → infinite recursion / stack overflow. That branch must actually compute: `math.Sqrt(x), nil` (or a Newton's-method loop).
- `return math.Sqrt(x), ErrorNegativeSqrt(x)` on the success path → returns a non-nil error when nothing failed. Contract: `nil` error = success. Must be `return math.Sqrt(x), nil`.
- Cosmetic: on the error path, convention returns the zero value (`return 0, ErrorNegativeSqrt(x)`), not `x`, since the caller shouldn't use the float when err != nil.

**Q: When do we use pointers and addresses, and when do we use this technique?**
Go passes everything by value (copies). A pointer is how you opt out of copying and touch the original. `&x` = address of x; `*p` = value at p.

Method receivers (the most common decision):
- Pointer receiver `func (t *T) M()` when: the method mutates the receiver; the struct is large (avoid copy cost); or *any* method on the type needs a pointer (then make them all pointer receivers for consistency).
- Value receiver `func (t T) M()` when the type is small and the method only reads — e.g. this file's `ErrorNegativeSqrt.Error()` (just a float64, read-only).
- When unsure: pointer receiver. It's the default for structs.

Other cases for pointers:
- A function must change a caller's variable → it takes `*T`, you pass `&x`.
- Constructors return `*T` so all callers share one instance (`func NewUser(...) *User`). This is why lesson 19's `run()` did `return &MyError{...}` — and `*MyError` was the type satisfying `error`.
- "Optional / maybe absent" values: `*T` can be `nil`, a plain struct can't.

When you do NOT need pointers:
- Slices, maps, channels — already reference types; a function can modify elements without `&` (appending is the exception — may reallocate, return the new slice).
- Small immutable values (`int`, `string`, `bool`, `time.Time`) — copying is cheap and simpler.
- When isolation is desirable — a copy stops the callee stomping your data.

| Situation | Pointer? |
|---|---|
| Method mutates its receiver | yes `*T` |
| Method only reads, small type | no, value |
| Type has any pointer-receiver method | yes, all methods `*T` |
| Function must change caller's variable | yes, pass `&x` |
| Passing a big struct around | yes, avoid copies |
| Value might be "absent" | yes `*T`, use `nil` |
| Passing a slice/map/channel | no, already a reference |
| Small immutable value | no, value |
