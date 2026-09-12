## main.go — 2026-09-12

**Q: Can you explain what's happening here?**
`rot13Reader` wraps another `io.Reader` (`r`). `main` builds one around a
`strings.NewReader` holding the ROT13-encoded string `"Lbh penpxrq gur pbqr!"`,
then passes `&r` to `io.Copy(os.Stdout, &r)`. For `io.Copy` to accept it,
`rot13Reader` must satisfy `io.Reader` by implementing:

```go
Read(p []byte) (n int, err error)
```

The pattern mirrors `gzip.Reader` (mentioned in `INFO.md`): it decorates an
existing reader, transforming the bytes it reads before handing them back,
without buffering the whole stream itself.

**Q: Can you help me complete the exercise?**
Two things were needed:
1. A `rot13` byte-shift helper — shifts letters by 13 within their case
   (`A-Z`, `a-z`), leaves everything else untouched:
   ```go
   func rot13(b byte) byte {
   	switch {
   	case b >= 'A' && b <= 'Z':
   		return 'A' + (b-'A'+13)%26
   	case b >= 'a' && b <= 'z':
   		return 'a' + (b-'a'+13)%26
   	default:
   		return b
   	}
   }
   ```
2. The `Read` method, which delegates to the wrapped reader to fill `p`,
   then transforms the bytes in place before returning:
   ```go
   func (r rot13Reader) Read(p []byte) (n int, err error) {
   	n, err = r.r.Read(p)
   	for i := 0; i < n; i++ {
   		p[i] = rot13(p[i])
   	}
   	return n, err
   }
   ```
A first draft had a typo in the return signature (`n int, err, error`
instead of `n int, err error`), which was a syntax error, not a logic bug.
After fixing it, `go run main.go` printed `You cracked the code!`.
