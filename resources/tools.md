# Tools & References

Essential tools and references for your architect journey.

Last Updated: March 21, 2026

---

## Development Environment

### Go Development

#### Essential
- **Go**: https://go.dev/dl/
  - Install: Latest stable version (1.22+)
  - Verify: `go version`

- **VS Code** (recommended) or **GoLand**
  - VS Code Go extension: https://marketplace.visualstudio.com/items?itemName=golang.go
  - Configure: Auto-format on save, linting

- **golangci-lint**: https://golangci-lint.run/
  - Install: `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`
  - Comprehensive linting tool

#### Useful Tools
- **delve** (debugger): https://github.com/go-delve/delve
- **go-critic**: https://github.com/go-critic/go-critic
- **gofmt** / **goimports**: Built-in formatting tools

---

### Databases

#### Development
- **PostgreSQL**: https://www.postgresql.org/
  - Install locally or use Docker
  - GUI: pgAdmin or DBeaver

- **SQLite**: Built-in with Go, no installation needed

- **Docker**: For running databases locally
  - PostgreSQL: `docker run -p 5432:5432 -e POSTGRES_PASSWORD=postgres postgres`
  - Redis: `docker run -p 6379:6379 redis`

#### Database Tools
- **DBeaver**: https://dbeaver.io/ (Universal database GUI)
- **TablePlus**: https://tableplus.com/ (Mac/Windows)
- **DataGrip**: https://www.jetbrains.com/datagrip/ (JetBrains)

---

### API Development & Testing

- **Postman**: https://www.postman.com/
  - API testing and documentation

- **Insomnia**: https://insomnia.rest/
  - Alternative to Postman

- **curl**: Command-line HTTP client
  - Built-in on Mac/Linux

- **httpie**: https://httpie.io/
  - User-friendly HTTP client

---

### Message Queues

- **RabbitMQ**: https://www.rabbitmq.com/
  - Docker: `docker run -p 5672:5672 -p 15672:15672 rabbitmq:management`

- **Redis**: https://redis.io/
  - Can be used as message queue

- **Kafka**: https://kafka.apache.org/
  - For Phase 2 distributed systems work

---

## Observability Stack

### Logging
- **Structured logging libraries**:
  - Go: `github.com/sirupsen/logrus` or `go.uber.org/zap`

- **Log aggregation**:
  - Loki: https://grafana.com/oss/loki/
  - Docker: See Grafana documentation

### Metrics
- **Prometheus**: https://prometheus.io/
  - Time-series database for metrics
  - Docker: `docker run -p 9090:9090 prom/prometheus`

- **Go Prometheus client**: `github.com/prometheus/client_golang`

### Dashboards
- **Grafana**: https://grafana.com/
  - Visualization and dashboards
  - Docker: `docker run -p 3000:3000 grafana/grafana`

### Tracing
- **Jaeger**: https://www.jaegertracing.io/
  - Distributed tracing

- **Zipkin**: https://zipkin.io/
  - Alternative to Jaeger

---

## Infrastructure & DevOps

### Containers & Orchestration
- **Docker**: https://www.docker.com/
  - Essential for local development

- **Docker Compose**: https://docs.docker.com/compose/
  - Multi-container applications

- **Kubernetes** (for Phase 2):
  - **minikube**: Local K8s cluster
  - **kind**: Kubernetes in Docker
  - **k9s**: Terminal UI for K8s

### Infrastructure as Code
- **Terraform**: https://www.terraform.io/
  - You already have experience with this

- **Ansible**: https://www.ansible.com/
  - Configuration management (if needed)

---

## Architecture & Design

### Diagramming

#### Essential
- **draw.io** (diagrams.net): https://www.diagrams.net/
  - Free, powerful, works offline
  - **RECOMMENDED to start**

- **Excalidraw**: https://excalidraw.com/
  - Hand-drawn style diagrams
  - Great for brainstorming

#### Advanced
- **Lucidchart**: https://www.lucidchart.com/
  - Professional tool (paid)

- **Mermaid**: https://mermaid.js.org/
  - Diagrams as code
  - Good for version control

#### For Architecture Diagrams
- **C4 Model templates**: https://c4model.com/
  - Use with any diagramming tool
  - Standard notation for architecture

### Documentation
- **Markdown editors**:
  - VS Code (with preview)
  - Typora: https://typora.io/
  - Obsidian: https://obsidian.md/ (great for notes)

- **Documentation sites**:
  - MkDocs: https://www.mkdocs.org/
  - Docusaurus: https://docusaurus.io/

---

## Version Control & Collaboration

### Git
- **Git**: https://git-scm.com/
  - Essential

- **GitHub**: https://github.com/
  - For your projects and portfolio

- **Git GUI** (optional):
  - GitKraken: https://www.gitkraken.com/
  - SourceTree: https://www.sourcetreeapp.com/

