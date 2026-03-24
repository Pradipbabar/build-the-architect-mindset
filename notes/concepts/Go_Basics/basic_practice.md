## 1. Packages

* A **package** is a collection of Go files.
* Every Go program starts with:

  ```go
  package main
  ```
* `main` package is required for executable programs.
* Other packages are reusable libraries.

---

## 2. Imports

* Used to include other packages.
* Example:

  ```go
  import "fmt"
  ```
* Multiple imports:

  ```go
  import (
      "fmt"
      "math"
  )
  ```

---

## 3. Exported Names

* Names starting with **capital letters are exported** (public).
* Example:

  ```go
  func Println() {} // exported
  func print() {}   // not exported
  ```

---

## 4. Functions (Multiple Return Values)

* Functions can return multiple values:

  ```go
  func divide(a, b int) (int, int) {
      return a / b, a % b
  }
  ```
* Named return values:

  ```go
  func split(sum int) (x, y int) {
      x = sum * 4 / 9
      y = sum - x
      return
  }
  ```

---

## 5. Variables (`var`, `:=`)

* Declare with `var`:

  ```go
  var x int = 10
  ```
* Type inference:

  ```go
  var x = 10
  ```
* Short declaration (`:=`) inside functions:

  ```go
  x := 10
  ```

---

## 6. Basic Types

* `int` → integers
* `float64` → decimal numbers
* `bool` → true/false
* `string` → text

Example:

```go
var a int = 10
var b float64 = 3.14
var c bool = true
var d string = "Hello"
```

---

## 7. Zero Values

* Variables without initialization get default values:

  * `int` → 0
  * `float64` → 0.0
  * `bool` → false
  * `string` → ""

---

## 8. Type Conversions

* Go does **not allow implicit conversion**.
* Must convert explicitly:

  ```go
  var i int = 10
  var f float64 = float64(i)
  ```

---

## 9. Constants

* Declared using `const`:

  ```go
  const Pi = 3.14
  ```
* Cannot be changed after declaration.

---

## 10. `for` Loops

* Only loop in Go:

  ```go
  for i := 0; i < 5; i++ {
      fmt.Println(i)
  }
  ```
* While-style:

  ```go
  i := 0
  for i < 5 {
      i++
  }
  ```
* Infinite loop:

  ```go
  for {
  }
  ```

---

## 11. `if` Statement

* Basic:

  ```go
  if x > 10 {
      fmt.Println("Big")
  }
  ```
* With short statement:

  ```go
  if x := 10; x > 5 {
      fmt.Println(x)
  }
  ```

---

## 12. `switch`

* Cleaner alternative to multiple `if`:

  ```go
  switch day {
  case "Monday":
      fmt.Println("Start")
  case "Friday":
      fmt.Println("End")
  default:
      fmt.Println("Other")
  }
  ```
* No need for `break`.

---

## 13. `defer`

* Delays execution until function ends:

  ```go
  defer fmt.Println("World")
  fmt.Println("Hello")
  ```
* Output:

  ```
  Hello
  World
  ```

---

## 14. Pointers

* Store memory addresses:

  ```go
  x := 10
  p := &x
  fmt.Println(*p) // value of x
  ```
* `&` → address
* `*` → value

---

## 15. Structs

* Custom data types:

  ```go
  type Person struct {
      Name string
      Age  int
  }
  ```
* Usage:

  ```go
  p := Person{Name: "Alice", Age: 30}
  ```

---

## 16. Arrays

* Fixed size:

  ```go
  var arr [3]int = [3]int{1, 2, 3}
  ```

---

## 17. Slices

* Dynamic arrays:

  ```go
  s := []int{1, 2, 3}
  ```
* Create with `make`:

  ```go
  s := make([]int, 5)
  ```
* Can store **any type**, including structs:

  ```go
  people := []Person{}
  ```

---

## 18. Maps

* Key-value storage:

  ```go
  m := make(map[string]int)
  m["Alice"] = 25
  ```
* Access:

  ```go
  fmt.Println(m["Alice"])
  ```

---

## 19. Function Values & Closures

* Functions are values:

  ```go
  add := func(a, b int) int {
      return a + b
  }
  ```
* Closures (capture variables):

  ```go
  func counter() func() int {
      i := 0
      return func() int {
          i++
          return i
      }
  }
  ```

---

## 🔑 Key Insight (Very Important)

* Go is **simple but strict**:

  * No implicit conversions
  * Only one loop (`for`)
  * Clear and predictable behavior
