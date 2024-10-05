## contains

**A Go package that checks if a value exists in a slice of strings, handling numeric types.**

This package provides a versatile `Contains` function that efficiently determines if a value is present within a slice of strings. It gracefully handles potential numeric values by converting the slice to the corresponding numeric type (int or float64) for comparison.

### Installation

```bash
go get -u github.com/your-username/contains
```

Replace `your-username` with your actual GitHub username or the location of your package.

### Usage

Import the `contains` package and leverage the `Contains` function:

```go
package main

import (
    "fmt"
    "github.com/your-username/contains"
)

func main() {
    stringSlice := []string{"apple", "banana", "10", "3.14"}

    fmt.Println(contains.MustContains(stringSlice, "apple"))        // true
    fmt.Println(contains.MustContains(stringSlice, "orange"))      // false
    fmt.Println(contains.MustContains(stringSlice, 10))             // true (converts "10" to int)
    fmt.Println(contains.MustContains(stringSlice, 3.14159))        // false (converts to float64 with slight precision loss)
}
```

**Important Note:** Precision loss might occur when converting strings representing floating-point numbers due to the inherent limitations of data types.

### Functions

* **`Contains(s []string, val any) (bool, error)`**
  - This is the primary function, taking a slice of strings (`s`) and a value (`val`) of any type.
  - It switches on the type of `val` to perform type-specific checks:
    - String: Calls `containsString` for direct comparison.
    - Int: Converts the string slice to an int slice and calls `containsInt`.
    - Float64: Converts the string slice to a float64 slice and calls `containsFloat64`.
    - Any other type: Panics with a log message indicating unsupported type.
* **`MustContains(s []string, val any) bool`**
  - This is similar to Contains, just Panics if an error occurs.
