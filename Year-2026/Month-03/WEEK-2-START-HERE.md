# Week 2: TGPL Chapters 1–4 + Gophercises #1 & #2 + Test Discipline

**Week**: Week 2 of Phase 1 — Module 1.1 (Go Fundamentals)
**Dates**: March 30 – April 5, 2026
**Total Hours Goal**: 12 hours (Module 1.1 requires 200 hours over ~16 weeks → average 12–13h/week)
**Location**: `Year-2026/Month-03/WEEK-2-START-HERE.md`
**Roadmap Reference**: `roadmap/PHASE-1-MODULES.md` → Module 1, Submodule 1.1 (Go Basics)

---

## ⚠️ Read This Before You Start

Week 1 gave you vocabulary. Week 2 gives you fluency. The shift is significant:

- **Week 1**: You followed guided Tour examples
- **Week 2**: You read a professional technical book and build independent programs

This week will feel harder. That is correct. The difficulty is the point.

### Module 1.1 at a Glance (Your Full Target)

From `roadmap/PHASE-1-MODULES.md` — ALL of these are required before Module 1.2:

| Requirement | Target | Your Status |
|---|---|---|
| TGPL Book — all chapters + exercises | Complete | 0% |
| Gophercises completed | 20+ | 0/20 |
| Phase 1 Projects on GitHub (>80% coverage) | 3 projects | 0/3 |
| Can explain goroutines, channels, select | Yes | Partial |
| Idiomatic Go — code review approved | Yes | Not started |
| External code review received | ≥1 project | Not started |
| Hours logged | 200 hours | ~10h |

**Week 2 advances**: TGPL progress, Gophercises count, testing habits.

---

## Week 2 Goals

**Primary Goal**: Complete TGPL Chapters 1–4 with typed examples and notes. Complete Gophercises #1 and #2 with tests.

**Secondary Goals**:
- Build your first independent mini-project from scratch, no guidance
- Write table-driven tests for every non-trivial function you write
- Establish code quality discipline (`go vet`, `gofmt`, `go test -cover`)
- Commit to GitHub daily with meaningful commit messages
- Log hours daily — never batch on Sunday

**Hours Target**: 12 hours (not 10 — Module 1.1 requires 200h total)

---

## Key Resources for This Week

| Resource | Use For | Link |
|---|---|---|
| TGPL Book (Donovan & Kernighan) | Main reading — Ch 1–4 | Your copy |
| Gophercises | Exercises #1, #2 | https://gophercises.com/ |
| pkg.go.dev | Documentation reference | https://pkg.go.dev/ |
| Go by Example — Testing | Testing patterns | https://gobyexample.com/testing |
| Go by Example — Command-Line Flags | Flag parsing for Gophercises | https://gobyexample.com/command-line-flags |
| Effective Go | Idiomatic style reference | https://go.dev/doc/effective_go |
| Go Blog — Table-Driven Tests | Testing discipline | https://go.dev/blog/subtests |

---

## Before You Write Any Code This Week

Run this on all code you wrote in Week 1:

```bash
# Format check — all files must pass before you write new code
gofmt -l .

# Vet check — all warnings must be resolved
go vet ./...

# Test coverage — know your baseline
go test -cover ./...
```

This is not optional. If your Week 1 code has `gofmt` or `go vet` issues, fix them Monday morning before starting Chapter 1. Clean code is a non-negotiable habit.

---

## Detailed Daily Plan

---

### Monday, March 30 (Target: 2.5 hours)

**Focus: TGPL Chapter 1 — A Tutorial**

Chapter 1 is a fast walkthrough of Go's most common features. The examples are short but dense. Do not skim. Type every example.

#### Session 1: TGPL Chapter 1 (90 min)

- [ ] Read Chapter 1 cover to cover
- [ ] Type out (do NOT copy-paste) every code example — even the ones that look obvious
- [ ] Run every example: `go run filename.go`
- [ ] For each example, answer in your notes: *"What would happen if I changed X?"* — then try it
- [ ] Take notes in: `notes/books/tgpl-chapter-1.md`

**Key topics in Chapter 1 to watch for**:
- `os.Args` for command-line input
- `bufio.Scanner` for reading stdin
- `http.Get` for a simple web fetch
- Named types (`type Celsius float64`) — this pattern appears throughout Go
- How Go handles errors vs. how other languages do

