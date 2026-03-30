# Week 3: TGPL Chapters 5–6 + Gophercises #3 & #4 + Function Mastery

**Week**: Week 3 of Phase 1 — Module 1.1 (Go Fundamentals)
**Dates**: April 6 – April 12, 2026
**Total Hours Goal**: 13 hours (Module 1.1 requires 200 hours over ~16 weeks → average 12–13h/week)
**Location**: `Year-2026/Month-03/WEEK-3-START-HERE.md`
**Roadmap Reference**: `roadmap/PHASE-1-MODULES.md` → Module 1, Submodule 1.1 (Go Basics)

---

## ⚠️ Read This Before You Start

Week 2 built fluency with types and data structures. Week 3 gives you the tools to write reusable code: functions and methods. This week's content is foundational for everything that follows.

- **Week 2**: Data structures, basic programs, testing discipline
- **Week 3**: Functions, methods, interfaces, design patterns

By the end of this week, you will understand how to design APIs in Go — how functions should be organized, how methods extend types, and how interfaces enable composition.

### Module 1.1 at a Glance (Progress Check)

From `roadmap/PHASE-1-MODULES.md` — ALL of these are required before Module 1.2:

| Requirement | Target | Your Status After Week 2 |
|---|---|---|
| TGPL Book — all chapters + exercises | Complete | ~20% (Ch 1-4) |
| Gophercises completed | 20+ | 2/20 |
| Phase 1 Projects on GitHub (>80% coverage) | 3 projects | 1/3 (Week 2 mini-project) |
| Can explain goroutines, channels, select | Yes | Not yet (Week 4) |
| Idiomatic Go — code review approved | Yes | Not started |
| External code review received | ≥1 project | Not started |
| Hours logged | 200 hours | ~22h |

**Week 3 advances**: Function mastery, method receivers, interfaces intro, Gophercises count to 4.

---

## Week 3 Goals

**Primary Goal**: Complete TGPL Chapters 5–6 with typed examples and deep understanding. Complete Gophercises #3 and #4 with tests. Fully understand function design patterns.

**Secondary Goals**:
- Write functions with clear, documented APIs (good parameter names, error handling)
- Understand methods and receivers deeply — when to use value receivers vs pointer receivers
- Begin thinking about interfaces — recognize when interfaces emerge from your code
- Start Project 1.1.1 planning: Log Parser CLI (launches Week 4)
- Maintain code quality discipline (`go vet`, `gofmt`, `go test -cover`)
- Commit to GitHub daily with meaningful commit messages
- Log hours daily — never batch on Sunday

**Hours Target**: 13 hours (not 12 — Module 1.1 requires 200h total over 16 weeks at 12.5h/week average)

---

## Key Resources for This Week

| Resource | Use For | Link |
|---|---|---|
| TGPL Book (Donovan & Kernighan) | Main reading — Ch 5–6 | Your copy |
| Gophercises | Exercises #3, #4 | https://gophercises.com/ |
| pkg.go.dev | Documentation reference | https://pkg.go.dev/ |
| Go by Example — Methods | Method receiver patterns | https://gobyexample.com/methods |
| Go by Example — Interfaces | Interface design | https://gobyexample.com/interfaces |
| Effective Go — Functions | Function best practices | https://go.dev/doc/effective_go |
| Go Blog — Receiver Types | Pointer vs value receivers | https://go.dev/blog/methods |

---

## Before You Write Any Code This Week

Check that all code from Week 2 still passes quality checks:

```bash
# Make sure nothing regressed
gofmt -l .

# All vet warnings must still be resolved
go vet ./...

# Test coverage should not have dropped
go test -cover ./...
```

Do NOT start Week 3 with quality debt from Week 2.

---

## Detailed Daily Plan

---

### Sunday, April 6 (Pre-week setup: 30 min)

**Complete the Week 2 Review from WEEK-2-START-HERE.md if not done**

