# NOTES

## main.go — 2026-09-19

**Q: How is this working?**
`ch := make(chan int, 3)` creates a channel with buffer capacity 3. All three
sends (`ch <- 1`, `ch <- 2`, `ch <- 11`) succeed immediately without blocking
since the buffer never fills past capacity. The three receives then pop
values off in FIFO order, printing `1 2 11`.

**Q: I added a change (extra `ch <- 4`/extra receive) — check the terminal for output**
The file was changed to 3 sends but 4 receives:
```go
ch := make(chan int, 3)
ch <- 1
ch <- 11
ch <- 2

fmt.Println(<-ch)
fmt.Println(<-ch)
fmt.Println(<-ch)
fmt.Println(<-ch)
```
Running it printed `1`, `11`, `2`, then hit the 4th `<-ch` which blocks
forever (buffer is empty, nothing else sends), producing:
```
fatal error: all goroutines are asleep - deadlock!
```
Note: this is actually *over-receiving* (4 receives for 3 sends), not
overfilling the *send* buffer. To trigger the overfill-on-send case
instead, add a 4th send with no extra receive (e.g. `ch <- 4` after the
existing three) — that blocks on the send side for the same underlying
reason (no goroutine can make progress).

**Q: So the length that we are passing is what length?**
The second argument to `make(chan int, 3)` is the channel's **buffer
capacity** — how many values it can hold in flight before a send blocks.
- A send blocks once the buffer already holds `capacity` un-received values.
- A receive blocks when the buffer is empty.
An *unbuffered* channel (`make(chan int)`, capacity 0) requires sender and
receiver to rendezvous — no queuing at all.

**Q: Can you give me an example of unbuffered?**
See `unbuffered_example.go` (tagged `//go:build ignore` so it doesn't
conflict with `main.go` in the same package; run with
`go run unbuffered_example.go`):
```go
ch := make(chan int) // unbuffered: capacity 0

go func() {
    ch <- 42 // blocks until main receives
}()

fmt.Println(<-ch) // blocks until goroutine sends
```
Needs a separate goroutine to send, since the unbuffered send can't
complete until a receiver is actively waiting at the same moment. Prints
`42`.

**Q: What's a channel exactly and when is it used?**
A channel is a typed conduit for passing values between goroutines with
built-in synchronization (Go's "share memory by communicating" idiom).
Created with `make(chan T)` or `make(chan T, n)`; `ch <- v` sends, `v :=
<-ch` receives, `close(ch)` signals no more values.

Common uses:
1. Handing off work/results between goroutines instead of shared variables.
2. Synchronization/signaling — e.g. an empty `chan struct{}` used just to
   say "done".
3. Worker pools / pipelines — goroutines reading tasks off one channel,
   writing results to another.
4. Rate limiting / capping concurrency — buffered channel used as a
   semaphore.
5. Fan-in/fan-out and `select` for multiplexing multiple channels
   (timeouts, cancellation, combining event sources).

Buffered vs. unbuffered tradeoff: unbuffered guarantees the sender knows
the receiver got the value at that exact instant (strict synchronization);
buffered decouples producer/consumer pacing at the cost of that guarantee.
