# Phase 1: Software Development Foundation - Detailed Modules

**Duration**: 12 Months (exactly)  
**Total Hours**: 500-600 hours at 10 hrs/week  
**Start**: March 21, 2026  
**End**: March 21, 2027  

---

## Module 1: Go Fundamentals (4 months, 200 hours)

**Period**: Mar 21 - Jun 21, 2026  
**Weekly Commitment**: 12-15 hours/week (higher concentration)

### Goals
- Build production-grade Go applications
- Master concurrency (goroutines, channels, sync)
- Write idiomatic, maintainable code
- Comprehensive testing skills

### Submodule Breakdown

#### 1.1: Go Basics (3 weeks, 50 hours)
- [ ] The Go Programming Language Book - Chapters 1-6
- [ ] Go syntax, types, functions
- [ ] Error handling patterns
- [ ] Package design
- [ ] Testing basics

**Deliverable**: Quiz yourself on all syntax

#### 1.2: Concurrency Mastery (4 weeks, 60 hours)
- [ ] TGPL Chapters 7-9 (focus on concurrency)
- [ ] Goroutines deep dive
- [ ] Channels (buffered, unbuffered, select)
- [ ] Sync primitives (WaitGroup, Mutex, Once)
- [ ] Context for cancellation
- [ ] Common pitfalls & deadlocks

**Deliverable**: Mini concurrency exercises (at least 5)

#### 1.3: Testing & Best Practices (3 weeks, 50 hours)
- [ ] Testing patterns (unit, integration, benchmarks)
- [ ] Table-driven tests
- [ ] Mocking strategies
- [ ] Test coverage analysis
- [ ] Code organization

**Deliverable**: Test 1 project to >80% coverage

#### 1.4: Build 3 Projects (8 weeks, 40 hours distributed)
Complete these in sequence:

**Project 1: Log Parser CLI** (2 weeks, 12 hours)
- Read log files, filter by pattern (regex)
- Output statistics (JSON, CSV, plaintext)
- Handle large files efficiently
- No databases, focus on I/O and CLI design

**Deliverable**: Working CLI, >80% test coverage, GitHub repo

**Project 2: Todo REST API** (3 weeks, 18 hours)
- REST CRUD endpoints
- SQLite persistence
- Proper HTTP status codes
- Input validation
- API key authentication
- Unit + integration tests

**Deliverable**: Completed API, >80% test coverage, GitHub repo

**Project 3: Web Crawler** (3 weeks, 18 hours)
- Crawl websites concurrently
- Follow links recursively
- Rate limiting (configurable req/sec)
- Respect robots.txt
- Stats reporting (concurrent, handled gracefully)
- Real concurrency tests

**Deliverable**: Working crawler, production-ready, GitHub repo

### Success Criteria (ALL required ✅)
- [ ] Completed TGPL book with all exercises
- [ ] Completed 20+ Gophercises
- [ ] 3 projects on GitHub with >80% test coverage
- [ ] Can explain goroutines, channels, and select
- [ ] Can write idiomatic Go code
- [ ] Code review approved by senior Go dev
- [ ] **200 hours logged and documented**
- [ ] **At least 1 project got external code review**

### Time Tracking Template
```
Week of [dates]:
- Mon: ___ hours on ___ (what specifically)
- Tue: ___ hours on ___
- Wed: ___ hours on ___
- Thu: ___ hours on ___
- Fri: ___ hours on ___
- Sat: ___ hours on ___
- Sun: ___ hours on ___
Total: ___ hours (goal: 12-15)

Deliverables completed:
- [ ] Specific thing 1
- [ ] Specific thing 2

Blockers:
-

Next week focus:
-
```

---

## Module 2: Database Design & SQL (2.5 months, 100 hours)

**Period**: Jun 21 - Sep 6, 2026  
**Weekly Commitment**: 10 hours/week

### Goals
- Design normalized relational schemas
- Write complex SQL queries
- Optimize slow queries
- Understand database fundamentals deeply

### Submodule Breakdown

#### 2.1: Database Fundamentals (2 weeks, 20 hours)
- [ ] ACID properties
- [ ] Relational algebra
- [ ] Normalization (1NF, 2NF, 3NF, BCNF)
- [ ] Keys (primary, foreign, candidate, composite)
- [ ] Constraints and referential integrity

**Deliverable**: Design 3 schemas with proper normalization

