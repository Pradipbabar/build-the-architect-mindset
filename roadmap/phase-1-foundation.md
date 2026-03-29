# Phase 1: Software Development Foundation

**Duration**: 12 months
**Goal**: Be able to build production-grade services independently in Go, with infrastructure awareness
**Time Commitment**: 10 hours/week (~600 hours total)

---

## Overview

You cannot architect systems you cannot build. This phase builds the fundamental software development skills required for any architect role, with early introduction to infrastructure concepts that architects must understand.

---

## Module 1: Master Go (4-6 months, ~200 hours)

### Learning Objectives
- [ ] Understand Go syntax, types, and control structures
- [ ] Master goroutines, channels, and concurrency patterns
- [ ] Write idiomatic Go code with proper error handling
- [ ] Create comprehensive tests (unit, table-driven tests, mocks, benchmarks)
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
- Rate limiting (configurable req/sec)
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

## Module 2: Infrastructure Fundamentals (2 months, ~80 hours)

**Why Now**: Architects must understand deployment and operations. Don't learn this later when it's too late.

### Learning Objectives
- [ ] Understand containerization and OCI standards
- [ ] Implement observability in applications (metrics, logs, traces)
- [ ] Know networking basics for distributed systems
- [ ] Compare REST vs gRPC for API design

### Resources

#### Containerization
- [ ] Docker documentation and best practices
- [ ] "Kubernetes in Action" (Chapter 1-3 for container basics)
- [ ] OCI specifications

#### Observability
- [ ] OpenTelemetry documentation
- [ ] Prometheus client libraries
- [ ] Structured logging with slog/zap

#### Networking & APIs
- [ ] gRPC documentation and tutorials
- [ ] TCP/IP basics for developers
- [ ] Service discovery concepts

### Projects

#### Project 1: Containerize All Applications
**Goal**: Learn Docker and container best practices

**Requirements**:
- Write production-grade Dockerfiles for all Phase 1 projects
- Multi-stage builds for optimization
- Security best practices (non-root user, minimal base images)
- Build and push to container registry
- Document container build process

**Skills Practiced**:
- Containerization
- Image optimization
- Security in containers
- Registry operations

**Status**: Not Started

---

#### Project 2: Add Observability
**Goal**: Implement three pillars of observability

**Requirements**:
- Add structured logging to all projects (JSON format)
- Implement Prometheus metrics collection
- Add basic tracing with OpenTelemetry
- Create health check endpoints
- Document observability setup

**Skills Practiced**:
- Logging patterns
- Metrics collection
- Tracing basics
- Health monitoring

**Status**: Not Started

---

#### Project 3: gRPC Alternative Implementation
**Goal**: Compare REST vs gRPC for API design

**Requirements**:
- Extend Todo API to support gRPC alongside REST
- Define proto3 schemas
- Implement gRPC server and client
- Compare performance and developer experience
- Document trade-offs and use cases

**Skills Practiced**:
- Protocol buffer design
- gRPC implementation
- API design comparison
- Performance analysis

**Status**: Not Started

---

#### Project 4: Networking Basics
**Goal**: Understand network layer for distributed systems

**Requirements**:
- Implement service discovery in Todo API
- Add load balancing concepts
- Understand TCP connection handling
- Document network architecture decisions

**Skills Practiced**:
- Service discovery
- Load balancing
- Network programming
- Connection management

**Status**: Not Started

---

### Success Criteria

- [ ] All projects containerized with production Dockerfiles
- [ ] Observability implemented (logs, metrics, traces)
- [ ] gRPC API implemented and compared to REST
- [ ] Networking concepts understood and documented
- [ ] Minimum 80 hours logged

---

## Module 3: Database Design & SQL (2-3 months, ~100 hours)

### Learning Objectives
- [ ] Design normalized relational schemas
- [ ] Write complex SQL queries with JOINs and aggregations
- [ ] Optimize slow queries using indexes and EXPLAIN
- [ ] Understand ACID properties and isolation levels
- [ ] Handle database transactions properly

### Resources

#### Books
- [ ] "SQL and Relational Theory" by C.J. Date
- [ ] "Database System Concepts" (Chapter 6-8 for SQL)
- [ ] PostgreSQL documentation

#### Online Resources
- [ ] SQLZoo - https://sqlzoo.net/
- [ ] LeetCode SQL problems
- [ ] PostgreSQL tutorial

### Projects

#### Project 1: Schema Design
**Goal**: Learn database design and normalization

**Requirements**:
- Design e-commerce database schema (users, products, orders, inventory)
- Apply proper normalization (3NF minimum)
- Identify appropriate indexes
- Write schema creation scripts
- Document design decisions

**Skills Practiced**:
- Schema design
- Normalization
- Index selection
- SQL DDL

**Status**: Not Started

---

#### Project 2: Query Optimization
**Goal**: Master SQL query writing and optimization

**Requirements**:
- Write complex queries (multiple JOINs, aggregations, window functions)
- Analyze query performance with EXPLAIN
- Optimize slow queries with indexes
- Handle complex business logic in SQL
- Document optimization results

**Skills Practiced**:
- Complex SQL
- Query optimization
- Index tuning
- Performance analysis

**Status**: Not Started

---

#### Project 3: Database Integration
**Goal**: Integrate database with Go applications

**Requirements**:
- Refactor Todo API to use PostgreSQL instead of SQLite
- Implement connection pooling
- Add database migrations
- Handle transactions properly
- Add database-specific error handling

**Skills Practiced**:
- Database drivers
- Connection management
- Migrations
- Transaction handling

**Status**: Not Started

---

### Success Criteria

- [ ] Can design normalized schemas for common domains
- [ ] Can write complex SQL without googling basic syntax
- [ ] Can optimize queries using EXPLAIN and indexes
- [ ] Understand isolation levels and their trade-offs
- [ ] All projects completed with proper database integration
- [ ] Minimum 100 hours logged

---

## Module 5: APIs & Integration Patterns (2-3 months, ~100 hours)

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

## Module 5: APIs & Integration Patterns (2-3 months, ~100 hours)

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

## Module 6: Distributed Systems Basics (2 months, ~120 hours)

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
- [ ] Minimum 120 hours logged

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
