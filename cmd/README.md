# Fibonacci Calculator
This project calculates the result of the Fibonacci Sequence of the given *n*-th number.

## Requirements
- [Installed Golang environment](https://golang.org/doc/install)

## Usage

### Get the n-th number
1. Start the server using `go run main.go`
2. Send a request using the url `curl "localhost:8000/fib?n=[YOUR-INPUT]"`

#### **Example**

Request: `curl "localhost:8000/fib?n=50"`

Output:
```shell
$ 12586269025
```

### Run tests

1. Run `go test ./... -v` 

Output
```shell
?       github.com/yasmindias/fibonacci-service [no test files]
=== RUN   TestGet
=== RUN   TestGet/success
=== RUN   TestGet/bad_request
--- PASS: TestGet (0.00s)
    --- PASS: TestGet/success (0.00s)
    --- PASS: TestGet/bad_request (0.00s)
=== RUN   TestFibonacciRecursion
--- PASS: TestFibonacciRecursion (0.00s)
PASS
ok      github.com/yasmindias/fibonacci-service/cmd     (cached)
```
