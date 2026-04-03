# April 4, 2026 — Day 4 of Week 2

**Date**: Friday, April 4, 2026
**Week**: Week 2 (Apr 1-7) — TGPL Chapters 1-4 + First Project Foundation
**Day Number**: 4 of 7 days
**Total Phase Hours Goal**: 60-80 hours | **Week Goal**: 5-6 hours | **Today's Goal**: 60-90 min

---

## Overview

Today is Day 4 of Week 2. You have:
- ✅ Read TGPL Chapters 1-4 (Apr 1-2)
- ✅ Planned and set up first project (Apr 3)
- ⏳ **Today**: Implement core project logic + first tests

---

## Today's Schedule
**Total Time**: 75-90 minutes

**Part 1: Core Logic Implementation (40 min)**
- [ ] Implement main functionality based on pseudocode
- [ ] Break into small functions (each 5-10 lines)
- [ ] Test each function as you go (manual testing)
- [ ] Commit: `"day-4: [project] core logic v1"`
- [ ] Example functions:
  - If Quiz Game: `parseCSV()`, `askQuestion()`, `startTimer()`
  - If Echo CLI: `parseFlags()`, `processText()`, `outputResult()`

**Part 2: First Tests (20 min)**
- [ ] Write 1-2 basic tests for your core functions
- [ ] Use table-driven tests (Go pattern):
  ```go
  func TestParseCSV(t *testing.T) {
      tests := []struct {
          name    string
          input   string
          want    int
      }{
          {"valid", "1,2", 2},
          {"empty", "", 0},
      }
      for _, tt := range tests {
          t.Run(tt.name, func(t *testing.T) {
              got := parseCSV(tt.input)
              if got != tt.want {
                  t.Errorf("got %d, want %d", got, tt.want)
              }
          })
      }
  }
  ```
- [ ] Run tests: `go test -v`

**Part 3: Code Quality (10 min)**
- [ ] `gofmt -w .`
- [ ] `go vet ./...`
- [ ] Fix any warnings
- [ ] Commit: `"day-4: [project] formatting and quality"`

**Part 4: Progress Review (5 min)**
- [ ] List what works
- [ ] List what's broken
- [ ] Plan for tomorrow (Saturday)

---

## Optional: If You Want to Do More

**Option 1: Incremental Approach** (If time is tight)
- Pick ONE critical function and implement it completely
- Write 1-2 test cases for it
- Focus on quality over quantity
- Finish with gofmt and go vet cleanup

**Option 2: Flexible Pace** (Work at your own speed)
- Just ensure by bedtime:
  - At least 1 core function implemented
  - 1-2 tests written and passing
  - Code passes `gofmt` and `go vet`
  - Meaningful commit made
  - Time and progress logged

---

## Implementation Guide by Project Type

### If Building Quiz Game (Gophercises #2):

**Functions to implement today**:
1. `parseCSV(filename string) []Problem` — Read CSV file, return slice of problems
2. `askQuestion(p Problem) bool` — Display question, read answer, return if correct

**Test focus**:
- Test CSV parsing with sample data
- Test question presentation

**Don't worry about today**:
- Timer (tackle tomorrow)
- Scoring (tackle tomorrow)

---

### If Building Echo CLI (Gophercises #1):

**Functions to implement today**:
1. `parseFlags() Options` — Parse -h, -n, -s flags
2. `processText(text string, opts Options) string` — Apply operations

**Test focus**:
- Test flag parsing
- Test text processing operations

---

### If Building Text Processor:

**Functions to implement today**:
1. `readFile(filename string) string` — Read file content
2. `countLines(text string) int` — Count lines
3. `countWords(text string) int` — Count words

**Test focus**:
- Test counting functions with various inputs

---

## Testing Patterns to Use

**Pattern 1: Table-Driven Tests** (recommended)
```go
func TestCountWords(t *testing.T) {
    tests := []struct {
        name string
        input string
        want int
    }{
        {"single word", "hello", 1},
        {"multiple words", "hello world", 2},
        {"empty", "", 0},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := CountWords(tt.input)
            if got != tt.want {
                t.Errorf("got %d, want %d", got, tt.want)
            }
        })
    }
}
```

**Pattern 2: Simple Tests** (quick, less code)
```go
func TestCountWords(t *testing.T) {
    got := CountWords("hello world")
    want := 2
    if got != want {
        t.Errorf("got %d, want %d", got, want)
    }
}
```

---

## Git Workflow for Today

**Commits should show progression**:
```
day-4: [project] implement parseCSV function
day-4: [project] add tests for parseCSV
day-4: [project] implement askQuestion function
day-4: [project] gofmt and go vet cleanup
```

**Before each commit**:
1. Run `go test -v` — should pass
2. Run `gofmt -w .` — should format
3. Run `go vet ./...` — should pass

---

## Code Quality Checklist

Before committing:
- [ ] Code compiles: `go build ./...`
- [ ] All tests pass: `go test -v`
- [ ] Formatted: `gofmt -w .`
- [ ] No vet warnings: `go vet ./...`
- [ ] Variable names are clear (no single letters except `i`, `j`, `t`)
- [ ] Functions have purpose and do one thing
- [ ] No println debugging (use t.Log() in tests)

---

## Common Pitfalls to Avoid

❌ **Don't**: Try to implement everything at once
✅ **Do**: Implement one function at a time, test it, then move on

❌ **Don't**: Write 20 lines of code without testing
✅ **Do**: Write 5-10 lines, test it, then add more

❌ **Don't**: Ignore test failures
✅ **Do**: Fix failing tests immediately — they tell you what's wrong

❌ **Don't**: Worry about perfection
✅ **Do**: Get something working first, polish later

---

## End of Day Checklist

- [ ] Core functionality partially implemented
- [ ] At least 1 function working and tested
- [ ] Code passes gofmt and go vet
- [ ] Tests written and passing
- [ ] Meaningful commits made
- [ ] Time logged
- [ ] Tomorrow's tasks planned

---

## End of Day Reflection

**What I built today**:
- 

**What's working**:
- 

**What's broken or incomplete**:
- 

**Tests passing**: ___ / ___ 

**What I learned**:
- 

**Time spent**: ___ min / 75-90 min goal

**Tomorrow (Day 5)**:
- Implement remaining functions
- Add more tests
- Polish and refactor

---

## Debugging Tips

**If tests fail**:
1. Read the error message carefully
2. Add `t.Log()` statements to see variable values
3. Run specific test: `go test -v -run TestFunctionName`
4. Try the function manually in a small test program

**If code doesn't compile**:
1. Read the error — Go is usually very clear
2. Check syntax: function signature, types, parentheses
3. Use `go vet` to catch issues early

**If you're stuck**:
1. Comment out the broken part
2. See if the rest works
3. Then tackle one small piece of the broken part
4. Ask for help if truly stuck

---

## Phase 1 Progress

**Hours logged so far**:
- Day 1 (Apr 1): ___ hours
- Day 2 (Apr 2): ___ hours
- Day 3 (Apr 3): ___ hours
- Day 4 (Apr 4): ___ hours (TODAY)
- **Week 2 Total**: ___ / 5-6 hours

**Toward Phase 1 Goal**: ___ / 60-80 hours

---

## Tomorrow's Preview (Day 5 - Saturday)

**Day 5 (April 5)**: 
- Complete remaining project functionality
- Add comprehensive tests
- Code review and refactoring
- Polish and quality checks

**What to have ready**:
- Most functions from Day 4 working and tested
- Clear idea of what's left to implement
- Any edge cases you've discovered

