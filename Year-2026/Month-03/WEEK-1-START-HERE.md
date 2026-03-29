# Week 1: Start Here - Your Principal Engineer Journey Action Plan

**Welcome to Day 1 of your Principal Engineer journey with Architect Mindset!**

This roadmap now includes leadership and strategic capabilities for Principal Engineer roles.

This document tells you exactly what to do in Week 1. Follow this step by step.

---

## Overview: What Week 1 Accomplishes

By the end of Week 1, you will have:
- Complete folder structure set up ✅ (DONE!)
- All necessary tools installed
- Books ordered/downloaded
- Baseline assessment of your Go knowledge
- First learning session completed
- Commitment declaration signed

**Time Required**: 10 hours total across 7 days

---

## Day 1 (Today): Setup & Commitment (2 hours)

### Step 1: Review Your Folder Structure (15 minutes)
You now have this structure created:
```
/learning-architect/
├── README.md                    ← Your main dashboard
├── COMMITMENT.md                ← Sign this today!
├── architect-career-roadmap.md  ← Full roadmap reference
├── WEEK-1-START-HERE.md        ← You are here
├── roadmap/                     ← Detailed phase plans
├── notes/                       ← Your learning notes
├── projects/                    ← Your code projects
├── system-designs/              ← Phase 2 designs
├── architecture-docs/           ← Phase 3 work
├── blog/                        ← Your blog posts
├── resources/                   ← Books, courses, tools
└── journal/                     ← Weekly journals
```

**Action**:
- [ ] Open the folder in your file explorer
- [ ] Open `README.md` and familiarize yourself with it
- [ ] Bookmark this folder location

---

### Step 2: Read and Sign COMMITMENT.md (45 minutes)

