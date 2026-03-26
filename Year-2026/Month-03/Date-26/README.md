# Day 4 — 26 March 2026

**Phase**: Phase 1 / Module 1.1 — Go Fundamentals  
**Focus**: Tour of Go — Concurrency  

---

## Context: Why Today Matters

Concurrency is the feature Go is most famous for. It is also the feature
most misunderstood by beginners. People read about goroutines, think they
understand them, and then write code that deadlocks, races, or silently
does nothing.

Today is not about writing clever concurrent code. Today is about
understanding the primitives — goroutines, channels, select — deeply
enough that you know what is actually happening when you use them.

Go's concurrency model is built on one philosophy:

> "Do not communicate by sharing memory; instead, share memory by communicating."

Keep that sentence in mind all day. It explains every design decision
in this section.

---

## Today's Single Goal

Complete the **Concurrency** section of the Tour of Go.

> https://go.dev/tour/concurrency/

This is the final major section of the Tour. After today, you will have
completed the entire Tour of Go.

---

## What You Will Cover

The Concurrency section covers these topics, in order:

1. Goroutines — what they are, how to launch them
2. Channels — typed conduits for communication between goroutines
3. Buffered channels — channels with capacity
4. Range and close — iterating over a channel, knowing when it is done
5. Select — waiting on multiple channel operations
6. Default select — non-blocking channel operations
7. sync.Mutex — protecting shared state when channels are not the right tool
8. Exercise: Equivalent Binary Trees (uses goroutines + channels + recursion)
9. Exercise: Web Crawler (concurrent crawling with goroutines)

---

## The Concepts That Trip People Up

Read these before you start. They will save you real confusion.

### Goroutines are not threads
A goroutine is a lightweight function executing concurrently with other
goroutines in the same address space. The Go runtime manages goroutines,
multiplexing them onto OS threads. You can have thousands of goroutines
where you could only have hundreds of threads. But "concurrent" does not
mean "parallel" — that depends on GOMAXPROCS.

```go
go f(x, y, z)  // launches f in a new goroutine
```

The arguments x, y, z are evaluated in the current goroutine.
The execution of f happens in the new goroutine. This distinction
matters when you are passing loop variables — a classic bug.

### Channels block by default
An unbuffered channel send blocks until a receiver is ready.
An unbuffered channel receive blocks until a sender is ready.
This is not a bug — it is the mechanism. It is how goroutines synchronise.

```go
ch := make(chan int)      // unbuffered — both sides block until ready
ch := make(chan int, 10)  // buffered  — send does not block until buffer is full
```

### The classic goroutine loop bug
This will bite you if you do not know it:

```go
// WRONG — all goroutines capture the same variable v
for _, v := range values {
    go func() {
        fmt.Println(v) // v is shared — race condition
    }()
}

// CORRECT — pass v as an argument, creating a copy per goroutine
for _, v := range values {
    go func(val string) {
        fmt.Println(val) // val is a copy — safe
    }(v)
}
```

### Select chooses randomly when multiple cases are ready
If more than one channel is ready at the same time in a select,
Go picks one at random. This is intentional. Do not write code that
depends on a specific order.

### Mutex is for shared state, channels are for communication
Use a sync.Mutex when multiple goroutines need to read/write shared
data and channels would be awkward. Use channels when goroutines need
to pass data or signals to each other. Both are valid — the right tool
depends on the situation.

---

## How to Work Through It

### The Rule
**Do not copy-paste.** Type every example. When a program deadlocks
or panics, do not just fix it and move on — understand *why* it happened.
A deadlock is information. Read the error message carefully.

### For each topic
1. Read the explanation carefully
2. Type the code in the editor
3. Run it — if it behaves unexpectedly, sit with it for 5 minutes
4. Ask yourself: *"What would happen if I changed X?"* Then try it.
5. Write down what surprised you in the Notes section below

### The exercises are harder today
The Binary Trees and Web Crawler exercises are significantly harder than
previous Tour exercises. Spend up to 20 minutes on each. If you are still
stuck, note your attempt and move on — do not lose the whole session.
You will return to these exact patterns when you build the Web Crawler
project later in Module 1.1.

---

## Code to Write Today

After finishing the Tour section, write this program **from scratch**.
No Tour reference. No AI. From memory.

**File**: `projects/phase-1/00-day4-concurrency/concurrency_practice.go`

```
Write a Go program that does the following:

1. Write a function called generate(nums ...int) <-chan int
   - It launches a goroutine that sends each number into a channel
   - It closes the channel when done
   - It returns the receive-only channel

2. Write a function called square(in <-chan int) <-chan int
   - It launches a goroutine that reads from in, squares each number,
     and sends the result into a new channel
   - It closes the output channel when the input channel is closed
   - It returns the receive-only output channel

3. In main():
   - Use generate() to create a channel of numbers 1 through 5
   - Pass that channel into square()
   - Range over the result and print each squared value
   - Expected output: 1, 4, 9, 16, 25 (one per line)

4. Chain a second square() call so numbers go through squaring twice:
   generate → square → square → print
   Expected output: 1, 16, 81, 256, 625

Allowed imports: "fmt"
```

