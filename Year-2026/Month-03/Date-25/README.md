# Day 3 — 25 March 2026

**Phase**: Phase 1 / Module 1.1 — Go Fundamentals  
**Focus**: Tour of Go — Methods and Interfaces  

---

## Context: Why Today Matters

Most Go beginners rush through Methods and Interfaces, get confused later,
and then spend weeks unlearning bad habits. This section is where Go starts
to feel different from other languages.

Do not rush it. One topic at a time. Full understanding before moving on.

---

## Today's Single Goal

Complete the **Methods and Interfaces** section of the Tour of Go.

> https://go.dev/tour/methods/

Do not move to the Concurrency section today. That is Day 4 or Day 5.

---

## What You Will Cover

The Methods and Interfaces section covers these topics, in order:

1. Methods (functions with a receiver)
2. Methods are functions (the subtle difference)
3. Pointer receivers vs value receivers
4. Pointers and interfaces
5. Interfaces (implicit satisfaction — no `implements` keyword)
6. Interface values (dynamic type and value)
7. Interface values with nil underlying values
8. Nil interface values
9. The empty interface (`interface{}` / `any`)
10. Type assertions
11. Type switches
12. Stringer interface (`fmt.Stringer`)
13. Errors (`error` interface)
14. Readers (`io.Reader` interface)
15. Images (`image.Image` interface — optional exercise)

---

## The Concepts That Trip People Up

Read these before you start. They will save you confusion.

### Pointer receivers vs value receivers
When you define a method with a pointer receiver (`*T`), Go can call it on
both a pointer and an addressable value. When you define it with a value
receiver (`T`), it works on both. But if your interface requires a pointer
receiver method, you must pass a pointer — not a value. This is one of the
most common beginner mistakes.

```go
// Value receiver — works on MyFloat directly
type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}
```

```go
// Pointer receiver — you need a *Vertex to satisfy an interface with Scale()
type Vertex struct{ X, Y float64 }

func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}
```

### Interfaces are satisfied implicitly
There is no `implements` keyword in Go. If your type has the methods an
interface requires, it satisfies that interface — automatically. This is
different from Java or C#. The interface and the type do not need to know
about each other.

### The error interface
This is one you will use every single day in Go. Memorise its shape:

```go
type error interface {
    Error() string
}
```

Any type that has an `Error() string` method is an `error`. That is it.

---

## How to Work Through It

### The Rule

### For each topic
1. Read the explanation on the page carefully
2. Type the code in the editor
3. Run it — study the output
4. Ask yourself: *"If someone asked me to explain this in 2 sentences, what would I say?"*
5. If you cannot answer → re-read before moving on
6. Write down what confused you in the Notes section below

### Do not skip exercises
The Tour has exercises marked with a pencil icon. Do them. They are short
and they are where the real understanding happens. Spend up to 15 minutes
on each exercise before moving on. If you are still stuck after 15 minutes,
note it down and continue — come back later.

---

## Code to Write Today

After finishing the Tour section, write this program **from scratch**.
No Tour reference. No AI. From memory.

**File**: `projects/phase-1/00-day3-methods/methods_practice.go`

```
Write a Go program that does the following:

1. Define an interface called Shape with two methods:
   - Area() float64
   - Perimeter() float64

2. Define two structs: Rectangle (width, height float64) and Circle (radius float64)

3. Implement the Shape interface on both Rectangle and Circle
   - Rectangle Area    = width * height
   - Rectangle Perimeter = 2 * (width + height)
   - Circle Area       = π * radius * radius  (use math.Pi)
   - Circle Perimeter  = 2 * π * radius

4. Write a function called printShape(s Shape) that prints:
   - The area and perimeter of the shape
   - The concrete type of the shape (hint: use a type switch)

5. In main(), create one Rectangle and one Circle, call printShape() on each

Allowed imports: "fmt", "math"
```

This tests: interfaces, method receivers, type switches — the core of today.

---

## Stretch Goal (Only If Time Allows)

If you finish early and want to go deeper, implement the `fmt.Stringer`
interface on both `Rectangle` and `Circle` so that `fmt.Println(rect)`
prints something meaningful like `Rectangle(width=4.0, height=3.0)`.

Do not do this at the cost of the main exercise. Main exercise first.

---

## Notes (Fill These In As You Go)

### Things that made sense immediately
- You can declare methods with pointer receivers.
- There are two reasons to use a pointer receiver.

    - The first is so that the method can modify the value that its receiver points to.

    - The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.
- An interface type is defined as a set of method signatures.

### Things that confused me
- You can declare a method on non-struct types, too.
- Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.
-

### Pointer receiver vs value receiver — my understanding in my own words
A value receiver means the method gets a copy of the original value, so anything I change inside the method does not affect the original struct.A pointer receiver means the method gets the actual memory address, so it can modify the original data directly and avoids copying large structs.


### What the error interface looks like — from memory
The error interface is just a type that has one method called Error() which returns a string message describing what went wrong.


### Questions I want to research later
- How to design small, composable interfaces (idiomatic Go style)
- Interface embedding (combining multiple interfaces)
- “Accept interfaces, return structs” pattern
- Creating a wrapper (adapter) around existing types

### One thing I want to remember from today
- Building reusable packages
- Designing flexible helper functions
---

## Time Log

| Session | Start | End | Duration | What I worked on |
|---------|-------|-----|----------|-----------------|
| 1 | | | | Tour — Methods to Interface values |
| 2 | | | | Tour — Type assertions to io.Reader |
| 3 | | | | methods_practice.go from scratch |

**Total today**: ___ hours  
**Running total**: ___ hours / 1500–2000 goal

---

## Concepts Check

Before you close your laptop, answer these without looking anything up.
Write your answers in your notes file, not here.

1. What is the difference between a pointer receiver and a value receiver?
   When would you use each?

2. How does Go know if a type satisfies an interface? Is there a keyword?

3. What is the zero value of an interface type?

4. What does a type assertion look like? What happens if it fails?

5. Write the `error` interface from memory. What method does it require?

If you cannot answer any of these, go back and re-read that specific topic
before calling Day 3 done.

---

## End of Day Checklist

- [ ] Completed all topics in the Methods and Interfaces section
- [ ] Typed every example (did not paste)
- [ ] Attempted all exercises in the Tour
- [ ] Wrote `methods_practice.go` from memory
- [ ] Program runs without errors (`go run methods_practice.go`)
- [ ] Answered the 5 concepts check questions in your notes
- [ ] Filled in the Notes section above
- [ ] Committed code to GitHub with a clear message
- [ ] Logged time honestly

---

## Git Commit Message for Today

```
day-3: tour of go methods and interfaces complete

- worked through methods, pointer receivers, interfaces, type switches
- wrote methods_practice.go from scratch (Shape interface, Rectangle, Circle)
- noted confusions around [fill in what confused you]
```

---

## Day 4 Preview

Tomorrow you move to **Concurrency** — goroutines, channels, and `select`.
This is the section Go is most famous for and where most beginners get lost.

Before Day 4, make sure you are solid on interfaces. Concurrency uses
interfaces heavily (`error`, `io.Reader`, context) and if interfaces are
still fuzzy, concurrency will feel impossible.

If today felt hard, that is fine. Methods and Interfaces is genuinely the
hardest conceptual section in the Tour. Tomorrow gets more concrete.

---

**Commitment check**: Did I write the code myself today? Yes / No  
**Honest hours logged**: ___