If you have not yet done the Week 2 Sunday reflection (journal entry, self-assessment, Week 3 planning):
- [ ] Complete Week 2 journal entry
- [ ] Fill in self-assessment confidence table
- [ ] Identify any Chapter 1-4 topics that rated 1-2 (plan to review)
- [ ] Update `roadmap/TIME-TRACKING.md` with Week 2 hours

This takes 30 minutes. Do not skip. Start Monday with a clean slate.

---

### Monday, April 6 (Target: 2.5 hours)

**Focus: TGPL Chapter 5 — Functions**

Chapter 5 teaches you how to write functions that are idiomatic, maintainable, and composable. This is where Go design starts.

#### Session 1: TGPL Chapter 5 (90 min)

- [ ] Read Chapter 5 cover to cover — do NOT skim
- [ ] Type out every code example — even ones that look obvious
- [ ] Run every example: `go run filename.go`
- [ ] For each example, answer in your notes: *"When would I use this pattern in real code?"*
- [ ] Take notes in: `notes/books/tgpl-chapter-5.md`

**Key topics in Chapter 5 to watch for**:
- Function declarations and type signatures
- Multiple return values and named returns
- Error handling patterns in Go (defer, custom error types)
- Anonymous functions and closures
- Variadic functions (`...Type`)
- Function values and higher-order functions
- Deferred function calls and cleanup

**Things beginners often struggle with in Chapter 5**:
- Understanding when to use named vs. unnamed return values
- Why Go encourages multiple return values instead of exceptions
- How `defer` interacts with panic/recover
- The difference between a function value and calling a function

### Critical Concept: Error Handling in Go

Go does NOT use exceptions. Every function that can fail returns an error as the last return value. This is foundational:

```go
// Pattern: (result, error) — learn this deeply
func Find(name string) (Person, error) {
    if name == "" {
        return Person{}, fmt.Errorf("name cannot be empty")
    }
    // ... find logic ...
    return person, nil
}

// Always check the error immediately
p, err := Find("Alice")
if err != nil {
    // Handle error
    return err
}
// Use p only here
```

Spend extra time understanding this. You will write this pattern thousands of times.

#### Session 2: Start Gophercises #3 (60 min)

URL: https://gophercises.com/ → Exercise 3

The URL shortener problem involves a bit of state management but focuses on the patterns you just learned.

- [ ] Read the full problem statement
- [ ] Plan before coding — write your plan in comments:

```go
// Plan:
// 1. Parse command-line flag for JSON file (default: urls.json)
// 2. Build a map[short] -> url
// 3. Start http server on localhost:8080
// 4. Handle GET /short → redirect to full URL
// 5. Handle 404 gracefully
```

- [ ] Start with the basic version (no persistence)
- [ ] Use the `net/url` package to parse URLs — read the docs first
- [ ] Use the `net/http` package (you saw this in Chapter 1, now you will use it more)
- [ ] Write at least 2 tests for URL parsing

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Tuesday, April 7 (Target: 2 hours)

**Focus: Finish Gophercises #3 + TGPL Chapter 6 Start**

#### Session 1: Finish Gophercises #3 + Write Tests (60 min)

- [ ] Complete URL shortener if not done
- [ ] Write table-driven tests for both URL parsing and HTTP routing
- [ ] Test that redirects work correctly (test the HTTP handler)
- [ ] Run `go test -v -cover ./...` — aim for >75% coverage
- [ ] Run `go vet ./...` — fix any issues
- [ ] Push finished Gophercises #3

**Testing tip**: To test HTTP handlers in Go, use the `net/http/httptest` package:

```go
import (
    "net/http/httptest"
    "testing"
)

func TestRedirect(t *testing.T) {
    // Create a mock request
    req := httptest.NewRequest("GET", "/short", nil)
    
    // Create a mock response writer
    w := httptest.NewRecorder()
    
    // Call your handler
    YourHandler(w, req)
    
    // Assert the response
    if w.Code != http.StatusMovedPermanently {
        t.Errorf("got %d, want %d", w.Code, http.StatusMovedPermanently)
    }
}
```

Read the `net/http/httptest` docs: https://pkg.go.dev/net/http/httptest

#### Session 2: TGPL Chapter 6 — Methods (60 min)

