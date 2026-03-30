# Week 5: Concurrency Mastery — Advanced Patterns + Project 1.1.2 Planning + Gophercises #7–8

**Week**: Week 5 of Phase 1 — Module 1.2 (Concurrency Mastery)
**Dates**: April 20 – April 26, 2026
**Total Hours Goal**: 13 hours (Module 1.2 requires 60 hours over 4 weeks → average 15h/week, but phased)
**Location**: `Year-2026/Month-03/WEEK-5-START-HERE.md`
**Roadmap Reference**: `roadmap/PHASE-1-MODULES.md` → Module 1.2 (Concurrency Mastery)

---

## ⚠️ Read This Before You Start

Week 4 gave you the foundations of concurrency: goroutines, channels, context. Week 5 is about mastery — understanding the subtle patterns that real concurrent systems use, recognizing pitfalls, and building systems that work correctly under load.

- **Week 4**: Foundations — goroutines, channels, basic patterns
- **Week 5**: Mastery — advanced patterns, error handling, system design, testing at scale

This week is not about learning new concepts. It is about deep practice and correct implementation. You will revisit everything from Week 4 and use it in more complex scenarios.

### Module 1.2 Progress Check

From `roadmap/PHASE-1-MODULES.md` — Module 1.2 (Concurrency Mastery):

| Requirement | Target | Your Status After Week 4 |
|---|---|---|
| TGPL Chapters 7–9 mastery | Deep understanding | Foundational (need depth) |
| Concurrency exercises completed | 5+ | 2 (Week 4 base + concurrent exercise) |
| Worker pool implementation | Can design & build | Not yet |
| Error handling in concurrent code | Production-ready | Basic understanding |
| Testing concurrent systems | Comprehensive | Basic (go test -race) |
| Hours logged | 200 hours total | ~61h |

**Week 5 advances**: Advanced patterns, worker pools, error handling strategies, comprehensive concurrent testing.

---

## Week 5 Goals

**Primary Goal**: Deep mastery of concurrency patterns. Build worker pool systems, implement robust error handling in concurrent code, test concurrent systems comprehensively.

**Secondary Goals**:
- Understand and implement advanced patterns: worker pools, fan-out/fan-in, rate limiting
- Master error propagation in concurrent code
- Test concurrent systems thoroughly (not just for races, but correctness under load)
- Complete Gophercises #7 and #8
- Begin architecture design for Project 1.1.2 (Todo REST API with concurrency)
- Consolidate understanding through teaching/explaining concepts
- Commit to GitHub daily with meaningful commit messages
- Log hours daily

**Hours Target**: 13 hours (consolidating rather than learning new concepts; depth over breadth)

**Realism**: This week is about "does my code work correctly?" rather than "does my code pass go vet". You will test concurrent code under stress, with timeouts, with errors, with cancellation.

---

## Key Resources for This Week

| Resource | Use For | Link |
|---|---|---|
| TGPL Book Chapters 7–9 | Rereading for mastery | Your copy (focus on examples) |
| Go by Example — all concurrency sections | Pattern reference | https://gobyexample.com/ |
| pkg.go.dev — sync, context, time | Deep API understanding | https://pkg.go.dev/ |
| Go Memory Model | Correctness verification | https://go.dev/ref/mem |
| Uber Go Style Guide — Concurrency | Real-world best practices | https://github.com/uber-go/guide/blob/master/style.md |
| Concurrency Patterns in Go | Advanced patterns | https://go.dev/blog/pipelines |
| Testing Go Concurrency | Testing strategies | https://www.youtube.com/watch?v=8hlmJbB0ZVo |

---

## Before You Write Any Code This Week

Check that all code from Week 4 still passes quality checks:

```bash
# Race detector — most important this week
go test -race -cover ./...

# Standard checks
gofmt -l .
go vet ./...
go test -v ./...
```

**NEW this week**: Run tests multiple times to catch intermittent failures:

```bash
# Run tests 10 times — catch race conditions that happen intermittently
for i in {1..10}; do go test -race ./... || break; done
```

---

## Detailed Daily Plan

### Monday, April 20 (Target: 2.5 hours)

