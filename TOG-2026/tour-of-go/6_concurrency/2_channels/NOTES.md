## main.go — 2026-09-18

**Q: Can you explain what's happening in main.go?**
This is the classic Tour of Go example showing goroutines and channels working together.

- `func sum(s []int, c chan int)` (line 5-11) computes the sum of a slice, then sends the result into the channel via `c <- sum` instead of returning it.
- `c := make(chan int)` creates an *unbuffered* channel of `int`. Unbuffered means a send blocks until someone receives, and vice versa.
- `go sum(s[:len(s)/2], c)` and `go sum(s[len(s)/2:], c)` launch two goroutines concurrently, each summing one half of `s := []int{7, 2, 8, -9, 4, 0}` (i.e. `{7, 2, 8}` and `{-9, 4, 0}`).
- `x, y := <-c, <-c` is where synchronization happens: main blocks on the first `<-c` until one goroutine sends its partial sum, then blocks on the second `<-c` until the other sends its sum. This naturally synchronizes execution — no `sync.WaitGroup` needed — because the unbuffered channel forces the sends to wait for receives.
- `x` and `y` are **not** guaranteed to correspond to a specific half — whichever goroutine finishes first sends first, so the order is nondeterministic. The total `x+y` is deterministic (17 + -5 = 12), but individual print order can be `17 -5 12` or `-5 17 12`.

**Key concept:** channels here serve two purposes at once — passing data *and* synchronizing goroutines.
