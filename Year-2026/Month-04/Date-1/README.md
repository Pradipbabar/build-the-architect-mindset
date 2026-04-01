# Day 3 (Week 2) — 1 April 2026

**Phase**: Phase 1 / Module 1.1 — Go Fundamentals  
**Focus**: TGPL Chapter 4 + Finish Gophercises #2 (Quiz Game)  
**Time Target**: 2.5 hours  
**Week**: Week 2 of Phase 1 (March 30 – April 5, 2026)  

---

## Context: Where You Are in Week 2

By end of Tuesday you completed:
- ✅ TGPL Chapter 1 — typed all examples, notes written
- ✅ TGPL Chapter 2 — program structure, zero values, pointers, named types
- ✅ TGPL Chapter 3 — integers, floats, strings, runes, iota, float comparison
- ✅ Gophercises #1 (Echo CLI) — done, table-driven tests written, pushed
- 🔄 Gophercises #2 (Quiz Game) — started, basic version in progress

You are **ahead of the original week plan**. Chapter 3 done on Tuesday means today is fully free for Chapter 4 — the most important chapter in the first half of the book — and finishing Gophercises #2 completely, including the timer stretch goal.

**Running total check** (update honestly):
- Hours logged so far this week: ___ / 12
- Phase 1 total: ___ / 200 hours

---

## Today's Big Picture

**Wednesday target**: 2.5 hours across two sessions:
1. **TGPL Chapter 4** — Composite Types: arrays, slices, maps, structs (90 min)
2. **Finish Gophercises #2** — Quiz Game with timer + tests (60 min)

**Before you start anything:**

```bash
gofmt -l .           # Should return nothing
go vet ./...         # Should return nothing
go test -cover ./... # Know your current baseline
```

Fix any issues before touching new code.

---

## Session 1: TGPL Chapter 4 — Composite Types (90 min)

This is the most important chapter in the first half of TGPL. Slices and maps appear in almost every Go program you will ever write. Do not rush it.

### The Rule
**Type every code example.** Run every example. Change something and predict the output before running.

### Key Topics in Chapter 4

| Concept | What to Understand |
|---|---|
| Arrays | Fixed-size, value type — copying an array copies all elements |
| Slices | Reference to an underlying array — three fields: pointer, length, capacity |
| `append()` | What happens when capacity is exceeded |
| `nil` slice vs empty slice | Both have `len=0` but they are not the same |
| Maps | Reference type, nil map vs empty map, two-value lookup |
| Struct embedding | Go's composition model — there is no inheritance |
| Struct tags | Used for JSON encoding/decoding |

### Critical Concepts — Work Through Each One

**1. Arrays are value types**
```go
a := [3]int{1, 2, 3}
b := a           // b is a full copy, not a reference
b[0] = 99
fmt.Println(a)   // What prints? Try it.
fmt.Println(b)
```
Arrays are rarely used directly in Go. Slices are used instead.

**2. Slices are not arrays**

A slice is a three-field struct internally:
- Pointer to the underlying array
- Length (`len`)
- Capacity (`cap`)

```go
s := []int{1, 2, 3}
t := s           // t points to the SAME underlying array
t[0] = 99
fmt.Println(s)   // What prints? Try it.
fmt.Println(t)
```

**3. `append()` and capacity growth**
```go
s := []int{1, 2, 3}
fmt.Println(len(s), cap(s)) // 3, 3

s = append(s, 4)
fmt.Println(len(s), cap(s)) // What happened to cap? Try it.
```
When `append` exceeds capacity, Go allocates a new underlying array. The original slice variable is unaffected — only the returned slice has the new array. This is a common source of bugs.

**4. nil slice vs empty slice**
```go
var s []int      // nil slice
t := []int{}     // empty slice

fmt.Println(len(s), len(t))   // both 0
fmt.Println(s == nil)          // true
fmt.Println(t == nil)          // false
```
Both are safe to `range` over and `append` to. The difference matters when you need to distinguish "not set" from "set but empty" (e.g., JSON encoding).

**5. Map — nil map vs empty map**
```go
var m map[string]int   // nil map — reading is safe, WRITING panics
m["key"] = 1           // panic: assignment to entry in nil map

m = make(map[string]int) // now safe to write
m["key"] = 1
```

**6. Two-value map lookup**
```go
v, ok := m["key"]
if !ok {
    // key does not exist — v is the zero value
}
```
Always use the two-value form when the key might not exist.

**7. Struct embedding**
```go
type Animal struct {
    Name string
}

type Dog struct {
    Animal        // embedded — Dog promotes Name and any Animal methods
    Breed string
}

d := Dog{Animal: Animal{Name: "Rex"}, Breed: "Labrador"}
fmt.Println(d.Name) // promoted field — no need for d.Animal.Name
```
This is Go's composition model. There is no class hierarchy, no `extends`.

### Where to Take Notes

Create file: `notes/books/tgpl-chapter-4.md`

