# Week 6: Project 1.1.2 Completion + Testing & Best Practices Intro + Gophercises #9–10

**Week**: Week 6 of Phase 1 — Module 1.2 → Module 1.3 Transition (Concurrency → Testing)
**Dates**: April 27 – May 3, 2026
**Total Hours Goal**: 13 hours (completing Module 1.2, starting Module 1.3)
**Location**: `Year-2026/Month-03/WEEK-6-START-HERE.md`
**Roadmap Reference**: `roadmap/PHASE-1-MODULES.md` → Module 1.2 (ending) → Module 1.3 (Testing & Best Practices)

---

## ⚠️ Read This Before You Start

Week 6 is a transition week. You finish the hardest project so far (Todo REST API with concurrency and persistence), shift your mindset from "can I make it work?" to "can I prove it works?", and begin serious testing practices.

- **Weeks 1-5**: Building functionality, learning Go, writing first projects
- **Week 6+**: Building confidence through testing, code review, documentation

This week is about completing Project 1.1.2 with production-quality testing and beginning your deep dive into testing practices.

### Module 1.2 → 1.3 Progress Check

From `roadmap/PHASE-1-MODULES.md`:

| Requirement | Target | Your Status After Week 5 |
|---|---|---|
| Concurrency exercises completed | 5+ | 6+ (including advanced patterns) |
| Project 1.1.1 (Log Parser CLI) | Complete | ✅ Complete |
| Project 1.1.2 (Todo REST API) | Complete | ⏳ In progress (structure done) |
| TGPL Chapters 1–9 + exercises | Complete | ~80% (Chapters read, exercises ~70%) |
| Gophercises completed | 8+ | 8/20 |
| Testing patterns understood | Basic | 0% (starts this week) |
| Code organization | Growing | Good structure emerging |
| Hours logged | 200 hours total | ~87h (on pace) |

**Week 6 advances**: Project 1.1.2 full completion with comprehensive testing, testing best practices begin, code organization solidifies.

---

## Week 6 Goals

**Primary Goal**: Complete Project 1.1.2 (Todo REST API) with >85% test coverage, including unit tests, integration tests, and concurrent load testing. Begin mastering testing patterns and strategies.

**Secondary Goals**:
- Understand and implement: table-driven tests, subtests, test helpers, mocking
- Build comprehensive integration tests with real SQLite database
- Test concurrent API under load (multiple simultaneous requests)
- Complete Gophercises #9 and #10
- Understand test coverage: what % means, limitations, path to 95%+
- Master error testing: happy path, sad path, edge cases
- Begin thinking about code review and documentation quality
- Commit to GitHub daily with meaningful commit messages
- Log hours daily

**Hours Target**: 13 hours (heavy on Project 1.1.2 completion and testing)

**Reality Check**: This week is about quality. Project 1.1.2 will take most of the week. That is intentional.

---

## Key Resources for This Week

| Resource | Use For | Link |
|---|---|---|
| TGPL Chapter 11 | Testing (if available) | Your copy |
| Go Testing Packages | Standard library testing | https://pkg.go.dev/testing |
| Go by Example — Testing | Basic patterns | https://gobyexample.com/testing |
| Table-Driven Tests | Go testing standard | https://github.com/golang/go/wiki/TableDrivenTests |
| Testing Best Practices | Comprehensive guide | https://go.dev/blog/subtests |
| Mocking in Go | testify/mock, interfaces | https://github.com/stretchr/testify |
| Integration Testing | Database, HTTP, real dependencies | Your own patterns |
| pkg.go.dev — httptest | HTTP testing | https://pkg.go.dev/net/http/httptest |
| pkg.go.dev — testing/quick | Property-based testing | https://pkg.go.dev/testing/quick |

---

## Before You Write Any Code This Week

Check that all code from Week 5 passes quality checks:

```bash
# Most important: no race conditions
go test -race -v -cover ./...

# Standard quality checks
gofmt -l .
go vet ./...
```

All Week 5 code should still be passing. If not, fix it today.

---

## Detailed Daily Plan

---

### Sunday, April 27 (Pre-week setup: 45 min)

**Complete the Week 5 Review if not done**

If you have not completed the Week 5 Sunday reflection:
- [ ] Complete Week 5 journal entry
- [ ] Complete concurrency mastery reflection (draw systems, explain patterns)
- [ ] Fill in self-assessment on concurrency topics
- [ ] Gather your PLAN.md for Project 1.1.2 (created in Week 5)
- [ ] Update `roadmap/TIME-TRACKING.md` with Week 5 hours

Then spend 15 minutes on Week 6 setup:
- [ ] Read this entire plan
- [ ] Understand: Project 1.1.2 is a major deliverable
- [ ] Set expectations for testing scope (unit, integration, load tests)

---

### Monday, April 28 (Target: 2.5 hours)

**Focus: Project 1.1.2 — Todo REST API Core Implementation**

This week you build the actual API. Use your PLAN.md from Week 5 as the guide.

