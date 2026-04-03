# Week 5: TGPL Chapters 10-12 + Advanced Concepts & Mini Project

**Week**: Week 5 of Phase 1 — Module 1 (Go Fundamentals Advanced)
**Dates**: April 20 – April 26, 2026
**Total Hours Goal**: 6-7 hours (final TGPL coverage + advanced concepts)
**Roadmap Reference**: `roadmap/PHASE-1-MODULES.md` → Module 1, Completing Fundamentals

---

## Week 5 Goals

**Primary Goal**: Complete TGPL (Chapters 10-12). Introduce advanced Go concepts.

**Secondary Goals**:
- Understand package design and best practices
- Master testing patterns and benchmarking
- Learn reflection basics (advanced)
- Start thinking about real-world Go architectures

**Hours Target**: 6-7 hours TOTAL for the week

---

## OPTION A: Chapter-by-Chapter + Mini Project

**Monday, April 20 — Chapter 10: Packages (1.5 hours)**
- [ ] Read TGPL Chapter 10 — Packages & Go Tools
- [ ] Type 3-4 examples (package organization, imports, visibility)
- [ ] Create: `notes/books/tgpl-chapter-10.md`
- [ ] Key topics: package structure, init(), unexported vs exported
- [ ] Commit: `"week-5: tgpl chapter 10 packages notes"`

**Tuesday, April 21 — Chapter 11: Testing (1 hour)**
- [ ] Read TGPL Chapter 11 — Testing
- [ ] Type 2-3 examples (unit tests, table-driven, benchmarks)
- [ ] Create: `notes/books/tgpl-chapter-11.md`
- [ ] Key topics: testing patterns, benchmarks, examples
- [ ] Commit: `"week-5: tgpl chapter 11 testing notes"`

**Wednesday, April 22 — Chapter 12: Reflection (1 hour)**
- [ ] Read TGPL Chapter 12 — Reflection
- [ ] Type 2-3 examples (reflect package basics)
- [ ] Create: `notes/books/tgpl-chapter-12.md`
- [ ] Key topics: reflection basics, type inspection, marshal/unmarshal
- [ ] Commit: `"week-5: tgpl chapter 12 reflection notes"`

**Thursday-Friday, April 23-24 — Advanced Concepts Intro (1.5 hours)**
- [ ] Study advanced patterns:
  - [ ] Error wrapping (errors.Is, errors.As): 20 min
  - [ ] Working with interfaces: 20 min
  - [ ] Performance optimization basics: 20 min
- [ ] Write 2-3 small examples: 30 min
- [ ] Create: `notes/concepts/Go_Advanced/error-handling-interfaces.md`
- [ ] Commit: `"week-5: advanced go concepts intro"`

**Saturday-Sunday, April 25-26 — Mini Project: Config Loader (1.5 hours)**
- [ ] Build a reusable config package:
  - [ ] Parses environment + flags
  - [ ] Uses reflection for validation
  - [ ] Well-tested with benchmarks
  - [ ] Properly packaged (reusable)
- [ ] Implementation: 1 hour
- [ ] Tests + benchmarks: 30 min
- [ ] Commit: `"week-5: config-loader package complete"`

---

## OPTION B: Chapters + Self-Study

**Monday-Wednesday, April 20-22**
- [ ] Read all 3 chapters: 2 hours
- [ ] Type examples: 1 hour
- [ ] Create notes: 30 min

**Thursday-Friday, April 23-24**
- [ ] Advanced concepts study: 1.5 hours

**Saturday-Sunday, April 25-26**
- [ ] Mini project (config loader): 1.5 hours

---

## OPTION C: Flexible

**Just ensure by end of week**:
- [ ] Chapters 10-12 fully read
- [ ] Notes created for all 3 chapters
- [ ] Advanced concepts introduction studied
- [ ] Mini project planned and started
- [ ] Code passes gofmt and go vet

---

## Key Topics This Week

### Chapter 10: Packages
- Package organization and structure
- `init()` function (initialization)
- Exported vs unexported (UPPERCASE vs lowercase)
- Documentation patterns
- Package visibility

### Chapter 11: Testing
- Unit tests (`*testing.T`)
- Table-driven tests
- Benchmarking (`*testing.B`)
- Examples (`TestExample`)
- Test coverage

### Chapter 12: Reflection
- `reflect.Type` and `reflect.Value`
- Type inspection at runtime
- Marshaling/unmarshaling with tags
- When to use reflection (and when not to!)

### Advanced Concepts to Explore
- **Error Wrapping**: Using `fmt.Errorf("%w", err)` for error chains
- **Interfaces**: Design by contract, composition over inheritance
- **Performance**: Profiling tools (pprof), benchmarking best practices
- **Goroutine Leaks**: Avoiding common concurrency mistakes
- **Context Package**: Request cancellation and timeouts

---

## Time Breakdown

| Task | Target | Actual |
|---|---|---|
| Chapter 10 reading + notes | 1.5 hours | ___ |
| Chapter 11 reading + notes | 1 hour | ___ |
| Chapter 12 reading + notes | 1 hour | ___ |
| Advanced concepts study | 1.5 hours | ___ |
| Mini project implementation | 1 hour | ___ |
| Quality checks + commits | 30 min | ___ |
| **Total** | **~6.5 hours** | ___ |

---

## End of Week Checklist

- [ ] Chapters 10-12 read completely
- [ ] Notes created for all 3 chapters
- [ ] Advanced concepts explored (error wrapping, interfaces, performance)
- [ ] Mini project (config loader) planned and implemented
- [ ] Mini project has tests + benchmarks
- [ ] All code passes gofmt and go vet
- [ ] All work committed to GitHub
- [ ] Total hours logged: ___ / 6-7

---

## Actual Time Spent

**Total hours: ___**

**Chapters that made sense**:

**Chapters that were confusing**:

**Advanced concepts I want to explore more**:

**Mini project progress**:

---

## Progress: TGPL Complete! 🎉

**After Week 5, you will have**:
- ✅ Read ALL 12 chapters of TGPL
- ✅ Created notes for every chapter
- ✅ Typed hundreds of examples
- ✅ Practiced concurrency patterns
- ✅ Tested your code thoroughly
- ✅ Started learning advanced concepts
- ✅ Built 3-4 projects (simple + mini project)

**You're now ready for real-world projects!**
