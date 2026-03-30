# Week 4: TGPL Chapters 7–9 + Project 1.1.1 Start + Concurrency Foundations

**Week**: Week 4 of Phase 1 — Module 1.1 → Module 1.2 Transition (Go Fundamentals → Concurrency)
**Dates**: April 13 – April 19, 2026
**Total Hours Goal**: 13 hours (Module 1.1/1.2 requires 200 hours over ~16 weeks → average 12.5h/week)
**Location**: `Year-2026/Month-03/WEEK-4-START-HERE.md`
**Roadmap Reference**: `roadmap/PHASE-1-MODULES.md` → Module 1, Submodule 1.2 (Concurrency)

---

## ⚠️ Read This Before You Start

Week 4 is where Go becomes truly powerful — and truly difficult. You leave behind sequential programming and enter concurrency.

- **Week 1-3**: Sequential programming, functions, types, methods
- **Week 4+**: Goroutines, channels, concurrent patterns, coordination

This week will change how you think about programs. Embrace the mental difficulty. This is the foundation for everything that makes Go special.

**WARNING**: Concurrency is hard. You will have race conditions. You will have deadlocks. Your goroutines will hang. This is normal and expected. Every developer goes through this phase.

### Module 1.1/1.2 Progress Check

From `roadmap/PHASE-1-MODULES.md` — Module 1.1 nearly complete, Module 1.2 begins:

| Requirement | Target | Your Status After Week 3 |
|---|---|---|
| TGPL Book — all chapters + exercises | Complete | ~50% (Ch 1-6) |
| Gophercises completed | 20+ | 4/20 |
| Phase 1 Projects on GitHub (>80% coverage) | 3 projects | 2/3 (+ Project 1.1.1 starts) |
| Can explain goroutines, channels, select | Yes | Not yet (THIS WEEK) |
| Idiomatic Go — code review approved | Yes | Not started |
| External code review received | ≥1 project | Not started |
| Hours logged | 200 hours | ~35h |

**Week 4 advances**: Concurrency foundations, concurrent programming patterns, Project 1.1.1 (Log Parser CLI) production-ready version.

---

## Week 4 Goals

**Primary Goal**: Complete TGPL Chapters 7–9 with deep focus on goroutines, channels, and select. Build Project 1.1.1 (Log Parser CLI) with streaming, error handling, and multiple output formats. Start your first real project.

**Secondary Goals**:
- Understand goroutines and how they scale beyond OS threads
- Master channels: unbuffered, buffered, directional, select
- Learn WaitGroup, sync.Mutex, sync.Once for coordination
- Write concurrent programs that are correct (no data races, no deadlocks)
- Build testing strategy for concurrent code
- Begin Gophercises #5 and #6 (likely concurrent exercises)
- Commit to GitHub daily with meaningful commit messages
- Log hours daily — never batch on Sunday

**Hours Target**: 13 hours (Module 1.2 heavily weighted; concurrency takes time to understand correctly)

**Reality Check**: This week may go slower than Week 2-3. Concurrency is harder. That is okay. Depth > speed.

---

## Key Resources for This Week

| Resource | Use For | Link |
|---|---|---|
| TGPL Book (Donovan & Kernighan) | Main reading — Ch 7–9 | Your copy |
| Gophercises | Exercises #5, #6 (if concurrent) | https://gophercises.com/ |
| Go by Example — Goroutines | Goroutine patterns | https://gobyexample.com/goroutines |
| Go by Example — Channels | Channel patterns | https://gobyexample.com/channels |
| Go by Example — Select | Select statement | https://gobyexample.com/select |
| pkg.go.dev — sync | Sync primitives | https://pkg.go.dev/sync |
| pkg.go.dev — context | Context and cancellation | https://pkg.go.dev/context |
| Go Memory Model | Understanding concurrency semantics | https://go.dev/ref/mem |
| Go Race Detector | Finding data races | `go test -race` |

---

## Before You Write Any Code This Week

Check that all code from Week 3 still passes quality checks and run the race detector:

```bash
# Make sure nothing regressed
gofmt -l .

# All vet warnings must still be resolved
go vet ./...

# Test coverage should not have dropped
go test -cover ./...

# NEW: Check for data races in previous concurrent code
go test -race ./...
```

Do NOT start Week 4 with quality debt from Week 3.

---

## Critical Pre-Reading: The Go Memory Model

Before you write ANY concurrent code, read this document fully:

