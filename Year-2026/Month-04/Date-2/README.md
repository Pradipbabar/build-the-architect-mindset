# Day 4 (Week 2) — 2 April 2026

**Phase**: Phase 1 / Module 1.1 — Go Fundamentals  
**Focus**: Code Quality + Advanced Practice + Gophercises #3 (Optional Stretch)  
**Time Target**: 2 hours  
**Week**: Week 2 of Phase 1 (March 30 – April 5, 2026)  

---

## Context: Where You Are in Week 2

By end of Wednesday (April 1) you completed:
- ✅ TGPL Chapter 1 — typed all examples, notes written
- ✅ TGPL Chapter 2 — program structure, zero values, pointers, named types
- ✅ TGPL Chapter 3 — integers, floats, strings, runes, iota, float comparison
- ✅ TGPL Chapter 4 — arrays, slices, maps, structs, embedding
- ✅ Gophercises #1 (Echo CLI) — done, table-driven tests written, pushed
- ✅ Gophercises #2 (Quiz Game) — complete with timer + tests, pushed

You are **significantly ahead** of the original week plan. Thursday's original goal (Chapter 4 + Gophercises #2) is already done. This gives you flexibility today.

**Running total check** (update honestly):
- Hours logged so far this week: ___ / 12
- Phase 1 total: ___ / 200 hours

---

## Today's Big Picture

**Thursday target**: 2 hours for one of these paths:

### Path A: Code Quality Deep Dive (Recommended) — 2 hours
Master the Go toolchain and ensure all code written this week meets professional standards. This is not busywork — these habits are how professional Go developers work.

### Path B: Stretch — Gophercises #3 (if you want more challenge) — 2 hours
Start Gophercises #3 (CSV merge) if you want to push further this week. Not required.

### Path C: Review & Experimentation — 2 hours
Go back through Chapter 4 concepts and experiment with advanced slice/map patterns. Deepen understanding rather than add new material.

---

## Session 1: Code Quality Audit (Recommended Path A) — 60 min

This is the most important session for your professional development. Every line of code you write this week must pass these checks before going to GitHub.

### The Four Mandatory Checks

**1. Format Check: `gofmt`**
```bash
cd /home/pradip/repos/build-the-architect-mindset/projects/phase-1/00-day-4-gophercises
gofmt -w .  # Fix all formatting issues
```

Go code has ONE standard format. Everyone's code looks the same. No debates about tabs vs spaces — the toolchain decides.

- [ ] All `.go` files are gofmt-compliant
- [ ] Run `gofmt -l .` — should output nothing

**2. Vet Check: `go vet`**
```bash
cd /home/pradip/repos/build-the-architect-mindset/projects/phase-1/00-day-4-gophercises
go vet ./...
```

`go vet` catches real bugs:
- Incorrect printf format strings
- Unreachable code
- Useless assignments
- Type mismatches that compile but are wrong

- [ ] All packages pass `go vet`
- [ ] No warnings or errors

**3. Test Coverage Check**
```bash
cd /home/pradip/repos/build-the-architect-mindset/projects/phase-1/00-day-4-gophercises
go test -v -cover ./...
```

Check coverage for each exercise:
- `01-echo`: Report coverage %
- `02-quiz`: Report coverage %

Target: >75% coverage minimum. 85%+ preferred.

- [ ] `01-echo` coverage: ___ %
- [ ] `02-quiz` coverage: ___ %

**4. Build Check**
```bash
cd /home/pradip/repos/build-the-architect-mindset/projects/phase-1/00-day-4-gophercises
# For each exercise, verify it compiles and runs
go run 01-echo/main.go hello world
go run 02-quiz/main.go -csv 02-quiz/problems.csv -limit 10
```

Both should run without errors.

- [ ] `01-echo` compiles and runs
- [ ] `02-quiz` compiles and runs

### Commit Your Quality Improvements

```bash
git add .
git commit -m "week-2: code quality audit - gofmt, go vet, test coverage reviewed"
git push
```

---

## Session 2: Choose One (60 min)

### Option A: Explore the Go Toolchain (Recommended for Understanding)

You use these tools daily for the next 12 months. Understand them deeply:

```bash
# 1. Learn what documentation looks like
go doc fmt
go doc fmt.Printf

# 2. Check for compiler optimization opportunities
go build -v 01-echo/main.go

# 3. Understand module dependencies
cat go.mod
go mod graph

# 4. See what `go generate` can do (advanced)
go doc go/types
```

Spend 30 minutes reading output, understanding what each tool tells you.

Write your findings in `notes/concepts/Go_Basics/Go_Toolchain.md`:

```markdown
# Go Toolchain Exploration

## go doc
- Purpose: ...
- How to use: ...
- When I'd use this: ...

## go build vs go run
- Difference: ...
- When to use each: ...

## go vet
- What it catches: ...
- Example of a bug it finds: ...

## go test -cover
- What coverage means: ...
- Why it matters: ...
```

### Option B: Gophercises #3 — CSV Merge (Stretch)

URL: https://gophercises.com/ → Exercise 3

