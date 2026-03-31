# Day 2 (Week 2) — 31 March 2026

**Phase**: Phase 1 / Module 1.1 — Go Fundamentals  
**Focus**: TGPL Chapters 2–3 + Gophercises #1 Complete  
**Time Target**: 2.5 hours  
**Week**: Week 2 of Phase 1 (March 30 – April 5, 2026)  

---

## Context: Building Momentum

Day 1 (Monday, March 30) gave you your first exposure to TGPL and Gophercises in Week 2. Today builds on that foundation by deepening your understanding of Go's core types and introducing control flow.

**Progress check**: Did you complete Monday's goals?
- [ ] TGPL Chapter 1 read and typed
- [ ] Notes created in `notes/books/TGPL/tgpl-chapter-1.md`
- [ ] Gophercises #1 started (Echo CLI problem understood)

If any of these are incomplete, finish them before moving forward. Week 2 is cumulative.

---

## Today's Big Picture

**Tuesday target**: 2.5 hours across two focused sessions:

1. **TGPL Chapters 2–3** — Program Structure, Data Types, and Values (120 min)
2. **Gophercises #1 Complete** — Echo CLI full solution with tests (30 min)

This is where Go's distinct features become apparent. Pay close attention to how Go handles types and structs—you'll see these patterns everywhere.

---

## Session 1: TGPL Chapters 2–3 — Program Structure & Data Types (120 min)

### Chapter 2: Program Structure

Chapter 2 explains how Go programs are organized: packages, imports, and the `main` package.

**For each section in Chapter 2:**

- [ ] Read the explanation
- [ ] Type every code example
- [ ] Run it: `go run filename.go`
- [ ] In your notes: *"How would I use this in a real program?"*

**Key Topics in Chapter 2 (Focus Here)**

| Concept | What to Understand |
|---|---|
| `package` declaration | Every Go file belongs to a package; `main` is the entry point |
| `import` statements | How to include libraries and your own packages |
| `func main()` | The entry point of any executable |
| `init()` function | Special initialization function that runs before `main()` |
| Exported names | Names starting with uppercase are exported; lowercase are package-private |

---

### Chapter 3: Data Types

Chapter 3 is dense. It covers all of Go's built-in types: integers, floats, strings, booleans, arrays, slices, maps, and structs. This is essential knowledge.

**For each type section in Chapter 3:**

- [ ] Read how the type works
- [ ] Type the example code
- [ ] Test edge cases: *"What happens if I…?"*
- [ ] Write one small example using each type

**Key Topics in Chapter 3 (Watch Closely)**

| Concept | What to Understand |
|---|---|
| `int`, `uint`, `float64` | Go's numeric types and when to use each |
| `string` | Strings are immutable sequences of bytes |
| `bool` | Boolean values and their operators |
| Arrays `[n]T` | Fixed-length sequences; rarely used directly |
| Slices `[]T` | Dynamic arrays; the most common collection type in Go |
| Maps `map[K]V` | Key-value pairs; Go's dictionary/hash table |
| Structs | Composite types grouping related data; foundation for Go's design |
| Pointers `*T` | References to values; Go uses them carefully (no pointer arithmetic) |

### Where to Take Notes

Update or create: `notes/books/TGPL/tgpl-chapter-2-3.md`

Structure:
```
# TGPL Chapters 2–3 — Program Structure & Data Types

## Chapter 2: Program Structure

### `package` declaration
- [Your explanation in your own words]

### `func main()`
- [Your explanation]

### Exported vs Unexported Names
- [Your explanation]

## Chapter 3: Data Types

### Numeric Types (int, uint, float64)
- [Your explanation with examples]

### Strings
- [Your explanation with examples]

### Arrays vs Slices
- [Key difference you'll remember]

### Maps
- [Your explanation with examples]

### Structs
- [Your explanation with examples]

### Pointers
- [Your explanation: Why does Go use them?]

## Examples I Typed
1. [Example description]
2. [Example description]
...

## Questions & Confusions
- [What parts are still unclear?]
- [What do I need to revisit?]
```

**Time allocation**:
- Reading Chapter 2: 20 min
- Typing Chapter 2 examples: 25 min
- Reading Chapter 3: 35 min
- Typing Chapter 3 examples: 30 min
- Taking notes: 10 min

---

## Session 2: Gophercises #1 — Complete the Echo CLI (30 min)

