# Go Pointers — Complete Guide

> **Tags:** #go #backend #pointers #learning
> **Status:** #reference
> **Related:** [[Go Basics]] [[Go Structs]] [[Go Methods]]

---

## What is a Pointer?

A pointer is a variable that holds a **memory address** instead of a value directly.

Think of memory as a row of numbered boxes:
- A **regular variable** stores data *in* a box
- A **pointer** stores the *address* of a box

```go
x := 42      // x is a box containing 42
p := &x      // p is a box containing the address of x
```

```mermaid
flowchart LR
    subgraph Memory
        direction LR
        boxX["Address: 0xc0001<br/>Variable: x<br/>Value: 42"]
        boxP["Address: 0xc0002<br/>Variable: p<br/>Value: 0xc0001"]
    end
    boxP -- "points to" --> boxX
```

`p` doesn't hold `42` — it holds the *address* of the box that holds `42`. The arrow is the pointer.

---

## The Three Core Operators

| Syntax | Name | What it does |
|--------|------|--------------|
| `&x` | Address-of | Returns the memory address of `x` |
| `*int` | Pointer type | Declares a type: "pointer to int" |
| `*p` | Dereference | Reads or writes the value at the address `p` holds |

---

## Step-by-Step Basics

### Step 1 — Regular variable

```go
x := 42
fmt.Println(x) // 42
```

### Step 2 — Get an address with `&`

```go
x := 42
p := &x                   // p holds the address of x
fmt.Println(p)            // 0xc0000b4010 (some address)
fmt.Printf("%T\n", p)     // *int
```

### Step 3 — Dereference with `*`

```go
x := 42
p := &x
fmt.Println(*p) // 42  ← follows the address, reads the value
```

### Step 4 — Modify through a pointer

```go
x := 42
p := &x

*p = 100           // change the value AT the address
fmt.Println(x)     // 100 ← x is now changed!
```

```mermaid
sequenceDiagram
    participant Code as Your code
    participant x as x (0xc0001)
    participant p as p (0xc0002)

    Code->>x: x := 42
    Note over x: value = 42
    Code->>p: p := &x
    Note over p: value = 0xc0001 (address of x)
    Code->>p: *p = 100
    p->>x: writes 100 to address 0xc0001
    Note over x: value = 100
    Code->>x: fmt.Println(x)
    Note over Code: prints 100
```

### Step 5 — Nil pointers

A pointer that points to nothing is `nil`. Dereferencing it will **panic**.

```go
var p *int
fmt.Println(p)   // <nil>
fmt.Println(*p)  // PANIC: runtime error: nil pointer dereference

// Always guard:
if p != nil {
    fmt.Println(*p)
}
```

```mermaid
flowchart TD
    A["var p *int"] --> B["p points to: nil<br/>(no address)"]
    B --> C{"*p ?"}
    C -->|"p == nil"| D["PANIC:<br/>nil pointer dereference"]
    C -->|"p != nil"| E["safe: read/write<br/>the value at p's address"]
```

---

## Pointers with Functions

Without a pointer, Go passes a **copy** of the value. The original is not changed.

```go
func double(n int) {
    n = n * 2  // modifies the copy, not the original
}

x := 5
double(x)
fmt.Println(x) // still 5
```

With a pointer, the function can modify the **original**:

```go
func double(n *int) {
    *n = *n * 2  // modifies the value at the address
}

x := 5
double(&x)
fmt.Println(x) // 10 ✓
```

```mermaid
flowchart LR
    subgraph without["Without pointer — pass by value"]
        direction TB
        x1["x = 5<br/>(0xc0001)"] -- "copies value" --> n1["n = 5<br/>(0xc0002, own box)"]
        n1 -- "n = n * 2" --> n1b["n = 10<br/>(only this box changes)"]
        x1b["x is still 5"]
    end
```

```mermaid
flowchart LR
    subgraph withp["With pointer — pass by address"]
        direction TB
        x2["x = 5<br/>(0xc0001)"] -- "copies the address" --> n2["n = 0xc0001<br/>(0xc0003, points back to x)"]
        n2 -- "*n = *n * 2<br/>writes through the address" --> x2b["x = 10<br/>(same box as x2, now updated)"]
    end
```