#### Session 1: API Implementation — Data Model & Store (90 min)

Start with the data layer. This ensures your concurrent access is safe from the start.

**Create**: `projects/phase-1/todo-rest-api/store.go`

```go
package main

import (
    "database/sql"
    "sync"
    "time"
    
    _ "github.com/mattn/go-sqlite3"
)

// Todo represents a todo item
type Todo struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Done      bool      `json:"done"`
    CreatedAt time.Time `json:"created_at"`
}

// Store manages todo persistence — thread-safe
type Store struct {
    db *sql.DB
    mu sync.RWMutex  // Protects concurrent access
}

// NewStore creates a new store with SQLite
func NewStore(dbPath string) (*Store, error) {
    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        return nil, err
    }
    
    // Create table if not exists
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS todos (
        id INTEGER PRIMARY KEY,
        title TEXT NOT NULL,
        done BOOLEAN DEFAULT FALSE,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    )`)
    if err != nil {
        return nil, err
    }
    
    return &Store{db: db}, nil
}

// Create adds a new todo
func (s *Store) Create(todo *Todo) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    result, err := s.db.Exec("INSERT INTO todos (title, done, created_at) VALUES (?, ?, ?)",
        todo.Title, todo.Done, time.Now())
    if err != nil {
        return err
    }
    
    id, err := result.LastInsertId()
    todo.ID = int(id)
    return err
}

// GetByID retrieves a todo by ID
func (s *Store) GetByID(id int) (*Todo, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    todo := &Todo{}
    err := s.db.QueryRow("SELECT id, title, done, created_at FROM todos WHERE id = ?", id).
        Scan(&todo.ID, &todo.Title, &todo.Done, &todo.CreatedAt)
    if err == sql.ErrNoRows {
        return nil, ErrNotFound
    }
    return todo, err
}

// GetAll retrieves all todos
func (s *Store) GetAll() ([]Todo, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    rows, err := s.db.Query("SELECT id, title, done, created_at FROM todos")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var todos []Todo
    for rows.Next() {
        var todo Todo
        if err := rows.Scan(&todo.ID, &todo.Title, &todo.Done, &todo.CreatedAt); err != nil {
            return nil, err
        }
        todos = append(todos, todo)
    }
    return todos, rows.Err()
}

// Update modifies an existing todo
func (s *Store) Update(todo *Todo) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    result, err := s.db.Exec("UPDATE todos SET title = ?, done = ? WHERE id = ?",
        todo.Title, todo.Done, todo.ID)
    if err != nil {
        return err
    }
    
    rows, err := result.RowsAffected()
    if rows == 0 {
        return ErrNotFound
    }
    return err
}

// Delete removes a todo
func (s *Store) Delete(id int) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    result, err := s.db.Exec("DELETE FROM todos WHERE id = ?", id)
    if err != nil {
        return err
    }
    
    rows, err := result.RowsAffected()
    if rows == 0 {
        return ErrNotFound
    }
    return err
}

// Close closes the database connection
func (s *Store) Close() error {
    return s.db.Close()
}
```

**Key Points**:
- [ ] Store has a mutex protecting database access
- [ ] Methods use RLock() for reads (multiple readers allowed)
- [ ] Methods use Lock() for writes (exclusive access)
- [ ] Proper error handling (ErrNotFound for missing items)
- [ ] Validation before database operations

**Create**: `projects/phase-1/todo-rest-api/errors.go`

```go
package main

import "errors"

var (
    ErrNotFound       = errors.New("todo not found")
    ErrInvalidTitle   = errors.New("title cannot be empty")
    ErrInvalidRequest = errors.New("invalid request")
)
```

**Create**: `projects/phase-1/todo-rest-api/main.go` (basic structure)

```go
package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "strings"
)

var (
    dbPath = flag.String("db", "todos.db", "Path to SQLite database")
    port   = flag.String("port", ":8080", "Server port")
)

func main() {
    flag.Parse()
    
    store, err := NewStore(*dbPath)
    if err != nil {
        log.Fatalf("Failed to init store: %v", err)
    }
    defer store.Close()
    
    mux := http.NewServeMux()
    
    // Routes
    mux.HandleFunc("/todos", handleTodos(store))
    mux.HandleFunc("/todos/", handleTodoByID(store))
    
    log.Printf("Starting server on %s", *port)
    log.Fatal(http.ListenAndServe(*port, mux))
}

// Handler functions will go in handler.go
```

**Actual time spent**: ___ hours
**What I accomplished**:

---

#### Session 2: Gophercises #9 Start (60 min)

URL: https://gophercises.com/ → Exercise 9

- [ ] Read the full problem
- [ ] Plan before coding
- [ ] Start implementing (don't need to finish today)
- [ ] Write 1 test

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Tuesday, April 29 (Target: 2 hours)

**Focus: Project 1.1.2 — HTTP Handlers + API Endpoints**

#### Session 1: Implement HTTP Handlers (90 min)

**Create**: `projects/phase-1/todo-rest-api/handler.go`

```go
package main

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"
)