This tests: goroutines, channels, range over channels, closing channels,
and the pipeline pattern — the most important concurrency pattern in Go.

---

## Stretch Goal (Only If Time Allows)

Add a `done` channel to your pipeline so that the main goroutine can
signal cancellation to all stages. When `done` is closed, all goroutines
should exit cleanly without leaking.

This is the foundation of how `context.Context` works under the hood.
Do not do this at the cost of the main exercise. Main exercise first.

---

## Notes (Fill These In As You Go)

### Things that made sense immediately
- A goroutine is a lightweight thread managed by the Go runtime.
- Channels are a typed conduit through which you can send and receive values with the channel operator, <- (The data flows in the direction of the arrow.)


### Things that confused me
- Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
- The select statement lets a goroutine wait on multiple communication operations

### What happened when a deadlock occurred (describe the error)
- The program stopped running and showed an error like: fatal error: all goroutines are asleep - deadlock!
    This happens when goroutines are waiting on each other (e.g., trying to receive from a channel with no sender), so nothing can proceed.

### Goroutines vs threads — my understanding in my own words
Goroutines are much lighter than OS threads and are managed by the Go runtime instead of the operating system. They use less memory, start faster, and can scale to thousands or millions. Threads are heavier, managed by the OS, and have more overhead. Go maps many goroutines onto a smaller number of threads efficiently.


### The pipeline pattern — explain it in one paragraph
The pipeline pattern in Go is a way of structuring concurrent programs where data flows through multiple stages connected by channels. Each stage runs in its own goroutine, processes incoming data, and sends the result to the next stage. This makes programs modular, easy to reason about, and efficient, since different stages can run in parallel.


### Questions I want to research later
- How does Go schedule goroutines internally
- When should I use buffered vs unbuffered channels?

### One thing I want to remember from today
- Don’t close a channel from the receiver side—only the sender should close it.

---

## Time Log

| Session | Start | End | Duration | What I worked on |
|---------|-------|-----|----------|-----------------|
| 1 | | | | Tour — Goroutines to Buffered channels |
| 2 | | | | Tour — Range/close, Select, Mutex, Exercises |
| 3 | | | | concurrency_practice.go from scratch |

**Total today**: ___ hours  
**Running total**: ___ hours / 1500–2000 goal

---

## Concepts Check

Before you close your laptop, answer these without looking anything up.
Write your answers in your notes file, not here.

1. What is a goroutine? How is it different from an OS thread?

2. What happens when you send to an unbuffered channel and no goroutine
   is receiving? What does the runtime do?

3. What is the difference between a buffered and unbuffered channel?
   When would you choose one over the other?

4. What does `close(ch)` signal? What does a receiver see after a
   channel is closed and drained?

5. What does `select` do when multiple cases are ready at the same time?

6. When would you reach for `sync.Mutex` instead of a channel?

If you cannot answer any of these, go back and re-read that topic
before calling Day 4 done.

---

## End of Tour of Go

If you have completed Days 2, 3, and 4 fully, you have finished the
entire Tour of Go — Basics, Methods and Interfaces, and Concurrency.

Take a moment to recognise that. Most people who bookmark the Tour
never finish it.

What you have now is a starting point, not mastery. You know the
vocabulary of Go. The next stage — starting Week 2 — is reading TGPL
and building real projects where that vocabulary becomes fluency.

---

## End of Day Checklist

- [ ] Completed all topics in the Concurrency section
- [ ] Typed every example (did not paste)
- [ ] Attempted both exercises (Binary Trees + Web Crawler)
- [ ] Wrote `concurrency_practice.go` from memory
- [ ] Program runs and prints correct output for both pipeline stages
- [ ] Answered the 6 concepts check questions in your notes
- [ ] Filled in the Notes section above
- [ ] Committed code to GitHub with a clear message
- [ ] Logged time honestly

---

## Git Commit Message for Today

```
day-4: tour of go concurrency section complete

- worked through goroutines, channels, select, mutex
- implemented generator and pipeline patterns from scratch
- completed Tour of Go (all 3 sections: basics, methods, concurrency)
- noted confusions around [fill in what confused you]
```

---

## Day 5 Preview

Tomorrow is your **first weekly review and journal entry**.

You are at the end of Week 1. Before moving to Week 2 (TGPL + Gophercises),
you need to consolidate what you have learned, assess honestly where you
stand, and plan the week ahead.

Day 5 is not a rest day — it is a reflection day. Open your
`journal/JOURNAL-TEMPLATE.md`, copy it to `journal/2026-03-week-12.md`,
and fill in your Week 1 summary. Then review `Year-2026/Month-03/WEEK-2-START-HERE.md`
to understand what is coming next week.

---

**Commitment check**: Did I write the code myself today? Yes / No  
**Honest hours logged**: ___