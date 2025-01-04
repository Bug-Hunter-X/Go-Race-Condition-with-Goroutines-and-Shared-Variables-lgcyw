```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var count int
        var mu sync.Mutex // Added mutex for synchronization
        for i := 0; i < 10; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        mu.Lock() // Acquire lock before accessing count
                        count++
                        mu.Unlock() // Release lock after accessing count
                }()
        }
        wg.Wait()
        fmt.Println(count) // Output will always be 10
}
```