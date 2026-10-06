## expenses.go — 2026-10-06

**Q: What's wrong with `var filtered = []Record` in `Filter`?**
`[]Record` is only a type, but the right-hand side of `=` must be a value. Use `var filtered []Record` (nil slice, idiomatic when you will append), `filtered := []Record{}` or `make([]Record, 0)`. The same function also had `filtered.append(record)`, which doesn't exist. `append` is a built-in that returns the new slice, so it has to be `filtered = append(filtered, record)`.

**Q: How can I give two optional types to a variable?**
Go has no union types. Options: an interface (`any`) plus a type switch, a generic constraint (`int | string`, usable only as a type parameter), a pointer for "value or nothing", a struct with two pointer fields, or a sealed interface with an unexported method.

**Q: How do I check if a key/value is present in a struct?**
Struct fields always exist, and unset ones hold their zero value. Compare to the zero value (`r.Category == ""`), or use a pointer field and check for `nil` when zero is a valid value. For a map, use the two-value form: `v, ok := m[key]`.

**Q: How can I reuse `r.Day >= p.From && r.Day <= p.To`?**
`ByDaysPeriod(p)` already builds that check. Build the predicate once and call it:
```go
inPeriod := ByDaysPeriod(p)
if inPeriod(record) { ... }
```
It can also be passed to `Filter`. Don't call `ByDaysPeriod` from inside itself (the commented-out lines 35 and 38 would recurse).

**Q: How do I know the category is ever present without mapping over the whole slice?**
With an unsorted slice you must scan it. `CategoryExpenses` already loops over every record, so set `isCategorySeen` in that loop, outside the period check:
```go
if c == record.Category {
    isCategorySeen = true
}
if inPeriod(record) && isCategory(record) {
    total += record.Amount
}
```
Putting the flag inside the period check made the "0 when nothing in period" test fail, because the error is meant to depend on the category only. `in` is a slice, not a map. A map lookup is O(1), but building it costs a scan anyway.

**Q: How do I create a formatted error in one step (lines 95–96)?**
`fmt.Errorf` formats and returns an error:
```go
return 0, fmt.Errorf("unknown category %s", c)
```
This replaces `Sprintf` plus `errors.New`. Remove the `"errors"` import afterwards, or it won't compile.

**Q: How do I avoid `if cond { return true }; return false` in `ByDaysPeriod` and `ByCategory`?**
The condition is already a `bool`, so return it directly:
```go
return r.Day >= p.From && r.Day <= p.To
return r.Category == c
```