**https://go.dev/ref/mem**

This explains the guarantees Go makes about memory visibility in concurrent programs. Key concepts:

- **Atomicity**: Incrementing a variable is NOT atomic in Go; use `sync/atomic` or mutex
- **Happens-before**: The memory model defines ordering guarantees for concurrent operations
- **Channels**: Channel operations have specific synchronization semantics (guaranteed to communicate)
- **Data races**: Two goroutines accessing the same memory without synchronization is a bug

You cannot write correct concurrent code without understanding these guarantees.

---

## Detailed Daily Plan

### Monday, April 13 (Target: 2.5 hours)

**Focus: TGPL Chapter 7 — Goroutines and Channels**

Chapter 7 is the foundation of Go's concurrency model. This chapter requires deeper study than earlier chapters.

#### Session 1: TGPL Chapter 7 — Goroutines Fundamentals (90 min)

- [ ] Read Chapter 7 cover to cover — **do NOT skim** — this is dense content
- [ ] Type out every code example — even ones that look trivial
- [ ] Run every example: `go run filename.go`
- [ ] For each example, answer in your notes: *"When would this pattern cause a data race?"* and *"How would I fix it?"*
- [ ] Take notes in: `notes/books/tgpl-chapter-7.md`

**Key topics in Chapter 7 to watch for**:
- Goroutines: lightweight concurrency primitive managed by the Go runtime
- Creating goroutines with the `go` keyword
- Channels: typed pipes for communication between goroutines
- Unbuffered vs. buffered channels (blocking behavior)
- Channel directions: send-only `chan<-` and receive-only `<-chan`
- The select statement: multiplexing on multiple channels
- Closing channels and detecting closed channels
- The for-range pattern with channels

**Things beginners often struggle with in Chapter 7**:
- Understanding when a goroutine blocks (on channel send/receive)
- Unbuffered channels block the sender until a receiver is ready
- If you close a channel with senders still trying to send: **panic**
- Receiving from a closed channel returns the zero value forever, but `x, ok := <-ch` tells you if channel is closed
- `select` with multiple cases: GO RANDOMLY PICKS ONE, not the first one listed

**Critical Concept: The Goroutine Leak**

A common bug in concurrent Go:

```go
// BAD: goroutine leak
func (a *App) fetch() {
    go func() {
        // This goroutine may never exit if there's nowhere
        // to read the result. The goroutine is a leak.
        result := someSlowOp()
        resultChan <- result
    }()
}

// GOOD: goroutine with bounded lifetime
func (a *App) fetch(ctx context.Context) {
    go func() {
        result := someSlowOp()
        select {
        case resultChan <- result:
            // Success
        case <-ctx.Done():
            // Context cancelled, exit cleanly
            return
        }
    }()
}
```

Spend extra time understanding goroutine lifetimes. A goroutine should have a clear entry AND exit point.

#### Session 2: Start Gophercises #5 (60 min)

URL: https://gophercises.com/ → Exercise 5

Gophercises #5 likely involves concurrency. Read the problem fully first.

- [ ] Read the full problem statement
- [ ] Plan before coding — write your plan in comments:

```go
// Plan:
// 1. [What is the problem asking for?]
// 2. [What goroutines will I create?]
// 3. [How will they communicate? (channels)]
// 4. [How will I know when they are all done?]
// 5. [What timeouts or context cancellation do I need?]
```

- [ ] Start with a simple version (possibly sequential first, then add goroutines)
- [ ] Write at least 2 tests (one for correct behavior, one for timeout/cancellation)
- [ ] Use `context.WithTimeout()` for bounded execution

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Tuesday, April 14 (Target: 2 hours)

**Focus: Finish Gophercises #5 + TGPL Chapter 8 Deep Dive**

#### Session 1: Finish Gophercises #5 + Write Tests (60 min)

- [ ] Complete the exercise
- [ ] Write table-driven tests that exercise both happy path and edge cases
- [ ] Include a test that verifies goroutines don't leak (this is hard — read up on it)
- [ ] Run `go test -v -cover ./...` — aim for >75% coverage
- [ ] Run `go test -race ./...` — **NO race detector warnings allowed**
- [ ] Push finished Gophercises #5

**Testing concurrent code is hard**. Common patterns:

```go
// Test that operation completes within timeout
func TestWithTimeout(t *testing.T) {
    ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
    defer cancel()
    
    result := doSomethingConcurrent(ctx)
    if result == nil {
        t.Fatal("operation timed out or failed")
    }
}

// Test with race detector
// Run: go test -race ./...
// The race detector will flag any data races
```

