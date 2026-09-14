# Exercism: testing and submitting a solution

## 1. Run the tests locally

`go test` treats a bare word as an import path, not a folder — so from the
repo root, `go test hello_world` fails with "package ... is not in std".

Always `cd` into the exercise's own directory (it has its own `go.mod`) and run:

```
cd go/<exercise-name>
go test ./...
```

## 2. Submit to Exercism

Workspace is already configured (`exercism configure` was run once), pointing
at this repo: `/Users/pratham82/dev/Go-tutorial`.

Once tests pass, from inside the exercise directory:

```
exercism submit <filename>.go
```

This uploads the solution to exercism.org. It prints a URL to the exercise
page — open it and click **Complete exercise** to mark it done.

## Common Go gotchas hit along the way

- `:=` and bare assignment (`x = 40`) are **statements** — only legal inside a
  function body. At package level you need a declaration: `const X = 40` or
  `var X = 40`.
- Lowercase identifiers (`ovenTime`) are unexported — invisible to test files
  expecting `OvenTime`. Match the exported name the tests use.
- `go run` requires `package main` + `func main()`. Exercism exercises are
  libraries (`package greeting`, etc.) with no `main` — use `go test`, not
  `go run`.