By now, you should understand the problem. Today is about completing a working solution.

### The Task (if you haven't started)

Gophercises #1: Echo — Create a command-line tool that:
- Takes command-line arguments
- Outputs them with a configurable separator (default: space)
- Has a `-n` flag to add a newline at the end (or `-N` to suppress it)

Example usage:
```bash
go run . hello world        # hello world
go run . -n hello world     # hello world\n
go run . -s=, hello world   # hello,world
```

### Approach

1. **Parse command-line arguments** using `flag` package
2. **Handle the `-n`, `-N`, and `-s` flags** appropriately
3. **Build your solution incrementally**:
   - Step 1: Read arguments and print them
   - Step 2: Add separator support
   - Step 3: Add `-n`/`-N` flags
4. **Write tests** using `testing` package (table-driven tests preferred)

### Skeleton

Create `projects/phase-1/00-day-4-gophercises/01-echo/main.go`

```go
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Define flags
	sep := flag.String("s", " ", "separator")
	newline := flag.Bool("n", true, "add newline at end")
	flag.Parse()

	args := flag.Args()

	// Construct output
	output := strings.Join(args, *sep)

	// Handle newline
	if *newline {
		fmt.Println(output)
	} else {
		fmt.Print(output)
	}
}
```

### Create Basic Tests

Create `projects/phase-1/00-day-4-gophercises/01-echo/main_test.go`

```go
package main

import (
	"testing"
)

// Test the logic (you may need to refactor main.go to make this testable)
func TestEcho(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		sep       string
		newline   bool
		expected  string
	}{
		{"single arg", []string{"hello"}, " ", true, "hello\n"},
		{"two args default sep", []string{"hello", "world"}, " ", true, "hello world\n"},
		{"custom separator", []string{"a", "b", "c"}, ",", true, "a,b,c\n"},
		{"no newline", []string{"hello"}, " ", false, "hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: Test your echo function
		})
	}
}
```

### Verify Your Solution

```bash
cd projects/phase-1/00-day-4-gophercises/01-echo
go run . hello world
go run . -s=, a b c
go run . -n=false test
go test -v
go vet ./...
gofmt -l .
```

### Commit to GitHub

```bash
git add projects/phase-1/00-day-4-gophercises/01-echo/
git commit -m "Gophercises #1: Echo CLI with tests"
git push origin master
```

---

## Checklist for End of Day

- [ ] TGPL Chapters 2–3 read completely
- [ ] All examples from Chapters 2–3 typed and tested
- [ ] Notes created/updated in `notes/books/TGPL/tgpl-chapter-2-3.md`
- [ ] Gophercises #1 solution complete and passing tests
- [ ] Code passes `gofmt`, `go vet`, and `go test`
- [ ] Changes committed to GitHub with meaningful message
- [ ] Hours logged for the day (Goal: 2.5 hours)

---

## If You Get Stuck

**Confused by structs?** → Read through Chapter 3 struct section again, then write a small custom struct (e.g., `type Person struct { Name string; Age int }`) and use it.

**Flag parsing confusing?** → Review `go run . -h` output and then write a small test program just for flags before tackling Gophercises #1 fully.

**Tests not working?** → Refactor your `main.go` to extract the echo logic into a separate function you can test, then write tests for that function.

---

## Looking Ahead (Week 2 Roadmap)

| Day | Date | Focus |
|---|---|---|
| Day 1 | Mon, Mar 30 | TGPL Ch. 1 + Gophercises #1 start |
| **Day 2** | **Tue, Mar 31** | **TGPL Ch. 2–3 + Gophercises #1 complete** |
| Day 3 | Wed, Apr 1 | TGPL Ch. 4 + Gophercises #2 |
| Day 4 | Thu, Apr 2 | Mini-project (guided): CLI todo app |
| Day 5 | Fri, Apr 3 | Mini-project (guided): Continue & testing |
| Day 6 | Sat, Apr 4 | Code review: Self-critique + cleanup |
| Day 7 | Sun, Apr 5 | Week review + log final hours |

---

## Time Tracking

**Today's target**: 2.5 hours

- Session 1 (TGPL 2–3): 2 hours
- Session 2 (Gophercises #1): 0.5 hours

**Week 2 cumulative**: ~5 hours (Goal by end of day: 5 / 12 hours)

Log your actual time at end of day in `roadmap/TIME-TRACKING.md`.