#### Session 2: TGPL Chapter 8 — Sync Primitives (60 min)

- [ ] Read Chapter 8: sync package, WaitGroup, Mutex, Once, etc.
- [ ] Take notes in: `notes/books/tgpl-chapter-8.md`
- [ ] For every concept, write a code example in your notes

**Key concepts in Chapter 8**:

| Concept | One-line explanation (write yours) |
|---|---|
| sync.WaitGroup — when/why to use | |
| sync.Mutex — protection of shared state | |
| sync.Once — guaranteed single execution | |
| sync.RWMutex — readers vs writers | |
| Channel vs Mutex for coordination | |
| Memory order guarantees from sync | |

**Critical Concept: Mutex vs Channel**

Go offers two tools for coordination:

```go
// Pattern 1: Channels (preferred for communication)
// Use when goroutines need to exchange data
func process(resultChan chan Result) {
    result := compute()
    resultChan <- result  // Send data AND sync
}

// Pattern 2: Mutex (preferred for protecting state)
// Use when multiple goroutines access shared state
type Counter struct {
    mu    sync.Mutex
    count int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    c.count++
    c.mu.Unlock()
}
```

Learn this distinction deeply. Choose the wrong tool and your code will be either wrong or inefficient.

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Wednesday, April 15 (Target: 2.5 hours)

**Focus: Context and Cancellation + Gophercises #6**

#### Session 1: TGPL Chapter 9 — Context & Cancellation (90 min)

- [ ] Read Chapter 9 (or relevant section on context)
- [ ] If TGPL doesn't cover context deeply, read: https://pkg.go.dev/context and https://go.dev/blog/context
- [ ] Take notes in: `notes/books/tgpl-chapter-9.md`
- [ ] Type every example and understand the cancellation patterns

**Key concepts in context**:

| Concept | What to Understand |
|---|---|
| context.Background() vs context.TODO() | Root contexts for background work vs RPC handlers |
| context.WithCancel() | Manual cancellation |
| context.WithTimeout() and With Deadline() | Time-based cancellation |
| context.WithValue() | Passing values through context (use sparingly) |
| Listening to `<-ctx.Done()` | Clean shutdown signals |

**Critical Pattern: Graceful Shutdown with Context**

This pattern appears in every production Go service:

```go
func worker(ctx context.Context, jobs <-chan Job) {
    for {
        select {
        case job, ok := <-jobs:
            if !ok {
                // Channel closed, exit cleanly
                return
            }
            processJob(job)
        case <-ctx.Done():
            // Context cancelled, stop accepting new jobs
            return
        }
    }
}

// In main:
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

// Start workers with context
for i := 0; i < numWorkers; i++ {
    go worker(ctx, jobs)
}

// When timeout expires, all workers exit gracefully
```

Master this pattern. You will use it constantly.

#### Session 2: Start Gophercises #6 (60 min)

URL: https://gophercises.com/ → Exercise 6

- [ ] Read the full problem statement
- [ ] Plan before coding:

```go
// Plan:
// 1. [What is the problem asking for?]
// 2. [Where will I use goroutines?]
// 3. [What context or cancellation do I need?]
// 4. [How will I handle errors from concurrent operations?]
```

- [ ] Build the basic version first
- [ ] Write 2 tests focusing on error handling and timeouts
- [ ] Use `-race` flag when testing: `go test -race ./...`

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Thursday, April 16 (Target: 2 hours)

**Focus: Finish Gophercises #6 + Project 1.1.1 Production Version**

#### Session 1: Finish Gophercises #6 + Race Testing (60 min)

- [ ] Complete the exercise
- [ ] Write comprehensive tests covering:
  - Happy path (everything works)
  - Timeout scenario (context times out)
  - Cancellation scenario (context is cancelled)
  - Error propagation (goroutine returns error, must be handled)
- [ ] Run `go test -race ./...` — fix any race conditions
- [ ] Aim for >80% coverage
- [ ] Push finished Gophercises #6

#### Session 2: Project 1.1.1 — Log Parser CLI (Production Version) (60 min)

This week you build the production version of your first real project. Week 2's mini-project was learning; this is different — you are building something with real requirements.

