# April 5, 2026 — Day 5 of Week 2

**Date**: Saturday, April 5, 2026
**Week**: Week 2 (Apr 1-7) — TGPL Chapters 1-4 + First Project Foundation
**Day Number**: 5 of 7 days
**Total Phase Hours Goal**: 60-80 hours | **Week Goal**: 5-6 hours | **Today's Goal**: 60-75 min

---

## Overview

Today is Day 5 of Week 2. You have:
- ✅ Read TGPL Chapters 1-4 (Apr 1-2)
- ✅ Planned and set up project (Apr 3)
- ✅ Implemented core logic + first tests (Apr 4)
- ⏳ **Today**: Complete project, add comprehensive tests, polish

---

## Today's Schedule
**Total Time**: 60-75 minutes

**Part 1: Finish Implementation (25 min)**
- [ ] Implement remaining functions from your plan
- [ ] Focus on: What's still not working or incomplete?
- [ ] Examples:
  - Quiz Game: Timer implementation, scoring
  - Echo CLI: All flag combinations
  - Text Processor: More complex operations
- [ ] Test each new function immediately: `go test -v`
- [ ] Commit: `"day-5: [project] complete core implementation"`

**Part 2: Comprehensive Testing (20 min)**
- [ ] Add more test cases for edge cases:
  - Empty inputs
  - Invalid data
  - Boundary conditions
  - Normal / typical usage
- [ ] Aim for 3-5 test cases per function
- [ ] Run full test suite: `go test -v`
- [ ] Target: 40%+ code coverage (we'll improve later)
- [ ] Commit: `"day-5: [project] comprehensive tests"`

**Part 3: Code Review & Refactoring (10 min)**
- [ ] Read through your own code
- [ ] Make variable names clearer if needed
- [ ] Remove any commented-out code
- [ ] Split large functions if they're doing too much
- [ ] Commit: `"day-5: [project] refactoring and cleanup"`

**Part 4: Quality & Documentation (5 min)**
- [ ] `gofmt -w .`
- [ ] `go vet ./...`
- [ ] Fix all warnings
- [ ] Add brief comments to complex functions
- [ ] Create or update README.md for the project

---

## Optional: If You Want to Do More

**Option 1: Focus on Testing** (If implementation is mostly done)
- Fix any major bugs blocking functionality (20 min)
- Write 5-8 comprehensive test cases covering all scenarios (30 min)
- Run full suite: `go test -v -cover`
- Final quality checks: gofmt, go vet

**Option 2: Flexible Pace** (Work at your own speed)
- Just ensure by bedtime:
  - All major functionality implemented
  - 5+ test cases written and passing
  - Code passes `gofmt` and `go vet`
  - README created for the project
  - All progression commits made
  - Ready to show someone your code

---

## Testing Strategy: From Basic to Comprehensive

**Level 1: Basic Testing** (what you likely did on Day 4)
```go
func TestExample(t *testing.T) {
    got := MyFunction("input")
    want := "expected"
    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
```

**Level 2: Table-Driven Testing** (what to do today)
```go
func TestExample(t *testing.T) {
    tests := []struct {
        name string
        input string
        want string
    }{
        {"normal case", "hello", "HELLO"},
        {"empty string", "", ""},
        {"with spaces", "hello world", "HELLO WORLD"},
        {"numbers", "123", "123"},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := MyFunction(tt.input)
            if got != tt.want {
                t.Errorf("got %q, want %q", got, tt.want)
            }
        })
    }
}
```

**Why table-driven**:
- ✅ Easy to add more test cases
- ✅ Clear what each test does (name field)
- ✅ Less repeated code
- ✅ Go language best practice

---

## Test Coverage Tips

**Check coverage**:
```bash
go test -cover ./...
```

**See which lines aren't tested**:
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out  # Opens in browser
```

**Aim for today**: 40%+ coverage
(We'll improve this next week)

---

## Example: Complete Quiz Game Test Structure

If you're building the Quiz Game, your tests might look like:

```go
package main

import "testing"

func TestParseProblems(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        wantLen int
    }{
        {
            "basic problems",
            "1+1,2\n2+2,4\n",
            2,
        },
        {
            "single problem",
            "5+5,10\n",
            1,
        },
        {
            "empty",
            "",
            0,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := ParseProblems(strings.NewReader(tt.input))
            if len(got) != tt.wantLen {
                t.Errorf("got %d problems, want %d", len(got), tt.wantLen)
            }
        })
    }
}
```

---

## Git Workflow for Today

Show progression with commits:
```
day-5: [project] implement remaining functions
day-5: [project] add comprehensive tests
day-5: [project] refactoring and cleanup
day-5: [project] quality checks and docs
```

**Commit often** — it shows your work progression.

---

## Project README Template

Create a `README.md` in your project directory:

```markdown
# [Project Name]

