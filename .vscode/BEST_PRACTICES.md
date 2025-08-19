# Go Codebase Best Practices

This document outlines the coding principles and best practices followed in this Go project to ensure clean, maintainable, and idiomatic code.

---

## 1. General Principles

- Keep it **simple** and **idiomatic** — follow Go's philosophy.
- Avoid overengineering; solve the problem in the simplest way possible.
- Readability and maintainability are prioritized over clever tricks.
- Code must be **tested**, **documented**, and **reviewed**.

---

## 2. Dependency Management

- Use **Go modules** (`go.mod`, `go.sum`) to manage dependencies.
- Avoid unnecessary external libraries. Prefer the Go standard library whenever possible.
- Use [Go Proxy](https://proxy.golang.org/) for dependency versioning consistency.
- Pin versions of all dependencies — avoid `replace` unless absolutely necessary.
- Run `go mod tidy` regularly to clean up unused dependencies.

---

## 3. Project Structure & Package Design

- Organize code into well-defined **packages** by responsibility, not by layer (e.g., `user`, not `handlers`).
- Keep packages **small**, **focused**, and **independent**.
- A package should not depend on a higher-level package (i.e., follow **dependency inversion**).
- Use internal packages (`/internal`) to hide implementation details.


---

## 4. Avoiding Cyclic Dependencies

- Always depend **inward** (towards abstractions), not outward (towards concrete implementations).
- Introduce **interfaces** in higher-level packages if a lower-level package needs to call back.
- Never let package A import package B if B also imports A (directly or indirectly).
- Break cycles using:
  - Interfaces
  - Inversion of control
  - Dependency injection
  - Event/message passing

🧠 **Tip**: Run `go list -json ./... | jq '.Imports, .Deps'` or use tools like `go mod graph` or `go-callvis` to visualize dependencies.

---

## 5. Function Naming Principles

- Use **clear, descriptive names** that express intent.
- Exported functions and types must be **capitalized** and have **GoDoc comments**.
- Prefer `verbNoun()` format for actions (e.g., `CreateUser`, `SendEmail`).
- Keep function names short **only if** their context is obvious.
- Avoid stutter (e.g., don't use `user.UserService`; use `user.Service`).

| Example          | ✅ Good        | 🚫 Bad           |
|------------------|---------------|------------------|
| CreateUser       | `CreateUser`  | `MakeNewUserFn`  |
| Start            | `Start`       | `BeginStartProc` |
| getUserByID      | `GetByID`     | `get_user_by_id` |

---

## 6. Interfaces

- Define **interfaces in the consumer package**, not the implementation.
- Keep interfaces **small** (usually 1-3 methods).
- Prefer composition over large, all-in-one interfaces.
- Use naming conventions:
  - `Reader`, `Writer`, `Store`, `Service`, `Handler`, etc.

---

## 7. Error Handling

- Handle errors explicitly. No silent failures.
- Wrap errors with context using `fmt.Errorf("context: %w", err)`.
- Avoid panics unless in truly unrecoverable situations.
- Use `errors.Is()` and `errors.As()` for error unwrapping.

---

## 8. Testing

- Use the standard `testing` package.
- Name tests clearly: `TestFunctionName_CaseDescription`
- Use table-driven tests for variations.
- Mock dependencies using interfaces — avoid real DBs or networks in unit tests.
- Place test files next to the source files: `example.go` → `example_test.go`

---

## 9. Linting & Formatting

- Enforce `gofmt`, `goimports`, and `golangci-lint`.
- Integrate formatters and linters into CI/CD.
- Prefer consistent formatting and idiomatic Go patterns.

---

## 10. Tooling

Recommended tools:

- `gofmt` – Formatting
- `goimports` – Import management
- `golangci-lint` – Linting
- `go test -cover` – Test coverage
- `go mod tidy` – Dependency cleanup
- `go-callvis` – Call graph visualization

---

## References

- [Effective Go](https://golang.org/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Go Proverbs by Rob Pike](https://go-proverbs.github.io/)
- [Uber Go Style Guide](https://github.com/uber-go/guide)

---

## Contribution

To propose updates to this guide, open a PR with your changes and rationale.
