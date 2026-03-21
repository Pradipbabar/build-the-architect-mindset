# Phase 2: System Design & Architecture

**Duration**: 12-18 months
**Goal**: Understand architectural patterns and design systems end-to-end
**Time Commitment**: 10 hours/week (~500-750 hours total)
**Prerequisites**: Phase 1 MUST be complete

---

## Overview

Now that you can build systems, you'll learn to design them. This phase focuses on architectural thinking, patterns, and system design skills.

---

## Module 1: Architectural Patterns (3-4 months, ~150 hours)

### Learning Objectives
- [ ] Understand monolith vs microservices trade-offs
- [ ] Know when to use event-driven architecture
- [ ] Understand CQRS and Event Sourcing
- [ ] Know distributed transaction patterns (Saga)
- [ ] Can diagram architectures clearly
- [ ] Write Architecture Decision Records (ADRs)

### Resources

#### Books
- [ ] "Building Microservices" by Sam Newman (2nd edition)
  - Status: Not Started
- [ ] "Software Architecture: The Hard Parts" by Ford et al.
  - Status: Not Started

#### Online
- [ ] microservices.io patterns - https://microservices.io/patterns/
- [ ] Martin Fowler's Architecture articles - https://martinfowler.com/architecture/

### Projects

#### Project 1: Monolith to Microservices
**Requirements**:
- Take one Phase 1 project (monolith)
- Identify service boundaries (don't just split randomly)
- Break into 2-3 microservices
- Document WHY you split where you did
- Implement service-to-service communication
- Compare complexity: monolith vs microservices

**Deliverables**:
- Architecture diagram (before/after)
- ADR explaining split decisions
- Working microservices
- Document trade-offs experienced

**Status**: Not Started
**GitHub**: (add link)

---

#### Project 2: Event-Driven System
**Requirements**:
- Build system with event-driven communication
- Use message broker (Kafka/RabbitMQ)
- Implement event sourcing for one domain
- Handle eventual consistency
- Add monitoring for event processing

**Deliverables**:
- Architecture diagram
- Event schema documentation
- Working implementation

**Status**: Not Started
**GitHub**: (add link)

---

#### Project 3: CQRS Implementation
**Requirements**:
- Implement CQRS pattern for specific use case
- Separate read and write models
- Use different databases for read/write if appropriate
- Handle synchronization between models
- Document when this pattern is useful vs overkill

**Deliverables**:
- Design document
- Implementation
- Performance comparison with simple CRUD

**Status**: Not Started
**GitHub**: (add link)

---

#### Project 4: ADR Practice
**Requirements**:
- Write ADRs for all major decisions in above projects
- Follow ADR template
- Include context, decision, consequences
- Review ADRs with peers if possible

**Status**: Not Started
**Location**: architecture-docs/adrs/

---

### Success Criteria

- [ ] Can explain when to use monolith vs microservices
- [ ] Understand event-driven architecture deeply
- [ ] Know CQRS and when NOT to use it
- [ ] Can write clear ADRs
- [ ] Can diagram architectures using standard notation
- [ ] Completed all pattern implementation projects
- [ ] Minimum 150 hours logged

---

## Module 2: Platform & Infrastructure Architecture (4-6 months, ~250 hours)

### Learning Objectives
- [ ] Design internal developer platforms
- [ ] Understand multi-tenancy patterns
- [ ] Deep Infrastructure as Code knowledge
- [ ] Know service mesh concepts
- [ ] Design for security (zero trust)
- [ ] Understand cost optimization

### Resources

#### Books
- [ ] "Team Topologies" by Skelton & Pais
  - Status: Not Started

#### Documentation
- [ ] CNCF Landscape - https://landscape.cncf.io/
- [ ] AWS Well-Architected Framework - https://aws.amazon.com/architecture/well-architected/
- [ ] Backstage docs - https://backstage.io/

### Projects

#### Project 1: Platform Design
**Requirements**:
- Design internal developer platform for team of 50+ engineers
- Include: API, CLI, self-service capabilities
- Document user journeys
- Design for multi-tenancy
- Include observability, secrets management, CI/CD

**Deliverables**:
- Comprehensive design document
- Architecture diagrams
- API specifications
- Security model

**Status**: Not Started
**Location**: architecture-docs/platform-design/

---

#### Project 2: Multi-Tenant System
**Requirements**:
- Design multi-tenant SaaS architecture
- Choose isolation strategy (database, schema, or row-level)
- Document trade-offs
- Implement prototype
- Include tenant provisioning, data isolation, billing hooks

**Deliverables**:
- Design document
- Prototype implementation
- Security analysis

**Status**: Not Started
**GitHub**: (add link)

---

#### Project 3: Zero Trust Security Implementation
**Requirements**:
- Implement zero trust security for one of your services
- Mutual TLS between services
- No implicit trust
- Policy-based access control
- Audit logging

**Deliverables**:
- Security architecture document
- Implementation
- Security audit report

**Status**: Not Started
**GitHub**: (add link)

---

#### Project 4: Observability Platform
**Requirements**:
- Set up complete observability stack
- Prometheus + Grafana for metrics
- Loki for logs
- Tempo for traces
- Create meaningful dashboards
- Set up alerts and SLOs

**Deliverables**:
- Architecture diagram
- Dashboard configurations
- Runbooks for common issues

**Status**: Not Started
**GitHub**: (add link)

---

#### Project 5: Cost Optimization Analysis
**Requirements**:
- Analyze costs for your infrastructure
- Identify optimization opportunities
- Calculate TCO for different approaches
- Present findings

**Deliverables**:
- Cost analysis document
- Optimization recommendations
- ROI calculations

**Status**: Not Started
**Location**: architecture-docs/cost-analysis/

---

### Success Criteria

- [ ] Can design platform for 50+ engineers
- [ ] Understand multi-tenancy deeply
- [ ] Can make build vs buy decisions with justification
- [ ] Know zero trust security principles
- [ ] Can design for cost efficiency
- [ ] Completed all platform projects
- [ ] Minimum 250 hours logged

---

## Module 3: System Design Practice (3-4 months, ~150 hours)

### Learning Objectives
- [ ] Gather requirements and constraints
- [ ] Do back-of-envelope calculations
- [ ] Design systems at scale
- [ ] Identify and handle failure modes
- [ ] Communicate designs clearly

### Resources

#### Books
- [ ] "System Design Interview" Vol 1 by Alex Xu
  - Status: Not Started
- [ ] "System Design Interview" Vol 2 by Alex Xu
  - Status: Not Started

#### Online
- [ ] System Design Primer - https://github.com/donnemartin/system-design-primer

### Practice Problems

For EACH problem below:
1. Write requirements document
2. Identify constraints (scale, latency, etc.)
3. Do capacity estimation
4. Create high-level design
5. Deep dive into 2-3 components
6. Identify failure modes
7. Explain trade-offs
8. Create architecture diagrams

#### Problem 1: URL Shortener
- [ ] Design document written
- [ ] Scale to 100M URLs/day
- [ ] Handle redirects with low latency

**Location**: system-designs/url-shortener/

---

#### Problem 2: Rate Limiter
- [ ] Design document written
- [ ] Multiple algorithms (token bucket, leaky bucket, fixed window)
- [ ] Distributed implementation

**Location**: system-designs/rate-limiter/

---

#### Problem 3: Notification System
- [ ] Design document written
- [ ] Email, SMS, push notifications
- [ ] Handle millions of notifications/day
- [ ] Prioritization and retry logic

**Location**: system-designs/notification-system/

---

#### Problem 4: Distributed Cache
- [ ] Design document written
- [ ] Consistent hashing
- [ ] Eviction policies
- [ ] Handle cache invalidation

**Location**: system-designs/distributed-cache/

---

#### Problem 5: API Gateway
- [ ] Design document written
- [ ] Routing, authentication, rate limiting
- [ ] High availability

**Location**: system-designs/api-gateway/

---

#### Problem 6: Logging System
- [ ] Design document written
- [ ] Ingest, store, query logs at scale
- [ ] Real-time and batch analysis

**Location**: system-designs/logging-system/

---

#### Problem 7: CI/CD Platform
- [ ] Design document written
- [ ] Handle thousands of builds/day
- [ ] Distributed build execution

**Location**: system-designs/cicd-platform/

---

#### Problem 8: Kubernetes Operator Platform
- [ ] Design document written
- [ ] Multi-tenant operator platform
- [ ] Relevant to your current work

**Location**: system-designs/k8s-operator-platform/

---

### Additional Practice
- [ ] Design 2+ systems from your company's domain
- [ ] Practice explaining designs in 45 minutes
- [ ] Get feedback on designs from senior engineers

### Success Criteria

- [ ] Completed all 8+ system designs
- [ ] Can design systems in 45-60 minutes
- [ ] Can explain trade-offs clearly
- [ ] Can do capacity estimation (back-of-envelope)
- [ ] Can identify failure modes
- [ ] Diagrams are clear and follow conventions
- [ ] Got feedback on at least 3 designs
- [ ] Minimum 150 hours logged

---

## Module 4: Networking & Security Deep Dive (2-3 months, ~100 hours)

### Learning Objectives
- [ ] Understand TCP/IP deeply
- [ ] Know DNS, load balancing, CDN
- [ ] Understand TLS/mTLS
- [ ] Design network architecture for K8s
- [ ] Know OWASP Top 10 and mitigations
- [ ] Implement zero trust principles

### Resources

#### Books
- [ ] "Computer Networking: A Top-Down Approach" by Kurose & Ross (Ch 1-6)
  - Status: Not Started

#### Online
- [ ] Cybrary Network+ free sections - https://www.cybrary.it/
- [ ] OWASP Top 10 - https://owasp.org/www-project-top-ten/

### Projects

#### Project 1: Build a Load Balancer
**Requirements**:
- Implement simple HTTP load balancer in Go
- Round-robin and least connections algorithms
- Health checks
- Handle failures gracefully

**Status**: Not Started
**GitHub**: (add link)

---

#### Project 2: mTLS Implementation
**Requirements**:
- Implement mutual TLS between services
- Certificate generation and rotation
- Handle certificate validation
- Document setup process

**Status**: Not Started
**GitHub**: (add link)

---

#### Project 3: Network Policies in K8s
**Requirements**:
- Set up K8s cluster with network policies
- Implement zero trust networking
- Only allow necessary service-to-service communication
- Document policies and reasoning

**Status**: Not Started
**Location**: architecture-docs/network-policies/

---

#### Project 4: Security Audit
**Requirements**:
- Audit one of your Phase 1 projects for security issues
- Check for OWASP Top 10 vulnerabilities
- Fix all findings
- Document findings and fixes

**Status**: Not Started
**Location**: architecture-docs/security-audit/

---

### Success Criteria

- [ ] Can troubleshoot network issues (DNS, routing, firewall)
- [ ] Understand TLS/mTLS deeply
- [ ] Can design secure network architecture
- [ ] Know OWASP Top 10 and mitigations
- [ ] Can perform security reviews
- [ ] Completed all networking/security projects
- [ ] Minimum 100 hours logged

---

## Phase 2 Overall Success Criteria

### Technical Skills
- [ ] Designed 10+ systems with full documentation
- [ ] Can explain architectural patterns and trade-offs
- [ ] Understand when NOT to use complex patterns
- [ ] Can create clear architecture diagrams
- [ ] Written 5+ ADRs for real decisions
- [ ] Deep knowledge in networking and security

### Communication Skills
- [ ] Written 6+ blog posts about architecture topics
- [ ] Presented at least 1 design to team/meetup
- [ ] Can explain technical concepts to non-technical audience
- [ ] Architecture docs are clear and actionable

### Metrics
- [ ] Minimum 500 hours logged
- [ ] All system designs documented
- [ ] Got feedback from architects/senior engineers
- [ ] Started applying learning at current job

### Mindset
- [ ] Think in trade-offs, not absolutes
- [ ] Consider business context in technical decisions
- [ ] Comfortable with ambiguity
- [ ] Know what you don't know

---

## If You Can't Check All Boxes Above

**DO NOT move to Phase 3.** You need this knowledge to do architecture work.

Fix gaps first, then proceed.

---

**Last Updated**: March 21, 2026
