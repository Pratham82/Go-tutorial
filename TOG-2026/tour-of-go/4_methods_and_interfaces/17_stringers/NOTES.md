## main.go — 2026-09-10

**Q: How is this working?**
Demonstrates `fmt.Stringer`:
```go
type Stringer interface {
	String() string
}
```
Any type with a `String() string` method satisfies it (implicitly — no `implements` keyword). `fmt` checks for this interface whenever it renders a value (`Println`, `Printf` `%v`/`%s`, etc.).

In this file:
```go
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
```
`Person` now implements `Stringer` (value receiver → both `Person` and `*Person` qualify).

`main`:
```go
a := Person{1, "Arthur Dent", 42}          // positional fields: id=1, Name="Arthur Dent", Age=42
z := Person{2, "Zaphod Beeblebrox", 9001}
fmt.Println(a, z)
```
`fmt.Println` does an internal type assertion like `if s, ok := arg.(Stringer); ok { use s.String() }`. Both match, so it calls `String()` instead of default struct formatting.

Output:
```
Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)
```

Without `String()` you'd get the default dump: `{1 Arthur Dent 42} {2 Zaphod Beeblebrox 9001}`. Note `id` is set but never printed — `String()` only uses `Name` and `Age`. Same idea as JS `toString()` / Python `__str__`, resolved structurally via the interface, not by inheritance.
