## census.go — 2026-09-30

**Q: How do I "map over" a `map[string]string` in Go?**
Go has no `map()` function — use `for ... range`:
```go
for k, v := range m { ... }   // key + value
for k := range m { ... }      // keys only
for _, v := range m { ... }   // values only
```
To "transform", build a new map/slice inside the loop. Iteration order is
random; sort keys (`slices.Sorted(maps.Keys(m))`) if order matters. `v` is a
copy — assign `m[k] = ...` to change the map.

**Q: Why didn't `for k,v := r.Address { k=="street" && v == 0 { ... } }` compile?**
Three mistakes:
- missing `range` → `for k, v := range r.Address`
- missing `if` before the condition
- `v == 0` compares a `string` to an `int` → use `v == ""`

The logic was also incomplete: if `"street"` isn't in the map at all, the loop
never returns `false`.

**Q: How do I write `if !r.Address["street"] { }`?**
`!` only works on `bool`. Go has no truthy/falsy values. Two options:
```go
if r.Address["street"] == "" { }      // missing OR empty
street, ok := r.Address["street"]     // "comma ok": did the key exist?
if !ok { }
```
A lookup of a missing key (even on a `nil` map) returns the zero value `""`.
That's why line 40 can be a single expression:
```go
return r.Name != "" && r.Address["street"] != ""
```
It covers the nil-map, empty-map, unknown-key and empty-street test cases all at once.
Age is optional (test "age is optional"), so there's no age check.

**Q: What's the zero value of a map?**
`nil`. Reading from a nil map, `len()` and `range` over it all work. Writing
to it panics (`assignment to entry in nil map`). `map[string]string{}` is
empty but **not** nil. The `Delete` test checks `Address != nil`, so an
empty map would fail it.

**Q: Why did `Delete` fail with `&Resident{...}` / `*r = &Resident{...}`?**
- `&Resident{...}` on its own line creates a *new* value and discards it. It
  doesn't touch the resident `r` points to.
- `*r = &Resident{...}` is a type mismatch: `*r` is a `Resident`, but
  `&Resident{...}` is a `*Resident`.
- Correct form (line 51): `*r = Resident{}` replaces the value behind the pointer.
- Multi-line composite literals need a trailing comma after the last field.

**Q: How does `*r = Resident{}` work without giving default values?**
Go never leaves memory uninitialized. Any field you don't set gets its type's
**zero value**:

| Type | Zero value |
|---|---|
| `string` | `""` |
| `int`, `float64` | `0` |
| `bool` | `false` |
| map, slice, pointer, chan, func, interface | `nil` |
| struct | each field at its zero value |

So `Resident{}` == `Resident{Name: "", Age: 0, Address: nil}`. It's also
future-proof: new fields get reset automatically.

**Q: In `Count`, why `resident.HasRequiredInfo()` and not `&resident...` or `*resident...`?**
On line 57, `residents` is a `[]*Resident`, so `resident` is already a `*Resident`.
That matches the pointer receiver `func (r *Resident)`, so no conversion is needed.

Go auto-converts for method calls anyway:

| You have | Receiver | Go does |
|---|---|---|
| `*T` | `*T` | nothing |
| `T` variable | `*T` | `(&v).M()` |
| `*T` | `T` | `(*p).M()` |

- `&resident.HasRequiredInfo()` parses as `&(resident.HasRequiredInfo())`,
  which takes the address of a `bool`. That's a compile error.
- Explicit dereference needs parentheses: `(*resident).HasRequiredInfo()`.
- You write `&`/`*` yourself when **creating/passing** a pointer
  (`return &Resident{...}`, line 14) or **replacing the whole value** behind
  one (`*r = Resident{}`, line 51). Method calls don't need them.
- Exception: Go can't auto-`&` a value it can't take the address of, such as a
  map element or a function's return value.
