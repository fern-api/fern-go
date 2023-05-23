# Roadmap

The following outlines a roadmap for the development of the Go generator.

## Model

- [x] Setup generator (i.e. `fern-go` binary)
- [x] Generate basic types (e.g. types, enums, and built-in types)
- [] Generate unions, visitors, unmarshalers, etc.
- [] Generate the IR
- [] Replace the manually-written IR with the generated IR
- [] Verify with round-trip tests (i.e. deserialize and re-serialize `ir.json` and verify its equivalent)
- [] Design and implement better generator output (i.e. generate into separate Go packages, files, etc)
- [] Generate a basic `go.mod`, `go.sum`, etc (similar to Fern's Typescript generator's `package.json`)
- [] Add integration tests (i.e. with hard-coded `ir.json` files and fixtures)
- [] API review

## Client

- [] Design client API