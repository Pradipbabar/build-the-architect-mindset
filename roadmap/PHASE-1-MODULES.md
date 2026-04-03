# Phase 1: Software Development Foundation - Detailed Modules

**Duration**: 12+ Months (flexible based on time availability)  
**Total Hours**: ~400+ hours at realistic weekly commitment  
**Start**: March 21, 2026  
**End**: Open-ended (prioritize depth over speed)

---

## ⚠️ REALISTIC NOTE FOR BUSY PROFESSIONALS

This roadmap is designed for people with **full-time work commitments**. Rather than the original aggressive timeline (12-15 hrs/week), we focus on **sustainable learning** at 5-7 hours per week. Quality always beats quantity.

---

## Module 1: Go Fundamentals (3-4 months, 60-80 hours)

**Period**: Mar 21 - Jun 21, 2026 (12 weeks)  
**Weekly Commitment**: 5-7 hours/week (sustainable with full-time work)  
**Realistic Distribution**: 
- Weekdays: 1-1.5 hours (evening study) = 5-7.5 hours
- Weekends: 1-2 hours (if desired)
- **OR** just weekends: 5-7 hours Saturday/Sunday

### Goals
- Build solid Go fundamentals (not production-grade yet)
- Understand concurrency concepts (goroutines, channels, basics)
- Write clean, testable code
- Complete core TGPL + practical projects

### Submodule Breakdown

