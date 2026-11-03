# Troubleshooting

Issues hit while building this project, written up as symptom → cause → fix. Environment: Windows 11, PowerShell.

---

## `errors.Is(err, syscall.ECONNREFUSED)` is always `false` on Windows

**Symptom:** checking for a refused connection explicitly returns `false`, even for a port with nothing listening:

```
PS> go run . -port 9999
debug refused? false
127.0.0.1:9999 closed
```

**Cause:** Windows uses Winsock error codes, not Unix ones. Extracting the raw OS error with `errors.As(err, &errno)` (`syscall.Errno`) showed `10061`:

| OS | "Connection refused" code | Constant |
|---|---|---|
| Linux | 111 | `ECONNREFUSED` |
| Windows | 10061 | `WSAECONNREFUSED` (Winsock) |

`syscall.ECONNREFUSED` holds the Unix value, so it never matches the Winsock error.

**Fix:** don't depend on OS error codes. Only check for a timeout (`errors.As` → `net.Error`, then `Timeout()`), which is portable, and treat every other dial error as `closed`. Verified the same code cross-compiles for Linux:

```powershell
$env:GOOS="linux"; go build -o bin/ .; Remove-Item Env:GOOS
```

`Remove-Item Env:GOOS` matters: if `GOOS` stays set to `linux`, the next `go run` builds a Linux binary that can't run on Windows.

---

## Closed port on localhost: instant on this machine, may take ~2s elsewhere

**Symptom (not reproduced yet):** on some Windows setups a refused connection to localhost reportedly takes ~2s instead of failing immediately, because the OS retries after the RST. On this machine `go run . -port 9999` returns `closed` instantly.

**Risk:** with a short timeout, a slow refusal could be misread as `filtered`.

**Next:** test by scanning the same closed port with a short and a long timeout once the `-timeout` flag exists.