// handleTodos handles GET (list) and POST (create)
func handleTodos(store *Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            getTodos(store, w, r)
        case http.MethodPost:
            createTodo(store, w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    }
}

// handleTodoByID handles GET (retrieve), PUT (update), DELETE
func handleTodoByID(store *Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Path: /todos/:id — extract ID
        parts := strings.Split(r.URL.Path, "/")
        if len(parts) < 3 {
            http.Error(w, "Invalid path", http.StatusBadRequest)
            return
        }
        
        id, err := strconv.Atoi(parts[2])
        if err != nil {
            http.Error(w, "Invalid ID", http.StatusBadRequest)
            return
        }
        
        switch r.Method {
        case http.MethodGet:
            getTodoByID(store, w, r, id)
        case http.MethodPut:
            updateTodo(store, w, r, id)
        case http.MethodDelete:
            deleteTodo(store, w, r, id)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    }
}

// Endpoint implementations
func getTodos(store *Store, w http.ResponseWriter, r *http.Request) {
    todos, err := store.GetAll()
    if err != nil {
        http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
}

func createTodo(store *Store, w http.ResponseWriter, r *http.Request) {
    var todo Todo
    if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
        http.Error(w, ErrInvalidRequest.Error(), http.StatusBadRequest)
        return
    }
    
    // Validate
    if strings.TrimSpace(todo.Title) == "" {
        http.Error(w, ErrInvalidTitle.Error(), http.StatusBadRequest)
        return
    }
    
    if err := store.Create(&todo); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(todo)
}

func getTodoByID(store *Store, w http.ResponseWriter, r *http.Request, id int) {
    todo, err := store.GetByID(id)
    if err == ErrNotFound {
        http.Error(w, "Todo not found", http.StatusNotFound)
        return
    }
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todo)
}

func updateTodo(store *Store, w http.ResponseWriter, r *http.Request, id int) {
    var todo Todo
    if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
        http.Error(w, ErrInvalidRequest.Error(), http.StatusBadRequest)
        return
    }
    todo.ID = id
    
    if strings.TrimSpace(todo.Title) == "" {
        http.Error(w, ErrInvalidTitle.Error(), http.StatusBadRequest)
        return
    }
    
    if err := store.Update(&todo); err == ErrNotFound {
        http.Error(w, "Todo not found", http.StatusNotFound)
        return
    } else if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todo)
}

func deleteTodo(store *Store, w http.ResponseWriter, r *http.Request, id int) {
    if err := store.Delete(id); err == ErrNotFound {
        http.Error(w, "Todo not found", http.StatusNotFound)
        return
    } else if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusNoContent)
}
```

#### Session 2: Finish Gophercises #9 (30 min)

- [ ] Complete Exercise 9
- [ ] Write tests
- [ ] Push to GitHub

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Wednesday, April 30 (Target: 2.5 hours)

**Focus: Project 1.1.2 — Testing Strategy & Unit Tests**

#### Session 1: Store Unit Tests (60 min)

**Create**: `projects/phase-1/todo-rest-api/store_test.go`

```go
package main

import (
    "os"
    "testing"
    "time"
)

func setupTest(t *testing.T) *Store {
    // Use in-memory sqlite for tests
    store, err := NewStore(":memory:")
    if err != nil {
        t.Fatalf("Failed to setup store: %v", err)
    }
    t.Cleanup(func() {
        store.Close()
    })
    return store
}

func TestCreateTodo(t *testing.T) {
    store := setupTest(t)
    
    tests := []struct {
        name      string
        todo      Todo
        wantErr   bool
    }{
        {
            name:    "valid todo",
            todo:    Todo{Title: "Buy milk", Done: false},
            wantErr: false,
        },
        {
            name:    "empty title",
            todo:    Todo{Title: "", Done: false},
            wantErr: false, // Store doesn't validate; handler does
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := store.Create(&tt.todo)
            if (err != nil) != tt.wantErr {
                t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
            }
            if err == nil && tt.todo.ID == 0 {
                t.Error("Create() did not set ID")
            }
        })
    }
}

func TestGetTodo(t *testing.T) {
    store := setupTest(t)
    
    // Setup: create a todo
    original := Todo{Title: "Test", Done: false}
    store.Create(&original)
    
    tests := []struct {
        name    string
        id      int
        wantErr bool
    }{
        {"existing", original.ID, false},
        {"not found", 999, true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            todo, err := store.GetByID(tt.id)
            if (err != nil) != tt.wantErr {
                t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
            }
            if err == nil && todo.Title != original.Title {
                t.Errorf("GetByID() = %v, want %v", todo.Title, original.Title)
            }
        })
    }
}

func TestUpdateTodo(t *testing.T) {
    store := setupTest(t)
    
    original := Todo{Title: "Original", Done: false}
    store.Create(&original)
    
    updated := original
    updated.Title = "Updated"
    updated.Done = true
    
    if err := store.Update(&updated); err != nil {
        t.Fatalf("Update() error = %v", err)
    }
    
    // Verify
    retrieved, _ := store.GetByID(updated.ID)
    if retrieved.Title != "Updated" || !retrieved.Done {
        t.Error("Update() did not persist changes")
    }
}