- [ ] Read Chapter 6 carefully — this chapter defines how Go handles OOP without inheritance
- [ ] Take notes in: `notes/books/tgpl-chapter-6.md`
- [ ] For every concept, write a 3-line code snippet in your notes to anchor it

**Key concepts in Chapter 6** — understand each before Thursday:

| Concept | One-line explanation (write yours) |
|---|---|
| Method declarations and receivers | |
| Pointer vs. value receivers | |
| Methods on non-struct types | |
| The method set | |
| Embedded types and promoted methods | |
| Interfaces — defining and implementing | |
| Interface satisfaction (implicit) | |
| Type assertions (`x.(T)`) | |
| Type switches (`switch x.(type)`) | |

**Critical Pattern: Interfaces Emerge From Similarity**

Unlike languages like Java or C#, Go does NOT enforce interface definitions at compile time. Instead, interfaces describe behavior:

```go
// If two types both have a Read([]byte) (int, error) method,
// they both satisfy io.Reader implicitly.
// You don't declare "implements io.Reader"
type io.Reader interface {
    Read(p []byte) (n int, err error)
}

type MyFile struct { /* ... */ }

func (f *MyFile) Read(p []byte) (int, error) {
    // implement reading
}

// MyFile now satisfies io.Reader without any declaration
```

This is a subtle but powerful difference from other languages. Spend time understanding it.

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Wednesday, April 8 (Target: 2.5 hours)

**Focus: Deep Dive on Methods + Start Gophercises #4**

#### Session 1: TGPL Chapter 6 Deep Dive — Receivers & Interfaces (90 min)

This session is pure focus on the hardest concepts in Chapter 6.

**Part A: Pointer vs. Value Receivers (30 min)**

Write out the decision tree:

```
Use a POINTER receiver (*T) when:
1. The method modifies the receiver
2. The type is large (e.g., a large struct with many fields)
3. For consistency, if ANY method on the type uses a pointer receiver

Use a VALUE receiver (T) when:
1. The method reads only (does not modify)
2. The type is small (int, string, small struct)
3. You want the method to work on copies
```

Take examples from Chapter 6 and annotate each with why the author chose pointer or value receiver.

**Part B: Interfaces in Practice (60 min)**

The biggest mistake beginners make: trying to define interfaces before knowing what they want.

The Go way: Write functions that operate on concrete types. If two types have similar behavior, extract an interface.

Example:

```go
// BAD: designing interfaces first
type Animal interface {
    Speak() string
    Move() string
}

type Dog struct{}
func (d *Dog) Speak() string { return "woof" }
func (d *Dog) Move() string { return "run" }

// GOOD: concrete types first, interface emerges
type Dog struct{}
func (d *Dog) Speak() string { return "woof" }
func (d *Dog) Move() string { return "run" }

type Cat struct{}
func (c *Cat) Speak() string { return "meow" }
func (c *Cat) Move() string { return "jump" }

// NOW we see both types have Speak() and Move()
// So we extract the interface:
type Speaker interface {
    Speak() string
}
```

Write 3 examples in your notes where interfaces emerge from concrete code.

#### Session 2: Start Gophercises #4 (60 min)

URL: https://gophercises.com/ → Exercise 4

This exercise involves link parsing from HTML. It will teach you how to use the `net/http` package to fetch pages and the `golang.org/x/net/html` package to parse them.

- [ ] Read the full problem statement
- [ ] Plan before coding — write your plan in comments:

```go
// Plan:
// 1. Read HTML from file or stdin
// 2. Parse HTML using golang.org/x/net/html
// 3. Find all <a> tags and extract href attributes
// 4. Print each link (or output JSON)
// 5. Handle malformed HTML gracefully
```

- [ ] You will need to import `golang.org/x/net/html` — run `go get golang.org/x/net/html` first
- [ ] Use the Tokenizer pattern from the package docs
- [ ] Write 2 tests parsing sample HTML

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Thursday, April 9 (Target: 2 hours)

**Focus: Finish Gophercises #4 + Method Design Deep Dive**