#### Session 2: Start Gophercises #1 — Echo CLI (60 min)

URL: https://gophercises.com/ → Exercise 1

- [ ] Read the full problem statement before writing any code (5 min)
- [ ] Write a plain-English plan in comments before any Go code (10 min):

```go
// Plan:
// 1. Read all command-line args using os.Args
// 2. Join them with spaces
// 3. Print to stdout
// 4. Handle edge case: no args provided
```

- [ ] Build the CLI tool using only `os` and `fmt` — no external libraries
- [ ] Run it manually with test inputs before writing formal tests
- [ ] Write your first 2 tests in `echo_test.go`

**Testing requirement**: Every function in Gophercises must have at least one test. No exceptions starting this week.

#### Commit (10 min)

```bash
git add .
git commit -m "week-2: tgpl chapter 1 notes + gophercises-01 echo cli start"
git push
```

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Tuesday, March 31 (Target: 2 hours)

**Focus: Finish Gophercises #1 with tests + TGPL Chapter 2**

#### Session 1: Finish Gophercises #1 + Write Tests (60 min)

- [ ] Complete the echo CLI if not done
- [ ] Write table-driven tests — this is the Go standard, not optional:

```go
func TestEcho(t *testing.T) {
    tests := []struct {
        name     string
        args     []string
        expected string
    }{
        {"single arg", []string{"hello"}, "hello"},
        {"multiple args", []string{"hello", "world"}, "hello world"},
        {"no args", []string{}, ""},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // your assertion here
        })
    }
}
```

- [ ] Run `go test -v -cover ./...` — aim for >70% coverage even at this stage
- [ ] Run `go vet ./...` — fix any issues
- [ ] Push finished Gophercises #1

#### Session 2: TGPL Chapter 2 — Program Structure (60 min)

- [ ] Read Chapter 2 carefully — this chapter defines how Go programs are structured
- [ ] Take notes in: `notes/books/tgpl-chapter-2.md`
- [ ] For every concept, write a 2-line code snippet in your notes to anchor it

**Key concepts in Chapter 2** — understand each before Thursday:

| Concept | One-line explanation (write yours) |
|---|---|
| Package declarations | |
| Short variable declaration `:=` vs `var` | |
| Zero values for each type | |
| Pointers — `*T` vs `&x` | |
| `new(T)` — what it returns, when to use it | |
| Named types (`type Celsius float64`) | |
| Blank identifier `_` | |
| Package-level vs function-level scope | |
| Import paths and package naming | |

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Wednesday, April 1 (Target: 2.5 hours)

**Focus: TGPL Chapter 3 + Start Gophercises #2**

#### Session 1: TGPL Chapter 3 — Basic Data Types (90 min)

- [ ] Read Chapter 3: integers, floats, complex numbers, booleans, strings, constants
- [ ] Take notes in: `notes/books/tgpl-chapter-3.md`
- [ ] Type and run the string examples — Go strings are immutable byte sequences, not character arrays. This distinction matters later.
- [ ] Pay close attention to **constants and `iota`** — this pattern appears in professional Go code constantly

**Things that surprise many beginners in Chapter 3**:
- Integer overflow behaviour in Go (no exceptions — it wraps silently)
- The difference between `len(s)` on a string vs. number of Unicode characters
- `rune` as an alias for `int32` representing a Unicode code point
- Why you cannot compare `float64` values with `==` reliably

Write your understanding of each in your notes.

#### Session 2: Start Gophercises #2 — Quiz Game (60 min)

URL: https://gophercises.com/ → Exercise 2

The quiz game reads a CSV file of questions/answers and quizzes the user.

- [ ] Read the full problem statement
- [ ] Plan before coding — write your plan in comments:

```go
// Plan:
// 1. Open CSV file from flag (-csv defaults to "problems.csv")
// 2. Parse each row: question, answer
// 3. Loop through questions, prompt user, read answer
// 4. Track correct count
// 5. Print final score
// Key: use -limit flag for timed version (stretch)
```

- [ ] Use the `flag` package for command-line flags — not `os.Args` directly
- [ ] Use `encoding/csv` to parse the file — read the docs first: https://pkg.go.dev/encoding/csv
- [ ] Build the basic version (no timer) before the timed version

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Thursday, April 2 (Target: 2 hours)

**Focus: TGPL Chapter 4 + Finish Gophercises #2**

#### Session 1: TGPL Chapter 4 — Composite Types (90 min)