func TestDeleteTodo(t *testing.T) {
    store := setupTest(t)
    
    todo := Todo{Title: "To delete", Done: false}
    store.Create(&todo)
    
    // Delete
    if err := store.Delete(todo.ID); err != nil {
        t.Fatalf("Delete() error = %v", err)
    }
    
    // Verify deleted
    _, err := store.GetByID(todo.ID)
    if err != ErrNotFound {
        t.Error("Delete() did not remove todo")
    }
}

func TestConcurrentAccess(t *testing.T) {
    store := setupTest(t)
    
    // Verify store handles concurrent writes safely
    for i := 0; i < 10; i++ {
        go func(n int) {
            store.Create(&Todo{Title: "Concurrent " + string(rune(n)), Done: false})
        }(i)
    }
    
    // Run with -race: go test -race ./...
}
```

#### Session 2: HTTP Handler Tests (60 min)

**Create**: `projects/phase-1/todo-rest-api/handler_test.go`

```go
package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func setupStoreForTest(t *testing.T) *Store {
    store, err := NewStore(":memory:")
    if err != nil {
        t.Fatalf("Failed to setup store: %v", err)
    }
    t.Cleanup(func() { store.Close() })
    return store
}

func TestGetTodos(t *testing.T) {
    store := setupStoreForTest(t)
    
    // Setup: create some todos
    store.Create(&Todo{Title: "Todo 1", Done: false})
    store.Create(&Todo{Title: "Todo 2", Done: true})
    
    req := httptest.NewRequest(http.MethodGet, "/todos", nil)
    w := httptest.NewRecorder()
    
    getTodos(store, w, req)
    
    if w.Code != http.StatusOK {
        t.Errorf("GET /todos status = %d, want %d", w.Code, http.StatusOK)
    }
    
    var todos []Todo
    if err := json.NewDecoder(w.Body).Decode(&todos); err != nil {
        t.Fatalf("Failed to decode response: %v", err)
    }
    
    if len(todos) != 2 {
        t.Errorf("Expected 2 todos, got %d", len(todos))
    }
}

func TestCreateTodoHandler(t *testing.T) {
    store := setupStoreForTest(t)
    
    tests := []struct {
        name       string
        body       interface{}
        wantStatus int
    }{
        {
            name:       "valid",
            body:       map[string]string{"title": "New todo"},
            wantStatus: http.StatusCreated,
        },
        {
            name:       "empty title",
            body:       map[string]string{"title": ""},
            wantStatus: http.StatusBadRequest,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            body, _ := json.Marshal(tt.body)
            req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))
            w := httptest.NewRecorder()
            
            createTodo(store, w, req)
            
            if w.Code != tt.wantStatus {
                t.Errorf("POST /todos status = %d, want %d", w.Code, tt.wantStatus)
            }
        })
    }
}

func TestGetTodoByIDHandler(t *testing.T) {
    store := setupStoreForTest(t)
    
    todo := Todo{Title: "Get me", Done: false}
    store.Create(&todo)
    
    tests := []struct {
        name       string
        id         int
        wantStatus int
    }{
        {"existing", todo.ID, http.StatusOK},
        {"not found", 999, http.StatusNotFound},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req := httptest.NewRequest(http.MethodGet, "/todos/", nil)
            w := httptest.NewRecorder()
            
            getTodoByID(store, w, req, tt.id)
            
            if w.Code != tt.wantStatus {
                t.Errorf("GET /todos/:id status = %d, want %d", w.Code, tt.wantStatus)
            }
        })
    }
}
```

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Thursday, May 1 (Target: 2 hours)

**Focus: Project 1.1.2 — Integration Tests + Concurrent Load Testing**

#### Session 1: Integration Tests (60 min)

**Create**: `projects/phase-1/todo-rest-api/integration_test.go`

```go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestEndToEndTodoAPI(t *testing.T) {
    store := setupStoreForTest(t)
    
    // Create handler
    mux := http.NewServeMux()
    mux.HandleFunc("/todos", handleTodos(store))
    mux.HandleFunc("/todos/", handleTodoByID(store))
    
    // Create a todo
    createBody, _ := json.Marshal(map[string]interface{}{
        "title": "Integration test todo",
        "done":  false,
    })
    
    req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(createBody))
    w := httptest.NewRecorder()
    mux.ServeHTTP(w, req)
    
    if w.Code != http.StatusCreated {
        t.Fatalf("Create failed: %d", w.Code)
    }
    
    var createdTodo Todo
    json.NewDecoder(w.Body).Decode(&createdTodo)
    
    // Get the todo
    getReq := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/todos/%d", createdTodo.ID), nil)
    w = httptest.NewRecorder()
    mux.ServeHTTP(getReq, w)
    
    if w.Code != http.StatusOK {
        t.Fatalf("Get failed: %d", w.Code)
    }
    
    var retrieved Todo
    json.NewDecoder(w.Body).Decode(&retrieved)
    if retrieved.Title != "Integration test todo" {
        t.Error("Retrieved todo does not match")
    }
    
    // Update the todo
    updateBody, _ := json.Marshal(map[string]interface{}{
        "title": "Updated title",
        "done":  true,
    })
    updateReq := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/todos/%d", createdTodo.ID), bytes.NewReader(updateBody))
    w = httptest.NewRecorder()
    mux.ServeHTTP(updateReq, w)
    
    if w.Code != http.StatusOK {
        t.Fatalf("Update failed: %d", w.Code)
    }
    
    // Delete the todo
    delReq := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/todos/%d", createdTodo.ID), nil)
    w = httptest.NewRecorder()
    mux.ServeHTTP(delReq, w)
    
    if w.Code != http.StatusNoContent {
        t.Fatalf("Delete failed: %d", w.Code)
    }
    
    // Verify deleted
    getReq = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/todos/%d", createdTodo.ID), nil)
    w = httptest.NewRecorder()
    mux.ServeHTTP(getReq, w)
    
    if w.Code != http.StatusNotFound {
        t.Error("Todo should be deleted")
    }
}

