## animal_magic.go — 2026-10-06

**Q: How do I import rand in Go?**
Use the standard library. In Go 1.22+ prefer `math/rand/v2` (`rand.IntN`, `rand.Float64`, `rand.Shuffle`, auto-seeded). The older `math/rand` uses `rand.Intn` and is auto-seeded since Go 1.20, so `rand.Seed` is deprecated. For security-sensitive values (tokens, keys) use `crypto/rand` instead.

```go
import "math/rand/v2"
```

**Q: Does `rand.IntN` take a range?**
No. `rand.IntN(n)` takes a single upper bound and returns a value in `[0, n)`. It panics if `n <= 0`. For a custom range, add an offset yourself.

**Q: How do I give a starting range?**
Add the start as an offset: `start + rand.IntN(end-start)` gives `[start, end)`. Add `+1` inside to include `end`.

**Q: How does this work?**
```go
var min, max = 1, 20
res := min + rand.IntN(max-min+1)
```
- `max-min+1` is `20`, the number of values from 1 to 20 inclusive (the `+1` is because both ends are included).
- `rand.IntN(20)` returns `0..19`.
- Adding `min` shifts that to `1..20`.

`min` and `max` are also built-in functions in Go 1.21+; declaring variables with those names shadows them inside the function. Legal, but `lo`/`hi` avoids the confusion. The `res` variable isn't needed; `return min + rand.IntN(max-min+1)` works directly.

**Q: How do I do the same for `GenerateWandEnergy` (`0.0 <= f < 12.0`)?**
`rand.Float64()` returns `[0.0, 1.0)`. Scale it by multiplying by the width of the range:

```go
return rand.Float64() * 12.0
```

General form for `[min, max)`: `min + rand.Float64()*(max-min)`.

**Q: How does `ShuffleAnimals` take the input string slice?**
It doesn't. The signature has no parameters (the test calls `ShuffleAnimals()`), so the eight animal names are declared inside the function. `rand.Shuffle(n, swap)` shuffles in place and returns nothing, so call it as a statement and then return the slice.

**Q: What was wrong with my first attempt?**
```go
func ShuffleAnimals(x []string) []string {
	return rand.Shuffle(len(x), func(i, j) {
		x[i], x[j] = x[j], x[i]
	})
}
```
1. Wrong signature: the test calls it with no arguments.
2. `rand.Shuffle` returns nothing, so `return rand.Shuffle(...)` doesn't compile ("no value used as value").
3. `func(i, j)` reads `i` and `j` as type names; parameters need a type: `func(i, j int)`.

**Final:**
```go
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}

	rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})

	return animals
}
```
The swap `animals[i], animals[j] = animals[j], animals[i]` uses Go's multiple assignment, so no temp variable is needed.
