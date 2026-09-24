# election_day.go — 2026-09-24

**Q: In the `VoteCount` function, how can I check if the pointer/counter is nil?**
Check the pointer parameter directly against `nil`, before doing any dereferencing:

```go
if counter == nil {
    return 0
}
```

You can't check `*counter == nil` — dereferencing (`*`) turns a `*int` into an `int`, and an `int` can never be `nil`. The nil check has to happen on the pointer itself, not on what it points to.

**Q: Check now (after adding the `counter == nil` check)**
The nil check on `counter` was right, but there was a leftover bug: `voteCounter` was declared as `var voteCounter *int` (zero value `nil`) and never given an address, so `*voteCounter = *counter` still panicked — you were dereferencing a nil pointer to write to it.

**Q: Check now (after simplifying)**
Removing the unused `voteCounter` pointer and returning `*counter` directly fixed it:

```go
func VoteCount(counter *int) int {
    if counter == nil {
        return 0
    }

    return *counter
}
```

No intermediate pointer was needed — once `counter` is confirmed non-nil, dereferencing it directly (`*counter`) is safe.

**Q: What was wrong before, when using `*voteCounter`?**
`var voteCounter *int` declares a pointer whose zero value is `nil` — it doesn't point to any `int` in memory. The line `*voteCounter = *counter` tried to **write** through that pointer ("go to the address `voteCounter` holds, and store this value there"), but `voteCounter` held no address. Dereferencing a nil pointer — for read or write — panics at runtime with `invalid memory address or nil pointer dereference`. The fix was to not use a second pointer at all: once `counter` itself is non-nil, dereference `counter` directly to read its value.
