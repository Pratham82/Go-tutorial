## main.go — 2026-09-20

**Q: Just explain what's happening here**

`SafeCounter` wraps a `map[string]int` (`v`) together with a `sync.Mutex` (`mu`). Go maps aren't safe for concurrent read/write, so the mutex is what makes this struct safe to share across goroutines.

- `Inc(key)` (main.go:16-21) locks `mu`, increments `c.v[key]`, then unlocks. The lock ensures only one goroutine at a time can touch the map, so 1000 concurrent increments don't race and corrupt it or get lost.
- `Value(key)` (main.go:24-29) locks `mu` and uses `defer c.mu.Unlock()` so the unlock always happens right before the function returns, then reads `c.v[key]`.
- `main()` creates a `SafeCounter`, launches 1000 goroutines each calling `c.Inc("somekey")` (all racing to increment the same key), then `time.Sleep(time.Second)` crudely waits for them to finish before reading the value — not a robust pattern (a `sync.WaitGroup` would be correct), but kept simple for this tour example.
- Prints `c.Value("somekey")`, which should be `1000` since the mutex prevented any lost updates.