## Overview
Brief description of what this program does.

## Building

\`\`\`bash
go build -o [program-name]
\`\`\`

## Running

\`\`\`bash
./[program-name] [args]
\`\`\`

## Examples

\`\`\`bash
$ ./program-name input1 input2
Result: something useful
\`\`\`

## Testing

\`\`\`bash
go test -v
\`\`\`

## Code Structure
- `main.go` - Entry point
- `functions.go` - Core logic
- `*_test.go` - Tests

## Learnings
- What I learned building this
- Challenges I overcame
```

---

## Code Quality Checklist

Before final commit:
- [ ] All tests pass: `go test -v`
- [ ] Code formatted: `gofmt -w .`
- [ ] No vet warnings: `go vet ./...`
- [ ] Variable names are clear and descriptive
- [ ] Functions are small and focused
- [ ] Comments explain complex logic
- [ ] No commented-out code left behind
- [ ] README created/updated
- [ ] 40%+ test coverage

---

## Refactoring Checklist

Ask yourself about each function:

**Is this function clear?**
- ✅ If yes, leave it
- ❌ If no, add a comment or rename variables

**Does this function do one thing?**
- ✅ If yes, great
- ❌ If no, split it into multiple functions

**Is this function too long?**
- ✅ If <20 lines, good
- ⚠️ If 20-50 lines, consider splitting
- ❌ If >50 lines, definitely split

**Are variable names clear?**
- ✅ `totalScore`, `userInput`, `isValid`
- ❌ `ts`, `ui`, `iv`, `x`, `y`, `temp`

---

## End of Date Checklist

- [ ] All core functionality implemented and working
- [ ] 5+ comprehensive test cases written and passing
- [ ] Code passes gofmt and go vet
- [ ] Code review and refactoring done
- [ ] README created for project
- [ ] All commits made with clear messages
- [ ] Time logged

---

## End of Week Reflection (For tomorrow)

**By the end of Day 5, reflect on**:

**What I built**:
- [ ] Completed project: _________________
- [ ] Core features: ___________________

**What's working well**:
- 

**What was challenging**:
- 

**What I learned**:
- About Go language
- About testing practices
- About my own coding style

**Time breakdown**:
- Day 1: ___ hours (reading)
- Day 2: ___ hours (reading)
- Day 3: ___ hours (planning)
- Day 4: ___ hours (coding)
- Day 5: ___ hours (polishing)
- **Week Total**: ___ / 5-6 hours

**Test coverage**: ___%

**Code quality**: 
- gofmt: ✓ / ✗
- go vet: ✓ / ✗
- Tests passing: ✓ / ✗

---

## Next Week (Days 6-7 / Sunday-Monday)

- Rest Day 6 (Sunday) - optional review
- Day 7 (Monday, April 7): Week review + plan Week 3

---

## Success Metrics for Week 2

By end of week, you should have:
- ✅ Read all 4 chapters thoroughly
- ✅ Built a working CLI tool or quiz game
- ✅ Written comprehensive tests
- ✅ Practiced code quality standards
- ✅ Experienced the Go development workflow
- ✅ Committed code to GitHub

**If you have all of these**: Week 2 is a success! 🎉

---

## Before You Close Your Work Today

**To-do**:
1. Run `go test -v` one final time — all passing?
2. Run `gofmt -w .` and `go vet ./...`
3. Make final commit: `"day-5: [project] complete and polished"`
4. Push to GitHub: `git push origin master`
5. Update learning journal with today's notes
6. Log total hours spent this week

**You're done!** Your first Go project is complete. Take a moment to appreciate it.