func TestConcurrentRequests(t *testing.T) {
    store := setupStoreForTest(t)
    mux := http.NewServeMux()
    mux.HandleFunc("/todos", handleTodos(store))
    
    // Send 50 concurrent POST requests
    done := make(chan bool, 50)
    errors := make(chan error, 50)
    
    for i := 0; i < 50; i++ {
        go func(id int) {
            body, _ := json.Marshal(map[string]string{
                "title": fmt.Sprintf("Concurrent todo %d", id),
            })
            req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))
            w := httptest.NewRecorder()
            mux.ServeHTTP(w, req)
            
            if w.Code != http.StatusCreated {
                select {
                case errors <- fmt.Errorf("request %d failed: %d", id, w.Code):
                default:
                }
            }
            done <- true
        }(i)
    }
    
    // Wait for all
    for i := 0; i < 50; i++ {
        <-done
    }
    
    close(errors)
    if len(errors) > 0 {
        t.Fatalf("Concurrent requests failed")
    }
    
    // Verify all were created
    req := httptest.NewRequest(http.MethodGet, "/todos", nil)
    w := httptest.NewRecorder()
    mux.ServeHTTP(w, req)
    
    var todos []Todo
    json.NewDecoder(w.Body).Decode(&todos)
    if len(todos) != 50 {
        t.Errorf("Expected 50 todos, got %d", len(todos))
    }
}
```

#### Session 2: Finish Gophercises #10 (60 min)

URL: https://gophercises.com/ → Exercise 10

- [ ] Read the full problem
- [ ] Plan before coding
- [ ] Build complete solution
- [ ] Write comprehensive tests
- [ ] Run `go test -race -v -cover ./...`
- [ ] Push to GitHub

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Friday, May 2 (Target: 1.5 hours)

**Focus: Project 1.1.2 — Code Quality & Documentation**

#### Session 1: Code Quality Audit (45 min)

Run on ALL Project 1.1.2 code:

```bash
# Format
gofmt -w ./projects/phase-1/todo-rest-api/

# Vet
go vet ./projects/phase-1/todo-rest-api/...

# Race detector (critical!)
go test -race -v -cover ./projects/phase-1/todo-rest-api/...

# Run multiple times to catch intermittent races
for i in {1..5}; do go test -race ./projects/phase-1/todo-rest-api/... || break; done
```

- [ ] All files pass `gofmt`
- [ ] All files pass `go vet`
- [ ] All tests pass with race detector **multiple times**
- [ ] Coverage: target >85%
- [ ] No dead code

#### Session 2: README & Documentation (45 min)

**Create**: `projects/phase-1/todo-rest-api/README.md`

```markdown
# Todo REST API

A production-ready REST API for managing todos, built with Go and SQLite.

## Features

- REST CRUD operations for todos
- SQLite persistence
- Concurrent request handling (thread-safe)
- Comprehensive test suite (>85% coverage)
- Graceful error handling

## Building

```bash
go build -o todo-api ./projects/phase-1/todo-rest-api/
```

## Running

```bash
# Start server on default port 8080
./todo-api

# Or specify custom port and database path
./todo-api -port :3000 -db /tmp/todos.db
```

## API Endpoints

### List todos
```bash
GET /todos
```

### Get single todo
```bash
GET /todos/:id
```

### Create todo
```bash
POST /todos
Content-Type: application/json

{"title": "Buy milk", "done": false}
```

### Update todo
```bash
PUT /todos/:id
Content-Type: application/json

{"title": "Buy milk and eggs", "done": true}
```

### Delete todo
```bash
DELETE /todos/:id
```

## Concurrency

The API is fully thread-safe. All database operations use sync.RWMutex:
- Multiple concurrent reads allowed
- Exclusive write access (maintains consistency)
- No goroutine leaks
- Passes `go test -race`

