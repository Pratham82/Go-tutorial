## main.go — 2026-09-17

**Q: Can you explain what's happening here?**
```go
func say(s string) {
	for range 5 {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

func main() {
	go say("hello r")
	say("hello")
}
```
- `go say("hello r")` starts `say("hello r")` as a goroutine — a lightweight concurrent thread of execution. `main()` doesn't wait for it; it immediately moves to the next line.
- `say("hello")` then runs synchronously on the main goroutine, printing "hello" 5 times with a 100µs sleep between each print, while `"hello r"` prints concurrently on its own goroutine.
- There's no synchronization (no `sync.WaitGroup`, channel, etc.), so:
  1. The two `say` calls interleave unpredictably — output order is non-deterministic.
  2. `main()` returns as soon as `say("hello")` finishes its 5 iterations, and when `main` returns the whole program exits immediately — even if the `"hello r"` goroutine hasn't finished printing yet.

**Q: The output is not uniform — sometimes "hello r" prints 5 times, sometimes not. What's the reason?**
This is the race condition inherent to the code:
- **Interleaving order varies** because both goroutines are runnable at the same time and the Go scheduler decides which one gets CPU time at each instant — this depends on OS thread scheduling, GOMAXPROCS, machine load, timer precision, etc. The very short `100 * time.Microsecond` sleep makes which goroutine wakes up first essentially arbitrary.
- **Fewer than 5 "hello r" prints can happen** because there's a race between the main goroutine finishing its 5 prints (and returning, which kills the whole process) and the `"hello r"` goroutine finishing its own 5 prints. Go does **not** wait for other goroutines when `main` returns. If the main goroutine finishes first, the program can exit before `"hello r"` has printed all 5 lines.

**Fix** — make `main` wait for the goroutine with a `sync.WaitGroup`:
```go
func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 5 {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go say("hello r", &wg)
	say("hello", nil) // or also pass wg and Add(2)
	wg.Wait()
}
```
This guarantees both goroutines finish before `main` returns, though the *interleaving order* of prints remains non-deterministic — that requires separate coordination (mutex/channel) if ordering matters.
</content>
