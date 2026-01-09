# DevOps Go Assignment: Site Status Checker

## Objective
Build a CLI tool using Go that checks the health status of provided websites. This will teach you about HTTP requests, error handling, and command-line argument parsing in Go.

## Level 1: The Basics (Synchronous)
1. **Parse Arguments**: Read the list of websites provided as command-line arguments.
   - *Hint*: Look at the `os` package, specifically `os.Args`.
2. **Check Status**: For each website, send a GET request.
   - *Hint*: Use the `net/http` package (`http.Get`).
3. **Report**: Print the result to the console.
   - Format: `[UP] http://example.com` (for 200 OK)
   - Format: `[DOWN] http://example.com` (for errors or non-200 codes)

## Level 2: Concurrency (The "Go" Way)
Once Level 1 is working, try to make it faster!
1. **Goroutines**: Spin up a new goroutine for each website check so they happen at the same time.
2. **Channels**: Use a channel to communicate the results back to the main function to print them.

## Example Usage
```bash
go run main.go https://google.com https://github.com https://nonexistent-site.local
```

## Expected Output
```text
[UP] https://google.com
[UP] https://github.com
[DOWN] https://nonexistent-site.local
```

## Resources
- [Go by Example: Command-Line Arguments](https://gobyexample.com/command-line-arguments)
- [Go by Example: HTTP Clients](https://gobyexample.com/http-clients)
- [Go by Example: Goroutines](https://gobyexample.com/goroutines)
