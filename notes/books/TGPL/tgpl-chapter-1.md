# TGPL Chapter 1 — A Tutorial (Detailed Notes)

---

## 1. Purpose of Chapter 1

Chapter 1 is a **quick tour of the Go language**, not a deep explanation. It introduces many concepts through small programs so you can **start writing useful programs immediately**. ([notes.shichao.io][1])

* Focus: learning by examples
* Covers real tasks: CLI tools, file processing, web requests, images
* Important: You are **not expected to understand everything fully yet**

👉 Think of this chapter as:

> “Exposure to Go, not mastery”

---

## 2. Structure of a Go Program

Basic Go program:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world")
}
```

### Explanation

* **package main** → entry point of program
* **import** → includes libraries
* **func main()** → execution starts here
* **fmt.Println()** → prints output

👉 Go programs are **compiled into machine code**, producing fast executables ([notes.shichao.io][1])

---

## 3. Key Concepts Introduced

---

### 3.1 Command-line Arguments (`os.Args`)

* A **slice of strings**
* Stores input passed from terminal

Example:

```bash
go run main.go hello world
```

```go
fmt.Println(os.Args[1]) // hello
```

✔ Key ideas:

* `os.Args[0]` → program name
* `os.Args[1:]` → user inputs
* Used in CLI tools

---

### 3.2 Looping and Iteration

Go uses **only one loop → `for`**

```go
for i := 0; i < len(os.Args); i++ {
    fmt.Println(os.Args[i])
}
```

Or better:

```go
for index, value := range os.Args {
    fmt.Println(index, value)
}
```

✔ `range` simplifies iteration

---

### 3.3 Strings and Formatting

* Strings can be concatenated using `+`
* Efficient way: `strings.Join()`

```go
fmt.Println(strings.Join(os.Args[1:], " "))
```

---

### 3.4 Maps (Key-Value Data Structure)

Used to count duplicates:

```go
counts := make(map[string]int)
counts[input]++
```

✔ Important ideas:

* Dynamic size
* Keys → unique
* Values → counts/data

---

### 3.5 Input Handling (`bufio.Scanner`)

Used for reading input line-by-line:

```go
scanner := bufio.NewScanner(os.Stdin)

for scanner.Scan() {
    line := scanner.Text()
}
```

✔ Features:

* Reads efficiently
* Handles large input streams
* Used in file reading and stdin

---

### 3.6 File Handling

```go
file, err := os.Open(filename)
```

✔ Concepts:

* Error handling using `if err != nil`
* Files treated like input streams

---

### 3.7 HTTP Requests (`http.Get`)

Fetch data from web:

```go
resp, err := http.Get(url)
```

✔ Concepts:

* Networking made simple
* Response contains webpage data
* Used for web clients

---

### 3.8 Concurrency (Introduced Lightly)

Go supports concurrency using **goroutines**:

```go
go fetch(url)
```

✔ Important idea:

* Multiple tasks run simultaneously
* Designed for modern multi-core systems ([liveBook][2])

---

### 3.9 Named Types

Custom types:

```go
type Celsius float64
```

✔ Benefits:

* Better readability
* Attach methods
* Strong type safety

---

## 4. Major Example Programs in Chapter 1

---

### 4.1 Echo Programs (echo1, echo2, echo3)

Purpose:

* Print command-line arguments

Concepts:

* `os.Args`
* loops
* string joining

---

### 4.2 Duplicate Line Counter (dup1, dup2, dup3)

Purpose:

* Count repeated lines

Concepts:

* maps
* file input
* `bufio.Scanner`

---

### 4.3 Lissajous (GIF Generator)

Purpose:

* Generate animated GIF

Concepts:

* loops
* math functions
* image creation

⚠️ Important:

* Math part is complex
* Focus on Go structure, not math

---

### 4.4 Fetch Program

Purpose:

* Download web content

Concepts:

* `http.Get`
* concurrency (parallel fetch)

---

### 4.5 Web Server Example

```go
http.HandleFunc("/", handler)
http.ListenAndServe("localhost:8000", nil)
```

Purpose:

* Start a web server

Concepts:

* HTTP handling
* request/response

---

## 5. Important Go Features Highlighted

---

### 5.1 Simplicity

* Clean syntax
* Minimal keywords

---

### 5.2 Fast Compilation

* Much faster than C/C++
* Produces single executable

---

### 5.3 Built-in Concurrency

* Goroutines
* Channels (introduced later)

---

### 5.4 Strong Typing

* Prevents errors
* Safer programs

---

### 5.5 Standard Library

Go provides powerful built-in packages:

* `fmt` → printing
* `os` → system interaction
* `bufio` → input
* `net/http` → networking

---

## 6. Error Handling in Go

Go uses **explicit error handling**:

```go
if err != nil {
    fmt.Println(err)
}
```

✔ No exceptions like Java/Python
✔ Errors are values


## 7. Common Confusions (Very Important)

---

### 1. Why so many concepts at once?

→ To give a **big picture of Go**

---

### 2. Why is Lissajous example hard?

→ Math is complex, not Go

---

### 3. Why no deep explanation?

→ Later chapters explain each topic in detail

---

## 8. Summary of Chapter 1

---

Chapter 1 teaches you:

* Basic Go program structure
* Command-line tools
* File and input handling
* Maps and data processing
* HTTP and networking
* Introduction to concurrency
* Real-world examples

👉 Overall:

> Chapter 1 = “Learn Go by seeing it in action”