**What This Project Does**:
- Reads log files and provides analytics
- Filters by log level (DEBUG, INFO, WARN, ERROR)
- Outputs statistics in multiple formats (JSON, CSV, plaintext)
- Handles large files efficiently (streaming, not loading into RAM)
- Exposed as a CLI tool with clear interface

**Your PLAN.md should already exist from Week 3**. Now code:

**Directory structure**:
```
projects/phase-1/log-parser-cli/
├── main.go           (CLI entry, flag parsing, orchestration)
├── parser.go         (reading and parsing log lines)
├── analyzer.go       (aggregating statistics)
├── formatter.go      (JSON, CSV, plaintext output writers)
├── main_test.go
├── parser_test.go
├── analyzer_test.go
├── formatter_test.go
├── README.md
└── example-logs.txt  (test data)
```

**Key requirements**:
- [ ] CLI flags: `-file` (path), `-level` (DEBUG|INFO|WARN|ERROR), `-format` (json|csv|text)
- [ ] Read files via streaming (don't load entire file into memory)
- [ ] Parse log records: look for pattern like `[TIMESTAMP] [LEVEL] message`
- [ ] Handle malformed lines gracefully (log and skip, do NOT crash)
- [ ] Output statistics: total lines, count per level, percentages
- [ ] Return non-zero exit code on errors

**Minimum Implementation**:
- [ ] Basic parsing (90 min is not enough for perfection)
- [ ] Works with real log files
- [ ] At least 60% test coverage
- [ ] Proper error handling

**Session Target**: Get the basic structure in place and 60%+ coverage. Friday and Saturday you will refine.

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Friday, April 17 (Target: 1.5 hours)

**Focus: Project 1.1.1 Refinement + Code Quality**

#### Session 1: Code Quality Audit (45 min)

Run this on ALL code written this week (Gophercises + Project):

```bash
# 1. Format every file
gofmt -w .

# 2. Check for common mistakes
go vet ./...

# 3. Check for race conditions
go test -race -cover ./...

# 4. Run tests verbosely
go test -v ./...
```

Fix every `go vet` AND every race condition warning. No exceptions. Race conditions are real bugs.

- [ ] All files pass `gofmt`
- [ ] All files pass `go vet`
- [ ] All packages pass `go test -race` with no warnings
- [ ] All packages have at least 70% test coverage
- [ ] No `TODO` or `FIXME` without context

**New Standard for Week 4+**: Every exported function in concurrent code must have a comment explaining goroutine safety:

```go
// ProcessWithTimeout processes items concurrently with a timeout.
// It is safe to call from multiple goroutines.
// Each goroutine runs independently; results are synchronized via the returned channel.
func ProcessWithTimeout(ctx context.Context, items []Item) <-chan Result {
    // ...
}
```

#### Session 2: Project 1.1.1 Enhancement (45 min)

Build out Log Parser CLI to handle more:

- [ ] Add support for filtering by level (not just output all)
- [ ] Implement at least 2 output formats (JSON and plaintext minimum)
- [ ] Add a feature: time-range filtering (optional but useful)
- [ ] Write README with clear usage examples
- [ ] Test with a real log file (generate one with > 1000 lines)

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Saturday, April 18 (Target: 2.5 hours)

**Focus: Project 1.1.1 Completion + Concurrent Exercises**

#### Session 1: Finish Log Parser CLI (90 min)

Complete everything for Project 1.1.1:

- [ ] All features working (from PLAN.md)
- [ ] `go test -race -v -cover ./...` passes with >80% coverage
- [ ] `go vet ./...` passes with no warnings
- [ ] README includes:
  - What it does (2 sentences)
  - How to build (`go build -o log-parser .`)
  - How to run (3 example commands)
  - Design decisions (1-2 sentences on why you chose your architecture)
- [ ] Commit structure: multiple commits, each meaningful
  - Example: `feat: log parser streaming reader`, `feat: json output formatter`, `fix: handle malformed lines`

**Push to GitHub**: `projects/phase-1/log-parser-cli/`

#### Session 2: Independent Concurrency Exercise (60 min)

This is NOT a Gophercises problem. You design and build it.

**Choose ONE** (these are harder than Week 3 mini-projects):

---

**Option A: Concurrent File Processor (Recommended)**

Build a CLI tool that processes files concurrently.

Requirements:
- Accept a directory path and file pattern (glob)
- Concurrently process each matched file (transform the content)
- Use 3-5 worker goroutines (configurable with `-workers` flag)
- Write outputs to a results directory
- Track success/failure counts
- Handle timeouts with `-timeout` flag
- Write at least 4 tests (happy path, timeout, error handling, empty directory)
- Target >85% coverage

Skills this uses: goroutines, channels, WaitGroup, context, file I/O, testing concurrent code

---

**Option B: Concurrent HTTP Scraper**

Build a tool that fetches URLs concurrently.

Requirements:
- Accept file with list of URLs (one per line)
- Fetch each URL concurrently (3-5 goroutines, configurable)
- Record status code and fetch time for each
- Support `-timeout` flag for individual requests
- Output summary: total URLs, successes, failures, average time
- Handle errors gracefully (bad URLs, network errors, timeouts)
- Write at least 4 tests (success, timeouts, network errors, empty list)
- Target >80% coverage

Skills this uses: goroutines, net/http, channels, sync.WaitGroup, time measurement, testing

---

**Option C: Concurrent Data Pipeline**

Build a tool that implements a pipeline pattern.

Requirements:
- Read numbers from stdin (or file with `-input` flag)
- Stage 1 (reader goroutine): Read numbers and send to channel
- Stage 2 (worker goroutines): Square the number, send to next channel
- Stage 3 (writer goroutine): Collect results and write summary
- Use 3 concurrent workers for squaring
- Support `-timeout` for the entire operation
- Output: original numbers, squared results, timing stats
- Write at least 5 tests (normal flow, empty input, timeout, errors)
- Target >85% coverage

Skills this uses: goroutines, channels, pipeline pattern, context, testing concurrent pipelines

---

**For whichever option you choose**:

- [ ] Write a design document (not code) — what goroutines? how do they communicate?
- [ ] Build it with goroutines, channels, and context (not sequential!)
- [ ] Write tests covering concurrency scenarios (timeout, cancellation, errors)
- [ ] Run `go test -race ./...` — **ZERO race conditions allowed**
- [ ] Run `go test -v -cover ./...` — target >80% coverage
- [ ] Add a README
- [ ] Push to GitHub: `projects/phase-1/week-4-concurrent-exercise/`
- [ ] Commit message: `feat: week-4 concurrent exercise — [option name]`

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Sunday, April 19 (Target: 1 hour)

**Focus: Weekly Review + Plan Week 5**

#### Part 1: Journal Entry (30 min)

Create your Week 4 journal:
```bash
cp journal/JOURNAL-TEMPLATE.md journal/2026-04-week-16.md
```

Fill in every section. Specific requirements for this week's journal:

**Technical Concepts section** — write 6+ specific things about concurrency. Examples:
- "Unbuffered channels block. I spent 2 hours debugging a deadlock caused by trying to send to a channel with no receiver."
- "The `select` statement picks a random ready case. I kept assuming it would pick the first one, but that is not how Go works."
- "Race conditions are subtle. `go test -race` found bugs that sequential thinking would never catch. I need to think about every shared memory access."
- "Context is the Go way to handle cancellation and timeouts. Once I understood `<-ctx.Done()`, many patterns became clear."
- "Goroutine leaks are real. A goroutine waiting forever on a channel is a bug. I now always have an exit condition."

**Struggles section** — be specific about concurrency learning:
- "Designing goroutine lifetimes is hard. I kept creating goroutines without clear exit conditions."
- "Testing concurrent code is a new skill. `go test -race` helped me find bugs I couldn't think of."
- "Channels feel magical at first but are just synchronized communication. Understanding the blocking semantics took a full day."

**Questions to Research**:
- "What are the performance implications of buffered vs unbuffered channels?"
- "How do I design a concurrent system that is both correct and efficient?"

#### Part 2: Self-Assessment (15 min)

Rate your current confidence on all TGPL topics (1 = shaky, 5 = expert):

| Topic | Chapter | Confidence |
|---|---|---|
| Goroutines (creation, lifecycle, scaling) | 7 | |
| Unbuffered channels | 7 | |
| Buffered channels | 7 | |
| Select statement and multiplexing | 7 | |
| Channel directions (send-only, receive-only) | 7 | |
| Closing channels and detection | 7 | |
| WaitGroup synchronization | 8 | |
| Mutex for protecting state | 8 | |
| Once guarantee | 8 | |
| Context and cancellation | 9 | |
| Timeouts with context | 9 | |
| Graceful shutdown patterns | 9 | |

Any rating of 1-2 → add focused review to Week 5. Any rating 3-4 → practice more this week or next.

#### Part 3: Retrospective on Project 1.1.1

Answer these honestly:

- [ ] What went well with the Log Parser project?
- [ ] What was difficult?
- [ ] How confident am I that this code will work on real log files?
- [ ] Would I feel comfortable showing this code in a job interview? Why/why not?
- [ ] What would I do differently if I built it again?

#### Part 4: Week 5 Planning (15 min)

Week 5 target (from PHASE-1-MODULES.md):
- Deepen concurrency mastery (reread Chapters 7-9 for depth)
- Build at least 2 more concurrent mini-projects or exercises
- Start preparation for Project 1.1.2 (Todo REST API with concurrency)
- Gophercises #7-8 (if you haven't started yet)
- Achieve 80%+ confidence on all concurrency topics

Write your 5 goals for Week 5 in `roadmap/weekly-goals.md`.

**Actual time spent**: ___ hours
**What I accomplished**:

---

## Week 4 Summary

### Total Time Logged
**Planned**: 13 hours
**Actual**: ___ hours
**Variance**: ___ hours (over/under)
**Running Total (Phase 1)**: ___ / 200 hours (Module 1.1 target, should be ~48h by end of Week 4)

### Deliverables Check

- [ ] TGPL Chapters 7–9 read with notes in `notes/books/tgpl-chapter-[7-9].md`
- [ ] Notes include design patterns: graceful shutdown, context usage, Mutex vs Channel
- [ ] Gophercises #5 — done + tests (including -race) + pushed
- [ ] Gophercises #6 — done + tests (including -race) + pushed
- [ ] Project 1.1.1 (Log Parser CLI) — complete, >80% coverage, pushed to GitHub
- [ ] Week 4 concurrent exercise — done + tests (including -race) + pushed
- [ ] All code passes `gofmt` and `go vet`
- [ ] All code passes `go test -race` with zero data race warnings
- [ ] All packages have >75% test coverage
- [ ] Daily commits to GitHub — at least 6 this week
- [ ] Week 4 journal filled in with concurrency learning
- [ ] Week 5 goals written in `roadmap/weekly-goals.md`
- [ ] Running hours total updated in `roadmap/TIME-TRACKING.md`
- [ ] Retrospective on Project 1.1.1 completed

### Gophercises Tracker (updated)

| # | Exercise | Status | Tests Written | Coverage | Race-Free |
|---|---|---|---|---|---|
| 1 | Echo CLI | Complete | Yes | ___ | Yes |
| 2 | Quiz Game | Complete | Yes | ___ | Yes |
| 3 | URL Shortener | Complete | Yes | ___ | Yes |
| 4 | Link Parser | Complete | Yes | ___ | Yes |
| 5 | [Likely concurrent] | Complete | Yes | ___ | Yes |
| 6 | [Likely concurrent] | Complete | Yes | ___ | Yes |
| 7 | | Not started | | | |
| ... | | | | | |
| 20 | | Not started | | | |

### Projects Completed

| Project | Status | Coverage | URL |
|---|---|---|---|
| Week 2 Mini-Project | Complete | >80% | projects/phase-1/week-2-mini-project/ |
| Week 3 Mini-Project | Complete | >80% | projects/phase-1/week-3-mini-project/ |
| Project 1.1.1: Log Parser CLI | Complete | >80% | projects/phase-1/log-parser-cli/ |
| Week 4 Concurrent Exercise | Complete | >80% | projects/phase-1/week-4-concurrent-exercise/ |

---

## Concepts You Must Be Able to Explain by End of Week 4

Answer these without looking anything up. Concurrency is hard; be honest about what you truly understand.

**TGPL Chapter 7 — Goroutines & Channels**
1. What is the difference between a thread (OS thread) and a goroutine?
2. What happens when you `go func(){ ... }()` and the function tries to access a variable from the enclosing scope?
3. What is the difference between these two?
   ```go
   ch := make(chan int)           // Unbuffered
   ch := make(chan int, 10)       // Buffered with capacity 10
   ```
4. When does sending to an unbuffered channel block? When does receiving?
5. What happens if you send to a closed channel? Receive from a closed channel?
6. How does `x, ok := <-ch` help you detect if a channel is closed?
7. What does `select` do and why is it useful for concurrent programs?
8. Can you have data races with channels? Explain your answer.

**TGPL Chapter 8 — Sync Primitives**
9. What is sync.WaitGroup and when would you use it instead of channels?
10. What does sync.Mutex protect and what happens if you forget to unlock?
11. What is sync.Once and what is it useful for?
12. Why might you use sync.RWMutex instead of sync.Mutex?

**TGPL Chapter 9 — Context & Cancellation**
13. What is context.Background() used for?
14. What is the difference between context.WithCancel() and context.WithTimeout()?
15. What does `<-ctx.Done()` return and how do you use it?
16. How do you pass a context through a chain of function calls?

**Concurrency Patterns**
17. Draw or describe the "graceful shutdown" pattern using context and goroutines.
18. What is a goroutine leak and how do you prevent one?
19. When would you use channels for communication vs sync.Mutex for protecting state?
20. How do you test concurrent code for correctness? (Hint: `go test -race`)

**Go Memory Model**
21. What guarantee does Go make about memory visibility when you send on a channel?
22. What is a data race and why does `go test -race` help you find them?
23. Can you read/write a variable from multiple goroutines without synchronization? When is it safe?

---

## Code Quality Standards — Updated for Concurrency

From `roadmap/PHASE-1-MODULES.md` — all deliverables require:

| Standard | Requirement | How to Check |
|---|---|---|
| Formatting | All files formatted | `gofmt -l .` returns nothing |
| Vet | No vet warnings | `go vet ./...` returns nothing |
| **Race-free** | **NO data races** | **`go test -race ./...` returns nothing** |
| Test coverage | >80% for projects, >75% for exercises | `go test -cover ./...` |
| Error handling | Every error is handled | Manual review |
| Docstrings | Exported functions have comments | `go doc ./...` |
| Goroutine safety | Concurrent code documents goroutine safety | Manual review |
| README | Every project has a README | Present in repo |

**NEW Standards for Week 4+**:

1. **Every concurrent function must document goroutine safety**:
   ```go
   // Process handles requests concurrently.
   // It is safe to call from multiple goroutines.
   func Process(req Request) Response {
   ```

2. **Every test of concurrent code must include a race-detector run**:
   ```bash
   go test -race ./...
   ```

3. **Every goroutine must have a clear exit condition** — no "fire and forget" goroutines

4. **Context should be passed through call stacks** — not stored or accessed globally

---

## If You Are Behind

**Didn't finish TGPL Chapters 7-9?**
→ These are foundational. Do NOT skip ahead.
→ Move unfinished chapters to start of Week 5.
→ Note in journal why and what you will change.

**Project 1.1.1 is a mess?**
→ This is expected. Your first real project will feel messy.
→ Focus on: making it work correctly, passing tests, no race conditions.
→ Code quality is secondary to correctness on first projects.
→ By Week 5, refactor and clean it up.

**Race conditions keep appearing in tests?**
→ Great! `go test -race` is doing its job.
→ Most races stem from:
  - Sending/receiving without proper coordination
  - Accessing shared state without mutex
  - Not waiting for goroutines to finish
→ Use channels when goroutines need to communicate.
→ Use mutexes when they need to access shared state.

**Gophercises #5 or #6 feel too hard?**
→ Concurrent exercises are harder than sequential ones.
→ Break the problem down: Design first, code second.
→ Test with smaller examples before building the full solution.
→ Use print statements / logging to understand goroutine flow.

**Cannot hit 13 hours?**
→ Log exactly what you got. Concurrency takes time.
→ 200-hour Module 1.1 target requires 12.5h/week. You are on pace but tight.
→ Consider if you need to extend Phase 1 or compress elsewhere.

**Confused about when to use goroutines?**
→ General rule: Use goroutines when operations are independent and can happen concurrently.
→ Bad: `go` for a single background task without coordination (goroutine leak)
→ Good: `go` to parallelize multiple independent operations with proper cleanup

---

## Next Week: Solidifying Concurrency + Project 1.1.2

Week 5 builds on concurrency foundations with more complex patterns:
- Reread Chapters 7-9 for depth
- Build 2-3 more concurrent mini-projects
- Understand common pitfalls and how to avoid them
- Prepare for Project 1.1.2 (Todo REST API)

**Week 5 Hard Mode**: Implement a worker pool and demonstrate load balancing.

By end of Week 5, you will be able to:
- Design concurrent systems that are both correct and efficient
- Identify and fix race conditions
- Choose appropriate synchronization primitives
- Build production-ready concurrent services

This is the power of Go. See you in Week 5.
