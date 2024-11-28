[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/rese/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/rese/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/rese)](https://pkg.go.dev/github.com/yyle88/rese)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/rese/master.svg)](https://coveralls.io/github/yyle88/rese?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/rese.svg)](https://github.com/yyle88/rese/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/rese)](https://goreportcard.com/report/github.com/yyle88/rese)

# rese

**rese** simplifies Go error handling and result extraction for multi-value function calls. It combines error and result checks into a single operation.

**rese** stands for **res** (result) + **err** (error).

## CHINESE README

[中文说明](README.zh.md)

## Installation

```bash
go get github.com/yyle88/rese
```

---

## Functions

| Function                                                  | Purpose                                                                                     | Returns                                                  |
|-----------------------------------------------------------|---------------------------------------------------------------------------------------------|----------------------------------------------------------|
| `V0(err error)`                                           | Checks the error and panics if it's not `nil`. No return value.                             | None                                                     |
| `V1[T1 any](v1 T1, err error) T1`                         | Checks the error, and if no error, returns `v1`.                                            | Returns the value of type `T1`                           |
| `V2[T1, T2 any](v1 T1, v2 T2, err error) (T1, T2)`        | Checks the error, and if no error, returns `v1` and `v2`.                                   | Returns `v1` of type `T1` and `v2` of type `T2`          |
| `P0(err error)`                                           | Checks the error and panics if it's not `nil`. No return value.                             | None                                                     |
| `P1[T1 any](v1 *T1, err error) *T1`                       | Checks the error, checks that `v1` is non-`nil`, and returns `v1`.                          | Returns a pointer to `v1` of type `T1`                   |
| `P2[T1, T2 any](v1 *T1, v2 *T2, err error) (*T1, *T2)`    | Checks the error, checks that `v1` and `v2` are non-`nil`, and returns `v1` and `v2`.       | Returns pointers to `v1` and `v2` of types `T1` and `T2` |
| `C0(err error)`                                           | Checks the error and panics if it's not `nil`. No return value.                             | None                                                     |
| `C1[T1 comparable](v1 T1, err error) T1`                  | Checks the error, checks that `v1` is not a zero value, and returns `v1`.                   | Returns `v1` of type `T1`                                |
| `C2[T1, T2 comparable](v1 T1, v2 T2, err error) (T1, T2)` | Checks the error, checks that `v1` and `v2` are not zero values, and returns `v1` and `v2`. | Returns `v1` of type `T1` and `v2` of type `T2`          |

## Examples

### Example 1: Simple error and result checking with `V1`

```go
package main

import (
	"fmt"
	"github.com/yyle88/rese"
)

func run() (string, error) {
	return "Hello, World!", nil
}

func main() {
	// Using V1 to check for error and get the result
	result := rese.V1(run()) 
	fmt.Println(result) // Outputs: Hello, World!
}
```

### Example 2: Ensuring non-nil pointers with `P1`

```go
package main

import (
	"fmt"
	"github.com/yyle88/rese"
)

func getSomething() (*int64, error) {
	v := int64(42)
	return &v, nil
}

func main() {
	// Using P1 to check error and ensure non-nil pointer
	ptr := rese.P1(getSomething()) 
	fmt.Println(*ptr) // Outputs: 42
}
```

### Example 3: Zero-value checking with `C1`

```go
package main

import (
	"fmt"
	"github.com/yyle88/rese"
)

func getInt() (int, error) {
	return 20, nil
}

func main() {
	// Using C1 to check error and ensure non-zero result
	num := rese.C1(getInt())
	fmt.Println("Received:", num) // Outputs: 20
}
```


## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Support

Welcome to contribute to this project by submitting pull requests or reporting issues.

If you find this package helpful, give it a star on GitHub!

**Thank you for your support!**

**Happy Coding with `rese`!** 🎉
