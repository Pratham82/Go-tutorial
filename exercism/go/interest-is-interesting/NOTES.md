## interest_is_interesting.go — 2026-10-02

**Q: How do you write a switch case in Go?**
No `break` needed: only the first matching case runs, and cases don't fall
through (use `fallthrough` explicitly if you ever want that). A switch with no
value (`switch { ... }`) acts like an if/else chain: each `case` is a boolean
expression, checked top to bottom. `InterestRate` (lines 7–19) uses this form:

```go
switch {
case balance < 0:
    interestRate = 3.213
case balance >= 0 && balance < 1000:
    interestRate = 0.5
...
}
```

Gotchas hit while writing it:
- The function returns the **rate**, not `balance * rate`. Don't compute in
  the switch what the caller computes later.
- Boundaries matter: `balance > 0` left `0` matching no case (so it returned the
  zero value `0`), and `balance <= 0` put `0` in the negative tier. The README
  says 0 belongs to the 0.5% tier, so the boundaries are `< 0` and `>= 0`.
- Because cases are checked in order, `balance >= 0` in the second case is
  already guaranteed once `balance < 0` failed. That half of the condition is
  redundant (same for `>= 1000` and `>= 5000`).

**Q: Why convert the rate to float64 instead of the balance to float32?**
`Interest` (line 26) mixes a `float64` balance with a `float32` rate, and Go
won't multiply different types. `float32` keeps only about 7 significant
digits, so `float32(balance)` loses precision on large balances. It only
passed because the tests compare with a tolerance. Converting the smaller
type up keeps full precision:

```go
return balance * float64(InterestRate(balance)) / 100
```

**Q: How do you write a while loop in Go?**
Go has no `while`. A `for` with only a condition works the same way:

```go
for balance < targetBalance {   // while (balance < targetBalance)
    yearsTook++
    balance = AnnualBalanceUpdate(balance)
}
```

Other forms: `for { }` is an infinite loop (exit with `break`/`return`);
`do...while` is `for { body; if !cond { break } }`.

**Q: What should happen to the `if balance >= targetBalance { return 0 }` block?**
Delete it. A `for cond` loop checks its condition **before the first
iteration**, so if the balance already meets the target the body runs zero
times and `yearsTook` stays `0`, the same result as the guard.

The guard was only needed in the earlier `for { ...; if done { break } }`
version, because that loop always runs its body at least once (it would have
returned `1` instead of `0`). Changing the loop shape made the guard
unnecessary.

Also: `balance` is a parameter and Go passes it in as a copy, so you can update it
directly in the loop. There's no need for a separate `predictedBalance`.