**Focus: Concurrency Pattern Deep Dive — Worker Pools**

Week 5 starts with the most common concurrent pattern in real systems: worker pools. You will understand this deeply and implement it.

#### Session 1: Worker Pool Pattern (90 min)

**What is a Worker Pool?**

A worker pool is a pattern where N goroutines (workers) wait for tasks on a channel, process them, and return results on another channel.

```
                    /→ Worker 1 ┐
    Tasks Channel --+-→ Worker 2 +-→ Results Channel
                    \→ Worker 3 ┘

- Multiple workers process tasks concurrently
- Tasks are distributed fairly (no single worker is overloaded)
- Workers are reused (not created/destroyed per task)
- Results come back in any order
```

**Why use worker pools?**
- Limit concurrent operations (don't create N goroutines for N tasks if N is huge)
- Reuse goroutines (cheaper than creating/destroying)
- Backpressure: tasks queue up if workers are busy
- Common in: API rate limiting, database query batching, file processing

**Read and Study**:
- [ ] Go Blog: "Pipelines" — https://go.dev/blog/pipelines (read fully, type examples)
- [ ] Read TGPL Chapter 8 sections on WaitGroup again
- [ ] Create file: `notes/books/concurrency-patterns.md`

**Implement a Worker Pool**:

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

// Task represents work to be done
type Task struct {
    ID   int
    Data string
}

// Result represents the outcome of processing a task
type Result struct {
    TaskID int
    Output string
    Error  error
}

// Worker processes tasks from the input channel and sends results to output
func worker(ctx context.Context, id int, tasks <-chan Task, results chan<- Result) {
    for {
        select {
        case task, ok := <-tasks:
            if !ok {
                // Tasks channel closed, worker exits
                return
            }
            // Process the task
            result := Result{
                TaskID: task.ID,
                Output: fmt.Sprintf("Worker %d processed: %s", id, task.Data),
            }
            // Send result (or exit if context cancelled)
            select {
            case results <- result:
            case <-ctx.Done():
                return
            }
        case <-ctx.Done():
            // Context cancelled, worker exits
            return
        }
    }
}

// WorkerPool creates N workers and processes tasks
func WorkerPool(ctx context.Context, numWorkers int, tasks []Task) []Result {
    tasksChan := make(chan Task, len(tasks))
    resultsChan := make(chan Result, len(tasks))
    
    // Start N workers
    var wg sync.WaitGroup
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            worker(ctx, workerID, tasksChan, resultsChan)
        }(i)
    }
    
    // Send tasks (in a goroutine so we don't block)
    go func() {
        for _, task := range tasks {
            select {
            case tasksChan <- task:
            case <-ctx.Done():
                close(tasksChan)
                return
            }
        }
        close(tasksChan) // Signal that no more tasks
    }()
    
    // Wait for workers to finish, then close results
    go func() {
        wg.Wait()
        close(resultsChan)
    }()
    
    // Collect results
    var results []Result
    for result := range resultsChan {
        results = append(results, result)
    }
    return results
}
```

Study this code. Understand every line. Ask yourself:
- Where could a goroutine leak occur if this code were wrong?
- What happens if a worker panics?
- How does context cancellation work through this system?
- Why is the tasks channel closed?

**Write this code in**: `projects/phase-1/concurrency-patterns/worker-pool.go`

#### Session 2: Gophercises #7 Start (60 min)

URL: https://gophercises.com/ → Exercise 7

Most Gophercises #7 problems involve concurrent processing. Read the problem fully.

- [ ] Read the full problem
- [ ] Plan before coding:

```go
// Plan:
// 1. [What work will be parallel?]
// 2. [How many workers do I need?]
// 3. [What messages flow between goroutines?]
// 4. [What errors can occur? How do I handle them?]
```

- [ ] Start implementing
- [ ] Write 2 tests (happy path, error handling)

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Tuesday, April 21 (Target: 2 hours)

**Focus: Error Handling in Concurrent Code + Finish Gophercises #7**

#### Session 1: Finish Gophercises #7 + Tests (60 min)

- [ ] Complete the exercise
- [ ] Write tests that cover:
  - Happy path (all work completes)
  - Partial failure (some tasks fail)
  - Timeout scenario (work times out)
  - Cancellation (context is cancelled)
- [ ] Run `go test -race -v -cover ./...` — aim for >75% coverage
- [ ] Push finished Gophercises #7

**Testing Error Scenarios**:

```go
func TestWithPartialFailure(t *testing.T) {
    // Create tasks where some will fail
    tasks := []Task{
        {ID: 1, Data: "good"},
        {ID: 2, Data: "bad"},    // Will error
        {ID: 3, Data: "good"},
    }
    
    results := ProcessTasks(context.Background(), tasks)
    
    // Count successes and failures
    var successes, failures int
    for _, r := range results {
        if r.Error != nil {
            failures++
        } else {
            successes++
        }
    }
    
    if failures != 1 || successes != 2 {
        t.Errorf("Expected 2 successes, 1 failure; got %d, %d", successes, failures)
    }
}
```

#### Session 2: Error Propagation in Concurrent Systems (60 min)

Error handling in concurrent code is the most neglected aspect. Learn it deeply.

**Pattern 1: Error Channel**

```go
// Result includes both value and error
type Result struct {
    Value interface{}
    Error error
}