**Before you start**:
- This exercise combines file I/O, slicing, and data structures
- Less hand-holding than #1 and #2 — you write the plan
- Use the same quality discipline: table-driven tests, `go vet`, coverage

**Steps**:
1. [ ] Read the problem statement completely
2. [ ] Write your plan in comments (no code yet)
3. [ ] Implement the basic version
4. [ ] Write 3+ table-driven tests
5. [ ] Verify: `go vet`, `gofmt`, `go test -cover`
6. [ ] Commit: `git add 03-csv && git commit -m "gophercises-03: csv merge"`

**Remember**: The exercise description is intentionally vague. Good. That's realistic. Professional engineers read incomplete specs and ask clarifying questions. Your question: "Wait, what should happen if a file doesn't exist?" Answer: try it and handle the edge case.

### Option C: Deep Dive — Chapter 4 Advanced Patterns

Pick ONE of these and spend 60 minutes getting very comfortable with it:

**1. Slice Tricks**
```go
// Reverse a slice
for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
    s[i], s[j] = s[j], s[i]
}

// Remove an element
s = append(s[:i], s[i+1:]...)

// Filter a slice
n := 0
for _, x := range s {
    if keep(x) {
        s[n] = x
        n++
    }
}
s = s[:n]
```

Type these. Understand them. Write 3 small test functions for each. Try variations.

**2. Interface{} — the Any Type**
```go
// This is how Go does dynamic typing
var v interface{}
v = 42
v = "hello"
v = []int{1, 2, 3}

// Type assertion
i, ok := v.(int)  // Is v actually an int?
```

Write 5 small test cases.

**3. Map Patterns**
```go
// Map of maps
m := make(map[string]map[string]int)
m["user1"]["score"] = 100

// Counting occurrences
counts := make(map[string]int)
for _, word := range words {
    counts[word]++
}

// Lookup with default
func getOrDefault(m map[string]int, key string, defaultVal int) int {
    if v, ok := m[key]; ok {
        return v
    }
    return defaultVal
}
```

Write tests. Make sure you understand the two-value form: `v, ok := m[key]`.

Put your experiments in a new file: `projects/phase-1/00-simple-projects/chapter4-advanced.go`

---

## Time Breakdown

| Task | Target | Actual |
|---|---|---|
| Session 1: Code Quality Audit | 60 min | ___ |
| Session 2: Toolchain OR Gophercises #3 OR Chapter 4 Deep Dive | 60 min | ___ |
| **Total** | **2 hours** | ___ |

---

## Quality Checklist (All Mandatory)

- [ ] All code passes `gofmt -l .` (returns nothing)
- [ ] All code passes `go vet ./...` (no warnings)
- [ ] Test coverage measured for all exercises
- [ ] Coverage % logged in this file
- [ ] Session 2 code (if written) also passes gofmt + go vet
- [ ] All new code committed with meaningful message
- [ ] Hours logged accurately — be honest

---

## Notes & Reflections (Fill In As You Work)

### Code Quality Findings

What issues did `go vet` find?
- 

What formatting changes did `gofmt` make?
- 

Which exercise has the best test coverage?
- 

---

### Session 2 Progress

**If you chose Option A (Toolchain)**:
- Most useful tool you discovered:
- Surprised by:
- Will use regularly:

**If you chose Option B (Gophercises #3)**:
- How different is #3 from #1 and #2?
- What tripped you up?
- Test coverage achieved: ____%

**If you chose Option C (Chapter 4 Deep Dive)**:
- Pattern you found most useful:
- Experiments that changed your understanding:
- Will practice more:

---

## End of Day Checklist

- [ ] Session 1: Code quality audit complete on all exercises
- [ ] `gofmt`, `go vet`, coverage metrics captured
- [ ] Session 2: Chose path A, B, or C and completed it
- [ ] All new/modified code follows quality standards
- [ ] All work committed to GitHub
- [ ] Hours logged accurately

---

## Actual Time Spent & Summary

**Total hours: ___ (Target: 2)**

**Session 1 (Code Quality)**:
- gofmt issues fixed: ___
- go vet warnings resolved: ___
- 01-echo coverage: ____%
- 02-quiz coverage: ____%

**Session 2 (Toolchain/Gophercises #3/Chapter 4)**:
- Option chosen: [ ] A — Toolchain  [ ] B — Gophercises #3  [ ] C — Chapter 4
- What I accomplished:

**What I struggled with**:

**What to revisit tomorrow (Friday)**:

---

## Tomorrow: Friday, April 3

**Focus**: If you did Path A and have extra time, start one Gophercises exercise from #3 onwards. If you did Path B or C, finish the assignment and ensure it passes all quality checks.

Final day of Week 2: wrap up any incomplete exercises, ensure everything is pushed to GitHub, and log your total hours. Week 3 (Monday–Friday next week) escalates to Chapters 5–7 where Go's concurrency model and interface design become the focus.

Remember: You are building habits that professional Go developers maintain for 20+ year careers. Discipline now (gofmt, go vet, tests) saves hours of debugging later.
