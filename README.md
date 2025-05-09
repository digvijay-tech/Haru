# Haru Programming Language

**Haru** is a lightweight, statically typed, interpreted programming language designed for safety, clarity, and developer ergonomics. Inspired by the simplicity of Go and the memory safety of Rust, Haru aims to make systems-level programming more accessible without sacrificing control.

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