#### Session 1: Finish Gophercises #4 + Add Tests (60 min)

- [ ] Complete link parser if not done
- [ ] Write table-driven tests for at least 3 different HTML snippets
- [ ] Test edge cases: malformed HTML, missing href, nested tags
- [ ] Run `go test -v -cover ./...` — aim for >75% coverage
- [ ] Run `go vet ./...` — fix any issues
- [ ] Push finished Gophercises #4

#### Session 2: Advanced Function & Method Patterns (60 min)

Now that you understand functions and methods, spend 60 minutes studying idiomatic Go design patterns.

**Pattern 1: The Builder Pattern**

Go does not have constructors. Instead, use functions that return configured values:

```go
type Config struct {
    Host string
    Port int
    Timeout time.Duration
}

// Option is a function that configures a Config
type Option func(*Config)

func WithHost(h string) Option {
    return func(c *Config) {
        c.Host = h
    }
}

// Constructor using options
func NewConfig(opts ...Option) Config {
    c := Config{
        Host: "localhost",
        Port: 8080,
        Timeout: 30 * time.Second,
    }
    for _, opt := range opts {
        opt(&c)
    }
    return c
}

// Usage:
config := NewConfig(WithHost("example.com"), WithPort(3000))
```

Study this pattern — it appears in many Go libraries.

**Pattern 2: Error Wrapping**

Modern Go (1.13+) uses error wrapping:

```go
if err != nil {
    return fmt.Errorf("failed to read config: %w", err)
}
```

The `%w` verb wraps the error. Later code can use `errors.Is()` and `errors.As()` to inspect wrapped errors.

Write 2 examples in your notes.

**Pattern 3: Interfaces for Testability**

Define interfaces for dependencies you want to mock:

```go
type UserStore interface {
    GetUser(id string) (User, error)
}

type Service struct {
    store UserStore
}

// In production:
service := &Service{store: database.NewPostgresUserStore()}

// In tests:
mockStore := &MockUserStore{}
service := &Service{store: mockStore}
```

This is how you make code testable.

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Friday, April 10 (Target: 1.5 hours)

**Focus: Code Quality + Project Planning**

This day is about quality and planning for the big project next week.

#### Session 1: Code Quality Audit (45 min)

Run this on all code written this week:

```bash
# 1. Format every file
gofmt -w .

# 2. Check for common mistakes
go vet ./...

# 3. Check test coverage across all packages
go test -cover ./...

# 4. Run tests verbosely to see what's actually passing
go test -v ./...
```

Fix every `go vet` warning. No exceptions.

- [ ] All files pass `gofmt`
- [ ] All files pass `go vet`
- [ ] All packages have at least 70% test coverage (80%+ preferred)
- [ ] No `TODO` or `FIXME` comments left without context

**Also check**: Do your Gophercises #3 and #4 have docstrings on exported functions? They should.

```go
// Find searches the map for the given ID and returns the Person.
// It returns an error if the ID is not found.
func (m Map) Find(id string) (Person, error) {
    // ...
}
```

#### Session 2: Project Planning — Log Parser CLI (45 min)

Next week you start Project 1.1.1: the Log Parser CLI. This is NOT a Gophercises problem — it is the first real project in your Portfolio. Spend 45 minutes planning.

**Create a file**: `projects/phase-1/week-4-log-parser-cli/PLAN.md`

Structure:

```markdown
# Log Parser CLI — Development Plan

## What It Does
- Reads a log file (any text format)
- Filters by log level (DEBUG, INFO, WARN, ERROR)
- Outputs statistics (JSON, CSV, or human-readable)
- Handles large files efficiently (stream processing)

## Requirements
- [ ] Accept `-file`, `-level`, `-format` flags
- [ ] Default: file=stdin, level=DEBUG, format=plaintext
- [ ] Process files in streaming fashion (not loading into memory)
- [ ] Parse lines like: [TIMESTAMP] [LEVEL] message
- [ ] Output statistics: total lines, per-level counts, percentages
- [ ] Error handling: missing file, malformed lines, invalid flags

## Design Decisions (write your reasoning)
1. How will I parse log lines? (regex? string.Split?)
2. Should I support custom log formats? (Not for MVP)
3. How will I define log levels? (const string? iota?)
4. What errors might occur? How will I handle them?

## Code Structure
```
log-parser/
├── main.go           (CLI entry, flags, output)
├── parser.go         (reading and parsing log lines)
├── formatter.go      (JSON, CSV, plaintext output)
├── main_test.go
├── parser_test.go
├── formatter_test.go
└── README.md

