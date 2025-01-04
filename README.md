# Go Race Condition Example

This repository demonstrates a common race condition in Go programs involving concurrent goroutines accessing a shared variable without proper synchronization mechanisms like mutexes or channels.

The `bug.go` file contains code that exhibits this race condition.  The `bugSolution.go` file provides a corrected version using a mutex to protect the shared `count` variable.

## Running the Code

1. Clone the repository:
   ```bash
git clone <repository_url>
```
2. Navigate to the repository directory:
   ```bash
cd <repository_name>
```
3. Run the buggy code:
   ```bash
go run bug.go
```
4. Run the corrected code:
   ```bash
go run bugSolution.go
```

Observe the differences in the output to understand how the race condition is resolved.
