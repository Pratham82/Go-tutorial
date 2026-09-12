## main.go — 2026-09-12

**Q: Can you explain what's happening here?**
This is a demo of the `image.Image` interface, not an exercise to fix.
`image.Image` (defined in `INFO.md`) requires three methods:
```go
type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```
The code uses the standard library's concrete implementation instead of
writing a custom one:
- `image.NewRGBA(image.Rect(0, 0, 100, 100))` allocates a 100x100
  `*image.RGBA` bitmap — corners (0,0) to (100,100).
- `m.Bounds()` returns that `image.Rectangle`, printed as
  `(0,0)-(100,100)`.
- `m.At(0, 0)` returns the pixel's `color.Color`; since the image is
  freshly zero-valued, that pixel is black/transparent. `.RGBA()` (a
  `color.Color` method) returns the four channels as
  `(r, g, b, a uint32)`, printing `0 0 0 0` here.

Both `color.Color` and `color.Model` are themselves interfaces
(`image/color` package) — this file just uses their predefined
implementations (`color.RGBA`, `color.RGBAModel`) to get familiar with the
`image.Image` interface's methods before likely implementing it on a
custom type in a later exercise.
