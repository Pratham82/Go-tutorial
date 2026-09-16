# Go Tutorial

Learning Go for fun. Repo for storing all Go code, exercises, and notes while learning — spans basics, the official Tour of Go, Exercism exercises, and a few standalone projects.

## Folders

- [`TOG-2026/tour-of-go`](TOG-2026/tour-of-go) — Current run through the official [Tour of Go](https://go.dev/tour/), organized by topic: basics, flow control, more types, methods & interfaces, generics.
- [`Tour of Go`](Tour%20of%20Go) — Earlier pass through the Tour of Go (Hello World, Basics modules).
- [`exercism/go`](exercism/go) — Solutions to [Exercism](https://exercism.org/tracks/go) Go track exercises (each has its own `go.mod`). Includes `cars-assemble`, `hello-world`, `welcome-to-tech-palace`, `annalyns-infiltration`, `vehicle-purchase`, `weather-forecast`, `lasagna`.
- [`go-job-prep`](go-job-prep) — Interview/job-prep notes, lessons, and reference material.
- [`go-projects`](go-projects) — Small standalone practice projects (e.g. [`01-calculator`](go-projects/01-calculator)).
- [`learn-go-fast`](learn-go-fast) — Tutorial-following code organized as runnable `cmd` programs.
- [`01-hello-world`](01-hello-world) — First Go program.

## Running exercises

Most Exercism exercises are their own Go module. From the repo root:

```bash
cd exercism/go/<exercise-name> && go test ./...
```

Or without `cd`, using Go 1.20+'s `-C` flag:

```bash
go test -C exercism/go/<exercise-name> ./...
```

## Resources

1. [Tour of Go](https://go.dev/tour/) — official interactive introduction to Go's syntax and features
2. [Exercism Go track](https://exercism.org/tracks/go) — free coding exercises with mentor feedback for practicing Go
3. [Gophercises](https://gophercises.com/) — free coding exercises focused on building small real-world Go programs
4. [Go by Example](https://gobyexample.com/) — hands-on introduction to Go through annotated example programs
5. [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests) — learn Go via test-driven development, from basics to concurrency
6. [Effective Go](https://go.dev/doc/effective_go) — official guide to writing idiomatic Go
7. [Go Proverbs](https://go-proverbs.github.io/) — Rob Pike's short idioms/philosophy behind Go's design
8. [Awesome Go](https://awesome-go.com/) — curated list of Go frameworks, libraries, and tools
9. [Go Standard Library docs](https://pkg.go.dev/std) — searchable reference for every package in the standard library
10. [100 Go Mistakes and How to Avoid Them](https://100go.co/) — common pitfalls and idiomatic fixes
11. [Exercism Go Track Concepts](https://exercism.org/tracks/go/concepts) — concept-by-concept breakdown backing the Exercism exercises