// Worker sends results with errors
func worker(tasksChan <-chan Task, resultsChan chan<- Result) {
    for task := range tasksChan {
        value, err := process(task)
        resultsChan <- Result{
            Value: value,
            Error: err,
        }
    }
}

// Caller checks errors
for result := range results {
    if result.Error != nil {
        log.Printf("Task failed: %v", result.Error)
        // Decide: continue, stop, retry?
    } else {
        use(result.Value)
    }
}
```

**Pattern 2: Early Exit on Error**

```go
// If ANY worker fails, stop everything
func ProcessWithEarlyExit(tasks []Task) ([]Result, error) {
    errChan := make(chan error, 1)  // Buffered: report error once
    // ... setup workers ...
    
    go func() {
        for result := range results {
            if result.Error != nil {
                select {
                case errChan <- result.Error:
                default: // Already reported error
                }
                // Caller will cancel context
                cancel()
            }
        }
    }()
    
    select {
    case err := <-errChan:
        return nil, err
    case <-time.After(timeout):
        return nil, fmt.Errorf("timeout")
    }
}
```

**Pattern 3: Error Aggregation**

```go
// Collect ALL errors, return them together
type AggregateError struct {
    Errors []error
}

func (ae AggregateError) Error() string {
    // Format all errors
}

// Collect errors from all workers
func ProcessWithAggregation(tasks []Task) ([]Result, *AggregateError) {
    results := /* ... run workers ... */
    
    var aggErr *AggregateError
    for _, r := range results {
        if r.Error != nil {
            if aggErr == nil {
                aggErr = &AggregateError{}
            }
            aggErr.Errors = append(aggErr.Errors, r.Error)
        }
    }
    
    if aggErr != nil {
        return results, aggErr
    }
    return results, nil
}
```

**Write this in**: `notes/books/concurrency-error-handling.md` with examples of all 3 patterns.

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Wednesday, April 22 (Target: 2.5 hours)

**Focus: Advanced Patterns — Fan-Out/Fan-In + Rate Limiting**

#### Session 1: Fan-Out/Fan-In Pattern (90 min)

Fan-out/Fan-in is how you distribute work and aggregate results. Essential for building scalable systems.

**Fan-Out**: One goroutine reads input and delegates to many workers.
**Fan-In**: Many workers produce results that are gathered into one channel.

```
Input → Splitter ──→ Worker 1 ────┐
                ├→ Worker 2 ├→ Merger → Output
                └→ Worker N ────┘
```

Study and implement:

```go
// Fan-out: split pipeline into multiple workers
func fanOut(in <-chan int, numWorkers int) []<-chan int {
    channels := make([]<-chan int, numWorkers)
    for i := 0; i < numWorkers; i++ {
        ch := make(chan int)
        channels[i] = ch
        
        go func(out chan<- int) {
            for num := range in {
                out <- num
            }
            close(out)
        }(ch)
    }
    return channels
}

