# Phase 1: Software Development Foundation

**Duration**: 12-18 months
**Goal**: Be able to build production-grade services independently in Go
**Time Commitment**: 10 hours/week (~500-750 hours total)

---

## Overview

You cannot architect systems you cannot build. This phase builds the fundamental software development skills required for any architect role.

---

## Module 1: Master Go (4-6 months, ~200 hours)

### Learning Objectives
- [ ] Understand Go syntax, types, and control structures
- [ ] Master goroutines, channels, and concurrency patterns
- [ ] Write idiomatic Go code with proper error handling
- [ ] Create comprehensive tests (unit, table-driven, benchmarks)
- [ ] Design clean package structures
- [ ] Debug and profile Go applications

### Resources

#### Books
- [ ] "The Go Programming Language" by Donovan & Kernighan
  - Status: Not Started
  - Progress: Ch 0/13
- [ ] "Learning Go" by Jon Bodner (Ch 5-13)
  - Status: Not Started
- [ ] Effective Go (online docs)
  - Status: Not Started

#### Courses/Tutorials
- [ ] Tour of Go - https://go.dev/tour/
- [ ] Gophercises - https://gophercises.com/
- [ ] Go by Example - https://gobyexample.com/

### Projects

#### Project 1: Log Parser (CLI Tool)
**Goal**: Learn file I/O, string manipulation, basic CLI design

**Requirements**:
- Read log files from command line arguments
- Filter logs by pattern (regex support)
- Output statistics (error count, warning count, by timestamp)
- Support multiple output formats (JSON, CSV, plain text)
- Handle large files efficiently

**Skills Practiced**:
- File I/O
- String manipulation
- CLI flag parsing
- Error handling
- Testing

**Status**: Not Started
**GitHub**: (add link when created)

---

#### Project 2: Todo API (REST API)
**Goal**: Learn HTTP, JSON, database basics, API design

**Requirements**:
- REST endpoints (CRUD operations)
- SQLite database for persistence
- Proper HTTP status codes
- Input validation
- Basic authentication (API key)
- Structured logging
- Unit and integration tests

**Skills Practiced**:
- HTTP server
- JSON encoding/decoding
- Database operations
- API design
- Testing HTTP handlers

**Status**: Not Started
**GitHub**: (add link when created)

---

#### Project 3: Web Crawler (Concurrency)
**Goal**: Master goroutines, channels, and concurrent programming

**Requirements**:
- Crawl websites concurrently
- Follow links up to N levels deep
- Rate limiting (max requests per second)
- Respect robots.txt
- Handle errors gracefully
- Report statistics (pages crawled, errors, time taken)

**Skills Practiced**:
- Goroutines
- Channels
- Sync primitives (WaitGroup, Mutex)
- Context for cancellation
- Concurrent error handling

**Status**: Not Started
**GitHub**: (add link when created)

---

#### Project 4: Comprehensive Testing
**Goal**: Write tests for all above projects

**Requirements**:
- Unit tests for all functions
- Table-driven tests
- Integration tests
- Benchmarks for critical paths
- >80% code coverage

**Skills Practiced**:
- Testing patterns
- Mocking
- Test organization
- Coverage analysis

**Status**: Not Started

---

### Success Criteria

- [ ] Completed TGPL cover to cover with all exercises
- [ ] Completed all 20 Gophercises
- [ ] Built 3+ projects (above) without AI assistance
- [ ] Can explain goroutines, channels, and select
- [ ] Can write idiomatic Go code following community standards
- [ ] Comfortable debugging with delve or print debugging
- [ ] Got code review from senior Go developer
- [ ] Minimum 200 hours logged

---

## Module 2: Database Design & SQL (2-3 months, ~100 hours)

### Learning Objectives
- [ ] Design normalized database schemas
- [ ] Write complex SQL queries
- [ ] Understand indexes and when to use them
- [ ] Know transaction isolation levels
- [ ] Optimize slow queries

### Resources

#### Courses
- [ ] SQLBolt - https://sqlbolt.com/ (interactive)
- [ ] Mode SQL Tutorial - https://mode.com/sql-tutorial/

#### Books
- [ ] "SQL Performance Explained" by Markus Winand
  - Status: Not Started

### Projects

#### Project 1: E-commerce Schema Design
**Requirements**:
- Design schema for: users, products, orders, inventory, reviews
- Include proper relationships (1:1, 1:N, N:M)
- Normalize to 3NF
- Add appropriate indexes
- Document design decisions

**Status**: Not Started

---

#### Project 2: Query Practice
**Requirements**:
- Write queries for:
  - Top 10 products by revenue
  - User purchase history with aggregations
  - Running totals
  - Window functions for rankings
  - Complex joins across 4+ tables

**Status**: Not Started

---

#### Project 3: Query Optimization
**Requirements**:
- Take slow queries and optimize them
- Use EXPLAIN to understand query plans
- Add indexes where appropriate
- Measure performance improvements

**Status**: Not Started

---

#### Project 4: Add Database to Todo API
**Requirements**:
- Refactor Todo API to use PostgreSQL
- Implement connection pooling
- Add migrations
- Handle transactions properly
- Add database tests

**Status**: Not Started
**GitHub**: (link to updated todo API)

---

### Success Criteria

- [ ] Can design schemas for common applications
- [ ] Write complex SQL without googling basic syntax
- [ ] Understand EXPLAIN output
- [ ] Know when to denormalize for performance
- [ ] Comfortable with transactions and isolation levels
- [ ] Minimum 100 hours logged

---

## Module 3: APIs & Integration Patterns (2-3 months, ~100 hours)

### Learning Objectives
- [ ] Design clean REST APIs
- [ ] Build gRPC services
- [ ] Implement authentication/authorization
- [ ] Handle API versioning
- [ ] Design for rate limiting and pagination

