## main.go — 2026-09-19

**Q: How is this code working and what are the core concepts used?**

Flow:
1. `tick := time.Tick(100 * time.Millisecond)` — returns a `<-chan Time` that receives a value every 100ms, forever.
2. `boom := time.After(500 * time.Millisecond)` — returns a `<-chan Time` that receives exactly one value after 500ms elapses.
3. The `for` loop runs a `select` repeatedly:
   - `case <-tick`: fires when 100ms has passed since the last tick — prints "tick."
   - `case <-boom`: fires once, at ~500ms — prints "BOOM." and `return`s (exits the program).
   - `default`: runs immediately if neither channel has a value ready — prints a "." and sleeps 50ms.

Since `select` with a `default` never blocks, the loop is essentially: "check if it's tick-time or boom-time; if not, print a dot and nap for 50ms." This produces a mix of `.` filler lines and periodic `tick.`/`BOOM.` lines, ending when `boom` fires at 500ms.

Core concepts:
- **`select`**: waits on multiple channel operations; picks whichever is ready. With multiple ready cases it picks pseudo-randomly.
- **`default` case**: makes `select` non-blocking — if no channel is ready, it runs `default` instead of waiting. Turns `select` from a blocking multiplexer into a polling mechanism.
- **`time.Tick(d)`**: convenience `<-chan Time` that fires repeatedly every `d` (a lightweight `time.NewTicker` wrapper; leaks on exit since it's never stopped — fine here since the program is short-lived).
- **`time.After(d)`**: a `<-chan Time` that fires once after `d`.
- **Closures**: `elapsed` closes over `start` to compute formatted elapsed time on each print.
- **Busy-polling with sleep**: since `select`+`default` doesn't block, `time.Sleep(50ms)` inside `default` throttles the polling rate instead of spinning the CPU.

General pattern: use `select` with `default` when you want to check channels without blocking — e.g. polling loops, health checks, or doing background work while waiting for events.
