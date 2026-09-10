## main.go — 2026-09-10

**Q: What's happening here?**
Demonstrates the *empty interface* `interface{}` — an interface listing zero methods, so every type satisfies it. It's the "holds anything" type.

```go
var i interface{}   // tuple (nil, nil)
describe(i)          // (<nil>, <nil>)

i = 42               // (42, int)
describe(i)          // (42, int)

i = "hello"          // ("hello", string)
describe(i)          // (hello, string)
```

Output:
```
(<nil>, <nil>)
(42, int)
(hello, string)
```

Same variable `i`, three different concrete types over its life (nil → int → string). Each assignment repopulates the `(value, type)` tuple — same mechanics as the earlier interface examples, just with no method-set restriction on what can go in. `describe(i interface{})` (line 18) accepts any argument; `%v` prints the value, `%T` the dynamic type via reflection.

From INFO.md: `any` is a built-in alias for `interface{}` (Go 1.18+), fully interchangeable. This is how `fmt.Println`/`Printf`, JSON decoding etc. accept arbitrary values (`...any`). To do type-specific work with the value inside, you need a type assertion (`i.(int)`) or type switch — next lesson.
