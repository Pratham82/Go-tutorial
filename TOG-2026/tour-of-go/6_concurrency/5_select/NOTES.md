## main.go — 2026-09-19

**Q: What's the use of select in go routines?**
`select` lets a goroutine wait on multiple channel operations at once and proceeds with whichever one is ready first — like a `switch` for channels.

In this file's loop:

```go
select {
case c <- x:
    x, y = y, x+y
case <-quit:
    fmt.Println("quit")
    return
}
```

`fibonacci` is simultaneously ready to **send** `x` on `c` or **receive** from `quit`. Whichever becomes ready first wins (random choice if both are ready at once).

Flow in this program:
1. `main` spawns a goroutine that reads from `c` 10 times (printing each value), then sends `0` on `quit`.
2. Each `fibonacci` iteration: since nothing reads `quit` yet, `c <- x` is the case that proceeds (channels are unbuffered, so it blocks until the receiver is ready).
3. This repeats 10 times, computing the next Fibonacci number each time.
4. After the 10th receive, the goroutine sends on `quit`. Now `<-quit` is the ready case, so `fibonacci` prints `"quit"` and returns.

Without `select`, `fibonacci` could only block on one channel operation at a time — it couldn't say "send a value, unless told to stop instead." This is the standard pattern for cancellation/shutdown signals alongside normal data flow.

**Q: Also check the output from the terminal**
Ran `go run main.go` — output matched expectations exactly:

```
0
1
1
2
3
5
8
13
21
34
quit
```

First 10 Fibonacci numbers (printed as each is received from `c`), followed by `quit` once the sender goroutine finishes its loop and signals via `quit`.
