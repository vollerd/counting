# Go Concurrent Counter Comparison

This script compares two approaches to implementing a concurrent counter in Go: mutex-based and array-segmented.

## Usage

Run the script:

```bash
go run main.go
```

This will execute both counter implementations and print their results and execution times.

### Approaches

Mutex-based: Uses a sync.Mutex to protect a shared counter. Easier to read, consistent, and easier to implement when the number of goroutines is known.

Array-segmented: Divides the counter into segments, one per goroutine, to reduce contention.

Feel free to modify the numGoroutines and numIncrements constants to experiment with different loads.
