# Go (backend) Resources

## Knowledge

- [Effective Go](https://go.dev/doc/effective_go)
  The canonical style/idiom guide from the Go team. Use for: naming, error idioms,
  interface design, when to use pointers, concurrency patterns.
- [Go blog: "Error handling and Go"](https://go.dev/blog/error-handling-and-go)
  Foundational post on the `error` interface and why Go has no exceptions. Use for:
  the mental model behind returning errors as values.
- [Go blog: "Working with Errors in Go 1.13"](https://go.dev/blog/go1.13-errors)
  Official explanation of `%w` wrapping, `errors.Is`, `errors.As`. Use for: modern
  error‑inspection code — this is the current best practice.
- [Go blog: "Errors are values"](https://go.dev/blog/errors-are-values)
  Rob Pike on treating errors as ordinary values you can program with. Use for:
  reducing repetitive `if err != nil` noise.
- [errors package docs](https://pkg.go.dev/errors)
  Primary source for `errors.New`, `Is`, `As`, `Join`, `Unwrap`. Use for: exact
  semantics when writing inspection logic.
- [Go Wiki: Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)
  The checklist Go reviewers actually apply. Use for: error strings, receiver
  names, doc comments, common mistakes.
- [Go by Example](https://gobyexample.com)
  Short, runnable snippets for one concept each (Mark McGranaghan). Use for: quick
  "how does X look in code" lookups.
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests) by Chris James
  TDD‑driven intro that builds toward real HTTP servers. Use for: testing habits
  and building a service test‑first.
- [The Go Programming Language](https://www.gopl.io/) by Donovan & Kernighan
  The reference book. Use for: deep, authoritative treatment of any core topic.
- [pkg.go.dev standard library](https://pkg.go.dev/std)
  Primary source for `net/http`, `encoding/json`, `context`, `testing`, `log/slog`.

## Wisdom (Communities)

- [r/golang](https://reddit.com/r/golang)
  Active, moderated against low‑effort posts. Use for: code review requests,
  "is this idiomatic?", job‑market questions.
- [Gophers Slack](https://invite.slack.golangbridge.org/)
  Official community Slack. Channels: #newbies, #reviews, #general. Use for:
  fast real‑time answers and getting code critiqued.
- [Go Forum](https://forum.golangbridge.org/)
  Longer‑form Q&A that stays searchable. Use for: design discussions.

## Gaps
- No chosen resource yet for interview‑style Go questions / take‑home patterns.
  Find one before the interview‑prep phase.