- [ ] Read Chapter 4: arrays, slices, maps, structs
- [ ] Take notes in: `notes/books/tgpl-chapter-4.md`
- [ ] This chapter is dense — go slower than Chapter 1

**Critical concepts in Chapter 4**:

**Arrays vs Slices**:
- Arrays have fixed size and are value types (copying an array copies all elements)
- Slices are references to an underlying array — this is how Go passes "arrays" in practice
- `append()` behaviour when capacity is exceeded (new underlying array allocated)
- Slice tricks: `s[i:j]`, `copy()`, nil slice vs empty slice

```go
var s []int        // nil slice — len 0, cap 0
s = []int{}        // empty slice — len 0, cap 0
// Both have len=0, but nil slice != empty slice
fmt.Println(s == nil)  // true for var s, false for []int{}
```

Write this distinction in your own words in your notes.

**Maps**:
- Maps are reference types, like slices
- Zero value of a map is nil — you cannot write to a nil map
- The two-value assignment for key lookup: `v, ok := m[key]`
- Iteration order is randomized by design

**Structs**:
- Struct embedding — this is Go's composition model (there is no inheritance)
- Anonymous fields and promoted methods
- Struct tags for JSON encoding

#### Session 2: Finish Gophercises #2 + Add Timer (30 min)

- [ ] Add the `-limit` flag for a timed quiz using `time.NewTimer`:

```go
// Hint: use select with two cases
select {
case <-timer.C:
    // time's up
case answer := <-answerCh:
    // process answer
}
```

- [ ] Write table-driven tests for quiz scoring logic
- [ ] Run `go test -v -cover ./...` — aim for >75% coverage
- [ ] Push to GitHub

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Friday, April 3 (Target: 1.5 hours)

**Focus: Code Quality + Go Toolchain Discipline**

This day is deliberately not about new content. It is about going back to everything you wrote this week and making it professional.

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

Fix every `go vet` warning. No exceptions. These are real bugs.

- [ ] All files pass `gofmt`
- [ ] All files pass `go vet`
- [ ] All packages have at least 60% test coverage (70%+ preferred)
- [ ] No `TODO` or `FIXME` comments left without a note in your journal

#### Session 2: Explore the Go Toolchain (45 min)

Spend 45 minutes systematically exploring tools you will use for the next 12 months:

```bash
# Documentation for any package
go doc fmt
go doc fmt.Fprintf
go doc os.Args

# List all available commands
go help

# Build without running
go build ./...

# Check module dependencies
go mod tidy
go list -m all
```

- [ ] Read the `go help` output fully — you should know what every subcommand does
- [ ] Look up 3 packages you used this week on https://pkg.go.dev/ — read the full docs page, not just the function signatures
- [ ] Bookmark the Go standard library index: https://pkg.go.dev/std

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Saturday, April 4 (Target: 2.5 hours)

**Focus: Independent Mini-Project — No Guidance**

This is your test. No Gophercises instructions. No Tour exercises. No example code. You choose the problem, plan it yourself, and build it.

Choose ONE:

---

**Option A: Word Frequency Counter (Recommended)**

Build a CLI tool that reads a text file and prints the top N most frequent words.

Requirements (you must implement all):
- Accept filename and `-n` flag (default: 10) for how many top words to show
- Strip punctuation from words
- Case-insensitive counting (`"The"` and `"the"` are the same word)
- Sort output by frequency descending, then alphabetically for ties
- Handle missing file gracefully with a useful error message
- Write at least 5 table-driven tests

Skills this uses: `os.Open`, `bufio.Scanner`, `strings`, `sort`, maps, structs, `flag`, `fmt`, error handling, testing

---

**Option B: Log Level Filter**

Build a CLI tool that reads a log file and filters lines by log level.

Requirements:
- Accept filename and `-level` flag (values: DEBUG, INFO, WARN, ERROR)
- Print only lines at or above the specified level
- Print total count of each level found (even filtered ones) at the end
- Handle malformed lines gracefully
- Write at least 5 table-driven tests

Skills this uses: `os.Open`, `bufio.Scanner`, `strings`, `flag`, `fmt`, error handling, testing

---

**Option C: Number Base Converter**

Build a CLI tool that converts numbers between bases (binary, octal, decimal, hex).

