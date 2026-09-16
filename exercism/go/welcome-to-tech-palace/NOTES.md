## welcome_to_tech_palace.go — 2026-09-16

**Q: How to remove all whitespace from a string, and how to replace characters in a string in Go?**
Use the `strings` package:
- Remove all whitespace: `strings.Join(strings.Fields(s), "")` — `strings.Fields` splits on any whitespace and drops empty strings, then `Join` with `""` collapses everything.
- Strip whitespace including exotic ones (e.g. non-breaking space `U+00A0`): `strings.Map` with `unicode.IsSpace`, returning `-1` to drop the rune.
- Replace characters: `strings.ReplaceAll(s, old, new)` for all occurrences, `strings.Replace(s, old, new, n)` to cap replacements, or `strings.NewReplacer(...)` for several different replacements at once.

**Q: What's the issue in this version of the cleanup function?**
```go
finalString := strings.Join(strings.Fields(withoutStars), "")
```
Joining with `""` (empty string) squashes every word together with no spaces at all — e.g. `"BUY NOW, SAVE 10%"` becomes `"BUYNOW,SAVE10%"`. Since the goal is to collapse extra whitespace/newlines down to single spaces (not eliminate spacing), the separator should be `" "` instead of `""`. The actual `CleanupMessage` in this file (line 19) already uses the correct `" "` separator.