#### 2.2: SQL Query Deep Dive (3 weeks, 30 hours)
- [ ] Complex SELECT queries
- [ ] JOINs (all types, performance implications)
- [ ] Aggregation & GROUP BY
- [ ] Window functions
- [ ] Subqueries & CTEs
- [ ] Set operations (UNION, INTERSECT, EXCEPT)

**Deliverable**: Solve 50+ SQL problems from SQL exercises

#### 2.3: Indexes & Performance (2 weeks, 20 hours)
- [ ] B-tree, Hash, GiST indexes
- [ ] EXPLAIN output analysis
- [ ] Query plan optimization
- [ ] Index selection strategies
- [ ] Common mistakes

**Deliverable**: Optimize 10 slow queries, measure improvements

#### 2.4: Transactions & Isolation (2 weeks, 15 hours)
- [ ] Isolation levels (Read Uncommitted, Read Committed, Repeatable Read, Serializable)
- [ ] Phantom reads, dirty reads, non-repeatable reads
- [ ] Deadlock analysis
- [ ] Locking mechanisms

**Deliverable**: Demonstrate each isolation level issue

#### 2.5: Building Projects (3 weeks, 15 hours)

**Project 1: Schema Design**
- E-commerce schema (users, products, orders, inventory)
- Properly normalized to 3NF
- Appropriate indexes identified
- Design document with rationale

**Project 2: Query Practice**
- Top 10 products by revenue
- User purchase history with aggregations
- Running totals and rankings
- Complex joins (4+ tables)

**Project 3: Database Integration**
- Refactor Todo API to use PostgreSQL
- Connection pooling
- Migrations (schema versioning)
- Transaction handling
- Proper cleanup/testing

### Success Criteria (ALL required ✅)
- [ ] Can design schemas for common domains
- [ ] Can write complex SQL without googling basic syntax
- [ ] Can read and optimize EXPLAIN output
- [ ] Understand isolation levels and trade-offs
- [ ] 3 projects completed with documentation
- [ ] **100 hours logged and documented**

---

## Module 3: APIs & Integration Patterns (2.5 months, 100 hours)

**Period**: Sep 6 - Dec 6, 2026  
**Weekly Commitment**: 10 hours/week

### Goals
- Design clean, maintainable REST APIs
- Build gRPC services
- Understand authentication/authorization
- API versioning and evolution

### Submodule Breakdown

#### 3.1: REST API Design (3 weeks, 30 hours)
- [ ] REST principles (resources, verbs, status codes)
- [ ] OpenAPI/Swagger specification
- [ ] Pagination, filtering, sorting
- [ ] API versioning strategies
- [ ] Error handling and responses
- [ ] API documentation best practices

**Deliverable**: Design REST API spec for multi-feature system

#### 3.2: gRPC & Protocol Buffers (2 weeks, 20 hours)
- [ ] Protocol Buffers basics
- [ ] Defining services and RPC methods
- [ ] Streaming (client, server, bidirectional)
- [ ] Error handling in gRPC
- [ ] Interceptors for cross-cutting concerns
- [ ] gRPC vs REST trade-offs

**Deliverable**: Convert REST API to gRPC, compare

#### 3.3: Authentication & Authorization (2 weeks, 20 hours)
- [ ] JWT (JSON Web Tokens)
- [ ] OAuth 2.0 fundamentals
- [ ] API keys and secrets management
- [ ] Role-based access control (RBAC)
- [ ] Securing credentials in code

**Deliverable**: Implement JWT in a service

#### 3.4: Advanced Patterns (2 weeks, 15 hours)
- [ ] Rate limiting (token bucket, sliding window)
- [ ] Request signing (AWS SigV4)
- [ ] Idempotency
- [ ] API versioning (URL, header, accept header)
- [ ] Backward compatibility

**Deliverable**: Design versioning strategy for evolving API

#### 3.5: Building Projects (3 weeks, 15 hours)

**Project 1: Enhanced REST API**
- Full REST design (resource-based, proper verbs)
- Pagination, filtering, sorting
- JWT authentication
- Rate limiting
- OpenAPI documentation
- >80% test coverage

**Project 2: gRPC Service**
- Protocol buffer definitions
- Streaming endpoints (at least 1)
- Interceptors for logging/auth
- Error handling with proper codes
- Performance comparable to REST

