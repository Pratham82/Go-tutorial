# Running tests

Each exercise is its own Go module (separate `go.mod`), so `go test ./...` from
this `exercism/go` root won't reach into them. Use one of the two approaches
below.

## From the `exercism/go` root

```sh
# Single test function
go test -C <exercise-dir> -run <TestName> -v ./...

# Whole file (all tests)
go test -C <exercise-dir> -v ./...
```

Example:

```sh
go test -C logs-logs-logs -run TestReplace -v ./...
go test -C logs-logs-logs -v ./...
```

## From inside the exercise folder

```sh
# Single test function
go test -run <TestName> -v ./...

# Whole file (all tests)
go test -v ./...
```

Example (from `exercism/go/logs-logs-logs`):

```sh
go test -run TestReplace -v ./...
go test -v ./...
```

`-run <TestName>` matches the test function name as a regex, so it also runs
matching subtests. `-v` prints each subtest's pass/fail; drop it for a plain
summary.
