# Learning Go

This repository contains the code snippets, mini projects, and exercises that accompany my Go course.  
Each `*section` directory mirrors a module in the course, so you can jump straight to the topic you are studying and run the examples locally.

## Getting Started

- Install Go 1.24 or newer (see `go.mod`).
- Clone the repo and run examples with `go run` from the project root, for example:
  ```
  go run ./3section/2-slices
  ```
- Execute all tests with:
  ```
  go test ./...
  ```

Some folders (notably `12section`) rely on SQLite (`github.com/mattn/go-sqlite3`). Building those examples requires a C toolchain. Networking examples in `16section` call `https://dummyjson.com/`, so they need outbound internet access unless you stub the client.

## Course Outline

| Directory | Focus |
|-----------|-------|
| `1section` | Language basics: variables, constants, iota-style enums, and the first CLI project. |
| `2section` | Control flow with `for`, `if/else`, and `switch`, plus a small challenge project. |
| `3section` | Collections and pointers: arrays, slices, maps, slice tricks, and a capstone exercise. |
| `4section` | Functions and error handling: multi-return, variadics, custom errors, `defer`, `panic`, and `recover`. |
| `5section` | Working with types: structs, methods, interfaces, custom `Stringer`s, generics, and a section project. |
| `6section` | Composition and embedding patterns that model richer domain types. |
| `7section` | Strings: Unicode, formatting, regex workflows, and Go templates. |
| `9section` | Concurrency fundamentals: goroutines, wait groups, channels (buffered/unbuffered), closing patterns, a downloader project, and race detection. |
| `10section` | File system utilities: reading/writing files, walking directories, temporary files, and the `embed` package. |
| `11section` | Encoding and decoding: JSON marshalling/unmarshalling, streaming with encoders/decoders, and Base64 helpers. |
| `12section` | Working with databases: connecting to SQLite, running queries, prepared statements, transactions, and a repository pattern. Sample databases live alongside the exercises. |
| `14section` | Testing patterns: capturing stdout, HTTP handler tests, and using `stretchr/testify` assertions. |
| `15section` | Time utilities: parsing/formatting, timers, tickers, randomness, scheduling, and time zone conversions. |
| `16section` | HTTP clients and testability: fetching remote data, refactoring for dependency injection, manual mocks, table-driven tests, and `testify/mock`. |

SQLite sample databases (`data.db`, `users_database.db`) are included for convenience.

## Working Through the Material

1. Pick the section that matches the lesson you just watched.
2. Run the example programs (`go run ./path/to/example`) to observe the behaviour.
3. Open the accompanying tests (`*_test.go`) to see idiomatic testing approaches.
4. Experiment by extending the snippets. Each directory is self-contained, so you can iterate freely.

Happy hacking!
