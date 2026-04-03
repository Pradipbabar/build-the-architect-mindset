# Week 7: Real-World Project #2 — Advanced Application Build

**Week**: Week 7 of Phase 1 — Module 1 (Advanced Projects)
**Dates**: May 4 – May 10, 2026
**Total Hours Goal**: 8 hours (larger, more complex project)
**Roadmap Reference**: `roadmap/PHASE-1-MODULES.md` → Module 1, Capstone Project

---

## Week 7 Goals

**Primary Goal**: Build a more sophisticated application that demonstrates mastery of Go fundamentals and production practices.

**Secondary Goals**:
- Build a project with multiple components and interfaces
- Apply advanced concurrency patterns
- Implement proper testing strategy
- Demonstrate package design skills
- Create a portfolio-worthy project
- Handle real-world complexity

**Hours Target**: 8 hours TOTAL for the week

---

## OPTION A: Full Week Advanced Project Development

**Monday, May 4 — Project Design & Architecture (1.5 hours)**
- [ ] Choose complex project (see ideas below)
- [ ] Design multi-package architecture
- [ ] Draw package dependency diagram
- [ ] Write interface contracts
- [ ] Plan testing strategy
- [ ] Commit: `"week-7: architecture and design"`

**Tuesday-Wednesday, May 5-6 — Foundation & Core Features (2.5 hours)**
- [ ] Implement base types and interfaces
- [ ] Build data structures and models
- [ ] Create main business logic components
- [ ] Day 1: Core data structures and interfaces (1.5h)
- [ ] Day 2: Main functionality (1h)

**Thursday, May 7 — Integration & Advanced Features (2 hours)**
- [ ] Integrate all packages
- [ ] Implement concurrency patterns (if applicable)
- [ ] Add advanced features (caching, validation, etc.)
- [ ] Error handling throughout

**Friday, May 8 — Comprehensive Testing (1.5 hours)**
- [ ] Unit tests (all packages): 45 min
- [ ] Integration tests: 30 min
- [ ] Benchmarks for performance-critical code: 15 min

**Saturday, May 9 — Documentation & Quality (1 hour)**
- [ ] README.md with architecture overview
- [ ] Usage examples and API documentation
- [ ] Code comments and godoc
- [ ] Fix any vet/fmt issues

**Sunday, May 10 — Polish & Final Review (30 min)**
- [ ] `gofmt`, `go vet`, `go test -race`
- [ ] Code review and refactoring
- [ ] Final cleanup
- [ ] Commit: `"week-7: advanced project complete"`

---

## OPTION B: Lighter Pace

**Monday-Tuesday, May 4-5**
- [ ] Design + foundation: 2.5 hours

**Wednesday-Thursday, May 6-7**
- [ ] Core features + integration: 2.5 hours

**Friday-Saturday, May 8-9**
- [ ] Advanced features + testing: 2.5 hours

**Sunday, May 10**
- [ ] Polish + finalize: 30 min

---

## OPTION C: Flexible

**Just ensure by end of week**:
- [ ] Complex multi-package project designed and started
- [ ] 2+ packages created with clear interfaces
- [ ] 1000+ lines of production code written
- [ ] 5+ integration tests written and passing
- [ ] 60%+ code coverage
- [ ] Concurrency handled safely (passes -race)
- [ ] All code quality checks passing
- [ ] Architecture documented in README
- [ ] Committed to GitHub with meaningful history

---

## Advanced Project Ideas (Pick ONE)

### Project 1: REST API — Expense Tracker ⭐ (Recommended)
**Complexity Level**: Advanced
**What it does**:
- Full REST API for personal expense tracking
- Users, expenses, categories, budgets
- Filtering and aggregation (by date, category, etc.)
- Budget tracking and alerts
- Currency conversion (optional: external API)
- JWT authentication (optional)

**Technologies**:
- `net/http` (or use routing library like `chi` or `gorilla/mux`)
- Database abstraction layer
- Concurrency for concurrent requests
- Error wrapping and validation
- Interface-based design for extensibility