```

## Test Strategy
- [ ] Table-driven tests for parsing each log level
- [ ] Tests for different output formats
- [ ] Test with real-world log samples
- [ ] Test edge cases: empty file, single line, missing fields
- [ ] Target >85% coverage

## Success Criteria
- Works correctly with real-world logs
- >85% test coverage
- All code passes `gofmt` and `go vet`
- README with usage examples
- GitHub repo with meaningful commit history
```

Fill in the plan. Do NOT code yet. Next week you will.

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Saturday, April 11 (Target: 2.5 hours)

**Focus: Independent Mini-Project #2 — Limited Guidance**

This is your second independent project. You have more freedom now, and you should feel the difference in confidence.

Choose ONE:

---

**Option A: Text Statistics Tool (Recommended)**

Build a CLI tool that analyzes a text file and prints statistics.

Requirements (you must implement all):
- Accept filename and `-stat` flag (options: lines, words, chars, sentences, paragraphs)
- Default: all statistics
- Count correctly (words = space-separated tokens, sentences = ends with `.!?`, etc)
- Output as table or JSON (add `-format` flag)
- Handle empty files, files with special characters
- Write at least 6 table-driven tests

Skills this use: file I/O, strings, formatting, methods, testing

---

**Option B: Duplicate File Finder**

Build a CLI tool that finds duplicate files by hash.

Requirements:
- Accept directory and `-hash` flag (options: md5, sha1, sha256)
- Recursively scan directory for all files
- Compute hash of each file
- Group files by hash and report duplicates
- Show file paths, sizes, and hashes
- Handle permission errors gracefully
- Write at least 5 table-driven tests

Skills this uses: file I/O, crypto/md5, path/filepath, structs, sorting, testing

---

**Option C: Configuration File Merger**

Build a CLI tool that merges multiple YAML/JSON configuration files.

Requirements:
- Accept multiple files as arguments: `tool file1.json file2.json file3.json`
- Merge them (later files override earlier ones)
- Validate that output is valid JSON/YAML
- Pretty-print the output
- Optionally save to file with `-output` flag
- Handle missing files, invalid JSON/YAML gracefully
- Write at least 6 table-driven tests

Skills this uses: encoding/json, os/fs, struct tags, error handling, testing

---

**For whichever option you choose**:

- [ ] Write a design document (not code) — 5 minutes
- [ ] Define the type structure you will use — 10 min (draw it out if needed)
- [ ] Build it (75 min) with tests as you go
- [ ] Run `go vet ./...` and fix everything
- [ ] Run `go test -v -cover ./...` — target >80% coverage
- [ ] Add a `README.md` (how to build, example usage)
- [ ] Push to GitHub: `projects/phase-1/week-3-mini-project/`
- [ ] Commit message: `feat: week-3 independent mini-project — [option name]`

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Sunday, April 12 (Target: 1 hour)

**Focus: Weekly Review + Plan Week 4**

#### Part 1: Journal Entry (30 min)

Create your Week 3 journal:
```bash
cp journal/JOURNAL-TEMPLATE.md journal/2026-04-week-14.md
```

Fill in every section. Specific requirements for this week's journal:

**Technical Concepts section** — write 5+ specific things. Examples:
- "Method receivers: I finally understood when to use pointers. The rule: use pointer receiver if ANY method on the type uses one, for consistency."
- "Interfaces emerge from similar behavior, not the other way around. I kept trying to design interfaces first, but Go wants concrete code first."
- "Error handling in Go feels verbose at first, but it forces you to think about every failure case. I wrote better error messages this week than I ever have."