// Worker processes values from a channel
func worker(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for num := range in {
            out <- process(num)
        }
        close(out)
    }()
    return out
}

// Fan-in: merge multiple channels into one
func fanIn(channels ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    
    for _, ch := range channels {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for num := range c {
                out <- num
            }
        }(ch)
    }
    
    go func() {
        wg.Wait()
        close(out)
    }()
    
    return out
}

// Usage:
inChan := generateNumbers() // Produces 1,2,3,...
workers := fanOut(inChan, 3) // 3 workers
results := fanIn(workers...)  // Merge back
for result := range results {
    fmt.Println(result)
}
```

Write this in: `projects/phase-1/concurrency-patterns/fan-in-fan-out.go`

Carefully trace through the code:
- When does each goroutine exit?
- What happens if a worker panics?
- Why must we use WaitGroup to close the output channel?

#### Session 2: Gophercises #8 Start (60 min)

URL: https://gophercises.com/ → Exercise 8

- [ ] Read the full problem
- [ ] Identify: is this a worker pool? fan-out/fan-in? rate limiting?
- [ ] Plan:

```go
// Plan:
// [Similar structure to previous exercises]
```

- [ ] Start implementing
- [ ] Write 2 tests

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Thursday, April 23 (Target: 2 hours)

**Focus: Rate Limiting + Finish Gophercises #8**

#### Session 1: Finish Gophercises #8 + Tests (60 min)

- [ ] Complete the exercise
- [ ] Write comprehensive tests covering edge cases
- [ ] Run `go test -race -v -cover ./...` — aim for >75% coverage, zero race conditions
- [ ] Push finished Gophercises #8

#### Session 2: Rate Limiting Pattern (60 min)

Rate limiting controls how fast operations happen. Common in: API calls, database queries, resource allocation.

**Pattern 1: Token Bucket (using time.Ticker)**

```go
// RateLimiter limits operations to N per second
type RateLimiter struct {
    tokens chan struct{}
}

func NewRateLimiter(rps float64) *RateLimiter {
    rl := &RateLimiter{
        tokens: make(chan struct{}, int(rps)),
    }
    
    // Generate tokens at the specified rate
    go func() {
        ticker := time.NewTicker(time.Second / time.Duration(rps))
        defer ticker.Stop()
        
        for range ticker.C {
            select {
            case rl.tokens <- struct{}{}:
            default:
                // Channel full, token dropped (maxed out)
            }
        }
    }()
    
    return rl
}

// Wait blocks until a token is available
func (rl *RateLimiter) Wait(ctx context.Context) error {
    select {
    case <-rl.tokens:
        return nil
    case <-ctx.Done():
        return ctx.Err()
    }
}

// Usage:
limiter := NewRateLimiter(10) // 10 operations per second

for _, url := range urls {
    if err := limiter.Wait(ctx); err != nil {
        break
    }
    go fetchURL(url)
}
```

**Pattern 2: Semaphore (fixed number of concurrent operations)**

```go
// Semaphore limits N concurrent operations
type Semaphore struct {
    sem chan struct{}
}

func NewSemaphore(max int) *Semaphore {
    return &Semaphore{
        sem: make(chan struct{}, max),
    }
}

func (s *Semaphore) Acquire(ctx context.Context) error {
    select {
    case s.sem <- struct{}{}:
        return nil
    case <-ctx.Done():
        return ctx.Err()
    }
}

func (s *Semaphore) Release() {
    <-s.sem
}

// Usage:
sem := NewSemaphore(5) // Max 5 concurrent

for _, item := range items {
    if err := sem.Acquire(ctx); err != nil {
        break
    }
    go func(i Item) {
        defer sem.Release()
        process(i)
    }(item)
}
```

Write examples in: `projects/phase-1/concurrency-patterns/rate-limiting.go`

Test both patterns:
- Verify rate limiting works (use time.Now() to verify timing)
- Test context cancellation stops waiting
- Test semaphore limits concurrent operations

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Friday, April 24 (Target: 1.5 hours)

**Focus: Testing Concurrent Code + Project 1.1.2 Planning**

#### Session 1: Stress Testing Concurrent Code (45 min)

Before you call concurrent code "correct," stress test it.

**Create**: `projects/phase-1/stress-test/main.go`

```go
package main

