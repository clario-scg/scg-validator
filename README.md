# scg-validator

A Laravel-inspired validation library for Go. It aggregates all field errors by default (like Laravel), supports bail to stop on first failure per field, and offers a familiar, declarative rule syntax.

- Aggregates errors per field (no panics). You get a Result with all errors or an error value you can return.
- Laravel-like bail behavior: add `bail` to stop validating a field after the first failed rule.
- Rich rule set: required, email, min/max, numeric, string, dates, files, and more.
- Confirmed rule: `<field>` must match `<field>_confirmation`.
- Custom messages and attribute names via simple setters.

## Install

```bash
go get github.com/next-trace/scg-validator
```

## Quickstart

Validate a simple payload using Laravel-style rule strings.

```go
package main

import (
	"fmt"

	"github.com/next-trace/scg-validator/validator"
)

func main() {
	v := validator.New()

	data := map[string]any{
		"name":  "John Doe",
		"email": "john@example.com",
		"age":   25,
	}

	rules := map[string]string{
		"name":  "required|min:2|max:50",
		"email": "required|email",
		"age":   "required|integer|min:18|max:120",
	}

	// Option A: get a detailed result with all errors aggregated
	res := v.ValidateWithResult(data, rules)
	if !res.IsValid() {
		fmt.Println("validation failed:")
		for field, errs := range res.Errors() {
			for _, msg := range errs {
				fmt.Printf("- %s: %s\n", field, msg)
			}
		}
		return
	}

	// Option B: use the simple error API
	if err := v.Validate(data, rules); err != nil {
		fmt.Println("validation error:", err.Error())
		return
	}

	fmt.Println("validation passed")
}
```

Notes:
- min/max apply to numbers or string lengths depending on the value type.
- The result API gives you full access to all errors per field.

## Special Behaviors

- Bail
  - Add `bail` to a field’s rule string to stop validating that field after the first failed rule.
  - Example:
    ```go
    v := validator.New()
    data := map[string]any{"email": "not-an-email"}
    rules := map[string]string{
    	// stops at the first failure for email
    	"email": "bail|required|email|min:10",
    }
    res := v.ValidateWithResult(data, rules)
    _ = res // res.Errors()["email"] will contain only the first error
    ```

- Confirmed
  - Validates that `<field>` equals `<field>_confirmation`.
  - Example:
    ```go
    v := validator.New()
    data := map[string]any{
    	"password":              "secret",
    	"password_confirmation": "secret",
    }
    rules := map[string]string{
    	"password": "required|confirmed",
    }
    if err := v.Validate(data, rules); err != nil {
    	fmt.Println("password not confirmed:", err)
    }
    ```

- Custom Messages and Attributes
  - Override any rule globally:
    ```go
    v := validator.New()
    v.SetCustomMessage("required", "The :attribute field is mandatory")
    v.SetCustomMessage("email", "Please provide a valid email address")
    ```
  - Set field-specific custom message for a rule using the key "<rule>.<field>":
    ```go
    v.SetCustomMessage("email.email", "The Email field must be a valid address")
    ```
  - Customize attribute names used in messages:
    ```go
    v.SetCustomAttribute("email", "Email")
    v.SetCustomAttribute("age", "Age")
    ```

## Advanced

- Database rules (exists, unique)
  - Implement contract.PresenceVerifier and register it per table. Example:
    ```go
    package main

    import (
    	"database/sql"
    	"fmt"

    	"github.com/next-trace/scg-validator/contract"
    	"github.com/next-trace/scg-validator/registry/database"
    	_ "modernc.org/sqlite" // or your driver
    )

    type DBVerifier struct{ db *sql.DB }
    func (d *DBVerifier) Exists(table, field string, value any) (bool, error) {
    	var cnt int
    	q := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?", table, field)
    	if err := d.db.QueryRow(q, value).Scan(&cnt); err != nil { return false, err }
    	return cnt > 0, nil
    }
    func (d *DBVerifier) Unique(table, field string, value any) (bool, error) {
    	exists, err := d.Exists(table, field, value)
    	return !exists, err
    }

    func main() {
    	db, _ := sql.Open("sqlite", ":memory:")
    	// create schema and seed...
    	database.RegisterPresenceVerifier("users", &DBVerifier{db: db})

    	v := validator.New()
    	data := map[string]any{"email": "john@example.com"}
    	rules := map[string]string{"email": "required|email|unique:users,email"}
    	_ = v.Validate(data, rules)
    }
    ```

- File rules (file, image, mimes)
  - Provided out of the box. Integrate with your file type detection as needed.

- Build tags
  - No special build tags are required. All rules are available by default. If you need to slim binaries, you can vendor and exclude packages at build time, but the library itself does not rely on build tags.

## Testing & Quality

[![Go Version](https://img.shields.io/badge/go-1.25%2B-00ADD8?logo=go)](https://go.dev/dl/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Run tests with race detector and coverage:

```bash
./scg test         # runs tests with -race and generates coverage.txt
./scg coverage     # opens coverage report
```

Linters, security, and formatting:

```bash
./scg lint
./scg lint-fix
./scg security
./scg format
```

## Versioning

This project follows [Semantic Versioning](https://semver.org/) (`MAJOR.MINOR.PATCH`).

- **MAJOR**: Breaking API changes
- **MINOR**: New features (backward-compatible)
- **PATCH**: Bug fixes and improvements (backward-compatible)

## License

MIT License. See the [LICENSE](LICENSE) file for details.

## Security

If you discover a security issue, please open a private issue or contact the maintainers. Avoid filing public exploits before a patch is available.

## Changelog

See Git history and releases for notable changes. Keep an eye on commit messages and PRs for rule additions and behavior changes.


## Package documentation

Each package in this repository now includes a package overview file (doc.go) to improve discoverability on pkg.go.dev and when browsing the source. Highlights:

- validator: High-level API for validating data using rule strings.
- facade: Public-facing facade over a pluggable ValidationEngine implementation.
- engine: Core execution engine that parses and evaluates rules.
- contract: Interfaces and small types to decouple components (e.g., ValidationEngine, Result, Registry).
- rules and subpackages: Built-in rule implementations organized by domain.
- registry and subpackages: Wiring for built-in rules and optional integrations like database and password helpers.
- message: Message resolver and default messages for rules and attributes.
- utils: Shared internal helpers (e.g., translation utilities).

You can view the rendered documentation via pkg.go.dev:
- https://pkg.go.dev/github.com/next-trace/scg-validator
- Or per-package pages (e.g., /validator, /facade, /engine, /contract, /rules, ...).
