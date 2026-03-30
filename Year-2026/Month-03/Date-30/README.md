# Day 1 (Week 2) — 30 March 2026

**Phase**: Phase 1 / Module 1.1 — Go Fundamentals  
**Focus**: TGPL Chapter 1 + Gophercises #1 Start  
**Time Target**: 2.5 hours  
**Week**: Week 2 of Phase 1 (March 30 – April 5, 2026)  

---

## Context: The Shift from Week 1 to Week 2

Week 1 gave you vocabulary through guided examples. Week 2 gives you fluency through professional technical content and independent building.

**The shift is significant:**
- Week 1: You followed the Tour of Go
- Week 2: You read a professional technical book (TGPL) and build independent programs

This week will feel harder. That is correct. The difficulty is the point.

---

## Today's Big Picture

**Monday target**: 2.5 hours across two sessions:
1. **TGPL Chapter 1** — Type every example, understand every concept (90 min)
2. **Gophercises #1 Start** — Echo CLI (60 min)

**Before you start anything:**

Run this on your Week 1 code to ensure code quality:

```bash
cd /path/to/projects/phase-1
gofmt -l .           # Format check
go vet ./...         # Vet check  
go test -cover ./... # Test coverage baseline
```

If your code has `gofmt` or `go vet` issues, fix them now. Clean code is a non-negotiable habit.

---

## Session 1: TGPL Chapter 1 — A Tutorial (90 min)

Chapter 1 is a fast walkthrough of Go's most common features. The examples are short but dense. Type every single one.

### The Rule
**Do NOT copy-paste.** Type every example yourself. This is how you build muscle memory.

### For each example in Chapter 1:

- [ ] Read the explanation carefully
- [ ] Type the code (do not copy-paste)
- [ ] Run it: `go run filename.go`
- [ ] Answer this in your notes: *"What would happen if I changed X?"* — then try it
- [ ] Move to the next example

### Key Topics in Chapter 1 (Watch for These)

| Concept | What to Understand |
|---|---|
| `os.Args` | How Go accesses command-line arguments |
| `bufio.Scanner` | How to read lines from stdin efficiently |
| `http.Get` | How to fetch web pages from Go code |
| Named types (`type Celsius float64`) | Go's way of creating new types from existing ones — appears everywhere |
| Error handling | How Go handles errors (no exceptions) |

### Where to Take Notes

Create file: `notes/books/tgpl-chapter-1.md`

Structure:
```
# TGPL Chapter 1 — A Tutorial

## Key Concepts
- **os.Args**: [Your explanation in your own words]
    - os.Args is a slice (list) of strings that stores command-line arguments passed to a Go program. The first element (os.Args[0]) is the program name, and the rest are inputs given by the user when running the program
- **bufio.Scanner**: [Your explanation]
    - bufio.Scanner is used to read input efficiently, usually from standard input or files. It reads data line by line (or word by word), making it easy to process user input or file content. Each call to Scan() reads the next line, and Text() returns that line.
- **http.Get**: [Your explanation]
    - http.Get is a function used to make HTTP requests (usually GET requests) to a URL. It fetches data from the web and returns a response, which can then be read (like a webpage’s HTML content).
- **Named types**: [Your explanation]
    - Named types are custom types created using the type keyword. They allow you to give a meaningful name to a type (e.g., type Celsius float64), improving code readability and enabling methods to be attached to them.

## Examples I Typed
1. prints command-line arguments
2. Uses http.Get to retrieve data from URLs and print the response.
...

## Questions & Confusions
- [What confused me?]
- [What do I need to revisit?]
```

**Time allocation**:
- Reading: 30 min
- Typing examples: 50 min  
- Taking notes: 10 min

---

## Session 2: Gophercises #1 — Echo CLI (60 min)

After finishing Chapter 1, start the first Gophercises problem.

### Problem

URL: https://gophercises.com/ → Exercise 1

Read the full problem before writing code (5 min).

### Plan Before Coding (10 min)

Before writing any Go code, write a comment plan:

```go
// Plan:
// 1. Read all command-line args using os.Args
// 2. Join them with spaces
// 3. Print to stdout
// 4. Handle edge case: no args provided
```

### Build the Tool (30 min)

- Use only `os` and `fmt` packages — no external libraries
- Run it manually with test inputs before writing formal tests
- Example test runs:
  ```bash
  go run echo.go arg1 arg2 arg3
  # Output: arg1 arg2 arg3
  
  go run echo.go hello
  # Output: hello
  
  go run echo.go
  # Output: [empty or handle gracefully]
  ```

### Write Your First Tests (15 min)

Create `echo_test.go` with **at least 2 test cases**.

**Testing requirement**: Every function in Gophercises must have at least one test. No exceptions starting this week.

### Example Test Structure

```go
package main

import (
    "testing"
)

func TestEcho(t *testing.T) {
    tests := []struct {
        name     string
        args     []string
        expected string
    }{
        {"single arg", []string{"hello"}, "hello"},
        {"multiple args", []string{"hello", "world"}, "hello world"},
        {"no args", []string{}, "[your behavior]"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Your assertion logic here
        })
    }
}
```

---

## Time Breakdown

| Task | Target | Actual |
|---|---|---|
| Session 1: TGPL Ch. 1 | 90 min | ___ |
| Session 2: Gophercises #1 Start | 60 min | ___ |
| **Total** | **2.5 hours** | ___ |

---

## Commit to GitHub (10 min after sessions)

When done with both sessions:

```bash
git add .
git commit -m "week-2: tgpl chapter 1 notes + gophercises-01 echo cli start"
git push
```

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

- [ ] Chapter 1 read completely + all examples typed
- [ ] Chapter 1 notes written in `notes/books/tgpl-chapter-1.md`
- [ ] Gophercises #1 started (basic echo.go written)
- [ ] Tests written for Gophercises #1
- [ ] `go vet ./...` passes with no warnings
- [ ] Code pushed to GitHub
- [ ] Hours logged accurately (be honest)

---

## Actual Time Spent & Summary

**Total hours: ___ (Target: 2.5)**

**What I accomplished**:

**What I struggled with**:

**What to revisit tomorrow**:
