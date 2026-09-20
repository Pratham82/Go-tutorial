# Tour of Go — Notes

## Setting up an exercise that imports `golang.org/x/tour/...`

Some exercises (e.g. `tree`, `reader`, `pic`) import packages from `golang.org/x/tour`.
Since each exercise folder is its own Go module, you need to generate `go.mod` and `go.sum`
before the import will resolve.

From inside the exercise directory:

```bash
go mod init <exercise-name>
go mod tidy
```

- `go mod init <exercise-name>` creates `go.mod` and declares the module name. The name is
  arbitrary for these exercises — just pick something descriptive (convention used here:
  kebab-case, suffixed with `-exercise`, e.g. `equivalent-binary-trees-exercise`).
- `go mod tidy` scans the code for imports, fetches `golang.org/x/tour`, and writes the
  `require` line into `go.mod` plus checksums into `go.sum`.

Verify both files exist:

```bash
ls -la go.*
```

Both `go.mod` and `go.sum` should be committed to git along with the exercise code.
