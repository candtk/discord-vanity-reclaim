# Discord Vanity Reclaim

A high-performance Go tool built to help Discord server owners reclaim their own vanity URLs during server transfers — before malicious snipers can register them in the brief window the URL goes unclaimed.

---

## The Problem

When transferring a Discord vanity URL between servers, there is an unavoidable window where the URL is released and unclaimed. In that gap — often under a second — automated snipers run by bad actors can steal the vanity permanently. This tool was built to solve that: by being faster than any sniper, the owner can reclaim their own URL the moment it becomes available.

---

## Technical Highlights

### Dual claiming strategies
Two independent HTTP claim modes are implemented, selectable via config:

- **fasthttp mode** — uses the [`valyala/fasthttp`](https://github.com/valyala/fasthttp) library for low-overhead HTTP/1.1 PATCH requests. Significantly faster than Go's `net/http` due to zero allocation design and connection reuse.
- **Raw TCP socket mode** — bypasses the HTTP client entirely. Pre-establishes a pool of persistent `*tls.Conn` connections to `discord.com:443` and writes raw HTTP/1.1 requests directly over the wire, eliminating all client-side overhead from connection setup and TLS handshaking.

### TLS fingerprinting
The TLS configuration is tuned to match real Discord client behaviour — specific cipher suites and elliptic curve preferences are hardcoded to produce a fingerprint consistent with the Discord iOS mobile client, reducing the chance of the connection being flagged or deprioritised server-side.

### Goroutine amplification
A configurable `amplify` parameter spawns N goroutines per target vanity, each running its own polling loop. With multiple vanities and high amplification, hundreds of concurrent claim attempts can be in-flight simultaneously.

### Proxy rotation with ratelimit awareness
Each goroutine pulls proxies from a shared buffered channel. On a `429 Too Many Requests`, the proxy is removed from circulation and re-added after a 30-second sleep (matching Discord's ratelimit window), preventing wasted checks against a blocked proxy.

### Pre-warmed socket pool
In socket mode, TLS connections to Discord are established upfront and held open in a channel-based pool. Claim goroutines pull a connection, write the payload, read the response, and return the connection — eliminating per-request TLS handshake latency entirely.

### Webhook notifications
Discord webhook payloads are sent asynchronously for three outcomes — success, ratelimited, and failed — with time-to-claim reported in each.

---

## Tech Stack

- **Language:** Go
- **HTTP:** `valyala/fasthttp`, `net/http`, raw `crypto/tls` sockets
- **Concurrency:** goroutines, buffered channels
- **Config:** YAML via `gopkg.in/yaml.v2`
- **Logging:** `gookit/color` terminal output

---

## Project Structure

```
discord-vanity-reclaim/
├── main.go       # Entry point, goroutine orchestration, proxy/socket channel setup
├── discord.go    # VanityCheck, FastHttpClaim, ClaimUsingSocket, TLS config
├── setup.go      # Config loading (YAML), proxy/vanity list parsing, CLI prompt
├── logging.go    # Coloured terminal logging, ASCII header, clear screen
├── webhook.go    # Discord webhook payloads for success/ratelimit/failure events
└── input/
    ├── config.yaml   # Configuration file
    ├── proxies.txt   # HTTP proxies (one per line)
    └── vanities.txt  # Target vanity URLs to monitor and claim
```

---

## Setup

**1. Clone and build**
```bash
git clone https://github.com/candtk/discord-vanity-reclaim.git
cd discord-vanity-reclaim
go build -o vanity-reclaim .
```

On macOS/Linux, make the binary executable:
```bash
chmod +x ./vanity-reclaim
```

**2. Configure `input/config.yaml`**

```yaml
main:
  token: "your_discord_token"
  guildid: 123456789
  webhook: "https://discord.com/api/webhooks/..."
  amplify: 6
  usesockets: false
  socketchannels: 5
  debug: false
```

**3. Add proxies and target vanities**

`input/proxies.txt` — one HTTP proxy per line (`host:port`):
```
203.0.113.1:8080
198.51.100.4:3128
```

`input/vanities.txt` — one vanity per line:
```
myserver
```

**4. Run**
```bash
./vanity-reclaim
```

---

## Configuration

| Option | Description |
|---|---|
| `token` | Discord token with vanity-change permissions on the target guild |
| `guildid` | ID of the server you are transferring the vanity to |
| `webhook` | Discord webhook URL for claim notifications |
| `amplify` | Goroutines per vanity — higher means more parallel claim attempts |
| `usesockets` | Enable raw TCP socket claiming mode (lower latency, experimental) |
| `socketchannels` | Number of pre-warmed TLS connections in the socket pool |
| `debug` | Print per-check timing and status to stdout |

---

## License

MIT