**Struggles section** — be specific:
- "Type assertions confused me. I spent 20 minutes understanding `v, ok := someInterface.(ConcreteType)`"
- "I wrote tests for HTTP handlers for the first time using httptest. Had to read the docs 3 times before I got it."

**Questions to Research** — list them:
- "How do I design interfaces that are actually useful? The chapter showed examples but I'm not confident."
- "When should I use type switches instead of type assertions?"

#### Part 2: Self-Assessment (15 min)

Rate your current confidence on all TGPL topics so far (1 = shaky, 5 = solid):

| Topic | Chapter | Confidence |
|---|---|---|
| Package structure and imports | 1 | |
| Functions and signatures | 5 | |
| Multiple return values and error handling | 5 | |
| Anonymous functions and closures | 5 | |
| Defer statement | 5 | |
| Methods on types | 6 | |
| Pointer vs value receivers | 6 | |
| Interfaces and satisfaction | 6 | |
| Type assertions and switches | 6 | |

Confidence ratings should be increasing as you progress. Any rating of 1-2 → add a focused review to Week 4.

#### Part 3: Week 4 Planning (15 min)

Open `roadmap/PHASE-1-MODULES.md` → Module 1.1 and confirm pace.

Week 4 target:
- TGPL Chapters 7–9 (concurrency — goroutines, channels, sync)
- Start Project 1.1.1: Log Parser CLI (2 weeks project)
- Complete Gophercises #5-6 OR dive straight into concurrency exercises
- Continue building testing discipline

Things to complete before Week 4:
- All code from Week 3 should pass `go vet` with no warnings
- All tests from Week 3 should be passing
- All code should be pushed to GitHub
- Project plan for Log Parser should be ready

Write your 5 goals for Week 4 in `roadmap/weekly-goals.md`.

**Actual time spent**: ___ hours
**What I accomplished**:

---

## Week 3 Summary

### Total Time Logged
**Planned**: 13 hours
**Actual**: ___ hours
**Variance**: ___ hours (over/under)
**Running Total (Phase 1)**: ___ / 200 hours (Module 1.1 target)

### Deliverables Check

- [ ] TGPL Chapters 5–6 read with notes taken in `notes/books/tgpl-chapter-[5-6].md`
- [ ] Notes include specific design patterns (builder, error wrapping, interfaces for testing)
- [ ] Gophercises #3 (URL shortener) — done + tests written + pushed
- [ ] Gophercises #4 (link parser) — done + tests written + pushed
- [ ] Mini-project #2 — done + tests + README + pushed to GitHub
- [ ] All code passes `gofmt` and `go vet`
- [ ] All packages have >75% test coverage (`go test -cover ./...`)
- [ ] Daily commits to GitHub — at least 6 this week
- [ ] Project 1.1.1 (Log Parser CLI) — design plan created and committed
- [ ] Week 3 journal filled in
- [ ] Week 4 goals written in `roadmap/weekly-goals.md`
- [ ] Running hours total updated in `roadmap/TIME-TRACKING.md`

### Gophercises Tracker (updated)

| # | Exercise | Status | Tests Written | Coverage |
|---|---|---|---|---|
| 1 | Echo CLI | Complete | Yes | ___ |
| 2 | Quiz Game | Complete | Yes | ___ |
| 3 | URL Shortener | Complete | Yes | ___ |
| 4 | Link Parser | Complete | Yes | ___ |
| 5 | | Not started | | |
| ... | | | | |
| 20 | | Not started | | |

---

## Concepts You Must Be Able to Explain by End of Week 3

Answer these without looking anything up. If you cannot, go back and restudy.

**TGPL Chapters 5**
1. What is the difference between returning `(result, error)` and throwing an exception? When is Go's approach better?
2. What does `fmt.Errorf("format: %w", err)` do differently than `fmt.Errorf("format: %v", err)`?
3. What is a closure and give one real-world example of when you would use one?
4. What is a variadic function? Give an example and explain why you would use one.
5. What does the `defer` statement do? Give three examples of how you would use it.

