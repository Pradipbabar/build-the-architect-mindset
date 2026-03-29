# Week 2: Continue Building — TGPL + Gophercises + First Real Project

**Week**: Week 2 of Phase 1 - Module 1.1 (Go Fundamentals)
**Dates**: March 28 – April 3, 2026
**Total Hours Goal**: 10 hours
**Location**: `Year-2026/Month-03/WEEK-2-START-HERE.md`

---

## Overview: What Week 2 Accomplishes

By the end of Week 2, you will have:
- Started reading The Go Programming Language (TGPL) — Chapters 1–2
- Completed Gophercises #1 and #2
- Deepened your Go syntax and control-flow understanding
- Started writing tests for your code (your first taste of TDD)
- Committed meaningful code to GitHub every day
- Logged 10 hours of verified, focused learning

**This week is about shifting from "I ran some examples" to "I can write Go independently."**

---

## What Phase 1 Requires of You (Context)

Before diving into tasks, understand what Module 1.1 demands overall:

From `roadmap/PHASE-1-MODULES.md`, Module 1.1 success criteria requires ALL of:

- [ ] Completed TGPL book with all exercises
- [ ] Completed 20+ Gophercises
- [ ] 3 projects on GitHub with >80% test coverage
- [ ] Can explain goroutines, channels, and `select`
- [ ] Can write idiomatic Go code
- [ ] Code review approved by senior Go dev
- [ ] **200 hours logged and documented**
- [ ] At least 1 project got external code review

Week 2 contributes directly to the first three. Every hour this week counts toward your 200-hour Module 1.1 target.

---

## Week 2 Goals

**Primary Goal**: Read TGPL Chapters 1–2 AND complete Gophercises #1 and #2

**Secondary Goals**:
- Start your weekly journal entry (copy the template, fill in Days 1–5)
- Write at least one test for every function you build
- Commit to GitHub every single day (even small commits count)
- Log hours daily — do not batch on Sunday

---

## Detailed Daily Plan

### Monday, March 30 (Target: 2 hours)

**Focus: TGPL Chapter 1 — A Tutorial**

- [ ] Read TGPL Chapter 1 cover to cover (60 min)
  - Do NOT skim. Type out each example yourself. Do not copy-paste.
  - Run every example: `go run filename.go`
  - Take notes in `notes/books/tgpl-chapter-1.md`
- [ ] Start Gophercises #1: `echo` clone (45 min)
  - URL: https://gophercises.com/
  - Build a CLI tool that echoes input back
  - Write it yourself — no looking at solutions first
- [ ] Commit your code to GitHub (15 min)
  - Commit message: `feat: gophercises-01 echo cli`
  - Push to `principal-engineer-portfolio/phase-1/`

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Tuesday, March 31 (Target: 1.5 hours)

**Focus: Finish Gophercises #1 + Start #2**

- [ ] Finish Gophercises #1 if not done — add tests (30 min)
  - Write at least 2 unit tests for your echo CLI
  - Run: `go test ./...`
- [ ] Start Gophercises #2: `quiz` game (60 min)
  - Read the problem statement fully before writing any code
  - Plan your approach (write a comment outline first)
  - Build and test as you go, not after

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Wednesday, April 1 (Target: 2 hours)

**Focus: TGPL Chapter 2 — Program Structure**

- [ ] Read TGPL Chapter 2 (60 min)
  - Focus on: variables, types, pointers, the `new` function, variable scope, and the blank identifier `_`
  - Take notes in `notes/books/tgpl-chapter-2.md`
  - Write a short code snippet for each new concept in your notes
- [ ] Finish Gophercises #2: `quiz` game (45 min)
  - Add a timer to the quiz (hint: `time.AfterFunc` or `time.NewTimer`)
  - This introduces you to the `time` package and basic concurrency patterns
- [ ] Commit all code (15 min)
  - Push both gophercises to GitHub

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Thursday, April 2 (Target: 1.5 hours)

**Focus: Write Tests + Explore Go Toolchain**

- [ ] Write tests for Gophercises #2 (45 min)
  - Practice table-driven tests for quiz scoring logic
  - Aim for >70% coverage even at this early stage
  - Run: `go test -cover ./...`
- [ ] Explore Go tooling (45 min)
  - Run `go vet` on your code — fix any warnings
  - Run `gofmt -l .` — ensure all files are formatted
  - Explore `go doc` for packages you've used (`fmt`, `os`, `bufio`, `time`)
  - Bookmark https://pkg.go.dev/ — your primary reference

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Friday, April 3 (Target: 1 hour)

**Focus: Community + Reflection**

- [ ] Post your progress in the Go Slack or Reddit r/golang (15 min)
  - Share what you built this week, ask one specific question
  - Example: "Just finished gophercises #1 and #2 — any tips for structuring CLI args better?"
- [ ] Review your notes from Monday–Thursday (15 min)
  - What concepts are still fuzzy? Write them down as questions.
- [ ] Update your journal entry for Days 1–5 (30 min)
  - File: `journal/2026-03-week-13.md` (copy from `JOURNAL-TEMPLATE.md`)
  - Fill in: time spent, what you learned, challenges, aha moments

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Saturday, April 4 (Target: 1.5 hours)

**Focus: Build Something Small on Your Own**

Choose ONE of these mini-projects to build with no guidance:

**Option A: Simple Calculator CLI**
- Input: two numbers and an operator (`+`, `-`, `*`, `/`)
- Output: result with error handling (e.g., divide by zero)
- Requirements: use `os.Args` or `bufio.Scanner`, write 3+ tests

**Option B: Word Frequency Counter**
- Input: a text file (use `os.Open`)
- Output: top 10 most frequent words with counts
- Requirements: maps, sorting by value, write 3+ tests

**Option C: FizzBuzz Extended**
- Go beyond the basic version: accept N and custom words for 3 and 5 via flags
- Requirements: `flag` package, write 5+ table-driven tests

- [ ] Pick one and build it from scratch (75 min)
- [ ] Write tests as you build (not after) (30 min)
- [ ] Push to GitHub under `phase-1/week-2-mini-project/`
- [ ] Add a `README.md` explaining what it does and how to run it

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Sunday, April 5 (Target: 0.5 hours)

**Focus: Weekly Review + Plan Week 3**

- [ ] Complete journal weekly summary section (20 min)
  - Total hours logged: ___
  - Did I hit 10 hours? Yes / No / Partially
  - What went well?
  - What was harder than expected?
  - Questions to research next week
- [ ] Plan Week 3 tasks (10 min)
  - Next: TGPL Chapters 3–4 + Gophercises #3 and #4
  - Check your `roadmap/PHASE-1-MODULES.md` → Submodule 1.1 for what's next
  - Write Week 3 goals in `roadmap/weekly-goals.md`

**Actual time spent**: ___ hours
**What I accomplished**:

---

## Week 2 Summary

### Total Time Logged
**Planned**: 10 hours
**Actual**: ___ hours
**Variance**: ___ hours (over/under)

### Deliverables to Check Off
- [ ] TGPL Chapters 1–2 read with notes taken
- [ ] Gophercises #1 (echo) — done + tested + pushed
- [ ] Gophercises #2 (quiz) — done + tested + pushed
- [ ] 1 mini-project built independently — done + pushed
- [ ] Tests written for all code (>70% coverage)
- [ ] Daily commits to GitHub (5+ commits this week)
- [ ] Journal entry filled in
- [ ] Week 3 planned

---

## Key Concepts You Should Understand by End of Week 2

If you can answer these without looking them up, you're on track:

**From TGPL Chapter 1**
- What does `package main` mean, and why is it required for executables?
- What is the difference between `var` declaration and `:=` short declaration?
- How do you read from standard input in Go?

**From TGPL Chapter 2**
- What is the zero value of an `int`? A `string`? A `bool`? A pointer?
- What does `new(T)` return?
- What is the blank identifier `_` and when would you use it?
- What is package-level scope vs. function-level scope?

**From Gophercises**
- How do you parse command-line flags in Go?
- How do you read a CSV file in Go?
- How do you use `time.NewTimer` or `time.AfterFunc`?

---

## If You're Behind

**Couldn't finish gophercises?**
→ That's okay. Carry them to Week 3. Don't skip ahead.
→ Do NOT move to Gophercises #3 until #1 and #2 are done.

**Couldn't hit 10 hours?**
→ Log honestly: How many hours did you get? ___
→ What blocked you? ___
→ Adjust your Week 3 schedule to make up hours.

**Stuck on the quiz timer?**
→ Look at Go's `time` package docs: https://pkg.go.dev/time
→ Try `time.After(duration)` — it returns a channel you can use in `select`
→ Spend 30 minutes debugging before asking for help

**Everything feels confusing?**
→ That's normal in Week 2. You're building mental models.
→ Don't skip steps to "feel faster." Confusion now = understanding later.
→ Write down what's confusing. That's your study list for this week.

---

## Week 3 Preview

Next week you'll:
1. Read TGPL Chapters 3–4 (basic data types, composite types)
2. Complete Gophercises #3 and #4
3. Start thinking about Project 1.1.1: **Log Parser CLI** (the first real project)
4. Begin writing more structured code with multiple files in one package
5. Continue daily commits and journaling

The Log Parser is where Module 1.1 gets serious. Week 3 is your warm-up for it.

---

## Commitment Check

Did you uphold your 5 commitments this week?

1. [ ] 10 focused hours (or honest partial count logged)
2. [ ] Sequential learning — only Module 1.1 content, nothing else
3. [ ] Built without AI writing the code for you
4. [ ] Wrote notes and journal entries (communication practice)
5. [ ] Tracked progress honestly in your time log

**Streak**: ___ weeks consistent

---

## Resources for This Week

| Resource | Use For |
|----------|---------|
| TGPL Book (Donovan & Kernighan) | Main reading — Chapters 1–2 |
| https://gophercises.com/ | Gophercises #1 and #2 |
| https://pkg.go.dev/time | Timer for quiz game |
| https://pkg.go.dev/flag | Flag parsing |
| https://go.dev/blog/laws-of-reflection | Save for later — just bookmark |
| https://gobyexample.com/command-line-flags | Quick reference for flags |
| https://gobyexample.com/testing | Quick reference for `testing` package |

---

## Notes
(Any additional thoughts, observations, or things to remember)




---

**Next**: After completing Week 2, read `roadmap/PHASE-1-MODULES.md` → Submodule 1.1 section for Week 3's exact tasks.

**Save journal as**: `journal/2026-03-week-13.md`

**GitHub target**: `principal-engineer-portfolio/phase-1/` — should have 3+ folders by end of week.