import (
    "context"
    "fmt"
    "math/rand"
    "sync"
    "testing"
    "time"
)

// StressTest runs a concurrent operation under load
func StressTest(t *testing.T, numGoroutines int, operationsPerGoroutine int) {
    var wg sync.WaitGroup
    errors := make(chan error, numGoroutines*operationsPerGoroutine)
    
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    // Launch many concurrent operations
    for i := 0; i < numGoroutines; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 0; j < operationsPerGoroutine; j++ {
                // Add some chaos: random delays, cancellations
                time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
                
                if rand.Float32() < 0.05 { // 5% cancel mid-operation
                    cancel()
                }
                
                result, err := YourConcurrentOperation(ctx)
                if err != nil && err != context.Canceled {
                    select {
                    case errors <- fmt.Errorf("goroutine %d op %d: %w", id, j, err):
                    default:
                    }
                }
            }
        }(i)
    }
    
    wg.Wait()
    close(errors)
    
    // Check for errors
    var errCount int
    for err := range errors {
        errCount++
        t.Logf("Error during stress test: %v", err)
    }
    
    if errCount > 0 {
        t.Fatalf("Stress test failed with %d errors", errCount)
    }
}

// Test it
func TestStress(t *testing.T) {
    // 100 goroutines, 100 ops each = 10,000 concurrent operations
    StressTest(t, 100, 100)
}
```

**Run it**:
```bash
go test -race -v -timeout 30s ./...
```

Run it multiple times. Concurrent bugs are often intermittent.

#### Session 2: Project 1.1.2 Architecture Planning (45 min)

Next week you build Project 1.1.2: Todo REST API. This week, design it.

**Create**: `projects/phase-1/todo-rest-api/PLAN.md`

Structure:

```markdown
# Todo REST API — Architecture Plan

## What It Does (from PHASE-1-MODULES.md)
- REST CRUD endpoints for todos
- SQLite persistence
- Proper HTTP status codes
- Input validation
- API key authentication
- Unit + integration tests

## REST Endpoints (design these)
- GET /todos — list all todos
- GET /todos/:id — get single todo
- POST /todos — create todo
- PUT /todos/:id — update todo
- DELETE /todos/:id — delete todo

## Data Model
```go
type Todo struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Done      bool      `json:"done"`
    CreatedAt time.Time `json:"created_at"`
}
```

## Concurrency Strategy
[ ] Question 1: How many concurrent requests will the API handle?
[ ] Question 2: Will database access be the bottleneck?
[ ] Question 3: Do I need a connection pool?
[ ] Question 4: Should handler goroutines write to shared state?

Answer: We will use:
- [ ] Handler goroutines (one per request, managed by net/http)
- [ ] Mutex to protect shared state OR connection pool
- [ ] Context for request cancellation

## Error Handling
- Invalid JSON: 400 Bad Request
- Not found: 404 Not Found
- Conflict (duplicate): 409 Conflict
- Server error: 500 Internal Server Error

## Testing Strategy
[ ] Unit tests for business logic
[ ] Integration tests with real HTTP handlers
[ ] Concurrency tests: multiple requests simultaneously
[ ] Error scenario tests

## Code Organization
```
todo-rest-api/
├── main.go              (server setup, routing)
├── handler.go           (HTTP handlers)
├── store.go             (SQLite access, thread-safe)
├── model.go             (Todo struct, validation)
├── auth.go              (API key checking)
├── main_test.go
├── handler_test.go
├── store_test.go
└── README.md
```

## Success Criteria
- All endpoints work correctly
- Concurrent requests handled correctly (>80% coverage, zero race conditions)
- Input validation prevents invalid data
- Proper error responses
- 85%+ test coverage

Fill in all sections. Do NOT code yet. Design first, code later.

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Saturday, April 25 (Target: 2.5 hours)

**Focus: Complex Concurrent Pattern Exercise + Project Prep**

