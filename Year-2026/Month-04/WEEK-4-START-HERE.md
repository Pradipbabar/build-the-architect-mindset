# Week 4: TGPL Chapters 7–9 + Concurrency Basics (Realistic Pace)

**Week**: Week 4 of Phase 1 — Module 1.2 (Concurrency Fundamentals)
**Dates**: April 13 – April 19, 2026
**Total Hours Goal**: 6-7 hours (concurrency is harder, take your time)
**Roadmap Reference**: `roadmap/PHASE-1-MODULES.md` → Module 1, Submodule 1.3

---

## Week 4 Goals

**Primary Goal**: Understand goroutines, channels, and basic concurrency patterns. Read Chapters 7-9.

**Hours Target**: 6-7 hours TOTAL for the week

---

## OPTION A: Weekday Spread + Weekend Project

**Monday, April 13 — Chapter 7 (1 hour)**
- [ ] Read TGPL Chapter 7 — Goroutines & Channels (part 1)
- [ ] Type 3-4 examples of goroutines
- [ ] Create: `notes/books/tgpl-chapter-7.md`
- [ ] Commit: `"week-4: tgpl chapter 7 goroutines notes"`

**Tuesday, April 14 — Chapter 8 (1.5 hours)**
- [ ] Read TGPL Chapter 8 — Channels Deep Dive
- [ ] Type 3-4 examples (send, receive, close, select)
- [ ] Create: `notes/books/tgpl-channel-8.md`
- [ ] Commit: `"week-4: tgpl chapter 8 channels notes"`

**Wednesday, April 15 — Chapter 9 (1 hour)**
- [ ] Read TGPL Chapter 9 — Concurrency Patterns
- [ ] Type 2-3 examples
- [ ] Create: `notes/books/tgpl-chapter-9.md`
- [ ] Commit: `"week-4: tgpl chapter 9 concurrency patterns"`

**Thursday, April 16 — Project Planning (30 min)**
- [ ] Decide on concurrent project (see below)
- [ ] Write plan in comments

**Friday-Sunday, April 17-19 — Build Concurrent Program (2.5-3 hours)**
- [ ] Friday: Basic implementation: 1.5 hours
- [ ] Saturday: Add goroutines/channels: 1 hour
- [ ] Sunday: Tests + quality checks: 30 min

---

## OPTION B: Weekend Focus

**Saturday, April 13 (3.5 hours)**
- [ ] Read Chapters 7-8: 2 hours
- [ ] Type examples: 1 hour
- [ ] Create notes: 30 min
- [ ] Commit: `"week-4: tgpl chapters 7-8 notes"`

**Sunday, April 14 (3.5 hours)**
- [ ] Read Chapter 9: 30 min
- [ ] Plan concurrent project: 30 min
- [ ] Implement: 2 hours
- [ ] Tests + quality checks: 30 min
- [ ] Commit: `"week-4: concurrent program complete"`

---

## OPTION C: Flexible

**Just ensure by end of week**:
- [ ] Chapters 7-9 fully read
- [ ] Notes created for all 3 chapters
- [ ] Concurrent program planned
- [ ] Basic implementation complete
- [ ] Uses goroutines and channels
- [ ] At least 2 tests written
- [ ] `go test -race` passes (no data races)
- [ ] Code passes gofmt and go vet

---

## Concurrent Program Ideas (Pick ONE)

### Project 1: Concurrent Downloader
- Download 3-5 URLs concurrently
- Report success/failure for each
- Use goroutines (one per URL)
- Collect results via channel
- Tests: 2-3 tests

### Project 2: Simple Worker Pool
- Create 3 workers from a channel
- Send jobs, workers process them
- Collect results
- Tests: 2-3 tests

### Project 3: File Processor with Goroutines
- Process multiple files concurrently
- One goroutine per file
- Collect results on channel
- Tests: 2-3 tests

---

## Important Notes

**Take your time with concurrency**. It's a bigger mental shift than Chapters 1-6. If you need 7-8 hours, that's fine. Log it and move forward. Don't rush through this.

---

## Time Breakdown

| Task | Target | Actual |
|---|---|---|
| TGPL Chapters 7-9 | 3 hours | ___ |
| Concurrent Program | 3 hours | ___ |
| Quality checks | 30 min | ___ |
| **Total** | **6-7 hours** | ___ |

---

## End of Week Checklist

- [ ] Chapters 7-9 read
- [ ] Notes created for each chapter
- [ ] Concurrent program planned
- [ ] Basic implementation complete
- [ ] Uses goroutines
- [ ] Uses channels
- [ ] At least 2 tests written + passing
- [ ] `go test -race` passes
- [ ] Code passes gofmt and go vet
- [ ] Project committed to GitHub
- [ ] Total hours logged: ___ / 6-7

---

## Actual Time Spent

**Total hours: ___**

**What I accomplished**:

**What I struggled with**:

**Concurrency concepts I understand**:

**Concurrency concepts to practice more**:
