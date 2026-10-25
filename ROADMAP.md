# Roadmap — TCP Port Scanner

My execution plan for this project, broken into phases. I'll check items off as I go.

## Phase 1 — Small Project

Everything lives in a single `main.go`.

- [ ] Set up project (`go mod init`) + hello world → verify: `go.mod` created, `go run .` prints output
- [ ] Check one hardcoded port with `net.DialTimeout` (plain `net.Dial` has no timeout and can hang on a firewalled port) → verify: start a local listener with `python -m http.server 8000`, port 8000 reports `open`
- [ ] Accept `-host` and `-port` via the `flag` package → verify: `go run . -host localhost -port 8000` reports `open`
- [ ] Extract a `scanPort(host string, port int, timeout time.Duration) string` function that classifies `open` / `closed` / `filtered` using `errors.As` + `net.Error.Timeout()` (try `errors.Is(err, syscall.ECONNREFUSED)` first and see why it fails on Windows — document it in a README Troubleshooting section) → verify: port 8000 shows `open`, an unused local port shows `closed`, `192.0.2.1` (RFC 5737 documentation IP, unroutable) shows `filtered`; same code builds for Linux (`$env:GOOS="linux"; go build`)
- [ ] Sequential port range scan via `-start` / `-end` flags, with input validation (1–65535, start ≤ end) → verify: prints open ports in order, invalid input exits with a clear error
- [ ] Replace ad-hoc `fmt.Println` with basic `log/slog` (`slog.NewTextHandler`) → verify: each port check shows up as a structured log line

## Phase 2 — Medium Project

Safety net first (package, tests, tooling), then features. From step 3 on, every feature must pass `just check` before push.

- [ ] Move scanning logic into `internal/scanner` (packages + exported names) → verify: `go run .` behaves exactly as before
- [ ] Table-driven tests against a local `net.Listener` (no dependency on `scanme.nmap.org`) — open listener → `open`, closed listener → `closed` → verify: `go test ./...` passes
- [ ] Tooling: `.golangci.yml` (v2 format), `govulncheck`, `gosec`, and a `Justfile` (`build`, `test`, `lint`, `format`, `security`, `check`, `run`, `clean`) wrapping `go mod tidy`, `go vet ./...`, `golangci-lint run`, `goimports -w .`, `go test -race ./...`, `govulncheck ./...` → verify: `just check` all green
- [ ] Concurrent scan with one goroutine per port (`sync.WaitGroup.Go()`), collect results safely and sort them by port → verify: `go test -race ./...` clean, output order matches the sequential version
- [ ] Replace `net.DialTimeout` with `net.Dialer.DialContext` + `context.WithTimeout`, configurable via `-timeout` flag → verify: scanning `192.0.2.1` doesn't hang past the configured timeout
- [ ] Goroutine pool with channel queue + `-workers` flag, passing the same `context.Context` through — comes after goroutine-per-port on purpose: scanning all 65535 ports at once can hit socket / ephemeral port limits, same reason production uses connection pools. `errgroup.SetLimit()` is the production-grade alternative to a hand-rolled pool → verify: `go test -race ./...` clean, never more than `-workers` dials in flight
- [ ] Pretty table output with summary (open count + scan duration) using `text/tabwriter` → verify: output formatted, totals match actual results

## Phase 3 — Large Project (optional)

Ordered from easiest to hardest.

- [ ] Build with version info (`-ldflags "-X main.version=..."`), log version at startup → verify: `just build` binary logs the injected version
- [ ] Track tool dependencies with `tool` directive in `go.mod` (goimports, govulncheck, gosec) instead of a `tools.go` hack — `golangci-lint` stays a separately installed binary (see the official install docs) to avoid pulling its large dependency tree into `go.mod` → verify: `go tool govulncheck ./...` runs
- [ ] Structured JSON logging: upgrade to `slog.JSONHandler` with fields (host, port, status, duration) → verify: logs are valid JSON, ready for CloudWatch/Loki ingestion
- [ ] Graceful shutdown: cancel an in-progress scan with Ctrl+C (`signal.NotifyContext`), reusing the `context.Context` already threaded through the worker pool, still print partial results → verify: pressing Ctrl+C mid-scan stops within ~1s, no leftover goroutines, partial results shown
- [ ] Metrics observability: write scan summary (ports scanned, open count, duration) in Prometheus text exposition format to a `.prom` file for node_exporter's textfile collector, hand-written with `fmt.Fprintf` (no `client_golang` dependency, to keep this pure Go). A short-lived CLI exits before Prometheus can scrape it, so a `/metrics` endpoint doesn't fit → verify: `promtool check metrics < scan.prom` passes

## Backlog

Not committing to these — just don't want to lose the idea.

- Banner grabbing — read the first bytes after connecting to identify the service behind an open port
- Rate limiting — throttle connection attempts so scanning stays a good network citizen
- JSON output mode — machine-readable results for feeding into other tools/CI
- Precise refused detection — platform-specific files (`refused_windows.go` / `refused_other.go` with `//go:build !windows`) to separate "connection refused" from other non-timeout errors
