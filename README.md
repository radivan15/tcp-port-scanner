# TCP Port Scanner

A simple CLI tool to scan TCP ports on a host — stdlib-first Go, external dependencies only when there's a concrete reason. A Go learning portfolio project for SRE/DevOps engineers who want to understand network programming in Go.

> Status: in progress

> **Disclaimer:** only scan hosts you own or have explicit permission to scan (e.g. `localhost`, or `scanme.nmap.org` for testing). Unauthorized port scanning may be illegal.

---

## What it does

Dials each port in a given range and classifies it as `open`, `closed`, or `filtered` (same model as `nmap`). Starts as a sequential single-port scanner using `net.DialTimeout`, then grows into a concurrent scanner using a goroutine worker pool (`-workers` flag) with a configurable per-port timeout (`-timeout`, via `net.Dialer.DialContext`) — finishing with a summary table (open-port count, scan duration). Ctrl+C cancellation with partial results is planned as an optional extra.

---

## Why I'm building it

As an SRE, I've used tools like `nmap` regularly for network debugging and incident response. I wanted to understand what happens under the hood — how concurrent scanning actually works and how Go handles thousands of TCP connections efficiently using goroutines and channels.

---

## Stack

| Component | Technology | Notes |
|---|---|---|
| Language | Go 1.26+ | Stdlib-first, dependencies only when justified |
| CLI flags | `flag` | Host, port range, `-workers`, `-timeout` |
| TCP dialing | `net` | Portable open/closed/filtered check via `net.Error.Timeout()` |
| Concurrency | goroutines, channels, `context` | Worker pool, per-port timeout |
| Output & logging | `text/tabwriter`, `log/slog` | Summary table, structured logs |
| Testing | `testing` + `net.Listener` | Table-driven, no dependency on external hosts |
| Tooling | `golangci-lint`, `govulncheck`, `gosec`, Justfile | `just check` before every push |

---

## Folder Structure

not available yet.

---

## Installation

not available yet.

---

## How to Run

not available yet.

---

## Roadmap

See [ROADMAP.md](ROADMAP.md) for the phase-by-phase build plan and backlog ideas.

---

## License

MIT — see [LICENSE](LICENSE).