The key difference: passing a value copies the *data*. Passing a pointer copies the *address*, but both point back to the same original box — so writes through the pointer are visible to the caller.

---

## Pointers with Structs

```go
type Person struct {
    Name string
    Age  int
}

// Takes a COPY — original unchanged
func birthday(p Person) {
    p.Age++
}

// Takes a POINTER — modifies the original
func birthdayPointer(p *Person) {
    p.Age++ // shorthand for (*p).Age++ — Go auto-dereferences struct pointers
}

func main() {
    person := Person{Name: "Prathamesh", Age: 28}

    birthday(person)
    fmt.Println(person.Age) // 28 — unchanged

    birthdayPointer(&person)
    fmt.Println(person.Age) // 29 ✓
}
```

> **Note:** `p.Age++` inside a pointer function is Go's syntactic sugar for `(*p).Age++`. You don't need to write the `*` explicitly for struct fields.

```mermaid
flowchart TD
    person["person: Person{Name: Prathamesh, Age: 28}<br/>(original, 0xc0001)"]

    person -- "birthday(person)<br/>passes a COPY" --> copy["copy inside birthday()<br/>Age becomes 29<br/>(0xc0002, thrown away on return)"]
    person -- "birthdayPointer(&person)<br/>passes the ADDRESS" --> same["p *Person = 0xc0001<br/>(*p).Age++ writes directly<br/>into the original"]
    same -.->|"mutates"| person
```

---

## `new()` vs `&T{}`

Two ways to create a heap-allocated pointer:

### `new(T)` — allocates and zeroes

```go
p := new(int)
fmt.Println(*p) // 0  ← zero value for int

p2 := new(Person)
fmt.Println(p2.Name) // ""
fmt.Println(p2.Age)  // 0
```

### `&T{}` — allocates and initialises (preferred for structs)

```go
person := &Person{Name: "Prathamesh", Age: 28}
fmt.Println(person.Name) // Prathamesh
```

> **Rule of thumb:** Use `new()` for primitives, `&T{}` for structs.

---

## Pointer Receivers on Methods

This is how Go handles object-oriented-style mutation.

```go
type Counter struct {
    count int
}

// Value receiver — gets a COPY, original unchanged
func (c Counter) IncWrong() {
    c.count++
}

// Pointer receiver — operates on the ORIGINAL
func (c *Counter) Inc() {
    c.count++
}

func (c *Counter) Value() int {
    return c.count
}

func main() {
    c := Counter{count: 0}

    c.IncWrong()
    fmt.Println(c.count) // 0 — unchanged

    c.Inc()
    c.Inc()
    fmt.Println(c.count) // 2 ✓
}
```

> **Rule of thumb:** If *any* method on a type needs to mutate state, use pointer receivers on *all* methods of that type for consistency.

```mermaid
flowchart LR
    c["c := Counter{count: 0}"]
    c -- "c.IncWrong()<br/>value receiver → COPY" --> copy["copy.count = 1<br/>(discarded)"]
    c -- "c.Inc()<br/>pointer receiver → same struct" --> c2["c.count = 1<br/>(persists)"]
    c2 -- "c.Inc() again" --> c3["c.count = 2<br/>(persists)"]
```

> Go automatically takes the address for you when you call a pointer-receiver method on an addressable value — that's why `c.Inc()` works even though `c` isn't `&c`.

---

## Stack vs Heap & Escape Analysis

Go manages memory automatically. You don't choose stack vs heap — the **compiler decides** using escape analysis.

| Location | Characteristics |
|----------|----------------|
| **Stack** | Fast, auto-freed when function returns |
| **Heap** | GC-managed, lives longer, slight overhead |