**Action**:
- [ ] Open `COMMITMENT.md`
- [ ] Read it completely (don't skip)
- [ ] Answer all questions honestly
- [ ] Fill in all blank sections
- [ ] Print or digitally sign it
- [ ] Save it

**This is critical.** If you can't commit, stop here and reconsider your goals.

---

### Step 3: Install Essential Tools (1 hour)

#### Install Go
- [ ] Download Go: https://go.dev/dl/
- [ ] Install latest stable version (1.22+)
- [ ] Verify: Open terminal and run `go version`
- [ ] Should see: `go version go1.22.x ...`

#### Install/Verify Git
- [ ] Check if installed: `git --version`
- [ ] If not installed: https://git-scm.com/downloads
- [ ] Configure git:
  ```bash
  git config --global user.name "Your Name"
  git config --global user.email "your.email@example.com"
  ```

#### Install VS Code (if you don't have it)
- [ ] Download: https://code.visualstudio.com/
- [ ] Install Go extension: Search "Go" in extensions, install official one
- [ ] Open a folder and verify Go extension works

#### Install Docker (for later, but do it now)
- [ ] Download: https://www.docker.com/products/docker-desktop
- [ ] Install Docker Desktop
- [ ] Verify: `docker --version`

#### Set Up GitHub
- [ ] Create account if you don't have one: https://github.com
- [ ] Create repository: `learning-architect-projects`
- [ ] Make it public (for your portfolio later)
- [ ] Clone it locally:
  ```bash
  cd /Users/pradip.babar/Learnings/learning-architect/projects
  git clone https://github.com/YOUR-USERNAME/learning-architect-projects.git
  ```

---

## Day 2: Acquire Learning Resources (1.5 hours)

### Books to Get

#### Option 1: Buy (Recommended if you have budget)
- [ ] "The Go Programming Language" by Donovan & Kernighan (~$40)
  - Amazon, or digital version

#### Option 2: Library
- [ ] Check local library for TGPL
- [ ] Request if not available

#### Option 3: Alternative Free Resources
If budget is an issue, start with:
- [ ] Bookmark "A Tour of Go": https://go.dev/tour/
- [ ] Bookmark "Go by Example": https://gobyexample.com/
- [ ] Bookmark "Effective Go": https://go.dev/doc/effective_go
- [ ] Save money for DDIA (most critical book, ~$50)

### Update Books Tracking
- [ ] Open `resources/books.md`
- [ ] Update status of books you acquired
- [ ] Set target date to start reading

---

## Day 3: Baseline Assessment (1.5 hours)

### Complete Tour of Go (Module 1 only)
- [x] Go to: https://go.dev/tour/
- [x] Complete "Basics" section only
- [x] Don't worry about understanding everything
- [x] Note what confuses you

### Document Your Current Level
- [x] Create file: `notes/concepts/go-baseline-assessment.md`
- [x] Write down:
  - What I already know about Go: ___
  - What confused me in Tour of Go: ___
  - What I need to learn first: ___
  - My current confidence level (1-10): ___

### Set Up Your First Project Folder
```bash
cd /Users/pradip.babar/Learnings/learning-architect/projects/phase-1
mkdir 00-mini-projects
cd 00-mini-projects
git init
```
- [x] Set up project folder

---

## Day 4: First Learning Session (2 hours)

### Read TGPL Chapter 1 (or Tour of Go Basics review)

If you have TGPL:
- [x] Read Chapter 1: Tutorial
- [x] Type out (don't copy-paste) the examples
- [x] Run each example
- [x] Take notes in `notes/books/tgpl-chapter-1.md`

If you don't have TGPL yet:
- [x] Complete entire Tour of Go
- [x] Do all exercises
- [x] Take notes

### Build Your First Tiny Program
- [x] Create file: `projects/phase-1/00-mini-projects/hello-world.go`
- [x] Write a program that:
  - Takes your name as input
  - Prints a greeting
  - Asks for your age
  - Prints your age in dog years (age * 7)

Example starter:
```go
package main

import "fmt"

func main() {
    // Your code here
}
```

- [x] Run it: `go run hello-world.go`
- [x] Test with different inputs
- [x] Add comments explaining what each part does

---

## Day 5: Create Your First Journal Entry (1 hour)

### Copy Template and Fill It Out
```bash
cp journal/JOURNAL-TEMPLATE.md journal/2026-03-week-12.md
```

- [x] Copy template and fill it out
- [x] Open `journal/2026-03-week-12.md`
- [x] Fill in Days 1-5 (what you did so far)
- [x] Be specific about time spent each day
- [x] Note what you learned
- [x] Note what you struggled with
- [x] Note questions you have

### Update README.md
- [x] Open `README.md`
- [x] Update "This Week" section with your actual progress
- [x] Update time tracking
- [x] Check off completed goals

---

## Day 6-7: Continue Learning + Weekly Review (2 hours)

### Saturday: Continue Learning (1.5 hours)
- [ ] Continue TGPL Chapter 1 or Tour of Go
- [ ] Build another tiny program (your choice):
  - Temperature converter (Celsius ↔ Fahrenheit)
  - Simple calculator (add, subtract, multiply, divide)
  - Number guessing game

### Sunday: Weekly Review (0.5 hours)
- [ ] Complete your journal entry for Week 1
- [ ] Fill out weekly summary section
- [ ] Update README.md with progress
- [ ] Calculate total hours spent this week
- [ ] Did you hit 10 hours? If not, why not?

### Plan Week 2
- [ ] Copy `roadmap/weekly-goals.md` to a new file
- [ ] Plan Week 2 goals (continue TGPL, more mini projects)
- [ ] Schedule your 10 hours for next week

---

## Week 1 Completion Checklist

By end of Week 1, you should have:

### Setup & Tools
- [ ] Folder structure created and understood
- [ ] Go installed and verified
- [ ] Git installed and configured
- [ ] GitHub account and repository created
- [ ] VS Code with Go extension set up
- [ ] Docker installed

### Commitment
- [ ] Read COMMITMENT.md completely
- [ ] Filled out all sections
- [ ] Signed and dated it
- [ ] Shared with accountability partner (if applicable)

### Learning
- [x] Completed baseline Go assessment (Tour of Go basics or TGPL Ch1)
- [x] Built 1-2 tiny programs
- [ ] Created notes on what you learned
- [ ] Identified what confused you

### Organization
- [x] First journal entry created
- [ ] README.md updated with your status
- [ ] Books status updated in resources/books.md
- [ ] Week 2 planned

### Time
- [x] Logged approximately 10 hours
- [x] Tracked time for each day
- [ ] Identified best time slots for learning

---

## If You're Behind Schedule

**Don't panic.** Adjust, don't abandon.

### Couldn't finish everything?
- That's okay for Week 1 (setup takes time)
- Carry over to Week 2
- The important thing: you started

### Couldn't get 10 hours?
- How many hours did you get? ___
- What blocked you? ___
- How will you prevent this next week? ___
- Adjust schedule, don't quit

### Feeling overwhelmed?
- This is normal
- You don't need to understand everything immediately
- Focus on consistency, not perfection
- Small progress every day adds up

---

## Week 2 Preview

Next week you'll:
1. Continue TGPL Chapters 2-3 (or finish Tour of Go + start TGPL)
2. Build 2-3 more mini CLI programs
3. Start thinking about your first real project (Log Parser)
4. Join a Go community (Reddit or Discord)
5. Continue daily journaling

---

## Questions or Stuck?

### Resources When Stuck
1. **Go Documentation**: https://pkg.go.dev/
2. **Go Forum**: https://forum.golangbridge.org/
3. **Stack Overflow**: https://stackoverflow.com/questions/tagged/go
4. **Reddit**: r/golang

### Remember the 30-Minute Rule
- When stuck, try to solve it yourself for 30 minutes
- Use documentation, not AI
- If still stuck after 30 minutes, ask for help
- Document what you tried

---

## Motivation for Hard Days

Read this when you feel like quitting:

**"The difference between who you are and who you want to be is what you do."**

You've already done more than most people who say "I want to learn":
- You've read a brutally honest assessment
- You've created a complete learning structure
- You've made a commitment
- You've taken the first step

**Keep going.**

The first week is the hardest. Then it gets easier as it becomes a habit.

---

## Final Checklist Before Week 2

- [ ] All tools installed and working
- [ ] COMMITMENT.md signed
- [ ] At least 1 program written
- [ ] Journal entry for Week 1 complete
- [ ] Week 2 planned
- [ ] You're still committed to this journey

**If all checked, you're ready for Week 2!**

---

## Your Week 1 Reflection

Take 5 minutes to write your thoughts:

**How do I feel about this journey after Week 1?**


**What excited me this week?**


**What scared me this week?**


**Am I still committed? (Be honest)**


**My focus for Week 2:**


---

**Congratulations on completing Week 1! You've started. Now keep going.**

**Next**: Continue with Week 2 following the plan in `roadmap/phase-1-foundation.md`
