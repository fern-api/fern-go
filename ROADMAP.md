# Roadmap

The following outlines a roadmap for the development of the Go generator.

## Model

- [x] Setup generator (i.e. `fern-go` binary)
- [x] Generate basic types (e.g. types, enums, and built-in types)
- [x] Generate unions, visitors, unmarshalers, etc
- [x] Generate undiscriminated unions
- [x] Handle object aliases (i.e. in `json.Unmarshaler`)
- [x] Generate the IR
- [x] Bootstrap the generator (i.e. Replace the manually-written IR with the generated IR)
- [x] Design and implement better generator output (i.e. generate into separate Go packages, files, etc)
- [x] Handle cross-package imports
- [x] Generate documentation for all relevant types
- [x] Add integration tests (i.e. with hard-coded `ir.json` files and fixtures)
- [x] Verify with round-trip tests (i.e. deserialize and re-serialize and verify it is equivalent)
- [x] Handle literal values in objects and [undiscriminated] unions
- [x] Add unsafe words to the Fern compiler (i.e. Go keywords).
- [x] Polish (e.g. better method receiver identifiers)
- [x] Generate a basic `go.mod`, `go.sum`, etc (similar to Fern's Typescript generator's `package.json`)
- [ ] API review

## Client

- [x] Design client API
- [x] Generate error types
- [x] Generate endpoint request types
- [x] Generate endpoint structures
- [x] Generate endpoint error decoders
- [x] Generate core utilities
- [x] Generate authorization options
- [x] Generate endpoint call method
- [x] Support path parameters (including service path parameters)
- [x] Generate client (w/ root endpoints and nested service endpoints)
- [x] Rename `Service` to `Client` for symmetry with other SDKs
- [x] Consolidate endpoint implementation in the un-exported client
- [x] Add an API error type so that all RPC errors preserve the status code
- [x] Serialize query parameters and headers
- [x] Improve structured error formatting, where possible.
- [x] Edit the file header to match Fern's other SDKs
- [x] Generate exported environments
- [x] Improve client constructor (i.e. reduce required parameters w/ options for environment and http client)
- [x] Support default environment settings (i.e. automatically set base URL based on service URL)
- [x] Introduce a better solution for using optional parameters (i.e. `stringPtr` helpers)
- [x] Generate documentation for all relevant types
- [x] Support file download and file upload RPCs
- [x] Revisit default generator output (i.e. whether to generate a `go.mod` by default or not)
- [x] Support custom error discriminiation strategy
- [ ] If a `Client` type exists, prefix the filename with a `_` to avoid conflicts
- [ ] Generate examples
- [ ] Support optional client/endpoint variables (e.g. namespace)
- [ ] Add coordinator logging for better user-facing console progress updates
- [ ] Better support for licenses (i.e. move implementation to the layer above the generator)
