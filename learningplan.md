# Go Backend Learning Plan (7 Days)

## Objective
Achieve backend-ready proficiency in Go within 7 days, with emphasis on:
- Memory model and pointers
- Concurrency (CSP model)
- Idiomatic error handling
- Testing (unit + integration)
- Backend systems (HTTP, DB, context)
- Performance and deployment

---

## Day 1 — Syntax Efficiency & Pointer Mental Model

### Core Concepts
- Variable shadowing (`:=`)
- Constants using `iota`
- Go module system (`go mod init`)
- Project structure basics (`/cmd`, `/internal`)
- Everything in Go is **pass by value**

### Pointer Deep Dive
- Pointer = copy of memory address
- Pointer vs value receiver:
  - Pointer → mutation, avoid copying large structs
  - Value → immutability, small structs

### Memory Model
- Stack vs Heap
- Escape analysis:
  ```bash
  go build -gcflags="-

# Day 2 — Composition over Inheritance

## Objective
Understand Go’s structural typing model and replace class-based thinking with composition, interfaces, and method-driven design.

---

## Structs & Methods

### Key Concepts
- Behavior is attached to types using **methods**
- Methods are functions with a **receiver**
- Receiver choice impacts:
  - mutation
  - performance
  - interface satisfaction

# Day 3 — Concurrency Model (CSP)

## Objective
Understand Go’s concurrency model based on **Communicating Sequential Processes (CSP)** and build safe, race-free concurrent systems.

---

## Goroutines

### Key Concepts
- Lightweight execution units managed by Go runtime
- Created using `go` keyword
- Not OS threads; cheaper and scalable

### Example
```go
func work() {
    fmt.Println("working")
}

func main() {
    go work()
}
````


# Day 4 — Idiomatic Error Handling & Testing

## Objective
Adopt Go’s explicit error handling model and build complete, testable logic with full branch coverage using idiomatic testing patterns.

---

## Error Handling

### Key Concepts
- No exceptions (`try/catch` does not exist)
- Errors are returned as values
- Use **early return** pattern

### Example
```go
func process(input int) (int, error) {
    if input < 0 {
        return 0, fmt.Errorf("invalid input: %d", input)
    }
    return input * 2, nil
}
````

