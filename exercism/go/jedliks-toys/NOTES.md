# jedliks_toys.go — 2026-09-25

**Q: Where do I start with `Drive`? I'd written it as `func (car Car) Drive() string` returning a `Sprintf` of the car.**
Three things were off:
- `Car` is already declared in `car.go` (same package). Declaring it again in `jedliks_toys.go` is a compile error, so only the methods belong there.
- The README's `// car is now Car{speed: 5, ... distance: 5}` is a comment describing the car's **state** after the call, not a return value.
- A value receiver `(car Car)` works on a **copy**, so any changes are thrown away when the method returns.

**Q: What should `Drive` return, actually?**
Nothing. The test calls `car.Drive()` on its own line and then compares the `car` struct to the expected one. `Drive` is a state-changing method, so it has no return type and uses a pointer receiver:

```go
func (car *Car) Drive() {
```

Compare with `DisplayDistance` / `DisplayBattery`, which *report* something and so do return a `string`.

**Q: But how do I update it?**
Assign to the fields through the receiver, like the README's `squareIt` example (`r.height = r.width`). Because `car` is a pointer, `car.distance = ...` changes the caller's real car. Go dereferences automatically, so you don't need `(*car).distance`.

**Q: Why does `./run-tests.sh jedliks-toys TestDrive -v` still say `[build failed]`?**
`-run` only filters which tests **execute**. Go still compiles the whole package first, including every test in `jedliks_toys_test.go`. Those tests reference `DisplayDistance`, `DisplayBattery` and `CanFinish`, so until those methods exist (even as `panic("")` stubs), nothing runs. A package compiles as a whole or not at all.

**Q: Getting the `Drive` guard right: bugs hit along the way**
- `car.distance = car.distance` / `car.distance + car.distance`: self-assignment or self-addition. A new car starts at 0, so it never moves. Distance grows by `car.speed`.
- Distance update placed **before** the `if`: the car "moved" even when it had no battery. Both updates must be inside the guard.
- `if car.battery > 1` / `postBatteryDrain > 1`: hard-coded numbers ignore each car's own drain, and `> 1` wrongly blocks cars that would end at 0% or 1%.
- `postBatteryDrain > 0` blocks battery 2 / drain 2, but ending at exactly 0% is allowed. The correct check is `>= 0`.

Final version:

```go
func (car *Car) Drive() {
	postBatteryDrain := car.battery - car.batteryDrain
	if postBatteryDrain >= 0 {
		car.battery = car.battery - car.batteryDrain
		car.distance = car.distance + car.speed
	}
}
```

**Q: With speed 5, why does `DisplayDistance` on a new car print `Driven 0 meters`? Shouldn't distance increase?**
No. `speed` is how far the car moves **per call to `Drive()`**, not how far it has already gone. `NewCar` doesn't set `distance`, so it has `int`'s zero value, `0`. The example never calls `Drive()`, so the car is still at the start line.

| Code | distance |
|---|---|
| `NewCar(5, 2)` | 0 |
| `car.Drive()` | 5 |
| `car.Drive()` | 10 |

**Q: `DisplayDistance` check**
Correct:

```go
func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}
```

It only reads the car, so a value receiver is fine. (Some people use a pointer receiver on every method of a type once any method needs one, for consistency.)

**Q: `DisplayBattery`: why didn't `"Battery at %d%"` or `"Battery at %d" + "%"` work?**
In `Sprintf`, `%` starts a format verb, so a lone trailing `%` is an error (`go vet`: `format % is missing verb at end of string`). `"%d" + "%"` is joined into `"%d%"` **before** `Sprintf` sees it, so it's the same broken string. A literal percent sign is written `%%`:

```go
return fmt.Sprintf("Battery at %d%%", car.battery)
```

The method name also had a typo at first (`DisaplyBattery`), which is what kept the whole package from building.

**Q: For `CanFinish`, with speed 5 and drain 2, does each drive cover 5 meters?**
Yes. One `Drive()` = `speed` meters for `batteryDrain` percent. `CanFinish` asks: with the battery it has **now** (`car.battery`, not `100`), can the car cover `trackDistance`?

**Q: Why did my track-first `CanFinish` fail?**
The approach was: drives needed = `trackDistance / car.speed` → battery used → battery left `>= 0`.
- **Integer division rounds down**, which under-counts drives needed. Track 61, speed 3 → `61 / 3 = 20` drives, but 20 × 3 = 60 m, 1 m short. The car really needs 21.
- Adding the remainder check with `||` (`left >= 0 || pendinDistance == 0`) let any evenly divisible track pass even with a dead battery (battery 30, track 100, speed 5 → true, expected false).
- Switching to `&&` (`left >= 0 && pendinDistance == 0`) passes all 4 tests but is still wrong: **any** track not divisible by speed returns false, even with a full battery (speed 5, drain 2, battery 100, track 7 → false, should be true). The tests just don't cover that case.

The remainder shouldn't be a pass/fail condition. It means "one more drive needed".

## Solution and how it works

`CanFinish`, battery-first (recommended). Rounding down is *correct* in this direction, because the car can't make a partial drive:

```go
func (car Car) CanFinish(trackDistance int) bool {
	drivesLeft := car.battery / car.batteryDrain
	maxDistance := drivesLeft * car.speed
	return maxDistance >= trackDistance
}
```

| speed | drain | battery | track | drives | max distance | result |
|---|---|---|---|---|---|---|
| 5 | 2 | 100 | 100 | 50 | 250 | true |
| 5 | 2 | 40 | 100 | 20 | 100 | true |
| 3 | 3 | 60 | 61 | 20 | 60 | false |
| 5 | 2 | 30 | 100 | 15 | 75 | false |
| 5 | 2 | 100 | 7 | 50 | 250 | true |

Track-first alternative: round **up** when there's a remainder, and check only the battery:

```go
drivesNeeded := trackDistance / car.speed
if trackDistance%car.speed > 0 {
	drivesNeeded++ // a partial drive still costs a full drive
}
return car.battery >= drivesNeeded*car.batteryDrain
```

Takeaway: when using integer division, check which direction the rounding goes. Rounding down is safe when counting what you **can afford** (drives the battery allows), and wrong when counting what you **need** (drives the track requires).
