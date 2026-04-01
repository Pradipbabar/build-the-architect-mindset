# 📘 Chapter 4: Composite Types (Go)

## 🔹 1. Introduction to Composite Types

Composite types are **data structures formed by combining simpler types**. They allow storing and organizing collections of data efficiently.

### Types covered:

* Arrays
* Slices
* Maps
* Structs
* JSON handling
* Text & HTML Templates ([Bookey][1])

### Classification:

* **Fixed-size types** → Arrays, Structs
* **Dynamic types** → Slices, Maps ([notes.shichao.io][2])

---

# 🔹 2. Arrays

## 📌 Definition

An array is a **fixed-length sequence of elements of the same type**.

```go
var a [3]int
```

## 📌 Key Features

* Size is part of the type → `[3]int` ≠ `[4]int`
* Elements stored in **contiguous memory**
* Access using index → `a[0]`
* Default values = zero values (0, "", false)

## 📌 Important Points

* Arrays are **value types** → copying creates a full copy
* Can be compared using `==` (if elements are comparable)
* Rarely used directly in Go (slices preferred)

## 📌 Example

```go
a := [3]int{1, 2, 3}
fmt.Println(a[1]) // 2
```

---

# 🔹 3. Slices

## 📌 Definition

A slice is a **dynamic, flexible view of an array**.

```go
var s []int
```

## 📌 Internal Structure (Very Important)

A slice has 3 components:

1. Pointer → to underlying array
2. Length → number of elements
3. Capacity → max elements it can hold ([notes.shichao.io][2])

---

## 📌 Creating Slices

```go
s := []int{1, 2, 3}
s := make([]int, 5)       // length = 5
s := make([]int, 5, 10)   // length = 5, capacity = 10
```

---

## 📌 Slice Operations

### 1. Slicing

```go
s := a[1:3] // from index 1 to 2
```

### 2. Append

```go
s = append(s, 4)
```

### 3. Copy

```go
copy(dest, src)
```

---

## 📌 Important Concepts

* Multiple slices can share same underlying array
* Changing one slice may affect another
* Slice is **reference-like**, but technically a value type

---

## 📌 Nil Slice

```go
var s []int // nil slice
```

---

## 📌 Comparison

* Cannot compare slices directly (except with `nil`)

---

# 🔹 4. Maps

## 📌 Definition

A map is a **collection of key-value pairs**.

```go
m := map[string]int{}
```

---

## 📌 Creating Maps

```go
m := make(map[string]int)
m := map[string]int{"a": 1, "b": 2}
```

---

## 📌 Operations

### Insert / Update

```go
m["key"] = 10
```

### Access

```go
value := m["key"]
```

### Delete

```go
delete(m, "key")
```

### Check existence

```go
v, ok := m["key"]
```

---

## 📌 Important Features

* Keys must be **comparable types**
* Values can be any type
* Accessing missing key → returns zero value
* Iteration order is **random**

---

## 📌 Example

```go
counts := make(map[string]int)
counts["go"]++
```

---

# 🔹 5. Structs

## 📌 Definition

A struct is a **collection of named fields (can be different types)**.

```go
type Employee struct {
    ID   int
    Name string
}
```

---

## 📌 Features

* Heterogeneous (different field types)
* Fixed size
* Fields accessed using dot notation

```go
e.Name = "John"
```

---

## 📌 Initialization

### 1. Zero value

```go
var e Employee
```

### 2. Literal

```go
e := Employee{ID: 1, Name: "John"}
```

---

## 📌 Struct Comparison

* Comparable if all fields are comparable

---

## 📌 Embedded Structs (Composition)

```go
type Manager struct {
    Employee
    Level int
}
```

---

# 🔹 6. JSON Handling

## 📌 What is JSON?

* Standard format for **data exchange**
* Uses key-value pairs and arrays ([notes.shichao.io][2])

---

## 📌 Go Support

Package: `encoding/json`

---

## 📌 Encoding (Go → JSON)

```go
data, _ := json.Marshal(v)
```

---

## 📌 Decoding (JSON → Go)

```go
json.Unmarshal(data, &v)
```

---

## 📌 Struct Tags (Very Important)

```go
type Person struct {
    Name string `json:"name"`
}
```

---

## 📌 Mapping

* JSON object → struct / map
* JSON array → slice
* JSON numbers, strings, bools → Go equivalents

---

# 🔹 7. Text and HTML Templates

## 📌 Purpose

Used to **generate formatted output dynamically** (like web pages or reports).

---

## 📌 Packages

* `text/template`
* `html/template`

---

## 📌 Basic Example

```go
tmpl, _ := template.New("test").Parse("Hello {{.Name}}")
tmpl.Execute(os.Stdout, data)
```

---

## 📌 Key Features

* Data-driven output
* Safe HTML escaping (in `html/template`)
* Used in web applications

---

# 🔹 8. Comparison Summary

| Type   | Size    | Data Type | Use Case              |
| ------ | ------- | --------- | --------------------- |
| Array  | Fixed   | Same      | Low-level, fixed data |
| Slice  | Dynamic | Same      | Most common (lists)   |
| Map    | Dynamic | Key-value | Fast lookup           |
| Struct | Fixed   | Mixed     | Modeling objects      |

---

# 🔹 9. Key Takeaways (Very Important)

* Arrays → fixed, rarely used
* Slices → most important, flexible lists
* Maps → fast key-value storage
* Structs → custom data models
* JSON → data exchange format
* Templates → dynamic output generation