#### Session 1: Advanced Concurrency Exercise — Choose ONE (90 min)

This is a challenging exercise. You are implementing one of the patterns you learned this week.


**Option A: Task Pipeline with Error Recovery (Recommended)**

Build a multi-stage pipeline that processes tasks, with error recovery.

Requirements:
- Stage 1 (reader): Read tasks from file, send to stage 2
- Stage 2 (validator): Validate tasks, send valid to stage 3, retry invalid up to 3 times
- Stage 3 (processor): Process validated tasks, return results
- Stage 4 (writer): Collect results and write to output file
- If any stage fails catastrophically, abort entire pipeline
- Support `-timeout` and context cancellation
- Write 5+ tests covering: happy path, stage failures, timeout, partial recovery

Skills: multi-stage pipelines, error propagation, retry logic, context cancellation

---

**Option B: Concurrent Download Manager**

Build a tool that downloads multiple files concurrently with resource limits.

Requirements:
- Accept list of URLs to download
- Use 5-10 concurrent workers (configurable with `-workers`)
- Limit bandwidth: max N MB/s total (configurable with `-bandwidth`)
- Retry failed downloads up to 3 times with exponential backoff
- Show progress: downloaded X of Y files, Z MB/s
- Support cancellation (Ctrl+C) with graceful shutdown
- Write 5+ tests: success, network errors, timeouts, cancellation, bandwidth limiting

Skills: rate limiting, worker pools, retry logic, progress tracking, signal handling

---

**Option C: Concurrent Cache with Expiration**

Build a thread-safe cache that automatically expires entries.

Requirements:
- Store key-value pairs with TTL (time-to-live)
- Support concurrent reads and writes
- Automatically expire entries after TTL
- Evict expired entries when cache is full
- Configurable: max size, eviction strategy (LRU or FIFO)
- Support context cancellation for all operations
- Write 6+ tests: reads/writes, expiration, eviction, concurrent stress, cancellation

Skills: sync.RWMutex for read-writer pattern, WaitGroup, goroutine lifecycle, testing concurrent data structures

---

**For whichever option you choose**:

- [ ] Write a design document first (5 min)
- [ ] Implement (75 min) with tests as you build
- [ ] Run `go test -race -v -cover ./...` multiple times
- [ ] Target >85% coverage, zero race conditions
- [ ] Add README with usage examples
- [ ] Push to GitHub: `projects/phase-1/week-5-advanced-pattern/`
- [ ] Commit message: `feat: week-5 advanced pattern — [option name]`

#### Session 2: Project 1.1.2 Todo API — Basic Structure (60 min)

Start building the API (full completion next week):

- [ ] Create directory: `projects/phase-1/todo-rest-api/`
- [ ] Copy your PLAN.md there
- [ ] Create basic file structure (main.go, handler.go, store.go, etc.)
- [ ] Implement basic routing with net/http
- [ ] Implement Todo struct and JSON marshalling
- [ ] Add TODO comments for unfinished sections

**Don't finish it this week** — just get structure and basic routing in place. Full implementation Week 6.

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Sunday, April 26 (Target: 1 hour)

**Focus: Weekly Review + Week 6 Planning**

#### Part 1: Concurrency Mastery Reflection (20 min)

Answer these deeply:

- [ ] Design a system with 4+ goroutines. Draw it. Explain where communication happens.
- [ ] List 5 patterns you learned this week: worker pool, fan-out/fan-in, rate limiting, error handling, ... (what's the 5th?)
- [ ] When you design concurrent code, what's your first question? (Hint: How many goroutines? When do they exit?)
- [ ] What race condition was hardest to understand this week? How did `go test -race` help you find it?

Write answers in your journal under "Concurrency Mastery Reflection".

#### Part 2: Journal Entry (25 min)

Create your Week 5 journal:
```bash
cp journal/JOURNAL-TEMPLATE.md journal/2026-04-week-17.md
```

**Technical Concepts section** — write 6+ specific examples from concurrency:
- "Worker pools: I finally understand why you separate task distribution from processing. Reduces context switching, improves throughput."
- "Error handling in concurrent code is subtle. I had to think about: Does one error cancel everything? Do I collect all errors? The answer depends on the use case."
- "Fan-out/fan-in feels like magic the first time, but it's just channels. Send to multiple channels, collect from multiple channels."
- "Rate limiting with token buckets or semaphores: I understood the concept, but the implementation details matter for correctness."

**Struggles section**:
- "I kept trying to optimize before the code was correct. `go test -race` forced me to focus on correctness first."
- "Testing concurrent code under stress is hard. I need better strategies for generating chaos."

**Concepts to Revisit**:
- List 3 topics you feel 3/5 confident on (not shaky, but not expert)
- Plan to return to them in Week 6

#### Part 3: Self-Assessment (10 min)

Rate your confidence on concurrency topics (1 = barely understand, 5 = can teach someone):

| Topic | Confidence |
|---|---|
| Goroutine lifecycle and cleanup | ___ |
| Channel blocking semantics | ___ |
| Select statement | ___ |
| WaitGroup synchronization | ___ |
| Mutex and protected state | ___ |
| Context and cancellation | ___ |
| Worker pools | ___ |
| Fan-out/fan-in patterns | ___ |
| Error propagation in concurrent systems | ___ |
| Rate limiting | ___ |
| Testing for race conditions | ___ |
| Testing concurrent systems under load | ___ |

Target: Most should be 4+. Any 1-2 → practice more.

#### Part 4: Week 6 Planning (5 min)

Week 6 target:
- Complete Project 1.1.2 (Todo REST API) — full implementation
- Complete Gophercises #9-10
- Optional: Deep dive on specific concurrency topic (pick your weakest 1-2)
- Prepare for Module 1.3 (Testing & Best Practices) which starts Week 7

Write 5 goals in `roadmap/weekly-goals.md`.

**Actual time spent**: ___ hours
**What I accomplished**:

---

## Week 5 Summary

### Total Time Logged
**Planned**: 13 hours
**Actual**: ___ hours
**Variance**: ___ hours (over/under)
**Running Total (Phase 1)**: ___ / 200 hours (should be ~74h by end of Week 5)

### Deliverables Check

- [ ] TGPL Chapters 7–9 reread for mastery (notes in `notes/books/concurrency-patterns.md`)
- [ ] Error handling patterns documented and implemented
- [ ] Worker pool pattern implemented and tested
- [ ] Fan-out/fan-in pattern implemented and tested
- [ ] Rate limiting (token bucket + semaphore) implemented
- [ ] Gophercises #7 — done + tests + race-detector pass + pushed
- [ ] Gophercises #8 — done + tests + race-detector pass + pushed
- [ ] Week 5 advanced pattern exercise — done + tests + race-detector pass + pushed
- [ ] Project 1.1.2 (Todo REST API) — design plan + basic structure (code started)
- [ ] All code passes `gofmt`, `go vet`, and `go test -race` with zero warnings
- [ ] All packages have >75% test coverage
- [ ] Stress tests run successfully with no race conditions or deadlocks
- [ ] Daily commits to GitHub — at least 6 this week
- [ ] Week 5 journal filled in with concurrency reflection
- [ ] Week 6 goals written in `roadmap/weekly-goals.md`
- [ ] Running hours total updated in `roadmap/TIME-TRACKING.md`

### Gophercises Tracker (updated)

| # | Exercise | Status | Tests Written | Coverage | Race-Free |
|---|---|---|---|---|---|
| 1–6 | [Previous] | Complete | Yes | ___ | Yes |
| 7 | [Concurrent] | Complete | Yes | ___ | Yes |
| 8 | [Concurrent] | Complete | Yes | ___ | Yes |
| 9–10 | | Not started | | | |
| ... | | | | | |
| 20 | | Not started | | | |

---

## Concepts You Must Be Able to Explain by End of Week 5

No looking anything up. These are mastery-level questions.

**Concurrency Patterns**
1. Draw a worker pool system. Explain: How many goroutines? When do they start? When do they exit?
2. Explain fan-out/fan-in. When would you use it instead of a worker pool?
3. Design a rate limiter that allows 100 operations per second. Show the code.
4. How do you handle errors from multiple concurrent operations?
5. Draw a graceful shutdown pattern: services running, signal received, cleanup happening, done.

**Advanced Understanding**
6. If you have 1000 tasks and 4 workers, explain what happens when you send a task to a worker's channel.
7. Why does closing a tasks channel matter? What happens if you don't close it?
8. Design a concurrent system that scales from 1 concurrent operation to 10,000 without creating 10,000 goroutines.
9. If one worker panics, what happens to the other workers and the main goroutine?
10. How do you prevent goroutine leaks in error scenarios?

**Testing & Verification**
11. You wrote concurrent code that passes `go test -race` but fails intermittently in production. What went wrong?
12. Design a test that verifies: operation completes in < 100ms, handles timeout gracefully, propagates errors correctly.
13. Your code has a race condition that `go test -race` doesn't catch (because it's rare). How do you find and fix it?