#### 1.1: Go Basics & Syntax (4 weeks, 20-24 hours)
- [ ] The Go Programming Language Book - Chapters 1-4
- [ ] Go syntax, types, pointers
- [ ] Error handling patterns
- [ ] Testing basics (table-driven tests)
- [ ] 2-3 small Gophercises (#1, #2, optionally #3)

**Deliverable**: Chapters 1-4 notes + 2-3 Gophercises with tests

#### 1.2: Functions, Methods & Interfaces (3 weeks, 15-18 hours)
- [ ] TGPL Chapters 5-6 (functions, methods, interfaces)
- [ ] Function design in Go
- [ ] Method receivers (value vs pointer)
- [ ] Interface concepts
- [ ] 1 small project: Simple CLI or text processor

**Deliverable**: Chapters 5-6 notes + 1 working project on GitHub

#### 1.3: Concurrency Foundations (3 weeks, 15-18 hours)
- [ ] TGPL Chapters 7-9 (concurrency, packages, testing)
- [ ] Goroutines basics
- [ ] Channels (unbuffered, buffered)
- [ ] Select statement
- [ ] 1 more small project: Concurrent program or enhanced CLI

**Deliverable**: Chapters 7-9 notes + 1 concurrent project on GitHub

#### 1.4: Code Quality & Review (2 weeks, 10-12 hours)
- [ ] Code quality practices (gofmt, go vet, test coverage)
- [ ] Idioms and best practices
- [ ] Code review of your own projects
- [ ] Self-assessment on fundamentals

**Deliverable**: All projects pass quality checks, GitHub ready

### Success Criteria (REALISTIC ✅)
- [x] Completed TGPL Chapters 1-9 with typed notes
- [x] Completed 2-3 Gophercises (#1, #2, optionally #3)
- [x] Built 2 projects on GitHub with >60% test coverage
- [x] Can explain goroutines, channels, select clearly
- [x] Can write clean, idiomatic Go code
- [x] **60-80 hours logged and documented**
- [x] Code passes gofmt, go vet, go test standards

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

## Module 2: Infrastructure Fundamentals (2 months, 80 hours)

**Period**: Jun 21 - Aug 21, 2026  
**Weekly Commitment**: 10 hours/week

### Goals
- Understand containerization and deployment
- Implement observability in applications
- Know networking basics for distributed systems
- Compare API technologies (REST vs gRPC)

### Submodule Breakdown

#### 2.1: Containerization & Docker (4 weeks, 40 hours)
- [ ] Docker fundamentals and OCI standards
- [ ] Writing production Dockerfiles
- [ ] Multi-stage builds and optimization
- [ ] Security best practices
- [ ] Container registries and deployment

**Deliverable**: All Phase 1 projects containerized

#### 2.2: Observability Implementation (3 weeks, 25 hours)
- [ ] Structured logging (slog, zap)
- [ ] Metrics collection (Prometheus client)
- [ ] Tracing basics (OpenTelemetry)
- [ ] Health checks and monitoring
- [ ] Observability best practices

**Deliverable**: Observability added to all projects

#### 2.3: gRPC & API Design (2 weeks, 15 hours)
- [ ] Protocol buffers and proto3
- [ ] gRPC service definition
- [ ] Streaming and error handling
- [ ] REST vs gRPC comparison
- [ ] When to use each

**Deliverable**: gRPC implementation in Todo API

#### 2.4: Networking Basics (1 week, 10 hours)
- [ ] TCP/IP fundamentals for developers
- [ ] Service discovery concepts
- [ ] Load balancing basics
- [ ] Connection handling

**Deliverable**: Networking concepts documented

### Success Criteria (ALL required ✅)
- [ ] All projects containerized with production Dockerfiles
- [ ] Observability implemented (logs, metrics, traces)
- [ ] gRPC API implemented alongside REST
- [ ] Networking concepts understood
- [ ] **80 hours logged and documented**

---

## Module 4: Database Design & SQL (2.5 months, 100 hours)

**Period**: Aug 21 - Nov 6, 2026  
**Weekly Commitment**: 10 hours/week

### Goals
- Design normalized relational schemas
- Write complex SQL queries
- Optimize slow queries
- Understand database fundamentals deeply

### Submodule Breakdown

#### 4.1: Database Fundamentals (2 weeks, 20 hours)
- [ ] ACID properties
- [ ] Relational algebra
- [ ] Normalization (1NF, 2NF, 3NF, BCNF)
- [ ] Keys (primary, foreign, candidate, composite)
- [ ] Constraints and referential integrity

**Deliverable**: Design 3 schemas with proper normalization

#### 4.2: SQL Query Deep Dive (3 weeks, 30 hours)
- [ ] Complex SELECT queries
- [ ] JOINs (all types, performance implications)
- [ ] Aggregation & GROUP BY
- [ ] Window functions
- [ ] Subqueries & CTEs
- [ ] Set operations (UNION, INTERSECT, EXCEPT)

**Deliverable**: Solve 50+ SQL problems from SQL exercises

#### 4.3: Indexes & Performance (2 weeks, 20 hours)
- [ ] B-tree, Hash, GiST indexes
- [ ] EXPLAIN output analysis
- [ ] Query plan optimization
- [ ] Index selection strategies
- [ ] Common mistakes

**Deliverable**: Optimize 10 slow queries, measure improvements

#### 4.4: Transactions & Isolation (2 weeks, 15 hours)
- [ ] Isolation levels (Read Uncommitted, Read Committed, Repeatable Read, Serializable)
- [ ] Phantom reads, dirty reads, non-repeatable reads
- [ ] Deadlock analysis
- [ ] Locking mechanisms

**Deliverable**: Demonstrate each isolation level issue

#### 4.5: Building Projects (3 weeks, 15 hours)

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

## Module 5: APIs & Integration Patterns (2.5 months, 100 hours)

**Period**: Nov 6, 2026 - Jan 21, 2027  
**Weekly Commitment**: 10 hours/week

### Goals
- Design clean, maintainable REST APIs
- Build gRPC services
- Understand authentication/authorization
- API versioning and evolution

### Submodule Breakdown

#### 5.1: REST API Design (3 weeks, 30 hours)
- [ ] REST principles (resources, verbs, status codes)
- [ ] OpenAPI/Swagger specification
- [ ] Pagination, filtering, sorting
- [ ] API versioning strategies
- [ ] Error handling and responses
- [ ] API documentation best practices

**Deliverable**: Design REST API spec for multi-feature system

#### 5.2: gRPC & Protocol Buffers (2 weeks, 20 hours)
- [ ] Protocol Buffers basics
- [ ] Defining services and RPC methods
- [ ] Streaming (client, server, bidirectional)
- [ ] Error handling in gRPC
- [ ] Interceptors for cross-cutting concerns
- [ ] gRPC vs REST trade-offs

**Deliverable**: Convert REST API to gRPC, compare

#### 5.3: Authentication & Authorization (2 weeks, 20 hours)
- [ ] JWT (JSON Web Tokens)
- [ ] OAuth 2.0 fundamentals
- [ ] API keys and secrets management
- [ ] Role-based access control (RBAC)
- [ ] Securing credentials in code

**Deliverable**: Implement JWT in a service

#### 5.4: Advanced Patterns (2 weeks, 15 hours)
- [ ] Rate limiting (token bucket, sliding window)
- [ ] Request signing (AWS SigV4)
- [ ] Idempotency
- [ ] API versioning (URL, header, accept header)
- [ ] Backward compatibility

**Deliverable**: Design versioning strategy for evolving API

#### 5.5: Building Projects (3 weeks, 15 hours)

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

## Module 6: Distributed Systems Basics (2 months, 120 hours)

**Period**: Jan 21 - Mar 21, 2027  
**Weekly Commitment**: 10-12 hours/week

### Goals
- Understand distributed systems fundamentals
- Know common patterns and pitfalls
- Can reason about distributed systems

### Submodule Breakdown

#### 6.1: Fundamentals (4 weeks, 40 hours)
- [ ] Fallacies of Distributed Computing
- [ ] Network partitions and failure modes
- [ ] Consensus algorithms (Raft basics)
- [ ] Message delivery guarantees (at-most-once, at-least-once, exactly-once)
- [ ] Event ordering and causality

**Deliverable**: Implement a simple consensus algorithm

#### 6.2: Databases & Consistency (3 weeks, 30 hours)
- [ ] CAP theorem and replication strategies
- [ ] Read replicas and eventual consistency
- [ ] Multi-master replication
- [ ] Data partitioning/sharding
- [ ] Handling splits and merges

**Deliverable**: Design a sharded database system

#### 6.3: Message Queues & Event Driven (3 weeks, 30 hours)
- [ ] Message brokers (Kafka, RabbitMQ)
- [ ] Partition management
- [ ] Consumer groups
- [ ] Dead letter queues
- [ ] Message ordering guarantees

**Deliverable**: Implement event-driven system

#### 6.4: Monitoring & Observability (2 weeks, 20 hours)
- [ ] Health checks and heartbeats
- [ ] Metrics collection
- [ ] Distributed tracing
- [ ] Log aggregation
- [ ] Alerting strategies

**Deliverable**: Instrument one of your projects fully

#### 6.5: Building Projects (4 weeks, 20 hours distributed)

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
- [ ] **120 hours logged and documented**

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
- [ ] Module 2 Summary (Infrastructure lessons learned)
- [ ] Module 4 Summary (Database lessons learned)
- [ ] Module 5 Summary (API lessons learned)
- [ ] Module 6 Summary (Distributed systems lessons learned)

### Time Tracking (MUST be verified ✅)
- [ ] 600+ hours logged in time tracking system
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