Requirements:
- Accept input number, from-base, and to-base as flags
- Support bases 2, 8, 10, 16
- Validate input is a valid number in the from-base
- Print the result and a human-readable explanation
- Write at least 8 table-driven tests covering edge cases

Skills this uses: `strconv`, `fmt`, `flag`, `math`, error handling, testing

---

**For whichever option you choose**:

- [ ] Write a comment plan before any Go code (10 min)
- [ ] Build it (75 min)
- [ ] Write tests as you build — not after (30 min)
- [ ] Run `go vet ./...` and fix everything
- [ ] Run `go test -v -cover ./...` — target >80% coverage
- [ ] Add a `README.md` (5 sentences: what it does, how to build, example usage)
- [ ] Push to GitHub: `phase-1/week-2-mini-project/`
- [ ] Commit message: `feat: week-2 independent mini-project — [option name]`

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Sunday, April 5 (Target: 1 hour)

**Focus: Weekly Review + Plan Week 3**

#### Part 1: Journal Entry (30 min)

Create your Week 2 journal:
```bash
cp journal/JOURNAL-TEMPLATE.md journal/2026-04-week-13.md
```

Fill in every section. Specific requirements for this week's journal:

**Technical Concepts section** — write at least 5 specific things, not vague summaries. Bad: "I learned about slices." Good: "Slices in Go are not arrays — they are a three-field struct (pointer, length, capacity) pointing at an underlying array. Appending beyond capacity allocates a new array, which means the original slice variable is unaffected."

**Struggles section** — list every concept you are not yet solid on. This is your Week 3 study list.

**Questions to Research** — list every question that came up this week that you did not have time to investigate fully.

#### Part 2: Self-Assessment (15 min)

Rate your current confidence on TGPL topics so far (1 = shaky, 5 = solid):

| Topic | Chapter | Confidence |
|---|---|---|
| Package structure and imports | 1 | |
| os.Args, bufio.Scanner | 1 | |
| Named types | 1 | |
| Short declarations vs var | 2 | |
| Pointers — `*T`, `&x`, `new(T)` | 2 | |
| Zero values | 2 | |
| Integer types and overflow | 3 | |
| String immutability, runes, bytes | 3 | |
| Constants and iota | 3 | |
| Arrays vs slices (the difference) | 4 | |
| append() and capacity growth | 4 | |
| Maps — nil vs empty, two-value lookup | 4 | |
| Struct embedding and composition | 4 | |

Any rating of 1 or 2 → add a focused review session to Week 3 before continuing to Chapter 5.

#### Part 3: Week 3 Planning (15 min)

Open `roadmap/PHASE-1-MODULES.md` → Module 1, Submodule 1.1 and confirm you are on pace.

Week 3 target:
- TGPL Chapters 5–6 (functions, methods)
- Gophercises #3 and #4
- Start planning Project 1.1.1: Log Parser CLI (your first real Module 1.1 project)

Write next week's 5 goals in `roadmap/weekly-goals.md`.

**Actual time spent**: ___ hours
**What I accomplished**:

---

## Week 2 Summary

### Total Time Logged
**Planned**: 12 hours
**Actual**: ___ hours
**Variance**: ___ hours (over/under)
**Running Total (Phase 1)**: ___ / 200 hours (Module 1.1 target)

### Deliverables Check

- [ ] TGPL Chapters 1–4 read with notes taken in `notes/books/tgpl-chapter-[1-4].md`
- [ ] Gophercises #1 (echo) — done + tests written + pushed to GitHub
- [ ] Gophercises #2 (quiz) — done + tests written + pushed to GitHub
- [ ] 1 independent mini-project — done + tests + README + pushed
- [ ] All code passes `gofmt` and `go vet`
- [ ] All packages have >70% test coverage (`go test -cover ./...`)
- [ ] Daily commits to GitHub — at least 5 this week
- [ ] Week 2 journal filled in
- [ ] Week 3 goals written in `roadmap/weekly-goals.md`
- [ ] Running hours total updated in `roadmap/TIME-TRACKING.md`

### Gophercises Tracker (running total)

| # | Exercise | Status | Tests Written | Coverage |
|---|---|---|---|---|
| 1 | Echo CLI | | | |
| 2 | Quiz Game | | | |
| 3 | | Not started | | |
| ... | | | | |
| 20 | | Not started | | |

Copy this table to your journal and keep it updated every week.

---

## Concepts You Must Be Able to Explain by End of Week 2