**TGPL Chapter 6**
6. What is the difference between these two?
   ```go
   type Dog struct{}
   
   func (d Dog) Bark() string { return "woof" }      // Value receiver
   func (d *Dog) Bark() string { return "woof" }     // Pointer receiver
   ```
7. Why would you want to use a pointer receiver?
8. What is an interface in Go and how is it different from interfaces in Java or C#?
9. How does Go determine if a type satisfies an interface? (Hint: NOT explicit)
10. What is the difference between `interface{}` and `any`? (Hint: they are the same in Go 1.18+)
11. What is a type assertion and what does `v, ok := x.(bool)` tell you?
12. What is a type switch and when would you use it instead of multiple type assertions?

**Function & Method Design**
13. Write a function that takes a variable number of arguments and returns an error. Include error handling.
14. Design a simple interface for "something that can be saved" and explain why you would use it.
15. When you see `func New...() *MyType`, what pattern is this and why is it common in Go?

---

## Code Quality Standards — Maintained from Week 2

From `roadmap/PHASE-1-MODULES.md` — all deliverables require:

| Standard | Requirement | How to Check |
|---|---|---|
| Formatting | All files formatted | `gofmt -l .` returns nothing |
| Vet | No vet warnings | `go vet ./...` returns nothing |
| Test coverage | >80% for projects, >75% for exercises | `go test -cover ./...` |
| Error handling | Every error is handled — no ignored errors | Manual review |
| Docstrings | Exported functions have comments | `go doc ./...` |
| README | Every project has a README | Present in repo |
| No dead code | Every function/type is used | Manual review |

**NEW Standard**: Every package should have a comment at the top explaining its purpose:

```go
// Package parser provides functions for parsing log files.
package parser

// A LogEntry represents a single line from a log file.
type LogEntry struct { ... }
```

**NEW Standard**: Design decisions should be documented in comments or a DESIGN.md if they are non-obvious:

```go
// We use a map instead of a slice for lookups because
// we frequently need to find entries by ID, and a map
// gives us O(1) lookup instead of O(n) scan.
var entries map[string]LogEntry
```

---

## If You Are Behind

**Didn't finish TGPL Chapters 5-6?**
→ These chapters are foundational. Do not skip ahead to concurrency without mastery.
→ Move unfinished chapters to the start of Week 4. Note in your journal why and what you will change.

**Couldn't hit 13 hours?**
→ Log exactly how many you got. Do not round up.
→ Identify which sessions you skipped or cut short.
→ The 200-hour Module 1.1 target requires 12.5h/week average. You need to find 1-2 hours somewhere.

**Gophercises #3 or #4 is confusing?**
→ Spend 30 minutes reading the problem statement multiple times before coding.
→ Sketch out the data structures first (what types do I need?).
→ Ask in your journal: "What is the minimum viable version that works?"
→ Build that before adding features.

**HTTP testing with httptest feels unfamiliar?**
→ It is. HTTP is abstractions on top of abstractions.
→ Read the `net/http` and `net/http/httptest` packages docs fully.
→ Write 3 simple handlers and test each one manually first.
→ Then write tests using httptest.

**Interfaces are confusing?**
→ Absolutely normal. Interfaces are subtle in Go.
→ Reread the "Interfaces Emerge From Similarity" section in this plan.
→ Write 5 small examples in your notes. No copy-pasting — type them.
→ Don't try to understand everything. Just understand: Go checks interface satisfaction implicitly at the call site.

---

## Next Week: The Big Week — Concurrency Begins

Week 4 introduces goroutines, channels, and the `select` statement. These are the hardest concepts in Go, but also the most powerful. By the end of Week 4, you will understand concurrent programming patterns that took years for developers in other languages to learn.

Prepare mentally: concurrency is hard. Embrace the difficulty.

**Week 4 teaser**:
- Goroutines: lightweight threads managed by the Go runtime
- Channels: how goroutines communicate
- Select: multiplexing on multiple channels
- Race conditions and the data race detector
- Timeout patterns and cancellation

See you in Week 4.