**Package structure**:
```
expense-tracker/
├── main.go
├── config/
│   └── config.go
├── handler/
│   ├── user.go
│   ├── expense.go
│   ├── category.go
│   └── middleware.go (auth, error handling)
├── service/
│   ├── user_service.go (business logic)
│   ├── expense_service.go
│   └── budget_service.go
├── storage/
│   ├── interface.go (Storage contract)
│   ├── memory.go (in-memory impl)
│   └── file.go (file-based impl)
├── model/
│   └── entities.go (all structs)
├── util/
│   ├── errors.go (error types)
│   └── validation.go
├── http_test.go
├── service_test.go
└── README.md
```

**Key advanced patterns**:
- Service layer for business logic (separation of concerns)
- Middleware for cross-cutting concerns
- Proper dependency injection
- Transaction-like operations
- Comprehensive error handling

**Tests needed**:
- User registration and login
- Expense CRUD operations
- Budget calculations and alerts
- Concurrent requests don't corrupt data
- All error cases handled properly
- Validation logic

**Estimated time**: 7-8 hours

### Project 2: CLI Tool — Mini Database with Query Engine
**Complexity Level**: Advanced
**What it does**:
- In-memory database with simple SQL-like query support
- Create tables, insert rows, query data
- WHERE clause filtering
- JOIN operations (simple)
- Aggregation functions (COUNT, SUM, AVG)
- Persistence to JSON

**Technologies**:
- Query parsing and execution
- Data structure design (tables, rows, indexes)
- Concurrent query handling
- Error handling and validation
- Reflection for dynamic column typing

**Package structure**:
```
mini-db/
├── main.go
├── parser/
│   ├── lexer.go (tokenize SQL)
│   ├── parser.go (parse SQL)
│   └── ast.go (abstract syntax tree)
├── engine/
│   ├── engine.go (query execution)
│   ├── table.go (table management)
│   └── query.go (query handler)
├── datatype/
│   ├── types.go (column types)
│   └── value.go (runtime values)
├── storage/
│   ├── interface.go
│   └── json_store.go
├── repl/
│   └── repl.go (read-eval-print loop)
├── engine_test.go
└── README.md
```

**Key advanced patterns**:
- Parser design (lexer + parser)
- Query planning and execution
- Table and index management
- Type inference and coercion
- Memory management for larger datasets

**Tests needed**:
- Query parsing
- SELECT operations with WHERE
- JOIN operations
- Aggregation functions
- INSERT and UPDATE operations
- Error cases (malformed queries, missing tables)
- Persistence/loading

**Estimated time**: 8+ hours

### Project 3: CLI Tool — Task Scheduling System
**Complexity Level**: Advanced
**What it does**:
- Schedule tasks to run at specific times
- Cron-like scheduling (simple)
- Task status tracking
- Job history and logs
- Task dependencies
- Concurrent task execution with coordination

**Technologies**:
- Time-based scheduling
- Goroutines and channel coordination
- Distributed execution patterns
- Status tracking and persistence
- Concurrent safety with sync.Mutex or channels

**Package structure**:
```
task-scheduler/
├── main.go
├── scheduler/
│   └── scheduler.go (main scheduler)
├── task/
│   ├── interface.go (Runnable interface)
│   ├── task.go (Task definition)
│   └── executor.go (Task executor)
├── storage/
│   ├── state.go (Task state tracking)
│   └── history.go (Execution history)
├── time/
│   └── parser.go (Cron expression parser)
├── concurrent/
│   ├── pool.go (Worker pool)
│   └── coordinator.go (Dependency coordination)
├── main_test.go
└── README.md
```

**Key advanced patterns**:
- Worker pool pattern
- Dependency resolution
- Time-based synchronization
- Proper shutdown and cleanup
- Concurrent safety at scale

**Tests needed**:
- Schedule parsing
- Task execution at correct time
- Task dependencies
- Worker pool coordination
- Failure and retry logic
- Concurrent execution safety
- State persistence

**Estimated time**: 8+ hours

### Project 4: Data Processing — Stream Analytics Tool
**Complexity Level**: Advanced
**What it does**:
- Process event streams from files/pipes
- Real-time statistics (moving avg, percentiles)
- Filtering and transformation
- Aggregation windows (time-based or count-based)
- Multiple output formats
- High-performance concurrent processing

