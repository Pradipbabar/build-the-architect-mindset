# Week 6: Real-World Project #1 — REST API or CLI Tool (Advanced)

**Week**: Week 6 of Phase 1 — Module 1 (Advanced Projects)
**Dates**: April 27 – May 3, 2026
**Total Hours Goal**: 7-8 hours (substantial production-like project)
**Roadmap Reference**: `roadmap/PHASE-1-MODULES.md` → Module 1, Real-World Projects

---

## Week 6 Goals

**Primary Goal**: Build a real-world, production-like Go project combining everything learned in Chapters 1-12.

**Secondary Goals**:
- Apply all TGPL knowledge (types, functions, methods, interfaces, concurrency, testing, reflection)
- Build proper package structure
- Implement comprehensive testing
- Handle errors gracefully with wrapping
- Focus on code quality and maintainability
- Think like a production engineer

**Hours Target**: 7-8 hours TOTAL for the week

---

## OPTION A: Full Week Project Development

**Monday, April 27 — Project Design & Architecture (1 hour)**
- [ ] Choose project (see ideas below)
- [ ] Design architecture (package structure, interfaces, data flow)
- [ ] Write pseudocode/comments
- [ ] Plan API or CLI interface
- [ ] Commit: `"week-6: project design documentation"`

**Tuesday-Wednesday, April 28-29 — Core Implementation (2.5 hours)**
- [ ] Build core data structures (proper types, interfaces)
- [ ] Implement main business logic
- [ ] Create proper package organization
- [ ] Day 1: Core functionality (1.5h)
- [ ] Day 2: Additional features (1h)

**Thursday, April 30 — Integration & Error Handling (1.5 hours)**
- [ ] Integrate components
- [ ] Comprehensive error handling with error wrapping
- [ ] Edge case management
- [ ] Configuration/flags support

**Friday, May 1 — Testing & Documentation (1.5 hours)**
- [ ] Unit tests (50%+ coverage minimum)
- [ ] Integration tests
- [ ] Benchmarks for critical paths
- [ ] README.md with usage examples
- [ ] Code comments and godoc

**Saturday, May 2 — Quality & Polish (1 hour)**
- [ ] `gofmt -w .` 
- [ ] `go vet ./...`
- [ ] `go test -race ./...`
- [ ] `go test -cover ./...`
- [ ] Fix all warnings/issues

**Sunday, May 3 — Final Review & Commit (30 min)**
- [ ] Code review of your own work
- [ ] Refactor unclear sections
- [ ] Final commit: `"week-6: real-world project complete"`
- [ ] Update project log

---

## OPTION B: Lighter Pace

**Monday-Tuesday, April 27-28**
- [ ] Design + core implementation: 2.5 hours

**Wednesday-Thursday, April 29-30**
- [ ] Continue implementation + error handling: 2.5 hours

**Friday-Sunday, May 1-3**
- [ ] Testing, packaging, quality: 2 hours

---

## OPTION C: Flexible

**Just ensure by end of week**:
- [ ] Substantial project implemented (200+ lines)
- [ ] Proper package structure (2+ packages)
- [ ] 3+ integration tests written
- [ ] 50%+ code coverage
- [ ] All code passes gofmt, go vet, go test -race
- [ ] README.md created with usage examples
- [ ] Committed to GitHub

---

## Real-World Project Ideas (Pick ONE)

### Project 1: REST API — Todo Service ⭐ (Recommended)
**What it does**:
- REST API for todo list management
- Create, read, update, delete (CRUD) operations
- Persistent storage (JSON file or in-memory)
- Proper HTTP status codes (201, 400, 404, etc.)
- Input validation and error handling
- Concurrent request handling

**Technologies**:
- `net/http` package (standard library)
- File I/O or struct-based storage
- Goroutines for concurrent requests handling
- Error wrapping with `fmt.Errorf("%w", ...)`
- Interfaces for storage abstraction

**Project structure**:
```
todo-api/
├── main.go
├── handler/
│   ├── todo.go (HTTP handlers)
│   └── response.go (JSON response helpers)
├── storage/
│   ├── interface.go (Storage interface definition)
│   ├── memory.go (In-memory implementation)
│   └── file.go (File-based implementation)
├── model/
│   └── todo.go (Todo struct definition)
├── main_test.go
├── storage/
│   └── storage_test.go
└── README.md
```

**Key patterns used**:
- Interface-based design for storage abstraction
- Error wrapping for context-aware error handling
- Goroutines for concurrent request handling
- Table-driven tests for HTTP handlers
- Proper HTTP routing and status codes

**Tests needed**:
- Create todo (valid input)
- Read todo (found and not found)
- Update todo (valid update, invalid ID)
- Delete todo
- List todos
- Error cases (invalid JSON, missing fields)
- Concurrent requests don't corrupt data

**Estimated time**: 5-7 hours implementation + tests

### Project 2: CLI Tool — Log Analyzer (Advanced)
**What it does**:
- Parse application log files
- Filter by level/date/keyword search
- Generate statistics and reports
- Multiple output formats (JSON, CSV, plaintext)
- Concurrent file processing for multiple files

**Technologies**:
- File I/O (streaming for large files)
- Regex pattern matching
- Concurrency with goroutines + channels
- Proper CLI interface with flags
- Error handling and validation

**Project structure**:
```
log-analyzer/
├── main.go
├── parser/
│   ├── log.go (Log entry struct and parser)
│   └── filter.go (Filter interface and implementations)
├── output/
│   ├── interface.go (Output interface)
│   ├── json.go (JSON formatter)
│   ├── csv.go (CSV formatter)
│   └── text.go (Text formatter)
├── worker/
│   └── processor.go (Concurrent processor with channels)
├── main_test.go
└── README.md
```