### Resources

#### Books
- [ ] "REST API Design Rulebook" by Mark Masse
  - Status: Not Started

#### Documentation
- [ ] Google API Design Guide - https://cloud.google.com/apis/design
- [ ] gRPC Go Tutorial - https://grpc.io/docs/languages/go/

### Projects

#### Project 1: Enhanced REST API
**Requirements**:
- Proper REST design (resource-based)
- Versioning (URL or header)
- Pagination
- Filtering and sorting
- Rate limiting
- JWT authentication
- Comprehensive error responses
- OpenAPI/Swagger documentation

**Status**: Not Started
**GitHub**: (add link when created)

---

#### Project 2: gRPC Service
**Requirements**:
- Same functionality as REST API but in gRPC
- Protocol buffer definitions
- Streaming support (at least one example)
- Error handling with status codes
- Interceptors for logging/auth

**Status**: Not Started
**GitHub**: (add link when created)

---

#### Project 3: API Gateway Pattern
**Requirements**:
- Simple API gateway that routes to multiple services
- Authentication at gateway level
- Rate limiting
- Request/response logging
- Metrics collection

**Status**: Not Started
**GitHub**: (add link when created)

---

### Success Criteria

- [ ] Can design REST APIs following best practices
- [ ] Understand REST vs gRPC trade-offs
- [ ] Implemented JWT authentication
- [ ] Written complete API documentation
- [ ] Can explain API versioning strategies
- [ ] Minimum 100 hours logged

---

## Module 4: Distributed Systems Basics (3-4 months, ~150 hours)

### Learning Objectives
- [ ] Understand CAP theorem and trade-offs
- [ ] Design for failure (retries, timeouts, circuit breakers)
- [ ] Implement async communication with message queues
- [ ] Add observability (logging, metrics, tracing)
- [ ] Know consistency models

### Resources

#### Books
- [ ] "Designing Data-Intensive Applications" by Martin Kleppmann (Ch 1-9)
  - **CRITICAL RESOURCE**
  - Status: Not Started
  - Progress: Ch 0/9

#### Online
- [ ] MIT 6.824 Distributed Systems (first 10 lectures) - YouTube
- [ ] Google SRE Book (monitoring, toil, SLOs chapters) - https://sre.google/sre-book/

### Projects

#### Project 1: Distributed Services
**Requirements**:
- Build 2-3 services that communicate
- One service uses HTTP (sync)
- One uses message queue (async) - Redis/RabbitMQ
- Implement retry logic with exponential backoff
- Add circuit breaker pattern
- Proper timeout handling

**Status**: Not Started
**GitHub**: (add link when created)

---

#### Project 2: Observability Stack
**Requirements**:
- Add structured logging to all services
- Instrument with Prometheus metrics
- Add distributed tracing (Jaeger/Zipkin)
- Create Grafana dashboards
- Set up alerts for error rates

**Status**: Not Started
**GitHub**: (add link when created)

---

#### Project 3: Chaos Engineering
**Requirements**:
- Deliberately cause failures in distributed system
- Test retry logic
- Test circuit breakers
- Test timeout handling
- Document failure scenarios and recovery

**Status**: Not Started
**GitHub**: (add link when created)

---

### Success Criteria

- [ ] Can explain CAP theorem with examples
- [ ] Implemented retry/timeout/circuit breaker patterns
- [ ] Built system with sync and async communication
- [ ] Full observability (logs, metrics, traces)
- [ ] Read DDIA chapters 1-9 completely
- [ ] Can discuss trade-offs in distributed systems
- [ ] Minimum 150 hours logged

---

## Phase 1 Overall Success Criteria

### Technical Skills
- [ ] Built 8-10 working projects, all on GitHub
- [ ] Can build REST and gRPC APIs from scratch
- [ ] Can design database schemas and write complex SQL
- [ ] Understand distributed systems fundamentals
- [ ] Can instrument code for observability

### Soft Skills
- [ ] Can explain technical concepts clearly in writing
- [ ] Written at least 3 blog posts about what you learned
- [ ] Got code reviewed by senior engineers

### Metrics
- [ ] Minimum 500 hours logged
- [ ] All projects have README with setup instructions
- [ ] All projects have tests with >70% coverage
- [ ] At least 1 project has production-grade quality

### Mindset
- [ ] Can build without AI crutch
- [ ] Comfortable with being stuck and debugging
- [ ] Understand WHY solutions work, not just THAT they work

---

## Principal Engineer Extensions - Phase 1

### System-Level Decision Making
- Choose libraries and frameworks based on long-term maintainability
- Consider operational complexity in technology selections
- Design with debugging and monitoring in mind from the start

### Trade-Off Analysis
- Balance development velocity vs. system reliability
- Evaluate memory/CPU efficiency in algorithm choices
- Consider security implications of dependencies

### Long-Term Technical Vision
- Structure code for future scalability and modification
- Design APIs that can evolve without breaking changes
- Choose patterns that support business growth

### Cross-Team Impact Awareness
- Design services with clear interfaces and contracts
- Consider operational monitoring needs
- Think about integration with existing systems

### Real-World Scenario Exercises
1. **Prototype to Production Evolution**: Transform a simple tool into a production service with proper error handling, configuration, and logging
2. **Zero-Downtime Migration**: Design a database schema change that can be deployed without service interruption
3. **API Versioning Strategy**: Modify an API to add features while maintaining backward compatibility

---

## If You Can't Check All Boxes Above

**DO NOT move to Phase 2.** Going forward without this foundation will waste your time.

Fix gaps first, then proceed.

---

**Last Updated**: March 21, 2026
