# airport_robot.go — 2026-09-28

**Q: Is my `SayHello` through step 1? (first attempt)**
No. There were three problems:
- `LanguageName(langugage string)` took an argument it doesn't need. A greeter already knows its own language, so the signature is `LanguageName() string`.
- `SayHello` hardcoded `g.LanguageName("German")` and built `"Hello %s"` itself. Every language got "German" and "Hello".
- It never called `Greet`, the interface method that produces the greeting.

`SayHello` shouldn't know anything about a particular language. It just asks the greeter for its language name and its greeting.

**Q: Check now (after `LanguageName()` fix)**
`fmt.Sprintf("I can speak %s: %s %s", greeting, name)` had three `%s` verbs but only two arguments. That isn't a compile error, but `go vet` flags it and the output contains `%!s(MISSING)`. `Greet` still wasn't being called, and the variable `greeting` actually held the language name.

**Q: Check now (calling both methods)**
`g.Greet(languageName)` passed the language name ("German") instead of the visitor's name, which would give "Hallo German!". The mix-up came from the misleading parameter name `Greet(langugage string)`. Renaming it to `Greet(visitorName string)` made the right argument obvious.

**Q: How do I add a function signature in a struct?**
You don't. Go structs hold only fields. Methods are declared outside the struct, at the top level of the file, using a **receiver**:

```go
type Italian struct{}

func (i Italian) LanguageName() string { ... }
```

`(i Italian)` attaches the method to `Italian`, and `i` works like `this`/`self`. Go has no `implements` keyword: a type satisfies an interface once it has every method with the exact signature. A struct also can't have a field and a method with the same name, which is why the earlier `LanguageName string` field clashed.

**Q: Check now (tests failing with "I can speak Ciao: Ciao Flora")**
`LanguageName()` returned the greeting word ("Ciao") and `Greet` borrowed it from there. Those are two separate facts, and each method should own one: `LanguageName()` returns "Italian", and `Greet` returns "Ciao …". The `!` also belongs in `Greet` (the spec says `Greet` returns `"Hallo {name}!"`), not in `SayHello`. The last failure was a typo: `"Protuguese"`.

**Solution and how it works**

```go
type Greeter interface {
	LanguageName() string
	Greet(visitorName string) string
}

func SayHello(visitorName string, g Greeter) string {
	languageName := g.LanguageName()
	greeting := g.Greet(visitorName)
	return fmt.Sprintf("I can speak %s: %s", languageName, greeting)
}
```

`SayHello` depends only on the `Greeter` interface. `Italian{}` and `Portuguese{}` are empty structs, because they have no data to store, and each one implements both methods (lines 18–36). Adding German later means a new type with those two methods, and `SayHello` doesn't change.

**Q: Can I write one `Greet` for several types, like `func (l Portuguese Italian) Greet(...)` with a switch?**
No. A method has exactly one receiver type, so each type needs its own methods. There are two alternatives:
- **One struct holding the data.** For example, store the language name and greeting word as fields, and write `LanguageName()`/`Greet()` once. Each language becomes a value instead of a type.
- **A type switch** (`switch v := g.(type)`) inside a function that takes an interface. This works against the design, though: every new language would mean editing that switch.

Rule of thumb: if types differ only in their data, use one struct with fields. If their behaviour differs, use separate types behind an interface.

**Q: Can `Greet` be generic and pick the string based on the language?**
Yes, with a **map**, e.g. `map[string]string` from language name to greeting word, looked up with the key from `LanguageName()`. Use the comma-ok form `v, ok := m[key]` to handle a language that isn't in the map.

Go generics (`[T any]`) don't help here. Methods can't have their own type parameters, and generics are for the *same logic across many types*. This case is *different data depending on a value*, which is a job for a map or struct fields.
