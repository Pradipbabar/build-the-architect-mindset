## 1. Methods (functions with a receiver)

A **method** is just a function with a **receiver argument**.

```go
type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

👉 `(v Vertex)` is the **receiver**

### Key points

* Methods attach behavior to types
* Can be defined on **structs or custom types**
* Must be defined in the **same package as the type**

---

## 2. Methods are functions (subtle difference)

Methods are basically functions with a receiver:

```go
func Abs(v Vertex) float64  // function
func (v Vertex) Abs() float64 // method
```

### Difference

| Function          | Method                 |
| ----------------- | ---------------------- |
| Explicit argument | Receiver bound to type |
| `Abs(v)`          | `v.Abs()`              |

👉 Methods give **object-like behavior without classes**

---

## 3. Pointer receivers vs value receivers

### Value receiver

```go
func (v Vertex) Scale(f float64)
```

* Works on a **copy**
* Cannot modify original

### Pointer receiver

```go
func (v *Vertex) Scale(f float64)
```

* Works on **actual value**
* Can modify original

### When to use pointer?

* Need to modify data
* Avoid copying large structs
* Consistency across methods

👉 Pointer receivers are more common ([Zone's Note][1])

---

## 4. Pointers and interfaces

* A type implements an interface based on its **method set**
* Method sets differ:

| Type | Method set                   |
| ---- | ---------------------------- |
| `T`  | value receiver methods       |
| `*T` | both value + pointer methods |

⚠️ Important:

* If method uses pointer receiver → only `*T` satisfies interface
* `T` alone will NOT

---

## 5. Interfaces (implicit satisfaction)

```go
type Reader interface {
    Read(p []byte) (int, error)
}
```

👉 A type implements an interface **automatically** if it has required methods

✔ No `implements` keyword
✔ No explicit declaration

This is called **implicit implementation** ([Stanza][2])

---

## 6. Interface values (dynamic type & value)

An interface internally stores:

```
(value, type)
```

Example:

```go
var i interface{}
i = 42
```

* Value → `42`
* Type → `int`

👉 Method calls are dispatched using the **underlying type** ([Luke Singham][3])

---

## 7. Interface values with nil underlying values

```go
var i interface{}
var p *MyStruct = nil
i = p
```

* Interface is **NOT nil**
* Underlying value is **nil**

👉 Calling methods may still work if handled properly

✔ Go allows methods on nil receivers

---

## 8. Nil interface values

```go
var i interface{}
```

* Both type and value are nil
* Calling any method → **panic**

👉 Important distinction:

| Case           | Is nil? |
| -------------- | ------- |
| `(nil, nil)`   | YES     |
| `(nil, *Type)` | NO      |

---

## 9. Empty interface (`interface{}` / `any`)

```go
var i interface{}
```

* Can hold **any type**
* Equivalent to `any` (Go 1.18+)

👉 Used for:

* Generic containers
* Unknown types
* Reflection

---

## 10. Type assertions

Extract underlying value:

```go
v := i.(int)
```

### Safe version:

```go
v, ok := i.(int)
```

* `ok = false` if wrong type
* Prevents panic

---

## 11. Type switches

```go
switch v := i.(type) {
case int:
    fmt.Println("int")
case string:
    fmt.Println("string")
}
```

👉 Cleaner way to handle multiple types

---

## 12. Stringer interface (`fmt.Stringer`)

```go
type Stringer interface {
    String() string
}
```

If implemented:

```go
func (p Person) String() string {
    return "Name: " + p.Name
}
```

👉 Automatically used by `fmt.Println`

✔ Custom printing behavior

---

## 13. Errors (`error` interface)

```go
type error interface {
    Error() string
}
```

Example:

```go
type MyError struct{}

func (e MyError) Error() string {
    return "something went wrong"
}
```

👉 Errors are just **interfaces**

✔ Encourages custom error types

---

## 14. Readers (`io.Reader` interface)

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

Used everywhere:

* Files
* Network
* Buffers

Example flow:

```go
r := strings.NewReader("hello")
buf := make([]byte, 8)
r.Read(buf)
```

👉 Core idea: **stream of data**

✔ Standard abstraction in Go

---

## 15. Images (`image.Image` interface) *(optional)*

```go
type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) Color
}
```

👉 Represents an image generically

Used in:

* Image processing
* Graphics libraries

---

# 🧩 Final Summary (Important Concepts)

### 🔑 Core ideas

* Go has **no classes** → uses methods + interfaces
* Interfaces define **behavior, not structure**
* Implementation is **implicit**
* Interfaces store `(type, value)` internally

---

### 🚀 Most important rules to remember

1. **Method = function + receiver**
2. **Pointer receiver = modify original**
3. **Interfaces are satisfied implicitly**
4. **Interface = (value, type)**
5. **Nil interface ≠ interface holding nil**
6. **Empty interface = any type**
7. **Type assertions & switches → runtime type handling**