### Code Review
- GitHub Pull Requests
- Read: https://google.github.io/eng-practices/review/

---

## Learning & Reference

### Go References
- **Go by Example**: https://gobyexample.com/
- **Go Playground**: https://go.dev/play/
  - Test code snippets

- **Effective Go**: https://go.dev/doc/effective_go
- **Go Wiki**: https://github.com/golang/go/wiki

### API References
- **HTTP Status Codes**: https://httpstatuses.com/
- **REST API Cheat Sheet**: https://github.com/RestCheatSheet/api-cheat-sheet

### System Design
- **System Design Primer**: https://github.com/donnemartin/system-design-primer
- **Latency Numbers**: https://gist.github.com/jboner/2841832
  - Every programmer should know these

### Engineering Blogs
- **Netflix TechBlog**: https://netflixtechblog.com/
- **Uber Engineering**: https://eng.uber.com/
- **Airbnb Engineering**: https://medium.com/airbnb-engineering
- **Martin Kleppmann**: https://martin.kleppmann.com/
- **High Scalability**: http://highscalability.com/

---

## Productivity & Time Tracking

### Time Tracking
- **Toggl**: https://toggl.com/ (Free tier available)
- **RescueTime**: https://www.rescuetime.com/
- **Simple spreadsheet**: Track manually

### Focus & Productivity
- **Pomodoro Timer**:
  - https://pomofocus.io/
  - Or use any timer app

- **Focus apps**:
  - Forest: https://www.forestapp.cc/
  - Cold Turkey (Windows): https://getcoldturkey.com/

### Note-Taking
- **Obsidian**: https://obsidian.md/
  - Markdown-based, great for learning
  - **RECOMMENDED**

- **Notion**: https://www.notion.so/
  - All-in-one workspace

- **Simple Markdown files**:
  - What we're using in this structure

---

## Blogging Platforms

### Free Options
- **dev.to**: https://dev.to/
  - Developer community
  - **RECOMMENDED to start**

- **Medium**: https://medium.com/
  - Large audience
  - Partner program for earnings

- **Hashnode**: https://hashnode.com/
  - Developer blogging platform

### Self-Hosted
- **Hugo**: https://gohugo.io/
  - Static site generator (Go-based!)

- **Jekyll**: https://jekyllrb.com/
  - GitHub Pages integration

---

## Online Playgrounds

### Code Testing
- **Go Playground**: https://go.dev/play/
- **SQLFiddle**: http://sqlfiddle.com/
- **DB Fiddle**: https://www.db-fiddle.com/

### System Design
- **Excalidraw**: For quick diagrams
- **draw.io**: For detailed architecture

---

## Communities

### Forums & Discussion
- **Reddit**:
  - r/golang
  - r/systemdesign
  - r/softwrearchitecture
  - r/devops

- **Stack Overflow**: https://stackoverflow.com/
  - For specific questions

- **Dev.to**: Community discussions

### Chat
- **Gophers Slack**: https://gophers.slack.com/
  - Go community

- **Software Architecture Discord**: (Search for current invite)

### Meetups
- **Meetup.com**: Find local tech meetups
  - Go meetups
  - DevOps meetups
  - Architecture meetups

---

## Quick Setup Checklist

### Week 1 Setup
- [ ] Install Go (latest version)
- [ ] Set up VS Code with Go extension
- [ ] Install Docker
- [ ] Create GitHub account (if you don't have one)
- [ ] Set up git configuration
- [ ] Install PostgreSQL (or via Docker)
- [ ] Install a Pomodoro timer
- [ ] Choose time tracking method
- [ ] Set up Obsidian or note-taking tool
- [ ] Bookmark all essential references

### Phase 1 Additional Setup
- [ ] Install golangci-lint
- [ ] Set up Postman or Insomnia
- [ ] Install Redis (via Docker)
- [ ] Set up RabbitMQ (via Docker)

### Phase 2 Additional Setup
- [ ] Set up Prometheus + Grafana
- [ ] Install Kubernetes (minikube or kind)
- [ ] Install k9s
- [ ] Choose diagramming tool (draw.io recommended)
- [ ] Set up blogging platform

---

## Tool Philosophy

### Keep It Simple
- Start with free, simple tools
- Don't over-invest in tools before you need them
- Master one tool before adding another
- Tools should help, not distract

### Essential vs. Nice-to-Have
**Essential**: Go, Docker, Git, VS Code, draw.io, Markdown
**Nice-to-Have**: Everything else

### When to Upgrade
- Upgrade when free tools become limiting
- Not before you've actually used the free version extensively
- Consider if complexity is worth the benefit

---

**Remember**:
- Tools are means to an end, not the end itself
- Focus on learning, not tool collection
- Master the basics before adding complexity
- The best tool is the one you'll actually use

---

**Last Updated**: March 21, 2026
**Review**: Update as you discover better tools or workflows
