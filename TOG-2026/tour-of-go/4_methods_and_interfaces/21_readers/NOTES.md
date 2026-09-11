## main.go — 2026-09-11

**Q: Please explain the working and what's happening here?**
Introduces `io.Reader`, the interface behind Go's streaming I/O:
```go
type Reader interface {
	Read(b []byte) (n int, err error)
}
```
Any type with that one method satisfies it — same structural rule as `Stringer`/`error`. Contract: fill up to `len(b)` bytes into `b`, return how many were actually written (`n`) and an error (`io.EOF` once the source is exhausted).

```go
r := strings.NewReader("Hello, Reader!")   // *strings.Reader, satisfies io.Reader
b := make([]byte, 8)                        // 8-byte buffer, reused every iteration
```

Loop (lines 13-21):
```go
for {
	n, err := r.Read(b)
	...
	if err == io.EOF {
		break
	}
}
```
Each `r.Read(b)` writes up to 8 bytes into `b` and returns `n` = bytes actually written this call. `b[:n]` is the only fresh data — `b` itself is 8 bytes long and reused/overwritten in place, so leftover bytes from a previous call can remain in slots beyond `n` (that's why the printed `b` array looks unchanged between the 2nd and 3rd iteration).

Actual output:
```
n = 8 err = <nil> b = [72 101 108 108 111 44 32 82]
b[:n] = "Hello, R"
n = 6 err = <nil> b = [101 97 100 101 114 33 32 82]
b[:n] = "eader!"
n = 0 err = EOF b = [101 97 100 101 114 33 32 82]
b[:n] = ""
```
Go's `Read` convention: EOF can be reported in a separate call from the last chunk of data (as it is here) rather than combined with it — callers must handle both patterns, which is why the loop checks `err == io.EOF` *after* using whatever data came back, not before.

**Q: Why is the word split on "R" then "eader" then ""?**
Pure function of buffer size vs. string length — nothing to do with the word itself.

`"Hello, Reader!"` is 14 characters:
```
H  e  l  l  o  ,     R  |  e  a  d  e  r  !
1  2  3  4  5  6  7  8  |  9 10 11 12 13 14
```
- 1st `Read(b)`: buffer holds max 8 bytes → grabs the first 8 → `"Hello, R"`. Coincidence that it lands right after "R".
- 2nd `Read(b)`: 14 − 8 = 6 remain → copies those 6 into the same buffer → `"eader!"` (`n = 6`); the buffer's last 2 slots keep stale bytes from the previous call, but `b[:n]` still shows the correct fresh substring.
- 3rd `Read(b)`: reader's position is now at the end, nothing left → `strings.Reader` reports this with a dedicated `n = 0, err = io.EOF` call. `b[:0]` is `""`.

Changing `make([]byte, 8)` to `make([]byte, 4)` would instead give `"Hell"`, `"o, R"`, `"eade"`, `"r!"`, `""` — the chunking is entirely buffer-size-driven.
