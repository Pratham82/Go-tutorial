## main.go — 2026-09-19

**Q: What's wrong with this for loop? `for i:= 0, i< n; i++ { c <- }`**
Three bugs:
1. `i := 0, i< n` uses a comma instead of `;` — Go's `for` clauses are separated by semicolons: `i := 0; i < n; i++`.
2. `c <-` is incomplete — a send needs a value, e.g. `c <- x`.
3. Missing the actual Fibonacci logic (updating `x, y`).

Fixed version used in the final code (see below) replaces the loop with `for range n` and sends `x` each iteration.

**Q: What's cap?**
`cap` is a built-in Go function returning capacity:
- For a channel made with `make(chan int, 10)`, `cap(c)` is `10` (the buffer size). `len(c)` is how many values are currently queued. Sending blocks once `len(c) == cap(c)`. An unbuffered channel has `cap` `0`.
- For a slice `make([]int, 3, 10)`, `cap` is `10` (space before reallocation), `len` is `3` (current elements).
In this file, `cap(c)` is used at line 17 (`go fibonacci(cap(c), c)`) to tell `fibonacci` how many numbers to generate, based on the channel's buffer size (10) rather than a hardcoded number.

**Q: What's happening in this file?**
- `fibonacci(n int, c chan int)` (lines 5-13): starts `x, y := 0, 1`, then `for range n` runs the body `n` times (range-over-int, Go 1.22+). Each iteration sends `x` on `c`, then advances `x, y = y, x+y`. After the loop, `close(c)` signals no more values will come.
- `main()` (lines 15-22): creates a buffered channel `c := make(chan int, 10)`, launches `fibonacci(cap(c), c)` as a goroutine (so it generates exactly 10 numbers), then `for i := range c` receives values as they arrive and automatically stops when `c` is closed. Each value is printed.
- Net effect: prints the first 10 Fibonacci numbers (0, 1, 1, 2, 3, 5, 8, 13, 21, 34), with producer and consumer running concurrently — the buffer lets `fibonacci` send without blocking until full, and `close` lets `main`'s range loop know when to stop.
