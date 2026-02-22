# 🦞 Lobster

Lobster: Browser-based "AirDrop" for Secrets

Lobster is a lightweight, secure tool designed for quick, cross-platform data sharing. Think of it as AirDrop for your browser, but focused on passwords and sensitive strings.

## How it works:

You send a password/secret to the backend → Lobster generates a unique, one-time UUID link → You open that link on any other device → The secret is retrieved and instantly wiped from memory.

## Key Features:

- Ephemeral: No databases. Secrets live only in RAM and vanish forever after the first read.

- Cross-Platform: Works on any device with a browser (PC, Mac, iPhone, Android).

- Zero-Trace: Designed for local deployment to ensure maximum privacy.

Tech Stack: Go (Golang), Gorilla Mux, Docker (Multi-stage builds).

> Ephemeral secret sharing tool — "AirDrop for your browser."

### 🚀 Quick Start

1.  **Clone the repo**
2.  **Build and Run with Docker:**
    ```bash
    docker build -t lobster .
    docker run -p 8080:8080 lobster
    ```
3.  **Usage**:

- Create: POST /api/links with {"password": "your-secret"}
- Retrieve: GET /api/links/{uuid} (Link is destroyed after use)
