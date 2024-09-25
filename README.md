## Matrix Multiplication

### Intro
Matrix multiplication is a fundamental operation in many scientific and engineering applications. This project showcases how to perform matrix multiplication in parallel using Go's concurrency primitives.

### Features

- Generates two random matrices of size N*M.
- Validates matrix dimensions before multiplication.
- Uses goroutines to parallelize the multiplication process.
- Uses waitgroups for managing the go routines
- Outputs the resulting matrix.

