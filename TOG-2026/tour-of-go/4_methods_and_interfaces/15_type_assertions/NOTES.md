## main.go — 2026-09-10

**Q: How is this working?**
A type assertion `i.(T)` pulls the concrete value back out of an interface by checking the type slot of its `(value, type)` tuple against `T`.

```go
var i interface{} = "hello"   // tuple: ("hello", string)
```

- Line 8 `s := i.(string)` — single-value form, matches → `s == "hello"` (a real string, not an interface). Prints `hello`.
- Line 11 `s, ok := i.(string)` — comma-ok form, matches → `s == "hello"`, `ok == true`. Prints `hello true`. Never panics.
- Line 14 `f, ok := i.(float64)` — comma-ok, mismatch → `f == 0` (zero value of float64), `ok == false`. Prints `0 false`. No panic.
- Line 17 `f = i.(float64)` — single-value form, mismatch → **panic**: `interface conversion: interface {} is string, not float64`. Single-value form has no way to report failure. This line is intentionally there to show the panic.

| form | on match | on mismatch |
|---|---|---|
| `t := i.(T)` | `t` = value | panic |
| `t, ok := i.(T)` | `t` = value, `ok` = true | `t` = zero value, `ok` = false (no panic) |

Use comma-ok when unsure of the type. Same shape as `v, ok := m[key]` for maps. Output stops at the panic:
```
hello
hello true
0 false
panic: interface conversion: interface {} is string, not float64
```

**Q: How to hold an "object" type in an interface? (`var i2 interface{} = {name: string}` doesn't work)**
Go has no bare object literal — every composite value needs a type.

Named struct (normal choice):
```go
type Person struct { Name string }
var i2 interface{} = Person{Name: "Prathamesh"}
p, ok := i2.(Person)
fmt.Println(p.Name, ok)   // Prathamesh true
```

Anonymous struct (closest to `{name: string}`, but verbose — the full `struct{...}` type must be repeated at every assertion, so only for one-off values):
```go
var i2 interface{} = struct{ Name string }{Name: "Prathamesh"}
p, ok := i2.(struct{ Name string })
```

Won't compile:
```go
var i2 interface{} = {Name: "Prathamesh"}        // missing type in composite literal
var i2 interface{} = struct{ Name string }       // that's a type, not a value
```

When the concrete type genuinely isn't known at compile time (e.g. decoded JSON), use `map[string]interface{}`:
```go
var i2 interface{} = map[string]interface{}{"name": "Prathamesh"}
m := i2.(map[string]interface{})
fmt.Println(m["name"])   // Prathamesh
```

**Q: But why did we get panic with `f` (line 17) when line 14 didn't, given both assert float64 on a string?**
The mismatch is identical — the difference is the *form* of the assertion, which decides what Go does on a mismatch.

- Line 14 `f, ok := i.(float64)` — two-result form. You asked for a value **and** a success flag, so Go has somewhere to report failure: `ok = false`, `f = 0` (zero value). No crash needed because the code is clearly prepared for failure.
- Line 17 `f = i.(float64)` — one-result form. No `ok` variable, so Go has **nowhere to report failure**. It can't produce a `float64` that isn't there; silently inventing one would be unsafe, so it panics: `interface conversion: interface {} is string, not float64`.

Not about `f` being special — the single-value form treats "type doesn't match" as fatal by design; comma-ok treats it as recoverable. Rule: know the type → one-result form (panic loudly on a wrong assumption); checking the type → `, ok`.
