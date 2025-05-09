# Haru Programming Language

**Haru** is a prototype programming language designed for simplicity, memory safety, and robust type checking. Built using Go and ANTLR, Haru combines Go's minimalism with Rust’s safety principles to create a practical, general-purpose language suitable for both systems-level and application-level programming.

This prototype features:
- Strong static typing with explicit mutability (`const`, `let`, `mut`)
- Functions with multi-value returns
- Fixed and dynamic arrays with bounds checking
- Safe pointer semantics (read-only / read-write)
- Structured control flow and scoped environments
- A custom interpreter written in Go, with ANTLR-based parsing
- CLI support with REPL and file execution modes

> **Note:** This is an experimental prototype built as part of an academic research project. The focus is on language design, not runtime performance.

## Requirements
* Go (go1.23.2)

## Getting Started
Build Haru interpreter:
```bash
go build -o ./haru ./cmd/harui/main.go
```

To run a Haru program:
```bash
haru run filename.haru
```

## Project Status
Actively developed as a proof-of-concept. Contributions and feedback are welcome.
