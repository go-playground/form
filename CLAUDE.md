# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

`github.com/go-playground/form/v4` — a Go library that decodes `url.Values` into Go structs and encodes Go structs back into `url.Values`. Designed for HTML form parsing. Zero non-test dependencies. Minimum Go 1.21.

## Commands

```bash
make              # lint + test + bench
make lint         # golangci-lint run --timeout 5m
make test         # go test -covermode=atomic -race ./...
make bench        # go test -run=NONE -bench=. -benchmem ./...

# Single test
go test -run TestDecoderNativeTime -race -v ./...
```

## Architecture

The library has symmetrical Decode and Encode pipelines sharing a reflection cache and configuration layer.

### Core Pipeline

1. **`form_decoder.go` / `form_encoder.go`** — Public `Decoder`/`Encoder` types, configuration methods, and entry points (`Decode`/`Encode`). These own a `sync.Pool` of worker instances (`*decoder`/`*encoder`) to avoid allocations.

2. **`decoder.go` / `encoder.go`** — Worker logic. `traverseStruct()` walks cached struct metadata and calls `setFieldByType()`, a large `reflect.Kind` switch that handles all Go types including nested structs, slices, arrays, maps, pointers, and `time.Time` (RFC3339).

3. **`cache.go`** — Thread-safe struct reflection cache using `atomic.Value` with copy-on-write writes. `parseStruct()` introspects fields once per type, reads the `form` tag, and sorts results (anonymous fields last). Shared by both decoder and encoder.

4. **`form.go`** — Shared constants: `Mode` (Implicit/Explicit controls whether untagged fields are processed) and `AnonymousMode` (Embed/Separate controls whether embedded struct fields are flattened or namespaced).

5. **`util.go`** — `ExtractType()` (exported, dereferences pointers/interfaces), `parseBool()` (extended: accepts on/off/yes/no/ok), `hasValue()` (zero-value check for omitempty).

### Key Design Decisions

- **Namespace paths**: Fields resolve via dot-separated paths (`Address.Name`) with bracket notation for indices/keys (`Items[0].Price`, `Map[key]`). Prefix/suffix are configurable.
- **Lazy bracket parsing**: `parseMapData()` only parses bracketed keys when encountering a slice/array/map, avoiding overhead for simple structs.
- **Error accumulation**: Decode errors are collected in `DecodeErrors` (map of namespace→error) rather than failing fast.
- **DOS protection**: `maxArraySize` (default 10000) caps dynamically-created slices/arrays.
- **Custom type functions**: `RegisterCustomTypeFunc` lets users override decode/encode for specific types. The functions receive `[]string` (decode) or `interface{}` (encode) and return parsed values.

### Extension Points

- `RegisterCustomTypeFunc(fn, types...)` — override decode/encode per type
- `RegisterTagNameFunc(fn)` — customize how struct tags map to field names
- `SetMode()` — `ModeExplicit` requires `form` tags on all fields
- `SetAnonymousMode()` — control embedded struct flattening (encoder only)

## Testing

All tests are white-box (same package). Tests use `github.com/go-playground/assert/v2` (dot-imported). `benchmarks/` contains comparative benchmarks in a separate package. `race_test.go`/`norace_test.go` provide a build-tagged `raceEnabled` constant.

Both `Decoder` and `Encoder` are safe for concurrent use — instantiate once, share across goroutines.