Answer these without looking anything up. If you cannot, go back.

**TGPL Chapter 1–2**
1. What is the difference between `:=` and `var`? When can you use each?
2. What is the zero value of `*int`? Of `map[string]int`? Of `[]string`?
3. What does `new(T)` return? How is it different from `var x T; &x`?
4. What is a named type? Give an example and explain why you would use one.
5. What is the blank identifier and give two scenarios where you would use it?

**TGPL Chapter 3**
6. Why should you not compare `float64` values with `==`?
7. What is a `rune`? How is it different from a `byte`?
8. Why is `len("hello")` sometimes not equal to the number of characters in the string?
9. What does `iota` do inside a `const` block?

**TGPL Chapter 4**
10. What is the difference between `var s []int` and `s := []int{}`?
11. What happens when you `append()` to a slice that has no remaining capacity?
12. What does `v, ok := m[key]` do when the key does not exist in the map?
13. Can you write to a nil map? What happens if you try?
14. What is struct embedding? How does it differ from traditional inheritance?

**Testing**
15. What is a table-driven test and why does the Go community prefer it?
16. What does `go test -cover ./...` tell you and what does it NOT tell you?

---

## Code Quality Standards — Non-Negotiable from Week 2 Onward

From `roadmap/PHASE-1-MODULES.md` — all projects require:

| Standard | Requirement | How to Check |
|---|---|---|
| Formatting | All files formatted | `gofmt -l .` returns nothing |
| Vet | No vet warnings | `go vet ./...` returns nothing |
| Test coverage | >80% for projects, >70% for exercises | `go test -cover ./...` |
| Error handling | Every error is handled — no `_` on errors | Manual review |
| Comments | Exported functions have doc comments | `go doc ./...` |
| README | Every project has a README | Present in repo |

These are not suggestions. Starting Week 2, every piece of code you push must meet these standards. This is what "idiomatic Go" means in practice.

---

## If You Are Behind

**Didn't finish TGPL Chapters 3–4?**
→ Do not skip ahead. TGPL builds sequentially. Chapter 5 (functions) relies on Chapter 4 (composite types).
→ Move unfinished chapters to the start of Week 3. Note in your journal why you fell behind and what you will change.

**Couldn't hit 12 hours?**
→ Log exactly how many you got. Do not round up.
→ Identify the specific days/sessions you missed. Write down why.
→ The 200-hour Module 1.1 target cannot be met at 8h/week. You need to average 12h+.

**Gophercises #2 timer is confusing?**
→ Start with `time.After(d)` — it returns a `<-chan time.Time` you can use directly in `select`
→ Read: https://pkg.go.dev/time#After
→ Spend 30 minutes stuck before looking at any examples

**Tests feel uncomfortable?**
→ Good. You are building a skill.
→ Start every test file by writing the test function signatures before the implementation function. This forces you to think about what you are testing.

---

## Week 3 Preview

By end of Week 3 you should have:
- TGPL Chapters 5–6 (functions, methods — these are dense)
- Gophercises #3 and #4 with tests
- A written plan for Project 1.1.1: Log Parser CLI

The Log Parser is your first real Module 1.1 project. It is not a mini-exercise. It requires multiple files, proper package structure, error handling, and >80% test coverage. Week 3 is your last preparation week before project work begins in Week 4.

---

## Commitment Tracker

Did you uphold all 5 commitments this week?

1. [ ] **12 focused hours** — no rounding up, no passive watching
2. [ ] **Sequential learning** — TGPL Ch 1-4 only, did not jump ahead
3. [ ] **No AI for writing code** — you wrote and debugged everything yourself
4. [ ] **Communication practice** — notes, journal, README, doc comments written
5. [ ] **Honest tracking** — every hour logged with specifics

**Streak**: ___ weeks consistent (Week 1 + Week 2 = 2 if you complete this)

---

## Notes

(Blockers, observations, questions to carry into Week 3)




---

**Phase 1 / Module 1.1 running hours**: ___ / 200
**Phase 1 total running hours**: ___ / 600
**Projects completed**: 0 / 3 (Log Parser, Todo API, Web Crawler)
**Gophercises completed**: ___ / 20

**Next**: After completing Week 2, open `roadmap/PHASE-1-MODULES.md` → Submodule 1.1 and verify you are on pace. Then open `Year-2026/Month-04/WEEK-3-START-HERE.md` (create this file on Sunday).