## Testing

Run all tests:
```bash
go test -v ./projects/phase-1/todo-rest-api/...
```

Run with race detector:
```bash
go test -race ./projects/phase-1/todo-rest-api/...
```

Coverage:
```bash
go test -cover ./projects/phase-1/todo-rest-api/...
```

## Architecture

- `main.go` — Server setup and routing
- `handler.go` — HTTP endpoint handlers
- `store.go` — SQLite persistence (thread-safe with mutex)
- `errors.go` — Error definitions
- `*_test.go` — Unit, handler, and integration tests
```

**Also create**: `projects/phase-1/todo-rest-api/DESIGN.md`

```markdown
# Todo API — Design Decisions

## Concurrency Strategy

The API uses sync.RWMutex to protect database access:

**Why RWMutex instead of regular Mutex?**
- GET /todos and GET /todos/:id are read operations
- Multiple concurrent reads are allowed (improves throughput)
- POST/PUT/DELETE are write operations (exclusive access)
- This pattern is common in databases and caches

**Why not connection pooling?**
- MVP doesn't need it (SQLite is single-file)
- RWMutex is sufficient for concurrent request handling
- Future: Could upgrade to postgresql + connection pool

## Error Handling

Every error scenario returns a meaningful HTTP status code:
- 400 Bad Request — Invalid input (empty title, malformed JSON)
- 404 Not Found — Todo doesn't exist
- 500 Internal Server Error — Database error

Structured errors in `errors.go` prevent duplicated error messages.

## Testing Strategy

Three levels of testing:

1. **Unit Tests** (store_test.go)
   - Test database operations in isolation
   - Use in-memory SQLite (:memory:)
   - Fast execution
   
2. **Handler Tests** (handler_test.go)
   - Test HTTP handlers with mock requests
   - Use httptest.NewRequest/NewRecorder
   - Verify status codes and response format
   
3. **Integration Tests** (integration_test.go)
   - End-to-end workflow testing
   - Concurrent request testing (verify thread-safety)
   - Real ServeHTTP behavior

**Coverage Target**: >85%

All tests pass `go test -race` with zero data race warnings.

## Future Improvements

1. **API Key Authentication** — Add -api-key flag
2. **Request Validation** — Structured validation middleware
3. **Logging** — Structured logging with timestamps
4. **Pagination** — Support limit/offset on list endpoint
5. **Filtering** — Filter by done status, created date range
6. **Database** — Migrate from SQLite to PostgreSQL
```

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Saturday, May 3 (Target: 2.5 hours)

**Focus: Testing Deep Dive — Table-Driven Tests & Mocking Patterns**

#### Session 1: Advanced Testing Patterns Study (90 min)

This session is about learning testing best practices that will level up ALL future code.

**Read and Study**:
- [ ] Go Blog — Subtests and Table-Driven Tests: https://go.dev/blog/subtests
- [ ] Uber Go Style Guide — Testing section: https://github.com/uber-go/guide/blob/master/style.md#testing
- [ ] Table-Driven Tests: https://github.com/golang/go/wiki/TableDrivenTests

**Create**: `notes/books/testing-best-practices.md`

Write your own understanding of:
1. Table-driven tests — why, when, how
2. Subtests using t.Run()
3. Test helpers and fixtures
4. Mocking strategies in Go
5. Assertions libraries (testify) — when to use
6. Code coverage metrics and their limitations
7. Benchmarking (go test -bench)
8. Property-based testing (testing/quick)

**Implement Examples**:

Create `projects/phase-1/testing-patterns/table_driven_test.go`:

```go
package patterns

import "testing"

func TestTableDriven(t *testing.T) {
    // THIS is the Go standard for testing
    tests := []struct {
        name      string
        input     int
        want      int
        wantErr   bool
    }{
        {name: "positive", input: 5, want: 5, wantErr: false},
        {name: "zero", input: 0, want: 0, wantErr: false},
        {name: "negative", input: -5, want: -5, wantErr: false},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := DoSomething(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("wantErr %v; got error %v", tt.wantErr, err)
            }
            if got != tt.want {
                t.Errorf("got %d, want %d", got, tt.want)
            }
        })
    }
}
```

Create `projects/phase-1/testing-patterns/mocking_test.go`:

```go
package patterns

import "testing"

// MockUserStore implements UserStore for testing
type MockUserStore struct {
    users map[int]User
    calls int
}

func (m *MockUserStore) GetUser(id int) (User, error) {
    m.calls++
    return m.users[id], nil
}