```
# TGPL Chapter 4 — Composite Types

## Arrays
- [Your explanation: value type, fixed size, rarely used directly]

## Slices
- Internal structure: [pointer, len, cap — in your own words]
- nil slice vs empty slice: [your explanation]
- append() and capacity growth: [what happens when cap is exceeded?]
- Slice tricks: s[i:j], copy(), nil check

## Maps
- nil map vs empty map: [your explanation]
- Two-value lookup: [your explanation]
- Iteration order: [what does Go guarantee?]

## Structs
- Embedding: [your explanation with example]
- Promoted fields and methods: [your explanation]
- Struct tags: [what are they used for?]

## Surprises / Confusions
- [What surprised you?]
- [What do you need to revisit?]
```

**Time allocation**:
- Reading + typing examples: 60 min
- Experimenting (change things, predict output): 20 min
- Taking notes: 10 min

---

## Session 2: Finish Gophercises #2 — Quiz Game (60 min)

You started the basic version on Tuesday. Today: complete it fully, add the timer, write tests, push.

### Step 1: Verify Basic Version Works (10 min)

```bash
go run main.go -csv problems.csv
```

Check manually:
- Questions print correctly
- Correct answers increment the count
- Final score prints: `"You scored X out of Y"`

Fix any bugs before adding the timer.

### Step 2: Add the `-limit` Flag and Timer (20 min)

```go
timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
flag.Parse()

timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
```

Inside your quiz loop, use a goroutine + channel so stdin does not block the timer:

```go
answerCh := make(chan string, 1)

go func() {
    var answer string
    fmt.Scan(&answer)
    answerCh <- answer
}()

select {
case <-timer.C:
    fmt.Printf("\nTime's up! You scored %d out of %d.\n", correct, total)
    return
case answer := <-answerCh:
    answer = strings.TrimSpace(strings.ToLower(answer))
    if answer == strings.TrimSpace(strings.ToLower(p.a)) {
        correct++
    }
}
```

Test it:
```bash
go run main.go -csv problems.csv -limit 5
# Answer slowly — timer should cut you off at 5 seconds
```

### Step 3: Write Tests (20 min)

Separate your pure logic from I/O so it is testable:

```go
func TestParseProblems(t *testing.T) {
    tests := []struct {
        name    string
        records [][]string
        want    []problem
    }{
        {
            name:    "basic parsing",
            records: [][]string{{"5+5", "10"}, {"1+1", "2"}},
            want:    []problem{{q: "5+5", a: "10"}, {q: "1+1", a: "2"}},
        },
        {
            name:    "trims answer whitespace",
            records: [][]string{{"5+5", " 10 "}},
            want:    []problem{{q: "5+5", a: "10"}},
        },
        {
            name:    "empty input",
            records: [][]string{},
            want:    []problem{},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := parseProblems(tt.records)
            // your assertion
        })
    }
}
```

### Step 4: Final Checks and Push (10 min)

```bash
go vet ./...
go test -v -cover ./...
git add .
git commit -m "week-2: gophercises-02 quiz game complete with timer + tests"
git push
```

---

## Time Breakdown

| Task | Target | Actual |
|---|---|---|
| Session 1: TGPL Ch. 4 | 90 min | ___ |
| Session 2: Finish Gophercises #2 | 60 min | ___ |
| **Total** | **2.5 hours** | ___ |

---

## Concepts to Verify From Yesterday (Chapter 3)

Answer these without looking anything up. If you cannot, spend 10 minutes reviewing before Session 1:

1. What does integer overflow do in Go? Does it panic?
2. What is a `rune`? How is it different from a `byte`?
3. Why does `len("日本語")` not return `3`?
4. What does `iota` do inside a `const` block? What happens when a new `const` block starts?
5. Why should you never compare `float64` values with `==`?

---

## Notes & Reflections (Fill In As You Work)

### Things that made sense immediately

-
-
-

### Things that confused me

-
-

### Questions to research later

-
-

---

## End of Day Checklist

- [ ] Chapter 4 read completely + all examples typed and run
- [ ] Experimented with slice sharing, `append()` capacity growth, nil map write panic
- [ ] Chapter 4 notes written in `notes/books/tgpl-chapter-4.md`
- [ ] Gophercises #2 basic version verified working
- [ ] Timer (`-limit` flag) implemented and manually tested
- [ ] `parseProblems` separated and independently testable
- [ ] At least 3 table-driven tests written and passing
- [ ] `go vet ./...` passes with no warnings
- [ ] `go test -cover ./...` run — coverage noted
- [ ] Code pushed to GitHub with a meaningful commit message
- [ ] Hours logged accurately — be honest

---

## Actual Time Spent & Summary

**Total hours: ___ (Target: 2.5)**

**What I accomplished**:

**What I struggled with**:

**What to revisit tomorrow (Thursday)**:

---

## Tomorrow: Thursday, April 2

**Focus**: TGPL Chapters 5–6 (Functions + Methods) + Gophercises #3 and #4.

Functions and methods are where Go's design philosophy becomes very clear — first-class functions, multiple return values, variadic functions, defer, and the method receiver pattern. Give Chapter 5 the same care you gave Chapter 4 today.