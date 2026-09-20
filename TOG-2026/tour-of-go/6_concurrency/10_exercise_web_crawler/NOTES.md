## main.go — 2026-09-20

**Q: Can you explain what's happening ? and what's the requirement of the exercise ?**

`Fetcher` is an interface with `Fetch(url) (body string, urls []string, err error)`. `Crawl(url, depth, fetcher)` recursively fetches a URL's body and the URLs it links to, printing what it finds, then recurses into those URLs with `depth-1` until `depth <= 0`. `fakeFetcher`/`fetcher` simulate a tiny site graph (including cycles, e.g. `pkg/` links back to `golang.org/`) so the exercise can be tested without real network calls.

The original `Crawl` was sequential and unguarded — per the TODOs at the top of the function, the exercise requires:
1. Fetching URLs in parallel (concurrently, not one after another).
2. Not fetching the same URL twice (needed because the fake site graph has cycles).

### Solution and how it works

Solution (main.go):
- Added a `visited` struct pairing a `map[string]bool` with a `sync.Mutex` (same pattern as the `SafeCounter` mutex exercise) to safely track which URLs have already been crawled across goroutines.
- In `Crawl`, lock the mutex, check `seen.v[url]`, and if not already visited mark it `true` — all under one lock so two goroutines can't both pass the check for the same URL and double-fetch.
- Each recursive call is launched as `go Crawl(...)` so sibling URLs are fetched in parallel instead of sequentially.
- A `sync.WaitGroup` tracks in-flight goroutines: `wg.Add(1)` before each `go Crawl(...)`, `defer wg.Done()` at the top of `Crawl`, and `wg.Wait()` in `main` so the program doesn't exit before all fetches complete.

Verified by running `go run main.go` — each URL (`golang.org/`, `pkg/`, `pkg/fmt/`, `pkg/os/`) is printed exactly once, plus one `not found: https://golang.org/cmd/` error, with no duplicates or infinite loop despite the cyclic links.
