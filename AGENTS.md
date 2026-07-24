# AGENTS.md

Guidance for coding agents working in `github.com/sourcegraph/go-jsonschema`, an
**experimental** Go library for reading JSON Schema (draft-07) documents and
generating Go types from them.

## Setup

The official test suite is vendored as a git submodule. After cloning, run:

```sh
git submodule update --init --recursive
```

## Build, test, run

```sh
go run ./cmd/go-jsonschema-compiler -pkg schema -o out.go schema.json
```

The CLI accepts one or more JSON Schema files; `-pkg` sets the emitted package
name and `-o` writes to a file instead of stdout.

## Conventions

- Standard Go formatting (`gofmt`); generated output is run through
  `go/format`.
- Compiler tests are golden-file based: each `compiler/testdata/<case>/` holds a
  `schema.json` input and a `want.go` expected output. Update the corresponding
  `want.go` when intentionally changing generated code.
- CI (`.github/workflows/go-ci.yml`) checks out submodules recursively and runs
  `go test ./...`; keep that command green.
