# Day 6 — 28 March 2026 (Saturday)

**Phase**: Phase 1 / Module 1.1 — Go Fundamentals  
**Focus**: Continue Learning - Build Additional Programs  
**Time Target**: 1.5 hours  
**Status**: ✅ Completed
**Rule**: Be honest. Do not inflate hours or progress. This log is for you.

---

## Context: Why Saturday Matters

You've completed the Tour of Go and implemented several core concepts (basics, methods/interfaces, concurrency). Saturday is about reinforcing what you learned by building another real program.

This is not review time — it's practice time. The best way to solidify concepts is to apply them in a different context.

---

## Today's Goals

1. Build a real program (temperature converter, calculator, or number game)
2. Use at least 2-3 concepts from the past 4 days
3. Test your program thoroughly
4. Document what you built
5. Identify any gaps in understanding

---

## Part 1: Choose Your Program (10 min)

Pick ONE of these (or propose your own):

### Option A: Temperature Converter
**What it does**:
- Asks user for temperature and unit (Celsius or Fahrenheit)
- Converts to the other unit
- Shows result with explanation

**Concepts**: Functions, variables, conditionals, user input

**Example Flow**:
```
Enter temperature: 32
Enter unit (C/F): F
Result: 32°F = 0°C

Explanation: Water freezes at this temperature.
```

### Option B: Simple Calculator
**What it does**:
- Takes two numbers and an operation (+, -, *, /)
- Performs calculation
- Shows result with validation

**Concepts**: Functions, error handling, type conversion, conditionals

### Option C: Number Guessing Game
**What it does**:
- Computer picks a random number (1-100)
- Player guesses, gets feedback (too high/low)
- Tracks number of attempts
- Declares winner with congratulations

**Concepts**: Loops, conditionals, random number generation, user input

### Pick Your Program
**I choose**: _________Number Guessing Game_______________

---

## Part 2: Plan Your Code (15 min)

Before you write code, write the logic:

```
1. What are the inputs? (user types something)
2. What are the outputs? (what does user see?)
3. What functions do I need?
4. What logic goes in each function?
5. Where might errors happen?
```

Write your plan here or in a separate notes file.

---

## Part 3: Build It (50 min)

### Create the file
```bash
cd /home/pradip/repos/build-the-architect-mindset/projects/phase-1/00-day2-basics
# or appropriate day folder
touch program-name.go
```

### Write your code
- Start with package main and imports
- Write your functions
- Write main()
- Add comments

### Run and test
```bash
go run program-name.go
```

Test with:
- Normal inputs
- Edge cases (0, negative numbers, empty input)
- Wrong types (letters when numbers expected)

### Debug as needed
Did something break? Good. That's how you learn.

---

## Part 4: Document What You Built (20 min)

Create a file: `projects/phase-1/PROGRAMS-BUILT.md`

Add an entry:

```markdown
## Program: [Name]
**Date**: March 28, 2026
**Time**: ___ minutes
**Concepts Used**: 
- List 3-5 concepts you used

**How It Works**:
Brief description (5 sentances max)

**Best Part**: What felt most natural?
**Hardest Part**: What was tricky?

**Lessons Learned**:
- Bullet 1
- Bullet 2
```

---

## Part 5: Reflect (15 min)

Answer in your journal thoughts or in console:

1. **Did this program feel easier than Day 1**? Why or why not?
2. **Which concept did you use most**? (functions, loops, conditionals, etc.)
3. **What would you add** if you had 2 more hours?
4. **Are you ready** for Day 7 (weekly review)?

---

## Notes

- Do not copy-paste code from the internet
- Do not skip testing
- Do not rush — take your time understanding each line
- Questions are welcome — note them for Monday

---

**Time Logged**: ___ hours

**Status**: Ready for Sunday review
