## booking_up_for_beauty.go — 2026-09-20

**Q: How does `time.Parse`/`time.Format` know which field is day vs month?**
Go uses one fixed reference moment: `Mon Jan 2 15:04:05 MST 2006`. Every layout
token is just "whatever that component equals in this reference moment" —
`1`/`01` = month (because the reference month is January, the 1st month),
`2`/`02` = day (because the reference day is the 2nd), `2006` = year, `15` =
hour (24h), `04` = minute, `05` = second, `Monday`/`Mon` = weekday,
`January`/`Jan` = month name. There's no magnitude-based inference (e.g. "13
can't be a month, so must be day") — it's purely positional. Whichever token
you place first in the layout is treated as the first field of the input, no
matter what value is actually there. If the layout says month-first and the
input value can't be a valid month, `Parse` just errors out.

**Q: Why did `time.Parse("Thursday, July 25, 2019 13:45:00", date)` fail with `cannot parse ", 2019 13:45:00" as "5"`?**
Because that layout was a copy of an *example* input, not built from the
reference moment. `Parse` only recognizes the specific magic substrings tied
to `Mon Jan 2 15:04:05 MST 2006` (`Monday`, `January`, `2`, `2006`, `15:04:05`,
etc.) — anything else in the layout is treated as literal text the real input
must match character-for-character. `Thursday`, `July`, `25`, `2019`,
`13:45:00` aren't recognized tokens, so Go tried to literally match them
against the real input and broke as soon as the characters diverged.
Correct layout for `"Thursday, July 25, 2019 13:45:00"`:
`"Monday, January 2, 2006 15:04:05"` (see `IsAfternoonAppointment`, line 43).

**Q: How do you get the hour/weekday out of a parsed time string?**
Once you have a `time.Time` (from `time.Parse`), call its accessor methods:
`t.Hour()`, `t.Minute()`, `t.Second()`, `t.Day()`, `t.Month()`, `t.Year()`,
`t.Weekday()`. `Weekday()` isn't extracted from the string — it's *computed*
from the date, since the input string may not even contain a weekday name.
`Month()` and `Weekday()` return their own named types (`time.Month`,
`time.Weekday`), not plain ints/strings.

**Q: How to compare a given time against the current time?**
Use `time.Now()` plus `.Before()` / `.After()` / `.Equal()` — never `==`, `<`,
`>` on `time.Time` directly. `time.Since(t)` (shorthand for
`time.Now().Sub(t)`) gives a `time.Duration`; positive means `t` is in the
past. Used in `HasPassed` (line 30-37):
```go
t, err := time.Parse("January 2, 2006 15:04:05", date)
elapsed := time.Since(t)
return elapsed > 0
```

**Q: How to format a `time.Time` back into a custom string (e.g. `Thursday, July 25, 2019, at 13:45`)?**
`t.Format(layout)` — the reverse of `Parse`, same token language. Build the
layout by mapping each piece of the target output to its reference-moment
token: `Monday` (weekday), `January 2, 2006` (date), `15:04` (24h hour:minute,
no seconds needed if not shown). Note `Parse` and `Format` calls in the same
function commonly need *different* layouts if the input string shape and the
desired output string shape differ (see `Description`, lines 60-74: parses
with `"1/2/2006 15:04:00"` — NB: trailing `00` here is a bug, should be `05`
for seconds — then formats with `"Monday, January 2, 2006"` and `"15:04"`).

**Q: How to build a `time.Time` from scratch (not by parsing a string)?**
`time.Date(year, month, day, hour, min, sec, nsec, loc)` — takes the
individual components directly. `month` must be a `time.Month` constant
(e.g. `time.September`), not a raw int. Used in `AnniversaryDate` (lines
78-81) to build "September 15 of the current year, midnight UTC":
```go
func AnniversaryDate() time.Time {
	year := time.Now().Year()
	return time.Date(year, time.September, 15, 0, 0, 0, 0, time.UTC)
}
```
The dynamic piece (current year) comes from `time.Now().Year()`; everything
else is fixed per the requirement.

**Q: Is there string interpolation like JS template literals in Go?**
No. Closest equivalents: `fmt.Sprintf(layout, args...)` (verb-based
placeholders like `%s`, `%d`, `%02d`) which *returns* a string (the `S`
prefix means "return as string", unlike `Printf` which writes to stdout), or
plain `+` concatenation. For turning a `time.Time` into a string, `t.Format(...)`
is preferred over manual `Sprintf`/concatenation of individual components —
it's the purpose-built tool.

**Known remaining issue (not yet fixed, tests still pass):** `Description`
(line 65) parses with `"1/2/2006 15:04:00"` — the trailing `00` is literal
text, not the seconds token `05`. Also there's a stray double-space bug noted
in the file's own comment (line 83-84): `"on  Monday"` vs `"on Monday"`.
Worth revisiting if more test cases are added.
