# TCP Port Scanner

A simple CLI tool to scan TCP ports on a host. Built with pure Go — no external dependencies.

> Built as a Go learning portfolio project — great for SRE/DevOps engineers who want to understand network programming in Go.

---

## Installation

Make sure [Go](https://go.dev/dl/) 1.23 or later is installed.

```bash
git clone https://github.com/radivan15/tcp-port-scanner.git
cd tcp-port-scanner
go build -o port-scanner .
```

---

## Usage

### Scan a single port
```bash
./port-scanner -host scanme.nmap.org -port 80
```

### Scan a port range (concurrent, default 100 workers)
```bash
./port-scanner -host scanme.nmap.org -start 1 -end 1024
```

### Scan with custom worker count
```bash
# Faster scan
./port-scanner -host scanme.nmap.org -start 1 -end 1024 -workers 200

# More resource-efficient
./port-scanner -host scanme.nmap.org -start 1 -end 1024 -workers 50
```

### Scan with custom timeout (in milliseconds)
```bash
# 300ms timeout — suitable for local or fast networks
./port-scanner -host 192.168.1.1 -start 1 -end 1024 -timeout 300

# 2000ms timeout — suitable for slow or remote networks
./port-scanner -host scanme.nmap.org -start 1 -end 1024 -timeout 2000
```

---

## Example Output

```
Scanning scanme.nmap.org port 1-1024 (workers: 100)...
Port 22    OPEN
Port 80    OPEN

Total: 2 open out of 1024 ports
Scan time: 1.043s
```

---

## Why I Built This

As an SRE, I've used tools like `nmap` regularly for network debugging and incident response. I wanted to understand what happens under the hood — how concurrent scanning actually works and how Go handles thousands of TCP connections efficiently using goroutines and channels.

---

## What I Learned

- How TCP connections work at the code level using `net.DialTimeout`
- The difference between naive goroutines (one per task) vs a goroutine pool (fixed workers + channel queue)
- How Go channels act as a job queue — similar to a task queue in distributed systems
- Using `sync.WaitGroup` to wait for all workers to finish
- How timeout and concurrency settings affect scan speed vs resource usage

---

## Roadmap

- [ ] Stage 1 — Scan a single port via CLI flag
- [ ] Stage 2 — Sequential port range scan
- [ ] Stage 3 — Concurrent scan with goroutines (one goroutine per port)
- [ ] Stage 4 — Goroutine pool with channel queue + `-workers` flag
- [ ] Stage 5 — Configurable per-port timeout via `-timeout` flag
- [ ] Stage 6 — Pretty table output with summary (open count + scan duration)

### Planned Features

- Scan a single port on a host
- Scan a port range (e.g. 1–1024)
- Concurrent scanning with goroutines (much faster)
- Control worker count (goroutines) via `-workers` flag
- Configurable per-port timeout via `-timeout` flag (in milliseconds)
- Pretty table output with open port count and scan duration

---

## License

MIT License — see [LICENSE](LICENSE)
