## main.go — 2026-09-20

**Q: Can you explain what I have to do in main.go? What are the requirements and how to get it done?**
This is the "Equivalent Binary Trees" exercise. `tree.New(k)` builds a randomized BST containing `k, 2k, ..., 10k`, each exactly once — so two trees built with the same `k` always hold the same *values*, just shaped differently. Three things to implement:
1. `Walk(t *tree.Tree, ch chan int)` — walk the tree, sending every node's `Value` to `ch`, then close `ch` when done.
2. `Same(t1, t2 *tree.Tree) bool` — use two goroutines running `Walk` on separate channels, then compare the sequences received to determine if the trees hold the same values.
3. `main()` — exercise both: print `Walk`'s output for `tree.New(1)`, and check `Same(tree.New(1), tree.New(1))` (true) vs `Same(tree.New(1), tree.New(2))` (false).

### Solution and how it works

Implemented as:
```go
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if ok1 != ok2 {
			return false
		}
		if !ok1 {
			return true
		}
		if v1 != v2 {
			return false
		}
	}
}
```
Verified with `go run main.go`: prints `1..10` in order, `Same(New(1), New(1))` → `true`, `Same(New(1), New(2))` → `false`.

**Q: Can you explain each function step by step — what's happening and why is it like that?**
- **`walk`** (unexported helper): base case `if t == nil { return }` stops recursion at leaves' nil children. Recurses left, sends the current value, recurses right (in-order traversal — order doesn't matter for correctness, just needs to be consistent). It's split from `Walk` so that only the outermost call closes the channel — if `walk` itself closed `ch`, every recursive call would try to close it and panic on the second attempt.
- **`Walk`** (exported, matches required signature): calls `walk` to fully traverse, then `close(ch)` exactly once after all values are sent. Closing tells readers ("no more values coming") so `range` or `v, ok := <-ch` know when to stop instead of blocking forever.
- **`Same`**: uses unbuffered channels (`make(chan int)`) so each `Walk` blocks on `ch <- t.Value` until read — this lets both trees be walked concurrently and compared value-by-value without collecting either into a slice first. Runs each `Walk` with `go` because calling them directly would deadlock (nothing would be reading yet). The loop uses `v, ok := <-ch` (two-value receive) to detect closed channels — `ok` becomes `false` once a channel is drained and closed. Check order matters: (1) `ok1 != ok2` catches trees with different value counts (one channel closes before the other) → `false`; (2) `!ok1` (both closed, everything matched so far) → `true`; (3) `v1 != v2` catches a value mismatch while both channels are still live → `false`. Checking `ok` before `v` avoids false matches against the zero value (`0`) that a closed channel keeps returning.
- **`main`**: `go Walk(tree.New(1), ch)` + `for v := range ch` prints values 1-10 (range auto-stops at channel close, no manual `ok` check needed). Then calls `Same` with a matching pair and a mismatching pair to demonstrate both outcomes.