**Project 3: API Gateway Pattern**
- Routes to multiple backend services
- Authentication at gateway level
- Request/response logging
- Rate limiting per user
- Metrics collection

### Success Criteria (ALL required ✅)
- [ ] Can design REST APIs following best practices
- [ ] Understand REST vs gRPC trade-offs deeply
- [ ] Can implement JWT authentication
- [ ] Can write OpenAPI specs
- [ ] 3 projects completed
- [ ] **100 hours logged and documented**

---

## Module 4: Distributed Systems Basics (3 months, 150 hours)

**Period**: Dec 6, 2026 - Mar 21, 2027  
**Weekly Commitment**: 10-12 hours/week

### Goals
- Understand distributed systems fundamentals
- Know common patterns and pitfalls
- Can reason about distributed systems

### Submodule Breakdown

#### 4.1: Fundamentals (4 weeks, 40 hours)
- [ ] Fallacies of Distributed Computing
- [ ] Network partitions and failure modes
- [ ] Consensus algorithms (Raft basics)
- [ ] Message delivery guarantees (at-most-once, at-least-once, exactly-once)
- [ ] Event ordering and causality

**Deliverable**: Implement a simple consensus algorithm

#### 4.2: Databases & Consistency (3 weeks, 35 hours)
- [ ] CAP theorem and replication strategies
- [ ] Read replicas and eventual consistency
- [ ] Multi-master replication
- [ ] Data partitioning/sharding
- [ ] Handling splits and merges

**Deliverable**: Design a sharded database system

#### 4.3: Message Queues & Event Driven (3 weeks, 35 hours)
- [ ] Message brokers (Kafka, RabbitMQ)
- [ ] Partition management
- [ ] Consumer groups
- [ ] Dead letter queues
- [ ] Message ordering guarantees

**Deliverable**: Implement event-driven system

#### 4.4: Monitoring & Observability (2 weeks, 20 hours)
- [ ] Health checks and heartbeats
- [ ] Metrics collection
- [ ] Distributed tracing
- [ ] Log aggregation
- [ ] Alerting strategies

**Deliverable**: Instrument one of your projects fully

#### 4.5: Building Projects (6 weeks, 20 hours distributed)

**Project 1: Multi-Service System**
- Multiple Go services
- Async communication with message queue
- Handle eventual consistency
- Proper error handling and retries

**Project 2: Fault Tolerance**
- Circuit breakers
- Timeouts and retries
- Graceful degradation
- Bulkheads

**Project 3: Observability Stack**
- Prometheus for metrics
- Loki for logs (if available)
- Traces with OpenTelemetry
- Meaningful dashboards and alerts

### Success Criteria (ALL required ✅)
- [ ] Can explain CAP theorem
- [ ] Understand message delivery guarantees
- [ ] Know consensus algorithm basics
- [ ] Can design a sharded system
- [ ] 3+ distributed projects completed
- [ ] **150 hours logged and documented**

---

## Phase 1 Completion Checklist

### Projects (ALL must be done ✅)
- [ ] Log Parser CLI - GitHub with >80% coverage
- [ ] Todo REST API - GitHub with >80% coverage
- [ ] Web Crawler - GitHub with >80% coverage
- [ ] E-commerce Schema - Database design doc
- [ ] Query Optimization - 10 queries optimized
- [ ] REST API - Full implementation
- [ ] gRPC Service - Full implementation
- [ ] API Gateway - Full implementation
- [ ] Multi-Service System - Implemented
- [ ] Observability Stack - Implemented

### Documentation (ALL must be done ✅)
- [ ] Module 1 Summary (Go lessons learned)
- [ ] Module 2 Summary (Database lessons learned)
- [ ] Module 3 Summary (API lessons learned)
- [ ] Module 4 Summary (Distributed systems lessons learned)

### Time Tracking (MUST be verified ✅)
- [ ] 500+ hours logged in time tracking system
- [ ] All hours categorized by module
- [ ] No more than 2 week gaps

### Code Quality (ALL must be met ✅)
- [ ] All projects have >80% test coverage
- [ ] At least 2 projects reviewed by senior dev
- [ ] Code follows Go best practices
- [ ] All projects documented with README

### Phase 1 COMPLETION CRITERIA
🎯 **ALL of the above must be ✅ to start Phase 2**

---

**Created**: March 22, 2026  
**For questions**: Review [ROADMAP-PROGRESS.md](ROADMAP-PROGRESS.md)
