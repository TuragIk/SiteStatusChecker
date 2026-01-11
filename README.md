<p align="center">
  <h1 align="center">Site Status Checker</h1>
</p>

<p align="center">
  <a href="https://go.dev/">
    <img src="https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go&logoColor=white" alt="Go">
  </a>
  <a href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License">
  </a>
</p>

<p align="center">
  <i>
    Site Status Checker is a lightweight, efficient command-line interface (CLI) tool built to monitor the health and reachability of websites. It leverages Go's standard library to perform HTTP requests and provide immediate status feedback, serving as a foundational project for understanding network operations in Go.
  </i>
</p>

<br>

## Features

* **URL Health Checks:** Validates the availability of multiple websites in a single execution.
* **Worker Pool Concurrency:** Uses a fixed pool of workers to efficiently check multiple sites in parallel.
* **Latency Reporting:** Displays the response time for each checked URL.
* **Status Code Analysis:** Distinguishes between successful responses (HTTP 200) and errors/outages.
* **CLI Arguments:** Accepts a dynamic list of target URLs via command-line arguments.
* **Robust Error Handling:** Gracefully manages network failures and invalid URLs without crashing.
* **File Input:** Supports reading URLs from a CSV file using the `-f` flag.
* **Piped Input:** Accepts URLs via standard input (stdin) for easy integration with other tools.
* **Configurable Timeout:** Custom timeout support via the `-t` flag (default: 5s).

## Tech Stack

| Component | Technology | Description |
| :--- | :--- | :--- |
| **Language** | `Go` (Golang) | Statically typed, compiled language known for efficiency. |
| **Networking** | `net/http` | Go's robust standard library for HTTP client/server implementations. |
| **Concurrency** | `sync` & `channels` | Worker pools and synchronization for efficient parallel processing. |
| **CLI & Flags** | `flag`, `os` | Standard library packages for CLI arguments and flag parsing. |
| **File I/O** | `encoding/csv`, `io` | Robust reading of files and standard input streams. |

## Quick Start

### Prerequisites

* [Go](https://go.dev/dl/) (Version 1.25 or higher recommended)

### Installation

1.  **Clone the Repository**
    ```bash
    git clone https://github.com/TuragIk/SiteStatusChecker.git
    cd SiteStatusChecker
    ```

2.  **Run the Application**
    Execute the tool by passing the URLs you wish to check as arguments:
    ```bash
    go run main.go https://google.com https://github.com
    ```

### Advanced Usage

**Read from File:**
```bash
go run main.go -f sites.csv
```

**Set Custom Timeout (seconds):**
```bash
go run main.go -t 2 https://slow-site.com
```

**Pipe from Stdin:**
```bash
cat sites.txt | go run main.go
```

**Combine Everything:**
```bash
echo "https://example.com" | go run main.go -f sites.csv -t 1 https://google.com
```

### Expected Output

```text
[UP] https://google.com (210ms)
[UP] https://github.com (95ms)
[DOWN] https://nonexistent-site.local (5ms)
```

## Project Structure

```text
SiteStatusChecker/
├── main.go            # Application Entry Point & Logic
├── go.mod             # Module Definition & Dependency Management
└── README.md          # Project Documentation
```
