# 🔢 Fibonacci CLI Generator

A simple, robust command-line interface (CLI) application built in Go to generate the Fibonacci sequence up to a specified number of terms.

This project was developed to practice fundamental Go concepts, including explicit error handling, table-driven testing, command-line argument parsing, and containerization with minimal base images.

## ✨ Features

* **Fast & Efficient:** Implemented in Go, compiled to a single, highly performant binary.
* **Robust Error Handling:** Checks for invalid inputs (e.g., non-integers, negative numbers).
* **Overflow Protection:** Includes explicit checks to prevent integer overflow for large input values, ensuring the application fails safely.
* **Containerized:** Utilizes a multi-stage Docker build to create a tiny, secure image based on `scratch`.

## 🚀 Getting Started

These instructions will get you a copy of the project up and running on your local machine.

### Prerequisites

* Go (version 1.21 or later)
* Docker or Podman (for containerized execution)

### 1. Build the Application

Navigate to the project directory and build the executable binary.

```bash
# 1. Initialize the Go module (if not already done)
go mod init fibonacci

# 2. Build the binary and name it 'fibonacci'
go build -o fibonacci
```

### 2. Run Locally
Execute the built application directly from your command line, passing the desired number of results as an argument.

```bash
./fibonacci 12
```
Output:

```bash
Fibonacci Sequence (12 results):
0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89
```

### Testing
```bash
manish.sharma2@M-L5061RWGK4 fibonacci % go test -v
=== RUN   TestFibonacci
--- PASS: TestFibonacci (0.00s)
PASS
ok      fibonacci       1.313s
```
