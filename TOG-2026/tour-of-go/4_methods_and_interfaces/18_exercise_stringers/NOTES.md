## main.go — 2026-09-11

**Q: How to give a type definition? (`type IPAddr { [4]byte }` didn't compile)**
Go's `type` declaration is `type <Name> <underlying-type>` — no braces unless the underlying type itself needs them (struct or interface). For this exercise:
```go
type IPAddr [4]byte
```
`IPAddr` is a distinct named type with underlying type `[4]byte`, so its values are indexable like an array (`ip[0]`…`ip[3]`).

General shapes:
```go
type Celsius float64            // named scalar
type IPAddr  [4]byte            // named array
type IntSlice []int             // named slice
type Handler func(int) error    // named function type
type Point struct { X, Y int }  // braces: struct needs them
type Reader interface {         // braces: interface needs them
	Read(p []byte) (int, error)
}
```

**Q: How is this working? (the finished exercise)**
Runs correctly:
```
loopback: 127.0.0.1
googleDNS: 8.8.8.8
```

- Line 5 `type IPAddr [4]byte` — named array type; each element is a `byte` (= `uint8`).
- Lines 8-11 the `String()` method:
  ```go
  func (ip IPAddr) String() string {
      return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
  }
  ```
  Value receiver + signature `String() string` → `IPAddr` satisfies `fmt.Stringer`. `%d` formats each byte as decimal. No infinite recursion: `Sprintf` here formats individual `byte`s with `%d`, not the `IPAddr` itself.
- Lines 14-17 map literal `map[string]IPAddr{ "loopback": {127,0,0,1}, ... }` — element type is known from the map type, so `IPAddr{...}` is elided; the 4 values fill the array positionally.
- Lines 19-21 `for name, ip := range hosts { fmt.Printf("%v: %v\n", name, ip) }` — `%v` on `ip` triggers the Stringer check: `fmt` finds `IPAddr.String()` and calls it, giving `127.0.0.1` instead of the default array dump `[127 0 0 1]`.
- Map iteration order is random, so the two output lines may swap between runs.