**Design Decisions**
14. Given a task: "Process 1 million log entries, categorize each one". Design 3 different concurrent approaches and compare them.
15. Your API needs to limit requests to 1000/second. Rate limit at the handler level or worker level? Why?

---

## Code Quality Standards — Maintained + Enhanced

| Standard | Requirement | How to Check |
|---|---|---|
| Formatting | All files formatted | `gofmt -l .` returns nothing |
| Vet | No vet warnings | `go vet ./...` returns nothing |
| **Race-free** | **ZERO data races** | **`go test -race ./...` multiple times** |
| Test coverage | >85% for projects, >75% for exercises | `go test -cover ./...` |
| Error handling | Every error handled | Manual review |
| Docstrings | Functions documented | `go doc ./...` |
| Goroutine safety | Clear documentation | "Safe to call from multiple goroutines" |
| Stress tested | Handles 10x normal load | Manual verification |
| Graceful shutdown | Clean exit on cancellation | Manual verification |

**NEW**: Every concurrent package should include a file: `doc.go` with package-level documentation explaining the concurrency model:

```go
// Package worker provides a worker pool for concurrent task processing.
//
// Concurrency Model:
// The WorkerPool uses N goroutines (workers) that process tasks from a channel.
// It is safe to send tasks from multiple goroutines. The pool handles synchronization
// internally. Workers exit gracefully when the task channel is closed.
//
// Error Handling:
// If a worker returns an error, it is included in the result. The pool continues
// processing other tasks. One failed task does not stop the pool.
package worker
```

---

## If You Are Behind

**Didn't finish advanced patterns?**
→ Don't panic. Patterns take time to master.
→ Move unfinished patterns to Week 6.
→ Focus on understanding one pattern deeply rather than knowing all patterns shallowly.

**Race conditions still appearing?**
→ Good! That means you're testing.
→ `go test -race` is your closest friend now.
→ Run tests 10+ times: `for i in {1..10}; do go test -race ./... || break; done`
→ If a race appears intermittently, you're debugging a real bug.

**Gophercises #7 or #8 very difficult?**
→ These are likely the hardest exercises so far.
→ Reread the problem 3 times before coding.
→ Test with a sequential version first, then add concurrency.
→ Use print statements to trace goroutine execution.

**Too many concurrent goroutines, program slower?**
→ You've hit a resource limit or context switching overhead.
→ Experiment: 4 workers vs 8 vs 16. Measure timing.
→ The optimal number depends on: CPU cores, task I/O vs CPU, system load.
→ This is a Week 6 optimization question.

**Can't hit 13 hours?**
→ This week is consolidation, not new learning.
→ 74h running total means you're on pace for the 200h goal.
→ If you're behind, focus on depth of understanding over hours.

---

## Next Week: Module 1.3 — Testing & Best Practices

Week 6 brings Module 1.3, but first you finish Project 1.1.2 (Todo REST API) and potentially start Unit + Integration testing deep dive.

**Week 6 teaser**:
- Test organization and patterns
- Integration testing with real databases
- Benchmarking and profiling
- Code coverage analysis
- Documentation practices

You're now deep into Go. Concurrency is becoming natural. Keep building. See you in Week 6.