**Key patterns used**:
- Interface-based design for filters and formatters
- Goroutines and channels for concurrent processing
- Error wrapping for context
- Table-driven tests
- Proper resource cleanup (defer)

**Tests needed**:
- Parse various log formats
- Filter operations (level, date, keyword)
- Output formatting (all 3 formats)
- Concurrent processing correctness
- Edge cases (empty files, malformed lines)

**Estimated time**: 5-7 hours (includes concurrency complexity)

### Project 3: CLI Tool — Configuration Manager
**What it does**:
- Manage application configuration files
- Support multiple formats (JSON, INI-like, simple key=value)
- Merge configurations (defaults + user overrides)
- Validate configuration against schema
- Generate example/template configs
- Support environment variable expansion

**Technologies**:
- File I/O and text parsing
- Reflection for schema validation
- Marshaling/unmarshaling
- Error wrapping
- Type conversions

**Project structure**:
```
config-manager/
├── main.go
├── parser/
│   ├── interface.go (Parser interface)
│   ├── json.go (JSON parser)
│   └── ini.go (INI parser)
├── model/
│   └── config.go (Config struct with tags)
├── validator/
│   └── schema.go (Schema validation using reflection)
├── merger/
│   └── merge.go (Config merging logic)
├── main_test.go
└── README.md
```

**Key patterns used**:
- Interface-based parser design (easy to extend)
- Reflection for dynamic validation
- Error wrapping for clear error messages
- Struct tags for metadata
- Table-driven tests

**Tests needed**:
- Parse all formats
- Merge operations (defaults + overrides)
- Validation against schema
- Error handling (invalid types, missing required fields)
- Environment variable expansion

**Estimated time**: 4-6 hours implementation + tests

### Project 4: CLI Tool — File Synchronizer (Concurrent)
**What it does**:
- Sync files between source and destination directories
- Detect changes using file hashing (MD5/SHA256)
- Concurrent file operations for speed
- Dry-run mode (show what would be done)
- Conflict resolution strategies
- Exclude patterns support

**Technologies**:
- File I/O (concurrent operations)
- Goroutines + channels for coordination
- Hash-based change detection (crypto package)
- Proper concurrency patterns (WaitGroup, sync.Mutex)
- CLI flags for options

**Estimated time**: 5-7 hours implementation + tests

---

## Minimum Requirements

Every project MUST have:
- [ ] Proper package structure (main + 2+ additional packages)
- [ ] At least one interface definition
- [ ] Clear function/method documentation (godoc style)
- [ ] 3+ types of tests (unit, integration, edge cases)
- [ ] Error handling using error wrapping (`fmt.Errorf("%w", ...)`)
- [ ] README.md with usage instructions and examples
- [ ] Passes `gofmt`, `go vet`, `go test -race`
- [ ] 50%+ code coverage
- [ ] Meaningful git commits throughout development

---

## Production Practices

Apply these practices to make the code production-ready:

**Error Handling**:
```go
if err != nil {
    return fmt.Errorf("failed to parse config: %w", err)
}
```

**Interfaces for Abstraction**:
```go
type Storage interface {
    Save(ctx context.Context, todo *Todo) error
    Get(id string) (*Todo, error)
    List() ([]*Todo, error)
}
```

**Concurrency Safety**:
- Use channels or sync.Mutex for shared data
- Run tests with `go test -race` to detect data races
- Use `context` package for cancellation and timeouts

**Code Quality**:
- Variable names: `todoCount` not `tc`
- Comments explain "why", not "what"
- No magic numbers (use named constants)
- Defer resource cleanup: `defer file.Close()`

---

## Time Breakdown

| Task | Target | Actual |
|---|---|---|
| Design + architecture | 1 hour | ___ |
| Core implementation | 2.5 hours | ___ |
| Integration + features | 1.5 hours | ___ |
| Testing | 1 hour | ___ |
| Documentation + polish | 1 hour | ___ |
| Quality checks | 30 min | ___ |
| **Total** | **~7.5 hours** | ___ |

---

## End of Week Checklist

- [ ] Project chosen and designed
- [ ] Proper package structure created (2+ packages)
- [ ] Interface(s) defined and used
- [ ] Core functionality implemented
- [ ] Comprehensive tests written (3+ types)
- [ ] Error handling with wrapping implemented
- [ ] 50%+ code coverage achieved
- [ ] All code passes gofmt, go vet, go test -race
- [ ] README.md with clear instructions and examples
- [ ] Meaningful commits throughout
- [ ] Total hours logged: ___ / 7-8

---

## Actual Time Spent

**Total hours: ___**

**Project chosen**: ___

**Architecture decisions made**:

**Interfaces defined**:

**Challenges faced**:

**What I learned about:**
- Package design
- Error handling in production code
- Testing strategies
- Code organization

**Code quality metrics**:
- Coverage: ___%
- Gofmt: ✓/✗
- Go vet: ✓/✗
- Race detector: ✓/✗

**Key takeaway**: 

---

## Project Completion Status

**Weeks 1-6 Progress**:
- ✅ TGPL Chapters 1-12 all completed
- ✅ Advanced concepts studied (error wrapping, interfaces, performance)
- ✅ Mini project completed (config loader or similar)
- ✅ First substantial real-world project built
- ✅ Production-level code quality practices applied
- ✅ ~55 hours of 60-80 hour goal complete

**Ready for Week 7 — Advanced project #2 or portfolio piece!**