```go
// Value stays on the STACK — never leaves the function
func localOnly() {
    x := 42
    fmt.Println(x) // x is gone after this returns
}

// Value ESCAPES to the HEAP — pointer outlives the function
func newPerson() *Person {
    p := &Person{Name: "Prathamesh"} // heap-allocated
    return p                          // safe! still valid after return
}
```

### See escape analysis yourself

```bash
go build -gcflags="-m" main.go
# Example output:
# ./main.go:8:7: &Person{...} escapes to heap
```

> Go's garbage collector handles freeing heap memory. You don't need to `free()` anything like in C.

```mermaid
flowchart TD
    A["Compiler sees a value"] --> B{"Does a pointer to it<br/>escape the function?<br/>(returned, stored globally,<br/>sent to a channel, etc.)"}
    B -->|"No"| C["Stack<br/>fast, freed automatically<br/>when function returns"]
    B -->|"Yes"| D["Heap<br/>GC-managed,<br/>lives as long as needed"]

    C --> C1["localOnly(): x := 42<br/>never leaves → stack"]
    D --> D1["newPerson(): p := &Person{...}; return p<br/>outlives the function → heap"]
```

This decision is called **escape analysis**, and it happens entirely at compile time — you never write `stack` or `heap` in your code.

---

## Pointers with Slices and Maps

Slices and maps in Go are already **reference-like** — internally a slice is a small struct holding a pointer to an underlying array, a length, and a capacity. This is why you rarely need `*[]int` or `*map[string]int`.

```go
func appendOne(s []int) []int {
    return append(s, 99)
}

nums := []int{1, 2, 3}
nums = appendOne(nums) // must reassign — append may return a NEW underlying array
```

```mermaid
flowchart LR
    subgraph slice["slice header (copied by value)"]
        direction TB
        ptr["ptr → underlying array"]
        len["len = 3"]
        cap["cap = 3"]
    end
    slice --> arr["[1, 2, 3]<br/>(shared until capacity is exceeded)"]
```

- Mutating an **existing element** (`s[0] = 100`) is visible to the caller — no pointer needed, because both copies of the slice header point at the same underlying array.
- **Appending past capacity** allocates a *new* underlying array — the caller's old slice header still points at the old array, which is why `append` must be reassigned.
- A pointer to a slice (`*[]int`) is only needed if you want a function to change the caller's length/cap/pointer itself (rare — e.g. a `reset()` helper).

Maps are always reference types too — passing a map to a function never needs a pointer to mutate its contents.

---

## Pointers to Pointers (`**T`)

Rare, but useful to recognise — a pointer can itself point to another pointer.

```go
x := 42
p := &x    // *int  — points to x
pp := &p   // **int — points to p, which points to x

fmt.Println(**pp) // 42 — dereference twice
```

```mermaid
flowchart LR
    pp["pp (**int)"] -- "points to" --> p["p (*int)"]
    p -- "points to" --> x["x = 42"]
```

You'll mostly see `**T` when a function needs to reassign the caller's pointer itself (e.g. replacing what `p` points to, not just the value at `*p`) — uncommon in everyday Go code, but shows up in some C-interop and low-level library code.

---

## Common Mistakes

### ❌ Mistake 1: Loop variable pointer trap

```go
// WRONG — all pointers point to the same variable i
func badPointers() []*int {
    result := []*int{}
    for i := 0; i < 3; i++ {
        result = append(result, &i) // &i is always the same address!
    }
    // All three dereference to 3 (the final value of i)
    return result
}

// CORRECT — capture a copy each iteration
func goodPointers() []*int {
    result := []*int{}
    for i := 0; i < 3; i++ {
        v := i               // new variable, new address
        result = append(result, &v)
    }
    return result
}
```

```mermaid
flowchart TD
    subgraph wrong["WRONG — one shared box"]
        direction TB
        i["i (single address, 0xc0001)<br/>reused every iteration:<br/>0 → 1 → 2 → 3"]
        p0["result[0]"] --> i
        p1["result[1]"] --> i
        p2["result[2]"] --> i
        note1["All three point at the SAME box.<br/>After the loop, i = 3, so all dereference to 3."]
    end
```