func TestServiceWithMock(t *testing.T) {
    mockStore := &MockUserStore{
        users: map[int]User{
            1: {ID: 1, Name: "Alice"},
        },
    }
    
    service := NewService(mockStore)
    user, _ := service.GetUser(1)
    
    if user.Name != "Alice" {
        t.Error("Expected Alice")
    }
    if mockStore.calls != 1 {
        t.Errorf("Expected 1 call, got %d", mockStore.calls)
    }
}
```

#### Session 2: Project 1.1.2 — Final Polish & Verification (60 min)

Complete final verification of Project 1.1.2:

- [ ] Run full test suite: `go test -v ./projects/phase-1/todo-rest-api/...`
- [ ] Verify coverage: `go test -cover ./projects/phase-1/todo-rest-api/...` (should be >85%)
- [ ] Race detector: `go test -race ./projects/phase-1/todo-rest-api/...` (run 3x)
- [ ] Manual testing: Start server, test endpoints with curl
- [ ] README complete with examples
- [ ] DESIGN.md explains architecture decisions
- [ ] All code committed with meaningful messages
- [ ] All files pass `gofmt` and `go vet`

**Push to GitHub**:
```bash
cd projects/phase-1/todo-rest-api/
# Run final checks
go test -race -v -cover ./...
git add .
git commit -m "feat: todo rest api - complete implementation with >85% test coverage"
git push
```

**Actual time spent**: ___ hours
**What I accomplished**:

---

### Sunday, May 4 (Target: 1 hour)

**Focus: Weekly Review + Plan Week 7**

#### Part 1: Project 1.1.2 Retrospective (20 min)

Answer these questions about Project 1.1.2:

- [ ] What was the hardest part of building this API?
- [ ] How did concurrency concerns affect your design?
- [ ] If you had to build it again, what would you do differently?
- [ ] How confident are you that this code works correctly under load?
- [ ] Would you show this code in a job interview? What parts are you proud of?

Write answers in your journal.

#### Part 2: Testing & Quality Reflection (15 min)

- [ ] How did testing this project differ from testing Week 1-2 code?
- [ ] What testing patterns are you now confident using?
- [ ] Which testing pattern is still confusing? (Plan to revisit)
- [ ] How useful was `go test -race`? When did it catch bugs?
- [ ] What's the relationship between test coverage percentage and actual correctness?

Write your reflections.

#### Part 3: Journal Entry (15 min)

Create your Week 6 journal:
```bash
cp journal/JOURNAL-TEMPLATE.md journal/2026-05-week-18.md
```

**Technical Concepts section**:
- "Mutex patterns: RWMutex vs regular Mutex. RWMutex allows concurrent reads but exclusive writes. This is fundamental to safe concurrent systems."
- "Table-driven tests are the Go standard. I finally understand why: clear test cases in a data structure, easy to add new cases, good error messages."
- "Integration testing is different from unit testing. It tests the entire system (handlers + store + database) together. This is what catches real bugs."
- "go test -race is powerful. It found race conditions I would never have thought of manually."

**Struggles**:
- "Concurrent request testing was complex. Had to use goroutines and channels to track completion."
- "SQLite concurrency was subtle. The RWMutex protects Go-level access, but SQLite has its own locking."

#### Part 4: Week 7 Planning (10 min)

Week 7 target (from PHASE-1-MODULES.md):
- Continue testing patterns (Module 1.3)
- Code organization and project structure
- Gophercises #11-12
- Optional: Project 1.1.3 (Web Crawler) start

Write 5 goals in `roadmap/weekly-goals.md`.

**Actual time spent**: ___ hours
**What I accomplished**:

---

## Week 6 Summary

### Total Time Logged
**Planned**: 13 hours
**Actual**: ___ hours
**Variance**: ___ hours (over/under)
**Running Total (Phase 1)**: ___ / 200 hours (should be ~100h by end of Week 6)

### Deliverables Check

- [ ] Project 1.1.2 (Todo REST API) — COMPLETE
  - [ ] All endpoints working (GET/POST/PUT/DELETE)
  - [ ] SQLite persistence
  - [ ] Thread-safe concurrent access
  - [ ] >85% test coverage
  - [ ] Zero race conditions (go test -race passes 3x)
- [ ] Comprehensive test suite:
  - [ ] Unit tests (store_test.go)
  - [ ] Handler tests (handler_test.go)
  - [ ] Integration tests (integration_test.go)
  - [ ] Concurrent load tests
- [ ] Gophercises #9 — complete + tests + pushed
- [ ] Gophercises #10 — complete + tests + pushed
- [ ] README with usage examples
- [ ] DESIGN.md with architecture decisions
- [ ] All code passes `gofmt`, `go vet`, `go test -race`
- [ ] Testing best practices documented in notes
- [ ] All code committed to GitHub with meaningful messages
- [ ] Week 6 journal filled in with project retrospective
- [ ] Week 7 goals written in `roadmap/weekly-goals.md`
- [ ] Running hours total updated in `roadmap/TIME-TRACKING.md`

### Gophercises Tracker (updated)

| # | Exercise | Status | Tests Written | Coverage | Race-Free |
|---|---|---|---|---|---|
| 1–8 | [Previous] | Complete | Yes | >75% | Yes |
| 9 | [API-related] | Complete | Yes | ___ | Yes |
| 10 | [API-related] | Complete | Yes | ___ | Yes |
| 11 | | Not started | | | |
| ... | | | | | |
| 20 | | Not started | | | |

### Projects Completed

| Project | Status | Coverage | Lines of Code | URL |
|---|---|---|---|---|
| Week 2 Mini-Project | Complete | >80% | ~200 | projects/phase-1/week-2-mini-project/ |
| Week 3 Mini-Project | Complete | >80% | ~250 | projects/phase-1/week-3-mini-project/ |
| Project 1.1.1: Log Parser CLI | Complete | >80% | ~400 | projects/phase-1/log-parser-cli/ |
| Week 4 Concurrent Exercise | Complete | >80% | ~350 | projects/phase-1/week-4-concurrent-exercise/ |
| Week 5 Advanced Pattern | Complete | >80% | ~450 | projects/phase-1/week-5-advanced-pattern/ |
| **Project 1.1.2: Todo REST API** | **Complete** | **>85%** | **~600** | **projects/phase-1/todo-rest-api/** |

---

## Concepts You Must Be Able to Explain by End of Week 6

Mastery-level questions. Answer without looking anything up.

**Testing Principles**
1. What is a table-driven test and why does Go favor them?
2. What is the difference between a unit test, integration test, and end-to-end test?
3. What does test coverage % actually tell you? What does it NOT tell you?
4. Why might 100% test coverage still miss bugs?
5. How do you test for race conditions? Why is `go test -race` important?

**Testing Patterns**
6. How do you use t.Run() (subtests) and why would you?
7. What is a test fixture and how do you implement one in Go?
8. How do you mock a dependency in Go? Give an example.
9. What is the interface-based testing pattern and why is it powerful?

**SQL & Data Access**
10. Why use RWMutex instead of regular Mutex for database access?
11. How do you handle concurrent writes to a single-file database safely?
12. What happens when two goroutines try to write to SQLite simultaneously?

**HTTP & REST**
13. How do you test HTTP handlers? Explain httptest.
14. What are appropriate HTTP status codes for: created, bad request, not found, server error?
15. How do you handle JSON encoding/decoding in Go?

**Project Architecture**
16. Design a data access layer that is thread-safe and testable.
17. Explain: Why separate handler.go, store.go, and errors.go?
18. How would you add database migrations to this API?

**Advanced**
19. Your concurrent API has 50% test coverage but passes all tests. Is it production-ready? Why/why not?
20. You see one race condition warning from `go test -race`. What must you do before shipping?

---

## Code Quality Standards — Enhanced for Testing

| Standard | Requirement | How to Check |
|---|---|---|
| Formatting | All files formatted | `gofmt -l .` returns nothing |
| Vet | No vet warnings | `go vet ./...` returns nothing |
| **Race-free** | **ZERO data races** | **`go test -race ./...` multiple runs** |
| **Test coverage** | **>85% for projects** | **`go test -cover ./...`** |
| Tests runnable | All tests pass | `go test -v ./...` |
| Error handling | Every error handled | Manual review |
| Docstrings | Functions documented | `go doc ./...` |
| Integration tests | End-to-end workflows tested | `*_integration_test.go` exists |
| README | Clear usage instructions | Present and complete |
| DESIGN.md | Architecture decisions explained | Present |

**NEW**: Every test file should have a setup function testing package scope uses:

```go
// helper to set up test fixtures
func setupTestDB(t *testing.T) *Store {
    store, err := NewStore(":memory:")
    if err != nil {
        t.Fatalf("setup failed: %v", err)
    }
    t.Cleanup(func() {
        store.Close()
    })
    return store
}
```

---

## If You Are Behind

**Project 1.1.2 taking too long?**
→ No problem. API design is complex.
→ Move incomplete parts to Week 7.
→ Focus on core endpoints + basic tests first.

**Test coverage below 85%?**
→ Identify what's not tested (run `go test -cover` with verbose output).
→ Write tests for missing paths (focus on error cases).
→ Understand: 95% coverage is realistic, 100% often isn't worth the effort.

**Race conditions still appearing?**
→ This is good news — you're finding them early.
→ Every race condition is a bug you would have found in production.
→ Use print statements to trace which goroutine accesses what when.
→ Consider: Do I need a Mutex here? Am I closing channels properly?

**Overwhelmed by testing patterns?**
→ Focus on table-driven tests + subtests first.
→ Mocking is secondary. Core patterns first.
→ You don't need testify or complex mocking frameworks yet.

**Gophercises #9 or #10 very hard?**
→ You're 50% through. These are the harder ones.
→ Break them into smaller parts.
→ Get basic version working, then add features.

---

## Next Week: Module 1.3 Launches — Testing & Organization

Week 7 focuses on:
- Code organization and package structure
- More testing patterns (benchmarks, property-based testing)
- Refactoring for testability
- Preparing for Project 1.1.3 (Web Crawler)

**Real soon now**: Concurrency comes together with testing. By Week 8, you will be able to build complex concurrent systems with high confidence.

You've built 6 projects in 6 weeks. This is real progress. See you in Week 7.
