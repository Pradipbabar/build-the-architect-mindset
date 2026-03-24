# Day 2 — 24 March 2026

**Phase**: Phase 1 / Module 1.1 — Go Fundamentals  
**Focus**: Tour of Go — Basics Section  
**Time Target**: 2 hours  

---

## Today's Single Goal

Complete the **Basics** section of the Tour of Go.

> https://go.dev/tour/basics/

Do not move to other sections today. Depth over breadth.

---

## What You Will Cover

The Basics section covers these topics, in order:

1. Packages
2. Imports
3. Exported names
4. Functions (including multiple return values)
5. Variables (`var`, `:=`)
6. Basic types (`int`, `float64`, `bool`, `string`)
7. Zero values
8. Type conversions
9. Constants
10. `for` loops (the only loop in Go)
11. `if` and `if` with a short statement
12. `switch`
13. `defer`
14. Pointers
15. Structs
16. Arrays
17. Slices
18. Maps
19. Function values and closures

---

## How to Work Through It

### The Rule
**Do not copy-paste.** Type every single example yourself. This is non-negotiable.

### For each topic
1. Read the explanation on the page
2. Type the code in the editor (do not paste)
3. Run it — see the output
4. Ask yourself: *"Can I explain what this does in one sentence?"*
5. If no → re-read. If yes → move on.
6. Write your confusion in the Notes section below

### When you are stuck
- Stay stuck for 5 minutes first — think it through
- Re-read the Tour page explanation
- Check https://go.dev/doc/effective_go for more context
- Write down the confusion — do not skip it

---

## Code to Write Today

After finishing the Tour section, write this small program **from scratch** — no reference, from memory.

**File**: `projects/phase-1/00-day2-basics/basics_practice.go`

```
Write a Go program that does the following:

1. Declares a struct called Person with fields: Name (string), Age (int)
2. Creates two Person values
3. Uses a for loop to print each person's name and age
4. Uses an if statement to print "Senior" if age > 40, else "Junior"
5. Uses a map to count how many times each label (Senior/Junior) appears
6. Prints the final count

No functions other than main. No imports other than "fmt".
```

This tests: structs, loops, if/else, maps — all from today's Basics section.

---

## Notes (Fill These In As You Go)

### Things that made sense immediately

- Defining data types
- Using the built-in `make` function
- Writing and understanding a `for` loop

### Things that confused me

- Adding struct data into a slice
- I assumed slices/arrays could only store default data types like `int`, `string`, or `bool`, not structs

### Questions I want to research later

- How to improve my problem-understanding capability
- How slices work internally and what types they can store

### One thing I want to remember from today

- Start thinking more deeply before jumping to conclusions


---

## Time Log

| Session | Start | End | Duration | What I worked on |
|---------|-------|-----|----------|-----------------|
| 1 | | | | Tour of Go — Packages to Constants |
| 2 | | | | Tour of Go — for/if/switch/defer to Closures |
| 3 | | | | basics_practice.go — writing from scratch |

**Total today**: ___ hours  
**Running total**: ___ hours / 1500–2000 goal

---

## End of Day Checklist

- [x] Completed all topics in the Tour of Go Basics section
- [x] Typed every example (did not paste)
- [x] Wrote `basics_practice.go` from memory
- [x] Program runs without errors (`go run basics_practice.go`)
- [x] Filled in the Notes section above
- [x] Committed code to GitHub with a clear message
- [x] Logged time honestly

---

## Git Commit Message for Today

```
day-2: tour of go basics section complete

- worked through all 19 topics in basics section
- wrote basics_practice.go from scratch (structs, maps, loops, if/else)
- noted confusions around [fill in what confused you]
```

---

## Day 3 Preview

Tomorrow you move to the **Methods and Interfaces** section of Tour of Go — the part most beginners skip and then struggle with for months.

Come to it rested. It requires focused attention.

---

**Commitment check**: Did I write the code myself today? Yes / No  
**Honest hours logged**: ___