**Technologies**:
- Concurrent stream processing
- Goroutines + channels (pipeline pattern)
- Buffering and backpressure handling
- Performance optimization
- Memory efficiency for large streams

**Key advanced patterns**:
- Pipeline pattern with goroutines
- Channel-based fan-out/fan-in
- Buffering strategies
- Backpressure handling
- Clean shutdown procedures

**Estimated time**: 8+ hours

---

## Production-Ready Checklist

Before declaring the project "complete":

**Architecture**:
- [ ] 3+ packages with clear responsibilities
- [ ] 2+ interfaces defined and used
- [ ] No circular dependencies
- [ ] Clear separation of concerns

**Error Handling**:
- [ ] Custom error types for domain errors
- [ ] Error wrapping with context
- [ ] Graceful degradation where applicable
- [ ] All error paths tested

**Testing**:
- [ ] 60%+ code coverage
- [ ] Unit tests for business logic
- [ ] Integration tests for components
- [ ] Benchmarks for performance-critical code
- [ ] Table-driven tests used appropriately
- [ ] All tests pass with `-race` flag

**Code Quality**:
- [ ] `gofmt -w .` passes
- [ ] `go vet ./...` passes
- [ ] `go test -race ./...` passes
- [ ] No unused variables/imports
- [ ] Meaningful variable/function names
- [ ] Comments on exported functions

**Documentation**:
- [ ] README.md with architecture overview
- [ ] Usage examples in README
- [ ] API documentation for key packages
- [ ] Godoc comments on public functions
- [ ] How to run tests and linters

**Version Control**:
- [ ] Meaningful commit messages
- [ ] Commits follow progression of development
- [ ] Clean git history (no WIP commits)
- [ ] Branch structure if applicable

---

## Time Breakdown

| Task | Target | Actual |
|---|---|---|
| Architecture & design | 1.5 hours | ___ |
| Core implementation | 2.5 hours | ___ |
| Integration & features | 2 hours | ___ |
| Testing | 1.5 hours | ___ |
| Documentation + polish | 1 hour | ___ |
| **Total** | **~8 hours** | ___ |

---

## End of Week Checklist

- [ ] Advanced project chosen and designed
- [ ] 3+ package structure with interfaces
- [ ] All core functionality implemented
- [ ] Comprehensive testing (60%+ coverage)
- [ ] Error handling with wrapping throughout
- [ ] All code passes gofmt, go vet, go test -race
- [ ] README with architecture documentation
- [ ] Examples and usage instructions included
- [ ] Meaningful git commit history
- [ ] Total hours logged: ___ / 8

---

## Actual Time Spent

**Total hours: ___**

**Project chosen**: ___

**Architecture overview**:

**Packages created**:

**Key interfaces designed**:

**Challenges encountered**:

**Solutions implemented**:

**What I learned**:

**Code quality metrics**:
- Coverage: ___%
- Package count: ___
- Lines of code: ___
- Test count: ___

---

## Phase 1 Completion Status

**Weeks 1-7 Progress**:
- ✅ TGPL Chapters 1-12 mastered
- ✅ Advanced concepts practiced (error wrapping, interfaces, concurrency)
- ✅ Mini project completed (config loader)
- ✅ 2 substantial real-world projects built
- ✅ Production-level code quality demonstrated
- ✅ Multi-package architecture designed and implemented
- ✅ Comprehensive testing strategy applied
- ✅ ~70 hours of 60-80 hour goal complete

---

## What's Next?

**Phase 1 is essentially complete!** You now have:
- Solid Go fundamentals from TGPL
- Production code experience
- Portfolio projects on GitHub
- Testing discipline
- Code quality standards

**Optional for remaining Phase 1 time**:
- Week 8-9: gRPC and protocol buffers intro
- Week 8-9: REST API with proper middleware and auth
- Week 8-9: Performance optimization and profiling
- Week 8-9: Contributions to open-source Go projects

**Next Phase (Phase 2)**:
- Ready to move to infrastructure/cloud architecture
- Can build distributed systems with confidence
- Can contribute to production Go codebases