```mermaid
flowchart TD
    subgraph correct["CORRECT — fresh box each time"]
        direction TB
        v0["v (0xc0010) = 0"]
        v1["v (0xc0011) = 1"]
        v2["v (0xc0012) = 2"]
        q0["result[0]"] --> v0
        q1["result[1]"] --> v1
        q2["result[2]"] --> v2
        note2["Each iteration declares a NEW v,<br/>so each pointer has its own box."]
    end
```

> **Note:** In Go 1.22+, loop variables are re-created each iteration, so this trap no longer applies for range loops. But it still applies for classic `for i := 0; ...` loops.

### ❌ Mistake 2: Nil pointer dereference

```go
// WRONG
func risky(p *int) {
    fmt.Println(*p) // PANIC if p is nil
}

// CORRECT
func safe(p *int) {
    if p == nil {
        fmt.Println("no value provided")
        return
    }
    fmt.Println(*p)
}
```

### ❌ Mistake 3: Over-using pointers for performance

```go
// Small structs — passing by value is often FASTER
// (pointer indirection causes cache misses)
type Point struct{ X, Y float64 }

func distanceByValue(p Point) float64 { ... }   // ✓ fine for small structs
func distanceByPointer(p *Point) float64 { ... } // unnecessary indirection

// Large structs — pointer is better
type BigConfig struct { /* 50+ fields */ }
func process(cfg *BigConfig) { ... } // ✓ avoids a big copy
```

---

## Full Cheat Sheet

```go
// --- Declarations ---
var p *int            // nil pointer to int
p := &x               // pointer to existing variable x
p := new(int)         // pointer to new zeroed int on heap
p := &Person{Age: 28} // pointer to new initialised struct on heap

// --- Reading & Writing ---
*p                    // dereference: get the value
*p = 42               // dereference: set the value
p.Name                // shorthand for (*p).Name on structs

// --- Checking ---
p == nil              // true if pointer is unset
p != nil              // safe to dereference

// --- Function signatures ---
func read(p *int) int { return *p }        // read via pointer
func mutate(p *int)   { *p = 99 }         // mutate via pointer
func (c *Counter) Inc() { c.count++ }     // pointer receiver method

// --- Printing ---
fmt.Println(p)        // prints the address: 0xc000...
fmt.Println(*p)       // prints the value at that address
```

---

## When to Use Pointers

| Use a pointer when... | Use a value when... |
|-----------------------|---------------------|
| You need to mutate the original | You only need to read |
| The struct is large (many fields) | The struct is small (2–3 fields) |
| You want to express "optional" (`nil`) | The value is always present |
| Implementing interface methods that mutate | Pure functions / computations |

---

## Quick Mental Model

```
&  →  "take me TO the address"     (variable → address)
*  →  "follow the address BACK"    (address → value)

They are opposites.
```

---

## Practice Exercises

Try these before moving on — they cover the traps above:

1. Write a `func swap(a, b *int)` that swaps the values at two addresses. Call it on two variables and confirm they swapped.
2. Write a `Stack` struct backed by a `[]int` with pointer-receiver methods `Push(v int)` and `Pop() (int, bool)`. Why must the receiver be a pointer here?
3. Predict the output, then run it and check yourself:
   ```go
   type Box struct{ Value int }
   func change(b Box)   { b.Value = 99 }
   func changeP(b *Box) { b.Value = 99 }

   box := Box{Value: 1}
   change(box)
   fmt.Println(box.Value)  // ?
   changeP(&box)
   fmt.Println(box.Value)  // ?
   ```
4. Run `go build -gcflags="-m"` on a small file with both a stack-only and a heap-escaping function. Confirm which lines the compiler reports as "escapes to heap".

---

## Related Topics to Explore Next

- [x] Pointers inside slices and maps
- [x] Pointers to pointers (`**T`)
- [ ] Interfaces with pointer types
- [ ] Concurrency and pointer safety (`sync.Mutex`)
- [ ] Unsafe pointers (`unsafe.Pointer`) — advanced
- [ ] Go garbage collector internals
