# April 3, 2026 — Day 3 of Week 2

**Date**: Thursday, April 3, 2026
**Week**: Week 2 (Apr 1-7) — TGPL Chapters 1-4 + First Project Foundation
**Day Number**: 3 of 7 days
**Total Phase Hours Goal**: 60-80 hours | **Week Goal**: 5-6 hours | **Today's Goal**: 45-60 min

---

## Overview

Today is Day 3 of Week 2. By now you should have:
- ✅ Read TGPL Chapters 1-2 (completed Apr 1)
- ✅ Read TGPL Chapters 3-4 (completed Apr 2)
- ⏳ **Today**: Quality audit of Chapters 1-4 + start first project

---

## Today's Schedule
**Total Time**: ~50-60 minutes

**Part 1: TGPL Chapters 1-4 Quality Audit (20 min)**
- [ ] Review your notes from chapters 1-4
- [ ] Identify 2-3 concepts that were unclear
- [ ] Re-read those specific sections (not full chapters)
- [ ] Update your notes with clarity
- [ ] Example: If interfaces confused you, re-read interface section only

**Part 2: First Project Planning (15 min)**
- [ ] Open `/projects/phase-1/00-day-4-gophercises/`
- [ ] Choose: **Gophercises #1 (Echo CLI)** OR **#2 (Quiz Game)**
- [ ] Read problem description
- [ ] Write pseudocode in comments (don't code yet)
- [ ] Design your approach

**Part 3: Project Setup (10 min)**
- [ ] Create directory: `/projects/phase-1/XX-day-X-project-name/`
- [ ] Create `main.go` with skeleton
- [ ] Create `go.mod` file
- [ ] Commit: `"day-3: [project-name] setup and planning"`

**Part 4: Code Quality Check (5 min)**
- [ ] `gofmt -l .`
- [ ] `go vet ./...`
- [ ] Log any issues found

---

## Optional: If You Want to Do More

**Option 1: Lighter Review** (If catching up on chapters)
- Skim chapters 1-4 notes instead of deep review
- Jot down 1 question per chapter to explore later
- Skip Part 1 and go straight to Part 2

**Option 2: Flexible Approach** (Work at your own pace)
- Just ensure by bedtime:
  - Reviewed notes from chapters 1-4 (at least skim)
  - Decided which project to start
  - Created project directory structure
  - Wrote basic go.mod file
  - Committed: `"day-3: [project] planning"`

---

## Project Options (Pick ONE for this week)

### Option 1: Quiz Game (Gophercises #2) ⭐
**What you'll build**: 
- Program that reads quiz questions/answers from CSV
- Asks questions one by one
- Counts correct answers
- Times the quiz (30 seconds default)

**Why it's good**:
- Uses file I/O, CSV parsing, basic concurrency
- Teaches testing patterns
- Practical CLI tool

**Relevant chapters**: 1-3 (syntax, data structures, functions)

**Time to complete**: 4-5 hours total (finish by end of Week 2)

**Start with**: File structure, CSV parsing, basic functions

---

### Option 2: Echo CLI (Gophercises #1)
**What you'll build**:
- CLI tool that processes command-line arguments
- Takes input and manipulates it (reverse, shuffle, sort)
- Outputs results

**Why it's good**:
- Simpler entry point
- Teaches CLI patterns and flags
- Good for testing practice

**Relevant chapters**: 1-3 syntax

**Time to complete**: 2-3 hours total

**Start with**: Parsing command-line flags, basic string manipulation

---

### Option 3: Text Processor from Scratch
**What you'll build**:
- CLI tool that reads text file
- Counts: lines, words, characters
- Sorts/finds patterns
- Outputs statistics

**Why it's good**:
- Combine multiple concepts
- File I/O practice
- Good for learning testing

**Time to complete**: 3-4 hours

---

## Today's Specific Tasks

| Time | Task | Checkpoint |
|------|------|-----------|
| 5 min | Set up focused work environment | Clean desk, no distractions |
| 10 min | Review TGPL chapters 1-4 notes | Note unclear concepts |
| 10 min | Choose project | Decision made |
| 15 min | Create project skeleton | go.mod created |
| 10 min | Write pseudocode | Comments in main.go |
| 5 min | Commit and quality check | gofmt + go vet pass |
| 5 min | Log in journal | Notes updated |

---

## Code Quality Standards (apply before each commit)

```bash
# Format your code
gofmt -w .

# Check for common mistakes
go vet ./...

# If you have tests, run them
go test -v

# Check code coverage (later)
go test -cover
```

---

## Git Commit Guidelines

**Today's commits should look like**:
```
day-3: [project-name] setup and planning
day-3: [project-name] skeleton and pseudocode
```

**Include in commit**:
- go.mod file
- main.go with skeleton + comments
- Possibly README with project description

---

## End of Day Reflection

**Before you close your work**, answer:

**What I accomplished**:
- [ ] Reviewed chapters 1-4 notes
- [ ] Chose project
- [ ] Set up directory structure
- [ ] Understood the problem

**What was confusing**:
- 

**What I'm excited to build tomorrow**:
- 

**Time spent**: ___ min / 50-60 min goal

**Next day goal**: Start implementing core logic

---

## Resource Links

**TGPL Chapters Reference**:
- Chapter 1: Getting Started (syntax, types basics)
- Chapter 2: Program Structure (packages, declarations, scope)
- Chapter 3: Data Types (integers, floats, strings, arrays, slices, maps)
- Chapter 4: Composite Types (structs, JSON marshaling)

**Gophercises** (if you're doing Quiz Game):
- `https://gophercises.com/#2`
- Problem statement + requirements
- Hints and solutions (use after you try!)

---

## Common Questions

**Q: Should I start coding today?**
A: No — today is planning, setup, and understanding. Coding starts tomorrow (Day 4).

**Q: Which project should I pick?**
A: Pick the one that sounds most interesting. You can't go wrong. Quiz Game (#2) is most practical.

**Q: What if I'm not done with chapters 1-4 review?**
A: That's OK. Review what you can, and don't obsess. You'll learn more by building code.

**Q: Can I use external packages?**
A: For now, stick to standard library. You can explore external packages next week.

---

## Phase 1 Progress

**Hours logged so far**:
- Day 1 (Apr 1): ___ hours
- Day 2 (Apr 2): ___ hours
- Day 3 (Apr 3): ___ hours (TODAY)
- **Week 2 Total**: ___ / 5-6 hours

**Toward Phase 1 Goal**: ___ / 60-80 hours

---

## Notes for Tomorrow (Day 4)

**Day 4 (Friday, April 4)**: Begin coding the project
- You'll have gone through setup and planning
- Start with core logic
- Write basic tests

**Prep tonight**:
- Make sure everything compiles (`go build`)
- Your pseudocode in comments is clear
- You understand the problem statement

