# TGPL Chapters 2–3 — Program Structure & Data Types

---

## Chapter 2: Program Structure

### 2.1 Names

* In Go, **names (identifiers)** are used for variables, functions, types, constants, and packages.

* A name must start with a letter or underscore (`_`), followed by letters, digits, or underscores.

* Go is **case-sensitive**, so `count` and `Count` are different.

* **Export rule (very important):**

  * Names starting with **uppercase** → exported (visible outside package)
  * Names starting with **lowercase** → unexported (package-private)

* Special names:

  * `_` (blank identifier): used to ignore values

**Key idea:** Names are simple but crucial because Go uses **capitalization for visibility instead of keywords like public/private.**

---

### 2.2 Declarations

* A **declaration** introduces a name and specifies its type and sometimes its value.

* There are 4 main kinds:

  * `var` → variables
  * `const` → constants
  * `type` → types
  * `func` → functions

Example:

```go
var x int
const Pi = 3.14
type Age int
func add(a int, b int) int { return a + b }
```

**Key idea:** Declarations define the **structure of a program before execution begins.**

---

### 2.3 Variables

* A **variable** stores a value of a specific type.

#### Ways to declare:

1. Explicit type:

```go
var x int = 10
```

2. Type inference:

```go
var x = 10
```

3. Short declaration (inside functions only):

```go
x := 10
```

* **Zero value concept (important in Go):**

  * If not initialized, variables get default values:

    * `int → 0`
    * `float → 0.0`
    * `bool → false`
    * `string → ""`

* Multiple variables:

```go
var a, b, c int
```

**Key idea:** Go avoids uninitialized garbage values — everything has a predictable default.

---

### 2.4 Assignments

* Assignment updates the value of a variable.

```go
x = 5
```

* Multiple assignment:

```go
a, b = b, a   // swap
```

* Assignment is **value-based**:

  * Copies the value (except for reference-like types such as slices, maps, pointers)

* Compound assignment:

```go
x += 2
```

* Increment/Decrement:

```go
x++
x--
```

(Note: these are statements, not expressions)

**Key idea:** Assignment in Go is simple and explicit, with no hidden conversions.

---

### 2.5 Type Declarations

* You can define **new named types**:

```go
type Celsius float64
type Fahrenheit float64
```

* Even if underlying type is same, they are **distinct types**:

```go
var c Celsius = 10
var f Fahrenheit = 10
```

* Requires explicit conversion:

```go
f = Fahrenheit(c)
```

**Why important?**

* Improves **type safety**
* Prevents mixing logically different values

**Key idea:** Type declarations create **stronger abstractions**, not just aliases.

---

### 2.6 Packages and Files

* A **package** is a collection of Go files in the same directory.

* Each file starts with:

```go
package main
```

* Rules:

  * All files in a directory → same package
  * One special package: `main` (entry point)

* Importing packages:

```go
import "fmt"
```

* Package structure enables:

  * Code reuse
  * Modularity
  * Encapsulation

**Key idea:** Packages are the **core unit of modularity in Go.**

---

### 2.7 Scope

* **Scope** determines where a name is accessible.

Types of scope:

1. **Local scope** → inside functions
2. **Package scope** → declared outside functions
3. **File scope** → imports and top-level declarations
4. **Block scope** → inside `{}`

* Inner scopes can **shadow** outer variables:

```go
x := 10
if true {
    x := 20  // different x
}
```

**Key idea:** Scope controls **visibility and lifetime of names**, preventing conflicts.

---

# Chapter 3: Basic Data Types

---

### 3.1 Integers

* Integer types:

  * Signed: `int8, int16, int32, int64`
  * Unsigned: `uint8, uint16, uint32, uint64`
  * Platform-dependent: `int`, `uint`

* Special:

  * `byte` = `uint8`
  * `rune` = `int32` (Unicode code point)

* Supports:

  * Arithmetic: `+ - * /`
  * Bitwise: `& | ^ << >>`

Example:

```go
var x int = 10
var y uint = 20
```

**Key idea:** Go enforces **explicit type conversions**, avoiding implicit mixing.

---

### 3.2 Floating-Point Numbers

* Types:

  * `float32`
  * `float64` (default)

* Based on IEEE 754 standard

Example:

```go
var pi float64 = 3.14159
```

* Scientific notation:

```go
1.23e4
```

**Key idea:** Floating-point values are approximate, not exact.

---

### 3.3 Complex Numbers

* Types:

  * `complex64`
  * `complex128`

Example:

```go
z := complex(2, 3)  // 2 + 3i
```

* Real and imaginary parts:

```go
real(z)
imag(z)
```

**Key idea:** Built-in support makes Go useful for scientific/math computations.

---

### 3.4 Booleans

* Type: `bool`
* Values:

  * `true`
  * `false`

Example:

```go
var flag bool = true
```

* Used in conditions:

```go
if flag {
    // do something
}
```

**Key idea:** No implicit conversion (e.g., 0 ≠ false)

---

### 3.5 Strings

* Immutable sequence of bytes

Example:

```go
s := "hello"
```

* Operations:

  * Concatenation:

    ```go
    s = s + " world"
    ```
  * Length:

    ```go
    len(s)
    ```

* UTF-8 encoded

* Indexing gives **bytes**, not characters

**Key idea:** Strings are immutable and UTF-8 aware, but byte-based internally.

---

### 3.6 Constants

* Declared using `const`

```go
const Pi = 3.14
```

* Must be known at compile time

* Can be **untyped constants**:

```go
const x = 10
```

* Typed constants:

```go
const y int = 10
```

* `iota` (important concept):

```go
const (
    A = iota
    B
    C
)
```

Values:

```
A = 0
B = 1
C = 2
```

**Key idea:** Constants improve **safety and clarity**, and `iota` helps generate sequences.

---

## Examples I Typed

```go
package main

import "fmt"

func main() {
    var x int = 10
    y := 20

    const Pi = 3.14

    z := complex(2, 3)

    fmt.Println(x, y, Pi, z)
}
```

```go
type Celsius float64

func main() {
    var c Celsius = 25
    fmt.Println(c)
}
```