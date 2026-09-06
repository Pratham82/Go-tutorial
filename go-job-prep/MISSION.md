# Mission: Go (Golang) for backend engineering

## Why
Land a paid job as a Go backend developer within 3–6 months (targeting mid‑2026).
The concrete outcome: be able to walk into an interview and a first week on the job
and confidently build, test, and ship an HTTP/JSON API service in Go without
hand‑holding.

## Success looks like
- Build a small REST API from scratch: routing, JSON encode/decode, request
  validation, structured logging, graceful shutdown.
- Handle errors the idiomatic Go way: wrapping, `errors.Is`/`errors.As`, sentinel
  vs. typed errors, and knowing when to return vs. log.
- Use goroutines and channels correctly for a real task (e.g. fan‑out to an
  upstream API with a timeout and `context` cancellation) without data races.
- Talk through a `struct`/interface design and justify value vs. pointer receivers.
- Write table‑driven tests with the standard `testing` package and run them with
  `go test -race`.
- Read an unfamiliar Go codebase and package it into a runnable service.

## Constraints
- Timeline: 3–6 months, part‑time study.
- Already a comfortable programmer in another language — skip "what is a variable",
  focus on Go idioms and the standard library.
- Has worked through the Tour of Go up to Methods & Interfaces.
- Learning style: short lessons, retrieval practice, one tangible win per session.

## Out of scope (for now)
- Generics deep‑dives beyond basic usage.
- Kubernetes operators / infra tooling (role target is API services).
- Web frameworks (Gin, Echo, Fiber) until the standard library `net/http` is solid.
- gRPC until REST/JSON is comfortable.
