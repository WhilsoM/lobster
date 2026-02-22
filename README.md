# 🦞 Lobster

**Lobster** is a simple tool I built to solve a common problem: sharing passwords or secrets between devices without leaving a trace. Think of it as **AirDrop for the browser**, but it works on everything.

I wanted something faster and more private than sending links to myself in Telegram "Saved Messages" or via email.

## How it works

1. You send a secret to the backend.
2. Lobster gives you a unique UUID link.
3. You open that link on another device (phone, laptop, whatever).
4. The secret is shown and **immediately deleted** from the server's memory.

## Why Lobster?

- **RAM only:** No databases. If the server restarts, everything is gone.
- **One-time only:** Once the link is opened, it’s destroyed. No "back" button, no history.
- **Privacy first:** Best used when deployed locally or in your private network.
- **Clean code:** I used Golang with a layered architecture (Handler -> Service -> Repository) and Dependency Injection.

## Tech Stack

- **Language:** Go (Golang)
- **Router:** Gorilla Mux
- **Containerization:** Docker
- **Concurrency:** Thread-safe `sync.RWMutex` for memory management.

---

## Quick Start

### 1. Run with Docker (Recommended)

```bash
# Build the image
docker build -t lobster .

# Start the container
docker run -p 8080:8080 lobster
```

### 2. Run locally

If you have Go installed:

```bash
go run cmd/main.go
```

The server will start at http://localhost:8080.

## API Usage

### Create a secret

Send your password as JSON:

```bash
curl -X POST http://localhost:8080/api/links \
     -H "Content-Type: application/json" \
     -d '{"password": "your-secret-string-here"}'
```

Response: {"id": "your-uuid-here"}

### Get the secret

Just hit the endpoint with the ID. Remember: it only works once!

```bash
curl http://localhost:8080/api/links/YOUR_UUID_HERE
```

Response: {"password": "your-secret-string-here"}

## Project Structure

I tried to keep things organized:

- cmd/ — Entry point.
- internal/handler/ — HTTP logic.
- internal/service/ — Business logic & Interfaces (DI).
- internal/repository/ — In-memory storage.
- utils/ — JSON response helpers.

Feel free to open an issue or a PR if you have ideas on how to make this even better!
