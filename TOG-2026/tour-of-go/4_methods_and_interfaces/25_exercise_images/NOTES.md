## main.go — 2026-09-12

**Q: What's needed to complete this exercise?**
The starter file had an empty `Image struct{}` and a dangling, incomplete
`func` declaration (no name/body) — it wouldn't compile. This exercise
builds on `24_images`: define your own type implementing `image.Image`
and pass it to `pic.ShowImage`. Needed:
1. Give `Image` fields for its dimensions (`w, h int`).
2. Implement the three interface methods:
   - `Bounds() image.Rectangle` → `image.Rect(0, 0, w, h)`
   - `ColorModel() color.Model` → `color.RGBAModel`
   - `At(x, y int) color.Color` → compute `v` and return
     `color.RGBA{v, v, 255, 255}` (per INFO.md's mapping from the earlier
     picture-generator exercise)
3. Remove the broken `func` stub.

**Q: I'm seeing an error on the `golang.org/x/tour/pic` import — what's the fix?**
Not a code bug — the `Image` implementation was correct. The problem was
this directory had no `go.mod`, so Go couldn't resolve
`golang.org/x/tour/pic`. Other exercises that import `golang.org/x/tour/...`
(e.g. `22_exercise_readers`) each have their own `go.mod`/`go.sum`. Fixed
by running, inside this directory:
```
go mod init images-exercise
go get golang.org/x/tour@v0.1.0
go mod tidy
```
After that, `go run main.go` compiled and `pic.ShowImage` printed the
rendered image data to stdout, confirming the `Image` type correctly
satisfies `image.Image`.
